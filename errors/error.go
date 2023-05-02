package errors

import "fmt"

type ErrorDetails struct {
	Level    int    `json:"level"`
	Code     int    `json:"code"`
	ErrorMsg string `json:"error"`
}

func (e ErrorDetails) Error() string {
	return e.ErrorMsg
}

func (e ErrorDetails) Formats(values ...interface{}) ErrorDetails {
	return ErrorDetails{
		Level:    e.Level,
		Code:     e.Code,
		ErrorMsg: fmt.Sprintf(e.ErrorMsg, values...),
	}
}

func createErrorDetails(code int, msg string, level int) ErrorDetails {
	return ErrorDetails{
		Level:    level,
		Code:     code,
		ErrorMsg: msg,
	}
}
