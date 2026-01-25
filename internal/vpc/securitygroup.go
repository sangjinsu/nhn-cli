package vpc

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListSecurityGroups() ([]SecurityGroup, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/security-groups"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SecurityGroupListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("보안 그룹 목록 조회 실패: %w", err)
	}

	return result.SecurityGroups, nil
}

func (c *Client) GetSecurityGroup(id string) (*SecurityGroup, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/security-groups/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SecurityGroupResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("보안 그룹 조회 실패: %w", err)
	}

	return &result.SecurityGroup, nil
}

func (c *Client) CreateSecurityGroup(name, description string) (*SecurityGroup, error) {
	reqBody := SecurityGroupCreateRequest{
		SecurityGroup: SecurityGroupCreateBody{
			Name:        name,
			Description: description,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/security-groups"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result SecurityGroupResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("보안 그룹 생성 실패: %w", err)
	}

	return &result.SecurityGroup, nil
}

func (c *Client) DeleteSecurityGroup(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/security-groups/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("보안 그룹 삭제 실패: %w", err)
	}

	return nil
}

func (c *Client) AddSecurityGroupRule(sgID, direction, protocol, etherType string, portMin, portMax *int, remoteIP, remoteGroupID, description string) (*SecurityGroupRule, error) {
	reqBody := SecurityGroupRuleCreateRequest{
		SecurityGroupRule: SecurityGroupRuleCreateBody{
			SecurityGroupID: sgID,
			Direction:       direction,
			EtherType:       etherType,
			Protocol:        protocol,
			PortRangeMin:    portMin,
			PortRangeMax:    portMax,
			RemoteIPPrefix:  remoteIP,
			RemoteGroupID:   remoteGroupID,
			Description:     description,
		},
	}

	resp, err := c.httpClient.Post(c.url("/v2.0/security-group-rules"), reqBody, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result struct {
		SecurityGroupRule SecurityGroupRule `json:"security_group_rule"`
	}
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("보안 그룹 규칙 추가 실패: %w", err)
	}

	return &result.SecurityGroupRule, nil
}

func (c *Client) DeleteSecurityGroupRule(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/security-group-rules/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("보안 그룹 규칙 삭제 실패: %w", err)
	}

	return nil
}
