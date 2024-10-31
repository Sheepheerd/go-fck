package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sheepheerd/go-fck/engine"
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

	reader := bufio.NewReader(file)

	word, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from the file:", err)
		return
	}
	word = strings.TrimSpace(word)
	fmt.Printf("The first word in the file is: '%s'\n", word)

	// Read in input from file and remove white space

	// Create lexer and call tokenize with the bf string of code

	// Get back a slice of tokens

	// Pass tokens to parser

	// Generate AST from parser
	engine.RunEngine()
}
