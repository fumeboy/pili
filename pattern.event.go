package pili

type Context interface {

}

func EVENT(e func(ctx Context) ierr) fnr {
	return func(in *router) *router {
		return &router{next: []*router{in}, matcher: nil, event: e}
	}
}
