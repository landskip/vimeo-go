package vimeo

import (
	"testing"
)

func TestShowInfomation(t *testing.T) {
	_, err := Me.ShowInfomation()
	if err != nil {
		t.Errorf("Cant not get own information. %s", err)
	}
}

func TestListVideos(t *testing.T) {
	_, err := Me.ListVideos(nil)
	if err != nil {
		t.Errorf("Cant not get own videos. %s", err)
	}
}
