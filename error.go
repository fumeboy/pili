package pili

import "fmt"

type ierr interface {
	display()
}

type find_sep_err struct {
	text []byte
	i int // 文本进行位置
	matchers []*matcher
	token *token
}

func (f find_sep_err) display() {
	fmt.Println("index:", f.i, `"`+string(f.text[f.i]) + `"`)
	fmt.Println("matchers:")
	for _,v := range f.matchers {
		fmt.Println("  ", `"`+string(v.name)+`"`)
	}
	if f.token != nil{
		fmt.Println("token:", f.token.identifier)
	}
}

type token_check_error struct {
	i int // 文本进行位置
	text string // 受检测文本
	token *token
}

func (this *token_check_error) display() {
	fmt.Println("index:",this.i)
	fmt.Println("text:", this.text)
	fmt.Println("token:", this.token.identifier)
}


