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

func (o *ListOptions) GetJSONBytes() ([]byte, error) {
	return nil, nil
}

func (o *ListOptions) GetValues() url.Values {
	values := url.Values{}
	values.Add("page", strconv.Itoa(o.Page))
	if p := o.PerPage; p != 0 {
		values.Add("per_page", strconv.Itoa(p))
	}
	if q := o.Query; q != "" {
		values.Add("query", q)
	}
	if f := o.Filter; f != "" {
		values.Add("filter", f)
	}
	if f := o.FilterEmbeddable; f != false {
		values.Add("filter_embeddable", strconv.FormatBool(f))
	}
	if f := o.FilterPlayable; f != false {
		values.Add("filter_playable", strconv.FormatBool(f))
	}
	if s := o.Sort; s != "" {
		values.Add("sort", s)
	}
	if d := o.Direction; d != "" {
		values.Add("Direction", d)
	}
	return values
}
