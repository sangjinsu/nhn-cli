package gamebase

import (
	"fmt"

	"nhncli/internal/client"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *client.HTTPClient
	baseURL    string
	appID      string
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

	appID := profile.GamebaseAppID
	secretKey := profile.GamebaseSecretKey
	if len(opts) > 0 {
		if opts[0].AppKey != "" {
			appID = opts[0].AppKey
		}
		if opts[0].SecretKey != "" {
			secretKey = opts[0].SecretKey
		}
	}

	if appID == "" {
		return nil, fmt.Errorf("Gamebase App ID가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if secretKey == "" {
		return nil, fmt.Errorf("Gamebase Secret Key가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}

	headers := map[string]string{
		"X-Secret-Key": secretKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://api-gamebase.nhncloudservice.com",
		appID:      appID,
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
