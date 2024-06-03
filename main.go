package main

import (
	getaccounts "get_eth_node/getAccounts"
	"time"
)

func main() {

	// trace.GetBlockTxTxR()
	for {
		getaccounts.GetAccount()
		time.Sleep(1 * time.Second)
	}

	// getpeers.GetPeers()

}
