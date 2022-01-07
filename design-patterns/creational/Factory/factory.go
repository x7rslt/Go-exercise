package Factory

import "fmt"

type Company interface {
	Service()
}

type Bytedance struct{

}

func (b *Bytedance)Service(){
	fmt.Println("字节跳动业务包含：抖音、今日头条、皮皮虾...")
}

type Tencent struct {

}

func (t *Tencent)Service(){
	fmt.Println("腾讯业务包含：微信、王者荣耀...")
}

func NewCompany(companyname string)Company{
	switch companyname {
	case "b":
		return &Bytedance{}
	case "t":
		return &Tencent{}
	}
	return nil
}