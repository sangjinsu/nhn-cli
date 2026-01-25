package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"nhncli/internal/config"
)

const CredentialsFileName = "credentials.json"

func GetCredentialsPath() (string, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, CredentialsFileName), nil
}

func LoadCache() (*CredentialsCache, error) {
	credPath, err := GetCredentialsPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(credPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &CredentialsCache{
				Profiles: make(map[string]*CachedToken),
			}, nil
		}
		return nil, fmt.Errorf("자격 증명 파일 읽기 실패: %w", err)
	}

	var cache CredentialsCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, fmt.Errorf("자격 증명 파일 파싱 실패: %w", err)
	}

	if cache.Profiles == nil {
		cache.Profiles = make(map[string]*CachedToken)
	}

	return &cache, nil
}

func (c *CredentialsCache) Save() error {
	if err := config.EnsureConfigDir(); err != nil {
		return err
	}

	credPath, err := GetCredentialsPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("자격 증명 직렬화 실패: %w", err)
	}

	if err := os.WriteFile(credPath, data, 0600); err != nil {
		return fmt.Errorf("자격 증명 파일 저장 실패: %w", err)
	}

	return nil
}

func ClearCache(profileName string) error {
	cache, err := LoadCache()
	if err != nil {
		return err
	}

	delete(cache.Profiles, profileName)
	return cache.Save()
}

func ClearAllCache() error {
	credPath, err := GetCredentialsPath()
	if err != nil {
		return err
	}

	if err := os.Remove(credPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("자격 증명 파일 삭제 실패: %w", err)
	}

	return nil
}
