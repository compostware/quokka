package cli

import (
	"testing"
	"strings"
)

const (
	key = "key1"
	desc = "desc1"
)

func TestRegisterCmd() {
	target := new CmdHandler()
	cmd := createStubCmd(key, desc)
	
	target.RegisterCmd(cmd)
	
	actual := target.cmds[key]
	if actual != cmd {
		t.Errorf("Unable to register, unexpected cmd in map: %s", actual)
	}
}


func createStubCmd(key, desc string) *stubCmd {
	stubCmd = stubCmd{key: key, desc: desc}
	return &stubCmd
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
	output = strings.Join(args, "-")
}

func (*stubCmd) PrintUsage() {
	// do nothing
}

