package auth

import "time"

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	TenantID  string    `json:"tenant_id,omitempty"`
}

type OAuthRequest struct {
	Auth OAuthAuth `json:"auth"`
}

type OAuthAuth struct {
	UserAccessKeyID string `json:"user_access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type OAuthResponse struct {
	Token struct {
		ID        string `json:"id"`
		ExpiresAt string `json:"expires_at"`
	} `json:"token"`
}

type IdentityRequest struct {
	Auth IdentityAuth `json:"auth"`
}

type IdentityAuth struct {
	TenantID            string                   `json:"tenantId"`
	PasswordCredentials IdentityPasswordCredentials `json:"passwordCredentials"`
}

type IdentityPasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IdentityResponse struct {
	Access struct {
		Token struct {
			ID      string `json:"id"`
			Expires string `json:"expires"`
			Tenant  struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				GroupID     string `json:"groupId"`
				Description string `json:"description"`
				Enabled     bool   `json:"enabled"`
				ProjectID   string `json:"project_id"`
			} `json:"tenant"`
		} `json:"token"`
		ServiceCatalog []ServiceCatalog `json:"serviceCatalog"`
	} `json:"access"`
}

type ServiceCatalog struct {
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	PublicURL string `json:"publicURL"`
	Region    string `json:"region"`
}

type CachedToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	TenantID    string `json:"tenant_id,omitempty"`
}

type CredentialsCache struct {
	Profiles map[string]*CachedToken `json:"profiles"`
}
