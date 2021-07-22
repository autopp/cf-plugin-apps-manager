// Copyright (C) 2021 Akira Tanimura (@autopp)
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type AppsManagerPlugin struct {
	browser Browser
	version *plugin.VersionType
	stdout  io.Writer
	err     error
}

type Browser interface {
	Open(url string) error
}

func NewAppsManagerPlugin(opts ...Option) *AppsManagerPlugin {
	p := &AppsManagerPlugin{}
	for _, o := range append([]Option{WithStdout(os.Stdout)}, opts...) {
		o(p)
	}
	return p
}

func (p *AppsManagerPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    "apps-manager",
		Version: *p.version,
		Commands: []plugin.Command{
			{
				Name:     "apps-manager",
				HelpText: "Open AppsManager in browser",
			},
		},
	}
}

func (p *AppsManagerPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	if ok, err := cliConnection.HasAPIEndpoint(); err != nil {
		p.setErr("Cannot check current API endpoint: %w", err)
		return
	} else if !ok {
		p.setErr("No API endpoint set")
		return
	}

	apiEndpoint, err := cliConnection.ApiEndpoint()
	if err != nil {
		p.setErr("Cannot get API endpoint: %w", err)
		return
	}

	apiURL, err := url.Parse(apiEndpoint)
	if err != nil {
		p.setErr("Cannot parse API endpoint URL: %w", err)
		return
	}

	amEndpoint := fmt.Sprintf("%s://apps%s", apiURL.Scheme, apiURL.Host[strings.IndexRune(apiURL.Host, '.'):])

	if ok, err := cliConnection.HasOrganization(); err != nil {
		p.setErr("Cannot check current organization: %w", err)
	} else if !ok {
		p.open(amEndpoint)
		return
	}

	org, err := cliConnection.GetCurrentOrg()
	if err != nil {
		p.setErr("Cannot get current organization: %w", err)
		return
	}
	orgURL := fmt.Sprintf("%s/organizations/%s", amEndpoint, org.Guid)

	if ok, err := cliConnection.HasSpace(); err != nil {
		p.setErr("Cannot check current space: %w", err)
		return
	} else if !ok {
		p.open(orgURL)
		return
	}

	space, err := cliConnection.GetCurrentSpace()
	if err != nil {
		p.setErr("Cannot get current space: %w", err)
		return
	}

	spaceURL := fmt.Sprintf("%s/spaces/%s", orgURL, space.Guid)

	p.open(spaceURL)
}

var versionPattern = regexp.MustCompile(`\Av(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)\z`)

func parseVersion(v string) *plugin.VersionType {
	var major, minor, build int
	if versionNums := versionPattern.FindStringSubmatch(v); versionNums != nil {
		major, _ = strconv.Atoi(versionNums[1])
		minor, _ = strconv.Atoi(versionNums[2])
		build, _ = strconv.Atoi(versionNums[3])
	}

	return &plugin.VersionType{
		Major: major,
		Minor: minor,
		Build: build,
	}
}

func (p *AppsManagerPlugin) Err() error {
	return p.err
}

func (p *AppsManagerPlugin) setErr(format string, a ...interface{}) {
	p.err = fmt.Errorf(format, a...)
}

func (p *AppsManagerPlugin) open(url string) {
	fmt.Fprintln(p.stdout, url)
	p.browser.Open(url)
}
