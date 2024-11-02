package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Sheepheerd/go-fck/engine"
	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/parser"
)

func main() {
	fmt.Println("go fck yourself")

	// Get filename from cli arg
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-fck <filename>")
		return
	}

	fileName := os.Args[1]

	// Open bf file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
		return
	}

	// Create lexer and call tokenize with the bf string of code
	tokens := lexer.Tokenize(operators)

	fmt.Println("Tokens:", tokens)
	// Get back a slice of tokens

	parsedTokens, symbolTable, err := parser.Parse(tokens)

	fmt.Printf("symbolTable: %v\n", symbolTable)
	if err != nil {
		fmt.Println("Problem parsing tokens")
	}

	fmt.Printf("parsedTokens: %v\n", parsedTokens)

	// Pass parsed tokens into engine

	engine.New().RunInstructions(parsedTokens, symbolTable)
}
