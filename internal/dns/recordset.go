package dns

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListRecordSets(zoneID string) ([]RecordSet, error) {
	url := fmt.Sprintf("%s/zones/%s/recordsets", c.baseURL, zoneID)
	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Record Set 목록 조회 실패: %w", err)
	}

	var result RecordSetListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.RecordsetList, nil
}

func (c *Client) GetRecordSet(zoneID, recordsetID string) (*RecordSet, error) {
	url := fmt.Sprintf("%s/zones/%s/recordsets?recordsetIdList=%s", c.baseURL, zoneID, recordsetID)
	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Record Set 조회 실패: %w", err)
	}

	var result RecordSetListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}
	if len(result.RecordsetList) == 0 {
		return nil, fmt.Errorf("Record Set '%s'을(를) 찾을 수 없습니다", recordsetID)
	}

	return &result.RecordsetList[0], nil
}

func (c *Client) CreateRecordSet(zoneID string, req *RecordSetCreateRequest) (*RecordSet, error) {
	url := fmt.Sprintf("%s/zones/%s/recordsets", c.baseURL, zoneID)
	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Record Set 생성 실패: %w", err)
	}

	var result RecordSetResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Recordset, nil
}

func (c *Client) UpdateRecordSet(zoneID, recordsetID string, req *RecordSetUpdateRequest) (*RecordSet, error) {
	url := fmt.Sprintf("%s/zones/%s/recordsets/%s", c.baseURL, zoneID, recordsetID)
	resp, err := c.httpClient.Put(url, req, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Record Set 수정 실패: %w", err)
	}

	var result RecordSetResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Recordset, nil
}

func (c *Client) DeleteRecordSet(zoneID, recordsetID string) error {
	url := fmt.Sprintf("%s/zones/%s/recordsets?recordsetIdList=%s", c.baseURL, zoneID, recordsetID)
	resp, err := c.httpClient.Delete(url, c.getOpts())
	if err != nil {
		return fmt.Errorf("Record Set 삭제 실패: %w", err)
	}

	var result struct {
		Header ResponseHeader `json:"header"`
	}
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}
	return checkResponse(result.Header)
}
