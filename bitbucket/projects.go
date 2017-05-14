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
	Key         string       `json:"key"`
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Avatar      string       `json:"avatar"`
	Public      bool         `json:"public,omitempty"`
	Type        string       `json:"type,omitempty"`
	Links       LinkSliceMap `json:"links,omitempty"`
}

type ProjectsService service

// List all of the accessible projects
// pagination is not currently working
// {Size:25, Limit:25, IsLastPage:false, Values:[]Project }
func (s *ProjectsService) List() (*ProjectList, error) {
	list := &ProjectList{}

	resp, err := s.client.Get("projects", nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving projects: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), list); err != nil {
		return nil, err
	}

	return list, nil
}

// Get retrieves the project of the given NAME
// bitbucket.Project{Key:"CHEF", ID:1234, Name:"Chef", Description:"Configuration Management", Public:false, Type:"NORMAL", Links:bitbucket.LinkSliceMap{"self":bitbucket.LinkSlice{bitbucket.Link{HREF:"https://stash.example.com/projects/CHEF"}}}}
func (s *ProjectsService) Get(name string) (*Project, error) {
	project := &Project{}

	resp, err := s.client.Get("projects/"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving project: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), project); err != nil {
		return nil, err
	}

	return project, err
}

// Create a project of the given NAME
func (s *ProjectsService) Create(key, name, description string) (*Project, error) {
	n := &Project{
		Key:         key,
		Name:        name,
		Description: description,
	}
	p := &Project{}

	resp, err := s.client.Post("projects", n)
	if err != nil {
		return nil, fmt.Errorf("retrieving project: %s", err)
	}

	if err := json.Unmarshal(resp.Body(), p); err != nil {
		return nil, err
	}

	return p, err
}

func (s *ProjectsService) Delete(name string) error {
	_, err := s.client.Delete("projects/" + name)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectsService) Update(name string, project *Project) (*Project, error) {
	p := &Project{}

	resp, err := s.client.Put("projects/"+name, project)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), p); err != nil {
		return nil, err
	}

	return p, nil
}
