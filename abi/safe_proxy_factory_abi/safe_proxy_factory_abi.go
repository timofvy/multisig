// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safe_proxy_factory_abi

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
)

// SafeProxyFactoryAbiMetaData contains all meta data concerning the SafeProxyFactoryAbi contract.
var SafeProxyFactoryAbiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"ProxyCreation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createChainSpecificProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIProxyCreationCallback\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"createProxyWithCallback\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SafeProxyFactoryAbiABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyFactoryAbiMetaData.ABI instead.
var SafeProxyFactoryAbiABI = SafeProxyFactoryAbiMetaData.ABI

// SafeProxyFactoryAbi is an auto generated Go binding around an Ethereum contract.
type SafeProxyFactoryAbi struct {
	SafeProxyFactoryAbiCaller     // Read-only binding to the contract
	SafeProxyFactoryAbiTransactor // Write-only binding to the contract
	SafeProxyFactoryAbiFilterer   // Log filterer for contract events
}

// SafeProxyFactoryAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyFactoryAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyFactoryAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFactoryAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxyFactoryAbiSession struct {
	Contract     *SafeProxyFactoryAbi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeProxyFactoryAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyFactoryAbiCallerSession struct {
	Contract *SafeProxyFactoryAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SafeProxyFactoryAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyFactoryAbiTransactorSession struct {
	Contract     *SafeProxyFactoryAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SafeProxyFactoryAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyFactoryAbiRaw struct {
	Contract *SafeProxyFactoryAbi // Generic contract binding to access the raw methods on
}

// SafeProxyFactoryAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyFactoryAbiCallerRaw struct {
	Contract *SafeProxyFactoryAbiCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyFactoryAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyFactoryAbiTransactorRaw struct {
	Contract *SafeProxyFactoryAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxyFactoryAbi creates a new instance of SafeProxyFactoryAbi, bound to a specific deployed contract.
func NewSafeProxyFactoryAbi(address common.Address, backend bind.ContractBackend) (*SafeProxyFactoryAbi, error) {
	contract, err := bindSafeProxyFactoryAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryAbi{SafeProxyFactoryAbiCaller: SafeProxyFactoryAbiCaller{contract: contract}, SafeProxyFactoryAbiTransactor: SafeProxyFactoryAbiTransactor{contract: contract}, SafeProxyFactoryAbiFilterer: SafeProxyFactoryAbiFilterer{contract: contract}}, nil
}

// NewSafeProxyFactoryAbiCaller creates a new read-only instance of SafeProxyFactoryAbi, bound to a specific deployed contract.
func NewSafeProxyFactoryAbiCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyFactoryAbiCaller, error) {
	contract, err := bindSafeProxyFactoryAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryAbiCaller{contract: contract}, nil
}

// NewSafeProxyFactoryAbiTransactor creates a new write-only instance of SafeProxyFactoryAbi, bound to a specific deployed contract.
func NewSafeProxyFactoryAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyFactoryAbiTransactor, error) {
	contract, err := bindSafeProxyFactoryAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryAbiTransactor{contract: contract}, nil
}

// NewSafeProxyFactoryAbiFilterer creates a new log filterer instance of SafeProxyFactoryAbi, bound to a specific deployed contract.
func NewSafeProxyFactoryAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFactoryAbiFilterer, error) {
	contract, err := bindSafeProxyFactoryAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryAbiFilterer{contract: contract}, nil
}

// bindSafeProxyFactoryAbi binds a generic wrapper to an already deployed contract.
func bindSafeProxyFactoryAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeProxyFactoryAbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactoryAbi.Contract.SafeProxyFactoryAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.SafeProxyFactoryAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.SafeProxyFactoryAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactoryAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.contract.Transact(opts, method, params...)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeProxyFactoryAbi.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactoryAbi.Contract.GetChainId(&_SafeProxyFactoryAbi.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiCallerSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactoryAbi.Contract.GetChainId(&_SafeProxyFactoryAbi.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactoryAbi.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactoryAbi.Contract.ProxyCreationCode(&_SafeProxyFactoryAbi.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiCallerSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactoryAbi.Contract.ProxyCreationCode(&_SafeProxyFactoryAbi.CallOpts)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactor) CreateChainSpecificProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.contract.Transact(opts, "createChainSpecificProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactorSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactor) CreateProxyWithCallback(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.contract.Transact(opts, "createProxyWithCallback", _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiSession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateProxyWithCallback(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactorSession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateProxyWithCallback(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.contract.Transact(opts, "createProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateProxyWithNonce(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiTransactorSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryAbi.Contract.CreateProxyWithNonce(&_SafeProxyFactoryAbi.TransactOpts, _singleton, initializer, saltNonce)
}

// SafeProxyFactoryAbiProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the SafeProxyFactoryAbi contract.
type SafeProxyFactoryAbiProxyCreationIterator struct {
	Event *SafeProxyFactoryAbiProxyCreation // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryAbiProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryAbiProxyCreation)
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
		it.Event = new(SafeProxyFactoryAbiProxyCreation)
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
func (it *SafeProxyFactoryAbiProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryAbiProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryAbiProxyCreation represents a ProxyCreation event raised by the SafeProxyFactoryAbi contract.
type SafeProxyFactoryAbiProxyCreation struct {
	Proxy     common.Address
	Singleton common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiFilterer) FilterProxyCreation(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryAbiProxyCreationIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryAbi.contract.FilterLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryAbiProxyCreationIterator{contract: _SafeProxyFactoryAbi.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryAbiProxyCreation, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryAbi.contract.WatchLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryAbiProxyCreation)
				if err := _SafeProxyFactoryAbi.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryAbi *SafeProxyFactoryAbiFilterer) ParseProxyCreation(log types.Log) (*SafeProxyFactoryAbiProxyCreation, error) {
	event := new(SafeProxyFactoryAbiProxyCreation)
	if err := _SafeProxyFactoryAbi.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
