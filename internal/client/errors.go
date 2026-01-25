package client

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	StatusCode int
	Code       string
	Message    string
	Details    string
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	BadRequest *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"badRequest"`
	ItemNotFound *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
	Forbidden *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"forbidden"`
	Unauthorized *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"unauthorized"`
	NeutronError *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"NeutronError"`
}

func ParseAPIError(statusCode int, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: statusCode,
		Message:    getDefaultErrorMessage(statusCode),
	}

	var errResp ErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		apiErr.Details = string(body)
		return apiErr
	}

	if errResp.Error.Message != "" {
		apiErr.Code = errResp.Error.Code
		apiErr.Message = errResp.Error.Message
	} else if errResp.BadRequest != nil {
		apiErr.Message = errResp.BadRequest.Message
	} else if errResp.ItemNotFound != nil {
		apiErr.Message = errResp.ItemNotFound.Message
	} else if errResp.Forbidden != nil {
		apiErr.Message = errResp.Forbidden.Message
	} else if errResp.Unauthorized != nil {
		apiErr.Message = errResp.Unauthorized.Message
	} else if errResp.NeutronError != nil {
		apiErr.Code = errResp.NeutronError.Type
		apiErr.Message = errResp.NeutronError.Message
	}

	return apiErr
}

func getDefaultErrorMessage(statusCode int) string {
	switch statusCode {
	case 400:
		return "잘못된 요청입니다"
	case 401:
		return "인증이 필요합니다"
	case 403:
		return "접근이 거부되었습니다"
	case 404:
		return "리소스를 찾을 수 없습니다"
	case 409:
		return "리소스 충돌이 발생했습니다"
	case 429:
		return "요청 한도를 초과했습니다"
	case 500:
		return "서버 내부 오류가 발생했습니다"
	case 502:
		return "게이트웨이 오류가 발생했습니다"
	case 503:
		return "서비스를 일시적으로 사용할 수 없습니다"
	default:
		return "알 수 없는 오류가 발생했습니다"
	}
}

func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 404
	}
	return false
}

func IsUnauthorized(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 401
	}
	return false
}

func IsForbidden(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 403
	}
	return false
}
