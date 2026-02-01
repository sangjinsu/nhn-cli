package auth

import (
	"os"
	"testing"
	"time"
)

func TestLoadCacheNonExistent(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cache, err := LoadCache()
	if err != nil {
		t.Fatalf("LoadCache() error = %v", err)
	}
	if cache == nil {
		t.Fatal("LoadCache() returned nil")
	}
	if len(cache.Profiles) != 0 {
		t.Errorf("expected 0 profiles, got %d", len(cache.Profiles))
	}
}

func TestSaveAndLoadCache(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// Create .nhn directory
	os.MkdirAll(tmpDir+"/.nhn", 0700)

	cache := &CredentialsCache{
		Profiles: map[string]*CachedToken{
			"default": {
				AccessToken: "token-abc",
				ExpiresAt:   time.Now().Add(1 * time.Hour).Unix(),
				TenantID:    "tenant-123",
			},
		},
	}

	if err := cache.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := LoadCache()
	if err != nil {
		t.Fatalf("LoadCache() error = %v", err)
	}

	cached, ok := loaded.Profiles["default"]
	if !ok {
		t.Fatal("profile 'default' not found in cache")
	}
	if cached.AccessToken != "token-abc" {
		t.Errorf("AccessToken = %q, want %q", cached.AccessToken, "token-abc")
	}
	if cached.TenantID != "tenant-123" {
		t.Errorf("TenantID = %q, want %q", cached.TenantID, "tenant-123")
	}
}

func TestClearCache(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	os.MkdirAll(tmpDir+"/.nhn", 0700)

	cache := &CredentialsCache{
		Profiles: map[string]*CachedToken{
			"default": {AccessToken: "token-1", ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), TenantID: "t1"},
			"prod":    {AccessToken: "token-2", ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), TenantID: "t2"},
		},
	}
	cache.Save()

	if err := ClearCache("default"); err != nil {
		t.Fatalf("ClearCache() error = %v", err)
	}

	loaded, _ := LoadCache()
	if _, ok := loaded.Profiles["default"]; ok {
		t.Error("default profile should have been cleared")
	}
	if _, ok := loaded.Profiles["prod"]; !ok {
		t.Error("prod profile should still exist")
	}
}

func TestClearAllCache(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	os.MkdirAll(tmpDir+"/.nhn", 0700)

	cache := &CredentialsCache{
		Profiles: map[string]*CachedToken{
			"default": {AccessToken: "token-1", ExpiresAt: time.Now().Add(1 * time.Hour).Unix()},
		},
	}
	cache.Save()

	if err := ClearAllCache(); err != nil {
		t.Fatalf("ClearAllCache() error = %v", err)
	}

	// Loading should return empty cache (file deleted)
	loaded, err := LoadCache()
	if err != nil {
		t.Fatalf("LoadCache() error = %v", err)
	}
	if len(loaded.Profiles) != 0 {
		t.Errorf("expected 0 profiles after ClearAllCache, got %d", len(loaded.Profiles))
	}
}
