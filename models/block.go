package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	MerkleProof   []byte
}

// constructor
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	//timestamp := time.Now().Unix()
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		MerkleProof:   []byte{},
	}
	block.MerkleProof = block.GetMerkleProof()

	return block
}

func (b *Block) toStr() string {
	return "Timestamp: " + string(b.Timestamp) +
		"\nHashTransactions: " + string(b.HashTransactions()) +
		"\nPrevBlockHash: " + string(b.PrevBlockHash) +
		"\nHash: " + string(b.Hash) +
		"\nMerkleProof: " + string(b.MerkleProof)
}

// TODO: Implement this function to add a new transaction to the block

// To do that, use the SetHash function. Feed the PrevBlockHash, Transactions, and Timestamp into the hash in this order
// concatenate them, and calculate a SHA256 hash on the concatenated combination.

func (b *Block) SetHash() {
	timestamp := []byte(string(b.Timestamp))

	var transactionsData []byte
	transactionsData = b.HashTransactions()

	// create a MerkleTree from the transactions

	// Concatenate PrevBlockHash, Transactions, and Timestamp
	data := append(b.PrevBlockHash, transactionsData...)
	data = append(data, timestamp...)

	// Calculate SHA256 hash
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

func (b *Block) HashTransactions() []byte {

	merkleTree := NewMerkleTree(b.Transactions)
	// Calculate SHA256 hash
	b.MerkleProof = merkleTree.RootNode.Data
	return merkleTree.RootNode.Data
}

func (b *Block) GetMerkleProof() []byte {
	//Merkle proof
	// Create a list to store the hashes of the transactions
	var transactionHashes [][]byte
	for _, tx := range b.Transactions {
		// Convert each transaction to bytes and calculate its hash
		// You might need to define the serialization of your transaction data
		// For example, JSON or other encoding format
		txHash := sha256.Sum256(tx.Data)
		transactionHashes = append(transactionHashes, txHash[:])
	}

	// Build the Merkle tree by repeatedly hashing pairs of transaction hashes
	for len(transactionHashes) > 1 {
		var nextLevelHashes [][]byte
		for i := 0; i < len(transactionHashes); i += 2 {
			// Concatenate the pair of hashes
			concatenatedHash := append(transactionHashes[i], transactionHashes[i+1]...)

			// Calculate the hash of the concatenated pair
			hash := sha256.Sum256(concatenatedHash)
			nextLevelHashes = append(nextLevelHashes, hash[:])
		}
		transactionHashes = nextLevelHashes
	}

	// The last remaining hash in the transactionHashes list is the Merkle root
	return transactionHashes[0]
}

func DisplayTransactions(b *Block) {
	fmt.Println("Transactions:")
	for index, tx := range b.Transactions {
		fmt.Printf("\t%d. %s \n", index, tx.toStr())
	}
}

func (block *Block) VerifyBlock() bool {
	// Verify the transaction using the Merkle Tree
	merkleTree := NewMerkleTree(block.Transactions)
	rootHash := merkleTree.RootNode.Data

	// Compare the calculated root hash with the stored Merkle Root in the block
	return string(rootHash) == string(block.MerkleProof)
}
