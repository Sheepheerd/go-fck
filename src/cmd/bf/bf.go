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

	var engineSymbols []lexer.Token
	var symbolTable map[int]int

	isBin := isBinFile(file)

	// reset scanner from previous bin file check
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error resetting file cursor:", err)
	}

	if isBin {
		fmt.Println("Got bin file")

		scanner := bufio.NewScanner(file)

		var stData string
		if scanner.Scan() {
			stData = scanner.Text()
		}

		symbolTable = parser.DeserializeSymbolTable(stData)

		var instrData string
		if scanner.Scan() {
			instrData = scanner.Text()
		}
		operators := []rune(instrData)
		engineSymbols = lexer.Tokenize(operators)

	} else {
		fmt.Println("Got regular file")

		operators, err := lexer.LexOperators(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tokens := lexer.Tokenize(operators)

		engineSymbols, symbolTable, err = parser.Parse(tokens)
		if err != nil {
			fmt.Println("Problem parsing tokens")
		}

	}

	engine.New().RunInstructions(engineSymbols, symbolTable)
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
