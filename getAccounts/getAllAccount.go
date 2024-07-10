package getaccounts

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccount() {

	client137, err := ethclient.Dial("http://192.168.110.157:8545")
	if err != nil {
		log.Fatal(err)
	}

	block3, _ := client137.BlockNumber(context.Background())
	fmt.Println("137block:", block3)

}
