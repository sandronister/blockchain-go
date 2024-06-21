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

	//chain.AddBlock("First block")

	iter := chain.GetIterator()

	block, err := iter.Next()

	catch.Handle(err)

	//block, err := chain.GetBlock([]byte("0000342dc11a9fd1833ed9fe18ca5627cedc56507de6698acfcafd301398cb35"))

	catch.Handle(err)

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Println("---------------------------------------------------------------------------------------")

	// chain.AddBlock("Second block")
	// fmt.Println("Blockchain initialized")

	// iter := chain.GetIterator()
	// // fmt.Printf("Blockchain initialized\n\n")
	// // fmt.Printf("Blockchain blocks\n\n")
	// // fmt.Printf("Genesis block\n\n")

	// block, err := iter.Next()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// fmt.Printf("Data: %s\n", block.Data)
	// fmt.Printf("Hash: %x\n", block.Hash)
	// fmt.Println()

	// block, err = iter.Next()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// fmt.Printf("Data: %s\n", block.Data)
	// fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Println()

	// block, err = iter.Next()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// fmt.Printf("Data: %s\n", block.Data)
	// fmt.Printf("Hash: %x\n", block.Hash)
	// fmt.Println()

	// block, err = iter.Next()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// fmt.Printf("Data: %s\n", block.Data)
	// fmt.Printf("Hash: %x\n", block.Hash)
	// fmt.Println()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// fmt.Printf("Data: %s\n", block.Data)
	// fmt.Printf("Hash: %x\n", block.Hash)
	// fmt.Println()

}
