package vimeo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Vimeo API client
type Client struct {
	Config Config
}

// Vimeo API response
type Response struct {
	Code int
	Body []byte
}

// Generate new VimeoClient object
func NewClient() (Client, error) {
	config, err := NewConfig()
	if err != nil {
		return Client{}, err
	}
	return Client{
		Config: config,
	}, nil
}

// Get vimeo API
func (c *Client) Get(path string, values url.Values) (Response, error) {
	req, err := http.NewRequest("GET",
		c.Config.BaseURL+path, nil)

	if err != nil {
		return Response{}, err
	}

	req.URL.RawQuery = values.Encode()
	req.Header.Add("Authorization", "Bearer "+c.Config.AccessToken)

	return getDoResponseBody(http.Client{}, req)
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
