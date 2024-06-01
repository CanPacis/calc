package main

import (
	"fmt"
	"strings"
)

type Expr interface {
	expr()
	String() string
}

type OpType int

const (
	OpTypeAdd OpType = iota
	OpTypeSub
	OpTypeMul
	OpTypeDiv
	OpTypeMod
	OpTypePow
)

var op_type_map = map[OpType]string{
	OpTypeAdd: "+",
	OpTypeSub: "-",
	OpTypeMul: "*",
	OpTypeDiv: "/",
	OpTypeMod: "%",
	OpTypePow: "^",
}

type FloatLiteralExpr struct {
	Value float64
}

func (expr FloatLiteralExpr) String() string {
	return fmt.Sprintf("%f", expr.Value)
}

func f_literal(value float64) FloatLiteralExpr {
	return FloatLiteralExpr{Value: value}
}

type ConstLiteralExpr struct {
	Name string
}

func (expr ConstLiteralExpr) String() string {
	return expr.Name
}

func c_literal(name string) ConstLiteralExpr {
	return ConstLiteralExpr{Name: name}
}

type BinaryExpr struct {
	Left  Expr
	Right Expr
	Op    OpType
}

func (expr BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", expr.Left.String(), op_type_map[expr.Op], expr.Right.String())
}

func binary_expr(left, right Expr, op OpType) BinaryExpr {
	return BinaryExpr{
		Left:  left,
		Right: right,
		Op:    op,
	}
}

type FnCallExpr struct {
	Name string
	Args []Expr
}

func (expr FnCallExpr) String() string {
	args := []string{}

	for _, arg := range expr.Args {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%s(%s)", expr.Name, strings.Join(args, " "))
}

func fn_call(name string, args ...Expr) FnCallExpr {
	return FnCallExpr{
		Name: name,
		Args: args,
	}
}

type GroupExpr struct {
	Expr Expr
}

func (expr GroupExpr) String() string {
	return fmt.Sprintf("(%s)", expr.Expr.String())
}

func group_expr(expr Expr) GroupExpr {
	return GroupExpr{
		Expr: expr,
	}
}

func (e FloatLiteralExpr) expr() {}
func (e ConstLiteralExpr) expr() {}
func (e BinaryExpr) expr()       {}
func (e FnCallExpr) expr()       {}
func (e GroupExpr) expr()        {}
