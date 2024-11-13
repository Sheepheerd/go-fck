package main

import (
	"fmt"
	"os"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/linker"
	"github.com/Sheepheerd/go-fck/parser"
)

func main() {
	// all bfc needs to do, is take a -o, lex, parse, and save output

	outFileIndex := len(os.Args) - 1
	dashOIndex := len(os.Args) - 2

	if os.Args[dashOIndex] != "-o" {
		fmt.Println("Bad output flag")
		return
	}

	outputFileName := os.Args[outFileIndex]

	fmt.Printf("outputFileName: %v\n", outputFileName)

	inputFileCount := dashOIndex - 1
	if inputFileCount < 1 {
		fmt.Println("No input files provided")
		return
	}

	files := make([]*os.File, inputFileCount)

	for i := 0; i < inputFileCount; i++ {
		file, err := os.Open(os.Args[i+1])
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()
		files[i] = file
	}

	fmt.Printf("all input files: %v\n", files)

	// Link files
	linkedFile, err := linker.Link(files)
	if err != nil {
		fmt.Println("Link error")
		os.Exit(0)
	}
	defer linker.CleanObjectFiles() // todo make this configurable

	// should be moved to lexer
	operators, err := lexer.LexOperators(linkedFile)
	if err != nil {
		fmt.Println(err.Error())
		return
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
