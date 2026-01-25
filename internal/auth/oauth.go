package auth

import (
	"fmt"
	"time"

	"nhncli/internal/client"
)

const OAuthEndpoint = "https://oauth.api.nhncloudservice.com/oauth2/token/create"

type OAuthAuthenticator struct {
	userAccessKeyID string
	secretAccessKey string
	debug           bool
	httpClient      *client.HTTPClient
}

func NewOAuthAuthenticator(userAccessKeyID, secretAccessKey string, debug bool) *OAuthAuthenticator {
	return &OAuthAuthenticator{
		userAccessKeyID: userAccessKeyID,
		secretAccessKey: secretAccessKey,
		debug:           debug,
		httpClient:      client.NewHTTPClient(debug),
	}
}

func (a *OAuthAuthenticator) GetToken() (*TokenResponse, error) {
	reqBody := OAuthRequest{
		Auth: OAuthAuth{
			UserAccessKeyID: a.userAccessKeyID,
			SecretAccessKey: a.secretAccessKey,
		},
	}

	resp, err := a.httpClient.Post(OAuthEndpoint, reqBody, nil)
	if err != nil {
		return nil, err
	}

	var oauthResp OAuthResponse
	if err := client.ReadJSON(resp, &oauthResp); err != nil {
		return nil, fmt.Errorf("OAuth 응답 처리 실패: %w", err)
	}

	expiresAt, err := time.Parse(time.RFC3339, oauthResp.Token.ExpiresAt)
	if err != nil {
		expiresAt = time.Now().Add(12 * time.Hour)
	}

	return &TokenResponse{
		Token:     oauthResp.Token.ID,
		ExpiresAt: expiresAt,
	}, nil
}

func (a *OAuthAuthenticator) GetTenantID() string {
	return ""
}
