package objectstorage

type Container struct {
	Name         string `json:"name"`
	Count        int    `json:"count"`
	Bytes        int64  `json:"bytes"`
	LastModified string `json:"last_modified,omitempty"`
}

type Object struct {
	Name         string `json:"name"`
	Hash         string `json:"hash"`
	Bytes        int64  `json:"bytes"`
	ContentType  string `json:"content_type"`
	LastModified string `json:"last_modified"`
}

type ContainerMetadata struct {
	ObjectCount string
	BytesUsed   string
	ReadACL     string
	WriteACL    string
}

type ObjectMetadata struct {
	ContentLength string
	ContentType   string
	ETag          string
	LastModified  string
}
