package monstercat

import (
	"net/url"
	"net/http"
	"path"
)

type Client struct {
	BaseURL *url.URL
	UserAgent string
	httpClient *http.Client
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