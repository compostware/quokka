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
	"testing"
	"strings"
)

const (
	key = "key1"
	desc = "desc1"
)

func TestRegisterCmd(t *testing.T) {
	target := initCmdHandler()
	cmd := createStubCmd(key, desc)
	
	target.RegisterCmd(cmd)
	
	actual := target.cmds[key]
	if actual != cmd {
		t.Errorf("Unable to register, unexpected cmd in map: %s", actual)
	}
}


func createStubCmd(key, desc string) *stubCmd {
	cmd := stubCmd{key: key, desc: desc}
	return &cmd
}

type stubCmd struct {
	key string
	desc string
	output string
}

func (*stubCmd) Key() string {
	return key
}

func (*stubCmd) Description() string {
	return desc
}

func (cmd *stubCmd) Execute(args []string) error {
	cmd.output = strings.Join(args, "-")
	return nil
}

func (*stubCmd) PrintUsage() {
	// do nothing
}

