package entity

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/sandronister/blockchain-go/pkg/utils"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &ProofOfWork{b, target}
}

func (p *ProofOfWork) InitData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			p.Block.PrevHash,
			p.Block.Data,
			utils.ToHex(int64(nonce)),
			utils.ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
}

func (p *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := p.InitData(nonce)
		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}

	}

	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := p.InitData(p.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1
}
