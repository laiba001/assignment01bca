package assignment01bca

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "sync"
)

// Block structure
type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

// Blockchain structure
type Blockchain struct {
    blocks []Block
    mu     sync.Mutex
}

// Create a new blockchain
func NewBlockchain() *Blockchain {
    return &Blockchain{}
}

// Create a new block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
    block := &Block{Transaction: transaction, Nonce: nonce, PreviousHash: previousHash}
    block.Hash = CalculateHash(block.Transaction, block.Nonce, block.PreviousHash)
    return block
}

// Calculate hash of a block
func CalculateHash(transaction string, nonce int, previousHash string) string {
    data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
    hash := sha256.New()
    hash.Write([]byte(data))
    return hex.EncodeToString(hash.Sum(nil))
}

// Add a new block to the blockchain
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
    bc.mu.Lock()
    defer bc.mu.Unlock()

    var previousHash string
    if len(bc.blocks) > 0 {
        previousHash = bc.blocks[len(bc.blocks)-1].Hash
    }

    newBlock := NewBlock(transaction, nonce, previousHash)
    bc.blocks = append(bc.blocks, *newBlock)
}

// List all blocks in the blockchain
func (bc *Blockchain) ListBlocks() {
    for i, block := range bc.blocks {
        fmt.Printf("Block %d: Transaction: %s, Nonce: %d, Previous Hash: %s, Hash: %s\n", 
            i, block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
    }
}

// Change transaction of a given block
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
    if index < 0 || index >= len(bc.blocks) {
        fmt.Println("Block index out of range")
        return
    }
    bc.blocks[index].Transaction = newTransaction
    // Recalculate the hash for the modified block
    bc.blocks[index].Hash = CalculateHash(bc.blocks[index].Transaction, bc.blocks[index].Nonce, bc.blocks[index].PreviousHash)
}

// Verify the blockchain
func (bc *Blockchain) VerifyChain() bool {
    for i := 1; i < len(bc.blocks); i++ {
        currentBlock := bc.blocks[i]
        previousBlock := bc.blocks[i-1]

        // Check if the previous hash is correct
        if currentBlock.PreviousHash != previousBlock.Hash {
            return false
        }
        // Check if the hash of the current block is correct
        if currentBlock.Hash != CalculateHash(currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash) {
            return false
        }
    }
    return true
}

