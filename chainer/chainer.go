package chainer

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	PrvHash string
	Data    string
	Hash    string
}

type BlockChain struct {
	Blocks []*Block
}

func MakeBlockChain() *BlockChain {
	genesisBlock := makeBlock("genesis hash", "genesis data")
	chain := []*Block{genesisBlock}
	return &BlockChain{chain}
}

func makeBlock(prvHash, data string) *Block {
	block := &Block{
		PrvHash: prvHash,
		Data:    data,
		Hash:    getHash(prvHash, data),
	}

	return block
}

func (chain *BlockChain) AppendBlock(data string) *Block {
	prvBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := makeBlock(prvBlock.Hash, data)
	chain.Blocks = append(chain.Blocks, newBlock)
	return newBlock
}

func getHash(prv, data string) string {
	hash := sha256.Sum256([]byte(prv + data))
	return hex.EncodeToString(hash[:])
}