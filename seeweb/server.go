package seeweb

import "time"

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
	Name         string      `json:"name"`
	Ipv4         string      `json:"ipv4"`
	Ipv6         string      `json:"ipv6"`
	Plan         string      `json:"plan"`
	PlanSize     *PlanSize   `json:"plan_size"`
	Location     string      `json:"location"`
	Notes        string      `json:"notes"`
	So           string      `json:"so"`
	CreationDate time.Time   `json:"creation_date"`
	DeletionDate interface{} `json:"deletion_date"`
	ActiveFlag   bool        `json:"active_flag"`
	Status       string      `json:"status"`
	APIVersion   string      `json:"api_version"`
	User         string      `json:"user"`
}

type SeewebCreateServerRequest struct {
	Plan     string `json:"plan"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Notes    string `json:"notes"`
	SSHKey   string `json:"ssh_key"`
}
type SeewebCreateServerResponse struct {
	Status   string  `json:"status"`
	ActionID int     `json:"action_id"`
	Server   *Server `json:"server"`
}

// Create creates a new server.
func (s *ServerService) Create(createServerRequest *SeewebCreateServerRequest) (*SeewebCreateServerResponse, *Response, error) {
	u := "/servers"
	v := new(SeewebCreateServerResponse)

	resp, err := s.client.newRequestDo("POST", u, &createServerRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
