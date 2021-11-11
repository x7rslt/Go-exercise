//软件包fmt使用与C的printf和scanf类似的功能实现格式化的I / O。格式“动词”是从C派生的，但更简单。
package fmt_test

import (
	"fmt"
	"io"
	"os"
	"strings"
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

func TestFmtFunc(t *testing.T) {
	a := 1
	b := 0
	fmt.Println("This is a test!")
	fmt.Print(fmt.Errorf("%d can't divided %b!", a, b))
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint:%v\n", err)
	}
	fmt.Print(n, "bytes writen.\n")
	content := "This"
	r := strings.NewReader(content)
	var i string
	n, err = fmt.Fscan(r, &i)
	fmt.Println(i)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscan:%v\n", err)
	}
	fmt.Println(n)

	//Fscanf扫描从r读取的文本，将连续的以空格分隔的值存储到由格式确定的连续的参数中。它返回成功解析的项目数。输入中的换行符必须与格式中的换行符匹配。
	var (
		j int
		B bool
		s string
	)
	r = strings.NewReader("5 true gophers")
	n, err = fmt.Fscanf(r, "%d %t %s\n", &j, &B, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf:%v\n", err)
	}
	fmt.Println(j, B, s)
	fmt.Println(n)
}
func TestPointer(t *testing.T) {
	value := 1
	fmt.Println(&value)
	n, err := fmt.Printf("%T1\n", &value)
	fmt.Println(n, err)
	p := &value
	*p = 2
	fmt.Println(p)
	fmt.Println(value)

	const name, age = "kim", 12
	s := fmt.Sprintf("%s is %d years old.", name, age)
	fmt.Println(s)
	io.WriteString(os.Stdout, s)
}

func TestFmtSscan(t *testing.T) {
	var (
		a string
		b int
	)
	fmt.Sscanf("Kim is 12 years old.", "%s is %d years old.", &a, &b)
	fmt.Println(a, b)
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintln(p.name, p.age)
}

func TestFmtstring(t *testing.T) {
	var xiao Person = Person{
		name: "xiaoyang",
		age:  18,
	}
	fmt.Printf("haha %s\n", xiao)

}
