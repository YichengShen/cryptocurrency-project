/*

	The telog package is an implementation of the Tamper-evident log.
	It supports two APIs:
		1. AddBlock: adds block to the end of the chain
		2. Check: takes the head and check whether there is some block that has been tampered

*/
package telog

import (
	"crypto/sha256"
	"fmt"
)

type hashPointer struct {
	pointer *block
	hash string
}

type block struct {
	hashPointer hashPointer
	data string
}

type Telog struct {
	head hashPointer
}

// Init initializes an empty head and a hash function used for telog with SHA-256.
func (t *Telog) Init() {
	t.head = hashPointer{}
}

// hashSha256 returns the hash digest of a block.
func (t *Telog) hashSha256(block block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", block))))

// AddBlock adds a block with data to the end of a tamper evident log.
func (t *Telog) AddBlock(data string) {
	// Create a new block.
	newBlock := block{
		// Use the old hash pointer of the head to connect the new block to the next latest block
		hashPointer: t.head,
		data: data,
	}
	
	// Hash the new block.
	newBlockHash := t.hashSha256(newBlock)

	// Connect the head to the new block with a hash pointer.
	t.head = hashPointer{
		pointer: &newBlock,
		hash: newBlockHash,
	}
}

// Check determines if the log has been tampered with,
// returning true if the log is valid and false if the log was tampered.
// TODO Check is returning false instead of true for an untampered log.
func (t *Telog) Check() bool {
	currentHashPointer := t.head
	emptyHashPointer := hashPointer{}

	// Execute as long as there is a non-empty hash pointer.
	for currentHashPointer != emptyHashPointer {
		// Access the block pointed to by the hash pointer
		currentBlock := *currentHashPointer.pointer

		// Rehash the block to check if it was tampered with.
		currentBlockHash := t.hashSha256(currentBlock)

		if currentBlockHash	!= currentHashPointer.hash {
			return false
		}

		// Iterate to next pointer
		currentHashPointer = currentBlock.hashPointer
	}
	return true
}