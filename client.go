package monstercat

import (
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	BaseURL    url.URL
	httpClient *http.Client
}

//New creates a new monstercat client
func New(baseURL string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		*u,
		http.DefaultClient,
	}, nil
}

func (c *Client) NewRequest(method, requestPath string) (*http.Request, error) {
	url := c.BaseURL
	url.Path = path.Join(url.Path, requestPath)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return req, err
}
