/*
	Copyright 2015 Lachlan Coote
	
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
	
		http://www.apache.org/licenses/LICENSE-2.0
		
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/


package realize

import (
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the transformation of an abstract spec into a concreate realization.
type Realizer interface {
	// Realizes a given Spec, generating a Realization. 
	Realize(spec *model.Spec) (*model.Realization, error)
}

// Constructor method for obtaining a reference to a Realizer.
func NewRealizer() Realizer {
	r := new(realizer)
	return r
}

type realizer struct {
	
}

func (* realizer) Realize(spec *model.Spec) (*model.Realization, error) {
	r := new(model.Realization)
	return r, nil
}

