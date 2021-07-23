package main

// Meta struct as sub-object of ServerInfos
type Meta struct {
	Customer        string `json:"customer"`
	HostingOwner    string `json:"hosting_owner"`
	HostingProvider string `json:"hosting_provider"`
	Login           string `json:"login"`
	SSH             bool   `json:"ssh,string"`
	System          string `json:"system"`
	Update          bool   `json:"update,string"`
	PrivateIP       string `json:"private_ip,omitempty"`
	JumpServer      string `json:"jump_server,omitempty"`
}

// ServerInfos as root object structure for consul's response body
type ServerInfos struct {
	ID              string
	Node            string
	Address         string
	Datacenter      string
	TaggedAddresses string
	Meta            Meta
	CreateIndex     int
	ModifyIndex     int
}

// Configuration struct for gonfig mapping
type Configuration struct {
	JumpServer      map[string]string `env:"jumpServer"`
	DefaultUsername string            `env:"defaultUsername"`
}
