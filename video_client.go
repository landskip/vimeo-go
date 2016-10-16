package vimeo

import (
	"net/http"
	"strconv"
)

// Video client
type VideoClient struct{}

// Create video
// VideoOptions is not required
func (c *VideoClient) CreateVideo(option *VideoOptions) (*Video, error) {
	var data QueryRequest
	if option != nil {
		data = option
	}
	video := Video{}
	err := query(http.MethodPost, "/me/videos", data, &video)
	return &video, err
}

// Show own video
func (c *VideoClient) ShowMyVideo(videoID int) (*Video, error) {
	video := Video{}
	err := query(http.MethodGet, "/me/videos/"+strconv.Itoa(videoID), nil, &video)
	return &video, err
}

// list own videos
// ListOptions is not required
func (c *VideoClient) ListMyVideos(option *ListOptions) (*VideoList, error) {
	var data QueryRequest
	if option != nil {
		data = option
	}
	videos := VideoList{}
	err := query(http.MethodGet, "/me/videos", data, &videos)
	return &videos, err
}
