package main

import (
	"math"
)

type Stack struct {
	Values  [1024]float64
	pointer int
}

func (s *Stack) Push(value float64) {
	s.Values[s.pointer] = value
	s.pointer++
}

func (s *Stack) Pop() float64 {
	value := s.Values[s.pointer-1]
	s.pointer--
	return value
}

func NewStack() Stack {
	return Stack{
		Values: [1024]float64{},
	}
}

type Vm struct {
	Version      uint32
	ConstantPool ConstantPool
	Instructions []Instruction
	Stack        Stack
}

func (vm Vm) Run() (float64, error) {
	for _, instruction := range vm.Instructions {
		switch instruction.Op {
		case OpConstant:
			constant := vm.ConstantPool.Get(int(instruction.Operands[0]))
			vm.Stack.Push(constant.(Float64Object).Value)
		case OpAdd:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(left + right)
		case OpSub:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(left - right)
		case OpMul:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(left * right)
		case OpDiv:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(left / right)
		case OpMod:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(math.Mod(left, right))
		case OpPow:
			right := vm.Stack.Pop()
			left := vm.Stack.Pop()

			vm.Stack.Push(math.Pow(left, right))
		case OpCall:
			fn := *builtin_fns.GetPointer(int(instruction.Operands[0]))
			args := []Float64Object{}

			for i := 0; i < int(instruction.Operands[1]); i++ {
				value := vm.Stack.Pop()
				args = append(args, Float64Object{value})
			}

			ret, err := fn(args...)
			if err != nil {
				return 0, err
			}
			vm.Stack.Push(ret.Value)
		}
	}

	return vm.Stack.Pop(), nil
}

func NewVm(input []byte) (*Vm, error) {
	deserializer := NewDeserializer(input)
	deserialized, err := deserializer.Deserialize()
	if err != nil {
		return nil, err
	}

	return &Vm{
		Version:      deserialized.Version,
		ConstantPool: deserialized.ConstantPool,
		Instructions: deserialized.Instructions,
		Stack:        NewStack(),
	}, nil
}

func NewVmFromCompiler(c *Compiler) *Vm {
	return &Vm{
		Version:      c.Version,
		ConstantPool: c.ConstantPool,
		Instructions: c.Instructions,
		Stack:        NewStack(),
	}
}
