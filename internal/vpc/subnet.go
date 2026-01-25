package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListSubnets(vpcID string) ([]Subnet, error) {
	url := c.url("/v2.0/vpcsubnets")
	if vpcID != "" {
		url += "?vpc_id=" + vpcID
	}

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SubnetListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("서브넷 목록 조회 실패: %w", err)
	}

	return result.VPCSubnets, nil
}

func (c *Client) GetSubnet(id string) (*Subnet, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/vpcsubnets/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SubnetResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("서브넷 조회 실패: %w", err)
	}

	return &result.VPCSubnet, nil
}

func (c *Client) CreateSubnet(vpcID, name, cidr string) (*Subnet, error) {
	reqBody := SubnetCreateRequest{
		VPCSubnet: SubnetCreateBody{
			VPCID: vpcID,
			Name:  name,
			CIDR:  cidr,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/vpcsubnets"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SubnetResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("서브넷 생성 실패: %w", err)
	}

	return &result.VPCSubnet, nil
}

func (c *Client) DeleteSubnet(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/vpcsubnets/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("서브넷 삭제 실패: %w", err)
	}

	return nil
}
