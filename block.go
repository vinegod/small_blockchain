package main

import (
	"fmt"
	"time"
)

const (
	version        = 1
	difficultyBits = 4 * 4
)

type MerkleNode struct {
	hash []byte
}

type BlockHeader struct {
	version       int32
	prevBlockHash []byte
	merkleRoot    *MerkleNode
	timestamp     int64
	bits          int
	nonce         int64
	blockHash     []byte
}

type Block struct {
	Height      int32
	blockSize   int32
	blockHeader *BlockHeader
	txCount     int32
	txs         string
}

func createNewBlock(txs string, height int32, prevHash []byte) *Block {
	block := &Block{height, 1, nil, 1, txs}
	mr := &MerkleNode{MerkleRootHash(block)}
	bh := &BlockHeader{version, prevHash, mr, time.Now().Unix(), difficultyBits, 0, []byte{}}
	block.blockHeader = bh
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	bh.blockHash = hash
	bh.nonce = nonce
	return block
}

func createGenesis() *Block {
	return createNewBlock("Genesis block", 0, []byte{})
}

func Print(block *Block) {
	fmt.Printf("Height: %d\n", block.Height)
	fmt.Printf("Blocksize: %d\n", block.blockSize)
	PrintHeader(block.blockHeader)
	fmt.Printf("txCount: %d\n", block.txCount)
	fmt.Printf("data: %s\n", block.txs)

}

func PrintHeader(header *BlockHeader) {
	fmt.Println("Blockheader: ")
	fmt.Println("\tversion: ", header.version)
	fmt.Printf("\tprevBlockhash: %x\n", header.prevBlockHash)
	fmt.Printf("\tmerkleRoot: %x\n", header.merkleRoot.hash)
	fmt.Printf("\ttimestamp: %x\n", header.timestamp)
	fmt.Printf("\tbits: %x\n", header.bits)
	fmt.Printf("\tnonce: %x\n", header.nonce)
	fmt.Printf("\tblockHash: %x\n", header.blockHash)
}
