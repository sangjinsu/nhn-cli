package auth

import (
	"fmt"
	"time"

	"nhncli/internal/client"
)

const IdentityEndpoint = "https://api-identity-infrastructure.nhncloudservice.com/v2.0/tokens"

type IdentityAuthenticator struct {
	tenantID   string
	username   string
	password   string
	debug      bool
	httpClient *client.HTTPClient
}

func NewIdentityAuthenticator(tenantID, username, password string, debug bool) *IdentityAuthenticator {
	return &IdentityAuthenticator{
		tenantID:   tenantID,
		username:   username,
		password:   password,
		debug:      debug,
		httpClient: client.NewHTTPClient(debug),
	}
}

func (a *IdentityAuthenticator) GetToken() (*TokenResponse, error) {
	reqBody := IdentityRequest{
		Auth: IdentityAuth{
			TenantID: a.tenantID,
			PasswordCredentials: IdentityPasswordCredentials{
				Username: a.username,
				Password: a.password,
			},
		},
	}

	resp, err := a.httpClient.Post(IdentityEndpoint, reqBody, nil)
	if err != nil {
		return nil, err
	}

	var identityResp IdentityResponse
	if err := client.ReadJSON(resp, &identityResp); err != nil {
		return nil, fmt.Errorf("Identity 응답 처리 실패: %w", err)
	}

	expiresAt, err := time.Parse(time.RFC3339, identityResp.Access.Token.Expires)
	if err != nil {
		expiresAt = time.Now().Add(12 * time.Hour)
	}

	return &TokenResponse{
		Token:     identityResp.Access.Token.ID,
		ExpiresAt: expiresAt,
		TenantID:  identityResp.Access.Token.Tenant.ID,
	}, nil
}

func (a *IdentityAuthenticator) GetTenantID() string {
	return a.tenantID
}
