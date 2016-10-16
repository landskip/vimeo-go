package vimeo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Group client
type GroupClient struct{}

// Option values for createGroup method.
type CreateGroupOption struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (o CreateGroupOption) GetValues() url.Values {
	return nil
}
func (o CreateGroupOption) GetJSONBytes() ([]byte, error) {
	return json.Marshal(o)
}

// Create group
// GroupOptions is not required
func (c *GroupClient) CreateGroup(option *CreateGroupOption) (*Group, error) {
	var data QueryRequest
	if option != nil {
		data = option
	}
	group := Group{}
	err := query(http.MethodPost, "/groups", data, &group)
	return &group, err
}

// Delete group
func (c *GroupClient) DeleteGroup(groupID int) error {
	return query(http.MethodDelete, "/groups/"+strconv.Itoa(groupID), nil, nil)
}

// Show own group
func (c *GroupClient) ShowMyGroup(groupID int) (*Group, error) {
	group := Group{}
	err := query(http.MethodGet, "/me/groups/"+strconv.Itoa(groupID), nil, &group)
	return &group, err
}

// list own groups
// ListOptions is not required
func (c *GroupClient) ListMyGroups(option *ListOptions) (*GroupList, error) {
	var data QueryRequest
	if option != nil {
		data = option
	}
	groups := GroupList{}
	err := query(http.MethodGet, "/me/groups", data, &groups)
	return &groups, err
}

func (c *GroupClient) ListAllGroups(option *ListOptions) (*GroupList, error) {
	groups := GroupList{}
	err := query(http.MethodGet, "/groups", option, &groups)
	return &groups, err
}

func (c *GroupClient) JoinGroup(groupID int) error {
	return query(http.MethodPut, "/me/groups/"+strconv.Itoa(groupID), nil, nil)
}

func (c *GroupClient) LeaveGroup(groupID int) error {
	return query(http.MethodDelete, "/me/groups/"+strconv.Itoa(groupID), nil, nil)
}
