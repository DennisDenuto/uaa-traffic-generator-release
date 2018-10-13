package config_test

import (
	. "uaa-traffic-generator/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/ginkgo/extensions/table"
	"uaa-traffic-generator/sender"
)

var _ = Describe("SenderFactory", func() {
	var config TrafficConfig
	BeforeEach(func() {
		config = TrafficConfig{}
	})

	table.DescribeTable("given a config for sending traffic",
		func(configCmd string, expectedSender sender.Sender) {
			config.UaaCommands = append(config.UaaCommands, UaaCommand{
				Cmd: configCmd,
			})

			senders := NewSenders(TrafficConfig{})

			Expect(senders).To(HaveLen(1))
			Expect(senders[0]).To(BeAssignableToTypeOf(expectedSender))
		},

		table.Entry("config with GetMe command", "GetMe", sender.GetMeSender{}),
	)
})
