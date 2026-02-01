package blockstorage

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListVolumes() ([]Volume, error) {
	resp, err := c.httpClient.Get(c.url("/volumes/detail"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VolumeListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("볼륨 목록 조회 실패: %w", err)
	}

	return result.Volumes, nil
}

func (c *Client) GetVolume(id string) (*Volume, error) {
	resp, err := c.httpClient.Get(c.url("/volumes/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VolumeResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("볼륨 조회 실패: %w", err)
	}

	return &result.Volume, nil
}

func (c *Client) CreateVolume(req *VolumeCreateRequest) (*Volume, error) {
	resp, err := c.httpClient.Post(c.url("/volumes"), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VolumeResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("볼륨 생성 실패: %w", err)
	}

	return &result.Volume, nil
}

func (c *Client) DeleteVolume(id string) error {
	resp, err := c.httpClient.Delete(c.url("/volumes/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("볼륨 삭제 실패: %w", err)
	}

	return nil
}

func (c *Client) ListVolumeTypes() ([]VolumeType, error) {
	resp, err := c.httpClient.Get(c.url("/types"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VolumeTypeListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("볼륨 타입 목록 조회 실패: %w", err)
	}

	return result.VolumeTypes, nil
}
