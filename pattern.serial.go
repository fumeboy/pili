package pili

func SERIAL(children ...fnr) fnr{
	return func(in *router) *router {
		for i := len(children)-1; i>=0; i--{
			in = children[i](in)
		}
		return in
	}
}
