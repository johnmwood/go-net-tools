package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/spf13/cobra"
)

type LocalNetwork struct {
	WAN_IP string `json: "WAN_IP"`
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:     "ping",
	Aliases: []string{"p"},
	Short:   "Send a simple request to an IP and get results of endpoint",
	Long:    `I will fill this out more when I figure out what options I want to include`, // TODO: fix after adding more functionality
	Run: func(cmd *cobra.Command, args []string) {
		ping()
	},
}

func ping() {
	file, _ := ioutil.ReadFile("./net_vars.json")

	netVars := LocalNetwork{}
	_ = json.Unmarshal([]byte(file), &netVars)

	ip := net.ParseIP(netVars.WAN_IP)
	fmt.Printf("%v", ip)
	// net.DialIP("tcp", ip)
}

/*
func promptIP() {
	return
}
*/
