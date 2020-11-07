// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20factory

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

// Erc20factoryABI is the input ABI used to generate the binding from.
const Erc20factoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"deployToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"deployToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"generatedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"symbolHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Erc20factoryBin is the compiled bytecode used for deploying new contracts.
var Erc20factoryBin = "0x608060405234801561001057600080fd5b50611264806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630712b9d01461005157806322123ef71461008a5780633cd6ee6d146101ea578063e486033914610290575b600080fd5b61006e6004803603602081101561006757600080fd5b50356102ec565b604080516001600160a01b039092168252519081900360200190f35b6101d6600480360360a08110156100a057600080fd5b6001600160a01b038235169160208101359160ff60408301351691908101906080810160608201356401000000008111156100da57600080fd5b8201836020820111156100ec57600080fd5b8035906020019184600183028401116401000000008311171561010e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561016157600080fd5b82018360208201111561017357600080fd5b8035906020019184600183028401116401000000008311171561019557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610316945050505050565b604080519115158252519081900360200190f35b61006e6004803603602081101561020057600080fd5b81019060208101813564010000000081111561021b57600080fd5b82018360208201111561022d57600080fd5b8035906020019184600183028401116401000000008311171561024f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106a9945050505050565b6102b6600480360360208110156102a657600080fd5b50356001600160a01b03166106bc565b6040805195865260208601949094526001600160a01b0390921684840152606084015260ff166080830152519081900360a00190f35b600081815481106102fc57600080fd5b6000918252602090912001546001600160a01b0316905081565b60006060868686868660405160200180866001600160a01b031681526020018581526020018460ff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015610381578181015183820152602001610369565b50505050905090810190601f1680156103ae5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156103e15781810151838201526020016103c9565b50505050905090810190601f16801561040e5780820380516001836020036101000a031916815260200191505b509750505050505050506040516020818303038152906040529050606061045060405180610a800160405280610a5c81526020016107a3610a5c9139836106f7565b9050600061045d826106a9565b90506001600160a01b0381166104a45760405162461bcd60e51b81526004018080602001828103825260308152602001806111ff6030913960400191505060405180910390fd5b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0383161790556104f4610774565b6040518060a00160405280886040516020018082805190602001908083835b602083106105325780518252601f199092019160209182019101610513565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001208152602001876040516020018082805190602001908083835b602083106105a55780518252601f199092019160209182019101610586565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012081526020018b6001600160a01b031681526020018a81526020018960ff1681525090508060016000846001600160a01b03166001600160a01b03168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506060820151816003015560808201518160040160006101000a81548160ff021916908360ff160217905550905050600194505050505095945050505050565b6000808251602084016000f09392505050565b60016020819052600091825260409091208054918101546002820154600383015460049093015491926001600160a01b039091169160ff1685565b6060806040519050835180825260208201818101602087015b81831015610728578051835260209283019201610710565b50855184518101855292509050808201602086015b8183101561075557805183526020928301920161073d565b508651929092011591909101601f01601f191660405250905092915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe608060405234801561001057600080fd5b50604051610a5c380380610a5c8339810160409081528151602080840151928401516060850151608086015160008054600160a060020a031916600160a060020a03871617905560038690556004805460ff191660ff85161790559086018051949692949093919092019161008b9160019190850190610106565b50805161009f906002906020840190610106565b50600160a060020a038516600081815260056020908152604080832088905580519283529082019290925280820186905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a150505050506101a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014757805160ff1916838001178555610174565b82800160010185558215610174579182015b82811115610174578251825591602001919060010190610159565b50610180929150610184565b5090565b61019e91905b80821115610180576000815560010161018a565b90565b6108ac806101b06000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a8578063095ea7b31461013257806318160ddd1461016a57806323b872dd1461019157806327e235e3146101bb578063313ce567146101dc5780635c658165146102075780638da5cb5b1461022e57806395d89b411461025f578063a9059cbb14610274575b600080fd5b3480156100b457600080fd5b506100bd610298565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f75781810151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013e57600080fd5b50610156600160a060020a0360043516602435610325565b604080519115158252519081900360200190f35b34801561017657600080fd5b5061017f61041f565b60408051918252519081900360200190f35b34801561019d57600080fd5b50610156600160a060020a0360043581169060243516604435610425565b3480156101c757600080fd5b5061017f600160a060020a036004351661066e565b3480156101e857600080fd5b506101f1610680565b6040805160ff9092168252519081900360200190f35b34801561021357600080fd5b5061017f600160a060020a0360043581169060243516610689565b34801561023a57600080fd5b506102436106a6565b60408051600160a060020a039092168252519081900360200190f35b34801561026b57600080fd5b506100bd6106b5565b34801561028057600080fd5b50610156600160a060020a036004351660243561070d565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561031d5780601f106102f25761010080835404028352916020019161031d565b820191906000526020600020905b81548152906001019060200180831161030057829003601f168201915b505050505081565b600080821161037e576040805160e560020a62461bcd02815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e2030000000604482015290519081900360640190fd5b336000908152600660209081526040808320600160a060020a03871684529091529020546103b2908363ffffffff61085216565b336000818152600660209081526040808320600160a060020a0389168085529083529281902094909455835192835282015280820184905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a150600192915050565b60035481565b600160a060020a0383166000908152600560205260408120548211156104bb576040805160e560020a62461bcd02815260206004820152602160248201527f6f776e657220646f6573206e6f74206861766520656e6f75676820746f6b656e60448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a038416600090815260066020908152604080832033845290915290205482111561055c576040805160e560020a62461bcd02815260206004820152602560248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820616c6c6f60448201527f77616e6365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600660209081526040808320338452909152902054610590908363ffffffff61086b16565b600160a060020a0385166000818152600660209081526040808320338452825280832094909455918152600590915220546105d1908363ffffffff61086b16565b600160a060020a038086166000908152600560205260408082209390935590851681522054610606908363ffffffff61085216565b600160a060020a0380851660008181526005602090815260409182902094909455805192881683529282015280820184905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a15060019392505050565b60056020526000908152604090205481565b60045460ff1681565b600660209081526000928352604080842090915290825290205481565b600054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561031d5780601f106102f25761010080835404028352916020019161031d565b3360009081526005602052604081205482111561079a576040805160e560020a62461bcd02815260206004820152602260248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820746f6b6560448201527f6e73000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600560205260409020546107ba908363ffffffff61086b16565b3360009081526005602052604080822092909255600160a060020a038516815220546107ec908363ffffffff61085216565b600160a060020a03841660008181526005602090815260409182902093909355805133815292830191909152818101849052517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a150600192915050565b60008282018381101561086457600080fd5b9392505050565b60008282111561087a57600080fd5b509003905600a165627a7a7230582025bc031d1ee00a03299cef0075e167cb75cb0c8b6d02b38db025f3376386289100296661696c656420746f206465706c6f7920636f6e74726163742c206c696b656c792072616e206f7574206f6620676173a264697066735822122098e42e827c78388827af3ec8db6cd2f9fd506deb77524d81ebec4adb5b66b20d64736f6c63430007040033"

