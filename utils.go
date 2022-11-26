package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func MerkleRootHash(block *Block) []byte {
	hash := sha256.Sum256([]byte(block.txs))
	hash = sha256.Sum256(hash[:])
	return hash[:]
}
