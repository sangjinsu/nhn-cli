package objectstorage

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"nhncli/internal/auth"
	"nhncli/internal/config"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	debug      bool
}

func getBaseURL(region, tenantID string) string {
	return fmt.Sprintf("https://%s-api-object-storage.nhncloudservice.com/v1/AUTH_%s", region, tenantID)
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
		return nil, fmt.Errorf("인증 실패: %w", err)
	}

	if tenantID == "" {
		tenantID = profile.TenantID
	}
	if tenantID == "" {
		return nil, fmt.Errorf("Tenant ID가 설정되지 않았습니다. 'nhn configure'로 Identity 인증 정보를 설정하세요")
	}

	return &Client{
		httpClient: &http.Client{Timeout: 5 * time.Minute},
		baseURL:    getBaseURL(region, tenantID),
		token:      token,
		debug:      debug,
	}, nil
}

func (c *Client) url(path string) string {
	return c.baseURL + path
}

func (c *Client) doRequest(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if c.debug {
		fmt.Printf("[DEBUG] %s %s\n", method, url)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	req.Header.Set("X-Auth-Token", c.token)
	req.Header.Set("User-Agent", "nhn-cli/0.1.0")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %w", err)
	}

	return resp, nil
}

func (c *Client) checkError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API 오류 (HTTP %d): %s", resp.StatusCode, string(body))
	}
	return nil
}
