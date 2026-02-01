package gamebase

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) GetLaunching() (*LaunchingInfo, error) {
	url := fmt.Sprintf("%s/tcgb-launching/v1.3/apps/%s/launching", c.baseURL, c.appID)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result LaunchingResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Launching, nil
}
