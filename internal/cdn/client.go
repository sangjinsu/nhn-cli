package cdn

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

	if profile.CDNAppKey == "" {
		return nil, fmt.Errorf("CDN AppKey가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if profile.CDNSecretKey == "" {
		return nil, fmt.Errorf("CDN Secret Key가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}

	headers := map[string]string{
		"Authorization": profile.CDNSecretKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://cdn.api.nhncloudservice.com",
		appKey:     profile.CDNAppKey,
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
