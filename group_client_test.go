package vimeo

import "testing"

func TestListMyGroups(t *testing.T) {
	SetTokenFromEnv()
	cli := new(GroupClient)
	_, err := cli.ListMyGroups(nil)
	if err != nil {
		t.Errorf("Cant not get own groups. Because %s.", err)
	}
}
