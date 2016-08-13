package vimeo

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	_, err := NewConfig()
	if err != nil {
		t.Errorf("Can not create config object.\n%v", err)
	}
}
