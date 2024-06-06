package main

import (
	getpeers "get_eth_node/getPeers"
	"math/big"
)

func main() {

	// trace.GetBlockTxTxR()
	// for {
	// 	getaccounts.GetAccount()
	// 	time.Sleep(1 * time.Second)
	// }

	getpeers.GetPeers()

	// client, _ := ethclient.Dial("https://taiko.blockpi.network/v1/rpc/public")

	// block, err := client.BlockByNumber(context.Background(), nil)
	// if err != nil {

	// 	log.Fatalf("Failed to get latest block: %v", err)
	// }

	// // 获取 gas price
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed to get suggested gas price: %v", err)
	// }

	// // 打印结果
	// fmt.Println("Current Gas Price: Gwei\n", gasPriceInGwei(gasPrice))
	// fmt.Printf("Latest Block Number: %d\n", block.Number().Uint64())
}

// 将 gas price 转换为 Gwei
func gasPriceInGwei(price *big.Int) *big.Float {

	gwei := new(big.Float).Quo(new(big.Float).SetInt(price), big.NewFloat(1e9))
	return gwei

}
