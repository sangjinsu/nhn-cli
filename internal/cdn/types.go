package cdn

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

// Service
type CDNService struct {
	Domain      string `json:"domain"`
	DomainAlias string `json:"domainAlias"`
	Region      string `json:"region"`
	Status      string `json:"status"`
	Description string `json:"description"`
	UseOrigin   string `json:"useOrigin"`
	OriginURL   string `json:"originUrl"`
	CreatedAt   string `json:"createDate"`
}

type ServiceListResponse struct {
	Header   ResponseHeader `json:"header"`
	Contents []CDNService   `json:"contents"`
}

type ServiceResponse struct {
	Header  ResponseHeader `json:"header"`
	Content CDNService     `json:"content"`
}

type ServiceCreateRequest struct {
	Domain      string `json:"domain"`
	DomainAlias string `json:"domainAlias,omitempty"`
	Region      string `json:"region,omitempty"`
	Description string `json:"description,omitempty"`
	UseOrigin   string `json:"useOrigin,omitempty"`
	OriginURL   string `json:"originUrl"`
}

// Purge
type PurgeRequest struct {
	PurgeType string   `json:"purgeType"`
	Items     []string `json:"items,omitempty"`
}

type PurgeResponse struct {
	Header ResponseHeader `json:"header"`
}

// Auth Token
type AuthTokenCreateRequest struct {
	SessionID          string `json:"sessionId,omitempty"`
	SinglePath         string `json:"singlePath,omitempty"`
	SingleWildcardPath string `json:"singleWildcardPath,omitempty"`
	DurationSeconds    int    `json:"durationSeconds,omitempty"`
}

type AuthTokenResponse struct {
	Header ResponseHeader `json:"header"`
	Token  string         `json:"token"`
}
