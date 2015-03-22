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


package model


const (
	defaultProvidesCapacity = 5;
	defaultRequiresCapacity = 5;
)



// A reusable component that can be referenced when building an image.
type Component struct {
	// The unique identifier of this component.
	Id string
	// The services provided by this component.
	Provides []*ProvidedService
	// The requirements of this component.
	Requires []Requirement
	// The syntax fragment associated with this component.
	Fragment string
}

// A reusable component that can be referenced when building an image.
type ProvidedService struct {
	// The identifier of the provided service.
	ServiceId string
}

// Constructor method for a component
func NewComponent(id, fragment string) (c *Component) {
	c = new(Component)
	c.Id = id
	c.Provides = make([]*ProvidedService, 0, defaultProvidesCapacity)
	c.Requires = make([]Requirement, 0, defaultRequiresCapacity)
	c.Fragment = fragment
	return c
}

// Adds a new provided service to the given component.
func (c *Component) AddProvides(provided *ProvidedService) {
	c.Provides = append(c.Provides, provided)
}

// Adds a new requirement to the given component
func (c *Component) AddRequires(req Requirement) {
	c.Requires = append(c.Requires, req)
}