// DeployErc20factory deploys a new Ethereum contract, binding an instance of Erc20factory to it.
func DeployErc20factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20factory, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20factoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Erc20factoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20factory{Erc20factoryCaller: Erc20factoryCaller{contract: contract}, Erc20factoryTransactor: Erc20factoryTransactor{contract: contract}, Erc20factoryFilterer: Erc20factoryFilterer{contract: contract}}, nil
}

// Erc20factory is an auto generated Go binding around an Ethereum contract.
type Erc20factory struct {
	Erc20factoryCaller     // Read-only binding to the contract
	Erc20factoryTransactor // Write-only binding to the contract
	Erc20factoryFilterer   // Log filterer for contract events
}

// Erc20factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20factorySession struct {
	Contract     *Erc20factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20factoryCallerSession struct {
	Contract *Erc20factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20factoryTransactorSession struct {
	Contract     *Erc20factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20factoryRaw struct {
	Contract *Erc20factory // Generic contract binding to access the raw methods on
}

// Erc20factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20factoryCallerRaw struct {
	Contract *Erc20factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20factoryTransactorRaw struct {
	Contract *Erc20factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20factory creates a new instance of Erc20factory, bound to a specific deployed contract.
func NewErc20factory(address common.Address, backend bind.ContractBackend) (*Erc20factory, error) {
	contract, err := bindErc20factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20factory{Erc20factoryCaller: Erc20factoryCaller{contract: contract}, Erc20factoryTransactor: Erc20factoryTransactor{contract: contract}, Erc20factoryFilterer: Erc20factoryFilterer{contract: contract}}, nil
}

// NewErc20factoryCaller creates a new read-only instance of Erc20factory, bound to a specific deployed contract.
func NewErc20factoryCaller(address common.Address, caller bind.ContractCaller) (*Erc20factoryCaller, error) {
	contract, err := bindErc20factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20factoryCaller{contract: contract}, nil
}

// NewErc20factoryTransactor creates a new write-only instance of Erc20factory, bound to a specific deployed contract.
func NewErc20factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20factoryTransactor, error) {
	contract, err := bindErc20factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20factoryTransactor{contract: contract}, nil
}

// NewErc20factoryFilterer creates a new log filterer instance of Erc20factory, bound to a specific deployed contract.
func NewErc20factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20factoryFilterer, error) {
	contract, err := bindErc20factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20factoryFilterer{contract: contract}, nil
}

// bindErc20factory binds a generic wrapper to an already deployed contract.
func bindErc20factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20factoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20factory *Erc20factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20factory.Contract.Erc20factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20factory *Erc20factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20factory.Contract.Erc20factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20factory *Erc20factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20factory.Contract.Erc20factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20factory *Erc20factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20factory *Erc20factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20factory *Erc20factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20factory.Contract.contract.Transact(opts, method, params...)
}

