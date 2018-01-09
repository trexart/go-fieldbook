package fieldbook

import (
	"fmt"
	"net/http"
)

var BASE_FIELDBOOK_URL = "https://api.fieldbook.com/v1"

type Error struct {
	code    int
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%v", err.Message)
}

type StatusCodeError struct {
	msg  string
	code int
}

func (cerr *StatusCodeError) Error() string {
	return fmt.Sprintf("%v", cerr.msg)
}

func (cerr *StatusCodeError) Code() int {
	if cerr == nil {
		return http.StatusOK
	}
	return cerr.code
}
