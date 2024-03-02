package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/simple_blockchain/models"
)

func main() {

	// init NewGenesisBlock(genesisData []byte) *Block with genesisData
	blockchain := models.NewBlockchain([]byte("Genesis Block"))
	for {
		fmt.Printf("\x1bc")
		fmt.Println("=== Simple Blockchain Menu ===")
		fmt.Println("1. Add new transactions")
		fmt.Println("2. See block")
		fmt.Println("3. See all blocks")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		reader := bufio.NewReader(os.Stdin)
		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			fmt.Printf("\x1bc")
			fmt.Println("Enter transaction data: ")

			// demo menu
			// TODO: add transaction to blockchain
			addTransactionToBlockchain(blockchain, inputTransactionData())

			fmt.Println("Transaction added successfully!")
		case "2":
			fmt.Printf("\x1bc")
			fmt.Println("=== Latest Block ===")
			printLatestBlock(blockchain)
			Blocks := blockchain.GetBlocks()
			i := len(Blocks) - 1
			schoice := "-1"
			for schoice != "3" {
				fmt.Println("1. See block before")
				fmt.Println("2. Verify this block")
				fmt.Println("3. Exit")
				fmt.Print("Enter your choice: ")
				reader := bufio.NewReader(os.Stdin)
				schoiceInput, _ := reader.ReadString('\n')
				schoice = strings.TrimSpace(schoiceInput)
				switch schoice {
				case "1":
					if i < 0 {
						fmt.Printf("\x1bc")
						fmt.Println("No block before")
					} else {
						i--
						fmt.Printf("\x1bc")
						fmt.Println("*Block information: \n")
						printBlock(Blocks[i])
					}
				case "2":
					// to do verify block
					if Blocks[i].VerifyBlock() {
						fmt.Println("Block is valid")
					} else {
						fmt.Println("Block is not valid")
					}
				case "3":
					break
				default:
					continue
				}

			}
		case "3":
			fmt.Printf("\x1bc")
			fmt.Println("=== Blockchain ===")
			// TODO: print all blocks in the blockchain

			printAllBlocks(blockchain)
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Enter transaction data function
func inputTransactionData() *models.Block {
	// loop to get transaction data into array
	transactions := []*models.Transaction{}
	for {
		fmt.Println("=== Enter transaction data: ===")

		fmt.Print("Enter name of sender: ")
		reader := bufio.NewReader(os.Stdin)
		sender, _ := reader.ReadString('\n')
		sender = strings.TrimSpace(sender)

		fmt.Print("Enter name of receiver: ")
		reader = bufio.NewReader(os.Stdin)
		receiver, _ := reader.ReadString('\n')
		receiver = strings.TrimSpace(receiver)

		fmt.Print("Enter amount: ")
		reader = bufio.NewReader(os.Stdin)
		amount, _ := reader.ReadString('\n')
		amount = strings.TrimSpace(amount)

		data := []byte(sender + " transfer " + receiver + ": " + amount)

		transactions = append(transactions, models.NewTransaction(data))

		fmt.Println("Transaction added successfully!")
		fmt.Print("Press 1 to stop, press other to add another transaction: ")
		reader = bufio.NewReader(os.Stdin)
		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)
		switch choice {
		case "1":
			return models.NewBlock(transactions, nil)
		default:
			continue
		}
	}

}

func addTransactionToBlockchain(blockchain *models.Blockchain, block *models.Block) {
	blockchain.AddBlock(block)
	fmt.Println("Transaction added successfully!")
}

func printLatestBlock(blockchain *models.Blockchain) {
	block := blockchain.GetLatestBlock()
	printBlock(block)
}

func printAllBlocks(blockchain *models.Blockchain) {
	for _, block := range blockchain.GetBlocks() {
		printBlock(block)
	}
}

func printBlock(block *models.Block) {

	fmt.Printf("Block: \n-Time: %d \n-Hash: %08b \n-PrevBlockHash: %08b\n", block.Timestamp, block.Hash, block.PrevBlockHash)
	models.DisplayTransactions(block)
}
