package Facade

import "fmt"

type CarModel struct {
}

func NewCarmodel() *CarModel {
	return &CarModel{}
}

func (m *CarModel)SetCarModel(){
	fmt.Println("Carmodel set model")
}

type CarEngine struct {

}

func NewCarEngine()*CarEngine{
	return &CarEngine{}
}

func (e *CarEngine)SetCarEngine(){
	fmt.Println("CarEngine set engine")
}

type CarBody struct {

}

func NewCarBody()*CarBody{
	return &CarBody{}
}
func (b *CarBody)SetCarBody(){
	fmt.Println("CarBody set body")

}


type CarFacade struct {
	model CarModel
	engine CarEngine
	body CarBody
}

func NewCarFacade()*CarFacade{
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (f *CarFacade)CreateCompleteCar(){
	f.model.SetCarModel()
	f.engine.SetCarEngine()
	f.body.SetCarBody()
}