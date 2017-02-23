package bitbucket

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type RepositoryList struct {
	Response
	Values []Repository
}

type Repository struct {
	Slug          string
	ID            int `json:"id"`
	Name          string
	ScmID         string `json:"scmId"`
	State         string
	StatusMessage string
	Forkable      bool
	Project       Project
	Public        bool
	Links         LinkSliceMap
}

func (c *Client) Repositories(project string) (*RepositoryList, error) {
	list := &RepositoryList{}

	resp, err := c.Execute(resty.MethodGet, "projects/%s/repos", project)
	if err != nil {
		return nil, fmt.Errorf("retrieving repositories for project %s: %s", project, err)
	}

	if err := json.Unmarshal(resp.Body(), list); err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Client) Repository(project, repo string) (*Repository, error) {
	repository := &Repository{}

	resp, err := c.Execute(resty.MethodGet, "projects/%s/repos/%s", project, repo)
	if err != nil {
		return nil, fmt.Errorf("retrieving repository for project %s repo %s: %s", project, repo, err)
	}

	if err := json.Unmarshal(resp.Body(), repository); err != nil {
		return nil, err
	}

	return repository, nil
}
