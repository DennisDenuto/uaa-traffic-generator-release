package runner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"uaa-traffic-generator/sender"
	"uaa-traffic-generator/sender/senderfakes"
	. "uaa-traffic-generator/runner"
	"github.com/cloudfoundry-community/go-uaa"
)

var _ = Describe("Runner", func() {
	var senders []sender.Sender
	var fakeSender *senderfakes.FakeSender
	var uaaApi *uaa.API

	BeforeEach(func() {
		fakeSender = &senderfakes.FakeSender{}
		senders = append(senders, fakeSender)
		uaaApi = &uaa.API{}
	})

	It("should run all provided senders", func() {
		RunAll(uaaApi, senders)
		Expect(fakeSender.SendCallCount()).To(Equal(1))
		Expect(fakeSender.SendArgsForCall(0)).To(Equal(uaaApi))

	})
})
