package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadNonExistentConfig(t *testing.T) {
	// Save and restore HOME
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}
	if len(cfg.Profiles) != 0 {
		t.Errorf("expected 0 profiles, got %d", len(cfg.Profiles))
	}
}

func TestSaveAndLoad(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := &Config{
		Profiles: map[string]*ProfileConfig{
			"default": {
				TenantID: "tenant-123",
				Username: "user@example.com",
				Password: "password",
				Region:   "KR1",
			},
			"prod": {
				UserAccessKeyID: "key-456",
				SecretAccessKey: "secret-789",
				Region:          "KR2",
			},
		},
	}

	if err := cfg.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Verify file exists
	configPath := filepath.Join(tmpDir, ConfigDirName, ConfigFileName)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	// Load back
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(loaded.Profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(loaded.Profiles))
	}

	def, err := loaded.GetProfile("default")
	if err != nil {
		t.Fatalf("GetProfile('default') error = %v", err)
	}
	if def.TenantID != "tenant-123" {
		t.Errorf("TenantID = %q, want %q", def.TenantID, "tenant-123")
	}
	if def.Region != "KR1" {
		t.Errorf("Region = %q, want %q", def.Region, "KR1")
	}
}

func TestGetProfileNotFound(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]*ProfileConfig{},
	}

	_, err := cfg.GetProfile("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent profile")
	}
}

func TestSetAndDeleteProfile(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]*ProfileConfig{},
	}

	cfg.SetProfile("test", &ProfileConfig{Region: "JP1"})
	if len(cfg.Profiles) != 1 {
		t.Errorf("expected 1 profile, got %d", len(cfg.Profiles))
	}

	p, err := cfg.GetProfile("test")
	if err != nil {
		t.Fatalf("GetProfile() error = %v", err)
	}
	if p.Region != "JP1" {
		t.Errorf("Region = %q, want %q", p.Region, "JP1")
	}

	cfg.DeleteProfile("test")
	if len(cfg.Profiles) != 0 {
		t.Errorf("expected 0 profiles after delete, got %d", len(cfg.Profiles))
	}
}

func TestListProfiles(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]*ProfileConfig{
			"a": {Region: "KR1"},
			"b": {Region: "KR2"},
		},
	}

	profiles := cfg.ListProfiles()
	if len(profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(profiles))
	}
}

func TestConfigDirPermissions(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	if err := EnsureConfigDir(); err != nil {
		t.Fatalf("EnsureConfigDir() error = %v", err)
	}

	configDir := filepath.Join(tmpDir, ConfigDirName)
	info, err := os.Stat(configDir)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}
	if !info.IsDir() {
		t.Error("config dir is not a directory")
	}
	// Check permissions (0700)
	perm := info.Mode().Perm()
	if perm != 0700 {
		t.Errorf("config dir permissions = %o, want 0700", perm)
	}
}
