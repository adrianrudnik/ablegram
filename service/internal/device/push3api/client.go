package push3api

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Uri   string
	Token string

	client *http.Client
}

func NewClient(uri string, token string) *Client {
	return &Client{
		Uri:   strings.TrimRight(strings.TrimSpace(uri), "/"),
		Token: token,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

// IsAuthorized checks if the client is authorized to access the Push 3 API.
func (c *Client) IsAuthorized() (bool, error) {
	if c.Token == "" {
		return false, nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.Uri, "/api/v1/access-allowed"), nil)
	if err != nil {
		return false, err
	}

	resp, err := c.executeRequest(req, true)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// executeRequest executes a request to the Push 3 API, with a possible token, if set.
func (c *Client) executeRequest(req *http.Request, authenticate bool) (*http.Response, error) {
	if authenticate && c.Token != "" {
		req.AddCookie(&http.Cookie{
			Name:     "Ableton-Challenge-Response-Token",
			Value:    c.Token,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
	}

	return c.client.Do(req)
}
