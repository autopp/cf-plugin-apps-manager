package plugin

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/autopp/cf-plugin-apps-manager/pkg/plugin/mock"
)

var _ = Describe("AppsManager", func() {
	var ctrl *gomock.Controller
	JustBeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	JustAfterEach(func() {
		ctrl.Finish()
	})

	DescribeTable("GetMetadata",
		func(v string, major, minor, build int) {
			b := mock.NewMockBrowser(ctrl)
			p := NewAppsManagerPlugin(b, v)
			Expect(p.GetMetadata().Version).To(Equal(plugin.VersionType{Major: major, Minor: minor, Build: build}))
		},
		Entry(`"v1.2.3" -> (1, 2, 3)`, "v1.2.3", 1, 2, 3),
		Entry(`"" -> (0, 0, 0)`, "", 0, 0, 0),
		Entry(`"v1.a.2" -> (0, 0, 0)`, "", 0, 0, 0),
	)
})
