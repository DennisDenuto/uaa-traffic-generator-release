package sender

import uaa_api "github.com/cloudfoundry-community/go-uaa"

//go:generate counterfeiter . TrafficSender
type TrafficSender interface {
	Send( * uaa_api.API)
}


type TrafficSenderFunc func(*uaa_api.API)

func (tsf TrafficSenderFunc) Send(api *uaa_api.API) {
	tsf(api)
}