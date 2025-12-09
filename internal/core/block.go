package core

import (
	"crypto/sha256"
	"fmt"
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
// It separates the metadata from the actual transaction data structure (though in this simple version, Data is included here).
type Header struct {
	TimeStamp     int64  // TimeStamp is the Unix time when the block was created.
	Data          []byte // Data contains the actual information stored in the block.
	PrevBlockHash Hash   // PrevBlockHash is the hash of the preceding block in the chain.
	Hash          Hash   // Hash is the SHA-256 hash of this block's header.
}

// NewBlock creates and returns a new Block with the specified data and previous block hash.
// It automatically sets the current timestamp and calculates the block's hash immediately upon creation.
func NewBlock(data string, prevBlockHash Hash) *Block {
	h := &Header{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          nil, // Will be calculated below
	}

	h.Hash = h.CalculateHash()

	return &Block{
		Header: h,
	}
}

// CalculateHash generates a SHA-256 hash of the block header fields.
// It combines the timestamp, data, and previous hash into a single string record
// and returns the resulting SHA-256 sum as a slice.
func (h *Header) CalculateHash() Hash {
	record := fmt.Sprintf("%d%s%s", h.TimeStamp, h.Data, h.PrevBlockHash)

	hash := sha256.Sum256([]byte(record))

	return hash[:]
}

// String implements the fmt.Stringer interface.
// It formats the block's details into a readable string with hex-encoded hashes.
func (b *Block) String() string {
	return fmt.Sprintf(
		"Time: %d\nData: %s\nPrev: %x\nHash: %x\n",
		b.Header.TimeStamp,
		b.Header.Data,
		b.Header.PrevBlockHash,
		b.Header.Hash,
	)
}