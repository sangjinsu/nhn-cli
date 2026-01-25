package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListRoutingTables() ([]RoutingTable, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/routingtables"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result RoutingTableListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("라우팅 테이블 목록 조회 실패: %w", err)
	}

	return result.RoutingTables, nil
}

func (c *Client) GetRoutingTable(id string) (*RoutingTable, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/routingtables/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result RoutingTableResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("라우팅 테이블 조회 실패: %w", err)
	}

	return &result.RoutingTable, nil
}
