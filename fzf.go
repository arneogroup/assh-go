package main

import (
	"fmt"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func getSelectionFromUser(consulList []ServerInfos) []int {
	idxs, _ := fuzzyfinder.FindMulti(consulList, func(i int) string {
		return fmt.Sprintf("[%s] ---- %s ---- %s", consulList[i].Node, consulList[i].Address, consulList[i].Meta.Customer)
	})
	return idxs
}
