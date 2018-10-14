package sender

import uaa_api "github.com/cloudfoundry-community/go-uaa"

//go:generate counterfeiter . TrafficSender
type TrafficSender interface {
	Send( * uaa_api.API)
}
