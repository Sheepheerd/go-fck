package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/parser"
)

func main() {
	// all bfc needs to do, is take a -o, lex, parse, and save output

	// bf in.bf -o out
	// this is a bad way to do this, but works for now
	// if os.Args[2] != "-o" {
	// 	fmt.Println("Only supported flag is -o")
	// 	return
	// }

	// if len(os.Args) != 2 {
	// 	fmt.Println("Bad args")
	// 	return
	// }

	inputFileName := os.Args[1]
	fmt.Printf("inputFileName: %v\n", inputFileName)
	outputFileName := os.Args[3]

	fmt.Printf("outputFileName: %v\n", outputFileName)

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// should be moved to lexer
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

	tokens := lexer.Tokenize(operators)
	// fmt.Println("Tokens:", tokens)

	parsedTokens, symbolTable, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println("Problem parsing tokens")
	}

	outFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Could not create output file")
	}

	outFile.WriteString("symbolTable=")
	outFile.WriteString(symbolTable.Serialize())

	outFile.WriteString("\n")
	outFile.WriteString(parsedTokens.Serialize())
}
