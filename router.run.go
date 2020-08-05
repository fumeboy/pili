package pili

func (this *router) match(s *state) (*router,ierr) {
	text := s.input[s.offset:]
	ltext := len(text)
	matchers := []*matcher{}
	mapp := map[string]*router{}
	for i,l := 0, len(this.next); i<l;i++{
		m := this.next[i].matcher
		matchers = append(matchers, m)
		mapp[m.name] = this.next[i]
	}
	if this.token == nil {
		j := 0
		for ;j<ltext;j++{
			if text[j] != ' '{break}
		}
		remains := text[j:]
		for i,l := 0, len(matchers);i<l;i++{
			r, ok := matchers[i].fn(remains)
			if ok {
				s.offset += len(r)+j
				return mapp[matchers[i].name], nil
			}
		}
	}else{
		var remains []byte
		for j := 0;j<=ltext;j++{
			remains = text[j:]
			for i,l := 0, len(matchers);i<l;i++{
				r, ok := matchers[i].fn(remains)
				if ok {
					s.offset += len(r)+j
					s.add(this.token, string(text[:j]))
					return mapp[matchers[i].name],nil
				}
			}
		}
	}
	if this.optional {
		return nil, nil
	}
	return nil, &find_sep_err{
		text: s.input,
		i:          s.offset,
		matchers: matchers,
		token: this.token,
	}
}

func RUN(s *state, f fnr) {
	rr := f(EOF()(nil))
	rr.clean()
	r := &router{next: []*router{rr}}
	//r.display(0)
	var err ierr
	for {
		if r == nil {
			break
		}
		if r.event != nil{
			if err := r.event(s.ctx); err != nil {
				break
			}
		}
		if r.redirect != nil {
			r = r.redirect
		}else if r.next == nil || len(r.next) == 0 {
			break
		}
		r, err = r.match(s)
		if err != nil || r == nil{
			break
		}
	}
	if err != nil {
		err.display()
	}else{
		s.Display()
	}
}