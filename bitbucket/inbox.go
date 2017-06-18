package bitbucket

import (
	"encoding/json"
	"fmt"
)

type InboxService service
type InboxCount struct {
	Count int
}

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

func (s *InboxService) Count() (int, error) {
	counts := &InboxCount{}
	resp, err := s.client.Get("inbox/pull-requests/count", nil)
	if err != nil {
		return 25, fmt.Errorf("retrieving inbox count: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), counts); err != nil {
		return 26, err
	}

	return counts.Count, nil
}
