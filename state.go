package pili

import "fmt"

type state struct {
	input []byte
	offset int
	len_in int

	history []*matcher
	ctx Context

	matched map[string][]string
}

func (this *state) add(token *token, data string) ierr {
	if token.check != nil {
		if err := token.check(data); err != nil {
			return err
		}
	}
	if this.matched[token.identifier] == nil{
		this.matched[token.identifier] = []string{}
	}
	this.matched[token.identifier] = append(this.matched[token.identifier], data)
	return nil
}
func (this *state) get(id string)[]string {
	r := this.matched[id]
	delete(this.matched, id)
	return r
}
func (this *state) Display() {
	fmt.Println("input:",`"`+string(this.input)+`"`)
	for k,v := range this.matched{
		fmt.Println("  token <"+k+">","=>")
		for _,vv := range v{
			fmt.Println("  -", `'`+string(vv)+`'`)
		}
	}
	fmt.Println("=== OVER")
}

func NEWstate(ctx Context, text string) *state {
	s := &state{input: []byte(text), ctx: ctx, matched: map[string][]string{}}
	s.len_in = len(s.input)
	return s
}
