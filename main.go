package main

import (
	"fmt"

	"github.com/swayam5342/moca/Lexer"
)

func main() {
	lex := Lexer.New("{}===+")
	for t := lex.NextToken(); t.Type != Lexer.EOF; t = lex.NextToken() {
		fmt.Printf("%+v\n", t)
	}
}
