package objectstorage

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListContainers() ([]Container, error) {
	resp, err := c.doRequest("GET", c.url("?format=json"), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return nil, fmt.Errorf("컨테이너 목록 조회 실패: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	if len(body) == 0 {
		return []Container{}, nil
	}

	var containers []Container
	if err := json.Unmarshal(body, &containers); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return containers, nil
}

func (c *Client) GetContainerMetadata(name string) (*ContainerMetadata, error) {
	resp, err := c.doRequest("HEAD", c.url("/"+name), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return nil, fmt.Errorf("컨테이너 메타데이터 조회 실패: %w", err)
	}

	return &ContainerMetadata{
		ObjectCount: resp.Header.Get("X-Container-Object-Count"),
		BytesUsed:   resp.Header.Get("X-Container-Bytes-Used"),
		ReadACL:     resp.Header.Get("X-Container-Read"),
		WriteACL:    resp.Header.Get("X-Container-Write"),
	}, nil
}

func (c *Client) CreateContainer(name string) error {
	resp, err := c.doRequest("PUT", c.url("/"+name), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return fmt.Errorf("컨테이너 생성 실패: %w", err)
	}

	return nil
}

func (c *Client) DeleteContainer(name string) error {
	resp, err := c.doRequest("DELETE", c.url("/"+name), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkError(resp); err != nil {
		return fmt.Errorf("컨테이너 삭제 실패: %w", err)
	}

	return nil
}
