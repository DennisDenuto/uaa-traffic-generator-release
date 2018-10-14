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

	senders := []sender.Sender{
		sender.GetMeSender{},
	}

	return senders, api, nil
}
