package plugin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	cf "code.cloudfoundry.org/cli/plugin"

	"github.com/autopp/cf-plugin-apps-manager/pkg/plugin"
)

var _ = Describe("AppsManager", func() {
	DescribeTable("GetMetadata",
		func(v string, major, minor, build int) {
			p := plugin.NewAppsManagerPlugin(nil, v)
			Expect(p.GetMetadata().Version).To(Equal(cf.VersionType{Major: major, Minor: minor, Build: build}))
		},
		Entry(`"v1.2.3" -> (1, 2, 3)`, "v1.2.3", 1, 2, 3),
		Entry(`"" -> (0, 0, 0)`, "", 0, 0, 0),
		Entry(`"v1.a.2" -> (0, 0, 0)`, "", 0, 0, 0),
	)
})
