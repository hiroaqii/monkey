package ast

import (
	"testing"

	"github.com/hiroaqii/monkey/token"
)

func TestString(t *testing.T) {
	Program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if Program.String() != "let myVar = anotherVar;" {
		t.Errorf("Program.String() wrong. got=%q", Program.String())
	}
}
