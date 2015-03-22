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

import (

)

const (
	defaultNumComponents = 5
)

// The internal representation of a realization, that is a spec that has been expanded and resolved to contain a 
// complete set of concrete and consistent components that satisfy a developer's initial requirements.
type Realization interface {
	// The set of components referenced by this realization in dependency order.
	Components() []ComponentReference
	// Appends the given component to this realization.
	AddComponent(component ComponentReference)
	// Sets the components for this realization.
	SetComponents(components []ComponentReference)
	// The fragment associated with this realization.
	Fragment() string
	// Sets the fragment for this realization.
	SetFragment(frgament string)
}

// A reference to a component referenced by a realization.
type ComponentReference interface {
	// The full unique identifier of the reference component.
	ComponentId() string
}

// Implementation of Realization
type realization struct {
	components []ComponentReference
	fragment string
}

// Implementation of ComponentReference
type componentReference struct {
	componentId string
}

// Constructs a new ComponentReference
func NewRealization() Realization {
	b := new(realization)
	b.components = make([]ComponentReference, 0, defaultNumComponents)
	return b
}

// Constructs a new ComponentReference
func NewComponentReference(componentId string) ComponentReference {
	c := new(componentReference)
	c.componentId = componentId
	return c
}

func (b *realization) Components() []ComponentReference {
	return b.components
}

func (b *realization) AddComponent(component ComponentReference) {
	b.components = append(b.components, component)
}

func (b *realization) SetComponents(components []ComponentReference) {
	b.components = components
}

func (b *realization) Fragment() string {
	return b.fragment
}

func (b *realization) SetFragment(fragment string) {
	b.fragment = fragment
}

func (c *componentReference) ComponentId() string {
	return c.componentId
}

