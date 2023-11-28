package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(fmt.Sprintf("%d", b.Timestamp))

	// create a MerkleTree from the transactions
	merkleTree := NewMerkleTree(b.Transactions)

	var transactionsData []byte
	// Concatenate PrevBlockHash, Transactions, and Timestamp
	data := append(transactionsData, b.PrevBlockHash...)
	data = append(data, merkleTree.RootNode.Data...)
	data = append(data, timestamp...)

	// Calculate SHA256 hash
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}
