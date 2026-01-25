package config

import "fmt"

func (p *ProfileConfig) Validate() error {
	// Identity 인증 필수
	if p.TenantID == "" {
		return fmt.Errorf("Tenant ID가 필요합니다")
	}
	if p.Username == "" {
		return fmt.Errorf("Username이 필요합니다")
	}
	if p.Password == "" {
		return fmt.Errorf("Password가 필요합니다")
	}

	// OAuth 인증 필수
	if p.UserAccessKeyID == "" {
		return fmt.Errorf("User Access Key ID가 필요합니다")
	}
	if p.SecretAccessKey == "" {
		return fmt.Errorf("Secret Access Key가 필요합니다")
	}

	if p.Region == "" {
		return fmt.Errorf("리전 설정이 필요합니다")
	}
	return nil
}

func (p *ProfileConfig) GetAuthTypeDisplay() string {
	return "Identity + OAuth"
}

func (p *ProfileConfig) GetMaskedCredentials() string {
	return p.Username
}
