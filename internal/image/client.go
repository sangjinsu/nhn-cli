package image

import (
	"fmt"

	"nhncli/internal/auth"
	"nhncli/internal/client"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *client.HTTPClient
	baseURL    string
	token      string
	debug      bool
}

func getBaseURL(region string) string {
	return fmt.Sprintf("https://%s-api-image-infrastructure.nhncloudservice.com", region)
}

func NewClient(profileName string, region string, debug bool) (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	profile, err := cfg.GetProfile(profileName)
	if err != nil {
		return nil, err
	}

	if region == "" {
		region = profile.Region
	}

	token, _, err := auth.GetAuthenticatedToken(profileName, profile, debug)
	if err != nil {
		return nil, fmt.Errorf("인증 실패: %w", err)
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    getBaseURL(region),
		token:      token,
		debug:      debug,
	}, nil
}

func (c *Client) getOpts() *client.RequestOption {
	return &client.RequestOption{
		Token: c.token,
	}
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s/v2%s", c.baseURL, path)
}
