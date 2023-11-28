package blockchain

import "time"

type Blockchain struct {
	blocks []*Block
}

// get last block
func (bc *Blockchain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.GetLastBlock()
	newBlock := NewBlock(transactions, prevBlock.GetHash())
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Transactions: transactions, PrevBlockHash: []byte(prevHash)}
	block.SetHash()
	return block
}
