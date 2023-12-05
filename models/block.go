package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Timestamp      int64
	Transactions   []*Transaction
	PrevBlockHash  []byte
	Hash           []byte
	MerkleRootNode *MerkleNode
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

	b.MerkleRootNode = merkleTree.RootNode
}

func (b *Block) GetHash() []byte {
	return b.Hash
}

func (block *Block) VerifyTransaction(indexTransaction int) bool {
	// Check if the specified transaction index exists
	if indexTransaction < 0 || indexTransaction >= len(block.Transactions) {
		return false
	}

	// Verify the transaction using the Merkle Tree
	merkleTree := NewMerkleTree(block.Transactions)
	rootHash := merkleTree.RootNode.Data

	// Compare the calculated root hash with the stored Merkle Root in the block
	return string(rootHash) == string(block.MerkleRootNode.Data)
}
