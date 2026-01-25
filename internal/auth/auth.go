package auth

import (
	"fmt"
	"time"

	"nhncli/internal/config"
)

type Authenticator interface {
	GetToken() (*TokenResponse, error)
	GetTenantID() string
}

func NewAuthenticator(profile *config.ProfileConfig, debug bool) (Authenticator, error) {
	switch profile.AuthType {
	case config.AuthTypeOAuth:
		return NewOAuthAuthenticator(profile.UserAccessKeyID, profile.SecretAccessKey, debug), nil
	case config.AuthTypeIdentity:
		return NewIdentityAuthenticator(profile.TenantID, profile.Username, profile.Password, debug), nil
	default:
		return nil, fmt.Errorf("지원하지 않는 인증 방식입니다: %s", profile.AuthType)
	}
}

func GetAuthenticatedToken(profileName string, profile *config.ProfileConfig, debug bool) (string, string, error) {
	cache, err := LoadCache()
	if err != nil {
		if debug {
			fmt.Printf("[DEBUG] 캐시 로드 실패: %v\n", err)
		}
		cache = &CredentialsCache{Profiles: make(map[string]*CachedToken)}
	}

	if cached, ok := cache.Profiles[profileName]; ok {
		if cached.ExpiresAt > time.Now().Add(5*time.Minute).Unix() {
			if debug {
				fmt.Println("[DEBUG] 캐시된 토큰 사용")
			}
			return cached.AccessToken, cached.TenantID, nil
		}
		if debug {
			fmt.Println("[DEBUG] 캐시된 토큰 만료됨")
		}
	}

	auth, err := NewAuthenticator(profile, debug)
	if err != nil {
		return "", "", err
	}

	token, err := auth.GetToken()
	if err != nil {
		return "", "", fmt.Errorf("토큰 발급 실패: %w", err)
	}

	cache.Profiles[profileName] = &CachedToken{
		AccessToken: token.Token,
		ExpiresAt:   token.ExpiresAt.Unix(),
		TenantID:    token.TenantID,
	}

	if err := cache.Save(); err != nil {
		if debug {
			fmt.Printf("[DEBUG] 캐시 저장 실패: %v\n", err)
		}
	}

	return token.Token, token.TenantID, nil
}
