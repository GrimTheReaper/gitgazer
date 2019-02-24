package main

import (
	"flag"

	"github.com/grimthereaper/gitgazer/network"

	"github.com/grimthereaper/gitgazer/github"
)

var flagPort int
var flagHost string
var flagToken string
var flagBuffer bool

func init() {
	// Using "flag" to slim down the external imports,
	//  for this test, the less imports we have the easier it will be to run..
	flag.IntVar(&flagPort, "port", 8080, "Port to bind the network to")
	flag.StringVar(&flagHost, "host", "", "Host to bind the network to")
	flag.StringVar(&flagToken, "token", "", "Github Token")
	flag.BoolVar(&flagBuffer, "buffer", true, "Buffer Github Results")
	flag.Parse()

	github.SetToken(flagToken)
	github.SetBuffering(flagBuffer)
}

func main() {
	// We don't care about the resulting server, since this app will only run one.
	_, err := network.Serve(flagHost, flagPort)
	if err != nil {
		panic(err)
	}
}
