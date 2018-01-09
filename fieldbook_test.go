package fieldbook

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const BOOK = "book1"

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	//client = NewClient(nil)
	//client.BaseURL, _ = url.Parse(server.URL)
	//client = NewClient("key-1", "pHASvYybHFSnnqWBQTWt", "5a540048e4121e0300be523e")
	client = NewClient("", "", BOOK)
	client.SetBaseURL(server.URL)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

func TestError_Error(t *testing.T) {
}

func TestStatusCodeError_Error(t *testing.T) {
}

func TestStatusCodeError_Code(t *testing.T) {
}
