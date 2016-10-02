package vimeo

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var _token string
var _url = "https://api.vimeo.com"

// Available APIs
var (
	Me = new(MeClient)
)

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

// Query for Vimeo APIs (supported GET, POST, PATCH, DELET)
// return ResponseCode, error
func query(method, path string, values url.Values, v interface{}) error {
	var body io.Reader
	if method == "POST" || method == "PATCH" || method == "PUT" {
		body = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest(method, _url+path, body)
	if err != nil {
		return err
	}

	switch method {
	case "GET", "DELETE":
		if values != nil {
			req.URL.RawQuery = values.Encode()
		}
	case "PATCH", "POST", "PUT":
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("Authorization", "Bearer "+_token)

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//TODO: check response code
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
