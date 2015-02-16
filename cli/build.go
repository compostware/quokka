package cli

import (
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

func (cmd *buildCmd) Execute(args []string) {
	cmdArgs := parseArgs(args)
	
	inputFile, err := os.Open(cmdArgs.inputFilename)
	handleErr(err)
	defer inputFile.Close()
	
	outputFile, err := os.Create(cmdArgs.outputFilename)
	handleErr(err)
	defer outputFile.Close()
	
	cmd.builder.Build(inputFile, outputFile)
}

func parseArgs(args []string) (parsed *buildCmdArgs) {
	if len(args) != 2 {
		os.Exit(1)
	}
	
	parsed = new(buildCmdArgs)
	parsed.inputFilename = args[0]
	parsed.outputFilename = args[1]
	return
}

func (*buildCmd) PrintUsage() {
	// TODO
}

func handleErr(err error) {
	if (err != nil) {
		panic(err)
	}
}

