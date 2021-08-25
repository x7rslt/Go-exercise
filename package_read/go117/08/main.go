package main

import "fmt"

//泛型队列
type queue[T any] []T

func (q *queue[T])enqueue(v T){
	*q = append(*q,v)


}
func (q *queue[T])dequeue()(T,bool){
	if len(*q)==0{
		var zero T
		return zero,false
	}
	r:= (*q)[0]
	*q=(*q)[1:]
	return r,true
}

func main(){
	q:=new(queue[string])
	q.enqueue("红烧肉")
	q.enqueue("清蒸鱼")
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())

}
