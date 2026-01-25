package compute

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListAvailabilityZones() ([]AvailabilityZone, error) {
	resp, err := c.httpClient.Get(c.url("/os-availability-zone"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result AvailabilityZoneListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("가용성 영역 목록 조회 실패: %w", err)
	}

	return result.AvailabilityZoneInfo, nil
}
