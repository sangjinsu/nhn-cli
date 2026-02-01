package pipeline

type PipelineExecuteResponse struct {
	Header ResponseHeader `json:"header"`
}

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}
