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
	h := initCmdHandler()
	h.RegisterCmd(NewBuildCmd())
	return h
}

func initCmdHandler() *cmdHandler {
	h := new(cmdHandler)
	h.cmds = make(map[string]Cmd)
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





