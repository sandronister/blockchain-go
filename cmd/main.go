package main

import (
	"fmt"
	"strconv"

	"github.com/sandronister/blockchain-go/internal/entity"
)

func main() {
	chain := entity.InitBlockChain()

	chain.AddBlock("Block Zeus")
	chain.AddBlock("Block Hades")
	chain.AddBlock("Block Poseidon")
	chain.AddBlock("Block Ares")
	chain.AddBlock("Block Athena")
	chain.AddBlock("Block Apollo")

	for _, block := range chain.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := entity.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
