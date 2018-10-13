package sender_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"uaa-traffic-generator/sender"
	"cloudfoundry-community/go-uaa/go-uaa"
)

var _ = Describe("Sender", func() {

	var fakeUaaServer *ghttp.Server

	BeforeEach(func() {
		fakeUaaServer = ghttp.NewServer()
		fakeUaaServer.AllowUnhandledRequests = true
	})

	AfterEach(func() {
		fakeUaaServer.Close()
	})

	Context("Send UserInfo Traffic", func() {
		BeforeEach(func() {
			fakeUaaServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWith(200, "{}"),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/userinfo", "scheme=openid"),
					ghttp.RespondWithJSONEncoded(200, uaa.UserInfo{}),
				),
			)
		})

		It("should request a token from the UAA", func() {

			sender.Send(fakeUaaServer.URL())
			Expect(fakeUaaServer.ReceivedRequests()).To(HaveLen(2))
		})
	})

})
