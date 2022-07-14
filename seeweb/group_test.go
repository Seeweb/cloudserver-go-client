package seeweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

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

func TestGroupCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &SeewebGroupCreateRequest{
		Notes:    "Test group 1",
		Password: "secret",
	}

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(SeewebGroupCreateRequest)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"status":"ok","group":{"id":2,"name":"eg100001","notes":"Test group 1","enabled":true}}`))
	})

	resp, _, err := client.Group.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebGroupCreateResponse{
		Status: "ok",
		Group: &Group{
			ID:      2,
			Name:    "eg100001",
			Notes:   "Test group 1",
			Enabled: true,
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestGroupDelete(t *testing.T) {
	setup()
	defer teardown()

	groupID := 2
	input := groupID

	mux.HandleFunc(fmt.Sprintf("/groups/%s", strconv.Itoa(groupID)), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(`{"status":"ok"}`))
	})

	resp, _, err := client.Group.Delete(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebGroupDeleteResponse{
		Status: "ok",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
