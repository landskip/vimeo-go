package vimeo

import (
	"testing"
	"net/http"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Errorf("Can not create client object.\n%v", err)
	}
}

func TestGet(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Errorf("Can not create client object.\n%v", err)
	}

	resp, err := c.Get("/videos/1", nil)
	if err != nil {
		t.Errorf("Can not call `GET /`.\n%v", err)
	}

	expectCode := http.StatusNotFound
	if resp.Code != expectCode {
		t.Errorf("Unexpected code.\nThe expected code is %d but code is %d",
			expectCode, resp.Code)
	}
}
