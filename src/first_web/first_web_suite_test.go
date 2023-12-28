package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFirstWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FirstWeb Suite")
}
