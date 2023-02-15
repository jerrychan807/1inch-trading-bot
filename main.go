package main

import (
	"github.com/jerrychan807/1inch-trading-bot/ethbasedclient"
	"github.com/jerrychan807/1inch-trading-bot/lib"
)

func main() {
	network := "matic"
	rpcUrl := "https://polygon-rpc.com"
	privateKey := "1cc4b74xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0af04"
	tokenAddr := "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174" // USDC
	toTokenAddr := "0xc2132D05D31c914a87C6611C10748AEb04B58e8F" // USDT
	amount := 0.5
	slippage := 1
	ethBasedClient := ethbasedclient.New(rpcUrl, privateKey)
	lib.BuyTokenByInch(ethBasedClient, tokenAddr, toTokenAddr, amount, int64(slippage), network)

}
