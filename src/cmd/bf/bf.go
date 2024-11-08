package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	if isBinFile(file) {
		fmt.Println("Got bin file")

	} else {
		fmt.Println("Got regular file")

		// reset scanner from bin file check

		operators, err := lexer.LexOperators(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tokens := lexer.Tokenize(operators)

		parsedTokens, symbolTable, err := parser.Parse(tokens)
		if err != nil {
			fmt.Println("Problem parsing tokens")
		}

		engine.New().RunInstructions(parsedTokens, symbolTable)
	}

}

func isBinFile(f *os.File) bool {
	scanner := bufio.NewScanner(f)

	var stLine string
	if scanner.Scan() {
		stLine = scanner.Text()
	}
	fmt.Printf("stLine: %v\n", stLine)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return strings.HasPrefix(stLine, "symbolTable=")
}
