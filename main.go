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
		fmt.Println("=== Simple Blockchain Menu ===")
		fmt.Println("1. Add new transactions")
		fmt.Println("2. See latest block")
		fmt.Println("3. See all blocks")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		reader := bufio.NewReader(os.Stdin)
		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			fmt.Print("Enter transaction data: ")
			// demo menu
			// TODO: add transaction to blockchain
			addTransactionToBlockchain(blockchain, inputTransactionData())

			fmt.Println("Transaction added successfully!")
		case "2":
			fmt.Println("=== Latest Block ===")
			printLatestBlock(blockchain)
		case "3":
			fmt.Println("=== Blockchain ===")
			// TODO: print all blocks in the blockchain
			printAllBlocks(blockchain)
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
		fmt.Print("=== Enter transaction data: ===")

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

		data := []byte(sender + receiver + amount)

		transactions = append(transactions, models.NewTransaction(data))

		fmt.Println("Transaction added successfully!")
		fmt.Println("Press 1 to stop, press other to add another transaction: ")
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
	fmt.Printf("Block: - time: %d - transaction: %s - Hash: %s - MerkleProof: %s - PrevBlockHash: %s\n", block.Timestamp, block.Transactions, block.Hash, block.MerkleProof, block.PrevBlockHash)
}

func printAllBlocks(blockchain *models.Blockchain) {
	for _, block := range blockchain.GetBlocks() {
		fmt.Printf("Block: - time: %d - transaction: %s - Hash: %s - MerkleProof: %s - PrevBlockHash: %s\n", block.Timestamp, block.Transactions, block.Hash, block.MerkleProof, block.PrevBlockHash)
	}
}
