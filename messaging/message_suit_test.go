package messaging

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMessagingSUIT(t *testing.T) {
	RegisterFailHandler(Fail)
	//junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
	junitReporter := reporters.NewJUnitReporter("../test-report/cireport.txt")
	RunSpecsWithDefaultAndCustomReporters(t, "Messaging Suit", []Reporter{junitReporter})
}
