package realize

import (
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the transformation of an abstract spec into a concreate realization.
type Realizer interface {
	// Realizes a given Spec, generating a Realization. 
	Realize(spec *model.Spec) *model.Realization
}

// Constructor method for obtaining a reference to a Realizer.
func NewRealizer() Realizer {
	r := new (realizer)
	return r
}

type realizer struct {
	
}

func (* realizer) Realize(spec *model.Spec) *model.Realization {
	r := new(model.Realization)
	return r
}

