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
	var validUserCredentials Credentials

	BeforeEach(func() {
		validUserCredentials = Credentials{
			Target:       "http://target",
			GrantType:    "password",
			ClientId:     "clientid",
			ClientSecret: "clientsecret",
			Username:     "username",
			UserPassword: "userpassword",
		}
		config = TrafficConfig{
			Credentials: validUserCredentials,
		}
	})

	table.DescribeTable("given a config for sending traffic",
		func(configCmd string, credentials Credentials, expectedSender sender.Sender) {
			config.UaaCommands = append(config.UaaCommands, UaaCommand{
				Cmd: configCmd,
			})

			senders, api, err := NewSenders(config)
			Expect(err).NotTo(HaveOccurred())

			Expect(senders).To(HaveLen(1))
			Expect(senders[0]).To(BeAssignableToTypeOf(expectedSender))
			Expect(api).To(Not(BeNil()))
			Expect(api.AuthenticatedClient).To(Not(BeNil()))
			Expect(api.TargetURL.String()).To(Equal("http://target"))
		},

		table.Entry("config with GetMe command", "GetMe", validUserCredentials, sender.GetMeSender{}),
		table.Entry("config with ListAllUsers command", "ListAllUsers", validUserCredentials, sender.ListAllUsersSender{}),
	)

	Context("multiple commands are provided", func() {
		BeforeEach(func() {
			config.UaaCommands = append(config.UaaCommands,
				UaaCommand{
					Cmd: "GetMe",
				},
				UaaCommand{
					Cmd: "ListAllUsers",
				},
			)

		})
		It("should create a new sender for each configured command", func() {
			senders, _, err := NewSenders(config)
			Expect(err).NotTo(HaveOccurred())

			Expect(senders).To(HaveLen(2))
		})
	})

	Context("Given an invalid target url", func() {
		BeforeEach(func() {
			validUserCredentials.Target = "not a url"
			config = TrafficConfig{
				Credentials: validUserCredentials,
			}
		})

		It("should return an error", func() {
			_, _, err := NewSenders(config)
			Expect(err).To(HaveOccurred())
		})
	})
})
