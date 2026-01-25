package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListVPCs() ([]VPC, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/vpcs"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VPCListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("VPC 목록 조회 실패: %w", err)
	}

	return result.VPCs, nil
}

func (c *Client) GetVPC(id string) (*VPC, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/vpcs/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VPCResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("VPC 조회 실패: %w", err)
	}

	return &result.VPC, nil
}

func (c *Client) CreateVPC(name, cidr string) (*VPC, error) {
	reqBody := VPCCreateRequest{
		VPC: VPCCreateBody{
			Name:   name,
			CIDRv4: cidr,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/vpcs"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VPCResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("VPC 생성 실패: %w", err)
	}

	return &result.VPC, nil
}

func (c *Client) UpdateVPC(id, name, cidr string) (*VPC, error) {
	reqBody := VPCUpdateRequest{
		VPC: VPCUpdateBody{},
	}

	if name != "" {
		reqBody.VPC.Name = name
	}
	if cidr != "" {
		reqBody.VPC.CIDRv4 = cidr
	}

	resp, err := c.httpClient.Put(c.url("/v2.0/vpcs/"+id), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result VPCResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("VPC 수정 실패: %w", err)
	}

	return &result.VPC, nil
}

func (c *Client) DeleteVPC(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/vpcs/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("VPC 삭제 실패: %w", err)
	}

	return nil
}
