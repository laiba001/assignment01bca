package main

import (
    "assignment01bca"
    "fmt"
)

func main() {
    blockchain := assignment01bca.NewBlockchain()

    blockchain.AddBlock("bob to alice", 1)
    blockchain.AddBlock("alice to charlie", 2)
    blockchain.ListBlocks()

    // Change a block transaction
    blockchain.ChangeBlock(1, "alice to dave")
    blockchain.ListBlocks()

    // Verify the blockchain
    if blockchain.VerifyChain() {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is invalid.")
    }
}

