//软件包fmt使用与C的printf和scanf类似的功能实现格式化的I / O。格式“动词”是从C派生的，但更简单。
package fmt_test

import (
	"fmt"
	"testing"
)

type Lehe struct {
}

type X string

func (x X) String() string { return fmt.Sprintf("<%s>", string(x)) }

func TestFmt(t *testing.T) {
	var l Lehe
	n := "leheheh\n"
	i := 10
	j := fmt.Sprintf("%1d\n", i)
	fmt.Println(len(j), j) //添加换行符
	fmt.Printf("%p\n", &n)
	fmt.Printf("%v\n", &n)
	fmt.Printf("%#v\n", l)

	g := 12.345
	fmt.Printf("%.1g\n", g)
	fmt.Print(g) //默认支持%v输出

	var m interface{} = 123
	fmt.Printf("%v\n", m) //无论任何verb，接口输出其内部值，而不是接口本身

	var v X
	v = "test"
	fmt.Printf("%[1]T", v)
	//v.String()

}

func TestFmtScan(t *testing.T) {
	integer := 23
	fmt.Println(integer)
	fmt.Printf("%T %T %#v\n", integer, &integer, integer)
	// Result: int *int
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %T %#v %#v", greats, greats, greats, &greats)
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)

	fmt.Println("", "test", "test2", "", "test3")
}
