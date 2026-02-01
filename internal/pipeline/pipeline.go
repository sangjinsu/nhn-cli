package pipeline

import (
	"fmt"
	"net/http"

	"nhncli/internal/client"
)

func (c *Client) ExecutePipeline(pipelineName string) error {
	url := fmt.Sprintf("%s/pipeline/v1.0/pipeline-exec/%s", c.baseURL, pipelineName)

	resp, err := c.httpClient.Post(url, nil, c.getOpts())
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		resp.Body.Close()
		return nil
	}

	var result PipelineExecuteResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return err
	}

	return checkResponse(result.Header)
}
