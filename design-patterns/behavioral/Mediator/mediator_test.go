package Mediator

import "testing"

func TestConcreteMediator(t *testing.T) {
	mediator := NewMediator()
	mediator.Bill.Respond()
}