// GeneratedTokens is a free data retrieval call binding the contract method 0x0712b9d0.
//
// Solidity: function generatedTokens(uint256 ) view returns(address)
func (_Erc20factory *Erc20factoryCaller) GeneratedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc20factory.contract.Call(opts, &out, "generatedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneratedTokens is a free data retrieval call binding the contract method 0x0712b9d0.
//
// Solidity: function generatedTokens(uint256 ) view returns(address)
func (_Erc20factory *Erc20factorySession) GeneratedTokens(arg0 *big.Int) (common.Address, error) {
	return _Erc20factory.Contract.GeneratedTokens(&_Erc20factory.CallOpts, arg0)
}

// GeneratedTokens is a free data retrieval call binding the contract method 0x0712b9d0.
//
// Solidity: function generatedTokens(uint256 ) view returns(address)
func (_Erc20factory *Erc20factoryCallerSession) GeneratedTokens(arg0 *big.Int) (common.Address, error) {
	return _Erc20factory.Contract.GeneratedTokens(&_Erc20factory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(bytes32 nameHash, bytes32 symbolHash, address owner, uint256 supply, uint8 decimals)
func (_Erc20factory *Erc20factoryCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	NameHash   [32]byte
	SymbolHash [32]byte
	Owner      common.Address
	Supply     *big.Int
	Decimals   uint8
}, error) {
	var out []interface{}
	err := _Erc20factory.contract.Call(opts, &out, "tokens", arg0)

	outstruct := new(struct {
		NameHash   [32]byte
		SymbolHash [32]byte
		Owner      common.Address
		Supply     *big.Int
		Decimals   uint8
	})

	outstruct.NameHash = out[0].([32]byte)
	outstruct.SymbolHash = out[1].([32]byte)
	outstruct.Owner = out[2].(common.Address)
	outstruct.Supply = out[3].(*big.Int)
	outstruct.Decimals = out[4].(uint8)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(bytes32 nameHash, bytes32 symbolHash, address owner, uint256 supply, uint8 decimals)
func (_Erc20factory *Erc20factorySession) Tokens(arg0 common.Address) (struct {
	NameHash   [32]byte
	SymbolHash [32]byte
	Owner      common.Address
	Supply     *big.Int
	Decimals   uint8
}, error) {
	return _Erc20factory.Contract.Tokens(&_Erc20factory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(bytes32 nameHash, bytes32 symbolHash, address owner, uint256 supply, uint8 decimals)
func (_Erc20factory *Erc20factoryCallerSession) Tokens(arg0 common.Address) (struct {
	NameHash   [32]byte
	SymbolHash [32]byte
	Owner      common.Address
	Supply     *big.Int
	Decimals   uint8
}, error) {
	return _Erc20factory.Contract.Tokens(&_Erc20factory.CallOpts, arg0)
}

// DeployToken is a paid mutator transaction binding the contract method 0x22123ef7.
//
// Solidity: function deployToken(address _owner, uint256 _supply, uint8 _decimals, string _name, string _symbol) returns(bool)
func (_Erc20factory *Erc20factoryTransactor) DeployToken(opts *bind.TransactOpts, _owner common.Address, _supply *big.Int, _decimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc20factory.contract.Transact(opts, "deployToken", _owner, _supply, _decimals, _name, _symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0x22123ef7.
//
// Solidity: function deployToken(address _owner, uint256 _supply, uint8 _decimals, string _name, string _symbol) returns(bool)
func (_Erc20factory *Erc20factorySession) DeployToken(_owner common.Address, _supply *big.Int, _decimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc20factory.Contract.DeployToken(&_Erc20factory.TransactOpts, _owner, _supply, _decimals, _name, _symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0x22123ef7.
//
// Solidity: function deployToken(address _owner, uint256 _supply, uint8 _decimals, string _name, string _symbol) returns(bool)
func (_Erc20factory *Erc20factoryTransactorSession) DeployToken(_owner common.Address, _supply *big.Int, _decimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc20factory.Contract.DeployToken(&_Erc20factory.TransactOpts, _owner, _supply, _decimals, _name, _symbol)
}

// DeployToken0 is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Erc20factory *Erc20factoryTransactor) DeployToken0(opts *bind.TransactOpts, _code []byte) (*types.Transaction, error) {
	return _Erc20factory.contract.Transact(opts, "deployToken0", _code)
}

// DeployToken0 is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Erc20factory *Erc20factorySession) DeployToken0(_code []byte) (*types.Transaction, error) {
	return _Erc20factory.Contract.DeployToken0(&_Erc20factory.TransactOpts, _code)
}

// DeployToken0 is a paid mutator transaction binding the contract method 0x3cd6ee6d.
//
// Solidity: function deployToken(bytes _code) returns(address)
func (_Erc20factory *Erc20factoryTransactorSession) DeployToken0(_code []byte) (*types.Transaction, error) {
	return _Erc20factory.Contract.DeployToken0(&_Erc20factory.TransactOpts, _code)
}
