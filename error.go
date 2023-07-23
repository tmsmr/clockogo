package clockogo

import (
	"encoding/json"
	"fmt"
)

type ErrorDetail struct {
	Message string      `json:"message"`
	Fields  interface{} `json:"fields"`
}

type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

type APIError struct {
	code int
	msg  ErrorBody
}

func NewAPIError(code int, body []byte) error {
	var errBody ErrorBody
	err := json.Unmarshal(body, &errBody)
	if err != nil {
		return err
	}
	return APIError{
		code: code,
		msg:  errBody,
	}
}

func (err APIError) Error() string {
	return fmt.Sprintf("API request failed with status %d: %s", err.code, err.msg.Error.Message)
}
