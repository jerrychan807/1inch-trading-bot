# 1inch-trading-bot
Trading bot on multichain using 1inch api      
This bot can be used to swap from One Crypto Asset to Another using 1inch Exchange

1inchApiV5 support chain list:

```go
Eth         Network = "eth"
Bsc         Network = "bsc"
Matic       Network = "matic"
Optimism    Network = "optimism"
Arbitrum    Network = "arbitrum"
GnosisChain Network = "gnosis"
Avalanche   Network = "avalanche"
Fantom      Network = "fantom"
Klaytn      Network = "klaytn"
Auror       Network = "auror"
```

# Usage:

1. Clone the repo
```sh
git clone https://github.com/jerrychan807/1inch-trading-bot.git
```
2. Install go packages
```sh
go mod tidy
```
3. Edit Swap paramters in main.go
- network: Network to use (matic, bsc, eth...)
- privateKey: Private Key of the account to use
- tokenAddr : Contract Address of Token to swap FROM
- toTokenAddress : Contract Address of Token to swap TO
- amount : Amount of Token to swap(Ethers)
- slippage : Slippage in percentage

Example Swap 0.5 USDC to USDT
```go
network := "matic"
rpcUrl := "https://polygon-rpc.com"
privateKey := "1ccxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx04"
tokenAddr := "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174" // USDC
toTokenAddr := "0xc2132D05D31c914a87C6611C10748AEb04B58e8F" // USDT
amount := 0.5
slippage := 1
```

4. run
```sh
go run main.go
```

Log:
```shell
time="2023-02-15T14:35:24+08:00" level=info msg="Try to Get TokenDecimals"
time="2023-02-15T14:35:25+08:00" level=info msg=amountInWei amountInWei=500000 amountInWeiStr=500000 usedTokenDecimals=6
time="2023-02-15T14:35:26+08:00" level=info msg="Allowance for 1inchRouter" Allowance=0
time="2023-02-15T14:35:26+08:00" level=info walletAddr=0x302c98exxxxxxxxxxxxxx45438C8
time="2023-02-15T14:35:27+08:00" level=info Data=0x095ea7b30000000000000000000000001111111254eeb25477b68fb85ed929f73a960582ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff ToAddr=0x2791bca1f2de4661ed88a30c99a7a9449aa84174
time="2023-02-15T14:35:34+08:00" level=info msg="approveTxHash Successfully" approveTxHash=0x4cb48264721926738bc847dbba2eb9e762ffba2178b6e4535af6596356df95c8
time="2023-02-15T14:35:34+08:00" level=info walletAddr=0x302cxxxxxxxxxxxxxxxxxxx45438C8
time="2023-02-15T14:35:42+08:00" level=info msg="swapTxHash Successfully" swapTxHash=0x9754ac23dbe6020157a68a270580b8929db8d791d2dc13d4d794b97bdf15c518
```

![20230215145054](https://raw.githubusercontent.com/jerrychan807/imggg/master/image/20230215145054.png)

# Refs:
- https://docs.1inch.io/docs/aggregation-protocol/api/swap-params
- https://github.com/jon4hz/go-1inch
- https://docs.1inch.io/docs/aggregation-protocol/api/swagger
- https://github.com/Mosiki/BuyTokensPancakeGolang