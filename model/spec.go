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
	defaultReqCapacity = 10
)


// The internal representation of a Spec, that is, an abstract, declarative description of a what a developer
// requires for their Docker image. 
type Spec struct {
	requires []Requirement
	fragment string
}

// Creates new instance
func NewSpec() (spec *Spec) {
	spec = new(Spec);
	spec.requires = make([]Requirement, 0, defaultReqCapacity)
	return;
}

// Access the requirements of this spec.
func (spec *Spec) Requires() []Requirement  {
	return spec.requires
}

// Access the syntax fragment of this spec.
func (spec *Spec) Fragment() string  {
	return spec.fragment
}

// Adds a new requirement to this spec.
func (spec *Spec) AddRequired(req Requirement) {
	spec.requires = append(spec.requires, req) 
}

// Sets the syntax fragment of this spec.
func (spec *Spec) SetFragment(fragment string) {
	spec.fragment = fragment 
}

