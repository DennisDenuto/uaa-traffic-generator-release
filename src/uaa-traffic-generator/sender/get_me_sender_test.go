package sender_test

import (
	. "uaa-traffic-generator/sender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/cloudfoundry-community/go-uaa"
)

var _ = Describe("GetMeSender", func() {
	var sender TrafficSender

	var api *uaa.API
	var fakeUaaServer *ghttp.Server

	BeforeEach(func() {
		sender = GetMeSender{}

		fakeUaaServer = ghttp.NewServer()
		fakeUaaServer.AllowUnhandledRequests = true

		var err error
		api, err = uaa.NewWithPasswordCredentials(fakeUaaServer.URL(), "", "", "", "", "", uaa.JSONWebToken, true)
		Expect(err).NotTo(HaveOccurred())
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

		It("should send traffic to the UAA", func() {
			sender.Send(api)
			Expect(fakeUaaServer.ReceivedRequests()).To(HaveLen(2))
		})
	})

})
