package gamebase

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) CreateBan(req *BanCreateRequest) error {
	url := fmt.Sprintf("%s/tcgb-member/v1.3/apps/%s/members/ban", c.baseURL, c.appID)

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return err
	}

	var result BanResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}

func (c *Client) ListBans(userID string) ([]BanInfo, error) {
	url := fmt.Sprintf("%s/tcgb-member/v1.3/apps/%s/members/bans?userId=%s", c.baseURL, c.appID, userID)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result BanListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Bans, nil
}

func (c *Client) ReleaseBan(userID string) error {
	url := fmt.Sprintf("%s/tcgb-member/v1.3/apps/%s/members/ban/release", c.baseURL, c.appID)

	req := &BanReleaseRequest{UserID: userID}

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return err
	}

	var result BanResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}
