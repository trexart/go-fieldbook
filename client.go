package fieldbook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

//Client is a Fieldbook API client
type Client struct {
	baseURL string
	key     string
	secret  string
	bookID  string
	hc      *http.Client
}

//NewClient creates and returns a new Fieldbook API client
func NewClient(key, secret, bookID string) *Client {
	return &Client{
		baseURL: BASE_FIELDBOOK_URL,
		key:     key,
		secret:  secret,
		bookID:  bookID,
		hc:      &http.Client{},
	}
}

//SetHTTPClient allows to set a different http client from the default
func (c *Client) SetHTTPClient(hc *http.Client) {
	c.hc = hc
}

func (c *Client) SetBaseURL(URL string) {
	c.baseURL = URL
}

func (c *Client) httpClient() *http.Client {
	return c.hc
}

func (c *Client) getURL() string {
	return fmt.Sprintf("%s/%s", c.baseURL, c.bookID)
}

func (c *Client) doReq(req *http.Request) ([]byte, http.Header, error) {
	client := c.httpClient()
	req.SetBasicAuth(c.key, c.secret)
	//req.Header.Add("Authorization", c.getAuth())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// handle errors
		errMsg := resp.Status
		log.Println(errMsg)
		var err error
		if resp.Body != nil {
			var slurp []byte
			slurp, err = ioutil.ReadAll(resp.Body)
			log.Println(string(slurp))
			if err != nil {
				return nil, resp.Header, err
			}
			ue := new(Error)
			plainUE := new(Error)
			if jerr := json.Unmarshal(slurp, ue); jerr == nil && !reflect.DeepEqual(ue, plainUE) {
				ue.code = resp.StatusCode
				err = ue
				log.Println(ue)
			} else {
				errMsg = string(slurp)
			}
		}
		if err == nil {
			log.Println("create status code error")
			err = &StatusCodeError{
				msg:  resp.Status,
				code: resp.StatusCode,
			}
			log.Println("Err: ", err.Error())
		}
		return nil, resp.Header, err
	}

	blob, err := ioutil.ReadAll(resp.Body)
	return blob, resp.Header, err
}
