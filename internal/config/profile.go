package config

import "fmt"

func (p *ProfileConfig) Validate() error {
	switch p.AuthType {
	case AuthTypeOAuth:
		if p.UserAccessKeyID == "" {
			return fmt.Errorf("User Access Key ID가 필요합니다")
		}
		if p.SecretAccessKey == "" {
			return fmt.Errorf("Secret Access Key가 필요합니다")
		}
	case AuthTypeIdentity:
		if p.TenantID == "" {
			return fmt.Errorf("Tenant ID가 필요합니다")
		}
		if p.Username == "" {
			return fmt.Errorf("Username이 필요합니다")
		}
		if p.Password == "" {
			return fmt.Errorf("Password가 필요합니다")
		}
	default:
		return fmt.Errorf("지원하지 않는 인증 방식입니다: %s", p.AuthType)
	}

	if p.Region == "" {
		return fmt.Errorf("리전 설정이 필요합니다")
	}

	return nil
}

func (p *ProfileConfig) GetAuthTypeDisplay() string {
	switch p.AuthType {
	case AuthTypeOAuth:
		return "OAuth"
	case AuthTypeIdentity:
		return "Identity"
	default:
		return string(p.AuthType)
	}
}

func (p *ProfileConfig) GetMaskedCredentials() string {
	switch p.AuthType {
	case AuthTypeOAuth:
		if len(p.UserAccessKeyID) > 8 {
			return p.UserAccessKeyID[:4] + "****" + p.UserAccessKeyID[len(p.UserAccessKeyID)-4:]
		}
		return "****"
	case AuthTypeIdentity:
		return p.Username
	default:
		return ""
	}
}
