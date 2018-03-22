// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package event_ipfs_storage

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EventIpfsStorageABI is the input ABI used to generate the binding from.
const EventIpfsStorageABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"EmitStringStorageEvent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_a\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_b\",\"type\":\"string\"}],\"name\":\"TextData\",\"type\":\"event\"}]"

// EventIpfsStorageBin is the compiled bytecode used for deploying new contracts.
const EventIpfsStorageBin = `60606040523415600e57600080fd5b6101358061001d6000396000f3006060604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e2bdf1868114610045575b600080fd5b341561005057600080fd5b61005861006c565b604051901515815260200160405180910390f35b60007f12308742bdfc772218894fe869825ebf18a809117fd27ffec52af8511d87949b604051604080825260038183018190527f616263000000000000000000000000000000000000000000000000000000000060608401526080602084018190528301527f646566000000000000000000000000000000000000000000000000000000000060a083015260c0909101905180910390a1506001905600a165627a7a7230582013f3d5a27620dced411e4373955622c55fd5bd83f928776ac08176b5897cd48f0029`

// DeployEventIpfsStorage deploys a new Ethereum contract, binding an instance of EventIpfsStorage to it.
func DeployEventIpfsStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventIpfsStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(EventIpfsStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventIpfsStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventIpfsStorage{EventIpfsStorageCaller: EventIpfsStorageCaller{contract: contract}, EventIpfsStorageTransactor: EventIpfsStorageTransactor{contract: contract}, EventIpfsStorageFilterer: EventIpfsStorageFilterer{contract: contract}}, nil
}

// EventIpfsStorage is an auto generated Go binding around an Ethereum contract.
type EventIpfsStorage struct {
	EventIpfsStorageCaller     // Read-only binding to the contract
	EventIpfsStorageTransactor // Write-only binding to the contract
	EventIpfsStorageFilterer   // Log filterer for contract events
}

// EventIpfsStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventIpfsStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventIpfsStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventIpfsStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventIpfsStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventIpfsStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventIpfsStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventIpfsStorageSession struct {
	Contract     *EventIpfsStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventIpfsStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventIpfsStorageCallerSession struct {
	Contract *EventIpfsStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EventIpfsStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventIpfsStorageTransactorSession struct {
	Contract     *EventIpfsStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EventIpfsStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventIpfsStorageRaw struct {
	Contract *EventIpfsStorage // Generic contract binding to access the raw methods on
}

// EventIpfsStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventIpfsStorageCallerRaw struct {
	Contract *EventIpfsStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EventIpfsStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventIpfsStorageTransactorRaw struct {
	Contract *EventIpfsStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventIpfsStorage creates a new instance of EventIpfsStorage, bound to a specific deployed contract.
func NewEventIpfsStorage(address common.Address, backend bind.ContractBackend) (*EventIpfsStorage, error) {
	contract, err := bindEventIpfsStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventIpfsStorage{EventIpfsStorageCaller: EventIpfsStorageCaller{contract: contract}, EventIpfsStorageTransactor: EventIpfsStorageTransactor{contract: contract}, EventIpfsStorageFilterer: EventIpfsStorageFilterer{contract: contract}}, nil
}

// NewEventIpfsStorageCaller creates a new read-only instance of EventIpfsStorage, bound to a specific deployed contract.
func NewEventIpfsStorageCaller(address common.Address, caller bind.ContractCaller) (*EventIpfsStorageCaller, error) {
	contract, err := bindEventIpfsStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventIpfsStorageCaller{contract: contract}, nil
}

// NewEventIpfsStorageTransactor creates a new write-only instance of EventIpfsStorage, bound to a specific deployed contract.
func NewEventIpfsStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EventIpfsStorageTransactor, error) {
	contract, err := bindEventIpfsStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventIpfsStorageTransactor{contract: contract}, nil
}

// NewEventIpfsStorageFilterer creates a new log filterer instance of EventIpfsStorage, bound to a specific deployed contract.
func NewEventIpfsStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EventIpfsStorageFilterer, error) {
	contract, err := bindEventIpfsStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventIpfsStorageFilterer{contract: contract}, nil
}

// bindEventIpfsStorage binds a generic wrapper to an already deployed contract.
func bindEventIpfsStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventIpfsStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventIpfsStorage *EventIpfsStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EventIpfsStorage.Contract.EventIpfsStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventIpfsStorage *EventIpfsStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.EventIpfsStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventIpfsStorage *EventIpfsStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.EventIpfsStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventIpfsStorage *EventIpfsStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EventIpfsStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventIpfsStorage *EventIpfsStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventIpfsStorage *EventIpfsStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.contract.Transact(opts, method, params...)
}

