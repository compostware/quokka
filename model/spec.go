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

