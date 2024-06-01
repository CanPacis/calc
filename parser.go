package main

import (
	"fmt"
	"path"
	"slices"
	"strconv"
	"strings"
)

type Parser struct {
	filepath string
	filename string
	input    []byte
	lexer    *Lexer
}

func NewParser(input []byte, filepath string) Parser {
	return Parser{
		filepath: filepath,
		filename: path.Base(filepath),
		input:    input,
	}
}

func (p Parser) expect(expected []token_type) (Token, error) {
	token := p.lexer.Next()

	expected_str := []string{}
	for _, e := range expected {
		expected_str = append(expected_str, token_map[e])
	}

	if !slices.Contains(expected, token.TokenType) {
		if len(expected_str) == 1 {
			return Illegal, fmt.Errorf("invalid '%s' token, was expecting a '%s' token at %d:%d", token, expected_str[0], token.Location.Line, token.Location.Col)
		} else {
			return Illegal, fmt.Errorf("invalid '%s' token, was expecting any of [%s] at %d:%d", token, strings.Join(expected_str, ", "), token.Location.Line, token.Location.Col)
		}
	}

	return token, nil
}

func (p *Parser) Parse() (Expr, error) {
	p.lexer = NewLexer(p.input)
	return p.parse_expr(false)
}

/*
"Arithmetic Expressions" {
expression = term  { ("+" | "-") term} .
term       = factor  { ("*"|"/") factor} .
factor     = constant | variable | "("  expression  ")"  | fn.
fn = variable "(" arg_list ")"
arg_list = expression | expression "," arg_list
variable   = "x" | "y" | "z" .
constant   = digit {digit} .
digit      = "0" | "1" | "..." | "9" .
}
*/

func (p *Parser) parse_expr(inline bool) (Expr, error) {
	term, err := p.parse_term()
	if err != nil {
		return nil, err
	}

	parse_right := func(op_type OpType) (BinaryExpr, error) {
		right, err := p.parse_term()
		if err != nil {
			return BinaryExpr{}, err
		}

		if !inline {
			_, err = p.expect([]token_type{PlusToken, MinusToken, StarToken, ForwardSlashToken, PercentToken, CaretToken, EOFToken})
		}

		return binary_expr(term, right, op_type), err
	}

	next := p.lexer.Next()

	switch next.TokenType {
	case EOFToken:
		return term, nil
	case PlusToken:
		return parse_right(OpTypeAdd)
	case MinusToken:
		return parse_right(OpTypeSub)
	default:
		p.lexer.Prev()

		if !inline {
			_, err = p.expect([]token_type{PlusToken, MinusToken, StarToken, ForwardSlashToken, PercentToken, CaretToken})
			return nil, err
		}

		return term, nil
	}
}

func (p *Parser) parse_term() (Expr, error) {
	factor, err := p.parse_factor()
	if err != nil {
		return nil, err
	}

	parse_right := func(op_type OpType) (BinaryExpr, error) {
		right, err := p.parse_factor()
		if err != nil {
			return BinaryExpr{}, err
		}

		return binary_expr(factor, right, op_type), nil
	}

	next := p.lexer.Next()

	switch next.TokenType {
	case EOFToken:
		return factor, nil
	case StarToken:
		return parse_right(OpTypeMul)
	case ForwardSlashToken:
		return parse_right(OpTypeDiv)
	case PercentToken:
		return parse_right(OpTypeMod)
	case CaretToken:
		return parse_right(OpTypePow)
	default:
		p.lexer.Prev()
		return factor, nil
	}
}

func (p *Parser) parse_factor() (Expr, error) {
	token, err := p.expect([]token_type{NumberToken, IdentifierToken, OpenParensToken})
	if err != nil {
		return nil, err
	}

	switch token.TokenType {
	case NumberToken:
		value, _ := strconv.ParseFloat(token.Literal, 64)
		return f_literal(value), nil
	case IdentifierToken:
		next := p.lexer.Next()

		if next.TokenType == OpenParensToken {
			return p.parse_call_expr(token.Literal)
		} else {
			p.lexer.Prev()
			return c_literal(token.Literal), nil
		}
	case OpenParensToken:
		expr, err := p.parse_expr(true)
		if err != nil {
			return nil, err
		}
		_, err = p.expect([]token_type{CloseParensToken})
		if err != nil {
			return nil, err
		}

		return group_expr(expr), nil
	}

	return nil, nil
}

func (p *Parser) parse_call_expr(name string) (FnCallExpr, error) {
	next := p.lexer.Next()
	if next.TokenType == CloseParensToken {
		return fn_call(name), nil
	}

	p.lexer.Prev()
	args, err := p.parse_arg_list()
	if err != nil {
		return FnCallExpr{}, err
	}

	_, err = p.expect([]token_type{CloseParensToken})
	if err != nil {
		return FnCallExpr{}, err
	}

	return fn_call(name, args...), nil
}

func (p *Parser) parse_arg_list() ([]Expr, error) {
	left, err := p.parse_expr(true)
	if err != nil {
		return nil, err
	}

	result := []Expr{left}
	next := p.lexer.Next()

	if next.TokenType == CommaToken {
		right, err := p.parse_arg_list()
		if err != nil {
			return nil, err
		}
		result = append(result, right...)
		return result, nil
	} else {
		p.lexer.Prev()
		return result, nil
	}
}
