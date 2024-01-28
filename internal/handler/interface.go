package handler

type ResponseInterface interface {
	SetOk(data interface{}) HTTPResponse
	SetOkWithStatus(status int, data interface{}) HTTPResponse
	SetError(err error, errCode int, message string) HTTPResponse
	SetErrorWithStatus(status int, err error, errCode int, message string) HTTPResponse
	HasError() bool
	GetData() interface{}
	GetError() error
	GetStatus() int
	GetErrCode() int
	GetErrorMessage() string
	GetErrorMessageVerbose() string
	HasNoContent() bool
}
