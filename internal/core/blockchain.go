package core

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

// dbPath is the directory where the BadgerDB data will be stored on disk.
const (
	dbPath = "./tmp/blocks"
)

// Blockchain maintains the state of the chain.
// LastHash stores the hash of the newest block in the chain.
// db is the pointer to the open BadgerDB connection.
type Blockchain struct {
	LastHash Hash
	db       *badger.DB
}

// InitBlockchain opens the database and loads the blockchain state.
// If the blockchain does not exist, it creates a new one with a Genesis block.
// If it does exist, it loads the last known hash from the DB.
func InitBlockchain() *Blockchain {
	var lastHash Hash

	// Set up BadgerDB options (specify path and disable default logging)
	opts := badger.DefaultOptions(dbPath)
	opts.Logger = nil

	// Open the database connection
	db, err := badger.Open(opts)
	Handle(err)

	// Start a Read-Write transaction to check/update the chain state
	err = db.Update(func(txn *badger.Txn) error {
		// Try to fetch the "lh" (Last Hash) key
		item, err := txn.Get([]byte("lh"))

		// Case: No blockchain found in DB
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found. Creating genesis block...")

			genesis := NewGenesisBlock()

			// Save Genesis block to DB
			err := txn.Set(genesis.Header.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			// Set the Last Hash key to point to Genesis
			err = txn.Set([]byte("lh"), genesis.Header.Hash)
			lastHash = genesis.Header.Hash
			return err

		} else {
			// Case: Blockchain exists, just load the tip
			item.Value(func(val []byte) error {
				lastHash = append([]byte{}, val...)
				return nil
			})
			return nil
		}
	})

	Handle(err)

	return &Blockchain{lastHash, db}
}

// AddBlock mines a new block containing the provided data and saves it to the database.
// It updates the "LastHash" key in the DB to point to this new block.
func (bc *Blockchain) AddBlock(data string) {
	lastHash := bc.LastHash

	// Mine the new block (expensive operation)
	newBlock := NewBlock(data, lastHash)

	// Store the new block in the database
	err := bc.db.Update(func(txn *badger.Txn) error {
		// Save the Block data (Key: BlockHash, Val: SerializedBlock)
		err := txn.Set(newBlock.Header.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}

		// Update the Last Hash key to point to this new block
		err = txn.Set([]byte("lh"), newBlock.Header.Hash)
		bc.LastHash = newBlock.Header.Hash
		return err
	})

	Handle(err)
}

// NewGenesisBlock creates the very first block in the chain.
// It has no previous hash (empty) and arbitrary data.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", Hash{})
}

// BlockchainIterator is a helper struct to traverse the blockchain.
// It keeps track of the "current" hash it is looking at.
type BlockchainIterator struct {
	CurrentHash Hash
	db			*badger.DB
}

// Iterator returns a new BlockchainIterator struct starting at the tip of the chain.
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{
		bc.LastHash,
		bc.db,
	}
}

// Next returns the next block from the chain (traversing backwards)
// and advances the iterator to the previous block.
// returns nil if there are no more blocks (Genesis reached).
func (iter *BlockchainIterator) Next() *Block {
	// Get the current block
	var block *Block

	// If we are already empty (Genesis prev hash is empty), stop.
	if len(iter.CurrentHash) == 0 {
		return nil
	}

	err := iter.db.View(func (txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			return err
		}

		// Deserialize the block
		err = item.Value(func(val []byte) error {
			block = DeserializeBlock(val)
			return nil
		})
		return err
	})

	Handle(err)
	
	// Move the cursor backward to the previous block
	iter.CurrentHash = block.Header.PrevBlockHash

	return block
}

// Close terminates the database connection.
// It is crucial to call this before the program exits to prevent data corruption.
func (bc *Blockchain) Close() {
	bc.db.Close()
}