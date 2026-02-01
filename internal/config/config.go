package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDirName  = ".nhn"
	ConfigFileName = "config.json"
)

type ProfileConfig struct {
	// Identity 인증 (OpenStack API용) - 필수
	TenantID string `json:"tenant_id"`
	Username string `json:"username"`
	Password string `json:"password"`

	// OAuth 인증 (NHN Cloud 고유 API용) - 선택
	UserAccessKeyID string `json:"user_access_key_id,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`

	// DNS Plus (AppKey 기반 인증) - 선택
	AppKey string `json:"app_key,omitempty"`

	// Pipeline
	PipelineAppKey string `json:"pipeline_app_key,omitempty"`

	// Deploy
	DeployAppKey string `json:"deploy_app_key,omitempty"`

	// CDN
	CDNAppKey    string `json:"cdn_app_key,omitempty"`
	CDNSecretKey string `json:"cdn_secret_key,omitempty"`

	// AppGuard
	AppGuardAppKey string `json:"appguard_app_key,omitempty"`

	// Gamebase
	GamebaseAppID     string `json:"gamebase_app_id,omitempty"`
	GamebaseSecretKey string `json:"gamebase_secret_key,omitempty"`

	Region string `json:"region"`
}

type Config struct {
	Profiles map[string]*ProfileConfig `json:"profiles"`
}

func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("홈 디렉토리를 찾을 수 없습니다: %w", err)
	}
	return filepath.Join(homeDir, ConfigDirName), nil
}

func GetConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, ConfigFileName), nil
}

func EnsureConfigDir() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(configDir, 0700)
}

func Load() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Profiles: make(map[string]*ProfileConfig),
			}, nil
		}
		return nil, fmt.Errorf("설정 파일 읽기 실패: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("설정 파일 파싱 실패: %w", err)
	}

	if config.Profiles == nil {
		config.Profiles = make(map[string]*ProfileConfig)
	}

	return &config, nil
}

func (c *Config) Save() error {
	if err := EnsureConfigDir(); err != nil {
		return err
	}

	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("설정 직렬화 실패: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("설정 파일 저장 실패: %w", err)
	}

	return nil
}

func (c *Config) GetProfile(name string) (*ProfileConfig, error) {
	profile, ok := c.Profiles[name]
	if !ok {
		return nil, fmt.Errorf("프로필 '%s'을(를) 찾을 수 없습니다. 'nhn configure'로 설정하세요", name)
	}
	return profile, nil
}

func (c *Config) SetProfile(name string, profile *ProfileConfig) {
	c.Profiles[name] = profile
}

func (c *Config) DeleteProfile(name string) {
	delete(c.Profiles, name)
}

func (c *Config) ListProfiles() []string {
	profiles := make([]string, 0, len(c.Profiles))
	for name := range c.Profiles {
		profiles = append(profiles, name)
	}
	return profiles
}
