package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

type LocalNetwork struct {
	WAN_IP string `json: "WAN_IP"`
}

func main() {
	file, _ := ioutil.ReadFile("./net_vars.json")

	netVars := LocalNetwork{}
	_ = json.Unmarshal([]byte(file), &netVars)

	ip := net.ParseIP(netVars.WAN_IP)
	net.DialIP("tcp", ip)
}
