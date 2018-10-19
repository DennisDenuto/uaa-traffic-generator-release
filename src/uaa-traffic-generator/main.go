package main

import (
	"github.com/golang/go/src/pkg/flag"
	"io/ioutil"
	"log"
	"uaa-traffic-generator/config"
	"uaa-traffic-generator/runner"
)

func main() {
	configPath := flag.String("config", "", "")
	flag.Parse()

	configJsonContent, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	trafficConfig := config.NewConfig(configJsonContent)

	senders, api, err := config.NewSenders(trafficConfig)
	if err != nil {
		log.Fatal(err)
	}
	runner.RunAll(api, senders)
}
