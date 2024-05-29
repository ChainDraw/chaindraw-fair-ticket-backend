package go2chain

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
	"gorm.io/gorm"
	"log"
	"strings"
)

type DB func(data types.Log)

var (
	WSS           string
	ListenAddress []common.Address
	client        *ethclient.Client
	events        map[string]DB
)

func ListerInit(db *gorm.DB) {
	//dsn   = "root:12345678@tcp(127.0.0.1:3306)/chaindraw?charset=utf8mb4&parseTime=True&loc=Local&timeout=10000ms"
	WSS = "wss://go.getblock.io/74d1785308b244db9c9fda86104694c5" // 合约部署所在链的WSS  wss://go.getblock.io/74d1785308b244db9c9fda86104694c5
	// WSS           = "ws://127.0.0.1:8545" // 合约部署所在链的WSS
	ListenAddress = []common.Address{
		common.HexToAddress("0x683A3c225FFbAAC03F25Eab457edeB202cEBEd26"), //factory合约地址
		common.HexToAddress("0xD2BDf4F1F8f667d91809594cbbdCc7b23a160656"), // LotteryMarket合约地址
	} // 监听的合约地址
	client, _ = ethclient.Dial(WSS) // 客户端
	events = map[string]DB{
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
		crypto.Keccak256Hash([]byte("event_nft_listed(address,address,uint256,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventNftListed{
				Seller:         data.Topics[1].Hex(),
				LotteryAddress: data.Topics[2].Hex(),
				TokenId:        data.Topics[3].Hex(),
			}
			tmp := &struct {
				price string
			}{}
			getLotteryMarketABI().UnpackIntoInterface(tmp, "event_nft_listed", data.Data)
			event.Price = tmp.price
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("NFTSold(address,address,uint256,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventNftSold{
				Buyer:          data.Topics[1].Hex(),
				LotteryAddress: data.Topics[2].Hex(),
				TokenId:        data.Topics[3].Hex(),
			}
			tmp := &struct {
				price string
			}{}
			getLotteryMarketABI().UnpackIntoInterface(tmp, "event_nft_sold", data.Data)
			event.Price = tmp.price
			db.Save(event)
		}),
		crypto.Keccak256Hash([]byte("NFTDelisted(address,address,uint256)")).Hex(): DB(func(data types.Log) {
			event := &model.EventNftDelisted{
				Seller:         data.Topics[1].Hex(),
				LotteryAddress: data.Topics[2].Hex(),
				TokenId:        data.Topics[3].Hex(),
			}
			db.Save(event)
		}),
	}
}

func Run() {
	query := ethereum.FilterQuery{
		Addresses: ListenAddress,
	}
	logs := make(chan types.Log, 100)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			events[vLog.Topics[0].Hex()](vLog)
		}
	}

}

func getLotteryEscrowFactoryABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryEscrowFactory.LotteryEscrowFactoryMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	return abi
}

func getLotteryEscrowABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryEscrow.LotteryEscrowMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	return abi
}

func getLotteryMarketABI() abi.ABI {
	abi, err := abi.JSON(strings.NewReader(LotteryMarket.LotteryMarketMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	return abi
}
