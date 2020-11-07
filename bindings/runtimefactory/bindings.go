// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package runtimefactory

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RuntimefactoryABI is the input ABI used to generate the binding from.
const RuntimefactoryABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"deployToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RuntimefactoryBin is the compiled bytecode used for deploying new contracts.
var RuntimefactoryBin = "0x608060405234801561001057600080fd5b50610132806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80633cd6ee6d14602d575b600080fd5b60cd60048036036020811015604157600080fd5b810190602081018135640100000000811115605b57600080fd5b820183602082011115606c57600080fd5b80359060200191846001830284011164010000000083111715608d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955060e9945050505050565b604080516001600160a01b039092168252519081900360200190f35b6000808251602084016000f0939250505056fea2646970667358221220e80340307ff5adda2653713ddbcfac16ec183ff57e4cb675e699c9091c037ad464736f6c63430007040033"

// DeployRuntimefactory deploys a new Ethereum contract, binding an instance of Runtimefactory to it.
func DeployRuntimefactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Runtimefactory, error) {
	parsed, err := abi.JSON(strings.NewReader(RuntimefactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RuntimefactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Runtimefactory{RuntimefactoryCaller: RuntimefactoryCaller{contract: contract}, RuntimefactoryTransactor: RuntimefactoryTransactor{contract: contract}, RuntimefactoryFilterer: RuntimefactoryFilterer{contract: contract}}, nil
}

// Runtimefactory is an auto generated Go binding around an Ethereum contract.
type Runtimefactory struct {
	RuntimefactoryCaller     // Read-only binding to the contract
	RuntimefactoryTransactor // Write-only binding to the contract
	RuntimefactoryFilterer   // Log filterer for contract events
}

// RuntimefactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RuntimefactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuntimefactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RuntimefactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuntimefactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RuntimefactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuntimefactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RuntimefactorySession struct {
	Contract     *Runtimefactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RuntimefactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RuntimefactoryCallerSession struct {
	Contract *RuntimefactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RuntimefactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RuntimefactoryTransactorSession struct {
	Contract     *RuntimefactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RuntimefactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RuntimefactoryRaw struct {
	Contract *Runtimefactory // Generic contract binding to access the raw methods on
}

// RuntimefactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RuntimefactoryCallerRaw struct {
	Contract *RuntimefactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RuntimefactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RuntimefactoryTransactorRaw struct {
	Contract *RuntimefactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRuntimefactory creates a new instance of Runtimefactory, bound to a specific deployed contract.
func NewRuntimefactory(address common.Address, backend bind.ContractBackend) (*Runtimefactory, error) {
	contract, err := bindRuntimefactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Runtimefactory{RuntimefactoryCaller: RuntimefactoryCaller{contract: contract}, RuntimefactoryTransactor: RuntimefactoryTransactor{contract: contract}, RuntimefactoryFilterer: RuntimefactoryFilterer{contract: contract}}, nil
}

// NewRuntimefactoryCaller creates a new read-only instance of Runtimefactory, bound to a specific deployed contract.
func NewRuntimefactoryCaller(address common.Address, caller bind.ContractCaller) (*RuntimefactoryCaller, error) {
	contract, err := bindRuntimefactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RuntimefactoryCaller{contract: contract}, nil
}

// NewRuntimefactoryTransactor creates a new write-only instance of Runtimefactory, bound to a specific deployed contract.
func NewRuntimefactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RuntimefactoryTransactor, error) {
	contract, err := bindRuntimefactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RuntimefactoryTransactor{contract: contract}, nil
}

// NewRuntimefactoryFilterer creates a new log filterer instance of Runtimefactory, bound to a specific deployed contract.
func NewRuntimefactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RuntimefactoryFilterer, error) {
	contract, err := bindRuntimefactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RuntimefactoryFilterer{contract: contract}, nil
}

// bindRuntimefactory binds a generic wrapper to an already deployed contract.
func bindRuntimefactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RuntimefactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Runtimefactory *RuntimefactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Runtimefactory.Contract.RuntimefactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Runtimefactory *RuntimefactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Runtimefactory.Contract.RuntimefactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Runtimefactory *RuntimefactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Runtimefactory.Contract.RuntimefactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Runtimefactory *RuntimefactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Runtimefactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Runtimefactory *RuntimefactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Runtimefactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Runtimefactory *RuntimefactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Runtimefactory.Contract.contract.Transact(opts, method, params...)
}

// DeployToken is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Runtimefactory *RuntimefactoryTransactor) DeployToken(opts *bind.TransactOpts, _code []byte) (*types.Transaction, error) {
	return _Runtimefactory.contract.Transact(opts, "deployToken", _code)
}

// DeployToken is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Runtimefactory *RuntimefactorySession) DeployToken(_code []byte) (*types.Transaction, error) {
	return _Runtimefactory.Contract.DeployToken(&_Runtimefactory.TransactOpts, _code)
}

// DeployToken is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Runtimefactory *RuntimefactoryTransactorSession) DeployToken(_code []byte) (*types.Transaction, error) {
	return _Runtimefactory.Contract.DeployToken(&_Runtimefactory.TransactOpts, _code)
}
