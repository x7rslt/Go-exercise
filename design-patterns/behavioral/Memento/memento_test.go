package Memento

import (
	"fmt"
	"testing"
)

func TestNewMemento(t *testing.T) {
	n := NewNumber(2)
	n.Double()
	n.Double()
	memento := n.CreateMemento()
	n.Half()
	fmt.Println(n.value)
	n.ReinstateMemento(memento)
	fmt.Println(n.value)

}
