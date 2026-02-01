package appguard

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type DashboardResponse struct {
	Header ResponseHeader   `json:"header"`
	Data   []DashboardEntry `json:"data"`
}

type DashboardEntry struct {
	DetectedDate string `json:"detectedDate"`
	DetectedCnt  int    `json:"detectedCnt"`
	BlockedCnt   int    `json:"blockedCnt"`
}
