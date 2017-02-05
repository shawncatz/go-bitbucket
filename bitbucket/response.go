package bitbucket

type Response struct {
	Size       int
	Limit      int
	IsLastPage bool
	Values     []interface{}
}
