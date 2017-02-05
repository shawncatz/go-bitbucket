package bitbucket

type Link struct {
	HREF string `json:"href"`
}

type LinkSlice []Link
type LinkSliceMap map[string]LinkSlice
