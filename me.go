package vimeo

import (
	"net/url"
	"strconv"
)

// Own client
// https://developer.vimeo.com/api/endpoints/me#/ondemand/pages
type MeClient struct{}

// Me
// TODO: add Metadata
type User struct {
	URI         string      `json:"uri"`
	Name        string      `json:"name"`
	Link        string      `json:"link"`
	Location    string      `json:"location"`
	Bio         string      `json:"bio"`
	CreatedTime string      `json:"created_time"`
	Account     string      `json:"account"`
	Pictures    string      `json:"pictures"`
	Websites    []string    `json:"websites"`
	Preferences Preferences `json:"preferences"`
}

// Preferences
type Preferences struct {
	Videos Privacy `json:"videos"`
}

// Privacy
type Privacy struct {
	Privacy string `json:"privacy"`
}

// View one user
func (c *MeClient) ShowInfomation() (*User, error) {
	user := User{}
	err := query("GET", "/me", nil, &user)
	return &user, err
}

// Filter to apply to the results
const (
	FilterEmbeddable = "embeddable"
	FilterPlayable   = "PLAYABLE"
	FilterAppOnly    = "app_only"
)

// Technique used to sort the results
const (
	SortDate         = "date"
	SortAlphabetical = "alphabetical"
	SortPlays        = "plays"
	SortLikes        = "likes"
	SortComments     = "comments"
	SortDuration     = "duration"
	SortDefault      = "default"
	SortModifiedTime = "modifed_time"
)

// The direction that the results are sorted
const (
	DirectionAscending  = "asc"
	DirectionDescending = "desc"
)

// List Options
type ListOptions struct {
	Page             int
	PerPage          int
	Query            string
	Filter           string
	FilterEmbeddable bool
	FilterPlayable   bool
	Sort             string
	Direction        string
}

// Create url values from list options
func (o *ListOptions) CreateValues() url.Values {
	values := url.Values{}
	values.Add("page", strconv.Itoa(o.Page))
	values.Add("per_page", strconv.Itoa(o.PerPage))
	values.Add("query", o.Query)
	values.Add("filter", o.Filter)
	values.Add("filter_embeddable", strconv.FormatBool(o.FilterEmbeddable))
	values.Add("filter_playable", strconv.FormatBool(o.FilterPlayable))
	values.Add("sort", o.Sort)
	values.Add("Direction", o.Direction)
	return values
}

// Video lists
type VideoList struct {}

// list own videos
// ListOptions is not required
func (c *MeClient) ListVideos(option *ListOptions) (*VideoList, error){
	var values url.Values
	if option != nil {
		values = option.CreateValues()
	}

	videos := VideoList{}
	err := query("GET", "/me/videos", values, &videos)
	return &videos, err
}
