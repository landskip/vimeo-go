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
