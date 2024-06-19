package main

import (
	"fmt"

	"github.com/sandronister/blockchain-go/configs"
	"github.com/sandronister/blockchain-go/internal/di"
	blockchain "github.com/sandronister/blockchain-go/internal/usecase/block-chain"
	"github.com/sandronister/blockchain-go/pkg/catch"
)

func main() {

	config, err := configs.LoadConfig(".")

	if err != nil {
		fmt.Println(err)
	}

	repository, err := di.NewBadgerRepository(config)
	catch.Handle(err)

	chain, err := blockchain.NewBlockChain(repository)
	catch.Handle(err)

	err = chain.Load()

	catch.Handle(err)

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
