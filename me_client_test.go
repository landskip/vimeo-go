package vimeo

import (
	"testing"
)

func TestShowInfomation(t *testing.T) {
    SetTokenFromEnv()
	cli := new(MeClient)
	_, err := cli.ShowInfomation()
	if err != nil {
		t.Errorf("Cant not get own information. Because %s.", err)
	}
}
