package dns

import (
	"fmt"

	"nhncli/internal/client"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *client.HTTPClient
	baseURL    string
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

	if profile.AppKey == "" {
		return nil, fmt.Errorf("DNS AppKey가 설정되지 않았습니다. 'nhn configure'로 AppKey를 설정하세요")
	}

	baseURL := fmt.Sprintf("https://dnsplus.api.nhncloudservice.com/dnsplus/v1.0/appkeys/%s", profile.AppKey)

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    baseURL,
		debug:      debug,
	}, nil
}

func (c *Client) url(path string) string {
	return c.baseURL + path
}

func (c *Client) getOpts() *client.RequestOption {
	return &client.RequestOption{}
}

func checkResponse(header ResponseHeader) error {
	if !header.IsSuccessful {
		return fmt.Errorf("API 오류 (코드: %d): %s", header.ResultCode, header.ResultMessage)
	}
	return nil
}
