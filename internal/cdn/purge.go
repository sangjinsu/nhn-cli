package cdn

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) Purge(domain string, purgeType string, items []string) error {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/services/%s/purge", c.baseURL, c.appKey, domain)

	req := &PurgeRequest{
		PurgeType: purgeType,
		Items:     items,
	}

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return err
	}

	var result PurgeResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}
