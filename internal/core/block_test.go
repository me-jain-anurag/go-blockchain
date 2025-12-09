package core

import (
	"testing"
)

// TestNewBlock verifies that a new block can be created and that it generates a valid hash.
func TestNewBlock(t *testing.T) {
	// 1. Setup: Create a "Genesis Block" scenario
	// The first block in a chain has no previous hash, so we use an empty byte slice.
	prevHash := []byte{}
	data := "Genesis Block"

	// 2. Execution: Call the function we are testing
	b := NewBlock(data, prevHash)

	// 3. Assertion: Verify the results
	// We check that the Hash field was actually populated by the constructor.
	if len(b.Header.Hash) == 0 {
		t.Error("NewBlock returned a block with an empty Hash")
	}

	// Optional: Also verify that the Data was set correctly
	if string(b.Header.Data) != data {
		t.Errorf("Expected block data to be %s, but got %s", data, b.Header.Data)
	}
}