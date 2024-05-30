// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LotteryEscrow

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LotteryEscrowMetaData contains all meta data concerning the LotteryEscrow contract.
var LotteryEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_organizer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_concertId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ticketType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ticketCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ddl\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LotteryEscrowError__DepositTimeOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LotteryEscrowError__alreadyJoin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ChainlinkVrf__RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"ChainlinkVrf__RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"organizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"LotteryEscrow__ClaimedFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"}],\"name\":\"LotteryEscrow__CompleteDraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"LotteryEscrow__Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nonWinner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"LotteryEscrow__NonWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"LotteryEscrow__Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"LotteryEscrow__Winner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allBuyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeDraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"concertId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ddl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWinner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingTicketCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startLottery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"url\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LotteryEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryEscrowMetaData.ABI instead.
var LotteryEscrowABI = LotteryEscrowMetaData.ABI

// LotteryEscrow is an auto generated Go binding around an Ethereum contract.
type LotteryEscrow struct {
	LotteryEscrowCaller     // Read-only binding to the contract
	LotteryEscrowTransactor // Write-only binding to the contract
	LotteryEscrowFilterer   // Log filterer for contract events
}

// LotteryEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotteryEscrowSession struct {
	Contract     *LotteryEscrow    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryEscrowCallerSession struct {
	Contract *LotteryEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LotteryEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryEscrowTransactorSession struct {
	Contract     *LotteryEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LotteryEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryEscrowRaw struct {
	Contract *LotteryEscrow // Generic contract binding to access the raw methods on
}

// LotteryEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryEscrowCallerRaw struct {
	Contract *LotteryEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryEscrowTransactorRaw struct {
	Contract *LotteryEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryEscrow creates a new instance of LotteryEscrow, bound to a specific deployed contract.
func NewLotteryEscrow(address common.Address, backend bind.ContractBackend) (*LotteryEscrow, error) {
	contract, err := bindLotteryEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrow{LotteryEscrowCaller: LotteryEscrowCaller{contract: contract}, LotteryEscrowTransactor: LotteryEscrowTransactor{contract: contract}, LotteryEscrowFilterer: LotteryEscrowFilterer{contract: contract}}, nil
}

// NewLotteryEscrowCaller creates a new read-only instance of LotteryEscrow, bound to a specific deployed contract.
func NewLotteryEscrowCaller(address common.Address, caller bind.ContractCaller) (*LotteryEscrowCaller, error) {
	contract, err := bindLotteryEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowCaller{contract: contract}, nil
}

// NewLotteryEscrowTransactor creates a new write-only instance of LotteryEscrow, bound to a specific deployed contract.
func NewLotteryEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryEscrowTransactor, error) {
	contract, err := bindLotteryEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowTransactor{contract: contract}, nil
}

// NewLotteryEscrowFilterer creates a new log filterer instance of LotteryEscrow, bound to a specific deployed contract.
func NewLotteryEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryEscrowFilterer, error) {
	contract, err := bindLotteryEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFilterer{contract: contract}, nil
}

// bindLotteryEscrow binds a generic wrapper to an already deployed contract.
func bindLotteryEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryEscrow *LotteryEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryEscrow.Contract.LotteryEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryEscrow *LotteryEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.LotteryEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryEscrow *LotteryEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.LotteryEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryEscrow *LotteryEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryEscrow *LotteryEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryEscrow *LotteryEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) Factory() (common.Address, error) {
	return _LotteryEscrow.Contract.Factory(&_LotteryEscrow.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) Factory() (common.Address, error) {
	return _LotteryEscrow.Contract.Factory(&_LotteryEscrow.CallOpts)
}

// AllBuyer is a free data retrieval call binding the contract method 0x5476d0de.
//
// Solidity: function allBuyer(uint256 ) view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) AllBuyer(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "allBuyer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllBuyer is a free data retrieval call binding the contract method 0x5476d0de.
//
// Solidity: function allBuyer(uint256 ) view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) AllBuyer(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.AllBuyer(&_LotteryEscrow.CallOpts, arg0)
}

// AllBuyer is a free data retrieval call binding the contract method 0x5476d0de.
//
// Solidity: function allBuyer(uint256 ) view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) AllBuyer(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.AllBuyer(&_LotteryEscrow.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LotteryEscrow.Contract.BalanceOf(&_LotteryEscrow.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LotteryEscrow.Contract.BalanceOf(&_LotteryEscrow.CallOpts, owner)
}

// CompleteDraw is a free data retrieval call binding the contract method 0x0bdeecbd.
//
// Solidity: function completeDraw() view returns(bool)
func (_LotteryEscrow *LotteryEscrowCaller) CompleteDraw(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "completeDraw")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompleteDraw is a free data retrieval call binding the contract method 0x0bdeecbd.
//
// Solidity: function completeDraw() view returns(bool)
func (_LotteryEscrow *LotteryEscrowSession) CompleteDraw() (bool, error) {
	return _LotteryEscrow.Contract.CompleteDraw(&_LotteryEscrow.CallOpts)
}

