package cli

import (
	"flag"
	"os"
	"runtime"

	"github.com/sandronister/blockchain-go/pkg/catch"
)

func (c *CommandLine) Run() {
	c.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		catch.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		catch.Handle(err)
	default:
		c.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {

		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}

		c.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}
}
