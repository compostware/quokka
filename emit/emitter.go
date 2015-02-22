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

package emit

import (
	"io"
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the emission of a Realization as a Dockerfile.
type Emitter interface {
	// Emits the given realization to the given output writer.
	Emit(realization *model.Realization, output io.Writer)
}

// Constructor method for obtaining a reference to a Emitter.
func NewEmitter() Emitter {
	e := new(emitter)
	return e
}

type emitter struct {
	
}

func (*emitter) Emit(realization *model.Realization, output io.Writer) {
	// nufin
}

