package seeweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestServerCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &SeewebServerCreateRequest{
		Plan:     "ECS1",
		Location: "it-fr2",
		Image:    "centos-7",
		Notes:    "my first server",
		SSHKey:   "public_key_label",
	}

	mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(SeewebServerCreateRequest)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"status":"ok","action_id":38,"server":{"name":"ec200016","ipv4":"","ipv6":"","plan":"ECS1","plan_size":{"core":"1","ram":"1024","disk":"20"},"location":"it-fr2","notes":"my first server","so":"centos-7","creation_date":"2019-04-30T15:19:48.535586+00:00","deletion_date":null,"active_flag":false,"status":"Booting","api_version":"v4","user":"admin","group":null}}`))
	})

	resp, _, err := client.Server.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebServerCreateResponse{
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
			DeletionDate: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			ActiveFlag:   false,
			Status:       "Booting",
			APIVersion:   "v4",
			User:         "admin",
			Group:        nil,
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

func TestServerList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/servers/"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","count":1,"server":[{"name":"ec200016","ipv4":"192.168.1.43","ipv6":"fe80::1","plan":"ECS1","plan_size":{"core":"1","ram":"1024","disk":"20"},"location":"it-fr2","notes":"my first server","so":"centos-7","creation_date":"2019-04-18T13:48:12.025548+00:00","deletion_date":null,"active_flag":true,"status":"Booted","api_version":"v4","user":"admin","virttype":"KVM"}]}`))
	})

	resp, _, err := client.Server.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebServerListResponse{
		Status: "ok",
		Count:  1,
		Server: []*Server{
			{
				Name: "ec200016",
				Ipv4: "192.168.1.43",
				Ipv6: "fe80::1",
				Plan: "ECS1",
				PlanSize: &PlanSize{
					Core: "1",
					RAM:  "1024",
					Disk: "20",
				},
				Location:     "it-fr2",
				Notes:        "my first server",
				So:           "centos-7",
				CreationDate: time.Date(2019, time.April, 18, 13, 48, 12, 25548000, time.FixedZone("", 0)),
				DeletionDate: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				ActiveFlag:   true,
				Status:       "Booted",
				APIVersion:   "v4",
				User:         "admin",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServerDelete(t *testing.T) {
	setup()
	defer teardown()

	serverName := "ec200016"
	input := serverName

	mux.HandleFunc(fmt.Sprintf("/servers/%s", serverName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(`{"status":"ok","action":{"id":39,"status":"in-progress","user":"admin","created_at":"2019-04-30T16:33:03.317800+00:00","started_at":"2019-04-30T16:33:03.317019+00:00","completed_at":null,"resource":"ec200016","resource_type":"ECS","type":"delete_server","progress":0}}`))
	})

	resp, _, err := client.Server.Delete(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebServerDeleteResponse{
		Status: "ok",
		Action: &Action{
			ID:           39,
			Status:       "in-progress",
			User:         "admin",
			CreatedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317800000, time.FixedZone("", 0)),
			StartedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317019000, time.FixedZone("", 0)),
			CompletedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Resource:     "ec200016",
			ResourceType: "ECS",
			Type:         "delete_server",
			Progress:     0,
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServerUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverName := "ec200016"
	input := &SeewebServerUpdateRequest{
		Note:  "update server name",
		Group: "eg103464",
	}

	mux.HandleFunc(fmt.Sprintf("/servers/%s", serverName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(SeewebServerUpdateRequest)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"status":"ok"}`))
	})

	resp, _, err := client.Server.Update(serverName, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebServerUpdateResponse{
		Status: "ok",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
