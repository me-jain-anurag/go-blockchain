package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/me-jain-anurag/go-blockchain/internal/core"
)

// CLI is the command-line interface manager.
// It holds a reference to the blockchain to perform actions on it.
type CLI struct {
	bc *core.Blockchain
}

// printUsage displays the available commands to the user.
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA: Adds a block to the blockchain")
	fmt.Println("  printchain				: Prints all the blocks of the blockchain")
}

// validateArgs ensures the user provided at least one command.
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1) // Exits the application cleanly
	}
}

// addBlock adds a new block to the chain.
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

// printchain iterates over the blockchain and prints every block.
func (cli *CLI) printChain() {
	iter := cli.bc.Iterator()

	for {
		block := iter.Next()
		
		if block == nil {
			break
		}
		// Since we  have implemented String() for block
		// we can print the block directly
		fmt.Println(block)

		// Validate PoW for the block
		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %v\n", pow.Validate())
		fmt.Println()
	}
}

// Run parses the command-line arguments and executes the appropriate logic.
func (cli *CLI) Run() {
	cli.validateArgs()

	// Define our subcommands using Go's "flag" package
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	
	// Define flags for subcommands
	// This says: "If the user uses 'addblock', look for a '-data' flag"
	addBlockData := addBlockCmd.String("data", "", "Block data")

	// Check the first argument (e.g., "addblock" or "printchain")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		core.Handle(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		core.Handle(err)
	default:
		cli.printUsage()
		os.Exit(1)
	}

	// Logic for 'addblock' command
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	// Logic for 'printchain' command
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
