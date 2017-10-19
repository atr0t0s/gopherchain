package main

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strconv"
	"container/list"
	"fmt"
	"os"
  "math/rand"
	"encoding/binary"
)

type Block struct {

	index int
	timestamp time.Time
	data string
	pow string
	previousHash string
	thisHash string

}

func genesis() Block {

	genesisPrevHash := "0000000000000000000000000000000000000000000000000000000000000000"
	genesisHash := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	genesisBlock := Block {0, time.Now(), "This is the Genesis block", "Genesis", genesisPrevHash, genesisHash}

	return genesisBlock

}

func nextBlock(lastBlock Block) Block {

	blockIndex := lastBlock.index + 1 													// append the blockchain height for this block
	blockData := "This is block " + strconv.Itoa(blockIndex)		// just some data (should be json)
	previousHash := lastBlock.thisHash 													// string that contains the hash of the previous block
	rand.Seed(time.Now().UnixNano()) 														// seeding the randomizer based on current time
	nonce := make([]byte, 4)																		// create a 4 byte long nonce variable
	rand.Read(nonce)																						// populate nonce with a random value
	newnonce := binary.LittleEndian.Uint32(nonce)								// get the 4 byte integer representation of the nonce

	var blockStringSum string
	// ---------------------- <Proof of work> ------------------------------ //

	for (previousHash != "") {																	// to keep an infinite loop while we do work
		var sha = sha256.New()																		// new SHA256 "object"
			str := fmt.Sprint(newnonce)															// get the unsigned integer into a string
			blockStringSum = str + previousHash											// append the nonce to the hash of the previous block.
																															// - normally this should contain the version, previous hash,
																															// - nonce, block header, etc.
																															
			sha.Write([]byte(blockStringSum))
			blockStringSum = hex.EncodeToString(sha.Sum(nil))

			fmt.Printf(blockStringSum)
			fmt.Printf("\n")

			if (blockStringSum[:4] == "0000") {
				break
			} else {
				newnonce++
				sha.Reset()
			}
		}

	t, err := os.Create("height")
	check(err)
	t.WriteString(strconv.Itoa(blockIndex))
	fmt.Printf("\nWritten to file:\n")
	t.Sync()

	proof := hex.EncodeToString(nonce)
	// ---------------------- </Proof of work> ------------------------------ //

	return Block {blockIndex, time.Now(), blockData, proof, previousHash, blockStringSum }

}


func check(e error) {
	if e != nil {
			panic(e)
	}
}

func main() {

	var blockchain = list.New()
	var genesisBlock = genesis()
	blockchain.PushBack(genesisBlock)
	var previousBlock = genesisBlock

	f, err := os.OpenFile("dat", os.O_APPEND|os.O_WRONLY|
os.O_CREATE, 0600)
	check(err)

	for e:= blockchain.Front(); e != nil; e = e.Next() {

		newBlock := nextBlock(previousBlock)

		blockchain.PushBack(newBlock)

		fmt.Printf("[HEIGHT]: ")
		fmt.Printf(strconv.Itoa(newBlock.index))
		fmt.Printf("\n")
		fmt.Printf("Block ")
		fmt.Printf(newBlock.thisHash)
		fmt.Printf(" has been added to the blockchain!\n")

		heightLabel := "[HEIGHT]: "
		height := strconv.Itoa(newBlock.index)
		blockLabel := "Block: "
		block := newBlock.thisHash

		blockInfo := heightLabel + height + " \n" + blockLabel + block + "\n"
		n, err := f.WriteString(blockInfo)
		check(err)
		fmt.Printf("wrote %d bytes\n\n", n)
		f.Sync()

		previousBlock = newBlock

		time.Sleep(3000 * time.Millisecond) //simulate block creation by delaying output

	}

}
