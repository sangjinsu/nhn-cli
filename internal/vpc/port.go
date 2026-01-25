package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListPorts() ([]Port, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/ports"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result PortListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("네트워크 인터페이스 목록 조회 실패: %w", err)
	}

	return result.Ports, nil
}

func (c *Client) GetPort(id string) (*Port, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/ports/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result PortResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("네트워크 인터페이스 조회 실패: %w", err)
	}

	return &result.Port, nil
}

func (c *Client) CreatePort(networkID, name string) (*Port, error) {
	reqBody := PortCreateRequest{
		Port: PortCreateBody{
			NetworkID:    networkID,
			Name:         name,
			AdminStateUp: true,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/ports"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result PortResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("네트워크 인터페이스 생성 실패: %w", err)
	}

	return &result.Port, nil
}

func (c *Client) DeletePort(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/ports/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("네트워크 인터페이스 삭제 실패: %w", err)
	}

	return nil
}

func (c *Client) GetPortByInstanceID(instanceID string) (*Port, error) {
	ports, err := c.ListPorts()
	if err != nil {
		return nil, err
	}

	for _, port := range ports {
		if port.DeviceID == instanceID {
			return &port, nil
		}
	}

	return nil, fmt.Errorf("인스턴스 %s의 네트워크 인터페이스를 찾을 수 없습니다", instanceID)
}
