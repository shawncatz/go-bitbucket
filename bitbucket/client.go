package bitbucket

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/resty.v0"
)

type Client struct {
	cfg  *Config
	rest *resty.Client
	test bool
}

type Config struct {
	URL      string
	Username string
	Password string
}

func New(cfg *Config) (*Client, error) {
	r := resty.New()
	r.RetryCount = 0
	r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))

	client := &Client{cfg: cfg, rest: r}

	return client, nil
}

func (c *Client) SetTesting() {
	c.test = true
	c.rest.SetDisableWarn(true)
}

func (c *Client) Execute(response interface{}, format string, args ...interface{}) error {
	url := c.cfg.URL + "/" + fmt.Sprintf(format, args...)

	if c.cfg.Password == "" {
		return errors.New("password is empty")
	}

	resp, err := c.rest.R().
		SetBasicAuth(c.cfg.Username, c.cfg.Password).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return err
	}

	if resp.RawResponse.StatusCode != 200 {
		return fmt.Errorf("response %d code", resp.RawResponse.StatusCode)
	}

	err = json.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return err
	}

	return err
}

//
//func (c *Client) executeGetPages(url string) (*Response, error) {
//	response := Response{}
//
//}
//
//func (c *Client) executeGet(url string) (*Response, error) {
//	response := Response{}
//	r, err := resty.R().
//		SetBasicAuth(c.cfg.Username, c.cfg.password).
//		SetHeader("Accept", "application/json").
//		Get(url)
//
//	if err != nil {
//		return err
//	}
//
//	if r.RawResponse.StatusCode != 200 {
//		return fmt.Errorf("response %d code", r.RawResponse.StatusCode)
//	}
//
//	err = json.Unmarshal(r.Body(), &response)
//	if err != nil {
//		return err
//	}
//
//	return &response, nil
//}
