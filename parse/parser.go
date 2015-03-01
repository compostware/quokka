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


package parse

import (
	"io"
	"github.com/compostware/quokka/model"
)

// An interface encapsulating the logic for the parsing raw input into a Spec.
type Parser interface {
	// Reads and parses the given input, generating a Spec.
	Parse(input io.Reader) (*model.Spec, error)
}

// Constructor method for obtaining a reference to a Parser.
func NewParser() Parser {
	p := new(parser)
	return p
}

type parser struct {
	
}

func (*parser) Parse(input io.Reader) (*model.Spec, error) {
	s := new(model.Spec)
	return s, nil
}

