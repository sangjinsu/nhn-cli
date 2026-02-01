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

func NewClient(profileName string, debug bool) (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	profile, err := cfg.GetProfile(profileName)
	if err != nil {
		return nil, err
	}

	if profile.GamebaseAppID == "" {
		return nil, fmt.Errorf("Gamebase App ID가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if profile.GamebaseSecretKey == "" {
		return nil, fmt.Errorf("Gamebase Secret Key가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}

	headers := map[string]string{
		"X-Secret-Key": profile.GamebaseSecretKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://api-gamebase.nhncloudservice.com",
		appID:      profile.GamebaseAppID,
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
