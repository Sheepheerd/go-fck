package parser

import (
	"errors"
	"fmt"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/stack"
)

var (
	ErrInvalidSyntax = errors.New("invalid syntax error")
)

func Parse(toks []lexer.Token) ([]lexer.Token, error) {

	isValid := hasValidSquareBraces(toks)

	fmt.Printf("isValid: %v\n", isValid)

	if !isValid {
		return nil, ErrInvalidSyntax
	}

	return toks, nil

}

func hasValidSquareBraces(toks []lexer.Token) bool {

	s := stack.New()

	for _, tok := range toks {

		if tok == lexer.LeftBracket {
			fmt.Println("pushed")
			s.Push("go") // this is dumb
		} else if tok == lexer.RightBracket {
			s.Pop()
			fmt.Println("popped")
		}

	}

	return s.Size() == 0

}
