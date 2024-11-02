package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"main.go/block"
	"main.go/transaction"
	"main.go/utils"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "TEH BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	transactionPool   []*transaction.Transaction
	Chain             []*block.Block `json:"chain"`
	BlockchainAddress string         `json:"blockchainAddress"`
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) CreateBlock(nonce int, previousHsh [32]byte) *block.Block {
	b := block.NewBlock(nonce, previousHsh, bc.transactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.transactionPool = []*transaction.Transaction{}
	return b
}

func (bc *Blockchain) CopyTransactionPool() []*transaction.Transaction {
	transactions := make([]*transaction.Transaction, 0)
	for _, tx := range bc.transactionPool {
		transactions = append(transactions, transaction.NewTransaction(
			tx.SenderBlockchainAddress,
			tx.RecipientBlockchainAddress,
			tx.Value,
		))
	}
	return transactions
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.BlockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining ,status=success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			value := t.Value
			if blockchainAddress == t.RecipientBlockchainAddress {
				totalAmount += value
			}
			if blockchainAddress == t.SenderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*transaction.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := block.Block{
		Timestamp:    0,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.transactionPool
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) AddTransaction(
	sender string,
	recipient string,
	value float32,
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
) bool {
	t := transaction.NewTransaction(sender, recipient, value)

	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}
	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Panicln("ERROR: Transaction signature verification failed")
	}
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
	t *transaction.Transaction) bool {
	m, err := json.Marshal(t)

	if err != nil {
		log.Panicln("Failed to marshal transaction:", err)
	}
	h := sha256.Sum256(m)

	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) Print() {
	for i, b := range bc.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &block.Block{}
	bc := new(Blockchain)
	bc.BlockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}
