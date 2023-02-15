package ethbasedclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jerrychan807/1inch-trading-bot/errorsutil"
	"log"
	"math/big"
)

type EthBasedClient struct {
	Client         *ethclient.Client
	PrivateKey     *ecdsa.PrivateKey
	PublicKeyECDSA *ecdsa.PublicKey
	Address        common.Address
	ChainID        *big.Int
	Transactor     *bind.TransactOpts
	Nonce          *big.Int
}

func New(rawurl, plainPrivateKey string) EthBasedClient {
	// connect with bsc
	client, err := ethclient.Dial(rawurl)
	errorsutil.HandleError(err)

	// create privateKey from string key
	privateKey, privateKeyErr := crypto.HexToECDSA(plainPrivateKey)
	errorsutil.HandleError(privateKeyErr)

	// generate public key and address from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// generate address from public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get chain id
	chainID, chainIDErr := client.ChainID(context.Background())
	errorsutil.HandleError(chainIDErr)

	// generate transactor for transactions management
	transactor, transactOptsErr := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	errorsutil.HandleError(transactOptsErr)

	ethBasedClientTemp := EthBasedClient{
		Client:         client,
		PrivateKey:     privateKey,
		PublicKeyECDSA: publicKeyECDSA,
		Address:        address,
		ChainID:        chainID,
		Transactor:     transactor,
	}

	return ethBasedClientTemp
}

func (ethBasedClient *EthBasedClient) ConfigureTransactor(value *big.Int, gasPrice *big.Int, gasLimit uint64) {

	if value.String() != "-1" {
		ethBasedClient.Transactor.Value = value
	}

	ethBasedClient.Transactor.GasPrice = gasPrice
	ethBasedClient.Transactor.GasLimit = gasLimit
	ethBasedClient.Transactor.Nonce = ethBasedClient.PendingNonce()
	ethBasedClient.Transactor.Context = context.Background()
}

func (ethBasedClient *EthBasedClient) Balance() *big.Int {
	// get current balance
	balance, balanceErr := ethBasedClient.Client.BalanceAt(context.Background(), ethBasedClient.Address, nil)
	errorsutil.HandleError(balanceErr)
	return balance
}

func (ethBasedClient *EthBasedClient) PendingNonce() *big.Int {
	// calculate next nonce
	nonce, nonceErr := ethBasedClient.Client.PendingNonceAt(context.Background(), ethBasedClient.Address)
	errorsutil.HandleError(nonceErr)
	return big.NewInt(int64(nonce))
}

func (ethBasedClient *EthBasedClient) PendingNonceUint64() uint64 {
	// calculate next nonce
	nonce, nonceErr := ethBasedClient.Client.PendingNonceAt(context.Background(), ethBasedClient.Address)
	errorsutil.HandleError(nonceErr)
	return nonce
}
