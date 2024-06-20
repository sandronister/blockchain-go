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

	defer repository.Close()

	chain, err := blockchain.NewBlockChain(repository)
	catch.Handle(err)

	chain.Init()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	chain.AddBlock("Fourth block")

	iter := chain.GetIterator()

	for {
		block, err := iter.Next()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
