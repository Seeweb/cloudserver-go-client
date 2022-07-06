package seeweb

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestServerCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &SeewebCreateServerRequest{
		Plan:     "ECS1",
		Location: "it-fr2",
		Image:    "centos-7",
		Notes:    "my first server",
		SSHKey:   "public_key_label",
	}

	mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(SeewebCreateServerRequest)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"status":"ok","action_id":38,"server":{"name":"ec200016","ipv4":"","ipv6":"","plan":"ECS1","plan_size":{"core":"1","ram":"1024","disk":"20"},"location":"it-fr2","notes":"my first server","so":"centos-7","creation_date":"2019-04-30T15:19:48.535586+00:00","deletion_date":null,"active_flag":false,"status":"Booting","api_version":"v4","user":"admin"}}`))
	})

	resp, _, err := client.Server.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebCreateServerResponse{
		Status:   "ok",
		ActionID: 38,
		Server: &Server{
			Name:         "ec200016",
			Ipv4:         "",
			Ipv6:         "",
			Plan:         "ECS1",
			Location:     "it-fr2",
			Notes:        "my first server",
			So:           "centos-7",
			CreationDate: time.Date(2019, time.April, 30, 15, 19, 48, 535586000, time.FixedZone("", 0)),
			DeletionDate: nil,
			ActiveFlag:   false,
			Status:       "Booting",
			APIVersion:   "v4",
			User:         "admin",
			PlanSize: &PlanSize{
				Core: "1",
				RAM:  "1024",
				Disk: "20",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
