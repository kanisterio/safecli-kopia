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

package s3

import (
	"testing"

	"github.com/kanisterio/safecli/test"
	"gopkg.in/check.v1"

	cli "github.com/kanisterio/safecli-kopia"
	"github.com/kanisterio/safecli-kopia/repository/storage/location"
)

func TestNewS3(t *testing.T) { check.TestingT(t) }

func newLocation(prefix, endpoint, region, bucket string, skipSSLVerify bool) location.Location {
	return location.Location{
		Prefix:           prefix,
		Endpoint:         endpoint,
		Region:           region,
		BucketName:       bucket,
		HasSkipSSLVerify: skipSSLVerify,
	}
}

// s3test is a test case for NewS3.
type s3test struct {
	Name        string
	Location    location.Location
	RepoPath    string
	ExpectedCLI []string
	ExpectedErr error
}

// newS3Test creates a new test case for NewS3.
func newS3Test(s3t s3test) test.ArgumentTest {
	return test.ArgumentTest{
		Name:        s3t.Name,
		Argument:    New(s3t.Location, s3t.RepoPath),
		ExpectedCLI: s3t.ExpectedCLI,
		ExpectedErr: s3t.ExpectedErr,
	}
}

// toArgTests converts a list of s3tests to a list of ArgumentTests.
func toArgTests(s3tests []s3test) []test.ArgumentTest {
	argTests := make([]test.ArgumentTest, len(s3tests))
	for i, s3t := range s3tests {
		argTests[i] = newS3Test(s3t)
	}
	return argTests
}

var _ = check.Suite(&test.ArgumentSuite{Cmd: "cmd", Arguments: toArgTests([]s3test{
	{
		Name:     "NewS3",
		Location: newLocation("prefix", "http://endpoint/path/", "region", "bucket", true),
		RepoPath: "repoPath",
		ExpectedCLI: []string{"cmd", "s3",
			"--region=region",
			"--bucket=bucket",
			"--endpoint=endpoint/path",
			"--prefix=prefix/repoPath/",
			"--disable-tls",
			"--disable-tls-verification",
		},
	},
	{
		Name:     "NewS3 w/o logger should not panic",
		Location: newLocation("prefix", "http://endpoint/path/", "region", "bucket", true),
		RepoPath: "repoPath",
		ExpectedCLI: []string{"cmd", "s3",
			"--region=region",
			"--bucket=bucket",
			"--endpoint=endpoint/path",
			"--prefix=prefix/repoPath/",
			"--disable-tls",
			"--disable-tls-verification",
		},
	},
	{
		Name:     "NewS3 with empty repoPath and https endpoint",
		Location: newLocation("prefix", "https://endpoint/path/", "region", "bucket", false),
		ExpectedCLI: []string{"cmd", "s3",
			"--region=region",
			"--bucket=bucket",
			"--endpoint=endpoint/path",
			"--prefix=prefix/",
		},
	},
	{
		Name:     "NewS3 with empty repoPath and endpoint",
		Location: newLocation("prefix", "", "region", "bucket", true),
		ExpectedCLI: []string{"cmd", "s3",
			"--region=region",
			"--bucket=bucket",
			"--prefix=prefix/",
			"--disable-tls-verification",
		},
	},
	{
		Name:     "NewS3 with empty repoPath, prefix and endpoint",
		Location: newLocation("", "", "region", "bucket", true),
		ExpectedCLI: []string{"cmd", "s3",
			"--region=region",
			"--bucket=bucket",
			"--prefix=",
			"--disable-tls-verification",
		},
	},
	{
		Name:        "NewS3 with empty repoPath, prefix, endpoint and bucket",
		ExpectedErr: cli.ErrInvalidBucketName,
	},
	{
		Name:     "NewS3 with empty logger should not panic",
		Location: newLocation("", "https://endpoint/path/", "", "bucket", false),
		ExpectedCLI: []string{"cmd", "s3",
			"--bucket=bucket",
			"--endpoint=endpoint/path",
			"--prefix=",
		},
	},
})})
