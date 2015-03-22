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

func (*stubBuildService) Build(input io.Reader, output io.Writer) error {
	// stub services writes read input into output writer
	_, err := io.Copy(output, input)
	return err
}

func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}

