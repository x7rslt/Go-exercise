package Visitor

import "testing"

func TestElement(t *testing.T) {
	e := new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}
