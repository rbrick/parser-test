package main

import (
	"strings"
)

type (
	TokenType int
	Token struct {
		tokenType TokenType
		value     rune
	}
)

const (
	// The token used for comments
	TokenComment TokenType = iota

	// The token used for assigning a key to a value
	TokenAssign

	// The token used for starting an array
	TokenArrayStart

	// The token used to close an array
	TokenArrayClose

	// Used for opening, and closing strings
	TokenDoubleQuote

	// Anything else that is not a keyword
	TokenAny

	// Represents nothing!
	TokenNone
)

// Turns an input string, into a set of tokens that are easy to work with
type Tokenizer struct {
	// The current position of the tokenizer
	Pos         int
	// The tokens that it has parsed
	Tokens      []Token
	// The input
	Input       string
	// The length of the input, here for convenience
	InputLength int
}


// Creates a new tokenizer
func NewTokenizer(s string, clean bool) *Tokenizer {
	// Clean line separators
	if clean {
		s = strings.Replace(strings.Replace(s, "\r", "", -1), "\n", "", -1)
	}
	return &Tokenizer{
		Pos: 0,
		Tokens: []Token{},
		Input: s,
		InputLength: len(s),
	}
}

func (t *Tokenizer) ParseToken() {
	// Get the char at the current position
	charAt := rune(t.Input[t.Pos])

	var token Token

	switch charAt {
	case '#':
		token = Token{tokenType: TokenComment, value: charAt}
	case '=':
		token = Token{tokenType: TokenAssign, value: charAt}
	case '[':
		token = Token{tokenType: TokenArrayStart, value: charAt}
	case ']':
		token = Token{tokenType: TokenArrayClose, value: charAt}
	case '"':
		token = Token{tokenType: TokenDoubleQuote, value: charAt}
	default:
		token = Token{tokenType: TokenAny, value: charAt}
	}

	// Append the position
	t.Tokens = append(t.Tokens, token)

	// Increment the position
	t.Pos++
}

func (t *Tokenizer) Tokenize() {
	for t.Pos < t.InputLength {
		t.ParseToken()
	}
}