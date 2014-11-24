package main

import (
	"github.com/compostware/quokka/cli"
)

func main() {
	handler := cli.NewCmdHandler()
	handler.Handle()
}