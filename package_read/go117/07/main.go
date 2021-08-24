package main

import "fmt"

type addable interface{
	comparable
}

type set[T addable]map[T]struct{}
func (s set[T])add(v T){s[v]=struct{}{}}
func (s set[T])contains(v T)bool{
	_,ok:=s[v]
	return ok
}

func (s set[T])len()int{return len(s)}
func(s set[T])delete(v T){delete(s,v)}
func (s set[T])iterate(f func(T)){for v:= range s{f(v)}}

func main(){
	s:=make(set[string])
	s.add("红烧肉")
	s.add("清蒸鱼")
	s.add("大闸蟹")
	fmt.Println("%v\n",s)
	if s.contains("大闸蟹"){
		fmt.Println("包含大闸蟹")
	}else{
		s.add("大闸蟹")
		fmt.Println("添加了大闸蟹")
	}
	fmt.Println("the len of set:%d\n", s.len())
	s.iterate(func(x string){fmt.Println("你点的菜："+x)})

}