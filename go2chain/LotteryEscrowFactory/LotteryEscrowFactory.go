// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LotteryEscrowFactory

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

// LotteryEscrowFactoryMetaData contains all meta data concerning the LotteryEscrowFactory contract.
var LotteryEscrowFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TicketTypeAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concertId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"}],\"name\":\"EscrowCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allEscrows\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_organizer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_concertId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ticketType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ticketCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ddl\",\"type\":\"uint256\"}],\"name\":\"createEscrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"escrows\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketType\",\"type\":\"uint256\"}],\"name\":\"getEscrowAddressByTicketType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LotteryEscrowFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryEscrowFactoryMetaData.ABI instead.
var LotteryEscrowFactoryABI = LotteryEscrowFactoryMetaData.ABI

// LotteryEscrowFactory is an auto generated Go binding around an Ethereum contract.
type LotteryEscrowFactory struct {
	LotteryEscrowFactoryCaller     // Read-only binding to the contract
	LotteryEscrowFactoryTransactor // Write-only binding to the contract
	LotteryEscrowFactoryFilterer   // Log filterer for contract events
}

// LotteryEscrowFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryEscrowFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryEscrowFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryEscrowFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryEscrowFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotteryEscrowFactorySession struct {
	Contract     *LotteryEscrowFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LotteryEscrowFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryEscrowFactoryCallerSession struct {
	Contract *LotteryEscrowFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LotteryEscrowFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryEscrowFactoryTransactorSession struct {
	Contract     *LotteryEscrowFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LotteryEscrowFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryEscrowFactoryRaw struct {
	Contract *LotteryEscrowFactory // Generic contract binding to access the raw methods on
}

// LotteryEscrowFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryEscrowFactoryCallerRaw struct {
	Contract *LotteryEscrowFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryEscrowFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryEscrowFactoryTransactorRaw struct {
	Contract *LotteryEscrowFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryEscrowFactory creates a new instance of LotteryEscrowFactory, bound to a specific deployed contract.
func NewLotteryEscrowFactory(address common.Address, backend bind.ContractBackend) (*LotteryEscrowFactory, error) {
	contract, err := bindLotteryEscrowFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFactory{LotteryEscrowFactoryCaller: LotteryEscrowFactoryCaller{contract: contract}, LotteryEscrowFactoryTransactor: LotteryEscrowFactoryTransactor{contract: contract}, LotteryEscrowFactoryFilterer: LotteryEscrowFactoryFilterer{contract: contract}}, nil
}

// NewLotteryEscrowFactoryCaller creates a new read-only instance of LotteryEscrowFactory, bound to a specific deployed contract.
func NewLotteryEscrowFactoryCaller(address common.Address, caller bind.ContractCaller) (*LotteryEscrowFactoryCaller, error) {
	contract, err := bindLotteryEscrowFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFactoryCaller{contract: contract}, nil
}

// NewLotteryEscrowFactoryTransactor creates a new write-only instance of LotteryEscrowFactory, bound to a specific deployed contract.
func NewLotteryEscrowFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryEscrowFactoryTransactor, error) {
	contract, err := bindLotteryEscrowFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFactoryTransactor{contract: contract}, nil
}

// NewLotteryEscrowFactoryFilterer creates a new log filterer instance of LotteryEscrowFactory, bound to a specific deployed contract.
func NewLotteryEscrowFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryEscrowFactoryFilterer, error) {
	contract, err := bindLotteryEscrowFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFactoryFilterer{contract: contract}, nil
}

// bindLotteryEscrowFactory binds a generic wrapper to an already deployed contract.
func bindLotteryEscrowFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryEscrowFactory *LotteryEscrowFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryEscrowFactory.Contract.LotteryEscrowFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryEscrowFactory *LotteryEscrowFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.LotteryEscrowFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryEscrowFactory *LotteryEscrowFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.LotteryEscrowFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryEscrowFactory *LotteryEscrowFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryEscrowFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.contract.Transact(opts, method, params...)
}

