package compute

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListInstances() ([]Instance, error) {
	resp, err := c.httpClient.Get(c.url("/servers/detail"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result InstanceListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("인스턴스 목록 조회 실패: %w", err)
	}

	return result.Servers, nil
}

func (c *Client) GetInstance(id string) (*Instance, error) {
	resp, err := c.httpClient.Get(c.url("/servers/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result InstanceResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("인스턴스 조회 실패: %w", err)
	}

	return &result.Server, nil
}

func (c *Client) CreateInstance(req *InstanceCreateRequest) (*Instance, error) {
	resp, err := c.httpClient.Post(c.url("/servers"), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result InstanceResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("인스턴스 생성 실패: %w", err)
	}

	return &result.Server, nil
}

func (c *Client) DeleteInstance(id string) error {
	resp, err := c.httpClient.Delete(c.url("/servers/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("인스턴스 삭제 실패: %w", err)
	}

	return nil
}

func (c *Client) StartInstance(id string) error {
	req := &InstanceActionRequest{
		Start: &struct{}{},
	}

	resp, err := c.httpClient.Post(c.url("/servers/"+id+"/action"), req, c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("인스턴스 시작 실패: %w", err)
	}

	return nil
}

func (c *Client) StopInstance(id string) error {
	req := &InstanceActionRequest{
		Stop: &struct{}{},
	}

	resp, err := c.httpClient.Post(c.url("/servers/"+id+"/action"), req, c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("인스턴스 중지 실패: %w", err)
	}

	return nil
}

func (c *Client) RebootInstance(id string, hard bool) error {
	rebootType := "SOFT"
	if hard {
		rebootType = "HARD"
	}

	req := &InstanceActionRequest{
		Reboot: &RebootAction{
			Type: rebootType,
		},
	}

	resp, err := c.httpClient.Post(c.url("/servers/"+id+"/action"), req, c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("인스턴스 재부팅 실패: %w", err)
	}

	return nil
}
