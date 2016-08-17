package router

type Router interface {
	Route(string) (string, error)
}
