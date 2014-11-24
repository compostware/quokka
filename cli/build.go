package cli

import (
	"io"
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

func (*buildCmd) Key() string {
	return buildCmdKey
}

func (*buildCmd) Description() string {
	return buildCmdDesc
}

func (cmd *buildCmd) Execute(args []string) {
	var input *io.Reader = nil // TODO
	var output *io.Writer = nil // TODO
	
	cmd.builder.Build(input, output)
}

func (*buildCmd) PrintUsage() {
	// TODO
}

