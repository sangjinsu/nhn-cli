package cdn

import (
	"fmt"

	"nhncli/internal/client"
)

func (c *Client) ListServices() ([]CDNService, error) {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/services", c.baseURL, c.appKey)

	resp, err := c.httpClient.Get(url, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ServiceListResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Contents, nil
}

func (c *Client) CreateService(req *ServiceCreateRequest) (*CDNService, error) {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/services", c.baseURL, c.appKey)

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ServiceResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Content, nil
}

func (c *Client) UpdateService(domain string, req *ServiceCreateRequest) (*CDNService, error) {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/services/%s", c.baseURL, c.appKey, domain)

	resp, err := c.httpClient.Put(url, req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result ServiceResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return &result.Content, nil
}

func (c *Client) DeleteService(domain string) error {
	url := fmt.Sprintf("%s/cdn/v2.0/appkeys/%s/services/%s", c.baseURL, c.appKey, domain)

	resp, err := c.httpClient.Delete(url, c.getOpts())
	if err != nil {
		return err
	}

	var result PurgeResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}
