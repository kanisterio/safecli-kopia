package location

import "strings"

type LocType string

const (
	LocTypeS3          LocType = "s3"
	LocTypes3Compliant LocType = "s3Compliant"
	LocTypeGCS         LocType = "gcs"
	LocTypeAzure       LocType = "azure"
	LocTypeFilestore   LocType = "filestore"
)

type Location struct {
	Type             LocType
	Region           string
	BucketName       string
	Endpoint         string
	Prefix           string
	HasSkipSSLVerify bool
	CredentialsFile  string
}

func (l Location) IsInsecureEndpoint() bool {
	return strings.HasPrefix(l.Endpoint, "http:")
}

func (l Location) IsPointInTypeSupported() bool {
	switch l.Type {
	case LocTypeAzure, LocTypeS3, LocTypes3Compliant:
		return true
	default:
		return false
	}
}
