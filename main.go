package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	blockchain "github.com/simple_blockchain/models"
)

func main() {
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
			fmt.Print("*Enter block data: ")
			// TODO: add transaction to blockchain
			// input transactions number of transaction
			fmt.Print(" --Enter number of transactions: ")
			fmt.Scanln(&numTxs)

			var transactions []blockchain.Transaction
			for i := 0; i < numTxs; i++ {
				var txHash []byte
				fmt.Print("Tx ", i, ": ")
				fmt.Scanln(&txHash)
				transaction := blockchain.Transaction{Data: txHash}
				transactions = append(transactions, transaction)
			}

			fmt.Println("Transaction added successfully!")
		case "2":
			// TODO: add transaction to blockchain

			// fmt.Printf("Block %d: %s\n", block.Index, block.Data)
		case "3":
			fmt.Println("=== Blockchain ===")
			// TODO: print all blocks in the blockchain

		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
