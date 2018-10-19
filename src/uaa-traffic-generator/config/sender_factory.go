package config

import (
	"github.com/cloudfoundry-community/go-uaa"
	"uaa-traffic-generator/sender"
)

func NewSenders(config TrafficConfig) ([]sender.TrafficSender, *uaa.API, error) {
	api, err := uaa.NewWithPasswordCredentials(config.Credentials.Target, "", config.Credentials.ClientId, config.Credentials.ClientSecret, config.Credentials.Username, config.Credentials.UserPassword, uaa.JSONWebToken, true)
	if err != nil {
		return nil, nil, err
	}

	var senders []sender.TrafficSender
	for _, cmd := range config.UaaCommands {
		switch cmd.Cmd {
		case GetMeCmd:
			senders = append(senders, buildTrafficSender(cmd.Loop, sender.GetMeSender{}))
		case ListAllUsersCmd:
			senders = append(senders, buildTrafficSender(cmd.Loop, sender.ListAllUsersSender{}))
		}

	}

	return senders, api, nil
}

func buildTrafficSender(loop int, trafficSender sender.TrafficSender) sender.TrafficSender {
	if loop > 0 {
		return sender.Loop(loop, trafficSender)
	}
	return trafficSender
}
