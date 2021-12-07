package main
import (
	"fmt"
	"github.com/mcuadros/go-defaults"
	"time"
)

type ExampleBasic struct {
	Foo bool   `default:"true"` //<-- StructTag with a default key
	Bar string `default:"33"`
	Qux int8    `default:"100""`
	Dur time.Duration `default:"1m"`
}

func NewExampleBasic() *ExampleBasic {
	example := new(ExampleBasic)
	defaults.SetDefaults(example) //<-- This set the defaults values

	return example
}

func main(){
	test := NewExampleBasic()
	fmt.Println(test.Foo) //Prints: true
	fmt.Println(test.Bar) //Prints: 33
	fmt.Println(test.Qux) //Prints:
	fmt.Println(test.Dur) //Prints: 1m0s

}
