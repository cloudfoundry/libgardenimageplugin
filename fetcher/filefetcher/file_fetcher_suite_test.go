package filefetcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTarFetcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Fetcher Suite")
}
