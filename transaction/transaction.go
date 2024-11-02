package transaction

import (
	"fmt"
	"strings"
)

type Transaction struct {
	SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	Value                      float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("senderBlockchainAddress	%s\n", t.SenderBlockchainAddress)
	fmt.Printf("RecipientBlockchainAddress	%s\n", t.RecipientBlockchainAddress)
	fmt.Printf("Value:	%.1f\n", t.Value)
}
