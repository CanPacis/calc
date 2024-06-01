package main

import (
	"fmt"
	"strings"
)

type Op byte

const (
	OpNoop Op = iota + 10
	OpConstant
	OpCall
	OpPop
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpPow
	OpExit
)

var op_map = map[Op]string{
	OpNoop:     "Noop",
	OpConstant: "Constant",
	OpCall:     "Call",
	OpPop:      "Pop",
	OpAdd:      "Add",
	OpSub:      "Sub",
	OpMul:      "Mul",
	OpDiv:      "Div",
	OpMod:      "Mod",
	OpPow:      "Pow",
	OpExit:     "Exit",
}

func (op Op) String() string {
	if len(op_map[op]) > 1 {
		return op_map[op]
	}

	return fmt.Sprintf("%d", op)
}

type Instruction struct {
	Op       Op
	Operands []uint32
}

func (i Instruction) String() string {
	operands := []string{}

	for _, operand := range i.Operands {
		operands = append(operands, fmt.Sprintf("%d", operand))
	}

	return fmt.Sprintf("%s(%s)", op_map[i.Op], strings.Join(operands, " "))
}

func (i Instruction) Serialize() []byte {
	result := []byte{byte(i.Op)}

	for _, operand := range i.Operands {
		result = append(result, uint32_to_bytes(operand)...)
	}

	return result
}

func NewInstruction(op Op, operands ...int) Instruction {
	u_operands := []uint32{}

	for _, operand := range operands {
		u_operands = append(u_operands, uint32(operand))
	}

	return Instruction{
		Op:       op,
		Operands: u_operands,
	}
}
