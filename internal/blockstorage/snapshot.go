package blockstorage

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListSnapshots() ([]Snapshot, error) {
	resp, err := c.httpClient.Get(c.url("/snapshots/detail"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SnapshotListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("스냅샷 목록 조회 실패: %w", err)
	}

	return result.Snapshots, nil
}

func (c *Client) GetSnapshot(id string) (*Snapshot, error) {
	resp, err := c.httpClient.Get(c.url("/snapshots/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SnapshotResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("스냅샷 조회 실패: %w", err)
	}

	return &result.Snapshot, nil
}

func (c *Client) CreateSnapshot(req *SnapshotCreateRequest) (*Snapshot, error) {
	resp, err := c.httpClient.Post(c.url("/snapshots"), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SnapshotResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("스냅샷 생성 실패: %w", err)
	}

	return &result.Snapshot, nil
}

func (c *Client) DeleteSnapshot(id string) error {
	resp, err := c.httpClient.Delete(c.url("/snapshots/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("스냅샷 삭제 실패: %w", err)
	}

	return nil
}
