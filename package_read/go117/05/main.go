package main

import (
	"fmt"
	"strconv"
)

type Price int

type ShowPrice interface{
	type int,int8,int16,int64
	String() string
}

func(p Price)String()string{
	return strconv.Itoa(int(p))
}


func showPriceList[T ShowPrice](s []T)(ret []string){
	for idx,v := range s{
		ret = append(ret,"第" + strconv.Itoa(idx+1)+"道菜价格是："+v.String())
	}
	return ret
}

func main(){
	fmt.Println(showPriceList([]Price{12,34,55,66,77,88}))
}
