package seeweb

// GroupService handles the communication with group
// related methods of the Seeweb API.
type GroupService service

type Group struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Notes   string `json:"notes"`
	Enabled bool   `json:"enabled"`
}
type SeewebGroupCreateRequest struct {
	Notes    string `json:"notes"`
	Password string `json:"password"`
}
type SeewebGroupCreateResponse struct {
	Status string `json:"status"`
	Group  *Group `json:"group"`
}

type SeewebGroupListResponse struct {
	Status string   `json:"status"`
	Groups []*Group `json:"groups"`
}

type SeewebGroupDeleteResponse struct {
	Status string `json:"status"`
}

// List lists all existing groups.
func (a *GroupService) List() (*SeewebGroupListResponse, *Response, error) {
	u := "/groups"
	v := new(SeewebGroupListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
