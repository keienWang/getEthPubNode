package trace

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/semaphore"
)

const maxConcurrentRequests = 10 // 最大并发请求数

func GetBlockTxTxR() {
	client, err := ethclient.Dial("http://192.168.110.146:8545")
	if err != nil {
		log.Fatal(err)
	}

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	startBlock := int64(5000000) // 从区块零开始
	sem := semaphore.NewWeighted(maxConcurrentRequests)

	var wg sync.WaitGroup

	for blockNum := startBlock; blockNum <= int64(latestBlock); blockNum++ {
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go func(blockNum int64) {
			defer wg.Done()
			defer sem.Release(1)
			processBlock(client, blockNum)
		}(blockNum)
	}

	wg.Wait()
	fmt.Println("All blocks and receipts have been processed.")
}

func processBlock(client *ethclient.Client, blockNum int64) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		log.Printf("Failed to fetch block %d: %v\n", blockNum, err)
		return
	}
	fmt.Println(block.Number())

	var wg sync.WaitGroup
	for _, tx := range block.Transactions() {
		wg.Add(1)
		go func(tx *types.Transaction) {
			defer wg.Done()
			processTransaction(client, tx)
		}(tx)
	}
	wg.Wait()
}

func processTransaction(client *ethclient.Client, tx *types.Transaction) {
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("Failed to fetch receipt for tx %s: %v\n", tx.Hash().Hex(), err)
		return
	}

	fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("Transaction Index: %d\n", receipt.TransactionIndex)
	fmt.Printf("Block Hash: %s\n", receipt.BlockHash.Hex())
	fmt.Printf("Cumulative Gas Used: %d\n", receipt.CumulativeGasUsed)
	fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
	fmt.Printf("Contract Address: %s\n", receipt.ContractAddress.Hex())
	fmt.Printf("Status: %d\n", receipt.Status)
	fmt.Printf("Effective Gas Price: %s\n", receipt.EffectiveGasPrice.String())
	fmt.Println()
}
