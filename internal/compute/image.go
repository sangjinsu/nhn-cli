package compute

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListImages() ([]Image, error) {
	resp, err := c.httpClient.Get(c.url("/images/detail"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ImageListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("이미지 목록 조회 실패: %w", err)
	}

	return result.Images, nil
}

func (c *Client) GetImage(id string) (*Image, error) {
	resp, err := c.httpClient.Get(c.url("/images/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ImageResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("이미지 조회 실패: %w", err)
	}

	return &result.Image, nil
}
