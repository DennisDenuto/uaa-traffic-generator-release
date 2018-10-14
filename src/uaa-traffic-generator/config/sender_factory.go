package config

import (
	"uaa-traffic-generator/sender"
	"github.com/cloudfoundry-community/go-uaa"
)

func NewSenders(config TrafficConfig) ([]sender.Sender, *uaa.API, error) {
	api, err := uaa.NewWithPasswordCredentials(config.Credentials.Target, "", config.Credentials.ClientId, config.Credentials.ClientSecret, config.Credentials.Username, config.Credentials.UserPassword, uaa.JSONWebToken, true)
	if err != nil {
		return nil, nil, err
	}

	var senders []sender.Sender
	for _, cmd := range config.UaaCommands {
		switch cmd.Cmd {
		case GetMeCmd:
			senders = append(senders, sender.GetMeSender{})
		case ListAllUsersCmd:
			senders = append(senders, sender.ListAllUsersSender{})
		}

	}

	return senders, api, nil
}
