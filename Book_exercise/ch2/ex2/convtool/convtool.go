//练习2.2：写一个类似于 cf 的通用的单位转换程序，从命令行参数或者标准输入（如果没有参数）获取数字，
//然后将每一个数字转换为以摄氏度和华氏度表示的温度，以英寸和米表示的长度单位，以磅和千克表示的重量，等等。
package convtool

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉ ", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g˚K", k)
}

type Inch float64
type Meter float64
type Usd float64
type Cny float64

func (u Usd) String() string {
	return fmt.Sprintf("%gUSD", u)
}
func (c Cny) String() string {
	return fmt.Sprintf("%gRMB", c)
}

func (i Inch) String() string {
	return fmt.Sprintf("%gin", i)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
