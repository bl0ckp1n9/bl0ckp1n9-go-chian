package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	// `json:""` - struct tag
	Data     string `json:"data"`
	Hash     string	`json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int	`json:"height"`
}

type blockchain struct {
	blocks []*Block
}

// Singleton
var b *blockchain
var once sync.Once // 오직 한번만 실행

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {

	newBlock := Block{data, "", getLastHash(),len(GetBlockchain().blocks)+1}
	newBlock.calculateHash()
	return &newBlock
}

func (chain *blockchain) AddBlock(data string) {
	chain.blocks = append(b.blocks, createBlock(data))
}

func (chain *blockchain) AllBlocks() []*Block {
	return chain.blocks
}

var ErrNotFound=errors.New("block not found")

func (chain *blockchain) GetBlock(height int) (*Block,error){
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1],nil
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
