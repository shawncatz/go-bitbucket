package bitbucket

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"emailAddress"`
	Display string `json:"displayName"`
	Active  bool   `json:"active"`
	Slug    string `json:"slug"`
	Type    string `json:"type"`
}
