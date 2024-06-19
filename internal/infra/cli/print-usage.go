package cli

import "fmt"

func (c *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -block BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  print - prints the blocks in the blockchain")
	fmt.Println("  validate - validate the blockchain")
	fmt.Println("  reindex - reindex the blockchain")
}
