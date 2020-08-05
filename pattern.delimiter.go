package pili

const (
	LF       = "\n"
	NotBlank = `\S`
)

func delimiter(text string) *matcher {
	return &matcher{
		name: text,
		fn:   func(i []byte) (o string, ok bool) {
			if text == string(i[:len(text)]) {
				return text, true
			}
			return "", false
		},
	}
}

func EOF() fnr {
	return func(in *router) *router {
		return &router{next: nil, matcher: &matcher{
			name: "EOF",
			fn: func(i []byte) (o string, ok bool) {
				if len(i) == 0{
					return "",true
				}else{
					return "", false
				}
			},
		}}
	}
}

func NOTBLANK() fnr {
	return func(in *router) *router {
		return &router{next: []*router{in}, matcher: &matcher{
			name: NotBlank,
			fn: func(i []byte) (o string, ok bool) {
				if i[0] != ' '{
					return "",true
				}else{
					return "", false
				}
			},
		}}
	}
}

func DELIMITER(text string) fnr {
	if len(text) == 0 {
		panic("非法的分割符")
	}
	return func(in *router) *router {
		return &router{next: []*router{in}, matcher: delimiter(text)}
	}
}