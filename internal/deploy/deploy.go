package deploy

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ExecuteDeploy(req *DeployExecuteRequest) (*DeployResult, error) {
	url := fmt.Sprintf("%s/api/v1.0/appkeys/%s/deployments", c.baseURL, c.appKey)

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result DeployExecuteResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Body, nil
}
