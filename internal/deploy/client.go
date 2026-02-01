package deploy

import (
	"fmt"

	"nhncli/internal/client"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *client.HTTPClient
	baseURL    string
	appKey     string
	headers    map[string]string
	debug      bool
}

func NewClient(profileName string, debug bool) (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	profile, err := cfg.GetProfile(profileName)
	if err != nil {
		return nil, err
	}

	if profile.DeployAppKey == "" {
		return nil, fmt.Errorf("Deploy AppKey가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if profile.UserAccessKeyID == "" || profile.SecretAccessKey == "" {
		return nil, fmt.Errorf("User Access Key ID/Secret이 설정되지 않았습니다")
	}

	headers := map[string]string{
		"X-TC-AUTHENTICATION-ID":     profile.UserAccessKeyID,
		"X-TC-AUTHENTICATION-SECRET": profile.SecretAccessKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://api-tcd.nhncloudservice.com",
		appKey:     profile.DeployAppKey,
		headers:    headers,
		debug:      debug,
	}, nil
}

func (c *Client) getOpts() *client.RequestOption {
	return &client.RequestOption{
		Headers: c.headers,
	}
}

func checkResponse(header ResponseHeader) error {
	if !header.IsSuccessful {
		return fmt.Errorf("API 오류 (코드: %d): %s", header.ResultCode, header.ResultMessage)
	}
	return nil
}
