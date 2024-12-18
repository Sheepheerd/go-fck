package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/stack"
)

var (
	ErrInvalidSyntax = errors.New("invalid syntax error")
)

type Tokens []lexer.Token

type SymbolTable map[int]int

func (s *SymbolTable) Serialize() string {
	var builder strings.Builder

	for k, v := range *s {
		builder.WriteString(fmt.Sprintf("%d:%d ", k, v))
	}

	return builder.String()
}

// gotta move this into a proper symbol table package eventually, should be its own type anyway
func DeserializeSymbolTable(input string) map[int]int {
	result := make(map[int]int)

	input = strings.TrimPrefix(input, "symbolTable=")

	pairs := strings.Split(input, " ")

	for _, pair := range pairs {
		values := strings.Split(pair, ":")
		if len(values) != 2 {
			continue
		}

		key, err1 := strconv.Atoi(values[0])
		value, err2 := strconv.Atoi(values[1])
		if err1 != nil || err2 != nil {
			continue
		}

		result[key] = value
	}

	return result
}

func (t *Tokens) Serialize() string {

	data := make([]rune, len(*t))

	for i, token := range *t {
		switch token {
		case lexer.LessThan:
			data[i] = '<'
		case lexer.GreaterThan:
			data[i] = '>'
		case lexer.Comma:
			data[i] = ','
		case lexer.Plus:
			data[i] = '+'
		case lexer.Minus:
			data[i] = '-'
		case lexer.Period:
			data[i] = '.'
		case lexer.LeftBracket:
			data[i] = '['
		case lexer.RightBracket:
			data[i] = ']'
		}
	}

	return string(data)
}

func Parse(toks []lexer.Token) (Tokens, SymbolTable, error) {
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
