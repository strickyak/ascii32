package ascii32

import (
	"fmt"
	"log"
	"strconv"
)

type (
	A__ struct {
		Emit func(r rune)
		Ops  [128]func(*A__)
		Data Stack
	}

	Stack struct {
		v []string
	}
)

func (o *Stack) Push(x string) {
	o.v = append(o.v, x)
}

func (o *Stack) Pop() string {
	z := o.v[len(o.v)-1]
	o.v = o.v[:len(o.v)-1]
	return z
}

func (o *Stack) String() string {
	return fmt.Sprintf("%#v", o.v)
}

func New__(emit func(r rune)) *A__ {
	z := &A__{
		Emit: emit,
	}
	z.Ops['+'] = func(o *A__) {
		b := o.Data.Pop()
		a := o.Data.Pop()
		x, xerr := strconv.ParseFloat(a, 64)
		if xerr != nil {
			log.Panicf("cannot ParseFloat %q: %v", a, xerr)
		}
		y, yerr := strconv.ParseFloat(b, 64)
		if yerr != nil {
			log.Panicf("cannot ParseFloat %q: %v", b, yerr)
		}
		o.Data.Push(fmt.Sprintf("%g", x+y))
	}
	z.Ops['!'] = func(o *A__) {
		s := o.Data.Pop()
		for _, r := range s {
			o.Emit('<')
			o.Emit(r)
			o.Emit('>')
		}
	}
	return z
}

func (o *A__) RunProgram(s string) {
	tt := Tokenize(s)

	// Debug tokenizer:
	for i, t := range tt {
		log.Printf("Tok[%d] = %v", i, t)
	}

	for i, t := range tt {
		log.Printf("Step[%d]: %v", i, t)
		o.Step(t)
	}

	log.Printf("Final Stack: %v", o.Data)
}

func (o *A__) Step(t Tok) {
	switch t.Type {
	case TokEOF, TokWhite:
		return
	case TokLetters:
		o.Data.Push(t.S)
	case TokOp:
		fn := o.Ops[t.S[0]]
		if fn == nil {
			log.Panicf("op not defined: <%c>", t.S[0])
		}
		fn(o)

	}
}
