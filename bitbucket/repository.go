package bitbucket

import (
	"fmt"
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
	resp := &RepositoryList{}
	err := c.Execute(*resp, "projects/%s/repos", project)
	if err != nil {
		return nil, fmt.Errorf("retrieving repositories for project %s: %s", project, err)
	}

	return resp, nil
}

func (c *Client) Repository(project, repo string) (*Repository, error) {
	resp := &Repository{}
	err := c.Execute(*resp, "projects/%s/repos/%s", project, repo)
	if err != nil {
		return nil, fmt.Errorf("retrieving repository for project %s repo %s: %s", project, repo, err)
	}

	return resp, nil
}
