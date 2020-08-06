package pili

const (
	LF       = "\n"
	NotBlank = `\NotBlank`
	BLANK_ = `\BLANK`
)

func delimiter(text string) *matcher {
	return &matcher{
		name: text,
		fn:   func(i []byte, offset int) (o string, ok bool) {
			if len(i) >= offset+len(text) && text == string(i[offset:offset+len(text)]) {
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
			fn: func(i []byte, offset int) (o string, ok bool) {
				if len(i) == offset{
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
			fn: func(i []byte, offset int) (o string, ok bool) {
				if i[offset] != ' '{
					return "",true
				}else{
					return "", false
				}
			},
		}}
	}
}
func BLANK() fnr {
	return func(in *router) *router {
		return &router{next: []*router{in}, matcher: &matcher{
			name: BLANK_,
			fn: func(i []byte, offset int) (o string, ok bool) {
				j := offset
				l := len(i)
				for ;j<l;j++{
					if i[j] != ' '{break}
				}
				if j > offset{
					return string(i[offset:j]),true
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