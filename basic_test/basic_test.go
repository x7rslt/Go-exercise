package basic_test

import (
	"fmt"
	"testing"
)

var str string
var str2 = ""

func TestVar(t *testing.T) {
	str3 := "" //注：只能定义的局部变量
	fmt.Println(str, str2, str3)
}

type Person struct {
	name string
}

func TestPointer(t *testing.T) {
	p := Person{"lehehe"}
	po := &p
	fmt.Println(po, *po) //注：&是指针，*是变量值

	p2 := &Person{"hahaha"}
	fmt.Println(p2.name) //指针成员变量可直接引用，编译器会自动隐形p转为（*p2）
	fmt.Println((*p2).name)
}

//[primary] 关于接口和类的说法，下面说法正确的是（）
//A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
//B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
//C. 类实现接口时，需要导入接口所在的包   //测试：将interface独立分包，不需要导入即可引用
//D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口
//
//参考答案：ABD

//这里说的类就是Go的结构体
type le interface {
	ha() string
}
type kaixin struct {
	name string
}

func (l *kaixin) ha(s string) string {
	return fmt.Sprintf("%s %s.", l.name, s)
}

func TestInterface(t *testing.T) {
	name := kaixin{"xiao"}
	status := name.ha("happy.")
	fmt.Println(status)
}

func TestStr(t *testing.T) {
	str := "123" + `abc`
	fmt.Println(str)
}

func TestFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println("lehehe")
	}
}
