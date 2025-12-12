package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// difficulty is a static constant that defines how hard it is to mine a block.
// In a real blockchain (like Bitcoin), this adjusts dynamically over time.
// For now, 18 is a reasonable number for quick testing on a laptop.
const difficulty = 18

// ProofOfWork represents the mechanism to secure the block.
// It holds a pointer to the block being mined and the target hash it needs to beat.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds a ProofOfWork structure for a specific block.
// It calculates the "target" hash based on the difficulty constant.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	
	// We shift the number 1 to the left by (256 - difficulty).
	// This creates a massive number that serves as the "ceiling".
	// Any hash we generate must be numerically smaller than this target.
	target.Lsh(target, uint(256-difficulty))

	return &ProofOfWork{b, target}
}

// prepareData merges the block fields with the nonce into a single byte slice.
// This byte slice is what gets hashed in the mining loop.
func (p *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			p.block.Header.PrevBlockHash,
			p.block.Header.Data,
			IntToBytes(p.block.Header.TimeStamp),
			IntToBytes(int64(difficulty)),
			IntToBytes(int64(nonce)),
		},
		[]byte{},
	)
}

// Run executes the mining process.
// It effectively loops forever, incrementing the nonce until it finds a hash
// that is lower than the target.
// Returns:
//   1. The winning nonce (int)
//   2. The resulting hash ([]byte)
func (p *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", p.block.Header.Data)

	// Loop until we reach the maximum integer size (practically infinite)
	for nonce < math.MaxInt64 {
		// 1. Prepare the data with the current nonce
		data := p.prepareData(nonce)
		
		// 2. Hash the data using SHA-256
		hash = sha256.Sum256(data)
		
		// 3. Convert the hash result to a big integer so we can compare it
		hashInt.SetBytes(hash[:])

		// 4. Compare our hash with the target
		// Cmp returns -1 if x < y
		if hashInt.Cmp(p.target) == -1 {
			fmt.Printf("\r%x", hash) // Print the winning hash
			break
		} else {
			nonce++ // Try again with a new number
		}
	}
	fmt.Printf("\n\n")

	return nonce, hash[:]
}

// Validate performs a quick check to ensure the block's hash satisfies the mining difficulty.
// Unlike Run() which searches for a nonce (CPU intensive), Validate() just checks
// if the existing nonce and data result in a valid hash (Instant).
// This is what other nodes in the network run to verify a block before accepting it.
func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int

	// 1. Re-create the data using the nonce stored in the block header
	data := p.prepareData(p.block.Header.Nonce)

	// 2. Hash the data
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	// 3. Check if the generated hash is less than the target difficulty
	// valid = (hashInt < target)
	return hashInt.Cmp(p.target) == -1
}

// IntToBytes is a utility function that converts an int64 to a byte slice.
// It uses BigEndian encoding, which is standard for network protocols.
func IntToBytes(num int64) []byte {
	buff := new(bytes.Buffer)
	// Write the number into the buffer using BigEndian order
	err := binary.Write(buff, binary.BigEndian, num)
	Handle(err)
	return buff.Bytes()
}