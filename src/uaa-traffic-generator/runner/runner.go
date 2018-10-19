package runner

import (
	"github.com/cloudfoundry-community/go-uaa"
	"sync"
	"uaa-traffic-generator/sender"
)

func RunAll(api *uaa.API, senders []sender.TrafficSender) {
	wg := sync.WaitGroup{}
	wg.Add(len(senders))

	for _, s := range senders {
		go func() {
			defer wg.Done()
			s.Send(api)
		}()
	}

	wg.Wait()
}
