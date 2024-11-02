package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"log"
	"main.go/utils"
)

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	Value                      float32 `json:"value"`
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		senderPrivateKey:           privateKey,
		senderPublicKey:            publicKey,
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, err := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	if err != nil {
		log.Panicln("Failed to generate signature:", err)
	}
	return &utils.Signature{
		R: r,
		S: s,
	}
}
