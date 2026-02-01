package gamebase

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ValidateToken(userID, accessToken string) (bool, error) {
	url := fmt.Sprintf("%s/tcgb-gateway/v1.3/apps/%s/members/%s/tokens/%s",
		c.baseURL, c.appID, userID, accessToken)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return false, err
	}

	var result TokenValidateResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return false, err
	}

	if err := checkResponse(result.Header); err != nil {
		return false, err
	}

	return result.Valid, nil
}
