package sender

import "github.com/cloudfoundry-community/go-uaa"

func Loop(times int, sender TrafficSender) TrafficSender {
	return TrafficSenderFunc(func(api *uaa.API) {
		for i := 0; i < times; i++ {
			sender.Send(api)
		}
	})
}
