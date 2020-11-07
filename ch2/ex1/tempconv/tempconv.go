//添加类型、常量和函数到tempconv包中，处理以开尔文为单位（K）的温度值，0K=-273.15˚c，变化1K和变化1˚c是等价的。
package tempconv

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
