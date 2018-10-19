package runner

import (
	"uaa-traffic-generator/sender"
	"github.com/cloudfoundry-community/go-uaa"
	"sync"
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
