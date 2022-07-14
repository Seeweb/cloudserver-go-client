package seeweb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPlanList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/plans/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","plans":[{"id":1,"name":"ECS1","cpu":"1","ram":"1024","disk":"20","hourly_price":0.017,"montly_price":12,"windows":false,"available":true,"available_regions":[{"id":1,"location":"it-fr2","description":"Frosinone"}]}]}`))
	})

	resp, _, err := client.Plan.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebPlanListResponse{
		Status: "ok",
		Plans: []*Plan{
			{
				ID:          1,
				Name:        "ECS1",
				CPU:         "1",
				RAM:         "1024",
				Disk:        "20",
				HourlyPrice: 0.017,
				MontlyPrice: 12,
				Windows:     false,
				Available:   true,
				AvailableRegions: []*AvailableRegions{
					{
						ID:          1,
						Location:    "it-fr2",
						Description: "Frosinone",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
