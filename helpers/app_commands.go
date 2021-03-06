package helpers

import (
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/cf"
	"time"
)

func PushApp(appName string, appPath string) {
	Expect(Cf("push", appName, "-p", appPath)).To(SayWithTimeout("App started", time.Minute*2))
}

func DeleteApp(appName string) {
	Expect(Cf("delete", appName, "-f")).To(SayWithTimeout("OK", time.Minute*2))
}
