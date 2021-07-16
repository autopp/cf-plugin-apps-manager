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
	"code.cloudfoundry.org/cli/plugin"
)

type Version plugin.VersionType

type AppsManagerPlugin struct {
	browser Browser
	version Version
}

type Browser interface {
	Open(url string)
}

func NewAppsManagerPlugin(b Browser, v Version) *AppsManagerPlugin {
	return &AppsManagerPlugin{browser: b, version: v}
}

func (p *AppsManagerPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    "AppsManager",
		Version: plugin.VersionType(p.version),
		Commands: []plugin.Command{
			{
				Name:     "apps-manager",
				HelpText: "Open AppsManager in browser",
			},
		},
	}
}

func (p *AppsManagerPlugin) Run(cliConnection plugin.CliConnection, args []string) {
}
