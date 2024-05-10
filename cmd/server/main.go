package main

import (
	"fmt"

	"github.com/sandronister/blockchain-go/configs"
	"github.com/sandronister/blockchain-go/internal/di"
	blockchain "github.com/sandronister/blockchain-go/internal/usecase/block-chain"
	"github.com/sandronister/blockchain-go/pkg/catch"
)

func main() {

	config, _ := configs.LoadConfig(".")

	repository, err := di.NewBadgerRepository(config)

	catch.Handle(err)

	chain, err := blockchain.NewBlockChain(repository)

	catch.Handle(err)

	chain.AddBlock("Block Zeus")
	chain.AddBlock("Block Hades")
	chain.AddBlock("Block Poseidon")
	chain.AddBlock("Block Ares")
	chain.AddBlock("Block Athena")
	chain.AddBlock("Block Apollo")

	iter := chain.GetIterator()

	for {

		block, err := iter.Next()

		if err != nil {
			break
		}

		fmt.Println("Block")
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Println()

	}

}