// AllEscrows is a free data retrieval call binding the contract method 0x1d3afee6.
//
// Solidity: function allEscrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCaller) AllEscrows(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrowFactory.contract.Call(opts, &out, "allEscrows", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllEscrows is a free data retrieval call binding the contract method 0x1d3afee6.
//
// Solidity: function allEscrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) AllEscrows(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.AllEscrows(&_LotteryEscrowFactory.CallOpts, arg0)
}

// AllEscrows is a free data retrieval call binding the contract method 0x1d3afee6.
//
// Solidity: function allEscrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCallerSession) AllEscrows(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.AllEscrows(&_LotteryEscrowFactory.CallOpts, arg0)
}

// Escrows is a free data retrieval call binding the contract method 0x012f52ee.
//
// Solidity: function escrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCaller) Escrows(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrowFactory.contract.Call(opts, &out, "escrows", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrows is a free data retrieval call binding the contract method 0x012f52ee.
//
// Solidity: function escrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) Escrows(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.Escrows(&_LotteryEscrowFactory.CallOpts, arg0)
}

// Escrows is a free data retrieval call binding the contract method 0x012f52ee.
//
// Solidity: function escrows(uint256 ) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCallerSession) Escrows(arg0 *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.Escrows(&_LotteryEscrowFactory.CallOpts, arg0)
}

// GetEscrowAddressByTicketType is a free data retrieval call binding the contract method 0xbb9faf66.
//
// Solidity: function getEscrowAddressByTicketType(uint256 ticketType) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCaller) GetEscrowAddressByTicketType(opts *bind.CallOpts, ticketType *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryEscrowFactory.contract.Call(opts, &out, "getEscrowAddressByTicketType", ticketType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEscrowAddressByTicketType is a free data retrieval call binding the contract method 0xbb9faf66.
//
// Solidity: function getEscrowAddressByTicketType(uint256 ticketType) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) GetEscrowAddressByTicketType(ticketType *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.GetEscrowAddressByTicketType(&_LotteryEscrowFactory.CallOpts, ticketType)
}

