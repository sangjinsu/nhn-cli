package compute

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
	tenantID   string
	debug      bool
}

func getBaseURL(region string) string {
	return fmt.Sprintf("https://%s-api-instance-infrastructure.nhncloudservice.com", region)
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

	token, tenantID, err := auth.GetAuthenticatedToken(profileName, profile, debug)
	if err != nil {
		return nil, err
	}

	if tenantID == "" && profile.TenantID != "" {
		tenantID = profile.TenantID
	}

	if tenantID == "" {
		return nil, fmt.Errorf("Tenant ID가 필요합니다. Identity 인증을 사용하거나 프로필에 Tenant ID를 설정하세요")
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    getBaseURL(region),
		token:      token,
		tenantID:   tenantID,
		debug:      debug,
	}, nil
}

func (c *Client) getOpts() *client.RequestOption {
	return &client.RequestOption{
		Token: c.token,
	}
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s/v2/%s%s", c.baseURL, c.tenantID, path)
}
