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

type ClientOption struct {
	AppKey    string
	SecretKey string
}

func NewClient(profileName string, debug bool, opts ...ClientOption) (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	profile, err := cfg.GetProfile(profileName)
	if err != nil {
		return nil, err
	}

	appKey := profile.CDNAppKey
	secretKey := profile.CDNSecretKey
	if len(opts) > 0 {
		if opts[0].AppKey != "" {
			appKey = opts[0].AppKey
		}
		if opts[0].SecretKey != "" {
			secretKey = opts[0].SecretKey
		}
	}

	if appKey == "" {
		return nil, fmt.Errorf("CDN AppKey가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if secretKey == "" {
		return nil, fmt.Errorf("CDN Secret Key가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}

	headers := map[string]string{
		"Authorization": secretKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://cdn.api.nhncloudservice.com",
		appKey:     appKey,
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