// GetEscrowAddressByTicketType is a free data retrieval call binding the contract method 0xbb9faf66.
//
// Solidity: function getEscrowAddressByTicketType(uint256 ticketType) view returns(address)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCallerSession) GetEscrowAddressByTicketType(ticketType *big.Int) (common.Address, error) {
	return _LotteryEscrowFactory.Contract.GetEscrowAddressByTicketType(&_LotteryEscrowFactory.CallOpts, ticketType)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LotteryEscrowFactory.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) IsRegistered(arg0 common.Address) (bool, error) {
	return _LotteryEscrowFactory.Contract.IsRegistered(&_LotteryEscrowFactory.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_LotteryEscrowFactory *LotteryEscrowFactoryCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _LotteryEscrowFactory.Contract.IsRegistered(&_LotteryEscrowFactory.CallOpts, arg0)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0x12f4a228.
//
// Solidity: function createEscrow(address _organizer, string _concertId, uint256 _ticketType, string _typeName, string _name, uint256 _price, string _url, uint256 _ticketCount, uint256 _ddl) returns(address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactor) CreateEscrow(opts *bind.TransactOpts, _organizer common.Address, _concertId string, _ticketType *big.Int, _typeName string, _name string, _price *big.Int, _url string, _ticketCount *big.Int, _ddl *big.Int) (*types.Transaction, error) {
	return _LotteryEscrowFactory.contract.Transact(opts, "createEscrow", _organizer, _concertId, _ticketType, _typeName, _name, _price, _url, _ticketCount, _ddl)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0x12f4a228.
//
// Solidity: function createEscrow(address _organizer, string _concertId, uint256 _ticketType, string _typeName, string _name, uint256 _price, string _url, uint256 _ticketCount, uint256 _ddl) returns(address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) CreateEscrow(_organizer common.Address, _concertId string, _ticketType *big.Int, _typeName string, _name string, _price *big.Int, _url string, _ticketCount *big.Int, _ddl *big.Int) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.CreateEscrow(&_LotteryEscrowFactory.TransactOpts, _organizer, _concertId, _ticketType, _typeName, _name, _price, _url, _ticketCount, _ddl)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0x12f4a228.
//
// Solidity: function createEscrow(address _organizer, string _concertId, uint256 _ticketType, string _typeName, string _name, uint256 _price, string _url, uint256 _ticketCount, uint256 _ddl) returns(address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactorSession) CreateEscrow(_organizer common.Address, _concertId string, _ticketType *big.Int, _typeName string, _name string, _price *big.Int, _url string, _ticketCount *big.Int, _ddl *big.Int) (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.CreateEscrow(&_LotteryEscrowFactory.TransactOpts, _organizer, _concertId, _ticketType, _typeName, _name, _price, _url, _ticketCount, _ddl)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryEscrowFactory.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrowFactory *LotteryEscrowFactorySession) Receive() (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.Receive(&_LotteryEscrowFactory.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LotteryEscrowFactory *LotteryEscrowFactoryTransactorSession) Receive() (*types.Transaction, error) {
	return _LotteryEscrowFactory.Contract.Receive(&_LotteryEscrowFactory.TransactOpts)
}

// LotteryEscrowFactoryEscrowCreatedIterator is returned from FilterEscrowCreated and is used to iterate over the raw logs and unpacked data for EscrowCreated events raised by the LotteryEscrowFactory contract.
type LotteryEscrowFactoryEscrowCreatedIterator struct {
	Event *LotteryEscrowFactoryEscrowCreated // Event containing the contract specifics and raw log

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
func (it *LotteryEscrowFactoryEscrowCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEscrowFactoryEscrowCreated)
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
		it.Event = new(LotteryEscrowFactoryEscrowCreated)
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
func (it *LotteryEscrowFactoryEscrowCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEscrowFactoryEscrowCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEscrowFactoryEscrowCreated represents a EscrowCreated event raised by the LotteryEscrowFactory contract.
type LotteryEscrowFactoryEscrowCreated struct {
	ConcertId     string
	TicketType    *big.Int
	EscrowAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEscrowCreated is a free log retrieval operation binding the contract event 0xed6859b036a1de681a479ac0a9a650bc1c5b8a11968fb44dc2c151eb1a3d8bb3.
//
// Solidity: event EscrowCreated(string concertId, uint256 indexed ticketType, address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactoryFilterer) FilterEscrowCreated(opts *bind.FilterOpts, ticketType []*big.Int) (*LotteryEscrowFactoryEscrowCreatedIterator, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrowFactory.contract.FilterLogs(opts, "EscrowCreated", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return &LotteryEscrowFactoryEscrowCreatedIterator{contract: _LotteryEscrowFactory.contract, event: "EscrowCreated", logs: logs, sub: sub}, nil
}

// WatchEscrowCreated is a free log subscription operation binding the contract event 0xed6859b036a1de681a479ac0a9a650bc1c5b8a11968fb44dc2c151eb1a3d8bb3.
//
// Solidity: event EscrowCreated(string concertId, uint256 indexed ticketType, address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactoryFilterer) WatchEscrowCreated(opts *bind.WatchOpts, sink chan<- *LotteryEscrowFactoryEscrowCreated, ticketType []*big.Int) (event.Subscription, error) {

	var ticketTypeRule []interface{}
	for _, ticketTypeItem := range ticketType {
		ticketTypeRule = append(ticketTypeRule, ticketTypeItem)
	}

	logs, sub, err := _LotteryEscrowFactory.contract.WatchLogs(opts, "EscrowCreated", ticketTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEscrowFactoryEscrowCreated)
				if err := _LotteryEscrowFactory.contract.UnpackLog(event, "EscrowCreated", log); err != nil {
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

// ParseEscrowCreated is a log parse operation binding the contract event 0xed6859b036a1de681a479ac0a9a650bc1c5b8a11968fb44dc2c151eb1a3d8bb3.
//
// Solidity: event EscrowCreated(string concertId, uint256 indexed ticketType, address escrowAddress)
func (_LotteryEscrowFactory *LotteryEscrowFactoryFilterer) ParseEscrowCreated(log types.Log) (*LotteryEscrowFactoryEscrowCreated, error) {
	event := new(LotteryEscrowFactoryEscrowCreated)
	if err := _LotteryEscrowFactory.contract.UnpackLog(event, "EscrowCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
