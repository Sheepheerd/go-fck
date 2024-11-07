package main

import (
	"fmt"

	"github.com/Sheepheerd/go-fck/engine"
)

func main() {
	e := engine.New()

	fmt.Println("I was able to make an engine", e)
}
