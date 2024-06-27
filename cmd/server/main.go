package main

import (
	"fmt"

	"github.com/sandronister/blockchain-go/configs"
	"github.com/sandronister/blockchain-go/internal/di"
	"github.com/sandronister/blockchain-go/internal/infra/cli"
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

	chain := blockchain.NewBlockChain(repository)

	catch.Handle(err)

	chain.Init()

	client := cli.NewCommandLine(chain)

	client.Run()
}
