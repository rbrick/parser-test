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

	// Anything else that is not a keyword
	TokenAny

	// Represents nothing!
	TokenNone
)

type Tokenizer struct {
	// The current position of the tokenizer
	Pos         int
	At          Token
	Tokens      []Token
	Input       string
	InputLength int
}


// Creates a new tokenizer
func NewTokenizer(s string, clean bool) *Tokenizer {
	if clean {
		s = strings.Replace(s, "\n", "", -1)
	}
	return &Tokenizer{
		Pos: 0,
		At: Token{tokenType: TokenNone, value: 0},
		Tokens: []Token{},
		Input: s,
		InputLength: len(s),
	}
}

func (t *Tokenizer) ParseToken() {
	// Get the char at the current position
	charAt := rune(t.Input[t.Pos])

	switch charAt {
	case '#':
		t.At = Token{tokenType: TokenComment, value: charAt}
	case '=':
		t.At = Token{tokenType: TokenAssign, value: charAt}
	case '[':
		t.At = Token{tokenType: TokenArrayStart, value: charAt}
	case ']':
		t.At = Token{tokenType: TokenArrayClose, value: charAt}
	default:
		t.At = Token{tokenType: TokenAny, value: charAt}
	}

	// Append the position
	t.Tokens = append(t.Tokens, t.At)

	// Increment the position
	t.Pos++
}

func (t *Tokenizer) Tokenize() {
	for t.Pos < t.InputLength {
		t.ParseToken()
	}
}