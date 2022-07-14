package seeweb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRegionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/regions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","regions":[{"id":1,"location":"it-fr2","description":"Frosinone"}]}`))
	})

	resp, _, err := client.Region.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebRegionListResponse{
		Status: "ok",
		Regions: []*Region{
			{
				ID:          1,
				Location:    "it-fr2",
				Description: "Frosinone",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
