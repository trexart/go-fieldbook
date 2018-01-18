package fieldbook

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var BASE_FIELDBOOK_URL = "https://api.fieldbook.com/v1"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	t.Time, err = time.Parse("2006-01-02", s)
	return nil
}

type QueryOptions struct {
	Exclude []string
	Include []string
	Expand  []string
}

func (options *QueryOptions) implement(req *http.Request) {
	q := req.URL.Query()
	if len(options.Exclude) > 0 {
		q.Add("exclude", strings.Join(options.Exclude, ","))
	}
	if len(options.Include) > 0 {
		q.Add("include", strings.Join(options.Include, ","))
	}
	if len(options.Expand) > 0 {
		q.Add("expand", strings.Join(options.Expand, ","))
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
