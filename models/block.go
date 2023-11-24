package models

import (
	"crypto/sha256"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

// TODO: Implement this function to add a new transaction to the block

// To do that, use the SetHash function. Feed the PrevBlockHash, Transactions, and Timestamp into the hash in this order
// concatenate them, and calculate a SHA256 hash on the concatenated combination.

func (b *Block) SetHash() {
	timestamp := []byte(string(b.Timestamp))

	var transactionsData []byte
	for _, tx := range b.Transactions {
		// Convert each transaction to bytes and concatenate
		// You might need to define the serialization of your transaction data
		// For example, JSON or other encoding format
		transactionsData = append(transactionsData, tx.Data...)
	}

	// Concatenate PrevBlockHash, Transactions, and Timestamp
	data := append(b.PrevBlockHash, transactionsData...)
	data = append(data, timestamp...)

	// Calculate SHA256 hash
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}
