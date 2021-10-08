package go2css_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGo2css(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go2css Suite")
}
