package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.Header.Get("X-Auth-Token") != "test-token" {
			t.Errorf("X-Auth-Token = %q, want %q", r.Header.Get("X-Auth-Token"), "test-token")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, err := client.Get(server.URL, &RequestOption{Token: "test-token"})
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	var result map[string]string
	if err := ReadJSON(resp, &result); err != nil {
		t.Fatalf("ReadJSON() error = %v", err)
	}
	if result["status"] != "ok" {
		t.Errorf("status = %q, want %q", result["status"], "ok")
	}
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", r.Header.Get("Content-Type"))
		}

		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "test" {
			t.Errorf("body.name = %q, want %q", body["name"], "test")
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": "123"})
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, err := client.Post(server.URL, map[string]string{"name": "test"}, &RequestOption{Token: "tok"})
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}

	var result map[string]string
	if err := ReadJSON(resp, &result); err != nil {
		t.Fatalf("ReadJSON() error = %v", err)
	}
	if result["id"] != "123" {
		t.Errorf("id = %q, want %q", result["id"], "123")
	}
}

func TestReadJSONError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"itemNotFound": map[string]interface{}{
				"message": "Resource not found",
				"code":    404,
			},
		})
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, err := client.Get(server.URL, nil)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	err = ReadJSON(resp, nil)
	if err == nil {
		t.Fatal("expected error for 404 response")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
}

func TestDeleteMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, err := client.Delete(server.URL, &RequestOption{Token: "tok"})
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if err := ReadJSON(resp, nil); err != nil {
		t.Fatalf("ReadJSON() error = %v", err)
	}
}

func TestPutMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"updated": "true"})
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, err := client.Put(server.URL, map[string]string{"name": "new"}, &RequestOption{Token: "tok"})
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	var result map[string]string
	if err := ReadJSON(resp, &result); err != nil {
		t.Fatalf("ReadJSON() error = %v", err)
	}
	if result["updated"] != "true" {
		t.Errorf("updated = %q, want %q", result["updated"], "true")
	}
}

func TestUserAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != "nhn-cli/0.1.0" {
			t.Errorf("User-Agent = %q, want %q", r.Header.Get("User-Agent"), "nhn-cli/0.1.0")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewHTTPClient(false)
	resp, _ := client.Get(server.URL, nil)
	ReadJSON(resp, nil)
}
