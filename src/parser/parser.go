package parser

import (
	"errors"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/stack"
)

var (
	ErrInvalidSyntax = errors.New("invalid syntax error")
)

func Parse(toks []lexer.Token) ([]lexer.Token, map[int]int, error) {
	symbolTable, err := createSymbolTable(toks)
	if err != nil {
		return nil, nil, ErrInvalidSyntax
	}

	return toks, symbolTable, nil

}

func createSymbolTable(toks []lexer.Token) (map[int]int, error) {

	s := stack.New()
	awaitingTokenAddresses := stack.New()

	symbolTable := make(map[int]int)

	for addr, tok := range toks {

		if tok == lexer.LeftBracket {
			awaitingTokenAddresses.Push(addr)
			symbolTable[addr] = -1
			s.Push("go")
		} else if tok == lexer.RightBracket {
			var tokenAddress int = awaitingTokenAddresses.Pop().(int)
			symbolTable[tokenAddress] = addr
			symbolTable[addr] = tokenAddress
			s.Pop()
		}
	}
	if s.Size() != 0 {
		return nil, ErrInvalidSyntax
	}

	return symbolTable, nil
}
