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


// A reusable component that can be referenced when building an image.
type Component struct {
	// The unique identifier of this component.
	Id string
	// The syntax fragment associated with this component.
	Fragment string
}

// Constructor method for a component
func NewComponent(id, fragment string) (c *Component) {
	c = new(Component)
	c.Id = id
	c.Fragment = fragment
	return c
}

