package engine

import (
	"bufio"
	"container/list"
	"fmt"
	"os"

	"github.com/Sheepheerd/go-fck/lexer"
	"github.com/Sheepheerd/go-fck/stack"
)

type Engine struct {
	instructionPointer int
	tapePointer        *list.Element
	tape               *list.List
	stack              *stack.Stack
}

func New() *Engine {
	e := &Engine{
		instructionPointer: 0,
		tape:               list.New(),
		stack:              stack.New(),
	}

	e.tapePointer = e.tape.PushBack(byte(0))

	return e
}

func (e *Engine) RunInstructions(parsedTokens []lexer.Token) {

	reader := bufio.NewReader(os.Stdin) // pass this in eventually

	for {
		token := parsedTokens[e.instructionPointer]
		switch token {
		case lexer.LessThan:
			e.decramentTapePointer()
			e.incramentInstructionPointer()
		case lexer.GreaterThan:
			e.incramentTapePointer()
			e.incramentInstructionPointer()
		case lexer.Plus:
			e.incramentCell()
			e.incramentInstructionPointer()
		case lexer.Minus:
			e.decramentCell()
			e.incramentInstructionPointer()
		case lexer.LeftBracket:
		case lexer.RightBracket:
		case lexer.Comma:
			e.putCellValue(*reader)
			e.incramentInstructionPointer()
		case lexer.Period:
			e.printCellValue()
			e.incramentInstructionPointer()
		}

	}

}

func (e *Engine) incramentInstructionPointer() {
	e.instructionPointer++
}

func (e *Engine) incramentTapePointer() {
	if e.tapePointer == e.tape.Back() {
		e.tapePointer = e.tape.PushBack(byte(0))
	} else {
		e.tapePointer = e.tapePointer.Next()
	}
}

func (e *Engine) decramentTapePointer() {
	if e.tapePointer != e.tape.Front() {
		e.tapePointer = e.tapePointer.Prev()
	}
}

func (e *Engine) incramentCell() {
	if val, ok := e.tapePointer.Value.(byte); ok {
		e.tapePointer.Value = val + 1
	} else {
		fmt.Println("Error: Expected byte value in tapePointer")
	}
}

func (e *Engine) decramentCell() {
	if val, ok := e.tapePointer.Value.(byte); ok {
		e.tapePointer.Value = val - 1
	} else {
		fmt.Println("Error: Expected byte value in tapePointer")
	}
}

func (e *Engine) printCellValue() {
	fmt.Println(e.tapePointer.Value)
}

func (e *Engine) putCellValue(reader bufio.Reader) {
	nextByte, err := reader.ReadByte()

	if err != nil {
		fmt.Println("Could not read data")
		return
	}

	e.tapePointer.Value = nextByte
}
