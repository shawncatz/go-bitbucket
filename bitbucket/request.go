package bitbucket

type Request struct {
	project     string
	repository  string
	pullRequest string

	client *Client
}

// Not sure I'm going to do this
//
//func (r *Request) Project(project string) *Request {
//	r.project = project
//	return r
//}
//
//func (r *Request) Repository(repo string) *Request {
//	r.repository = repo
//	return r
//}
//
//func (r *Request) Go() (*Response, error) {
//	if r.project == "" {
//		return r.client.Projects()
//	}
//
//	if r.repository == "" {
//		// do project list
//	}
//
//	if r.pullRequest == "" {
//		// do pull request list
//	}
//
//	return nil, nil
//}
