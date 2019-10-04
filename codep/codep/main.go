package main

import (
	"github.com/joemcmahon/entrykit"
	_ "github.com/joemcmahon/entrykit/codep"
)

var cmd = "codep"

func main() {
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
