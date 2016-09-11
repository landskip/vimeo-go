package vimeo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var _token string
var _url = "https://api.vimeo.com"

// Vimeo API response
type Response struct {
	Code int
	Body []byte
}

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

// Post vimeo API
func (c *Client) Post(path string, values url.Values) (Response, error) {
	req, err := http.NewRequest("POST",
		c.Config.BaseURL+path,
		strings.NewReader(values.Encode()))

	if err != nil {
		return Response{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+c.Config.AccessToken)

	return getDoResponseBody(http.Client{}, req)
}

func getDoResponseBody(cli http.Client, req *http.Request) (Response, error) {
	var result Response
	resp, err := cli.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.Body = b
	result.Code = resp.StatusCode
	return result, nil
}
