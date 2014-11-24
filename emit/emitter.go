package emit

import (
	"io"
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the emission of a Realization as a Dockerfile.
type Emitter interface {
	// Emits the given realization to the given output writer.
	Emit(realization *model.Realization, output *io.Writer)
}

// Constructor method for obtaining a reference to a Emitter.
func NewEmitter() Emitter {
	e := new(emitter)
	return e
}

type emitter struct {
	
}

func (*emitter) Emit(realization *model.Realization, output *io.Writer) {
	// nufin
}

