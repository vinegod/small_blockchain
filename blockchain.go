package main

type Blockchain struct {
	Blocks []*Block
}

func createBlockchain() *Blockchain {
	bc := Blockchain{[]*Block{createGenesis()}}
	return &bc
}

func (blockchain *Blockchain) AddBlock(txs string) {
	blocksnum := len(blockchain.Blocks)
	prevhash := blockchain.Blocks[blocksnum-1].blockHeader.blockHash

	newBlock := createNewBlock(txs, int32(blocksnum), prevhash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
