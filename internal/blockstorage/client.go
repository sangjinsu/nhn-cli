package blockstorage

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
	return fmt.Sprintf("https://%s-api-block-storage-infrastructure.nhncloudservice.com", region)
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
		return nil, fmt.Errorf("Identity 인증 실패: %w", err)
	}

	if tenantID == "" {
		tenantID = profile.TenantID
	}
	if tenantID == "" {
		return nil, fmt.Errorf("Tenant ID가 설정되지 않았습니다. 'nhn configure'로 Identity 인증 정보를 설정하세요")
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
