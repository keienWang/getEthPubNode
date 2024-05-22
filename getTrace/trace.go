package gettrace

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetTrace() {

	client, err := ethclient.Dial("http://192.168.110.146:8545")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(5000001))
	var txHashs []string
	for _, tx := range block.Transactions() {
		txHashs = append(txHashs, tx.Hash().String())
	}
	fmt.Println(time.Now().String())
	txscript, err := GetTransactions(txHashs)
	fmt.Println(txscript)
	fmt.Println(time.Now().String())

}

func GetTransactions(txHashs []string) ([]*types.Transaction, error) {

	client, err := ethclient.Dial("http://192.168.110.146:8545")
	if err != nil {
		log.Fatal(err)
	}

	name := "eth_getTransactionByHash"
	// 结果数组存储的是每个请求的结果指针，也就是引用
	rets := []*types.Transaction{}

	// 获取 hash 数组的长度，方便在循环中逐个实例化 BatchElem
	size := len(txHashs)

	reqs := []rpc.BatchElem{}

	for i := 0; i < size; i++ {
		ret := types.Transaction{}
		// 实例化每个 BatchElem
		req := rpc.BatchElem{
			Method: name,
			Args:   []interface{}{txHashs[i]},
			// &ret 传入单个请求的结果引用，这样是保证它在函数内部被修改值后，回到函数外来，值仍有效
			Result: &ret,
		}
		reqs = append(reqs, req)  // 将每个 BatchElem 添加到 BatchElem 数组
		rets = append(rets, &ret) // 每个请求的结果引用添加到结果数组中
	}
	err1 := client.Client().BatchCall(reqs) // 传入 BatchElem 数组，发起批量请求
	return rets, err1
}
