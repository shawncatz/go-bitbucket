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

func (u *User) Mention() string {
	return "@" + u.Name
}

type UserAssoc struct {
	Role         string `json:"role"`
	Approved     bool   `json:"approved"`
	Status       string `json:"status"`
	LastReviewed string `json:"lastReviewedCommit"`
	User         User
}

func (ua *UserAssoc) Mention() string {
	return ua.User.Mention()
}
