package testrpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestRpc() {
	client, err := rpc.Dial("http://192.168.110.146:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RPC: %v", err)
	}
	defer client.Close()

	startBlock := 5000000
	endBlock := 5001000
	var totalDuration time.Duration
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 使用一个通道来限制并行数
	concurrency := 100

	sem := make(chan struct{}, concurrency)

	// 记录开始时间
	startTime := time.Now()

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		wg.Add(1)
		go func(blockNumber int) {
			defer wg.Done()
			sem <- struct{}{} // 获取一个令牌
			start := time.Now()

			var block *types.Block
			err := client.CallContext(context.Background(), &block, "eth_getBlockByNumber", fmt.Sprintf("0x%x", blockNumber), true)
			if err != nil {
				log.Fatalf("Failed to call eth_getBlockByNumber: %v", err)
			}

			// 获取所有交易
			var transactions []types.Transaction
			for _, tx := range block.Transactions() {
				transactions = append(transactions, *tx)
			}

			callDuration := time.Since(start)
			mu.Lock()
			totalDuration += callDuration
			mu.Unlock()
			<-sem // 释放一个令牌
		}(blockNumber)
	}

	wg.Wait()

	// 记录结束时间
	endTime := time.Now()

	// 计算并打印总的运行时间
	totalTime := endTime.Sub(startTime)
	fmt.Printf("Total time taken for fetching blocks from %d to %d: %v\n", startBlock, endBlock, totalTime)

}
