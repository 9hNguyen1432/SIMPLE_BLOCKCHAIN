package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	pkBlockChain "github.com/simple_blockchain/models"
)

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
			// TODO: add transaction to blockchain
			// input transactions number of transaction
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
				println(string(txHash))
				transaction := pkBlockChain.Transaction{Data: txHash}
				transactions = append(transactions, &transaction)
			}

			blockChain.AddBlock(transactions)
			fmt.Println("-----> Transaction added successfully!")
		case "2":
			latestBlock := blockChain.GetLastBlock()
			fmt.Print("*Latest block information: \n")
			fmt.Println("  Thông tin latest block:")
			fmt.Println("  Timestamp:", latestBlock.Timestamp)
			fmt.Println("  PrevBlockHash:", hex.EncodeToString(latestBlock.PrevBlockHash))
			fmt.Println("  Hash:", hex.EncodeToString(latestBlock.Hash))

			fmt.Println("  Các giao dịch trong block:")
			for i, tx := range latestBlock.Transactions {
				fmt.Printf("   %d. %v \n", i+1, string(tx.Data))
			}
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
