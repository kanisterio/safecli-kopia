// Copyright 2024 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fs

import (
	"testing"

	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/test"
	"gopkg.in/check.v1"

	"github.com/kanisterio/safecli-kopia"
)

func TestFilesystemOptions(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&test.ArgumentSuite{Cmd: "cmd", Arguments: []test.ArgumentTest{
	{
		Name:        "optRepoPath",
		Argument:    command.NewArguments(optRepoPath("/path/to/repo")),
		ExpectedCLI: []string{"cmd", "--path=/path/to/repo"},
	},
	{
		Name:        "Invalid RepoPath",
		Argument:    command.NewArguments(optRepoPath("")),
		ExpectedErr: cli.ErrInvalidRepoPath,
	},
}})
