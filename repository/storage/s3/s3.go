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
	"strings"

	"github.com/kanisterio/safecli/command"

	"github.com/kanisterio/safecli-kopia/internal"
	"github.com/kanisterio/safecli-kopia/repository/storage/location"
)

// New creates a new subcommand for the S3 storage.
func New(location location.Location, repoPathPrefix string) command.Applier {
	endpoint := resolveS3Endpoint(location.Endpoint)
	prefix := internal.GenerateFullRepoPath(location.Prefix, repoPathPrefix)
	return command.NewArguments(subcmdS3,
		optRegion(location.Region),
		optBucket(location.BucketName),
		optEndpoint(endpoint),
		optPrefix(prefix),
		optDisableTLS(location.IsInsecureEndpoint()),
		optDisableTLSVerify(location.HasSkipSSLVerify),
	)
}

// resolveS3Endpoint removes the trailing slash and
// protocol from provided endpoint and
// returns the absolute endpoint string.
func resolveS3Endpoint(endpoint string) string {
	if endpoint == "" {
		return ""
	}

	if strings.HasSuffix(endpoint, "/") {
		endpoint = strings.TrimRight(endpoint, "/")
	}

	sp := strings.SplitN(endpoint, "://", 2)

	return sp[len(sp)-1]
}
