package objectstorage

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListObjects(container string) ([]Object, error) {
	resp, err := c.doRequest("GET", c.url("/"+container+"?format=json"), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return nil, fmt.Errorf("오브젝트 목록 조회 실패: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	if len(body) == 0 {
		return []Object{}, nil
	}

	var objects []Object
	if err := json.Unmarshal(body, &objects); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return objects, nil
}

func (c *Client) GetObjectMetadata(container, object string) (*ObjectMetadata, error) {
	resp, err := c.doRequest("HEAD", c.url("/"+container+"/"+object), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return nil, fmt.Errorf("오브젝트 메타데이터 조회 실패: %w", err)
	}

	return &ObjectMetadata{
		ContentLength: resp.Header.Get("Content-Length"),
		ContentType:   resp.Header.Get("Content-Type"),
		ETag:          resp.Header.Get("ETag"),
		LastModified:  resp.Header.Get("Last-Modified"),
	}, nil
}

func (c *Client) UploadObject(container, objectName string, reader io.Reader, contentType string) error {
	headers := map[string]string{}
	if contentType != "" {
		headers["Content-Type"] = contentType
	} else {
		headers["Content-Type"] = "application/octet-stream"
	}

	resp, err := c.doRequest("PUT", c.url("/"+container+"/"+objectName), reader, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return fmt.Errorf("오브젝트 업로드 실패: %w", err)
	}

	return nil
}

func (c *Client) DownloadObject(container, object string) (io.ReadCloser, error) {
	resp, err := c.doRequest("GET", c.url("/"+container+"/"+object), nil, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("오브젝트 다운로드 실패 (HTTP %d): %s", resp.StatusCode, string(body))
	}

	return resp.Body, nil
}

func (c *Client) DeleteObject(container, object string) error {
	resp, err := c.doRequest("DELETE", c.url("/"+container+"/"+object), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return fmt.Errorf("오브젝트 삭제 실패: %w", err)
	}

	return nil
}
