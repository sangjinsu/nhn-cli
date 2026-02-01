package dns

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListZones() ([]Zone, error) {
	resp, err := c.httpClient.Get(c.url("/zones"), c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Zone 목록 조회 실패: %w", err)
	}

	var result ZoneListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.ZoneList, nil
}

func (c *Client) GetZone(zoneID string) (*Zone, error) {
	url := fmt.Sprintf("%s/zones?zoneIdList=%s", c.baseURL, zoneID)
	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Zone 조회 실패: %w", err)
	}

	var result ZoneListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}
	if len(result.ZoneList) == 0 {
		return nil, fmt.Errorf("Zone '%s'을(를) 찾을 수 없습니다", zoneID)
	}

	return &result.ZoneList[0], nil
}

func (c *Client) CreateZone(req *ZoneCreateRequest) (*Zone, error) {
	resp, err := c.httpClient.Post(c.url("/zones"), req, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Zone 생성 실패: %w", err)
	}

	var result ZoneResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Zone, nil
}

func (c *Client) UpdateZone(zoneID string, req *ZoneUpdateRequest) (*Zone, error) {
	url := fmt.Sprintf("%s/zones/%s", c.baseURL, zoneID)
	resp, err := c.httpClient.Put(url, req, c.getOpts())
	if err != nil {
		return nil, fmt.Errorf("Zone 수정 실패: %w", err)
	}

	var result ZoneResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}
	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Zone, nil
}

func (c *Client) DeleteZone(zoneID string) error {
	url := fmt.Sprintf("%s/zones/async?zoneIdList=%s", c.baseURL, zoneID)
	resp, err := c.httpClient.Delete(url, c.getOpts())
	if err != nil {
		return fmt.Errorf("Zone 삭제 실패: %w", err)
	}

	var result struct {
		Header ResponseHeader `json:"header"`
	}
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}
	return checkResponse(result.Header)
}
