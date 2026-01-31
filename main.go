package main

import "fmt"

type Token struct {
	Type    string
	Literal string
}
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func main() {
	input := "=+(){},;"
	lexer := New(input)
	for tok := lexer.NextToken(); tok.Type != "EOF"; tok = lexer.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{Type: "ASSIGN", Literal: string(l.ch)}
	case '+':
		tok = Token{Type: "PLUS", Literal: string(l.ch)}
	case '(':
		tok = Token{Type: "LPAREN", Literal: string(l.ch)}
	case ')':
		tok = Token{Type: "RPAREN", Literal: string(l.ch)}
	case '{':
		tok = Token{Type: "LBRACE", Literal: string(l.ch)}
	case '}':
		tok = Token{Type: "RBRACE", Literal: string(l.ch)}
	case ',':
		tok = Token{Type: "COMMA", Literal: string(l.ch)}
	case ';':
		tok = Token{Type: "SEMICOLON", Literal: string(l.ch)}
	case 0:
		tok = Token{Type: "EOF", Literal: ""}
	default:
		tok = Token{Type: "ILLEGAL", Literal: string(l.ch)}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
