// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LotteryMarket

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

// LotteryMarketMetaData contains all meta data concerning the LotteryMarket contract.
var LotteryMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NFTListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NFTSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"delistNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"listNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lotteryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"setFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LotteryMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMarketMetaData.ABI instead.
var LotteryMarketABI = LotteryMarketMetaData.ABI

// LotteryMarket is an auto generated Go binding around an Ethereum contract.
type LotteryMarket struct {
	LotteryMarketCaller     // Read-only binding to the contract
	LotteryMarketTransactor // Write-only binding to the contract
	LotteryMarketFilterer   // Log filterer for contract events
}

// LotteryMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotteryMarketSession struct {
	Contract     *LotteryMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryMarketCallerSession struct {
	Contract *LotteryMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LotteryMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryMarketTransactorSession struct {
	Contract     *LotteryMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LotteryMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryMarketRaw struct {
	Contract *LotteryMarket // Generic contract binding to access the raw methods on
}

// LotteryMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryMarketCallerRaw struct {
	Contract *LotteryMarketCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryMarketTransactorRaw struct {
	Contract *LotteryMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryMarket creates a new instance of LotteryMarket, bound to a specific deployed contract.
func NewLotteryMarket(address common.Address, backend bind.ContractBackend) (*LotteryMarket, error) {
	contract, err := bindLotteryMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LotteryMarket{LotteryMarketCaller: LotteryMarketCaller{contract: contract}, LotteryMarketTransactor: LotteryMarketTransactor{contract: contract}, LotteryMarketFilterer: LotteryMarketFilterer{contract: contract}}, nil
}

// NewLotteryMarketCaller creates a new read-only instance of LotteryMarket, bound to a specific deployed contract.
func NewLotteryMarketCaller(address common.Address, caller bind.ContractCaller) (*LotteryMarketCaller, error) {
	contract, err := bindLotteryMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketCaller{contract: contract}, nil
}

// NewLotteryMarketTransactor creates a new write-only instance of LotteryMarket, bound to a specific deployed contract.
func NewLotteryMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryMarketTransactor, error) {
	contract, err := bindLotteryMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketTransactor{contract: contract}, nil
}

// NewLotteryMarketFilterer creates a new log filterer instance of LotteryMarket, bound to a specific deployed contract.
func NewLotteryMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryMarketFilterer, error) {
	contract, err := bindLotteryMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketFilterer{contract: contract}, nil
}

// bindLotteryMarket binds a generic wrapper to an already deployed contract.
func bindLotteryMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryMarket *LotteryMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryMarket.Contract.LotteryMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryMarket *LotteryMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryMarket.Contract.LotteryMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryMarket *LotteryMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryMarket.Contract.LotteryMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryMarket *LotteryMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryMarket *LotteryMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryMarket *LotteryMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryMarket.Contract.contract.Transact(opts, method, params...)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_LotteryMarket *LotteryMarketCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryMarket.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_LotteryMarket *LotteryMarketSession) FactoryAddress() (common.Address, error) {
	return _LotteryMarket.Contract.FactoryAddress(&_LotteryMarket.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_LotteryMarket *LotteryMarketCallerSession) FactoryAddress() (common.Address, error) {
	return _LotteryMarket.Contract.FactoryAddress(&_LotteryMarket.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0x0007df30.
//
// Solidity: function listings(address , uint256 ) view returns(address seller, address lotteryAddress, uint256 tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Seller         common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Price          *big.Int
}, error) {
	var out []interface{}
	err := _LotteryMarket.contract.Call(opts, &out, "listings", arg0, arg1)

	outstruct := new(struct {
		Seller         common.Address
		LotteryAddress common.Address
		TokenId        *big.Int
		Price          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LotteryAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0x0007df30.
//
// Solidity: function listings(address , uint256 ) view returns(address seller, address lotteryAddress, uint256 tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketSession) Listings(arg0 common.Address, arg1 *big.Int) (struct {
	Seller         common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Price          *big.Int
}, error) {
	return _LotteryMarket.Contract.Listings(&_LotteryMarket.CallOpts, arg0, arg1)
}

// Listings is a free data retrieval call binding the contract method 0x0007df30.
//
// Solidity: function listings(address , uint256 ) view returns(address seller, address lotteryAddress, uint256 tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketCallerSession) Listings(arg0 common.Address, arg1 *big.Int) (struct {
	Seller         common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Price          *big.Int
}, error) {
	return _LotteryMarket.Contract.Listings(&_LotteryMarket.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryMarket *LotteryMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LotteryMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryMarket *LotteryMarketSession) Owner() (common.Address, error) {
	return _LotteryMarket.Contract.Owner(&_LotteryMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LotteryMarket *LotteryMarketCallerSession) Owner() (common.Address, error) {
	return _LotteryMarket.Contract.Owner(&_LotteryMarket.CallOpts)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address lotteryAddress, uint256 tokenId) payable returns()
func (_LotteryMarket *LotteryMarketTransactor) BuyNFT(opts *bind.TransactOpts, lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "buyNFT", lotteryAddress, tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address lotteryAddress, uint256 tokenId) payable returns()
func (_LotteryMarket *LotteryMarketSession) BuyNFT(lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.BuyNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address lotteryAddress, uint256 tokenId) payable returns()
func (_LotteryMarket *LotteryMarketTransactorSession) BuyNFT(lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.BuyNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xa17a7094.
//
// Solidity: function delistNFT(address lotteryAddress, uint256 tokenId) returns()
func (_LotteryMarket *LotteryMarketTransactor) DelistNFT(opts *bind.TransactOpts, lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "delistNFT", lotteryAddress, tokenId)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xa17a7094.
//
// Solidity: function delistNFT(address lotteryAddress, uint256 tokenId) returns()
func (_LotteryMarket *LotteryMarketSession) DelistNFT(lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.DelistNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xa17a7094.
//
// Solidity: function delistNFT(address lotteryAddress, uint256 tokenId) returns()
func (_LotteryMarket *LotteryMarketTransactorSession) DelistNFT(lotteryAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.DelistNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address lotteryAddress, uint256 tokenId, uint256 price) returns()
func (_LotteryMarket *LotteryMarketTransactor) ListNFT(opts *bind.TransactOpts, lotteryAddress common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "listNFT", lotteryAddress, tokenId, price)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address lotteryAddress, uint256 tokenId, uint256 price) returns()
func (_LotteryMarket *LotteryMarketSession) ListNFT(lotteryAddress common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.ListNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId, price)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address lotteryAddress, uint256 tokenId, uint256 price) returns()
func (_LotteryMarket *LotteryMarketTransactorSession) ListNFT(lotteryAddress common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _LotteryMarket.Contract.ListNFT(&_LotteryMarket.TransactOpts, lotteryAddress, tokenId, price)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_LotteryMarket *LotteryMarketTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_LotteryMarket *LotteryMarketSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryMarket.Contract.OnERC721Received(&_LotteryMarket.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_LotteryMarket *LotteryMarketTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryMarket.Contract.OnERC721Received(&_LotteryMarket.TransactOpts, operator, from, tokenId, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LotteryMarket *LotteryMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LotteryMarket *LotteryMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _LotteryMarket.Contract.RenounceOwnership(&_LotteryMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LotteryMarket *LotteryMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LotteryMarket.Contract.RenounceOwnership(&_LotteryMarket.TransactOpts)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_LotteryMarket *LotteryMarketTransactor) SetFactoryAddress(opts *bind.TransactOpts, _factoryAddress common.Address) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "setFactoryAddress", _factoryAddress)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_LotteryMarket *LotteryMarketSession) SetFactoryAddress(_factoryAddress common.Address) (*types.Transaction, error) {
	return _LotteryMarket.Contract.SetFactoryAddress(&_LotteryMarket.TransactOpts, _factoryAddress)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_LotteryMarket *LotteryMarketTransactorSession) SetFactoryAddress(_factoryAddress common.Address) (*types.Transaction, error) {
	return _LotteryMarket.Contract.SetFactoryAddress(&_LotteryMarket.TransactOpts, _factoryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LotteryMarket *LotteryMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LotteryMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LotteryMarket *LotteryMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LotteryMarket.Contract.TransferOwnership(&_LotteryMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LotteryMarket *LotteryMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LotteryMarket.Contract.TransferOwnership(&_LotteryMarket.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryMarket *LotteryMarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryMarket.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryMarket *LotteryMarketSession) Receive() (*types.Transaction, error) {
	return _LotteryMarket.Contract.Receive(&_LotteryMarket.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryMarket *LotteryMarketTransactorSession) Receive() (*types.Transaction, error) {
	return _LotteryMarket.Contract.Receive(&_LotteryMarket.TransactOpts)
}

// LotteryMarketNFTDelistedIterator is returned from FilterNFTDelisted and is used to iterate over the raw logs and unpacked data for NFTDelisted events raised by the LotteryMarket contract.
type LotteryMarketNFTDelistedIterator struct {
	Event *LotteryMarketNFTDelisted // Event containing the contract specifics and raw log

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
func (it *LotteryMarketNFTDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryMarketNFTDelisted)
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
		it.Event = new(LotteryMarketNFTDelisted)
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
func (it *LotteryMarketNFTDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryMarketNFTDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryMarketNFTDelisted represents a NFTDelisted event raised by the LotteryMarket contract.
type LotteryMarketNFTDelisted struct {
	Seller         common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNFTDelisted is a free log retrieval operation binding the contract event 0xfc314074f528769d77ad4a29ede7c622894f632282e936ebefff2646a598c2ee.
//
// Solidity: event NFTDelisted(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId)
func (_LotteryMarket *LotteryMarketFilterer) FilterNFTDelisted(opts *bind.FilterOpts, seller []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (*LotteryMarketNFTDelistedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.FilterLogs(opts, "NFTDelisted", sellerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketNFTDelistedIterator{contract: _LotteryMarket.contract, event: "NFTDelisted", logs: logs, sub: sub}, nil
}

// WatchNFTDelisted is a free log subscription operation binding the contract event 0xfc314074f528769d77ad4a29ede7c622894f632282e936ebefff2646a598c2ee.
//
// Solidity: event NFTDelisted(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId)
func (_LotteryMarket *LotteryMarketFilterer) WatchNFTDelisted(opts *bind.WatchOpts, sink chan<- *LotteryMarketNFTDelisted, seller []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.WatchLogs(opts, "NFTDelisted", sellerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryMarketNFTDelisted)
				if err := _LotteryMarket.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
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

// ParseNFTDelisted is a log parse operation binding the contract event 0xfc314074f528769d77ad4a29ede7c622894f632282e936ebefff2646a598c2ee.
//
// Solidity: event NFTDelisted(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId)
func (_LotteryMarket *LotteryMarketFilterer) ParseNFTDelisted(log types.Log) (*LotteryMarketNFTDelisted, error) {
	event := new(LotteryMarketNFTDelisted)
	if err := _LotteryMarket.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryMarketNFTListedIterator is returned from FilterNFTListed and is used to iterate over the raw logs and unpacked data for NFTListed events raised by the LotteryMarket contract.
type LotteryMarketNFTListedIterator struct {
	Event *LotteryMarketNFTListed // Event containing the contract specifics and raw log

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
func (it *LotteryMarketNFTListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryMarketNFTListed)
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
		it.Event = new(LotteryMarketNFTListed)
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
func (it *LotteryMarketNFTListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryMarketNFTListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryMarketNFTListed represents a NFTListed event raised by the LotteryMarket contract.
type LotteryMarketNFTListed struct {
	Seller         common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNFTListed is a free log retrieval operation binding the contract event 0xbeab3a2bb824b124a8a1eb465eec003338d61b414db132d37e9b3a984fdcf010.
//
// Solidity: event NFTListed(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) FilterNFTListed(opts *bind.FilterOpts, seller []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (*LotteryMarketNFTListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.FilterLogs(opts, "NFTListed", sellerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketNFTListedIterator{contract: _LotteryMarket.contract, event: "NFTListed", logs: logs, sub: sub}, nil
}

// WatchNFTListed is a free log subscription operation binding the contract event 0xbeab3a2bb824b124a8a1eb465eec003338d61b414db132d37e9b3a984fdcf010.
//
// Solidity: event NFTListed(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) WatchNFTListed(opts *bind.WatchOpts, sink chan<- *LotteryMarketNFTListed, seller []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.WatchLogs(opts, "NFTListed", sellerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryMarketNFTListed)
				if err := _LotteryMarket.contract.UnpackLog(event, "NFTListed", log); err != nil {
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

// ParseNFTListed is a log parse operation binding the contract event 0xbeab3a2bb824b124a8a1eb465eec003338d61b414db132d37e9b3a984fdcf010.
//
// Solidity: event NFTListed(address indexed seller, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) ParseNFTListed(log types.Log) (*LotteryMarketNFTListed, error) {
	event := new(LotteryMarketNFTListed)
	if err := _LotteryMarket.contract.UnpackLog(event, "NFTListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryMarketNFTSoldIterator is returned from FilterNFTSold and is used to iterate over the raw logs and unpacked data for NFTSold events raised by the LotteryMarket contract.
type LotteryMarketNFTSoldIterator struct {
	Event *LotteryMarketNFTSold // Event containing the contract specifics and raw log

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
func (it *LotteryMarketNFTSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryMarketNFTSold)
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
		it.Event = new(LotteryMarketNFTSold)
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
func (it *LotteryMarketNFTSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryMarketNFTSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryMarketNFTSold represents a NFTSold event raised by the LotteryMarket contract.
type LotteryMarketNFTSold struct {
	Buyer          common.Address
	LotteryAddress common.Address
	TokenId        *big.Int
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNFTSold is a free log retrieval operation binding the contract event 0xf6ccc0cf89c21c54724ea6ef2ec01da8fe33a1cebe06607e160fd78483eba302.
//
// Solidity: event NFTSold(address indexed buyer, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) FilterNFTSold(opts *bind.FilterOpts, buyer []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (*LotteryMarketNFTSoldIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.FilterLogs(opts, "NFTSold", buyerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketNFTSoldIterator{contract: _LotteryMarket.contract, event: "NFTSold", logs: logs, sub: sub}, nil
}

// WatchNFTSold is a free log subscription operation binding the contract event 0xf6ccc0cf89c21c54724ea6ef2ec01da8fe33a1cebe06607e160fd78483eba302.
//
// Solidity: event NFTSold(address indexed buyer, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) WatchNFTSold(opts *bind.WatchOpts, sink chan<- *LotteryMarketNFTSold, buyer []common.Address, lotteryAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var lotteryAddressRule []interface{}
	for _, lotteryAddressItem := range lotteryAddress {
		lotteryAddressRule = append(lotteryAddressRule, lotteryAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryMarket.contract.WatchLogs(opts, "NFTSold", buyerRule, lotteryAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryMarketNFTSold)
				if err := _LotteryMarket.contract.UnpackLog(event, "NFTSold", log); err != nil {
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

// ParseNFTSold is a log parse operation binding the contract event 0xf6ccc0cf89c21c54724ea6ef2ec01da8fe33a1cebe06607e160fd78483eba302.
//
// Solidity: event NFTSold(address indexed buyer, address indexed lotteryAddress, uint256 indexed tokenId, uint256 price)
func (_LotteryMarket *LotteryMarketFilterer) ParseNFTSold(log types.Log) (*LotteryMarketNFTSold, error) {
	event := new(LotteryMarketNFTSold)
	if err := _LotteryMarket.contract.UnpackLog(event, "NFTSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LotteryMarket contract.
type LotteryMarketOwnershipTransferredIterator struct {
	Event *LotteryMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LotteryMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryMarketOwnershipTransferred)
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
		it.Event = new(LotteryMarketOwnershipTransferred)
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
func (it *LotteryMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryMarketOwnershipTransferred represents a OwnershipTransferred event raised by the LotteryMarket contract.
type LotteryMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LotteryMarket *LotteryMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LotteryMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LotteryMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LotteryMarketOwnershipTransferredIterator{contract: _LotteryMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LotteryMarket *LotteryMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LotteryMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LotteryMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryMarketOwnershipTransferred)
				if err := _LotteryMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LotteryMarket *LotteryMarketFilterer) ParseOwnershipTransferred(log types.Log) (*LotteryMarketOwnershipTransferred, error) {
	event := new(LotteryMarketOwnershipTransferred)
	if err := _LotteryMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
