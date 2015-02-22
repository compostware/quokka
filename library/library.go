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

package library

import (
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the publishing and retrieval of components from the library.
type LibraryClient interface {
	// Publishes a component to the library. This component might overwrite an existing component with the
	// same unique identifier.
	Publish(component *model.Component)
	// Retrieves a component from the library by its unique identifier.
	Retrieve(componentId string) *model.Component
}

// Constructor method for obtaining a reference to a LibraryClient.
func NewLibraryClient() LibraryClient {
	lib := new(inMemoryLibrary)
	lib.components = make(map[string]*model.Component)
	return lib
}

// A simple initial library (and client) implementation that stores components in memory in a map, keyed by
// id.
type inMemoryLibrary struct {
	components map[string]*model.Component
}

func (lib *inMemoryLibrary) Publish(component *model.Component) {
	lib.components[component.Id] = component
}

func (lib *inMemoryLibrary) Retrieve(componentId string) *model.Component {
	return lib.components[componentId]
	
}

