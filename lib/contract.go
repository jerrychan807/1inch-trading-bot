package lib

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/jerrychan807/1inch-trading-bot/contracts/ERC20"
	"github.com/jerrychan807/1inch-trading-bot/util"
	"log"
	"math/big"
)

// Token Token信息
type Token struct {
	ContractAddr string // Token合约地址
	Name         string // Token名字
	Symbol       string // Token符号
	Decimals     uint8  // Token小数位
}

// @title 获得Rpc客户端实例
func getRpcClient(rpcUrl string) *ethclient.Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// @title 获取erc20Token的一个contract实例
// @param client *ethclient.Client "rpc连接client"
func GetTokenInstance(contractAddr common.Address, client *ethclient.Client) *erc20.Erc20 {
	instance, err := erc20.NewErc20(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

// @title 查询erc20Token的代币名称
// @param instance *erc20.Erc20 "erc20实例"
func GetTokenName(instance *erc20.Erc20) string {
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return name
}

// @title 查询erc20Token的代币符号
// @param instance *erc20.Erc20 "erc20实例"
func GetTokenSymbol(instance *erc20.Erc20) string {
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return symbol
}

// @title 查询erc20Token的精度/小数点
// @param instance *erc20.Erc20 "erc20实例"
func GetTokenDecimals(instance *erc20.Erc20) uint8 {
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return decimals
}

// @title 查询某地址的erc20Token余额
// @param instance *erc20.Erc20 "erc20实例"
func GetAddrBalance(instance *erc20.Erc20, Addr common.Address) string {
	balanceWei, err := instance.BalanceOf(&bind.CallOpts{}, Addr)
	if err != nil {
		log.Fatal(err)
	}

	balance := util.WeiToEther(balanceWei)
	//fmt.Printf("[*] balance: %s \n", balance)
	return balance
}

func GetAddrBalanceWei(instance *erc20.Erc20, Addr common.Address) *big.Int {
	balanceWei, err := instance.BalanceOf(&bind.CallOpts{}, Addr)
	if err != nil {
		log.Fatal(err)
	}
	return balanceWei
}
