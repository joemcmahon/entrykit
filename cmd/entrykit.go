package main

import (
	"fmt"
	"os"

	"github.com/joemcmahon/entrykit"

	_ "github.com/joemcmahon/entrykit/codep"
	_ "github.com/joemcmahon/entrykit/prehook"
	_ "github.com/joemcmahon/entrykit/render"
	_ "github.com/joemcmahon/entrykit/switch"
)

var Version string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
