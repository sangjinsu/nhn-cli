package cdn

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) CreateAuthToken(req *AuthTokenCreateRequest) (string, error) {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/auth-token", c.baseURL, c.appKey)

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return "", err
	}

	var result AuthTokenResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return "", err
	}

	if err := checkResponse(result.Header); err != nil {
		return "", err
	}

	return result.Token, nil
}
