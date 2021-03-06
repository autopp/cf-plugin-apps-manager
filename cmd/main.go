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

package main

import (
	"fmt"
	"os"

	cf "code.cloudfoundry.org/cli/plugin"
	"github.com/autopp/cf-plugin-apps-manager/pkg/browser"
	"github.com/autopp/cf-plugin-apps-manager/pkg/plugin"
)

var version = "HEAD"

type AppsManagerPlugin struct{}

func main() {
	b := browser.New()
	p := plugin.NewAppsManagerPlugin(plugin.WithBrowser(b), plugin.WithVersion(version))
	cf.Start(p)

	if err := p.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
