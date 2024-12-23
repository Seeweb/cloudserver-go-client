package seeweb

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestActionGet(t *testing.T) {
	setup()
	defer teardown()

	actionId := "39"
	input := actionId

	mux.HandleFunc(fmt.Sprintf("/actions/%s", actionId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","action":{"id":39,"status":"completed","user":"admin","created_at":"2019-04-30T16:33:03.317800+00:00","started_at":"2019-04-30T16:33:03.317019+00:00","completed_at":"2019-04-30T16:34:04.159922+00:00","resource":"ec200016","resource_type":"ECS","type":"delete_server","progress":100}}`))
	})

	resp, _, err := client.Action.Get(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebActionGetResponse{
		Status: "ok",
		Action: &Action{
			ID:           39,
			Status:       "completed",
			User:         "admin",
			CreatedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317800000, time.FixedZone("", 0)),
			StartedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317019000, time.FixedZone("", 0)),
			CompletedAt:  time.Date(2019, time.April, 30, 16, 34, 4, 159922000, time.FixedZone("", 0)),
			Resource:     "ec200016",
			ResourceType: "ECS",
			Type:         "delete_server",
			Progress:     100,
		},
	}

	if resp.Status != want.Status {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

	if !equalStructWithDatesFn(*resp.Action, *want.Action) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestActionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/actions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","actions":[{"id":39,"status":"completed","user":"admin","created_at":"2019-04-30T16:33:03.317800+00:00","started_at":"2019-04-30T16:33:03.317019+00:00","completed_at":"2019-04-30T16:34:04.159922+00:00","resource":"ec200016","resource_type":"ECS","type":"delete_server","progress":100},{"id":38,"status":"completed","user":"admin","created_at":"2019-04-30T15:19:48.539503+00:00","started_at":"2019-04-30T15:19:48.534100+00:00","completed_at":"2019-04-30T15:20:10.112664+00:00","resource":"ec200016","resource_type":"ECS","type":"create_server","progress":100}]}`))
	})

	resp, _, err := client.Action.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebActionListResponse{
		Status: "ok",
		Actions: []*Action{
			{
				ID:           39,
				Status:       "completed",
				User:         "admin",
				CreatedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317800000, time.FixedZone("", 0)),
				StartedAt:    time.Date(2019, time.April, 30, 16, 33, 3, 317019000, time.FixedZone("", 0)),
				CompletedAt:  time.Date(2019, time.April, 30, 16, 34, 4, 159922000, time.FixedZone("", 0)),
				Resource:     "ec200016",
				ResourceType: "ECS",
				Type:         "delete_server",
				Progress:     100,
			},
			{
				ID:           38,
				Status:       "completed",
				User:         "admin",
				CreatedAt:    time.Date(2019, time.April, 30, 15, 19, 48, 539503000, time.FixedZone("", 0)),
				StartedAt:    time.Date(2019, time.April, 30, 15, 19, 48, 534100000, time.FixedZone("", 0)),
				CompletedAt:  time.Date(2019, time.April, 30, 15, 20, 10, 112664000, time.FixedZone("", 0)),
				Resource:     "ec200016",
				ResourceType: "ECS",
				Type:         "create_server",
				Progress:     100,
			},
		},
	}

	if resp.Status != want.Status {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

	for i := 0; i < len(want.Actions); i++ {
		if !equalStructWithDatesFn(*resp.Actions[i], *want.Actions[i]) {
			t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
		}
	}
}
