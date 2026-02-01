package gamebase

import (
	"fmt"
	"strings"

	"nhncli/internal/client"
)

func (c *Client) GetMember(userID string) (*Member, error) {
	url := fmt.Sprintf("%s/tcgb-member/v1.3/apps/%s/members/%s", c.baseURL, c.appID, userID)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result MemberResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Member, nil
}

func (c *Client) ListMembers(userIDs []string) ([]Member, error) {
	url := fmt.Sprintf("%s/tcgb-member/v1.3/apps/%s/members?userIds=%s",
		c.baseURL, c.appID, strings.Join(userIDs, ","))

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result MemberListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Members, nil
}

func (c *Client) WithdrawMember(userID string) error {
	url := fmt.Sprintf("%s/tcgb-gateway/v1.3/apps/%s/members/%s/withdraw", c.baseURL, c.appID, userID)

	resp, err := c.httpClient.Delete(url, c.getOpts())
	if err != nil {
		return err
	}

	var result BanResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}
