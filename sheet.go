package fieldbook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//ListSheets returns a string slice of all sheets in book
func (c *Client) ListSheets() ([]string, error) {
	fullURL := fmt.Sprintf("%s", c.getURL())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	slurp, _, err := c.doReq(req)
	if err != nil {
		return nil, err
	}

	//log.Println(string(slurp))

	var sheets []string
	if err := json.Unmarshal(slurp, &sheets); err != nil {
		return nil, err
	}

	return sheets, nil
}
