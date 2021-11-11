package iota_test

import (
	"fmt"
	"testing"
)

const(
	c0=iota
	c1=iota
	c2=iota
)

const(
	d0 =iota
	d1
	d2
)

const(
	e0 =iota+1
	e1
	e2
)

//Skip value
const(
	f0 = iota+1
	_
	f1
	f2
)

func TestIotaEnum(t *testing.T){
	fmt.Println(c0,c1,c2)// "0 1 2"
	fmt.Println(d0,d1,d2)// "0 1 2"
	fmt.Println(e0,e1,e2)// "1 2 3"
	fmt.Println(f0,f1,f2)// "1 3 4"
}

//Complete enum type with strings [best practice]
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func TestTypeIotaEnum(t *testing.T){
	var d Direction = North
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
	// Output: North goes up.
}