package main

import (
	"fmt"
)

func main() {
	bc := createBlockchain()

	bc.AddBlock("Den sent 0 coins to Block")
	bc.AddBlock("Den sent 5 coins to Block")

	for _, block := range bc.Blocks {
		Print(block)
		fmt.Println()
	}
}
