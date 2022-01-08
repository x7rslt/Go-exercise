package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}


type Rice struct {

}

func (r *Rice)Cook(){
	fmt.Println("This is rice.")
}

type Tomato struct {

}
func (t *Tomato)Cook(){
	fmt.Println("This is tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable()Lunch
}

type SimpleLunchFactory struct {

}
func NewSimpleLunchFactory() LunchFactory{
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory)CreateFood()Lunch{
	return &Rice{}
}
func (s *SimpleLunchFactory)CreateVegetable()Lunch{
	return &Tomato{}
}