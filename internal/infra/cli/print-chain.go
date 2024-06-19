package cli

import (
	"fmt"
	"strconv"

	"github.com/sandronister/blockchain-go/internal/entity"
)

func (c *CommandLine) printChain() {
	iter := c.blockchain.GetIterator()

	for {
		block, err := iter.Next()

		if err != nil {
			break
		}

		fmt.Println("Block")
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		pow := entity.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}
