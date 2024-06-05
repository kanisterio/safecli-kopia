// Copyright 2024 The Kanister Autho
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

package repository

import (
	"time"

	"github.com/kanisterio/safecli-kopia/args"
	"github.com/kanisterio/safecli-kopia/repository/storage/location"
)

var (
	common = args.Common{
		RepoPassword:   "encr-key",
		ConfigFilePath: "path/kopia.config",
		LogDirectory:   "cache/log",
	}

	cache = args.Cache{
		CacheDirectory:           "/tmp/cache.dir",
		ContentCacheSizeLimitMB:  0,
		MetadataCacheSizeLimitMB: 0,
	}
)

var (
	retentionMode   = "Locked"
	retentionPeriod = 15 * time.Minute

	locFS = location.Location{
		Type:   "filestore",
		Prefix: "test-prefix",
	}

	locAzure = location.Location{
		Type:       "azure",
		BucketName: "test-bucket",
		Prefix:     "test-prefix",
	}

	locGCS = location.Location{
		Type:            "gcs",
		BucketName:      "test-bucket",
		Prefix:          "test-prefix",
		CredentialsFile: "/tmp/creds.txt",
	}

	locS3 = location.Location{
		Type:             "s3",
		Endpoint:         "test-endpoint",
		Region:           "test-region",
		BucketName:       "test-bucket",
		Prefix:           "test-prefix",
		HasSkipSSLVerify: false,
	}

	locS3Compliant = location.Location{
		Type:             "s3Compliant",
		Endpoint:         "test-endpoint",
		Region:           "test-region",
		BucketName:       "test-bucket",
		Prefix:           "test-prefix",
		HasSkipSSLVerify: false,
	}

	locFTP = location.Location{
		Type: "ftp",
	}
)
