package bitbucket

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type RepoList struct {
	Response
	Values []Repo
}

type Repo struct {
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

type ReposService service

func (s *ReposService) List(project string) (*RepoList, error) {
	list := &RepoList{}

	resp, err := s.client.Execute(resty.MethodGet, "projects/%s/repos", project)
	if err != nil {
		return nil, fmt.Errorf("retrieving repositories for project %s: %s", project, err)
	}

	if err := json.Unmarshal(resp.Body(), list); err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ReposService) Get(project, repo string) (*Repo, error) {
	repository := &Repo{}

	resp, err := s.client.Execute(resty.MethodGet, "projects/%s/repos/%s", project, repo)
	if err != nil {
		return nil, fmt.Errorf("retrieving repository for project %s repo %s: %s", project, repo, err)
	}

	if err := json.Unmarshal(resp.Body(), repository); err != nil {
		return nil, err
	}

	return repository, nil
}
