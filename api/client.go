package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Client communicates with the https://swapi.dev REST API.
type Client struct {
	base string
	http *http.Client
}

// NewClient returns a configured HTTP Client
func NewClient(c *http.Client) *Client {
	if c == nil {
		c = http.DefaultClient
	}

	return &Client{base: "https://swapi.dev/api", http: c}
}

// NewRequest builds the http request for the given URL
func (c *Client) NewRequest(ctx context.Context, url string) (*http.Request, error) {
	if len(url) == 0 {
		return nil, errors.New("invalid empty-string url")
	}

	if url[0] == '/' { // Assume the user has given a relative path.
		url = c.base + url
	}

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return r.WithContext(ctx), nil
}

// Do the request
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.http.Do(r)
	if err != nil {
		return nil, err
	}

	// close the response body when we return
	defer func() {
		_ = resp.Body.Close()
	}()

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, fmt.Errorf("unable to parse JSON [%s %s]: %v", r.Method, r.URL.RequestURI(), err)
		}
	}

	return resp, nil
}
