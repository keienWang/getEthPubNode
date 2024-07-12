package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ConnectToEthereumNode 连接到以太坊节点
func ConnectToEthereumNode(nodeURL string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetAccountBalance 查询账户余额
func GetAccountBalance(client *ethclient.Client, address string) (*big.Float, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return etherValue, nil
}
