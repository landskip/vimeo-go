package vimeo

import (
	"net/url"
	"strconv"
)

// Video client
type VideoClient struct{}

// Create video
// VideoOptions is not required
func (c *VideoClient) CreateVideo(option *VideoOptions) (*Video, error) {
	var values url.Values
	if option != nil {
		values = option.createValues()
	}
	video := Video{}
	err := query("POST", "/me/videos", values, &video)
	return &video, err
}

// Show own video
func (c *VideoClient) ShowMyVideo(videoID int) (*Video, error) {
	video := Video{}
	err := query("GET", "/me/videos/"+strconv.Itoa(videoID), nil, &video)
	return &video, err
}

// list own videos
// ListOptions is not required
func (c *VideoClient) ListMyVideos(option *ListOptions) (*VideoList, error) {
	var values url.Values
	if option != nil {
		values = option.createValues()
	}
	videos := VideoList{}
	err := query("GET", "/me/videos", values, &videos)
	return &videos, err
}
