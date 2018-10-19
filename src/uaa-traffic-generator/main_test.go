package main_test

import (
	"encoding/json"
	"github.com/cloudfoundry-community/go-uaa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"io/ioutil"
	"os"
	"os/exec"
	"uaa-traffic-generator/config"
)

var _ = Describe("Main", func() {
	var trafficGeneratorSession *gexec.Session
	var fakeUaaServer *ghttp.Server

	var command *exec.Cmd

	BeforeEach(func() {
		fakeUaaServer = ghttp.NewServer()
		configPath := generateTrafficConfigFile(fakeUaaServer.URL())
		command = exec.Command(uaaTrafficGeneratorPath, "-config", configPath)

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

	JustBeforeEach(func() {
		var err error
		trafficGeneratorSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		trafficGeneratorSession.Kill()
	})

	It("should send traffic to the UAA", func() {
		Eventually(trafficGeneratorSession).Should(gexec.Exit(0))
		Expect(fakeUaaServer.ReceivedRequests()).To(HaveLen(2))
	})
})

func generateTrafficConfigFile(fakeUaaServerUrl string) string {
	trafficConfig := config.TrafficConfig{
		UaaCommands: []config.UaaCommand{
			{Cmd: "GetMe"},
		},
		Credentials: config.Credentials{
			Target: fakeUaaServerUrl,
		},
	}
	tempFile, err := ioutil.TempFile(os.TempDir(), "trafficconfig")
	Expect(err).NotTo(HaveOccurred())
	trafficConfigJson, err := json.Marshal(trafficConfig)
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(tempFile.Name(), trafficConfigJson, os.ModePerm)
	Expect(err).NotTo(HaveOccurred())
	return tempFile.Name()
}
