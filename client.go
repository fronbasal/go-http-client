package go_http_client

import "net/http"

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
func New(url string) (Client, error) {
	// Create an empty HTTP Client
	c := &http.Client{}
	// Create a new http.Request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Client{}, err
	}
	// Set a custom UA
	r.Header.Set("User-Agent", UserAgent)
	// Allow GZIP
	r.Header.Set("Accept-Encoding", "gzip")

	return Client{Client: c, Request: r}, nil
}

// Do dispatches the Request with the Client
func (c *Client) Do() (resp *http.Response, err error) {
	return c.Client.Do(c.Request)
}

// SetBasicAuth sets a username and a password to the Request object
func (c *Client) SetBasicAuth(username, password string) {
	c.Request.SetBasicAuth(username, password)
}