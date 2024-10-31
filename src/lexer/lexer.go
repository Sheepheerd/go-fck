package main

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
