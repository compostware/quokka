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
	"github.com/compostware/quokka/library"
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the emission of a Realization as a Dockerfile.
type Emitter interface {
	// Emits the given realization to the given output writer.
	Emit(realization *model.Realization, output io.Writer) error
}

type emitter struct {
	library library.LibraryClient
}

type emittingWriter struct {
	output io.Writer
	err error
}

// Constructor method for obtaining a reference to a Emitter.
func NewEmitter(library library.LibraryClient) Emitter {
	e := new(emitter)
	e.library = library
	return e
}

func (e *emitter) Emit(realization *model.Realization, output io.Writer) (err error) {
	err = e.emitComponents(realization.Components, output)
	if err != nil {
		return
	}
	
	err = e.emitFragment(realization.Fragment, output)
	return
}

func (e *emitter) emitComponents(refs []*model.ComponentReference, output io.Writer) error {
	if refs == nil {
		return nil
	}
	
	for _, ref := range refs {
		err := e.emitComponent(ref, output)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (e *emitter) emitComponent(ref *model.ComponentReference, output io.Writer) error {
	compId := ref.ComponentId
	
	comp, err := e.library.Retrieve(compId)
	if err != nil {
		return err
	}

	return e.emitFragment(comp.Fragment, output)
}

func (e *emitter) emitFragment(fragment string, output io.Writer) error {
	if len(fragment) == 0 {
		return nil
	}
	
	_, err := io.WriteString(output, fragment)
	return err
}



