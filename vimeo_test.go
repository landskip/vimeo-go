package vimeo

import (
	"testing"
)

func TestSetToken(t *testing.T) {
	expectToken := "test-token"
	SetToken(expectToken)

	if _token != expectToken {
		t.Errorf("Token is invalid. expected token was %s", expectToken)
	}
}

func TestSetTokenFromEnv(t *testing.T) {
}

func TestQuery(t *testing.T) {
	// TODO: GET method
	// TODO: POST method
	// TODO: PUT method
	// TODO: DELETE method
}
