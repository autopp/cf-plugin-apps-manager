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

package browser

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

type Browser struct{}

func New() *Browser {
	return &Browser{}
}

func (*Browser) Open(url string) error {
	var browserCmd []string
	switch runtime.GOOS {
	case "windows":
		browserCmd = []string{"cmd", "/c", "start"}
	case "darwin":
		browserCmd = []string{"open"}
	case "linux":
		browserCmd = []string{"xdg-open"}
	default:
		return fmt.Errorf("unsupported os: %s", runtime.GOOS)
	}

	name := browserCmd[0]
	args := append(browserCmd[1:], url)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return exec.CommandContext(ctx, name, args...).Run()
}
