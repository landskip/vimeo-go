package vimeo

import (
	"net/url"
	"strconv"
)

// Group client
type GroupClient struct{}

// Create group
// GroupOptions is not required
func (c *GroupClient) CreateGroup(name, description string) (*Group, error) {
	var values url.Values
	values.Add("name", name)
	values.Add("description", description)

	group := Group{}
	err := query("POST", "/groups", values, &group)
	return &group, err
}

func (c *GroupClient) DeleteGroup(groupID int) error {
	return query("DELETE", "/groups/"+strconv.Itoa(groupID), nil, nil)
}

// Show own group
func (c *GroupClient) ShowMyGroup(groupID int) (*Group, error) {
	group := Group{}
	err := query("GET", "/me/groups/"+strconv.Itoa(groupID), nil, &group)
	return &group, err
}

// list own groups
// ListOptions is not required
func (c *GroupClient) ListMyGroups(option *ListOptions) (*GroupList, error) {
	var values url.Values
	if option != nil {
		values = option.createValues()
	}
	groups := GroupList{}
	err := query("GET", "/me/groups", values, &groups)
	return &groups, err
}

func (c *GroupClient) ListAllGroups(option *ListOptions) (*GroupList, error) {
	var values url.Values
	if option != nil {
		values = option.createValues()
	}
	groups := GroupList{}
	err := query("GET", "/groups", values, &groups)
	return &groups, err
}

func (c *GroupClient) JoinGroup(groupID int) error {
	return query("PUT", "/me/groups/"+strconv.Itoa(groupID), nil, nil)
}

func (c *GroupClient) LeaveGroup(groupID int) error {
	return query("DELETE", "/me/groups/"+strconv.Itoa(groupID), nil, nil)
}
