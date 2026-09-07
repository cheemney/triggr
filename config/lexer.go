package config

import "unicode"

type tokenKind int

const (
	tokEOF tokenKind = iota
	tokIdent
	tokString
	tokLBrace
	tokRBrace
)

type token struct {
	kind tokenKind
	text string
}

type lexer struct {
	input []rune
	pos   int
}

func newLexer(src string) *lexer {
	return &lexer{input: []rune(src)}
}

func (l *lexer) next() token {
	l.skipSpace()
	if l.pos >= len(l.input) {
		return token{kind: tokEOF}
	}

	switch ch := l.input[l.pos]; {
	case ch == '{':
		l.pos++
		return token{kind: tokLBrace, text: "{"}
	case ch == '}':
		l.pos++
		return token{kind: tokRBrace, text: "}"}
	case ch == '"':
		return l.readString()
	default:
		return l.readIdent()
	}
}

func (l *lexer) skipSpace() {
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos++
	}
}

func (l *lexer) readString() token {
	l.pos++ // skip opening quote
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '"' {
		l.pos++
	}
	text := string(l.input[start:l.pos])
	l.pos++ // skip closing quote
	return token{kind: tokString, text: text}
}

func (l *lexer) readIdent() token {
	start := l.pos
	for l.pos < len(l.input) &&
		!unicode.IsSpace(l.input[l.pos]) &&
		l.input[l.pos] != '{' &&
		l.input[l.pos] != '}' {
		l.pos++
	}
	return token{kind: tokIdent, text: string(l.input[start:l.pos])}
}
