package getaccounts

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccount() {

	client, err := ethclient.Dial("http://192.168.110.146:8545")
	if err != nil {
		log.Fatal(err)
	}

	hash, _ := client.TransactionReceipt(context.Background(), common.HexToHash("0x3f97c969ddf71f515ce5373b1f8e76e9fd7016611d8ce455881009414301789e"))

	fmt.Println(hash)

}
