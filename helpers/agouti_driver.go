package helpers

import (
	"fmt"
	"os/exec"

	"github.com/onsi/ginkgo"
	"github.com/sclevine/agouti"
)

func AgoutiDriver() *agouti.WebDriver {
	if _, err := exec.LookPath("phantomjs"); err == nil {
		fmt.Fprintln(ginkgo.GinkgoWriter, "WARNING: using phantomjs, which is flaky in CI, but is more convenient during development")
		return agouti.PhantomJS()
	} else {
		return agouti.Selenium(agouti.Browser("firefox"))
	}
}
