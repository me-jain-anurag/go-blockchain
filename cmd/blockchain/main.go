package main

import (
	"github.com/me-jain-anurag/go-blockchain/internal/core"
)

func main() {
	// 1. Initialize the Blockchain (Open DB)
	bc := core.InitBlockchain()
	
	// 2. Ensure DB is closed when CLI finishes
	defer bc.Close()

	// 3. Start the CLI
	cli := CLI{bc}
	cli.Run()
}