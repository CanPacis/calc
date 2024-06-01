package main

import (
	"fmt"
	"strings"
)

type Object interface {
	GetValue() interface{}
}

type Float64Object struct {
	Value float64
}

func (obj Float64Object) GetValue() interface{} {
	return obj.Value
}

type ConstantPool struct {
	Values  [512]Object
	pointer int
}

func (p ConstantPool) Has(value Object) int {
	for i, v := range p.Values {
		if v == nil {
			return -1
		}

		if v.GetValue() == value.GetValue() {
			return i
		}
	}

	return -1
}

func (p *ConstantPool) Add(value Object) int {
	has := p.Has(value)

	if has >= 0 {
		return has
	}

	p.Values[p.pointer] = value
	p.pointer++

	return p.pointer - 1
}

func (p *ConstantPool) Get(index int) Object {
	return p.Values[index]
}

func (p ConstantPool) String() string {
	result := []string{}

	for i, constant := range p.Values {
		if constant == nil {
			break
		}

		result = append(result, fmt.Sprintf("%d: %+v", i, constant))
	}

	return fmt.Sprintf("[%s]", strings.Join(result, " "))
}

func (p ConstantPool) Serialize() []byte {
	result := []byte{}

	for _, obj := range p.Values[:p.pointer] {
		result = append(result, float64_to_bytes(obj.GetValue().(float64))...)
	}

	return result
}

func NewConstantPool() ConstantPool {
	return ConstantPool{
		Values:  [512]Object{},
		pointer: 0,
	}
}
