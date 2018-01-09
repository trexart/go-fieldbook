package fieldbook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//ListRecords returns all records for selected sheet and query options
func (c *Client) ListRecords(sheet string, records interface{}, options *QueryOptions) error {
	fullURL := fmt.Sprintf("%s/%s", c.getURL(), sheet)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	if options != nil {
		options.implement(req)
	}

	slurp, _, err := c.doReq(req)
	if err != nil {
		return err
	}

	//log.Println(string(slurp))

	if err := json.Unmarshal(slurp, &records); err != nil {
		return err
	}

	//log.Printf("%v", records)

	return nil
}

//GetRecord returns the record in the selected sheet with the id
func (c *Client) GetRecord(sheet string, id int, record interface{}, options *QueryOptions) error {
	fullURL := fmt.Sprintf("%s/%s/%v", c.getURL(), sheet, id)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	if options != nil {
		options.implement(req)
	}

	slurp, _, err := c.doReq(req)
	if err != nil {
		return err
	}

	//log.Println(string(slurp))

	if err := json.Unmarshal(slurp, &record); err != nil {
		return err
	}

	//log.Printf("%v", records)

	return nil
}

//CreateRecord adds a new record to the selected sheet
func (c *Client) CreateRecord(sheet string, record interface{}) error {
	fullURL := fmt.Sprintf("%s/%s", c.getURL(), sheet)

	jsonStr, err := json.Marshal(record)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	slurp, _, err := c.doReq(req)
	if err != nil {
		return err
	}

	//log.Println(string(slurp))

	if err := json.Unmarshal(slurp, &record); err != nil {
		return err
	}

	return nil
}

//UpdateRecord uses PATCH to update only passed in fields for record identified by id in the selected sheet
func (c *Client) UpdateRecord(sheet string, id int, record interface{}) error {
	fullURL := fmt.Sprintf("%s/%s/%v", c.getURL(), sheet, id)

	jsonStr, err := json.Marshal(record)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", fullURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	slurp, _, err := c.doReq(req)
	if err != nil {
		return err
	}

	//log.Println(string(slurp))

	if err := json.Unmarshal(slurp, &record); err != nil {
		return err
	}

	return nil
}

//DeleteRecord removes the record with id in selected sheet
func (c *Client) DeleteRecord(sheet string, id int) error {
	fullURL := fmt.Sprintf("%s/%s/%v", c.getURL(), sheet, id)

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	_, _, err = c.doReq(req)
	if err != nil {
		return err
	}

	//log.Println(string(slurp))

	return nil
}
