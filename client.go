package go_http_client

import (
	"github.com/pkg/errors"
	"net/http"
)

// set a custom UA
const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/53.0.2785.143 " +
	"Safari/537.36"

// Client holds the Client and Request structs from the default HTTP lib
type Client struct {
	Client  *http.Client
	Request *http.Request
}

// New creates a new HTTP Client with a Client and a Request
func New(url string) (*Client, error) {
	// Create an empty HTTP Client
	c := &http.Client{}
	// Create a new http.Request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &Client{}, err
	}
	// Set a custom UA
	r.Header.Set("User-Agent", UserAgent)
	// Allow GZIP
	r.Header.Set("Accept-Encoding", "gzip")

	return &Client{Client: c, Request: r}, nil
}

// NewMethod creates a new HTTP Client with a specified HTTP Method
func NewMethod(url, method string) (*Client, error) {
	client, err := New(url)
	if err != nil {
		return &Client{}, nil
	}
	client.Request.Method = method
	return client, err
}

// Do dispatches the Request with the Client and checks if it returned status 200
func (c *Client) Do() (resp *http.Response, err error) {
	resp, err = c.Client.Do(c.Request)
	if resp.StatusCode < 200 || resp.StatusCode > 201 {
		return resp, errors.New("Did not receive status 200")
	}
	return resp, err
}

// SetBasicAuth sets a username and a password to the Request object
func (c *Client) SetBasicAuth(username, password string) {
	c.Request.SetBasicAuth(username, password)
}
