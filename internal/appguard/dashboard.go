package appguard

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) GetDashboard(targetDate string, os int, targetType int) ([]DashboardEntry, error) {
	url := fmt.Sprintf("%s/appguard/v1.0/appkeys/%s/dashboard?targetDate=%s&os=%d&targetType=%d",
		c.baseURL, c.appKey, targetDate, os, targetType)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result DashboardResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Data, nil
}
