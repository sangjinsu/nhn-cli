package dns

type ResponseHeader struct {
	IsSuccessful  bool   `json:"isSuccessful"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

// Zone

type Zone struct {
	ZoneID         string `json:"zoneId"`
	ZoneName       string `json:"zoneName"`
	ZoneStatus     string `json:"zoneStatus"`
	Description    string `json:"description"`
	RecordsetCount int    `json:"recordsetCount"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

type ZoneListResponse struct {
	Header     ResponseHeader `json:"header"`
	TotalCount int            `json:"totalCount"`
	ZoneList   []Zone         `json:"zoneList"`
}

type ZoneResponse struct {
	Header ResponseHeader `json:"header"`
	Zone   Zone           `json:"zone"`
}

type ZoneCreateRequest struct {
	Zone ZoneCreateBody `json:"zone"`
}

type ZoneCreateBody struct {
	ZoneName    string `json:"zoneName"`
	Description string `json:"description,omitempty"`
}

type ZoneUpdateRequest struct {
	Zone ZoneUpdateBody `json:"zone"`
}

type ZoneUpdateBody struct {
	Description string `json:"description,omitempty"`
}

// Record Set

type RecordSet struct {
	RecordsetID     string   `json:"recordsetId"`
	RecordsetName   string   `json:"recordsetName"`
	RecordsetType   string   `json:"recordsetType"`
	RecordsetTTL    int      `json:"recordsetTtl"`
	RecordsetStatus string   `json:"recordsetStatus"`
	RecordList      []Record `json:"recordList"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
}

type Record struct {
	RecordDisabled bool   `json:"recordDisabled"`
	RecordContent  string `json:"recordContent"`
}

type RecordSetListResponse struct {
	Header        ResponseHeader `json:"header"`
	TotalCount    int            `json:"totalCount"`
	RecordsetList []RecordSet    `json:"recordsetList"`
}

type RecordSetResponse struct {
	Header    ResponseHeader `json:"header"`
	Recordset RecordSet      `json:"recordset"`
}

type RecordSetCreateRequest struct {
	Recordset RecordSetCreateBody `json:"recordset"`
}

type RecordSetCreateBody struct {
	RecordsetName string   `json:"recordsetName"`
	RecordsetType string   `json:"recordsetType"`
	RecordsetTTL  int      `json:"recordsetTtl"`
	RecordList    []Record `json:"recordList"`
}

type RecordSetUpdateRequest struct {
	Recordset RecordSetUpdateBody `json:"recordset"`
}

type RecordSetUpdateBody struct {
	RecordsetType string   `json:"recordsetType,omitempty"`
	RecordsetTTL  int      `json:"recordsetTtl,omitempty"`
	RecordList    []Record `json:"recordList,omitempty"`
}
