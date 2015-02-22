package cli

import (
	"bytes"
	"io"
    "io/ioutil"
    "os"
    "path"
    "testing"
	"github.com/compostware/quokka/services"
)

const (
	inputFileName = "input"
	outputFileName = "output"
	inputText = "turgid existence"
)

var (
	inputData = []byte("turgid existence")
)

func TestForValidInputAndOutputFiles(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	checkErr(err)
	defer os.RemoveAll(dir)
	
	inputPath := path.Join(dir, inputFileName)
	outputPath := path.Join(dir, outputFileName)
	
	err = ioutil.WriteFile(inputPath, inputData, os.ModeTemporary | os.ModePerm)
	checkErr(err)
	
	target := createTestCommand()
	
	args := []string { inputPath, outputPath }
	
	target.Execute(args)
	
	outputData, err := ioutil.ReadFile(outputPath)
	checkErr(err)
	
	if bytes.Compare(inputData[:], outputData[:]) != 0 {
		t.Errorf("Input file not correctly written to outputfile")
	} 
}

func createTestCommand() *buildCmd {
	cmd := new(buildCmd)
	cmd.builder = createStubService()
	return cmd
}

func createTmpFile() *os.File {
	f, err := ioutil.TempFile("", "")
	if (err != nil) {
		panic(err)
	}
	return f
}

func createStubService() services.BuildService {
	return new(stubBuildService)
}

type stubBuildService struct {
	
}

func (*stubBuildService) Build(input io.Reader, output io.Writer) {
	// stub services writes read input into output writer
	io.Copy(output, input)
}

func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}

