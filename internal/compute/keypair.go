package compute

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListKeypairs() ([]Keypair, error) {
	resp, err := c.httpClient.Get(c.url("/os-keypairs"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result KeypairListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("키페어 목록 조회 실패: %w", err)
	}

	keypairs := make([]Keypair, len(result.Keypairs))
	for i, kw := range result.Keypairs {
		keypairs[i] = kw.Keypair
	}

	return keypairs, nil
}

func (c *Client) CreateKeypair(name, publicKey string) (*KeypairCreated, error) {
	reqBody := KeypairCreateRequest{
		Keypair: KeypairCreateBody{
			Name:      name,
			PublicKey: publicKey,
		},
	}

	resp, err := c.httpClient.Post(c.url("/os-keypairs"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result KeypairResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("키페어 생성 실패: %w", err)
	}

	return &result.Keypair, nil
}

func (c *Client) DeleteKeypair(name string) error {
	resp, err := c.httpClient.Delete(c.url("/os-keypairs/"+name), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("키페어 삭제 실패: %w", err)
	}

	return nil
}
