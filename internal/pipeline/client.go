package pipeline

import (
	"fmt"

	"nhncli/internal/client"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *client.HTTPClient
	baseURL    string
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

	if profile.PipelineAppKey == "" {
		return nil, fmt.Errorf("Pipeline AppKey가 설정되지 않았습니다. 'nhn configure'로 설정하세요")
	}
	if profile.UserAccessKeyID == "" || profile.SecretAccessKey == "" {
		return nil, fmt.Errorf("User Access Key ID/Secret이 설정되지 않았습니다")
	}

	headers := map[string]string{
		"X-NHN-REGION":               profile.Region,
		"X-NHN-APPKEY":               profile.PipelineAppKey,
		"X-TC-AUTHENTICATION-ID":     profile.UserAccessKeyID,
		"X-TC-AUTHENTICATION-SECRET": profile.SecretAccessKey,
	}

	return &Client{
		httpClient: client.NewHTTPClient(debug),
		baseURL:    "https://kr1-pipeline.api.nhncloudservice.com",
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
