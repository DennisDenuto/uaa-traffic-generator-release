package config_test

import (
	. "uaa-traffic-generator/config"

	"github.com/cloudfoundry-community/go-uaa"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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
		func(configCmd string, credentials Credentials, expectedSender sender.TrafficSender) {
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

	Context("A command with a Loop property", func() {
		BeforeEach(func() {
			config.UaaCommands = append(config.UaaCommands,
				UaaCommand{
					Cmd:  "GetMe",
					Loop: 2,
				},
				UaaCommand{
					Cmd:  "ListAllUsers",
					Loop: 2,
				},
			)

		})

		It("should create a new sender that will Loop", func() {
			senders, _, err := NewSenders(config)
			Expect(err).NotTo(HaveOccurred())

			Expect(senders[0]).To(BeAssignableToTypeOf(sender.TrafficSenderFunc(func(api *uaa.API) {})))
			Expect(senders[1]).To(BeAssignableToTypeOf(sender.TrafficSenderFunc(func(api *uaa.API) {})))
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
