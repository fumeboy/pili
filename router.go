package pili

import "fmt"

type fnr = func(in *router) *router
type matcher struct {
	name string
	fn func(i []byte, offset int) (o string,ok bool)
}

type router struct {
	matcher *matcher

	next []*router
	redirect *router
	optional bool

	token *token
	event func(ctx Context) ierr
}

func (this *router) clean() {
	toclean := []*router{}
	new_next := []*router{}
	for i,l := 0, len(this.next);i<l;i++{
		v := this.next[i]
		if v!= nil{
			if v.matcher == nil {
				toclean = append(toclean, v)
			}else{
				new_next = append(new_next, v)
				v.clean()
			}
		}
	}
	for i,l := 0, len(toclean);i<l;i++{
		v := toclean[i]
		v.clean()
		if v.matcher != nil {
			this.matcher = v.matcher
		}
		if v.token != nil {
			if this.token != nil {
				panic("token repeat")
			}
			this.token = v.token
		}
		if v.event != nil {
			if this.event != nil {
				panic("event repeat")
			}
			this.event = v.event
		}
		new_next = append(new_next, v.next...)
	}
	this.next = new_next
}

func (this *router) display(intend int) {
	intend_ := ""
	for i:=0;i<intend;i++{
		intend_+=" "
	}
	if this.token != nil {
		fmt.Println(intend_,"token:",this.token.identifier)
	}
	if this.matcher != nil {
		fmt.Println(intend_,"matcher:","\""+this.matcher.name+"\"")
	}else{
		fmt.Println(intend_,"matcher:","nil")
	}
	if this.redirect != nil{
		fmt.Println(intend_,"redirect:","\""+this.redirect.matcher.name+"\"")
	}
	for i,l := 0, len(this.next);i<l;i++{
		v := this.next[i]
		if v != nil {
			v.display(intend+1)
		}
		if i > 0 {
			fmt.Println(intend_,"---")
		}
	}
}