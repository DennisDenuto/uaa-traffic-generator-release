package runner_test

import (
	"github.com/cloudfoundry-community/go-uaa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "uaa-traffic-generator/runner"
	"uaa-traffic-generator/sender"
	"uaa-traffic-generator/sender/senderfakes"
)

var _ = Describe("Runner", func() {
	var senders []sender.TrafficSender
	var fakeSender *senderfakes.FakeTrafficSender
	var uaaApi *uaa.API

	BeforeEach(func() {
		fakeSender = &senderfakes.FakeTrafficSender{}
		senders = append(senders, fakeSender)
		uaaApi = &uaa.API{}
	})

	It("should run all provided senders", func() {
		RunAll(uaaApi, senders)
		Expect(fakeSender.SendCallCount()).To(Equal(1))
		Expect(fakeSender.SendArgsForCall(0)).To(Equal(uaaApi))

	})
})
