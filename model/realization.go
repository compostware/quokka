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


// The internal representation of a realization, that is a spec that has been expanded and resolved to contain a 
// complete set of concrete and consistent components that satisfy a developer's initial requirements.
type Realization struct {
	// The set of components referenced by this realization in dependency order.
	Components []*ComponentReference
	// The fragment associated with this realization.
	Fragment string
}

// A reference to a component referenced by a realization.
type ComponentReference struct {
	// The full unique identifier of the reference component.
	ComponentId string
}

