package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const OAuthEndpoint = "https://oauth.api.nhncloudservice.com/oauth2/token/create"

type OAuthAuthenticator struct {
	userAccessKeyID string
	secretAccessKey string
	debug           bool
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewOAuthAuthenticator(userAccessKeyID, secretAccessKey string, debug bool) *OAuthAuthenticator {
	return &OAuthAuthenticator{
		userAccessKeyID: userAccessKeyID,
		secretAccessKey: secretAccessKey,
		debug:           debug,
	}
}

func (a *OAuthAuthenticator) GetToken() (*TokenResponse, error) {
	// Basic Auth 헤더 생성
	credentials := a.userAccessKeyID + ":" + a.secretAccessKey
	basicAuth := base64.StdEncoding.EncodeToString([]byte(credentials))

	// Form data 생성
	body := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest(http.MethodPost, OAuthEndpoint, body)
	if err != nil {
		return nil, fmt.Errorf("OAuth 요청 생성 실패: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+basicAuth)

	if a.debug {
		fmt.Printf("[DEBUG] POST %s\n", OAuthEndpoint)
		fmt.Printf("[DEBUG] Authorization: Basic %s...\n", basicAuth[:20])
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("OAuth 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("OAuth 응답 읽기 실패: %w", err)
	}

	if a.debug {
		fmt.Printf("[DEBUG] Response Status: %d\n", resp.StatusCode)
		fmt.Printf("[DEBUG] Response Body: %s\n", string(respBody))
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("OAuth 인증 실패 [%d]: %s", resp.StatusCode, string(respBody))
	}

	var oauthResp OAuthTokenResponse
	if err := json.Unmarshal(respBody, &oauthResp); err != nil {
		return nil, fmt.Errorf("OAuth 응답 파싱 실패: %w", err)
	}

	// expires_in은 초 단위 (기본 86400초 = 24시간)
	expiresAt := time.Now().Add(time.Duration(oauthResp.ExpiresIn) * time.Second)

	return &TokenResponse{
		Token:     oauthResp.AccessToken,
		ExpiresAt: expiresAt,
	}, nil
}

func (a *OAuthAuthenticator) GetTenantID() string {
	return ""
}
