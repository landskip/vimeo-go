package vimeo

import (
	"net/url"
	"strconv"
)

// Respose of GET list
type List struct {
	Page   int `json:"page"`
	Paging struct {
		First    string      `json:"first"`
		Last     string      `json:"last"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"paging"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// Request of GET List
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

// Create url values from list options
func (o *ListOptions) createValues() url.Values {
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
