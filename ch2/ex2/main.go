//导入convtool包，转换开尔文温度
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"Go_exercise/ch2/ex2/convtool"
)

func tranfer(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "cf:%v\n", err)
		os.Exit(0)
	}
	f := convtool.Fahrenheit(t)
	c := convtool.Celsius(t)
	k := convtool.Kelvin(t)
	i := convtool.Inch(t)
	m := convtool.Meter(t)
	u := convtool.Usd(t)
	cn := convtool.Cny(t)
	fmt.Printf("%s = %s,%s = %s\n", f, convtool.FToC(f), c, convtool.CToF(c))
	fmt.Printf("%s = %s,%s = %s\n", k, convtool.KToC(k), c, convtool.CToK(c))
	fmt.Printf("%s = %s,%s = %s\n", i, convtool.IToM(i), m, convtool.MToI(m))
	fmt.Printf("%s = %s,%s = %s\n", u, convtool.UToC(u), cn, convtool.CToU(cn))
	os.Exit(0)
}

func main() {

	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			tranfer(arg)
		}
	}
	fmt.Println("Please input something.")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		tranfer(input.Text())
	}
}
