package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client    *http.Client
	debug     bool
	userAgent string
}

type RequestOption struct {
	Headers map[string]string
	Token   string
}

func NewHTTPClient(debug bool) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		debug:     debug,
		userAgent: "nhn-cli/0.1.0",
	}
}

func (c *HTTPClient) Do(method, url string, body interface{}, opts *RequestOption) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("요청 본문 직렬화 실패: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)

		if c.debug {
			fmt.Printf("[DEBUG] Request Body: %s\n", string(jsonBody))
		}
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if opts != nil {
		if opts.Token != "" {
			req.Header.Set("X-Auth-Token", opts.Token)
		}
		for k, v := range opts.Headers {
			req.Header.Set(k, v)
		}
	}

	if c.debug {
		fmt.Printf("[DEBUG] %s %s\n", method, url)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %w", err)
	}

	return resp, nil
}

func (c *HTTPClient) Get(url string, opts *RequestOption) (*http.Response, error) {
	return c.Do(http.MethodGet, url, nil, opts)
}

func (c *HTTPClient) Post(url string, body interface{}, opts *RequestOption) (*http.Response, error) {
	return c.Do(http.MethodPost, url, body, opts)
}

func (c *HTTPClient) Put(url string, body interface{}, opts *RequestOption) (*http.Response, error) {
	return c.Do(http.MethodPut, url, body, opts)
}

func (c *HTTPClient) Delete(url string, opts *RequestOption) (*http.Response, error) {
	return c.Do(http.MethodDelete, url, nil, opts)
}

func ReadJSON(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("응답 읽기 실패: %w", err)
	}

	if resp.StatusCode >= 400 {
		return ParseAPIError(resp.StatusCode, body)
	}

	if v != nil && len(body) > 0 {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("응답 파싱 실패: %w", err)
		}
	}

	return nil
}
