package appresponse

import (
	"enigmacamp.com/gosql/logger"
	"net/http"
)

type ResponseMessage struct {
	Description string      `json:"message"`
	Data        interface{} `json:"data"`
}
type ErrorMessage struct {
	Code             int    `json:"code"`
	ErrorDescription string `json:"message"`
}

func logError(errMessage *ErrorMessage) {
	logger.Logger.Error().Msgf("code: %d, message: %s", errMessage.Code, errMessage.ErrorDescription)
}

func NewResponseMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		description, data,
	}
}
func NewSimpleResponseMessage(data interface{}) ResponseMessage {
	return ResponseMessage{
		"0", data,
	}
}
func NewUnauthorizedError(err error) *ErrorMessage {
	em := &ErrorMessage{
		Code:             http.StatusUnauthorized,
		ErrorDescription: err.Error(),
	}
	logError(em)
	return em
}
func NewInternalServerError(err error) *ErrorMessage {
	em := &ErrorMessage{
		Code:             http.StatusInternalServerError,
		ErrorDescription: err.Error(),
	}
	logError(em)
	return em
}
func NewBadRequestError(err error) *ErrorMessage {
	em := &ErrorMessage{
		Code:             http.StatusBadRequest,
		ErrorDescription: err.Error(),
	}
	logError(em)
	return em
}
