package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	pkBlockChain "github.com/simple_blockchain/models"
)

func printBlock(block *pkBlockChain.Block) {
	fmt.Println("  Thông tin block:")
	fmt.Println("  Timestamp:", block.Timestamp)
	fmt.Println("  PrevBlockHash:", hex.EncodeToString(block.PrevBlockHash))
	fmt.Println("  Hash:", hex.EncodeToString(block.Hash))

	fmt.Println("  Các giao dịch trong block:")
	for i, tx := range block.Transactions {
		fmt.Printf("   %d. %v \n", i+1, string(tx.Data))
	}
}

func main() {
	var blockChain pkBlockChain.Blockchain

	for {
		fmt.Println("=== Simple Blockchain Menu ===")
		fmt.Println("1. Add new block")
		fmt.Println("2. See latest block")
		fmt.Println("3. See all blocks")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)
		switch choice {
		case "1":
			var numTxs int
			fmt.Print("*Enter block data: \n")
			fmt.Print(" --Enter number of transactions: ")
			fmt.Scanln(&numTxs)
			var transactions []*pkBlockChain.Transaction
			for i := 0; i < numTxs; i++ {
				fmt.Print(" --Tx ", i, ": ")
				reader := bufio.NewReader(os.Stdin)
				inputString, err := reader.ReadString('\n')

				if err != nil {
					fmt.Println("Lỗi khi đọc dữ liệu:", err)
					return
				}

				inputString = strings.TrimSuffix(inputString, "\n")

				txHash := []byte(inputString)
				transaction := pkBlockChain.Transaction{Data: txHash}
				transactions = append(transactions, &transaction)
			}

			blockChain.AddBlock(transactions)
			fmt.Println("-----> Transaction added successfully!")
		case "2":
			latestBlock := blockChain.GetLastBlock()
			fmt.Print("*Latest block information: \n")
			printBlock(latestBlock)
		case "3":
			fmt.Println("=== Blockchain ===")
			// TODO: print all blocks in the blockchain
			for i, block := range blockChain.Blocks {
				fmt.Println("Block", i+1)
				printBlock(block)
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
