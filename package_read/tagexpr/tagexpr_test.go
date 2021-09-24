package tagexpr_test

import (
	"fmt"
	tagexpr "github.com/bytedance/go-tagexpr/v2"
	"testing"
)

//An interesting go struct tag expression syntax for field validation, etc.
func TestTage(t *testing.T){
	Example()
}
func Example() {
	type T struct {
		A  int             `tagexpr:"$<0||$>=100"`
		B  string          `tagexpr:"len($)>1 && regexp('^\\w*$')"`
		C  bool            `tagexpr:"expr1:(f.g)$>0 && $; expr2:'C must be true when T.f.g>0'"`
		d  []string        `tagexpr:"@:len($)>0 && $[0]=='D'; msg:sprintf('invalid d: %v',$)"`
		e  map[string]int  `tagexpr:"len($)==$['len']"`
		e2 map[string]*int `tagexpr:"len($)==$['len']"`
		f  struct {
			g int `tagexpr:"$"`
		}
	}

	vm := tagexpr.New("tagexpr")
	t := &T{
		A:  107,
		B:  "abc",
		C:  true,
		d:  []string{"x", "y"},
		e:  map[string]int{"len": 1},
		e2: map[string]*int{"len": new(int)},
		f: struct {
			g int `tagexpr:"$"`
		}{1},
	}

	tagExpr, err := vm.Run(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(tagExpr.Eval("A"))
	fmt.Println(tagExpr.Eval("B"))
	fmt.Println(tagExpr.Eval("C@expr1"))
	fmt.Println(tagExpr.Eval("C@expr2"))
	if !tagExpr.Eval("d").(bool) {
		fmt.Println(tagExpr.Eval("d@msg"))
	}
	fmt.Println(tagExpr.Eval("e"))
	fmt.Println(tagExpr.Eval("e2"))
	fmt.Println(tagExpr.Eval("f.g"))

	// Output:
	// true
	// true
	// true
	// C must be true when T.f.g>0
	// invalid d: [x y]
	// true
	// false
	// 1
}