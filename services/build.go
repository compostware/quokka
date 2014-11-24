package services

import (
	"fmt"
	"io"
	"github.com/compostware/quokka/parse"
	"github.com/compostware/quokka/realize"
	"github.com/compostware/quokka/emit"
)


// A service interface providing mechanisms to build Dockerfiles from input specs.
type BuildService interface {
	// Builds a Dockerfile from an inputted spec and writes the results to an output file.
	Build(input *io.Reader, output *io.Writer)
}

// A constuctor method for creating a new BuildService.
func NewBuildService() BuildService {
	b := new(builder)
	b.parser = parse.NewParser()
	b.realizer = realize.NewRealizer()
	b.emitter = emit.NewEmitter()
	return b
}

type builder struct {
	parser parse.Parser
	realizer realize.Realizer
	emitter emit.Emitter
}

func (b *builder) Build(input *io.Reader, output *io.Writer) {
	fmt.Println("Building...")
	
	// parse
	spec := b.parser.Parse(input)
	
	// realize
	realization := b.realizer.Realize(spec)
	
	// emit
	b.emitter.Emit(realization, output)
}

