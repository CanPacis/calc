package main

import "fmt"

type Compiler struct {
	ConstantPool ConstantPool
	Expr         Expr
	Instructions []Instruction
	Version      uint32
}

func (c *Compiler) Compile() ([]byte, error) {
	err := c.compile_expr(c.Expr)
	return c.serialize(), err
}

func (c Compiler) serialize() []byte {
	result := []byte{}
	result = append(result, []byte("calc.arc")...)
	result = append(result, uint32_to_bytes(c.Version)...)

	constants := c.ConstantPool.Serialize()
	result = append(result, uint32_to_bytes(uint32(len(constants)))...)
	result = append(result, constants...)

	for _, instruction := range c.Instructions {
		serialized := instruction.Serialize()
		result = append(result, uint32_to_bytes(uint32(len(serialized)))...)
		result = append(result, serialized...)
	}

	return result
}

func (c *Compiler) compile_expr(e Expr) error {
	switch expr := e.(type) {
	case FloatLiteralExpr:
		return c.compile_f_literal_expr(expr)
	case ConstLiteralExpr:
		return c.compile_c_literal_expr(expr)
	case FnCallExpr:
		return c.compile_call_expr(expr)
	case BinaryExpr:
		return c.compile_binary_expr(expr)
	default:
		return fmt.Errorf("unknown expression %s", expr.String())
	}
}

func (c *Compiler) compile_f_literal_expr(expr FloatLiteralExpr) error {
	index := c.ConstantPool.Add(Float64Object(expr))
	c.Instructions = append(c.Instructions, NewInstruction(OpConstant, index))
	return nil
}

func (c *Compiler) compile_c_literal_expr(expr ConstLiteralExpr) error {
	builtin, ok := builtin_consts[expr.Name]

	if !ok {
		return fmt.Errorf("constant '%s' does not exist", expr.Name)
	}

	index := c.ConstantPool.Add(builtin)
	c.Instructions = append(c.Instructions, NewInstruction(OpConstant, index))
	return nil
}

func (c *Compiler) compile_call_expr(expr FnCallExpr) error {
	builtin, ok := builtin_fns[expr.Name]

	if !ok {
		return fmt.Errorf("function '%s' does not exist", expr.Name)
	}

	for _, arg := range expr.Args {
		if err := c.compile_expr(arg); err != nil {
			return err
		}
	}

	c.Instructions = append(c.Instructions, NewInstruction(OpCall, builtin.Pointer, len(expr.Args)))

	return nil
}

func (c *Compiler) compile_binary_expr(expr BinaryExpr) error {
	if err := c.compile_expr(expr.Left); err != nil {
		return err
	}

	if err := c.compile_expr(expr.Right); err != nil {
		return err
	}

	switch expr.Op {
	case OpTypeAdd:
		c.Instructions = append(c.Instructions, NewInstruction(OpAdd))
	case OpTypeSub:
		c.Instructions = append(c.Instructions, NewInstruction(OpSub))
	case OpTypeMul:
		c.Instructions = append(c.Instructions, NewInstruction(OpMul))
	case OpTypeDiv:
		c.Instructions = append(c.Instructions, NewInstruction(OpDiv))
	case OpTypeMod:
		c.Instructions = append(c.Instructions, NewInstruction(OpMod))
	case OpTypePow:
		c.Instructions = append(c.Instructions, NewInstruction(OpPow))
	}

	return nil
}

func NewCompiler(expr Expr) *Compiler {
	return &Compiler{
		ConstantPool: NewConstantPool(),
		Instructions: []Instruction{},
		Expr:         expr,
		Version:      1,
	}
}
