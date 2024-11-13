package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Token int

var tokens []Token

const (
	LessThan Token = iota
	GreaterThan
	Plus
	Minus
	LeftBracket
	RightBracket
	Comma
	Period
)

func Tokenize(tokensSlice []rune) []Token {
	var tokens []Token

	for _, token := range tokensSlice {
		switch token {
		case '<':
			tokens = append(tokens, LessThan)
		case '>':
			tokens = append(tokens, GreaterThan)
		case '+':
			tokens = append(tokens, Plus)
		case '-':
			tokens = append(tokens, Minus)
		case '[':
			tokens = append(tokens, LeftBracket)
		case ']':
			tokens = append(tokens, RightBracket)
		case ',':
			tokens = append(tokens, Comma)
		case '.':
			tokens = append(tokens, Period)
		}
	}
	return tokens
}

var (
	ErrCouldNotLex = errors.New("could not lex operators")
)

func LexOperators(f *os.File) ([]rune, error) {
	scanner := bufio.NewScanner(f)
	var operators []rune

	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char != ' ' {
				operators = append(operators, char)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, ErrCouldNotLex
	}

	return operators, nil
}
