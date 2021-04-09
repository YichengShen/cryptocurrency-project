/*

	The telog package is an implementation of the Tamper-evident log.
	It supports two APIs:
		1. AddBlock: Add block to the end of the chain.
		2. Check: Iterates through a log and checks whether there is some block that has been tampered.

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

// Init initializes an empty head used for the tamper evident log with SHA-256.
func (t *Telog) Init() {
	t.head = hashPointer{}
}

// hashSha256 returns the hash digest of a block.
func (t *Telog) hashSha256(block block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", block))))
}

// AddBlock adds a block with data to the right end of a tamper evident log, where the left end of the log is the first
// data block added and the right end of the log is the last data block added.
func (t *Telog) AddBlock(data string) {
	// Create a new block.
	newBlock := block{
		// Use the old hash pointer of the head to connect the new block to the right-most block
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

// Check determines if the log has been tampered with, returning true if the log is valid and false if the log was
// tampered with.
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