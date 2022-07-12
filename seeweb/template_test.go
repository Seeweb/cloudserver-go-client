package seeweb

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTemplateGet(t *testing.T) {
	setup()
	defer teardown()

	templateId := "9"
	input := templateId

	mux.HandleFunc(fmt.Sprintf("/templates/%s", templateId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","template":{"id":9,"name":"ei100006","creation_date":"2019-04-24T11:12:44.121251+00:00","active_flag":true,"status":"Created","uuid":"6e71411b-5c94-202c-c4fc-bacd3c864b1b","notes":"Template server staging"}}`))
	})

	resp, _, err := client.Template.Get(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebTemplateGetResponse{
		Status: "ok",
		Template: &Template{
			ID:           9,
			Name:         "ei100006",
			CreationDate: time.Date(2019, time.April, 24, 11, 12, 44, 121251000, time.FixedZone("", 0)),
			ActiveFlag:   true,
			Status:       "Created",
			UUID:         "6e71411b-5c94-202c-c4fc-bacd3c864b1b",
			Notes:        "Template server staging",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTemplateList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/templates/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"status":"ok","templates":[{"id":9,"name":"ei100006","creation_date":"2019-04-24T11:12:44.121251+00:00","active_flag":true,"status":"Created","uuid":"6e71411b-5c94-202c-c4fc-bacd3c864b1b","notes":"Template server staging"},{"id":10,"name":"ei100007","creation_date":"2019-04-24T11:28:58.821111+00:00","active_flag":true,"status":"Created","uuid":"b3f8115b-1090-0173-b37e-b27685bd9dfc","notes":"Template server produzione"}]}`))
	})

	resp, _, err := client.Template.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &SeewebTemplateListResponse{
		Status: "ok",
		Templates: []*Template{
			{
				ID:           9,
				Name:         "ei100006",
				CreationDate: time.Date(2019, time.April, 24, 11, 12, 44, 121251000, time.FixedZone("", 0)),
				ActiveFlag:   true,
				Status:       "Created",
				UUID:         "6e71411b-5c94-202c-c4fc-bacd3c864b1b",
				Notes:        "Template server staging",
			}, {
				ID:           10,
				Name:         "ei100007",
				CreationDate: time.Date(2019, time.April, 24, 11, 28, 58, 821111000, time.FixedZone("", 0)),
				ActiveFlag:   true,
				Status:       "Created",
				UUID:         "b3f8115b-1090-0173-b37e-b27685bd9dfc",
				Notes:        "Template server produzione",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
