package runner

import (
	"uaa-traffic-generator/sender"
	"github.com/cloudfoundry-community/go-uaa"
)

func RunAll(api *uaa.API, senders []sender.Sender) {

	for _, sender := range senders{
		sender.Send(api)
	}
}