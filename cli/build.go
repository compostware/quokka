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

package cli

import (
	"fmt"
	"os"
	"github.com/compostware/quokka/services"
)


const (
	buildCmdKey = "build"
	buildCmdDesc = "Compiles a spec into a Dockerfile"
)


// Constructor for a Cmd that encapsulating the "build"-ing of a Dockerfile from a Spec. Exposes the BuildService
// thru the CLI.
func NewBuildCmd() Cmd {
	cmd := new(buildCmd)
	cmd.builder = services.NewBuildService()
	return cmd
}

type buildCmd struct {
	builder services.BuildService
}

type buildCmdArgs struct {
	inputFilename, outputFilename string
}

func (*buildCmd) Key() string {
	return buildCmdKey
}

func (*buildCmd) Description() string {
	return buildCmdDesc
}

func (cmd *buildCmd) Execute(args []string) error {
	cmdArgs, err := parseArgs(args)
	if err != nil {
		return err
	}
	
	inputFile, err := os.Open(cmdArgs.inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	
	outputFile, err := os.Create(cmdArgs.outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	
	return cmd.builder.Build(inputFile, outputFile)
}

func parseArgs(args []string) (*buildCmdArgs, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Insufficient number of arguments for build command: %n", len(args))
	}
	
	parsed := new(buildCmdArgs)
	parsed.inputFilename = args[0]
	parsed.outputFilename = args[1]
	
	return parsed, nil
}

func (*buildCmd) PrintUsage() {
	// TODO
}

