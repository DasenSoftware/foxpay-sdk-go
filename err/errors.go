package errors

import (
	"fmt"
)

type FoxPaySDKError struct {
	Code      int32
	Message   string
	RequestId string
}

func (e *FoxPaySDKError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[FoxPaySDKError] Code=%v, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[FoxPaySDKError] Code=%v, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewFoxPaySDKError(code int32, message, requestId string) error {
	return &FoxPaySDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *FoxPaySDKError) GetCode() int32 {
	return e.Code
}

func (e *FoxPaySDKError) GetMessage() string {
	return e.Message
}

func (e *FoxPaySDKError) GetRequestId() string {
	return e.RequestId
}
