/*

	The telog package is an implementation of the Tamper-evident log.
	It supports two APIs:
		1. addBlock: adds block to the end of the chain
		2. check: takes the head and check whether there is some block that has been tampered

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

// Returns a empty hash pointer
// Used when creating the genesis block
func CreateEmptyHashPointer() hashPointer {
	return hashPointer{}
}

// Returns the hash digest
func hashSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Adds a new block into the log
// Returns the head hash pointer
// Leave the hashPointer empty if the block is a genesis block
func AddBlock(data string, head hashPointer) hashPointer {
	newBlock := block{
		hashPointer: head,
		data: data,
	}

	hashDigest := hashSha256(newBlock)

	newHead := hashPointer{
		pointer: &newBlock,
		hash: hashDigest,
	}

	return newHead
}

// Recursively check if the log is tampered
func Check(head hashPointer) bool {
	// Base Case: Reached genesis block
	if head == (hashPointer{}) {
		fmt.Println("reached genesis")
		return true
	}

	prevHashPointer := (*head.pointer).hashPointer
	// Recursive Case
	if Check(prevHashPointer) {
		fmt.Println("Previous block", *head.pointer)
		fmt.Println("-----")
		prevBlockHash := hashSha256(*head.pointer)
		if head.hash == prevBlockHash {
			return true
		}
	}

	return false
}

// TODO: add attack to test check