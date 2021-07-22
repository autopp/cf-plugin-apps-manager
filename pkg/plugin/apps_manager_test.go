package plugin

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/cli/plugin"
	plugin_models "code.cloudfoundry.org/cli/plugin/models"
	"github.com/autopp/cf-plugin-apps-manager/pkg/plugin/mock"
)

var _ = Describe("AppsManager", func() {
	var ctrl *gomock.Controller
	var b *mock.MockBrowser
	JustBeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		b = mock.NewMockBrowser(ctrl)
	})

	JustAfterEach(func() {
		ctrl.Finish()
	})

	DescribeTable("GetMetadata",
		func(v string, major, minor, build int) {
			p := NewAppsManagerPlugin(WithBrowser(b), WithVersion(v))
			Expect(p.GetMetadata().Version).To(Equal(plugin.VersionType{Major: major, Minor: minor, Build: build}))
		},
		Entry(`"v1.2.3" -> (1, 2, 3)`, "v1.2.3", 1, 2, 3),
		Entry(`"" -> (0, 0, 0)`, "", 0, 0, 0),
		Entry(`"v1.a.2" -> (0, 0, 0)`, "", 0, 0, 0),
	)

	apiEndpoint := "https://api.cf.example.com"
	amEndpoint := "https://apps.cf.example.com"
	orgName := "myorg"
	orgGUID := "guid-of-myorg"
	spaceName := "myspace"
	spaceGUID := "guid-of-myspace"
	DescribeTable("Run",
		func(hasAPI, hasOrg, hasSpace bool, expectedURL string) {
			// Arrange
			cliConnection := mock.NewMockCliConnection(ctrl)
			cliConnection.EXPECT().HasAPIEndpoint().AnyTimes().Return(hasAPI, nil)
			if hasAPI {
				cliConnection.EXPECT().ApiEndpoint().Return(apiEndpoint, nil)
			}
			cliConnection.EXPECT().HasOrganization().AnyTimes().Return(hasOrg, nil)
			if hasOrg {
				cliConnection.EXPECT().GetCurrentOrg().Return(plugin_models.Organization{
					OrganizationFields: plugin_models.OrganizationFields{
						Name: orgName,
						Guid: orgGUID,
					},
				}, nil)
			}

			cliConnection.EXPECT().HasSpace().AnyTimes().Return(hasSpace, nil)
			if hasSpace {
				cliConnection.EXPECT().GetCurrentSpace().Return(plugin_models.Space{
					SpaceFields: plugin_models.SpaceFields{
						Name: spaceName,
						Guid: spaceGUID,
					},
				}, nil)
			}
			if expectedURL != "" {
				b.EXPECT().Open(expectedURL).Times(1).Return(nil)
			}

			// Act
			p := NewAppsManagerPlugin(WithBrowser(b), WithVersion("v0.1.0"))
			p.Run(cliConnection, []string{"apps-manager"})

			// Assert
			if expectedURL != "" {
				Expect(p.Err()).NotTo(HaveOccurred())
			} else {
				Expect(p.Err()).To(HaveOccurred())
			}
		},
		Entry(`org and space are set -> open the space URL`,
			true, true, true, amEndpoint+"/organizations/"+orgGUID+"/spaces/"+spaceGUID),
		Entry(`org is set but space is not set -> open the org URL`,
			true, true, false, amEndpoint+"/organizations/"+orgGUID),
		Entry(`org and space are not set -> open the am URL`,
			true, false, false, amEndpoint),
		Entry(`API endpoint is not set -> don't open and set error`,
			false, false, false, ""),
	)
})
