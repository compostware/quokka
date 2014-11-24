package parse

import (
	"io"
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the parsing raw input into a Spec.
type Parser interface {
	// Reads and parses the given input, generating a Spec.
	Parse(input *io.Reader) *model.Spec
}

// Constructor method for obtaining a reference to a Parser.
func NewParser() Parser {
	p := new(parser)
	return p
}

type parser struct {
	
}

func (*parser) Parse(input *io.Reader) *model.Spec {
	s := new(model.Spec)
	return s
}

