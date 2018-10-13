package config

import "uaa-traffic-generator/sender"

func NewSenders(config TrafficConfig) []sender.Sender {
	return []sender.Sender{
		sender.GetMeSender{},
	}
}
