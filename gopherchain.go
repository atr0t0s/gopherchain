package main

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strconv"
	"container/list"
	"fmt"
)

type Block struct {

	index int
	timestamp time.Time
	data string
	previousHash string

}

func genesis() Block {

	var sha = sha256.New()
	sha.Write([]byte("This is Gopher, building your chain!"))

	genesisBlock := Block {0, time.Now(), "This is the Genesis block", hex.EncodeToString(sha.Sum(nil))}

	return genesisBlock

}

func nextBlock(lastBlock Block) Block {

	blockIndex := lastBlock.index + 1
	blockTime := time.Now().String()
	blockData := "This is block " + strconv.Itoa(blockIndex)
	previousHash := lastBlock.previousHash
	blockString := strconv.Itoa(blockIndex) + blockTime + blockData + previousHash

	var sha = sha256.New()
	sha.Write([]byte(blockString))

	return Block {blockIndex, time.Now(), blockData, hex.EncodeToString(sha.Sum(nil)) }

}

func main() {

	var blockchain = list.New()
	var genesisBlock = genesis()
	blockchain.PushBack(genesisBlock)
	var previousBlock = genesisBlock

	for e:= blockchain.Front(); e != nil; e = e.Next() {
		newBlock := nextBlock(previousBlock)
		blockchain.PushBack(newBlock)

		fmt.Printf("[HEIGHT]: ")
		fmt.Printf(strconv.Itoa(previousBlock.index))
		fmt.Printf("\n")
		fmt.Printf("Block ")
		fmt.Printf(previousBlock.previousHash)
		fmt.Printf(" has been added to the blockchain!\n")

		previousBlock = newBlock

		time.Sleep(1000 * time.Millisecond) //simulate block creation by delaying output

	}

}

