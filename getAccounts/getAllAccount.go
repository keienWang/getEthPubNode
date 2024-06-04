package getaccounts

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccount() {

	client, err := ethclient.Dial("http://192.168.110.153:8545")
	if err != nil {
		log.Fatal(err)
	}

	block, _ := client.BlockNumber(context.Background())
	fmt.Println(block)

}
