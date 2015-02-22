package model

import (

)


// The internal representation of a realization, that is a spec that has been expanded and resolved to contain a 
// complete set of concrete and consistent components that satisfy a developer's initial requirements.
type Realization struct {
	// The set of components referenced by this realization in dependency order.
	Components []ComponentReference
	// The fragment associated with this realization.
	Fragment string
}

// A reference to a component referenced by a realization.
type ComponentReference struct {
	// The full unique identifier of the reference component.
	ComponentId string
}

