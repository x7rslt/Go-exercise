//添加类型、常量和函数到tempconv包中，处理以开尔文为单位（K）的温度值，0K=-273.15˚c，变化1K和变化1˚c是等价的。
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
