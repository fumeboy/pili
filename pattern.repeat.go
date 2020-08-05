package pili

func REPEAT(left string, repeated fnr, sep string, right string) fnr {
	return func(in *router) *router {
		var a *router = &router{}
		var b *router
		if left != "" {
			a.next = []*router{}
			a.matcher = delimiter(left)
		}
		if right == ""{
			b = &router{
				matcher:  nil,
				next:     []*router{
					{
						matcher: delimiter(sep),
						redirect: a,
					},
					in,
				},
			}
		}else{
			b = &router{
				matcher:  nil,
				next:     []*router{
					{
						matcher: delimiter(sep),
						redirect: a,
					},
					DELIMITER(right)(in),
				},
			}
		}
		if left == "" {
			a.redirect = repeated(b)
		}else{
			a.next = append(a.next, repeated(b))
		}
		return a
	}
}
