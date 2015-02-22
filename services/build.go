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


package services

import (
	"fmt"
	"io"
	"github.com/compostware/quokka/parse"
	"github.com/compostware/quokka/realize"
	"github.com/compostware/quokka/emit"
)


// A service interface providing mechanisms to build Dockerfiles from input specs.
type BuildService interface {
	// Builds a Dockerfile from an inputted spec and writes the results to an output file.
	Build(input io.Reader, output io.Writer)
}

// A constuctor method for creating a new BuildService.
func NewBuildService() BuildService {
	b := new(builder)
	b.parser = parse.NewParser()
	b.realizer = realize.NewRealizer()
	b.emitter = emit.NewEmitter()
	return b
}

type builder struct {
	parser parse.Parser
	realizer realize.Realizer
	emitter emit.Emitter
}

func (b *builder) Build(input io.Reader, output io.Writer) {
	fmt.Println("Building...")
	
	// parse
	spec := b.parser.Parse(input)
	
	// realize
	realization := b.realizer.Realize(spec)
	
	// emit
	b.emitter.Emit(realization, output)
}

