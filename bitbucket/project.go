package bitbucket

import (
	"encoding/json"
	"fmt"
)

type ProjectList struct {
	Response
	Values []Project
}

type Project struct {
	Key         string
	ID          int `json:"id"`
	Name        string
	Description string
	Public      bool
	Type        string
	Links       LinkSliceMap
}

// Projects lists all of the accessible projects
// pagination is not currently working
// {Size:25, Limit:25, IsLastPage:false, Values:[]Project }
func (c *Client) Projects() (*ProjectList, error) {
	body, err := c.Execute("projects")
	if err != nil {
		return nil, fmt.Errorf("retrieving projects: %s", err)
	}

	resp := ProjectList{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

// Project retrieves the project of the given NAME
// bitbucket.Project{Key:"CHEF", ID:1234, Name:"Chef", Description:"Configuration Management", Public:false, Type:"NORMAL", Links:bitbucket.LinkSliceMap{"self":bitbucket.LinkSlice{bitbucket.Link{HREF:"https://stash.example.com/projects/CHEF"}}}}
func (c *Client) Project(name string) (*Project, error) {
	body, err := c.Execute("projects/%s", name)
	if err != nil {
		return nil, fmt.Errorf("retrieving project: %s", err)
	}

	resp := Project{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