// CompleteDraw is a free data retrieval call binding the contract method 0x0bdeecbd.
//
// Solidity: function completeDraw() view returns(bool)
func (_LotteryEscrow *LotteryEscrowCallerSession) CompleteDraw() (bool, error) {
	return _LotteryEscrow.Contract.CompleteDraw(&_LotteryEscrow.CallOpts)
}

// ConcertId is a free data retrieval call binding the contract method 0xbd73f75f.
//
// Solidity: function concertId() view returns(string)
func (_LotteryEscrow *LotteryEscrowCaller) ConcertId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "concertId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ConcertId is a free data retrieval call binding the contract method 0xbd73f75f.
//
// Solidity: function concertId() view returns(string)
func (_LotteryEscrow *LotteryEscrowSession) ConcertId() (string, error) {
	return _LotteryEscrow.Contract.ConcertId(&_LotteryEscrow.CallOpts)
}

// ConcertId is a free data retrieval call binding the contract method 0xbd73f75f.
//
// Solidity: function concertId() view returns(string)
func (_LotteryEscrow *LotteryEscrowCallerSession) ConcertId() (string, error) {
	return _LotteryEscrow.Contract.ConcertId(&_LotteryEscrow.CallOpts)
}

// Ddl is a free data retrieval call binding the contract method 0xf70e97b7.
//
// Solidity: function ddl() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) Ddl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "ddl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ddl is a free data retrieval call binding the contract method 0xf70e97b7.
//
// Solidity: function ddl() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) Ddl() (*big.Int, error) {
	return _LotteryEscrow.Contract.Ddl(&_LotteryEscrow.CallOpts)
}

