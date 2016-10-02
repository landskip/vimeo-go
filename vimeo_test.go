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

func TestTokenFromEnv(t *testing.T)  {
    err := SetTokenFromEnv()
    if err != nil {
        t.Errorf("Token can get from environment. Please set VIMEO_TOKEN.")
    }
}

// TODO: mock
func TestQuery(t *testing.T) {
	// TODO: GET method
	// TODO: POST method
	// TODO: PUT method
	// TODO: DELETE method
}
