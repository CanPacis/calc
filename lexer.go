package main

import (
	"fmt"
	"unicode"
)

type token_type int

const (
	EOFToken token_type = iota
	PlusToken
	MinusToken
	ForwardSlashToken
	StarToken
	PercentToken
	CaretToken
	OpenParensToken
	CloseParensToken
	CommaToken

	NumberToken
	IdentifierToken
	IllegalToken
)

type Location struct {
	Line int
	Col  int
}

type Token struct {
	TokenType token_type
	Literal   string
	Location  Location
}

var token_map = map[token_type]string{
	NumberToken:       "Number",
	IllegalToken:      "Illegal",
	IdentifierToken:   "Identifier",
	EOFToken:          "EOF",
	CommaToken:        ",",
	PlusToken:         "+",
	MinusToken:        "-",
	ForwardSlashToken: "/",
	StarToken:         "*",
	PercentToken:      "%",
	CaretToken:        "^",
	OpenParensToken:   "(",
	CloseParensToken:  ")",
}

func (t Token) String() string {
	if t.TokenType > CommaToken {
		return fmt.Sprintf("%s(%s)", token_map[t.TokenType], t.Literal)
	}

	return token_map[t.TokenType]
}

type Lexer struct {
	runes         []rune
	offset        int
	current_token Token
	location      Location
}

func NewLexer(input []byte) *Lexer {
	return &Lexer{
		runes: []rune(string(input)),
		location: Location{
			Line: 1,
			Col:  0,
		},
	}
}

const eof_rune rune = -1

func (l Lexer) current_rune() rune {
	if l.offset >= len(l.runes) {
		return eof_rune
	}
	return l.runes[l.offset]
}

func (l Lexer) next_rune() rune {
	if l.offset+1 >= len(l.runes) {
		return eof_rune
	}
	return l.runes[l.offset+1]
}

func (l *Lexer) advance() {
	l.offset++
	l.location.Col++
}

func (l Lexer) create_token(typ token_type, literal string) Token {
	return Token{
		TokenType: typ,
		Literal:   literal,
		Location:  l.location,
	}
}

func (l *Lexer) Next() Token {
	current_rune := l.current_rune()

	switch current_rune {
	case '+':
		l.advance()
		l.current_token = l.create_token(PlusToken, "+")
		return l.current_token
	case '-':
		if unicode.IsNumber(l.next_rune()) {
			return l.lex_number()
		} else {
			l.advance()
			l.current_token = l.create_token(MinusToken, "-")
			return l.current_token
		}
	case '/':
		l.advance()
		l.current_token = l.create_token(ForwardSlashToken, "/")
		return l.current_token
	case '*':
		l.advance()
		l.current_token = l.create_token(StarToken, "*")
		return l.current_token
	case '%':
		l.advance()
		l.current_token = l.create_token(PercentToken, "%")
		return l.current_token
	case '^':
		l.advance()
		l.current_token = l.create_token(CaretToken, "^")
		return l.current_token
	case '(':
		l.advance()
		l.current_token = l.create_token(OpenParensToken, "(")
		return l.current_token
	case ')':
		l.advance()
		l.current_token = l.create_token(CloseParensToken, "(")
		return l.current_token
	case ',':
		l.advance()
		l.current_token = l.create_token(CommaToken, ",")
		return l.current_token
	case eof_rune, 0:
		l.advance()
		l.current_token = l.create_token(EOFToken, "")
		return l.current_token
	default:
		if unicode.IsSpace(current_rune) {
			l.skip_whitespace()
			return l.Next()
		} else if unicode.IsNumber(current_rune) {
			return l.lex_number()
		} else {
			return l.lex_identifier()
		}
	}
}

func (l Lexer) CurrentToken() Token {
	return l.current_token
}

func (l *Lexer) Prev() {
	size := len(l.current_token.Literal)
	l.offset -= size
	l.location.Col -= size
}

func (l *Lexer) skip_whitespace() {
	for unicode.IsSpace(l.current_rune()) {
		if l.current_rune() == '\r' || l.current_rune() == '\n' {
			l.location.Line++
			l.location.Col = 0
		}
		l.advance()
	}
}

func (l *Lexer) lex_number() Token {
	value := []rune{}

	if l.current_rune() == '-' {
		l.advance()
		value = append(value, '-')
	}

	for unicode.IsNumber(l.current_rune()) {
		value = append(value, l.current_rune())
		l.advance()
	}

	if l.current_rune() == '.' {
		l.advance()
		value = append(value, '.')
	}

	for unicode.IsNumber(l.current_rune()) {
		value = append(value, l.current_rune())
		l.advance()
	}

	l.current_token = l.create_token(NumberToken, string(value))
	return l.current_token
}

func (l *Lexer) lex_identifier() Token {
	current := l.current_rune()
	if !identifier_major(current) {
		l.advance()
		l.current_token = Illegal
		return l.current_token
	}

	value := []rune{}

	for identifier_minor(current) {
		value = append(value, current)
		l.advance()
		current = l.current_rune()
	}

	l.current_token = l.create_token(IdentifierToken, string(value))
	return l.current_token
}

func identifier_major(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func identifier_minor(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

var Illegal = Token{
	TokenType: IllegalToken,
	Literal:   "",
}
