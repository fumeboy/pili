package pili

type token struct {
	identifier string
	check func(text string)ierr
}
func TOKEN(text string, check func(text string)ierr) fnr {
	if len(text) == 0 {
		panic("非法的标识符")
	}
	return func(in *router) *router {
		return &router{next: []*router{in}, matcher: nil, token: &token{identifier: text, check: check}}
	}
}