package fieldbook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type TestRecord struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func TestClient_ListRecords(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s", BOOK, SHEET_NAME),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[
				{
				  "id": 1,
				  "first_name": "John",
				  "last_name": "Doe"
				},
				{
					"id": 2,
					"first_name": "Jack",
					"last_name": "Jackson"
				},
				{
					"id": 3,
					"first_name": "Jeff",
					"last_name": "Jefferson"
				}
			  ]`)
		},
	)

	var records []TestRecord
	err := client.ListRecords(SHEET_NAME, &records, nil)
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

	mux.HandleFunc(fmt.Sprintf("/%s/%s/%v", BOOK, SHEET_NAME, 1),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{
				  "id": 1,
				  "first_name": "John",
				  "last_name": "Doe"
				}`)
		},
	)

	var record TestRecord
	err := client.GetRecord(SHEET_NAME, 1, &record, nil)
	if err != nil {
		t.Errorf("getRecord returned error: %v", err)
	}

	if record.ID != 1 {
		t.Errorf("getRecords returned %+v, want %+v",
			record.ID, 1)
	}
	if record.FirstName != "John" {
		t.Errorf("getRecords returned %+v, want %+v",
			record.FirstName, "John")
	}
	if record.LastName != "Doe" {
		t.Errorf("getRecords returned %+v, want %+v",
			record.LastName, "Doe")
	}

	log.Printf("%v", record)
}

func TestClient_CreateRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s", BOOK, SHEET_NAME),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			var record TestRecord
			json.NewDecoder(r.Body).Decode(&record)
			fmt.Fprintf(w, `{
				  "id": 4,
				  "first_name": "%s",
				  "last_name": "%s"
				}`, record.FirstName, record.LastName)
		},
	)

	record := TestRecord{
		FirstName: "Test Create",
	}

	err := client.CreateRecord(SHEET_NAME, &record)
	if err != nil {
		t.Errorf("createRecord returned error: %v", err)
	}

	if record.ID == 0 {
		t.Errorf("createRecord didn't properly return the newly set ID")
	}
	if record.ID != 4 {
		t.Errorf("createRecord isn't returned the correct ID")
	}
	if record.FirstName != "Test Create" {
		t.Errorf("createRecord isn't returned the correct first name")
	}
}

func TestClient_UpdateRecord(t *testing.T) {
	setup()
	defer teardown()

	record := TestRecord{
		ID:        1,
		FirstName: "Test Create 1",
	}

	mux.HandleFunc(fmt.Sprintf("/%s/%s/%v", BOOK, SHEET_NAME, record.ID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PATCH")
			var record TestRecord
			json.NewDecoder(r.Body).Decode(&record)
			fmt.Fprintf(w, `{
				  "id": %v,
				  "first_name": "%s",
				  "last_name": "%s"
				}`, record.ID, record.FirstName, record.LastName)
		},
	)

	err := client.UpdateRecord(SHEET_NAME, record.ID, &record)
	if err != nil {
		t.Errorf("updateRecord returned error: %v", err)
	}
	if record.ID != 1 {
		t.Errorf("updateRecord isn't returned the correct ID")
	}
	if record.FirstName != "Test Create 1" {
		t.Errorf("updateRecord isn't returned the correct name")
	}
}

func TestClient_DeleteRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s/%v", BOOK, SHEET_NAME, 1),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		},
	)

	err := client.DeleteRecord(SHEET_NAME, 1)
	if err != nil {
		t.Errorf("deleteRecord returned error: %v", err)
	}
}
