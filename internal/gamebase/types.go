package gamebase

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

// Member
type Member struct {
	UserID        string `json:"userId"`
	Valid         string `json:"valid"`
	AppID         string `json:"appId"`
	RegDate       string `json:"regDate"`
	LastLoginDate string `json:"lastLoginDate"`
}

type MemberResponse struct {
	Header ResponseHeader `json:"header"`
	Member Member         `json:"member"`
}

type MemberListResponse struct {
	Header  ResponseHeader `json:"header"`
	Members []Member       `json:"members"`
}

// Ban
type BanInfo struct {
	UserID    string `json:"userId"`
	BanType   string `json:"banType"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
	Reason    string `json:"reason"`
}

type BanCreateRequest struct {
	UserID    string `json:"userId"`
	BanType   string `json:"banType,omitempty"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
	Reason    string `json:"reason,omitempty"`
	Message   string `json:"message,omitempty"`
}

type BanListResponse struct {
	Header ResponseHeader `json:"header"`
	Bans   []BanInfo      `json:"banList"`
}

type BanResponse struct {
	Header ResponseHeader `json:"header"`
}

type BanReleaseRequest struct {
	UserID string `json:"userId"`
}

// Launching
type LaunchingInfo struct {
	Status struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
}

type LaunchingResponse struct {
	Header    ResponseHeader `json:"header"`
	Launching LaunchingInfo  `json:"launching"`
}

// Auth
type TokenValidateResponse struct {
	Header ResponseHeader `json:"header"`
	Valid  bool           `json:"valid"`
}
