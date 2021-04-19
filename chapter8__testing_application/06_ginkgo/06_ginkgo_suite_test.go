package 06_ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test06Ginkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "06Ginkgo Suite")
}
