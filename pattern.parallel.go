package pili

func PARALLEL(paths ...fnr) fnr {
	return func(in *router) *router {
		inputs := []*router{}
		for _,v:=range paths{
			inputs = append(inputs, v(in))
		}
		return &router{next: inputs, matcher: nil}
	}
}
func OPTION(paths ...fnr) fnr {
	return func(in *router) *router {
		inputs := []*router{}
		for _,v:=range paths{
			inputs = append(inputs, v(in))
		}
		inputs = append(inputs, in)
		return &router{next: inputs, matcher: nil, optional: true}
	}
}