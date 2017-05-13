package bitbucket

import (
	"errors"
	"fmt"
	"gopkg.in/resty.v0"
	"net/http"
)

type Client struct {
	URL      string
	Username string
	Password string

	http   *http.Client
	common service // re-use single struct
	rest   *resty.Client
	test   bool

	Projects *ProjectsService
	Repos    *ReposService
}

type service struct {
	client *Client
}

func NewClient(URL, username, password string) *Client {
	r := resty.New()
	r.RetryCount = 0
	r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))

	c := &Client{URL: URL, Username: username, Password: password, rest: r}
	c.common.client = c

	c.Projects = (*ProjectsService)(&c.common)

	return c
}

func (c *Client) setTesting() {
	c.test = true
	c.rest.SetDisableWarn(true)
}

func (c *Client) Execute(method string, format string, args ...interface{}) (*resty.Response, error) {
	url := c.URL + "/" + fmt.Sprintf(format, args...)

	if c.Password == "" {
		return nil, errors.New("password is empty")
	}

	resp, err := c.rest.R().
		SetBasicAuth(c.Username, c.Password).
		SetHeader("Accept", "application/json").
		Execute(method, url)

	if err != nil {
		return nil, err
	}

	if resp.RawResponse.StatusCode != 200 {
		return nil, fmt.Errorf("response %d code", resp.RawResponse.StatusCode)
	}

	return resp, nil
}
