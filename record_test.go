package fieldbook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type TestRecord struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Updated bool   `json:"is_updated"`
}

func TestClient_ListRecords(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/products", BOOK),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[
				{
				  "id": 1,
				  "name": "Test 1",
				  "is_updated": false
				},
				{
					"id": 2,
					"name": "Test 2",
					"is_updated": false
				},
				{
					"id": 3,
					"name": "Test 3",
					"is_updated": true
				}
			  ]`)
		},
	)

	var records []TestRecord
	err := client.ListRecords("products", &records, nil)
	if err != nil {
		t.Errorf("listRecords returned error: %v", err)
	}

	want := 3
	if len(records) != want {
		t.Errorf("listRecords returned %+v, want %+v",
			len(records), want)
	}

	log.Printf("%v", records)
}

func TestClient_GetRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/products/%v", BOOK, 1),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{
				  "id": 1,
				  "name": "Test 1",
				  "is_updated": false
				}`)
		},
	)

	var record TestRecord
	err := client.GetRecord("products", 1, &record, nil)
	if err != nil {
		t.Errorf("getRecord returned error: %v", err)
	}

	if record.ID != 1 {
		t.Errorf("getRecords returned %+v, want %+v",
			record.ID, 1)
	}
	if record.Name != "Test 1" {
		t.Errorf("getRecords returned %+v, want %+v",
			record.Name, "Test 1")
	}
	if record.Updated != false {
		t.Errorf("getRecords returned %+v, want %+v",
			record.Updated, false)
	}

	log.Printf("%v", record)
}

func TestClient_CreateRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/products", BOOK),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			var record TestRecord
			json.NewDecoder(r.Body).Decode(&record)
			fmt.Fprintf(w, `{
				  "id": 4,
				  "name": "%s",
				  "is_updated": %v
				}`, record.Name, record.Updated)
		},
	)

	record := TestRecord{
		Name: "Test Create",
	}

	err := client.CreateRecord("products", &record)
	if err != nil {
		t.Errorf("createRecord returned error: %v", err)
	}

	if record.ID == 0 {
		t.Errorf("createRecord didn't properly return the newly set ID")
	}
	if record.ID != 4 {
		t.Errorf("createRecord isn't returned the correct ID")
	}
	if record.Name != "Test Create" {
		t.Errorf("createRecord isn't returned the correct name")
	}
}

func TestClient_UpdateRecord(t *testing.T) {
	setup()
	defer teardown()

	record := TestRecord{
		ID:   1,
		Name: "Test Create 1",
	}

	mux.HandleFunc(fmt.Sprintf("/%s/products/%v", BOOK, record.ID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PATCH")
			var record TestRecord
			json.NewDecoder(r.Body).Decode(&record)
			fmt.Fprintf(w, `{
				  "id": %v,
				  "name": "%s",
				  "is_updated": %v
				}`, record.ID, record.Name, record.Updated)
		},
	)

	err := client.UpdateRecord("products", record.ID, &record)
	if err != nil {
		t.Errorf("updateRecord returned error: %v", err)
	}
	if record.ID != 1 {
		t.Errorf("updateRecord isn't returned the correct ID")
	}
	if record.Name != "Test Create 1" {
		t.Errorf("updateRecord isn't returned the correct name")
	}
}

func TestClient_DeleteRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/products/%v", BOOK, 1),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		},
	)

	err := client.DeleteRecord("products", 1)
	if err != nil {
		t.Errorf("deleteRecord returned error: %v", err)
	}
}
