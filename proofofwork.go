package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const (
	maxNonce = math.MaxInt64
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.blockHeader.bits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(int64(pow.block.blockHeader.version)),
			pow.block.blockHeader.prevBlockHash,
			pow.block.blockHeader.merkleRoot.hash,
			IntToHex(pow.block.blockHeader.timestamp),
			IntToHex(int64(pow.block.blockHeader.bits)),
			IntToHex(nonce),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.txs)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("%x", hash)
	fmt.Print("\n\n")

	return nonce, hash[:]
}
