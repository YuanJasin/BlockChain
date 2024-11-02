package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"main.go/transaction"
	"time"
)

type Block struct {
	Timestamp    int64                      `json:"timestamp"`
	Nonce        int                        `json:"nonce"`
	PreviousHash [32]byte                   `json:"previous_hash"`
	Transactions []*transaction.Transaction `json:"transactions"`
}

func (b *Block) Print() {
	fmt.Printf("timestamp			%d\n", b.Timestamp)
	fmt.Printf("nonce       		%d\n", b.Nonce)
	fmt.Printf("previousHash		%x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	// 把块内的所有数据转化位json
	m, _ := json.Marshal(b)
	//fmt.Println(string(m))
	// 通过sha256算法处理json
	return sha256.Sum256([]byte(m))
}

func NewBlock(nonce int, previousHsh [32]byte, transaction []*transaction.Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHsh,
		Transactions: transaction,
	}
}
