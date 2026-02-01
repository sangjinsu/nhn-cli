package deploy

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"nhncli/internal/client"
)

func (c *Client) ExecuteDeploy(req *DeployExecuteRequest) (*DeployResult, error) {
	url := fmt.Sprintf("%s/api/v1.0/appkeys/%s/deployments", c.baseURL, c.appKey)

	resp, err := c.httpClient.Post(url, req, c.getOpts())
	if err != nil {
		return nil, err
	}

	var result DeployExecuteResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Body, nil
}

func (c *Client) UploadBinary(req *BinaryUploadRequest) (*BinaryResult, error) {
	url := fmt.Sprintf("%s/api/v1.0/projects/%s/artifacts/%d/binary-group/%d",
		c.baseURL, c.appKey, req.ArtifactID, req.BinaryGroupKey)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// applicationType (필수)
	writer.WriteField("applicationType", req.ApplicationType)

	// binaryFile (필수)
	f, err := os.Open(req.BinaryFile)
	if err != nil {
		return nil, fmt.Errorf("파일 열기 실패: %w", err)
	}
	defer f.Close()

	part, err := writer.CreateFormFile("binaryFile", filepath.Base(req.BinaryFile))
	if err != nil {
		return nil, fmt.Errorf("multipart 파일 생성 실패: %w", err)
	}
	if _, err := io.Copy(part, f); err != nil {
		return nil, fmt.Errorf("파일 복사 실패: %w", err)
	}

	// 선택 필드
	if req.Version != "" {
		writer.WriteField("version", req.Version)
	}
	if req.Description != "" {
		writer.WriteField("description", req.Description)
	}
	if req.OsType != "" {
		writer.WriteField("osType", req.OsType)
	}
	if req.Fix {
		writer.WriteField("fix", strconv.FormatBool(req.Fix))
	}

	// metaFile (선택)
	if req.MetaFile != "" {
		mf, err := os.Open(req.MetaFile)
		if err != nil {
			return nil, fmt.Errorf("메타 파일 열기 실패: %w", err)
		}
		defer mf.Close()

		metaPart, err := writer.CreateFormFile("metaFile", filepath.Base(req.MetaFile))
		if err != nil {
			return nil, fmt.Errorf("multipart 메타 파일 생성 실패: %w", err)
		}
		if _, err := io.Copy(metaPart, mf); err != nil {
			return nil, fmt.Errorf("메타 파일 복사 실패: %w", err)
		}
	}

	writer.Close()

	httpReq, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	for k, v := range c.headers {
		httpReq.Header.Set(k, v)
	}

	if c.debug {
		fmt.Printf("[DEBUG] POST %s\n", url)
	}

	resp, err := c.httpClient.GetRawClient().Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %w", err)
	}

	var result BinaryUploadResponse
	if err := client.ReadJSON(resp, &result); err != nil {
		return nil, err
	}

	if err := checkResponse(result.Header); err != nil {
		return nil, err
	}

	return result.Body, nil
}
