package seeweb

import (
	"net/http"
	"reflect"
	"testing"
)

// List lists all existing groups
func TestGroupList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","groups":[{"id":2,"name":"eg100001","notes":"Test group 1","enabled":true}]}`))
	})

	resp, _, err := client.Group.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebGroupListResponse{
		Status: "ok",
		Groups: []*Group{
			{
				ID:      2,
				Name:    "eg100001",
				Notes:   "Test group 1",
				Enabled: true,
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
