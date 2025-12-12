package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Hash represents a cryptographic SHA-256 hash as a byte slice.
type Hash []byte

// Block represents a single unit of data in the blockchain.
// It contains a pointer to its Header which holds the metadata.
type Block struct {
	Header *Header
}

// Header contains the metadata and cryptographic information for a block.
// It separates the metadata from the actual transaction data structure.
type Header struct {
	TimeStamp     int64  // TimeStamp is the Unix time when the block was created.
	Data          []byte // Data contains the actual information stored in the block.
	PrevBlockHash Hash   // PrevBlockHash is the hash of the preceding block in the chain.
	Hash          Hash   // Hash is the SHA-256 hash of this block's header.
	Nonce         int    // Nonce is the number used once to satisfy the Proof of Work condition.
}

// NewBlock creates and returns a new Block with the specified data and previous block hash.
// It performs the Proof of Work mining process to find a valid hash and nonce before returning.
func NewBlock(data string, prevBlockHash Hash) *Block {
	h := &Header{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          nil, // Will be calculated via Proof of Work
		Nonce:         0,
	}

	// Create a temporary block wrapper to pass to the PoW engine
	block := &Block{Header: h}

	// Initialize the Proof of Work engine
	pow := NewProofOfWork(block)

	// Mine the block (this is a blocking operation)
	nonce, hash := pow.Run()

	// Update the header with the winning nonce and hash
	h.Hash = hash
	h.Nonce = nonce

	return block
}

// String implements the fmt.Stringer interface.
// It formats the block's details into a readable string with hex-encoded hashes.
func (b *Block) String() string {
	return fmt.Sprintf(
		"Time: %d\nData: %s\nPrev: %x\nHash: %x\nNonce: %d\n",
		b.Header.TimeStamp,
		b.Header.Data,
		b.Header.PrevBlockHash,
		b.Header.Hash,
		b.Header.Nonce,
	)
}

// Serialize converts the Block structure into a byte slice.
// This is necessary because the database can only store pure bytes, not Go structs.
// It uses the encoding/gob package for efficient binary serialization.
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	Handle(err)

	return result.Bytes()
}

// DeserializeBlock takes a byte slice and decodes it back into a Block struct.
// It reverses the Serialize process.
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)
	Handle(err)

	return &block
}

// Handle is a helper function to check for errors and panic if one occurs.
// In a production app, we might handle this more gracefully, but for now, we stop execution.
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}