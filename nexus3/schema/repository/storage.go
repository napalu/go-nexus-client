package repository

const (
	StorageWritePolicyAllow     StorageWritePolicy = "ALLOW"
	StorageWritePolicyAllowOnce StorageWritePolicy = "ALLOW_ONCE"
	StorageWritePolicyAllowDeny StorageWritePolicy = "DENY"
)

type StorageWritePolicy string

// HostedStorage contains repository storage for hosted
type HostedStorage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName"`

	// StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`

	// WritePolicy controls if deployments of and updates to assets are allowed
	WritePolicy *StorageWritePolicy `json:"writePolicy,omitempty"`

	// Allow redeploying 'latest' tag defer to WritePolicy otherwise
	LatestPolicy *bool `json:"latestPolicy,omitempty"`
}

// Storage contains repository storage
type Storage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName"`

	// StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
}
