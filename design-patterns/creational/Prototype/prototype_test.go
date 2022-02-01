package Prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "出去浪啊"
	p:= ConcretePrototype{name: name}
	newProto := p.Clone()
	assert.Equal(t, name,newProto.Name())
}
