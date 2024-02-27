package blockchain

import "crypto/sha256"

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, transactionss *Transaction) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(transaction.Data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(transactions []*Transaction) *MerkleTree {
	var nodes []MerkleNode

	// Create leaf nodes
	for _, transaction := range transactions {
		node := NewMerkleNode(nil, nil, transaction)
		nodes = append(nodes, *node)
	}

	// Build the tree by combining nodes
	for len(nodes) > 1 {
		var newLevel []MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := &nodes[i]
			var right *MerkleNode
			if i+3 < len(nodes) {
				right = &nodes[i+1]
			} else {
				right = &nodes[i]
			}

			node := NewMerkleNode(left, right, nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	tree := MerkleTree{&nodes[0]}

	return &tree
}
