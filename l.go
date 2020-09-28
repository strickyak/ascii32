package ascii32

import (
	"strings"
)

type TokType int

const (
	TokEOF TokType = iota
	TokLetters
	TokOp
	TokWhite
)

type Tok struct {
	Line int
	Col  int
	S    string
	Type TokType
}

type lex struct {
	Line    int
	Col     int
	Rest    []rune
	Program []rune
}

func Tokenize(s string) []Tok {
	o := newLex(s)
	var z []Tok
	for {
		t := o.Next()
		// log.Printf("NEXT ----> %v", t)
		z = append(z, t)
		if t.Type == TokEOF {
			break
		}
	}
	return z
}

func newLex(s string) *lex {
	return &lex{
		Line:    1,
		Col:     1,
		Rest:    []rune(s),
		Program: []rune(s),
	}
}

func (o *lex) Next() Tok {
	l, c := o.Line, o.Col // current line & col

	if len(o.Rest) == 0 {
		return Tok{
			Line: l,
			Col:  c,
			S:    "",
			Type: TokEOF,
		}
	}

	var buf strings.Builder
	for len(o.Rest) > 0 && IsWhite(o.Rest[0]) {
		buf.WriteRune(o.Rest[0])
		if IsNewLine(o.Rest[0]) {
			o.Line++
			o.Col = 1
		} else {
			o.Col++
		}
		o.Rest = o.Rest[1:]
	}
	if buf.Len() > 0 {
		return Tok{
			Line: l,
			Col:  c,
			S:    buf.String(),
			Type: TokWhite,
		}
	}

	for len(o.Rest) > 0 && IsLetter(o.Rest[0]) {
		buf.WriteRune(o.Rest[0])
		o.Col++
		o.Rest = o.Rest[1:]
	}
	if buf.Len() > 0 {
		return Tok{
			Line: l,
			Col:  c,
			S:    buf.String(),
			Type: TokLetters,
		}
	}

	o.Col++
	str := string([]rune{o.Rest[0]})
	o.Rest = o.Rest[1:]
	return Tok{
		Line: l,
		Col:  c,
		S:    str,
		Type: TokOp,
	}
}

func IsNewLine(r rune) bool {
	return r == '\n'
}
func IsWhite(r rune) bool {
	return r <= ' '
}
func IsLetter(r rune) bool {
	switch {
	case 'A' <= r && r <= 'Z':
		return true
	case 'a' <= r && r <= 'z':
		return true
	case '0' <= r && r <= '9':
		return true
	}
	switch r {
	case '_', '.', ',':
		return true
	}
	return false
}
