package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getConsulEnvVar(token string) string {
	consulToken := os.Getenv(token)
	if consulToken == "" {
		log.Fatal("Token not found. Please ensure " + token + " is part of your env")
	}
	return consulToken
}

func getListFromConsul() []ServerInfos {
	// curl -s --header "X-Consul-Token: ${CONSUL_HTTP_TOKEN}" "https://${CONSUL_HTTP_ADDR}/v1/catalog/nodes"

	req, err := http.NewRequest("GET", os.ExpandEnv("https://"+getConsulEnvVar("CONSUL_HTTP_ADDR")+"/v1/catalog/nodes"), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Consul-Token", os.ExpandEnv(getConsulEnvVar("CONSUL_HTTP_TOKEN")))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result []ServerInfos
	json.Unmarshal([]byte(body), &result)
	return result
}
