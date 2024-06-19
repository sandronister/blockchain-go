package cli

import "fmt"

func (c *CommandLine) addBlock(data string) {
	c.blockchain.AddBlock(data)
	fmt.Println("Block added!")
}
