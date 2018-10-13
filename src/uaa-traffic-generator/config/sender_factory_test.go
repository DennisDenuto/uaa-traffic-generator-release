package config_test

import (
	. "uaa-traffic-generator/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"uaa-traffic-generator/sender"
)

var _ = Describe("SenderFactory", func() {
	var config TrafficConfig
	BeforeEach(func() {
		config = TrafficConfig{}
	})

	Context("given a config for sending 'GetMe' traffic", func() {
		BeforeEach(func() {
			config.UaaCommands = append(config.UaaCommands, UaaCommand{
				Cmd: "GetMe",
			})
		})
		It("should build correct sender", func() {
			senders := NewSenders(TrafficConfig{})

			Expect(senders).To(HaveLen(1))
			Expect(senders[0]).To(BeAssignableToTypeOf(sender.GetMeSender{}))
		})
	})
})
