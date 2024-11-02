package main

import (
	"fmt"
	"main.go/blockchain"
	"main.go/wallet"
)

func main() {
	// 创建钱包
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// 创建交易请求(用于验证)
	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0)

	// 创建区块链
	chain := blockchain.NewBlockchain(walletM.BlockchainAddress())
	// 添加交易到区块链交易池(用于存储)
	isAdded := chain.AddTransaction(
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
		walletA.PublicKey(),
		t.GenerateSignature(),
	)

	fmt.Println("Added?", isAdded)
}
