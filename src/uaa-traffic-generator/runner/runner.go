package runner

import (
	"uaa-traffic-generator/sender"
	"github.com/cloudfoundry-community/go-uaa"
)

func RunAll(api *uaa.API, senders []sender.TrafficSender) {

	for _, s := range senders {
		s.Send(api)
	}
}
