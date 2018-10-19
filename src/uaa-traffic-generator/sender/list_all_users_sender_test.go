package sender_test

import (
	. "uaa-traffic-generator/sender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/cloudfoundry-community/go-uaa"
	"golang.org/x/oauth2"
)

var _ = Describe("ListAllUsersSender", func() {
	var sender TrafficSender

	var api *uaa.API
	var fakeUaaServer *ghttp.Server

	BeforeEach(func() {
		sender = ListAllUsersSender{}

		fakeUaaServer = ghttp.NewServer()
		fakeUaaServer.AllowUnhandledRequests = true

		var err error
		api, err = uaa.NewWithPasswordCredentials(fakeUaaServer.URL(), "", "", "", "", "", uaa.JSONWebToken, true)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		fakeUaaServer.Close()
	})

	FContext("Send ListAllUsers Traffic", func() {
		BeforeEach(func() {
			fakeUaaServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/oauth/token"),
					ghttp.RespondWithJSONEncoded(200, oauth2.Token{AccessToken: "abc"}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/Users", "count=100&sortOrder=ascending&startIndex=1"),
					ghttp.RespondWithJSONEncoded(200, paginatedUserList{}),
				),
			)
		})

		It("should send traffic to the UAA", func() {
			sender.Send(api)
			Expect(fakeUaaServer.ReceivedRequests()).To(HaveLen(3))
		})
	})
})


type paginatedUserList struct {
	uaa.Page
	Resources []uaa.User   `json:"resources"`
	Schemas   []string `json:"schemas"`
}
