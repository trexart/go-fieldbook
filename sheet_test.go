package fieldbook

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_listSheets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s", BOOK),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[
				"products",
				"styles",
				"colours",
				"sizes",
				"product_variations"
			  ]`)
		},
	)

	sheets, err := client.listSheets()
	if err != nil {
		t.Errorf("listSheets returned error: %v", err)
	}

	want := 5
	if len(sheets) != want {
		t.Errorf("listSheets returned %+v, want %+v",
			len(sheets), want)
	}
}
