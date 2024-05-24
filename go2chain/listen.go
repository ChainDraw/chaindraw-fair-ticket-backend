package main

import (
	"chaindraw-fair-ticket-backend/go2chain/LotteryEscrow"
	"chaindraw-fair-ticket-backend/go2chain/LotteryEscrowFactory"
	"chaindraw-fair-ticket-backend/go2chain/LotteryMarket"
	model "chaindraw-fair-ticket-backend/model/event"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

type DB func(data types.Log)

var (
	dsn   = "root:12345678@tcp(127.0.0.1:3306)/chaindraw?charset=utf8mb4&parseTime=True&loc=Local&timeout=10000ms"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	WSS           = "ws://127.0.0.1:8545" // 合约部署所在链的WSS地址
	ListenAddress = []common.Address{
		common.HexToAddress("0x8464135c8f25da09e49bc8782676a84730c318bc"),
	} // 监听的合约
	client, _ = ethclient.Dial(WSS) // 客户端
	events    = map[string]DB{
		crypto.Keccak256Hash([]byte("EscrowCreated(uint256,uint256,address)")).Hex(): DB(func(data types.Log) {
			event := &model.EventEscrowCreated{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				EscrowAddress common.Address
			}{}
			getLotteryEscrowFactoryABI().UnpackIntoInterface(tmp, "EscrowCreated", data.Data)
			event.EscrowAddress = tmp.EscrowAddress.Hex()
			ListenAddress = append(ListenAddress, tmp.EscrowAddress)
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__Deposited(uint256,uint256,address,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowDeposited{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				Buyer common.Address
				Money common.Address
			}{}
			getLotteryEscrowABI().UnpackIntoInterface(tmp, "LotteryEscrow__Deposited", data.Data)
			event.Buyer = tmp.Buyer.Hex()
			event.Money = tmp.Money.Hex()
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__Refunded(uint256,uint256,address,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowRefunded{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				Buyer common.Address
				Money common.Address
			}{}
			getLotteryEscrowABI().UnpackIntoInterface(tmp, "LotteryEscrow__Refunded", data.Data)
			event.Buyer = tmp.Buyer.Hex()
			event.Money = tmp.Money.Hex()
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__ClaimedFund(uint256,uint256,address,address,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowClaimedfund{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				Organizer common.Address
				Winner    common.Address
				Money     string
			}{}
			getLotteryEscrowABI().UnpackIntoInterface(tmp, "LotteryEscrow__ClaimedFund", data.Data)
			event.Organizer = tmp.Organizer.Hex()
			event.Winner = tmp.Winner.Hex()
			event.Money = tmp.Money
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__NonWinner(uint256,uint256,address,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowNonwinner{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				NonWinner common.Address
				Money     string
			}{}
			getLotteryEscrowABI().UnpackIntoInterface(tmp, "LotteryEscrow__NonWinner", data.Data)
			event.NonWinner = tmp.NonWinner.Hex()
			event.Money = tmp.Money
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__Winner(uint256,uint256,address)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowWinner{
				ConcertId:  data.Topics[1].Hex(),
				TicketType: data.Topics[2].Hex(),
			}
			tmp := &struct {
				Winner common.Address
			}{}
			getLotteryEscrowABI().UnpackIntoInterface(tmp, "LotteryEscrow__Winner", data.Data)
			event.Winner = tmp.Winner.Hex()
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("LotteryEscrow__CompleteDraw(address)")).Hex(): DB(func(data types.Log) {
			event := &model.EventLotteryescrowCompletedraw{}
			getLotteryEscrowABI().UnpackIntoInterface(event, "LotteryEscrow__CompleteDraw", data.Data)
			db.Save(event)
		}),
	}
)

// TODO 1.chan 设置chan缓存容量 2.Log 记录监听日志
func main() {
	query := ethereum.FilterQuery{
		Addresses: ListenAddress,
	}
	logs := make(chan types.Log) //TODO chan
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err) // TODO Log
	}
	for {
		select {
		case err := <-sub.Err(): // TODO Log
			log.Fatal(err)
		case vLog := <-logs:
			events[vLog.Topics[0].Hex()](vLog)
		}
	}
}

func getLotteryEscrowFactoryABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryEscrowFactory.LotteryEscrowFactoryMetaData.ABI))
	if err != nil {
		log.Fatal(err) // TODO Log
	}
	return abi
}

func getLotteryEscrowABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryEscrow.LotteryEscrowMetaData.ABI))
	if err != nil {
		log.Fatal(err) // TODO Log
	}
	return abi
}

func getLotteryMarketABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryMarket.LotteryMarketMetaData.ABI))
	if err != nil {
		log.Fatal(err) // TODO Log
	}
	return abi
}
