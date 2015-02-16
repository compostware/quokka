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

func (cmd *stubCmd) Execute(args []string) {
	cmd.output = strings.Join(args, "-")
}

func (*stubCmd) PrintUsage() {
	// do nothing
}

