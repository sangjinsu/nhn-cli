package deploy

type DeployExecuteRequest struct {
	ArtifactID    int    `json:"artifactId"`
	ServerGroupID int    `json:"serverGroupId"`
	ConcurrentNum int    `json:"concurrentNum,omitempty"`
	NextWhenFail  bool   `json:"nextWhenFail,omitempty"`
	DeployNote    string `json:"deployNote,omitempty"`
}

type DeployExecuteResponse struct {
	Header ResponseHeader `json:"header"`
	Body   *DeployResult  `json:"body,omitempty"`
}

type DeployResult struct {
	DeploymentID int    `json:"deploymentId"`
	Status       string `json:"status"`
}

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type BinaryUploadRequest struct {
	ArtifactID      int
	BinaryGroupKey  int
	ApplicationType string // "client" or "server"
	BinaryFile      string // 파일 경로
	Version         string
	Description     string
	OsType          string
	MetaFile        string // plist 파일 경로
	Fix             bool
}

type BinaryUploadResponse struct {
	Header ResponseHeader `json:"header"`
	Body   *BinaryResult  `json:"body,omitempty"`
}

type BinaryResult struct {
	DownloadUrl string `json:"downloadUrl"`
	BinaryKey   string `json:"binaryKey"`
}
