package main

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenNumber   TokenType = "NUMBER"
	TokenOperator TokenType = "OPERATOR"
	TokenEOF      TokenType = "EOF"
	TokenInvalid  TokenType = "INVALID"
	TokenReturn   TokenType = "RETURN"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) nextToken() Token {
	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	ch := l.input[l.pos]

	// Ignorar espacios
	if unicode.IsSpace(rune(ch)) {
		l.pos++
		return l.nextToken()
	}

	// Leer digitos hasta que se acaben o lleguemos all final del archivo
	if unicode.IsDigit(rune(ch)) {
		start := l.pos
		for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
			l.pos++
		}

		return Token{Type: TokenNumber, Value: l.input[start:l.pos]}
	}

	if ch == '+' || ch == '-' || ch == '*' || ch == '/' { // Esto se puede hacer mejor seguro
		l.pos++
		return Token{Type: TokenOperator, Value: string(ch)}
	}

	if unicode.IsLetter(rune(ch)) {
		word := ""
		for l.pos < len(l.input) && unicode.IsLetter(rune(l.input[l.pos])) {
			word += string(l.input[l.pos])
			l.pos++
		}

		word = strings.ToLower(word)
		if word == "return" {
			l.pos++
			return Token{Type: TokenReturn, Value: "0"} // Esta parte va a haber que cambiarla para manejar todo lo que sean keywords
		} else {
			l.pos++
			return Token{Type: TokenInvalid, Value: word}
		}
	}

	l.pos++
	return Token{Type: TokenInvalid, Value: string(ch)}
}
