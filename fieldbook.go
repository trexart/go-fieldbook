package fieldbook

import (
	"fmt"
	"net/http"
	"strings"
)

var BASE_FIELDBOOK_URL = "https://api.fieldbook.com/v1"

type QueryOptions struct {
	exclude []string
	include []string
	expand  []string
}

func (options *QueryOptions) implement(req *http.Request) {
	q := req.URL.Query()
	if len(options.exclude) > 0 {
		q.Add("exclude", strings.Join(options.exclude, ","))
	}
	if len(options.include) > 0 {
		q.Add("include", strings.Join(options.include, ","))
	}
	if len(options.expand) > 0 {
		q.Add("expand", strings.Join(options.expand, ","))
	}
	req.URL.RawQuery = q.Encode()
}

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
