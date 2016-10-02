package vimeo

import "testing"

func TestListMyVideos(t *testing.T) {
	SetTokenFromEnv()
	cli := new(VideoClient)
	_, err := cli.ListMyVideos(nil)
	if err != nil {
		t.Errorf("Cant not get own videos. Because %s.", err)
	}
}

// TODO: createVideo
// TODO: ShowMyVideo
