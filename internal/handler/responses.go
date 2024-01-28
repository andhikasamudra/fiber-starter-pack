package handler

import (
	"fmt"
	internalError "github.com/andhikasamudra/fiber-starter-pack/internal/error"
	"net/http"
)

type HTTPResponse struct {
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
	Err       error       `json:"-"`
	ErrCode   int         `json:"-"`
	Message   string      `json:"message"`
	NoContent bool        `json:"-"`
}

func NewResponse() *HTTPResponse {
	return &HTTPResponse{}
}

// SetOk ...
func (res HTTPResponse) SetOk(data interface{}) HTTPResponse {
	res.Status = http.StatusOK
	res.Data = data
	res.Message = "success"

	return res
}

// SetOkWithStatus ...
func (res HTTPResponse) SetOkWithStatus(status int, data interface{}) HTTPResponse {
	res.Status = status
	res.Data = data
	res.Message = "success"
	res.NoContent = status == http.StatusNoContent

	return res
}

// SetError ...
func (res HTTPResponse) SetError(err error, errCode int, message string) HTTPResponse {
	res.Status = http.StatusInternalServerError
	res.Err = err
	res.ErrCode = errCode
	res.Message = message
	res.Data = nil

	return res
}

// SetErrorWithStatus ...
func (res HTTPResponse) SetErrorWithStatus(status int, err error, errCode int, message string) HTTPResponse {
	res.Status = status
	res.Err = err
	res.ErrCode = errCode
	res.Message = message

	return res
}

// HasError ...
func (res HTTPResponse) HasError() bool {
	return res.Err != nil
}

// GetData ...
func (res HTTPResponse) GetData() interface{} {
	return res.Data
}

// GetError ...
func (res HTTPResponse) GetError() error {
	return res.Err
}

// GetStatus ...
func (res HTTPResponse) GetStatus() int {
	if res.Status != 0 {
		return res.Status
	}
	return http.StatusInternalServerError
}

// GetErrCode ...
func (res HTTPResponse) GetErrCode() int {
	if res.ErrCode != 0 {
		return res.ErrCode
	}

	return internalError.UnknownError
}

// GetErrorMessage get error message from message or error object
func (res HTTPResponse) GetErrorMessage() string {
	if res.Message != "" {
		return res.Message
	}

	return res.Err.Error()
}

// GetErrorMessageVerbose get full string with error code, message and error object
func (res HTTPResponse) GetErrorMessageVerbose() string {
	return fmt.Sprintf("Error Code: %d, Message: %s. Detail: %s", res.ErrCode, res.Message, res.Err.Error())
}

// HasNoContent ...
func (res HTTPResponse) HasNoContent() bool {
	return res.NoContent
}
