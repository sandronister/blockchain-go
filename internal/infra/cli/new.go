package cli

import (
	blockchain "github.com/sandronister/blockchain-go/internal/usecase/block-chain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func NewCommandLine(chain *blockchain.BlockChain) *CommandLine {
	return &CommandLine{blockchain: chain}
}
