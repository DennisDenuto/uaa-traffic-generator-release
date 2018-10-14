package main

import (
	"github.com/golang/go/src/pkg/flag"
	"uaa-traffic-generator/config"
	"io/ioutil"
	"uaa-traffic-generator/runner"
)

func main() {
	configPath := flag.String("config", "", "")
	flag.Parse()

	configJsonContent, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	trafficConfig := config.NewConfig(configJsonContent)

	senders, api, err := config.NewSenders(trafficConfig)
	if err != nil {
		panic(err)
	}
	runner.RunAll(api, senders)
}