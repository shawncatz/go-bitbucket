package bitbucket

import (
	"encoding/json"
	"fmt"
)

type InboxService service

func (s *InboxService) List() (*PullRequestList, error) {
	list := &PullRequestList{}

	resp, err := s.client.Get("inbox/pull-requests", nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving inbox: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), list); err != nil {
		return nil, err
	}

	return list, nil
}
