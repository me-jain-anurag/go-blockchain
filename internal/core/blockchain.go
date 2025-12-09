package core

// Blockchain maintains an ordered sequence of Blocks.
// It is essentially a linked list where every block points to the previous one via hash.
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates a new Blockchain instance with a Genesis Block.
// This function initializes the chain so it's never empty.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{NewGenesisBlock()},
	}
}

// AddBlock wraps the logic of creating a new block and appending it to the chain.
// It automatically handles finding the previous block's hash.
func (bc *Blockchain) AddBlock(data string) {
	// 1. Get the previous block (the current tip of the chain)
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// 2. Create the new block using the previous block's hash
	newBlock := NewBlock(data, prevBlock.Header.Hash)

	// 3. Append to the chain
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock creates the very first block in the chain.
// It has no previous hash (empty) and arbitrary data.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", Hash{})
}

// GetBlocks returns the slice of blocks.
// This is useful for iterating over the chain from outside the package.
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}
