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

// GetAuthenticatedToken Identity 인증을 사용하여 토큰을 발급받습니다.
// VPC, Compute 등 OpenStack 기반 API는 Identity 인증만 지원합니다.
func GetAuthenticatedToken(profileName string, profile *config.ProfileConfig, debug bool) (string, string, error) {
	cache, err := LoadCache()
	if err != nil {
		if debug {
			fmt.Printf("[DEBUG] 캐시 로드 실패: %v\n", err)
		}
		cache = &CredentialsCache{Profiles: make(map[string]*CachedToken)}
	}

	if cached, ok := cache.Profiles[profileName]; ok {
		// Identity 토큰은 항상 TenantID가 있음
		// TenantID가 없으면 이전 OAuth 토큰이므로 새로 발급
		if cached.TenantID != "" && cached.ExpiresAt > time.Now().Add(5*time.Minute).Unix() {
			if debug {
				fmt.Println("[DEBUG] 캐시된 Identity 토큰 사용")
			}
			return cached.AccessToken, cached.TenantID, nil
		}
		if debug {
			if cached.TenantID == "" {
				fmt.Println("[DEBUG] 캐시된 토큰에 TenantID 없음, 새 토큰 발급")
			} else {
				fmt.Println("[DEBUG] 캐시된 토큰 만료됨")
			}
		}
	}

	// Identity 인증 사용 (OpenStack API용)
	identityAuth := NewIdentityAuthenticator(
		profile.TenantID,
		profile.Username,
		profile.Password,
		debug,
	)

	token, err := identityAuth.GetToken()
	if err != nil {
		return "", "", fmt.Errorf("Identity 인증 실패: %w", err)
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
