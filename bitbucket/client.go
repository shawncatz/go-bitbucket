package bitbucket

import (
	"fmt"
	"gopkg.in/resty.v0"
)

type Client struct {
	cfg  *Config
	rest *resty.Client
}

func New(config string) (*Client, error) {
	cfg, err := newConfig(config)

	if err != nil {
		return nil, err
	}

	r := resty.New()
	r.RetryCount = 0
	r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))

	client := &Client{cfg: cfg, rest: r}

	return client, nil
}

func (c *Client) Execute(format string, args ...interface{}) ([]byte, error) {
	url := c.cfg.URL + "/" + fmt.Sprintf(format, args...)

	resp, err := resty.R().
		SetBasicAuth(c.cfg.Username, c.cfg.password).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.RawResponse.StatusCode != 200 {
		return nil, fmt.Errorf("response %d code", resp.RawResponse.StatusCode)
	}

	return resp.Body(), err
}
