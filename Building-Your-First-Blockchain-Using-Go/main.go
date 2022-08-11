package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data          []byte
	PrevBlockHash []byte
	Timestamp     int64
	Nonce         int
}

func (b *Block) String() string {
	return fmt.Sprintf("Block %x\nPrev. block: %x\nData: %s\nNonce: %d\nTimestamp: %d\n",
		b.Hash(), b.PrevBlockHash, b.Data, b.Nonce, b.Timestamp)
}

func (b *Block) Hash() []byte {
	data := append(b.PrevBlockHash, b.Data...)
	data = append(data, b.Timestamp, b.Nonce)
	return sha256.Sum256(data)
}

func mineBlock(block *Block) {
	for {
		hash := block.Hash()
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 && hash[3] == 0 {
			fmt.Printf("Mined block: %x\n", hash)
			break
		}

		block.Nonce++
	}
}

type Blockchain struct {
	Blocks       []*Block
	GenesisBlock *Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash())
	bc.Blocks = append(bc.Blocks, newBlock)
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")

	for _, block := range bc.Blocks {
		fmt.Println(block.String())
	}
}
