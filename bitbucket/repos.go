package bitbucket

import (
	"encoding/json"
	"fmt"
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

type ReposListOptions struct {
	ListOptions
}

type ReposService service

func (s *ReposService) List(project string, options *ReposListOptions) (*RepoList, error) {
	list := &RepoList{}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", options.Limit),
		"start": fmt.Sprintf("%d", options.Start),
	}

	resp, err := s.client.Get("projects/"+project+"/repos", params)
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

	resp, err := s.client.Get("projects/"+project+"/repos/"+repo, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving repository for project %s repo %s: %s", project, repo, err)
	}

	if err := json.Unmarshal(resp.Body(), repository); err != nil {
		return nil, err
	}

	return repository, nil
}

func (s *ReposService) Create(project, key, name string) (*Repo, error) {
	n := &Repo{
		Slug: key,
		Name: name,
		Project: Project{
			Key: project,
		},
	}

	resp, err := s.client.Post("projects/"+project+"/repos", n)
	if err != nil {
		return nil, fmt.Errorf("creating repo: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), n); err != nil {
		return nil, err
	}

	return n, nil
}
