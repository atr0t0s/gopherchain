package main

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strconv"
	"container/list"
	"fmt"
  "math/rand"
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

func doWork(newBlock Block) string {

	var sha = sha256.New()
	tryString := "Try it out!"

	nonce := make([]byte, 4)
  rand.Read(nonce)

	hashString := "00000" + newBlock.previousHash[5:len(newBlock.previousHash)]

	for ((tryString[:5]) != hashString[:5]) {


			sha.Write([]byte(tryString))
			tryString = hex.EncodeToString(sha.Sum(nil))
			tryString = tryString + hex.EncodeToString(nonce)
			fmt.Printf(hex.EncodeToString(nonce))
			//fmt.Printf(tryString)
			fmt.Printf("\n\n")

			nonce[0]++
	}

	return "Found!"
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
		proof := doWork(newBlock)

		if (proof == "Found!") {
			blockchain.PushBack(newBlock)

			fmt.Printf("[HEIGHT]: ")
			fmt.Printf(strconv.Itoa(previousBlock.index))
			fmt.Printf("\n")
			fmt.Printf("Block ")
			fmt.Printf(previousBlock.previousHash)
			fmt.Printf(" has been added to the blockchain!\n")

			proof = "Unfound"

		}

		previousBlock = newBlock

		time.Sleep(10000 * time.Millisecond) //simulate block creation by delaying output

	}

}
