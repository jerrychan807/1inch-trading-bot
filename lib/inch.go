package lib

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jerrychan807/1inch-trading-bot/ethbasedclient"
	"github.com/jerrychan807/1inch-trading-bot/ethutils"
	go1inch "github.com/jon4hz/go-1inch"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

// @title 1inch路由器的合约地址
func GetApproveSpender(network string) string {
	client := go1inch.NewClient()
	res, _, err := client.ApproveSpender(context.Background(), go1inch.Network(network))
	if err != nil {
		fmt.Println(err)
	}
	Logger.WithFields(logrus.Fields{"1inchRouter": res.Address}).Info("1inchRouter Contract Address")
	return res.Address
}

// @title 查询授权的额度
func GetApproveAllowance(tokenAddr string, walletAddr string, network string) string {
	client := go1inch.NewClient()
	res, _, err := client.ApproveAllowance(context.Background(), go1inch.Network(network), tokenAddr, walletAddr)
	if err != nil {
		fmt.Println(err)
	}
	Logger.WithFields(logrus.Fields{"Allowance": res.Allowance}).Info("Allowance for 1inchRouter")
	return res.Allowance
}

// @title 返回授权给1inchRouter的tx数据
func GetApproveTx(tokenAddr string, network string) *go1inch.ApproveTransactionRes {
	client := go1inch.NewClient()
	res, _, err := client.ApproveTransaction(
		context.Background(),
		go1inch.Network(network),
		tokenAddr,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	Logger.WithFields(logrus.Fields{"ToAddr": res.To, "Data": res.Data}).Info("")
	return res
}

func GetSwapTxData(tokenAddr string, toTokenAddr string, amount string, walletAddr string, slippage int64, network string) *go1inch.SwapRes {
	client := go1inch.NewClient()

	res, _, err := client.Swap(context.Background(), go1inch.Network(network),
		tokenAddr,
		toTokenAddr,
		amount,
		walletAddr,
		slippage,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// @title 给1inch路由授权
func ApproveTokenByInch(ethBasedClient ethbasedclient.EthBasedClient, tokenAddr string, network string) error {
	Logger.WithFields(logrus.Fields{"walletAddr": ethBasedClient.Address}).Info("")
	nonce := ethBasedClient.PendingNonceUint64()
	approveRes := GetApproveTx(tokenAddr, network)

	gasLimit := uint64(210000)
	ToAddr := common.HexToAddress(approveRes.To)
	ethValueInt, _ := strconv.Atoi(approveRes.Value)
	ethValue := big.NewInt(int64(ethValueInt))
	data := common.FromHex(approveRes.Data)
	gasPriceInt, _ := strconv.Atoi(approveRes.GasPrice)
	gasPrice := big.NewInt(int64(gasPriceInt))
	chainID, _ := ethBasedClient.Client.NetworkID(context.Background())
	t := types.NewTransaction(nonce, ToAddr, ethValue, gasLimit, gasPrice, data)
	s := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(t, s, ethBasedClient.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Logger.Info(gasPriceErr)
	err = ethBasedClient.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		Logger.WithFields(logrus.Fields{"err": err, "chainID": chainID}).Info("SendTransaction,ApproveTokenByInch Faild")
		return err
	}
	txHash := signedTx.Hash().Hex()

	_, waitErr := bind.WaitMined(context.Background(), ethBasedClient.Client, signedTx)
	if waitErr != nil {
		Logger.WithFields(logrus.Fields{"approveTxHash": txHash}).Info("approveTxHash Faild")
		return err
	} else {
		Logger.WithFields(logrus.Fields{"approveTxHash": txHash}).Info("approveTxHash Successfully")
	}
	return nil
}

// @title 通过1inch交易
func SwapTokenByInch(ethBasedClient ethbasedclient.EthBasedClient, tokenAddr string, toTokenAddr string, amount string, walletAddr string, slippage int64, network string) error {
	Logger.WithFields(logrus.Fields{"walletAddr": ethBasedClient.Address}).Info("")
	nonce := ethBasedClient.PendingNonceUint64()

	swapRes := GetSwapTxData(tokenAddr, toTokenAddr, amount, walletAddr, slippage, network) // 获取swap交易参数
	ToAddr := common.HexToAddress(swapRes.Tx.To)
	gasLimit := uint64(swapRes.Tx.Gas)
	gasPriceInt, _ := strconv.Atoi(swapRes.Tx.GasPrice)
	gasPrice := big.NewInt(int64(gasPriceInt))
	ethValueInt, _ := strconv.Atoi(swapRes.Tx.Value)
	ethValue := big.NewInt(int64(ethValueInt))
	data := common.FromHex(swapRes.Tx.Data)
	chainID, _ := ethBasedClient.Client.NetworkID(context.Background())

	t := types.NewTransaction(nonce, ToAddr, ethValue, gasLimit, gasPrice, data)
	s := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(t, s, ethBasedClient.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ethBasedClient.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		Logger.WithFields(logrus.Fields{"err": err, "chainID": chainID}).Info("SendTransaction")
		return err
	}
	txHash := signedTx.Hash().Hex()

	_, waitErr := bind.WaitMined(context.Background(), ethBasedClient.Client, signedTx)
	if waitErr != nil {
		Logger.WithFields(logrus.Fields{"swapTxHash": txHash}).Info("swapTxHash Faild")
		return waitErr
	} else {
		Logger.WithFields(logrus.Fields{"swapTxHash": txHash}).Info("swapTxHash Successfully")
	}
	return nil
}

// @title 通过1inch交易
func BuyTokenByInch(ethBasedClient ethbasedclient.EthBasedClient, tokenAddr string, toTokenAddr string, amount float64, slippage int64, network string) error {
	usedTokenContractAddr := common.HexToAddress(tokenAddr)
	usedTokenIns := GetTokenInstance(usedTokenContractAddr, ethBasedClient.Client)
	Logger.Info("Try to Get TokenDecimals")
	usedTokenDecimals := GetTokenDecimals(usedTokenIns)
	amountInWei := ethutils.EtherToWeiByDecimal(big.NewFloat(amount), int(usedTokenDecimals))
	amountInWeiStr := amountInWei.String()
	Logger.WithFields(logrus.Fields{"amountInWei": amountInWei, "amountInWeiStr": amountInWeiStr, "usedTokenDecimals": usedTokenDecimals}).Info("amountInWei")

	// 检查授权
	allowance := GetApproveAllowance(tokenAddr, ethBasedClient.Address.String(), network)
	if allowance == "0" {
		approveErr := ApproveTokenByInch(ethBasedClient, tokenAddr, network)
		if approveErr != nil {
			return approveErr
		}
	}

	err := SwapTokenByInch(ethBasedClient, tokenAddr, toTokenAddr, amountInWeiStr, ethBasedClient.Address.String(), slippage, network)
	if err != nil {
		return err
	}
	return nil
}

func CheckBuyTokenStatus(ethBasedClient ethbasedclient.EthBasedClient, toTokenAddr string) bool {
	toTokenContractAddr := common.HexToAddress(toTokenAddr)

	toTokenIns := GetTokenInstance(toTokenContractAddr, ethBasedClient.Client)
	balance := GetAddrBalanceWei(toTokenIns, ethBasedClient.Address)
	if balance.Cmp(big.NewInt(0)) == 1 { // 余额大于0
		Logger.WithFields(logrus.Fields{"walletAddr": ethBasedClient.Address, "balance": balance}).Info("buyToken Successfully")
		return true
	} else {
		Logger.WithFields(logrus.Fields{"walletAddr": ethBasedClient.Address, "balance": balance}).Info("buyToken Faild")
		return false
	}
}

