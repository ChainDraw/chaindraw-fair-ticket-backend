package main

import (
	"chaindraw-fair-ticket-backend/go2chain/LotteryEscrowFactory"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var (
	WalletPrivateKey                     = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	AnvilRPC                             = "http://127.0.0.1:8545"
	AnvilChainID                         = 31337
	LotteryEscrowFactory_ContractAddress = common.HexToAddress("0x2de080e97b0cae9825375d31f5d0ed5751fdf16d")
	LotteryEscrow_ContractAddress        = common.HexToAddress("0x153327E2B1Df854f53fE1928C66165361b50Fcc1")
	SimulateAddress                      = common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")
	client, _                            = ethclient.Dial(AnvilRPC)
)

func main() {
	wLotteryEscrowFactory()
}

//func wLotteryEscrow() {
//	auth = contructTx(WalletPrivateKey, AnvilChainID)
//
//	lotteryEscrow, err := LotteryEscrow.NewLotteryEscrow(LotteryEscrow_ContractAddress, client)
//	if err != nil {
//		fmt.Println("NewLotteryEscrowFactory错误!")
//	}
//
//	LotteryEscrowFactory.
//}

func wLotteryEscrowFactory() {
	// 交易
	auth := contructTx(WalletPrivateKey, AnvilChainID)

	// 实例
	lotteryEscrowFactory, err := LotteryEscrowFactory.NewLotteryEscrowFactory(LotteryEscrowFactory_ContractAddress, client)
	if err != nil {
		fmt.Println("NewLotteryEscrowFactory错误!")
	}

	// 交互
	transaction, err := lotteryEscrowFactory.CreateEscrow(auth, SimulateAddress, "1", big.NewInt(int64(32)), "C级", "南看台", big.NewInt(int64(188)), "http://89Y45", big.NewInt(int64(3888)), big.NewInt(int64(1700000000)))
	if err != nil {
		fmt.Println("CreateEscrow发生错误!")
	}
	fmt.Println("交易哈希:", transaction.Hash().Hex())

}

func contructTx(WalletPrivateKey string, ChainID int) *bind.TransactOpts {

	privateKey, address := wallet(WalletPrivateKey)
	chainID := big.NewInt((int64(ChainID)))

	// nonce
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	nonce, _ := client.PendingNonceAt(context.Background(), address)
	auth.Nonce = big.NewInt(int64(nonce))

	// value
	auth.Value = big.NewInt(0) // in wei

	// gasLimit
	auth.GasLimit = uint64(3000000) // in units

	// gasPrice
	// auth.GasPrice = big.NewInt(1e9) // only for legacy transactions
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.GasFeeCap = gasPrice

	// gasTip
	gasTipCap, _ := client.SuggestGasTipCap(context.Background())
	auth.GasTipCap = gasTipCap

	return auth
}

func wallet(privateKey string) (*ecdsa.PrivateKey, common.Address) {
	privateKeyECDSA, _ := crypto.HexToECDSA(privateKey)
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKeyECDSA, address
}
