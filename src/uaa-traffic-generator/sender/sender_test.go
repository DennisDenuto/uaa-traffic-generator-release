package sender_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"uaa-traffic-generator/sender"
)

var _ = Describe("Sender", func() {

	var fakeUaaServer *ghttp.Server

	BeforeEach(func() {
		fakeUaaServer = ghttp.NewServer()
	})

	AfterEach(func() {
		fakeUaaServer.Close()
	})

	Context("Send Obtain Token", func() {
		BeforeEach(func() {
			fakeUaaServer.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/oauth/token", "a=a"),
			))
		})

		It("should request a token from the UAA", func() {

			sender.Send()
			Expect(fakeUaaServer.ReceivedRequests()).To(Equal(1))
		})
	})


})
