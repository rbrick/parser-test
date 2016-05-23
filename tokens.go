package main

import (
	"strings"
)

type (
	TokenType int
	Token     struct {
		tokenType TokenType
		value     string
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

	// Reprents a String
	TokenString

	// Represents an integer
	TokenInteger

	// Represents a float
	TokenFloat

	// Anything else that is not a keyword
	TokenAny
)

// Turns an input string, into a set of tokens that are easy to work with
type Tokenizer struct {
	// The current position of the tokenizer
	Pos         int
	// The previous token we parsed
	Prev        Token
	// The token at the current position
	At          Token
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
		s = strings.Replace(s, "\r", "", -1)
	}
	return &Tokenizer{
		Pos:         0,
		Tokens:      []Token{},
		Input:       s,
		InputLength: len(s),
	}
}

func (t *Tokenizer) ParseToken() {
	// Get the char at the current position
	charAt := string(t.Input[t.Pos])

	var token Token

	switch rune(charAt[0]) {
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

	t.Tokens = append(t.Tokens, token)

	// Swap
	t.Prev = t.At
	t.At = token

	// Increment the position
	t.Pos++
}

func (t *Tokenizer) ParseString() {
	//token := Token{tokenType: TokenString, value: ""}
	t.ParseToken()

	for t.At.value[0] != '"' {
		// Ignore character escapes
		if t.At.value[0] == '\\'  {
			t.ParseToken()
		}
	}
}

func (t *Tokenizer) Tokenize() {
	for t.Pos < t.InputLength {
		t.ParseToken()

	}
}
