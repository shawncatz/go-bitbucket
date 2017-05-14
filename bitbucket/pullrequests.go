package bitbucket

import (
	"encoding/json"
	"fmt"
)

type PullRequestList struct {
	Response
	Values []PullRequest
}

type PullRequest struct {
	ID           int          `json:"id,omitempty"`
	Version      int          `json:"version"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	State        string       `json:"state"`
	Open         bool         `json:"open"`
	Closed       bool         `json:"closed"`
	Locked       bool         `json:"locked"`
	Created      int          `json:"createdDate"`
	Updated      int          `json:"updatedDate"`
	Links        LinkSliceMap `json:"links"`
	To           Ref          `json:"toRef"`
	From         Ref          `json:"fromRef"`
	Author       UserAssoc    `json:"author"`
	Reviewers    []*UserAssoc `json:"reviewers"`
	Participants []*UserAssoc `json:"participants"`
}

type UserAssoc struct {
	Role         string `json:"role"`
	Approved     bool   `json:"approved"`
	Status       string `json:"status"`
	LastReviewed string `json:"lastReviewedCommit"`
	User         User
}

type Ref struct {
	ID         string `json:"id"`
	Repository struct {
		Slug    string `json:"slug"`
		Name    string `json:"name"`
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
	} `json:"repository"`
}

type PullRequestsService service

func (s *PullRequestsService) List(project, repo string) (*PullRequestList, error) {
	list := &PullRequestList{}

	resp, err := s.client.Get("projects/"+project+"/repos/"+repo+"/pull-requests", nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving pull requests: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), list); err != nil {
		return nil, err
	}

	return list, nil
}
