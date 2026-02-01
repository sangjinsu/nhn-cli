package loadbalancer

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListLoadBalancers() ([]LoadBalancer, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/lbaas/loadbalancers"), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result LoadBalancerListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("로드 밸런서 목록 조회 실패: %w", err)
	}

	return result.Loadbalancers, nil
}

func (c *Client) GetLoadBalancer(id string) (*LoadBalancer, error) {
	resp, err := c.httpClient.Get(c.url("/v2.0/lbaas/loadbalancers/"+id), c.getOpts())
	if err != nil {
		return nil, err
	}

	var result LoadBalancerResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("로드 밸런서 조회 실패: %w", err)
	}

	return &result.Loadbalancer, nil
}

func (c *Client) CreateLoadBalancer(req *LoadBalancerCreateRequest) (*LoadBalancer, error) {
	resp, err := c.httpClient.Post(c.url("/v2.0/lbaas/loadbalancers"), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result LoadBalancerResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("로드 밸런서 생성 실패: %w", err)
	}

	return &result.Loadbalancer, nil
}

func (c *Client) UpdateLoadBalancer(id string, req *LoadBalancerUpdateRequest) (*LoadBalancer, error) {
	resp, err := c.httpClient.Put(c.url("/v2.0/lbaas/loadbalancers/"+id), req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result LoadBalancerResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("로드 밸런서 수정 실패: %w", err)
	}

	return &result.Loadbalancer, nil
}

func (c *Client) DeleteLoadBalancer(id string) error {
	resp, err := c.httpClient.Delete(c.url("/v2.0/lbaas/loadbalancers/"+id), c.getOpts())
	if err != nil {
		return err
	}

	if err := client.ReadJSON(resp, nil); err != nil {
		return fmt.Errorf("로드 밸런서 삭제 실패: %w", err)
	}

	return nil
}
