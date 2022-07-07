package seeweb

import (
	"fmt"
	"time"
)

// ServerService handles the communication with server
// related methods of the Seeweb API.
type ServerService service

type PlanSize struct {
	Core string `json:"core"`
	RAM  string `json:"ram"`
	Disk string `json:"disk"`
}

// Server represents a server.
type Server struct {
	Name         string    `json:"name"`
	Ipv4         string    `json:"ipv4"`
	Ipv6         string    `json:"ipv6"`
	Plan         string    `json:"plan"`
	PlanSize     *PlanSize `json:"plan_size"`
	Location     string    `json:"location"`
	Notes        string    `json:"notes"`
	So           string    `json:"so"`
	CreationDate time.Time `json:"creation_date"`
	DeletionDate time.Time `json:"deletion_date"`
	ActiveFlag   bool      `json:"active_flag"`
	Status       string    `json:"status"`
	APIVersion   string    `json:"api_version"`
	User         string    `json:"user"`
}

type Action struct {
	ID           int       `json:"id"`
	Status       string    `json:"status"`
	User         string    `json:"user"`
	CreatedAt    time.Time `json:"created_at"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	Resource     string    `json:"resource"`
	ResourceType string    `json:"resource_type"`
	Type         string    `json:"type"`
	Progress     int       `json:"progress"`
}

type SeewebServerCreateRequest struct {
	Plan     string `json:"plan"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Notes    string `json:"notes"`
	SSHKey   string `json:"ssh_key"`
}
type SeewebServerCreateResponse struct {
	Status   string  `json:"status"`
	ActionID int     `json:"action_id"`
	Server   *Server `json:"server"`
}

type SeewebServerListResponse struct {
	Status string    `json:"status"`
	Count  int       `json:"count"`
	Server []*Server `json:"server"`
}

type SeewebServerDeleteResponse struct {
	Status string  `json:"status"`
	Action *Action `json:"action"`
}

type SeewebServerUpdateRequest struct {
	Note  string `json:"note"`
	Group string `json:"group"`
}

type SeewebServerUpdateResponse struct {
	Status string `json:"status"`
}

// Create creates a new server.
func (s *ServerService) Create(createServerRequest *SeewebServerCreateRequest) (*SeewebServerCreateResponse, *Response, error) {
	u := "/servers"
	v := new(SeewebServerCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, &createServerRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// List lists all existing servers.
func (s *ServerService) List() (*SeewebServerListResponse, *Response, error) {
	u := "/servers"
	v := new(SeewebServerListResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing server.
func (s *ServerService) Delete(name string) (*SeewebServerDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s", name)
	v := new(SeewebServerDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing server.
func (s *ServerService) Update(name string, updateServerRequest *SeewebServerUpdateRequest) (*SeewebServerUpdateResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s", name)
	v := new(SeewebServerUpdateResponse)

	resp, err := s.client.newRequestDo("PUT", u, &updateServerRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
