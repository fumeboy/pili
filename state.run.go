package pili

func (this *router) match(s *state) (*router, ierr) {
	matchers := []*matcher{}
	mapp := map[string]*router{}
	for i, l := 0, len(this.next); i < l; i++ {
		m := this.next[i].matcher
		matchers = append(matchers, m)
		mapp[m.name] = this.next[i]
	}
	if this.token == nil {
		j := s.offset
		for ; j < s.len_in; j++ {
			if s.input[j] != ' ' {
				break
			}
		}
		for i, l := 0, len(matchers); i < l; i++ {
			if matchers[i].name == BLANK_ {
				if j > s.offset {
					s.offset = j
					s.history = append(s.history, matchers[i])
					return mapp[matchers[i].name], nil
				}
			}else{
				r, ok := matchers[i].fn(s.input, j)
				if ok {
					s.offset = len(r) + j
					s.history = append(s.history, matchers[i])
					return mapp[matchers[i].name], nil
				}
			}
		}
	} else {
		for j := s.offset; j <= s.len_in; j++ {
			for i, l := 0, len(matchers); i < l; i++ {
				r, ok := matchers[i].fn(s.input, j)
				if ok {
					s.add(this.token, string(s.input[s.offset:j]))
					s.offset = len(r) + j
					s.history = append(s.history, matchers[i])
					return mapp[matchers[i].name], nil
				}
			}
		}
	}
	if this.optional {
		return nil, nil
	}
	return nil, &find_sep_err{
		text:     s.input,
		i:        s.offset,
		history:  s.history,
		matchers: matchers,
		token:    this.token,
	}
}

func (s *state) RUN(f fnr) {
	rr := f(EOF()(nil))
	rr.clean()
	r := &router{next: []*router{rr}}
	//r.display(0)
	var err ierr
	for {
		if r == nil {
			break
		}
		if r.event != nil {
			if err := r.event(s.ctx); err != nil {
				break
			}
		}
		if r.redirect != nil {
			r = r.redirect
		} else if r.next == nil || len(r.next) == 0 {
			break
		}
		r, err = r.match(s)
		if err != nil || r == nil {
			break
		}
	}
	if err != nil {
		err.display()
	} else {
		s.Display()
	}
}
