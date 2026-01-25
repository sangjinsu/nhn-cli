package compute

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListFlavors() ([]Flavor, error) {
	resp, err := c.httpClient.Get(c.url("/flavors/detail"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FlavorListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("인스턴스 타입 목록 조회 실패: %w", err)
	}

	return result.Flavors, nil
}

func (c *Client) GetFlavor(id string) (*Flavor, error) {
	resp, err := c.httpClient.Get(c.url("/flavors/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result FlavorResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("인스턴스 타입 조회 실패: %w", err)
	}

	return &result.Flavor, nil
}
