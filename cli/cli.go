package cli

import (
	"os"
	"fmt"
)

// Encapsulates a single command accessed via a specific key.
type Cmd interface {
	// The unique key of this command.
	Key() string
	// An extended user-friendly description of this command.
	Description() string
	// Executes this command with the given array of arguments. These arguments will not contain: the command key; or
	// the CLI executable.
	Execute(args []string)
	// Print the usage of flags, arguments etc. for this command.
	PrintUsage()
}


// A registry of commands and handler for inputs from the CLI. 
type CmdHandler interface {
	// Registers the given Cmd in this handler.
	RegisterCmd(Cmd)
	// Identifies and executes the appropriate Cmd based on input specified in the command line arguments
	// of the executing process.
	Handle()
	// Executes the Cmd identified by the specified key, passing it the specified arguments.
	HandleCmd(key string, args []string)
	// Prints a complete descriptipon of a ussage of the CLI to stdout, based on the set of registered Cmds.
	PrintUsage()
}

// Constructor method for a new CmdHandler.
func NewCmdHandler() CmdHandler {
	h := new(cmdHandler)
	h.cmds = make(map[string]Cmd)
	h.RegisterCmd(NewBuildCmd())
	return h
}

type cmdHandler struct {
	cmds map[string]Cmd
}

func (h *cmdHandler) RegisterCmd(cmd Cmd) {
	key := cmd.Key()
	h.cmds[key] = cmd
}

func (h *cmdHandler) Handle() {
	args := os.Args
	
	if args == nil || len(args) < 2 {
		h.handleUsageError("Insufficient arguments")
		return
	} 
	
	key := args[1]
	cmdArgs := args[2:]
	
	h.HandleCmd(key, cmdArgs)
}

func (h *cmdHandler) PrintUsage() {
	// TODO
}

func (h *cmdHandler) HandleCmd(key string, args []string) {
	cmd := h.cmds[key]
	
	if cmd == nil {
		h.handleUsageError("Unrecognized command: " + key)
		return
	}
	
	cmd.Execute(args)
}

// Prints the given message with an error, accompanied by a print-out of general usage for the CLI..
func (h *cmdHandler) handleUsageError(message string) {
	fmt.Println(message)
	h.PrintUsage()
}





