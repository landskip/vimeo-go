package vimeo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var _token string
var _url = "https://api.vimeo.com"

// Set access token
func SetToken(token string) {
	_token = token
}

// Set VIMEO_TOKEN to token in environment
func SetTokenFromEnv() error {
	_token = os.Getenv("VIMEO_TOKEN")
	if _token == "" {
		return errors.New("VIMEO_TOKEN is not set.")
	}
	return nil
}

// Request for query method
type QueryRequest interface {
	GetValues() url.Values         // for GET & DELETE
	GetJSONBytes() ([]byte, error) // for POST & PATCH & PUT
}

// Query for Vimeo APIs (supported GET, POST, PATCH, DELETE, PUT)
// return response code & error
func query(method, path string, data QueryRequest, result interface{}) error {
	var body io.Reader
	if data != nil && (method == http.MethodPost || method == http.MethodPatch || method == http.MethodPut) {
		if j, err := data.GetJSONBytes(); err != nil {
			return err
		} else if j != nil {
			body = bytes.NewBuffer(j)
		}
	}

	req, err := http.NewRequest(method, _url+path, body)
	if err != nil {
		return err
	}

	switch method {
	case http.MethodGet, http.MethodDelete:
		if data != nil {
			if v := data.GetValues(); v != nil {
				req.URL.RawQuery = v.Encode()
			}
		}
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", "Bearer "+_token)

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code != http.StatusOK && code != http.StatusCreated {
		return errors.New(http.StatusText(code))
	}

	if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else if result != nil {
		return json.Unmarshal(b, result)
	}
	return nil
}
