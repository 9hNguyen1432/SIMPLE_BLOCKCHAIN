package models

type Blockchain struct {
	blocks []*Block
}

// TODO: implement this function to add a new block to the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	block.PrevBlockHash = bc.GetLatestBlock().Hash
	block.SetHash()
	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func NewBlockchain(genesisData []byte) *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock(genesisData)}}
}

func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

func NewGenesisBlock(genesisData []byte) *Block {
	// The Genesis block has no previous block, so set PrevBlockHash to nil or an empty byte slice
	// You can also use a predefined constant for the Genesis block's PrevBlockHash
	emptyPrevBlockHash := []byte{}

	// Create the Genesis block
	genesisBlock := NewBlock([]*Transaction{NewTransaction(genesisData)}, emptyPrevBlockHash)
	genesisBlock.SetHash()

	return genesisBlock
}