// Ddl is a free data retrieval call binding the contract method 0xf70e97b7.
//
// Solidity: function ddl() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) Ddl() (*big.Int, error) {
	return _LotteryEscrow.Contract.Ddl(&_LotteryEscrow.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _LotteryEscrow.Contract.Deposits(&_LotteryEscrow.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _LotteryEscrow.Contract.Deposits(&_LotteryEscrow.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.GetApproved(&_LotteryEscrow.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.GetApproved(&_LotteryEscrow.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryEscrow *LotteryEscrowSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LotteryEscrow.Contract.IsApprovedForAll(&_LotteryEscrow.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LotteryEscrow.Contract.IsApprovedForAll(&_LotteryEscrow.CallOpts, owner, operator)
}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address ) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCaller) IsWinner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "isWinner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address ) view returns(bool)
func (_LotteryEscrow *LotteryEscrowSession) IsWinner(arg0 common.Address) (bool, error) {
	return _LotteryEscrow.Contract.IsWinner(&_LotteryEscrow.CallOpts, arg0)
}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address ) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCallerSession) IsWinner(arg0 common.Address) (bool, error) {
	return _LotteryEscrow.Contract.IsWinner(&_LotteryEscrow.CallOpts, arg0)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) LastRequestId() (*big.Int, error) {
	return _LotteryEscrow.Contract.LastRequestId(&_LotteryEscrow.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) LastRequestId() (*big.Int, error) {
	return _LotteryEscrow.Contract.LastRequestId(&_LotteryEscrow.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "linkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) LinkToken() (common.Address, error) {
	return _LotteryEscrow.Contract.LinkToken(&_LotteryEscrow.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) LinkToken() (common.Address, error) {
	return _LotteryEscrow.Contract.LinkToken(&_LotteryEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryEscrow *LotteryEscrowCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryEscrow *LotteryEscrowSession) Name() (string, error) {
	return _LotteryEscrow.Contract.Name(&_LotteryEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryEscrow *LotteryEscrowCallerSession) Name() (string, error) {
	return _LotteryEscrow.Contract.Name(&_LotteryEscrow.CallOpts)
}

// Organizer is a free data retrieval call binding the contract method 0x61203265.
//
// Solidity: function organizer() view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) Organizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "organizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Organizer is a free data retrieval call binding the contract method 0x61203265.
//
// Solidity: function organizer() view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) Organizer() (common.Address, error) {
	return _LotteryEscrow.Contract.Organizer(&_LotteryEscrow.CallOpts)
}

// Organizer is a free data retrieval call binding the contract method 0x61203265.
//
// Solidity: function organizer() view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) Organizer() (common.Address, error) {
	return _LotteryEscrow.Contract.Organizer(&_LotteryEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) Owner() (common.Address, error) {
	return _LotteryEscrow.Contract.Owner(&_LotteryEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) Owner() (common.Address, error) {
	return _LotteryEscrow.Contract.Owner(&_LotteryEscrow.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.OwnerOf(&_LotteryEscrow.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryEscrow *LotteryEscrowCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LotteryEscrow.Contract.OwnerOf(&_LotteryEscrow.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) Price() (*big.Int, error) {
	return _LotteryEscrow.Contract.Price(&_LotteryEscrow.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) Price() (*big.Int, error) {
	return _LotteryEscrow.Contract.Price(&_LotteryEscrow.CallOpts)
}

// RemainingTicketCount is a free data retrieval call binding the contract method 0x00099854.
//
// Solidity: function remainingTicketCount() view returns(uint32)
func (_LotteryEscrow *LotteryEscrowCaller) RemainingTicketCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "remainingTicketCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RemainingTicketCount is a free data retrieval call binding the contract method 0x00099854.
//
// Solidity: function remainingTicketCount() view returns(uint32)
func (_LotteryEscrow *LotteryEscrowSession) RemainingTicketCount() (uint32, error) {
	return _LotteryEscrow.Contract.RemainingTicketCount(&_LotteryEscrow.CallOpts)
}

// RemainingTicketCount is a free data retrieval call binding the contract method 0x00099854.
//
// Solidity: function remainingTicketCount() view returns(uint32)
func (_LotteryEscrow *LotteryEscrowCallerSession) RemainingTicketCount() (uint32, error) {
	return _LotteryEscrow.Contract.RemainingTicketCount(&_LotteryEscrow.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _LotteryEscrow.Contract.RequestIds(&_LotteryEscrow.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _LotteryEscrow.Contract.RequestIds(&_LotteryEscrow.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(uint256 paid, bool fulfilled)
func (_LotteryEscrow *LotteryEscrowCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Paid      *big.Int
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(struct {
		Paid      *big.Int
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(uint256 paid, bool fulfilled)
func (_LotteryEscrow *LotteryEscrowSession) SRequests(arg0 *big.Int) (struct {
	Paid      *big.Int
	Fulfilled bool
}, error) {
	return _LotteryEscrow.Contract.SRequests(&_LotteryEscrow.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(uint256 paid, bool fulfilled)
func (_LotteryEscrow *LotteryEscrowCallerSession) SRequests(arg0 *big.Int) (struct {
	Paid      *big.Int
	Fulfilled bool
}, error) {
	return _LotteryEscrow.Contract.SRequests(&_LotteryEscrow.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryEscrow *LotteryEscrowSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LotteryEscrow.Contract.SupportsInterface(&_LotteryEscrow.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryEscrow *LotteryEscrowCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LotteryEscrow.Contract.SupportsInterface(&_LotteryEscrow.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryEscrow *LotteryEscrowCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryEscrow *LotteryEscrowSession) Symbol() (string, error) {
	return _LotteryEscrow.Contract.Symbol(&_LotteryEscrow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryEscrow *LotteryEscrowCallerSession) Symbol() (string, error) {
	return _LotteryEscrow.Contract.Symbol(&_LotteryEscrow.CallOpts)
}

// TicketCount is a free data retrieval call binding the contract method 0xcfbd900f.
//
// Solidity: function ticketCount() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) TicketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "ticketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketCount is a free data retrieval call binding the contract method 0xcfbd900f.
//
// Solidity: function ticketCount() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) TicketCount() (*big.Int, error) {
	return _LotteryEscrow.Contract.TicketCount(&_LotteryEscrow.CallOpts)
}

// TicketCount is a free data retrieval call binding the contract method 0xcfbd900f.
//
// Solidity: function ticketCount() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) TicketCount() (*big.Int, error) {
	return _LotteryEscrow.Contract.TicketCount(&_LotteryEscrow.CallOpts)
}

// TicketType is a free data retrieval call binding the contract method 0x0c737add.
//
// Solidity: function ticketType() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCaller) TicketType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "ticketType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketType is a free data retrieval call binding the contract method 0x0c737add.
//
// Solidity: function ticketType() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowSession) TicketType() (*big.Int, error) {
	return _LotteryEscrow.Contract.TicketType(&_LotteryEscrow.CallOpts)
}

// TicketType is a free data retrieval call binding the contract method 0x0c737add.
//
// Solidity: function ticketType() view returns(uint256)
func (_LotteryEscrow *LotteryEscrowCallerSession) TicketType() (*big.Int, error) {
	return _LotteryEscrow.Contract.TicketType(&_LotteryEscrow.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryEscrow *LotteryEscrowCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryEscrow *LotteryEscrowSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LotteryEscrow.Contract.TokenURI(&_LotteryEscrow.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryEscrow *LotteryEscrowCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LotteryEscrow.Contract.TokenURI(&_LotteryEscrow.CallOpts, tokenId)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_LotteryEscrow *LotteryEscrowCaller) Url(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryEscrow.contract.Call(opts, &out, "url")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_LotteryEscrow *LotteryEscrowSession) Url() (string, error) {
	return _LotteryEscrow.Contract.Url(&_LotteryEscrow.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_LotteryEscrow *LotteryEscrowCallerSession) Url() (string, error) {
	return _LotteryEscrow.Contract.Url(&_LotteryEscrow.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LotteryEscrow *LotteryEscrowTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LotteryEscrow *LotteryEscrowSession) AcceptOwnership() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.AcceptOwnership(&_LotteryEscrow.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.AcceptOwnership(&_LotteryEscrow.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Approve(&_LotteryEscrow.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Approve(&_LotteryEscrow.TransactOpts, to, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LotteryEscrow *LotteryEscrowTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LotteryEscrow *LotteryEscrowSession) Deposit() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Deposit(&_LotteryEscrow.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) Deposit() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Deposit(&_LotteryEscrow.TransactOpts)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "rawFulfillRandomWords", _requestId, _randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_LotteryEscrow *LotteryEscrowSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.RawFulfillRandomWords(&_LotteryEscrow.TransactOpts, _requestId, _randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.RawFulfillRandomWords(&_LotteryEscrow.TransactOpts, _requestId, _randomWords)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address participant) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) Refund(opts *bind.TransactOpts, participant common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "refund", participant)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address participant) returns()
func (_LotteryEscrow *LotteryEscrowSession) Refund(participant common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Refund(&_LotteryEscrow.TransactOpts, participant)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address participant) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) Refund(participant common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Refund(&_LotteryEscrow.TransactOpts, participant)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SafeTransferFrom(&_LotteryEscrow.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SafeTransferFrom(&_LotteryEscrow.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryEscrow *LotteryEscrowSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SafeTransferFrom0(&_LotteryEscrow.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SafeTransferFrom0(&_LotteryEscrow.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryEscrow *LotteryEscrowSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SetApprovalForAll(&_LotteryEscrow.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.SetApprovalForAll(&_LotteryEscrow.TransactOpts, operator, approved)
}

// StartLottery is a paid mutator transaction binding the contract method 0x160344e2.
//
// Solidity: function startLottery() returns()
func (_LotteryEscrow *LotteryEscrowTransactor) StartLottery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "startLottery")
}

// StartLottery is a paid mutator transaction binding the contract method 0x160344e2.
//
// Solidity: function startLottery() returns()
func (_LotteryEscrow *LotteryEscrowSession) StartLottery() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.StartLottery(&_LotteryEscrow.TransactOpts)
}

// StartLottery is a paid mutator transaction binding the contract method 0x160344e2.
//
// Solidity: function startLottery() returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) StartLottery() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.StartLottery(&_LotteryEscrow.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.TransferFrom(&_LotteryEscrow.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.TransferFrom(&_LotteryEscrow.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_LotteryEscrow *LotteryEscrowSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.TransferOwnership(&_LotteryEscrow.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.TransferOwnership(&_LotteryEscrow.TransactOpts, to)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x7a8042bd.
//
// Solidity: function withdrawLink(uint256 _amount) returns()
func (_LotteryEscrow *LotteryEscrowTransactor) WithdrawLink(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.contract.Transact(opts, "withdrawLink", _amount)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x7a8042bd.
//
// Solidity: function withdrawLink(uint256 _amount) returns()
func (_LotteryEscrow *LotteryEscrowSession) WithdrawLink(_amount *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.WithdrawLink(&_LotteryEscrow.TransactOpts, _amount)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x7a8042bd.
//
// Solidity: function withdrawLink(uint256 _amount) returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) WithdrawLink(_amount *big.Int) (*types.Transaction, error) {
	return _LotteryEscrow.Contract.WithdrawLink(&_LotteryEscrow.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrow *LotteryEscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrow *LotteryEscrowSession) Receive() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Receive(&_LotteryEscrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrow *LotteryEscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _LotteryEscrow.Contract.Receive(&_LotteryEscrow.TransactOpts)
}

// LotteryEscrowApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LotteryEscrow contract.
type LotteryEscrowApprovalIterator struct {
	Event *LotteryEscrowApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowApproval represents a Approval event raised by the LotteryEscrow contract.
type LotteryEscrowApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LotteryEscrowApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowApprovalIterator{contract: _LotteryEscrow.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LotteryEscrowApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowApproval)
				if err := _LotteryEscrow.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseApproval(log types.Log) (*LotteryEscrowApproval, error) {
	event := new(LotteryEscrowApproval)
	if err := _LotteryEscrow.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LotteryEscrow contract.
type LotteryEscrowApprovalForAllIterator struct {
	Event *LotteryEscrowApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowApprovalForAll represents a ApprovalForAll event raised by the LotteryEscrow contract.
type LotteryEscrowApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LotteryEscrowApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowApprovalForAllIterator{contract: _LotteryEscrow.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LotteryEscrowApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowApprovalForAll)
				if err := _LotteryEscrow.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseApprovalForAll(log types.Log) (*LotteryEscrowApprovalForAll, error) {
	event := new(LotteryEscrowApprovalForAll)
	if err := _LotteryEscrow.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the LotteryEscrow contract.
type LotteryEscrowBatchMetadataUpdateIterator struct {
	Event *LotteryEscrowBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the LotteryEscrow contract.
type LotteryEscrowBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*LotteryEscrowBatchMetadataUpdateIterator, error) {

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowBatchMetadataUpdateIterator{contract: _LotteryEscrow.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LotteryEscrowBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowBatchMetadataUpdate)
				if err := _LotteryEscrow.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseBatchMetadataUpdate(log types.Log) (*LotteryEscrowBatchMetadataUpdate, error) {
	event := new(LotteryEscrowBatchMetadataUpdate)
	if err := _LotteryEscrow.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowChainlinkVrfRequestFulfilledIterator is returned from FilterChainlinkVrfRequestFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkVrfRequestFulfilled events raised by the LotteryEscrow contract.
type LotteryEscrowChainlinkVrfRequestFulfilledIterator struct {
	Event *LotteryEscrowChainlinkVrfRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowChainlinkVrfRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowChainlinkVrfRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowChainlinkVrfRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowChainlinkVrfRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowChainlinkVrfRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowChainlinkVrfRequestFulfilled represents a ChainlinkVrfRequestFulfilled event raised by the LotteryEscrow contract.
type LotteryEscrowChainlinkVrfRequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Payment     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChainlinkVrfRequestFulfilled is a free log retrieval operation binding the contract event 0x85c29d2ad0f1a812f307411c654a2fbfdd3304c770203c84643b0a726f491eea.
//
// Solidity: event ChainlinkVrf__RequestFulfilled(uint256 requestId, uint256[] randomWords, uint256 payment)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterChainlinkVrfRequestFulfilled(opts *bind.FilterOpts) (*LotteryEscrowChainlinkVrfRequestFulfilledIterator, error) {

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "ChainlinkVrf__RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowChainlinkVrfRequestFulfilledIterator{contract: _LotteryEscrow.contract, event: "ChainlinkVrf__RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkVrfRequestFulfilled is a free log subscription operation binding the contract event 0x85c29d2ad0f1a812f307411c654a2fbfdd3304c770203c84643b0a726f491eea.
//
// Solidity: event ChainlinkVrf__RequestFulfilled(uint256 requestId, uint256[] randomWords, uint256 payment)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchChainlinkVrfRequestFulfilled(opts *bind.WatchOpts, sink chan<- *LotteryEscrowChainlinkVrfRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "ChainlinkVrf__RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowChainlinkVrfRequestFulfilled)
				if err := _LotteryEscrow.contract.UnpackLog(event, "ChainlinkVrf__RequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainlinkVrfRequestFulfilled is a log parse operation binding the contract event 0x85c29d2ad0f1a812f307411c654a2fbfdd3304c770203c84643b0a726f491eea.
//
// Solidity: event ChainlinkVrf__RequestFulfilled(uint256 requestId, uint256[] randomWords, uint256 payment)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseChainlinkVrfRequestFulfilled(log types.Log) (*LotteryEscrowChainlinkVrfRequestFulfilled, error) {
	event := new(LotteryEscrowChainlinkVrfRequestFulfilled)
	if err := _LotteryEscrow.contract.UnpackLog(event, "ChainlinkVrf__RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowChainlinkVrfRequestSentIterator is returned from FilterChainlinkVrfRequestSent and is used to iterate over the raw logs and unpacked data for ChainlinkVrfRequestSent events raised by the LotteryEscrow contract.
type LotteryEscrowChainlinkVrfRequestSentIterator struct {
	Event *LotteryEscrowChainlinkVrfRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowChainlinkVrfRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowChainlinkVrfRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowChainlinkVrfRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowChainlinkVrfRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowChainlinkVrfRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowChainlinkVrfRequestSent represents a ChainlinkVrfRequestSent event raised by the LotteryEscrow contract.
type LotteryEscrowChainlinkVrfRequestSent struct {
	RequestId *big.Int
	NumWords  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainlinkVrfRequestSent is a free log retrieval operation binding the contract event 0x4461b2777889df5916f85d2cecf8eafcc1bb8f0ecee41ddca4b9497a07ad4c96.
//
// Solidity: event ChainlinkVrf__RequestSent(uint256 requestId, uint32 numWords)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterChainlinkVrfRequestSent(opts *bind.FilterOpts) (*LotteryEscrowChainlinkVrfRequestSentIterator, error) {

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "ChainlinkVrf__RequestSent")
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowChainlinkVrfRequestSentIterator{contract: _LotteryEscrow.contract, event: "ChainlinkVrf__RequestSent", logs: logs, sub: sub}, nil
}

// WatchChainlinkVrfRequestSent is a free log subscription operation binding the contract event 0x4461b2777889df5916f85d2cecf8eafcc1bb8f0ecee41ddca4b9497a07ad4c96.
//
// Solidity: event ChainlinkVrf__RequestSent(uint256 requestId, uint32 numWords)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchChainlinkVrfRequestSent(opts *bind.WatchOpts, sink chan<- *LotteryEscrowChainlinkVrfRequestSent) (event.Subscription, error) {

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "ChainlinkVrf__RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowChainlinkVrfRequestSent)
				if err := _LotteryEscrow.contract.UnpackLog(event, "ChainlinkVrf__RequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainlinkVrfRequestSent is a log parse operation binding the contract event 0x4461b2777889df5916f85d2cecf8eafcc1bb8f0ecee41ddca4b9497a07ad4c96.
//
// Solidity: event ChainlinkVrf__RequestSent(uint256 requestId, uint32 numWords)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseChainlinkVrfRequestSent(log types.Log) (*LotteryEscrowChainlinkVrfRequestSent, error) {
	event := new(LotteryEscrowChainlinkVrfRequestSent)
	if err := _LotteryEscrow.contract.UnpackLog(event, "ChainlinkVrf__RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowClaimedFundIterator is returned from FilterLotteryEscrowClaimedFund and is used to iterate over the raw logs and unpacked data for LotteryEscrowClaimedFund events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowClaimedFundIterator struct {
	Event *LotteryEscrowLotteryEscrowClaimedFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowClaimedFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowClaimedFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowClaimedFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowClaimedFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowClaimedFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowClaimedFund represents a LotteryEscrowClaimedFund event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowClaimedFund struct {
	ConcertId  string
	TicketType *big.Int
	Organizer  common.Address
	Winner     common.Address
	Money      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowClaimedFund is a free log retrieval operation binding the contract event 0x55afc2b602fdf10b8b9c5ec43d0c93e0abf8d0486fe82aa9ec6d599591876610.
//
// Solidity: event LotteryEscrow__ClaimedFund(string concertId, uint256 indexed ticketType, address organizer, address winner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowClaimedFund(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowLotteryEscrowClaimedFundIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__ClaimedFund", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowClaimedFundIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__ClaimedFund", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowClaimedFund is a free log subscription operation binding the contract event 0x55afc2b602fdf10b8b9c5ec43d0c93e0abf8d0486fe82aa9ec6d599591876610.
//
// Solidity: event LotteryEscrow__ClaimedFund(string concertId, uint256 indexed ticketType, address organizer, address winner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowClaimedFund(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowClaimedFund, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__ClaimedFund", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowClaimedFund)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__ClaimedFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowClaimedFund is a log parse operation binding the contract event 0x55afc2b602fdf10b8b9c5ec43d0c93e0abf8d0486fe82aa9ec6d599591876610.
//
// Solidity: event LotteryEscrow__ClaimedFund(string concertId, uint256 indexed ticketType, address organizer, address winner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowClaimedFund(log types.Log) (*LotteryEscrowLotteryEscrowClaimedFund, error) {
	event := new(LotteryEscrowLotteryEscrowClaimedFund)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__ClaimedFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowCompleteDrawIterator is returned from FilterLotteryEscrowCompleteDraw and is used to iterate over the raw logs and unpacked data for LotteryEscrowCompleteDraw events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowCompleteDrawIterator struct {
	Event *LotteryEscrowLotteryEscrowCompleteDraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowCompleteDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowCompleteDraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowCompleteDraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowCompleteDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowCompleteDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowCompleteDraw represents a LotteryEscrowCompleteDraw event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowCompleteDraw struct {
	LotteryAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowCompleteDraw is a free log retrieval operation binding the contract event 0x52a6fbe07b021d2ea40c304f5e2501643d8cf53c3b51cf784c4d9dc9d1bcd7c2.
//
// Solidity: event LotteryEscrow__CompleteDraw(address lotteryAddress)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowCompleteDraw(opts *bind.FilterOpts) (*LotteryEscrowLotteryEscrowCompleteDrawIterator, error) {

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__CompleteDraw")
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowCompleteDrawIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__CompleteDraw", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowCompleteDraw is a free log subscription operation binding the contract event 0x52a6fbe07b021d2ea40c304f5e2501643d8cf53c3b51cf784c4d9dc9d1bcd7c2.
//
// Solidity: event LotteryEscrow__CompleteDraw(address lotteryAddress)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowCompleteDraw(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowCompleteDraw) (event.Subscription, error) {

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__CompleteDraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowCompleteDraw)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__CompleteDraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowCompleteDraw is a log parse operation binding the contract event 0x52a6fbe07b021d2ea40c304f5e2501643d8cf53c3b51cf784c4d9dc9d1bcd7c2.
//
// Solidity: event LotteryEscrow__CompleteDraw(address lotteryAddress)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowCompleteDraw(log types.Log) (*LotteryEscrowLotteryEscrowCompleteDraw, error) {
	event := new(LotteryEscrowLotteryEscrowCompleteDraw)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__CompleteDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowDepositedIterator is returned from FilterLotteryEscrowDeposited and is used to iterate over the raw logs and unpacked data for LotteryEscrowDeposited events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowDepositedIterator struct {
	Event *LotteryEscrowLotteryEscrowDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowDeposited represents a LotteryEscrowDeposited event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowDeposited struct {
	ConcertId  string
	TicketType *big.Int
	Buyer      common.Address
	Money      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowDeposited is a free log retrieval operation binding the contract event 0x4d6ce4580c5dacbd2299d8ff3d4e3a3c48eb2d078d17a7e679e1078bb1769b74.
//
// Solidity: event LotteryEscrow__Deposited(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowDeposited(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowLotteryEscrowDepositedIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__Deposited", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowDepositedIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__Deposited", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowDeposited is a free log subscription operation binding the contract event 0x4d6ce4580c5dacbd2299d8ff3d4e3a3c48eb2d078d17a7e679e1078bb1769b74.
//
// Solidity: event LotteryEscrow__Deposited(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowDeposited(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowDeposited, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__Deposited", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowDeposited)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowDeposited is a log parse operation binding the contract event 0x4d6ce4580c5dacbd2299d8ff3d4e3a3c48eb2d078d17a7e679e1078bb1769b74.
//
// Solidity: event LotteryEscrow__Deposited(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowDeposited(log types.Log) (*LotteryEscrowLotteryEscrowDeposited, error) {
	event := new(LotteryEscrowLotteryEscrowDeposited)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowNonWinnerIterator is returned from FilterLotteryEscrowNonWinner and is used to iterate over the raw logs and unpacked data for LotteryEscrowNonWinner events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowNonWinnerIterator struct {
	Event *LotteryEscrowLotteryEscrowNonWinner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowNonWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowNonWinner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowNonWinner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowNonWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowNonWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowNonWinner represents a LotteryEscrowNonWinner event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowNonWinner struct {
	ConcertId  string
	TicketType *big.Int
	NonWinner  common.Address
	Money      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowNonWinner is a free log retrieval operation binding the contract event 0x1295862289ea0b345ce149601493b6287e4d2e428d7def6919d6255231cfa741.
//
// Solidity: event LotteryEscrow__NonWinner(string concertId, uint256 indexed ticketType, address nonWinner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowNonWinner(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowLotteryEscrowNonWinnerIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__NonWinner", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowNonWinnerIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__NonWinner", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowNonWinner is a free log subscription operation binding the contract event 0x1295862289ea0b345ce149601493b6287e4d2e428d7def6919d6255231cfa741.
//
// Solidity: event LotteryEscrow__NonWinner(string concertId, uint256 indexed ticketType, address nonWinner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowNonWinner(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowNonWinner, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__NonWinner", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowNonWinner)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__NonWinner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowNonWinner is a log parse operation binding the contract event 0x1295862289ea0b345ce149601493b6287e4d2e428d7def6919d6255231cfa741.
//
// Solidity: event LotteryEscrow__NonWinner(string concertId, uint256 indexed ticketType, address nonWinner, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowNonWinner(log types.Log) (*LotteryEscrowLotteryEscrowNonWinner, error) {
	event := new(LotteryEscrowLotteryEscrowNonWinner)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__NonWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowRefundedIterator is returned from FilterLotteryEscrowRefunded and is used to iterate over the raw logs and unpacked data for LotteryEscrowRefunded events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowRefundedIterator struct {
	Event *LotteryEscrowLotteryEscrowRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowRefunded represents a LotteryEscrowRefunded event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowRefunded struct {
	ConcertId  string
	TicketType *big.Int
	Buyer      common.Address
	Money      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowRefunded is a free log retrieval operation binding the contract event 0x3ec87ffb561b22adf9a03fb21194f35302ad5744806919bd41cda5c63011b6ea.
//
// Solidity: event LotteryEscrow__Refunded(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowRefunded(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowLotteryEscrowRefundedIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__Refunded", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowRefundedIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__Refunded", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowRefunded is a free log subscription operation binding the contract event 0x3ec87ffb561b22adf9a03fb21194f35302ad5744806919bd41cda5c63011b6ea.
//
// Solidity: event LotteryEscrow__Refunded(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowRefunded(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowRefunded, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__Refunded", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowRefunded)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowRefunded is a log parse operation binding the contract event 0x3ec87ffb561b22adf9a03fb21194f35302ad5744806919bd41cda5c63011b6ea.
//
// Solidity: event LotteryEscrow__Refunded(string concertId, uint256 indexed ticketType, address buyer, uint256 money)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowRefunded(log types.Log) (*LotteryEscrowLotteryEscrowRefunded, error) {
	event := new(LotteryEscrowLotteryEscrowRefunded)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowLotteryEscrowWinnerIterator is returned from FilterLotteryEscrowWinner and is used to iterate over the raw logs and unpacked data for LotteryEscrowWinner events raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowWinnerIterator struct {
	Event *LotteryEscrowLotteryEscrowWinner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowLotteryEscrowWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowLotteryEscrowWinner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowLotteryEscrowWinner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowLotteryEscrowWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowLotteryEscrowWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowLotteryEscrowWinner represents a LotteryEscrowWinner event raised by the LotteryEscrow contract.
type LotteryEscrowLotteryEscrowWinner struct {
	ConcertId  string
	TicketType *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLotteryEscrowWinner is a free log retrieval operation binding the contract event 0xaab09bf34c5ca80680ba85d2bbc5ae2668156ed988ddf090b1b6300439be0e74.
//
// Solidity: event LotteryEscrow__Winner(string concertId, uint256 indexed ticketType, address winner)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterLotteryEscrowWinner(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowLotteryEscrowWinnerIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "LotteryEscrow__Winner", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowLotteryEscrowWinnerIterator{contract: _LotteryEscrow.contract, event: "LotteryEscrow__Winner", logs: logs, sub: sub}, nil
}

// WatchLotteryEscrowWinner is a free log subscription operation binding the contract event 0xaab09bf34c5ca80680ba85d2bbc5ae2668156ed988ddf090b1b6300439be0e74.
//
// Solidity: event LotteryEscrow__Winner(string concertId, uint256 indexed ticketType, address winner)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchLotteryEscrowWinner(opts *bind.WatchOpts, sink chan<- *LotteryEscrowLotteryEscrowWinner, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "LotteryEscrow__Winner", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowLotteryEscrowWinner)
				if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Winner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryEscrowWinner is a log parse operation binding the contract event 0xaab09bf34c5ca80680ba85d2bbc5ae2668156ed988ddf090b1b6300439be0e74.
//
// Solidity: event LotteryEscrow__Winner(string concertId, uint256 indexed ticketType, address winner)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseLotteryEscrowWinner(log types.Log) (*LotteryEscrowLotteryEscrowWinner, error) {
	event := new(LotteryEscrowLotteryEscrowWinner)
	if err := _LotteryEscrow.contract.UnpackLog(event, "LotteryEscrow__Winner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the LotteryEscrow contract.
type LotteryEscrowMetadataUpdateIterator struct {
	Event *LotteryEscrowMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowMetadataUpdate represents a MetadataUpdate event raised by the LotteryEscrow contract.
type LotteryEscrowMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*LotteryEscrowMetadataUpdateIterator, error) {

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowMetadataUpdateIterator{contract: _LotteryEscrow.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LotteryEscrowMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowMetadataUpdate)
				if err := _LotteryEscrow.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseMetadataUpdate(log types.Log) (*LotteryEscrowMetadataUpdate, error) {
	event := new(LotteryEscrowMetadataUpdate)
	if err := _LotteryEscrow.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the LotteryEscrow contract.
type LotteryEscrowOwnershipTransferRequestedIterator struct {
	Event *LotteryEscrowOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the LotteryEscrow contract.
type LotteryEscrowOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LotteryEscrowOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowOwnershipTransferRequestedIterator{contract: _LotteryEscrow.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LotteryEscrowOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowOwnershipTransferRequested)
				if err := _LotteryEscrow.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseOwnershipTransferRequested(log types.Log) (*LotteryEscrowOwnershipTransferRequested, error) {
	event := new(LotteryEscrowOwnershipTransferRequested)
	if err := _LotteryEscrow.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LotteryEscrow contract.
type LotteryEscrowOwnershipTransferredIterator struct {
	Event *LotteryEscrowOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the LotteryEscrow contract.
type LotteryEscrowOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LotteryEscrowOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowOwnershipTransferredIterator{contract: _LotteryEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LotteryEscrowOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowOwnershipTransferred)
				if err := _LotteryEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*LotteryEscrowOwnershipTransferred, error) {
	event := new(LotteryEscrowOwnershipTransferred)
	if err := _LotteryEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryEscrowTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LotteryEscrow contract.
type LotteryEscrowTransferIterator struct {
	Event *LotteryEscrowTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LotteryEscrowTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LotteryEscrowTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LotteryEscrowTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowTransfer represents a Transfer event raised by the LotteryEscrow contract.
type LotteryEscrowTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LotteryEscrowTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryEscrow.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowTransferIterator{contract: _LotteryEscrow.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LotteryEscrowTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryEscrow.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowTransfer)
				if err := _LotteryEscrow.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryEscrow *LotteryEscrowFilterer) ParseTransfer(log types.Log) (*LotteryEscrowTransfer, error) {
	event := new(LotteryEscrowTransfer)
	if err := _LotteryEscrow.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