// EmitStringStorageEvent is a paid mutator transaction binding the contract method 0xe2bdf186.
//
// Solidity: function EmitStringStorageEvent() returns(bool)
func (_EventIpfsStorage *EventIpfsStorageTransactor) EmitStringStorageEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventIpfsStorage.contract.Transact(opts, "EmitStringStorageEvent")
}

// EmitStringStorageEvent is a paid mutator transaction binding the contract method 0xe2bdf186.
//
// Solidity: function EmitStringStorageEvent() returns(bool)
func (_EventIpfsStorage *EventIpfsStorageSession) EmitStringStorageEvent() (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.EmitStringStorageEvent(&_EventIpfsStorage.TransactOpts)
}

// EmitStringStorageEvent is a paid mutator transaction binding the contract method 0xe2bdf186.
//
// Solidity: function EmitStringStorageEvent() returns(bool)
func (_EventIpfsStorage *EventIpfsStorageTransactorSession) EmitStringStorageEvent() (*types.Transaction, error) {
	return _EventIpfsStorage.Contract.EmitStringStorageEvent(&_EventIpfsStorage.TransactOpts)
}

// EventIpfsStorageTextDataIterator is returned from FilterTextData and is used to iterate over the raw logs and unpacked data for TextData events raised by the EventIpfsStorage contract.
type EventIpfsStorageTextDataIterator struct {
	Event *EventIpfsStorageTextData // Event containing the contract specifics and raw log

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
func (it *EventIpfsStorageTextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventIpfsStorageTextData)
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
		it.Event = new(EventIpfsStorageTextData)
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
func (it *EventIpfsStorageTextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventIpfsStorageTextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventIpfsStorageTextData represents a TextData event raised by the EventIpfsStorage contract.
type EventIpfsStorageTextData struct {
	A   string
	B   string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTextData is a free log retrieval operation binding the contract event 0x12308742bdfc772218894fe869825ebf18a809117fd27ffec52af8511d87949b.
//
// Solidity: event TextData(_a string, _b string)
func (_EventIpfsStorage *EventIpfsStorageFilterer) FilterTextData(opts *bind.FilterOpts) (*EventIpfsStorageTextDataIterator, error) {

	logs, sub, err := _EventIpfsStorage.contract.FilterLogs(opts, "TextData")
	if err != nil {
		return nil, err
	}
	return &EventIpfsStorageTextDataIterator{contract: _EventIpfsStorage.contract, event: "TextData", logs: logs, sub: sub}, nil
}

// WatchTextData is a free log subscription operation binding the contract event 0x12308742bdfc772218894fe869825ebf18a809117fd27ffec52af8511d87949b.
//
// Solidity: event TextData(_a string, _b string)
func (_EventIpfsStorage *EventIpfsStorageFilterer) WatchTextData(opts *bind.WatchOpts, sink chan<- *EventIpfsStorageTextData) (event.Subscription, error) {

	logs, sub, err := _EventIpfsStorage.contract.WatchLogs(opts, "TextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventIpfsStorageTextData)
				if err := _EventIpfsStorage.contract.UnpackLog(event, "TextData", log); err != nil {
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
