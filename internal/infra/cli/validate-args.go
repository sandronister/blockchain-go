package cli

import (
	"os"
	"runtime"
)

func (c *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		runtime.Goexit()
	}
}
