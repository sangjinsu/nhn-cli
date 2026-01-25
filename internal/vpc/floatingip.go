package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

const DefaultExternalNetworkID = "b04b1c31-f2e9-4ae0-a264-02b7d61ad618"

func (c *Client) ListFloatingIPs() ([]FloatingIP, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/floatingips"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FloatingIPListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("플로팅 IP 목록 조회 실패: %w", err)
	}

	return result.FloatingIPs, nil
}

func (c *Client) GetFloatingIP(id string) (*FloatingIP, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/floatingips/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FloatingIPResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("플로팅 IP 조회 실패: %w", err)
	}

	return &result.FloatingIP, nil
}

func (c *Client) CreateFloatingIP(networkID string) (*FloatingIP, error) {
	if networkID == "" {
		networkID = DefaultExternalNetworkID
	}

	reqBody := FloatingIPCreateRequest{
		FloatingIP: FloatingIPCreateBody{
			FloatingNetworkID: networkID,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/floatingips"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FloatingIPResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("플로팅 IP 생성 실패: %w", err)
	}

	return &result.FloatingIP, nil
}

func (c *Client) AssociateFloatingIP(floatingIPID, portID string) (*FloatingIP, error) {
	reqBody := FloatingIPUpdateRequest{
		FloatingIP: FloatingIPUpdateBody{
			PortID: &portID,
		},
	}

	resp, err := c.httpClient.Put(c.url("/v2.0/floatingips/"+floatingIPID), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FloatingIPResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("플로팅 IP 연결 실패: %w", err)
	}

	return &result.FloatingIP, nil
}

func (c *Client) DisassociateFloatingIP(floatingIPID string) (*FloatingIP, error) {
	reqBody := FloatingIPUpdateRequest{
		FloatingIP: FloatingIPUpdateBody{
			PortID: nil,
		},
	}

	resp, err := c.httpClient.Put(c.url("/v2.0/floatingips/"+floatingIPID), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FloatingIPResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("플로팅 IP 연결 해제 실패: %w", err)
	}

	return &result.FloatingIP, nil
}

func (c *Client) DeleteFloatingIP(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/floatingips/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("플로팅 IP 삭제 실패: %w", err)
	}

	return nil
}
