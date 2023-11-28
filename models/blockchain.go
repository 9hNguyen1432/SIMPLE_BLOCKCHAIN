package blockchain

import (
	"time"
)

type Blockchain struct {
	blocks []*Block
}

// get last block
func (bc *Blockchain) GetLastBlock() *Block {
	if bc.blocks == nil {
		bc.blocks = []*Block{}
		return nil
	}
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.GetLastBlock()

	var preHash []byte
	if prevBlock == nil {
		preHash = make([]byte, 32)
	} else {
		preHash = prevBlock.Hash
	}

	newBlock := NewBlock(transactions, preHash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Transactions: transactions, PrevBlockHash: []byte(prevHash)}
	block.SetHash()
	return block
}
