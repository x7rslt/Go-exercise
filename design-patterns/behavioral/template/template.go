package template

import "fmt"

type WorkerInterface interface {
	GetUp()
	Work()
	Sleep()
}


type Worker struct {
	WorkerInterface
}

func NewWorker(w WorkerInterface)*Worker{
	return &Worker{w}
}

func (w *Worker)Daily(){
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct {

}

func (c *Coder)GetUp(){
	fmt.Println("Coder GetUp")

}

func (c *Coder)Work(){
	fmt.Println("Coder Coding")
}

func (c *Coder)Sleep(){
	fmt.Println("Coder Sleep")
}