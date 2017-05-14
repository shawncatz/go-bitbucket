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

	Inbox        *InboxService
	Projects     *ProjectsService
	Repos        *ReposService
	PullRequests *PullRequestsService
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
	c.Repos = (*ReposService)(&c.common)
	c.PullRequests = (*PullRequestsService)(&c.common)
	c.Inbox = (*InboxService)(&c.common)

	return c
}

func (c *Client) setTesting() {
	c.test = true
	c.rest.SetDisableWarn(true)
}

func (c *Client) Execute(method string, uri string, body interface{}) (*resty.Response, error) {
	url := c.URL + "/" + uri

	if c.Password == "" {
		return nil, errors.New("password is empty")
	}

	req := c.rest.R().
		SetBasicAuth(c.Username, c.Password).
		SetHeader("Accept", "application/json")

	switch method {
	case resty.MethodPost, resty.MethodPut:
		req = req.SetBody(body)
	case resty.MethodGet:
		req = req.SetQueryParams(body.(map[string]string))
	default:
		// nothing
	}

	resp, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}

	if resp.RawResponse.StatusCode < 200 || resp.RawResponse.StatusCode > 299 {
		return resp, fmt.Errorf("%d:%s", resp.RawResponse.StatusCode, resp.RawResponse.Status)
	}

	return resp, nil
}

func (c *Client) Get(uri string, params map[string]string) (*resty.Response, error) {
	return c.Execute(resty.MethodGet, uri, params)
}

func (c *Client) Delete(uri string) (*resty.Response, error) {
	return c.Execute(resty.MethodDelete, uri, nil)
}

func (c *Client) Post(uri string, body interface{}) (*resty.Response, error) {
	return c.Execute(resty.MethodPost, uri, body)
}

func (c *Client) Put(uri string, body interface{}) (*resty.Response, error) {
	return c.Execute(resty.MethodPut, uri, body)
}
