package loadbalancer

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListListeners() ([]Listener, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/lbaas/listeners"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ListenerListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("리스너 목록 조회 실패: %w", err)
	}

	return result.Listeners, nil
}

func (c *Client) GetListener(id string) (*Listener, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/lbaas/listeners/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ListenerResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("리스너 조회 실패: %w", err)
	}

	return &result.Listener, nil
}

func (c *Client) CreateListener(req *ListenerCreateRequest) (*Listener, error) {
	resp, err := c.httpClient.Post(c.url("/v2.0/lbaas/listeners"), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ListenerResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("리스너 생성 실패: %w", err)
	}

	return &result.Listener, nil
}

func (c *Client) DeleteListener(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/lbaas/listeners/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("리스너 삭제 실패: %w", err)
	}

	return nil
}
