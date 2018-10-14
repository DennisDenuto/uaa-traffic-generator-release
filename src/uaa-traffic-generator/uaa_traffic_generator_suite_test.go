package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/onsi/gomega/gexec"
)

func TestUaaTrafficGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UaaTrafficGenerator Suite")
}

var uaaTrafficGeneratorPath string

var _ = BeforeSuite(func() {
	var err error
	uaaTrafficGeneratorPath, err = gexec.Build("uaa-traffic-generator")
	Expect(err).NotTo(HaveOccurred())
})
