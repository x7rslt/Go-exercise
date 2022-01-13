package Bridage

import (
	"fmt"
	"time"
)

type Draw interface{
	DrawCircle(radius,x,y int)
}

type RedCircle struct {

}

func (r *RedCircle)DrawCircle(radius,x,y int){
	fmt.Println("RedCircle radius、x、y：",radius,x,y)

}
type YellowCircle struct{

}

func (c *YellowCircle)DrawCircle(radius,x,y int){
	fmt.Println("YellowCircle radius,x,y:",radius,x,y)
}

type Shape struct{
	draw Draw
}

func (s *Shape)Shape(d Draw){
	s.draw = d
	time.Now().Unix()
}

type Circle struct {
	shape Shape
	x int
	y int
	radius int
}

func (c *Circle)Constructor(x,y,radius int,d Draw){
	c.shape.Shape(d)
	c.x = x
	c.y = y
	c.radius = radius
}

func (c *Circle)Draw(){
	c.shape.draw.DrawCircle(c.radius,c.x,c.y)
}

