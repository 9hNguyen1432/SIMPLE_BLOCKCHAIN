package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/simple_blockchain/models"
// )

// func main() {
// 	transaction := models.Transaction{
// 		Data: []byte("Hello World!"),
// 	}

// 	fmt.Println(string(transaction.Data))

// 	for {
// 		fmt.Println("=== Simple Blockchain Menu ===")
// 		fmt.Println("1. Add new transaction")
// 		fmt.Println("2. See latest block")
// 		fmt.Println("3. See all blocks")
// 		fmt.Println("4. Exit")

// 		fmt.Print("Enter your choice: ")
// 		reader := bufio.NewReader(os.Stdin)
// 		choiceInput, _ := reader.ReadString('\n')
// 		choice := strings.TrimSpace(choiceInput)

// 		switch choice {
// 		case "1":
// 			fmt.Print("Enter transaction data: ")
// 			// data, _ := reader.ReadString('\n')
// 			// demo menu
// 			// TODO: add transaction to blockchain

// 			fmt.Println("Transaction added successfully!")
// 		case "2":
// 			// TODO: add transaction to blockchain

// 			// fmt.Printf("Block %d: %s\n", block.Index, block.Data)
// 		case "3":
// 			fmt.Println("=== Blockchain ===")
// 			// TODO: print all blocks in the blockchain

// 		case "4":
// 			fmt.Println("Exiting...")
// 			return
// 		default:
// 			fmt.Println("Invalid choice. Please try again.")
// 		}
// 	}
// }
