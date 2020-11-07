// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymentchannel

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

// PaymentchannelABI is the input ABI used to generate the binding from.
const PaymentchannelABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recoveredAddress\",\"type\":\"address\"}],\"name\":\"DestinationProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ErcChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"EthChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingChannelValue\",\"type\":\"uint256\"}],\"name\":\"MicroPaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_moderator\",\"type\":\"address\"}],\"name\":\"ModeratorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_moderator\",\"type\":\"address\"}],\"name\":\"ModeratorUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recoveredAddress\",\"type\":\"address\"}],\"name\":\"SourceProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"closeErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultState\",\"outputs\":[{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ercChannels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openDate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"sourceProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destinationProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ethChannels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openDate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"sourceProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destinationProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"expireErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_channelValueInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationInDays\",\"type\":\"uint256\"}],\"name\":\"openErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_channelValueInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationInDays\",\"type\":\"uint256\"}],\"name\":\"openEthChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_moderator\",\"type\":\"address\"}],\"name\":\"setModerator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitErcDestinationProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitErcMicroPaymentProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitErcSourceProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitEthDestinationProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitEthSourceProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_moderator\",\"type\":\"address\"}],\"name\":\"unsetModerator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// PaymentchannelBin is the compiled bytecode used for deploying new contracts.
var PaymentchannelBin = "0x60c0604052601c60808190527f19457468657265756d205369676e6564204d6573736167653a0a33320000000060a090815262000040916003919062000089565b506004805460ff191660011761ff00191690553480156200006057600080fd5b5060008054336001600160a01b0319918216811790925560018054909116909117905562000135565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620000c157600085556200010c565b82601f10620000dc57805160ff19168380011785556200010c565b828001600101855582156200010c579182015b828111156200010c578251825591602001919060010190620000ef565b506200011a9291506200011e565b5090565b5b808211156200011a57600081556001016200011f565b611e2480620001456000396000f3fe60806040526004361061014f5760003560e01c806391cca3db116100b6578063ed15652f1161006f578063ed15652f146105f7578063f2fde38b1461062a578063f851a4401461065d578063f8be342e14610672578063fcf0301d146106a8578063ffa1ad74146106ed57610156565b806391cca3db146104785780639ef3ae9d1461048d578063a0ef91df146104d2578063aa13a612146104e7578063b1653e6714610519578063c21d92d8146105b257610156565b8063704b6c0211610108578063704b6c021461031257806375bba1891461034557806378bdd839146103785780637baccc6d146103bd5780638da5cb5b1461040e5780638db538021461043f57610156565b8063054f7d9c1461015857806314d0f1ba146101815780632b2b4a75146101b45780633f73cfd7146101fd5780633f93cc451461023657806349df728c146102df57610156565b3661015657005b005b34801561016457600080fd5b5061016d610777565b604080519115158252519081900360200190f35b34801561018d57600080fd5b5061016d600480360360208110156101a457600080fd5b50356001600160a01b0316610787565b3480156101c057600080fd5b5061016d600480360360808110156101d757600080fd5b506001600160a01b0381358116916020810135909116906040810135906060013561079c565b34801561020957600080fd5b5061016d6004803603604081101561022057600080fd5b50803590602001356001600160a01b031661095a565b34801561024257600080fd5b506102606004803603602081101561025957600080fd5b5035610b02565b604051808b6001600160a01b031681526020018a6001600160a01b03168152602001896001600160a01b03168152602001888152602001878152602001868152602001858152602001841515815260200183151581526020018260038111156102c557fe5b81526020019a505050505050505050505060405180910390f35b3480156102eb57600080fd5b5061016d6004803603602081101561030257600080fd5b50356001600160a01b0316610b6c565b34801561031e57600080fd5b5061016d6004803603602081101561033557600080fd5b50356001600160a01b0316610cd2565b34801561035157600080fd5b5061016d6004803603602081101561036857600080fd5b50356001600160a01b0316610d5e565b34801561038457600080fd5b5061016d600480360360a081101561039b57600080fd5b5080359060ff6020820135169060408101359060608101359060800135610deb565b3480156103c957600080fd5b5061016d600480360360e08110156103e057600080fd5b5080359060ff6020820135169060408101359060608101359060808101359060a08101359060c00135610fa0565b34801561041a57600080fd5b50610423611271565b604080516001600160a01b039092168252519081900360200190f35b34801561044b57600080fd5b5061016d6004803603604081101561046257600080fd5b50803590602001356001600160a01b0316611280565b34801561048457600080fd5b5061016d6113e9565b34801561049957600080fd5b5061016d600480360360a08110156104b057600080fd5b5080359060ff60208201351690604081013590606081013590608001356113f2565b3480156104de57600080fd5b5061016d6115b8565b61016d600480360360608110156104fd57600080fd5b506001600160a01b038135169060208101359060400135611651565b34801561052557600080fd5b506105436004803603602081101561053c57600080fd5b50356117b2565b604051808a6001600160a01b03168152602001896001600160a01b031681526020018881526020018781526020018681526020018581526020018415158152602001831515815260200182600381111561059957fe5b8152602001995050505050505050505060405180910390f35b3480156105be57600080fd5b5061016d600480360360a08110156105d557600080fd5b5080359060ff6020820135169060408101359060608101359060800135611813565b34801561060357600080fd5b5061016d6004803603602081101561061a57600080fd5b50356001600160a01b0316611996565b34801561063657600080fd5b5061016d6004803603602081101561064d57600080fd5b50356001600160a01b0316611a20565b34801561066957600080fd5b50610423611ac8565b34801561067e57600080fd5b50610687611ad7565b6040518082600381111561069757fe5b815260200191505060405180910390f35b3480156106b457600080fd5b5061016d600480360360a08110156106cb57600080fd5b5080359060ff6020820135169060408101359060608101359060800135611ae5565b3480156106f957600080fd5b50610702611c89565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561073c578181015183820152602001610724565b50505050905090810190601f1680156107695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600154600160a01b900460ff1681565b60026020526000908152604090205460ff1681565b6040805133606090811b6020808401919091526bffffffffffffffffffffffff1987831b811660348501529188901b909116604883015242605c8084019190915283518084039091018152607c9092018352815191810191909120600081815260089092529181205490919060ff161561081557600080fd5b6000818152600860209081526040808320805460ff1916600190811790915560069283905281842080546001600160a01b03199081163317825591810180546001600160a01b038c811691851691909117909155600282018054918d169190931617909155600381018890554262015180880281016004830155600582015590910183905551879183917f41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb7327629190a2604080516323b872dd60e01b81523360048201523060248201526044810187905290516001600160a01b038316916323b872dd9160648083019260209291908290030181600087803b15801561091857600080fd5b505af115801561092c573d6000803e3d6000fd5b505050506040513d602081101561094257600080fd5b505161094d57600080fd5b5060019695505050505050565b60008281526008602052604081205460ff1661097557600080fd5b6000808481526006602052604090206007015462010000900460ff16600381111561099c57fe5b146109a657600080fd5b6000838152600660205260409020546001600160a01b031633146109c957600080fd5b6000838152600660205260409020600201546001600160a01b038381169116146109f257600080fd5b600083815260066020526040812060038101805492905560070180546001919062ff0000191662010000830217905550604051839085907fb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe90600090a26040516001600160a01b03851690600080516020611dcf83398151915290600090a26040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b158015610ac057600080fd5b505af1158015610ad4573d6000803e3d6000fd5b505050506040513d6020811015610aea57600080fd5b5051610af557600080fd5b6001925050505b92915050565b6006602081905260009182526040909120805460018201546002830154600384015460048501546005860154968601546007909601546001600160a01b039586169794861696939095169491939092909160ff80821691610100810482169162010000909104168a565b600080546001600160a01b0316331480610b9057506001546001600160a01b031633145b610b9957600080fd5b60045460ff16610ba857600080fd5b604080516370a0823160e01b8152306004820152905183916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b158015610bf357600080fd5b505afa158015610c07573d6000803e3d6000fd5b505050506040513d6020811015610c1d57600080fd5b50516040519091506001600160a01b03851690600080516020611dcf83398151915290600090a26040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610c9357600080fd5b505af1158015610ca7573d6000803e3d6000fd5b505050506040513d6020811015610cbd57600080fd5b5051610cc857600080fd5b5060019392505050565b600080546001600160a01b03163314610cea57600080fd5b6001546001600160a01b0383811691161415610d0557600080fd5b600180546001600160a01b0384166001600160a01b0319909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a1506001919050565b600080546001600160a01b0316331480610d8257506001546001600160a01b031633145b610d8b57600080fd5b6001600160a01b038216600081815260026020908152604091829020805460ff19166001179055815192835290517f07abfa597e8dcfbbfa28a386fa0728caf3360f382d92356232125424bcaa6a539281900390910190a1506001919050565b60008181526008602052604081205460ff16610e0657600080fd5b6000808381526006602052604090206007015462010000900460ff166003811115610e2d57fe5b14610e3757600080fd5b60008281526006602052604090206007015460ff1615610e5657600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610eb2573d6000803e3d6000fd5b505060408051601f1901516000868152600660205291909120549092506001600160a01b038084169116149050610ee557fe5b60008381526006602090815260408083206007018054600160ff199182168117909255600984528285208c86528452828520338652909352818420805490931617909155516001600160a01b0383169185917fef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a679190a3610f66836001611caf565b1561094d57600083815260066020526040902060070180546002919062ff0000191662010000835b02179055505060019695505050505050565b60008381526008602052604081205460ff16610fbb57600080fd5b6000808581526006602052604090206007015462010000900460ff166003811115610fe257fe5b14610fec57600080fd5b60008481526006602052604090206003015461100757600080fd5b6000848152600660205260409020600101546001600160a01b0316331461102d57600080fd5b6000848484604051602001808481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209050600060038260405160200180838054600181600116156101000203166002900480156110cb5780601f106110a95761010080835404028352918201916110cb565b820191906000526020600020905b8154815290600101906020018083116110b7575b505091825250604080518083038152602092830190915280519101209150508981146110f357fe5b600060018b8b8b8b60405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561114f573d6000803e3d6000fd5b505060408051601f19015160008a8152600660205291909120549092506001600160a01b03808416911614905061118257fe5b60008781526006602052604081206003015461119e9087611db9565b600089815260066020526040808220600381018490556002015490519293506001600160a01b0316918291600080516020611dcf83398151915291a26040805163a9059cbb60e01b81523360048201526024810189905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b15801561122957600080fd5b505af115801561123d573d6000803e3d6000fd5b505050506040513d602081101561125357600080fd5b505161125e57600080fd5b5060019c9b505050505050505050505050565b6000546001600160a01b031681565b60008281526008602052604081205460ff1661129b57600080fd5b600260008481526006602052604090206007015462010000900460ff1660038111156112c357fe5b146112cd57600080fd5b6000838152600660205260409020600101546001600160a01b031633146112f357600080fd5b6000838152600660205260409020600201546001600160a01b0383811691161461131c57600080fd5b6000838152600660205260408120600380820180549390556007909101805462ff0000191662010000830217905550604051839085907fceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba7290600090a26040516001600160a01b03851690600080516020611dcf83398151915290600090a26040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b158015610ac057600080fd5b60045460ff1681565b60008181526008602052604081205460ff1661140d57600080fd5b6000808381526005602052604090206006015462010000900460ff16600381111561143457fe5b1461143e57600080fd5b60008281526005602052604090206006015460ff161561145d57600080fd5b6000828152600960209081526040808320898452825280832033845290915290205460ff161561148c57600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156114e8573d6000803e3d6000fd5b505060408051601f1901516000868152600560205291909120549092506001600160a01b03808416911614905061151b57fe5b6000838152600560209081526040808320600601805460ff19166001179055600982528083208a84528252808320338452909152516001600160a01b0383169185917fef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a679190a361158c836000611caf565b1561094d57600083815260056020526040902060060180546002919062ff000019166201000083610f8e565b600080546001600160a01b03163314806115dc57506001546001600160a01b031633145b6115e557600080fd5b60045460ff166115f457600080fd5b6040517fe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac90600090a160405133904780156108fc02916000818181858888f19350505050158015611649573d6000803e3d6000fd5b506001905090565b600082341461165f57600080fd5b6040805133606090811b6020808401919091529087901b6bffffffffffffffffffffffff19166034830152604882018690524260688084018290528451808503909101815260889093018452825192820192909220600081815260089092529290205490919060ff16156116d257600080fd5b6000818152600860209081526040808320805460ff19166001908117909155600592839052922080546001600160a01b03199081163317825592810180546001600160a01b038b16941693909317909255600282018790554262015180870201600380840191909155600480840186905591830184905590546006909201805461010090930460ff1692909162ff000019909116906201000090849081111561177757fe5b021790555060405181907ec725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa90600090a250600195945050505050565b600560208190526000918252604090912080546001820154600283015460038401546004850154958501546006909501546001600160a01b0394851696939094169491939092919060ff808216916101008104821691620100009091041689565b60008181526008602052604081205460ff1661182e57600080fd5b6000808381526006602052604090206007015462010000900460ff16600381111561185557fe5b1461185f57600080fd5b600082815260066020526040902060070154610100900460ff161561188357600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156118df573d6000803e3d6000fd5b505060408051601f1901516000868152600660205291909120600101549092506001600160a01b03808416911614905061191557fe5b6000838152600660209081526040808320600701805461ff001916610100179055600982528083208a84528252808320338452909152808220805460ff19166001179055516001600160a01b0383169185917f4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac9190a3610f66836001611caf565b600080546001600160a01b03163314806119ba57506001546001600160a01b031633145b6119c357600080fd5b6001600160a01b038216600081815260026020908152604091829020805460ff19169055815192835290517f6e69dd57dae2d984f9d5b55a45dac79a60c83270d517cfd720af1eda0dd2d09b9281900390910190a1506001919050565b600080546001600160a01b03163314611a3857600080fd5b6000546001600160a01b03838116911614801590611a5e57506001600160a01b03821615155b611a6757600080fd5b600080546001600160a01b0319166001600160a01b03841690811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a1506001919050565b6001546001600160a01b031681565b600454610100900460ff1681565b60008181526008602052604081205460ff16611b0057600080fd5b6000808381526005602052604090206006015462010000900460ff166003811115611b2757fe5b14611b3157600080fd5b600082815260056020526040902060060154610100900460ff1615611b5557600080fd5b6000828152600960209081526040808320898452825280832033845290915290205460ff1615611b8457600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611be0573d6000803e3d6000fd5b505060408051601f1901516000868152600560205291909120600101549092506001600160a01b038084169116149050611c1657fe5b6000838152600560209081526040808320600601805461ff001916610100179055600982528083208a84528252808320338452909152516001600160a01b0383169185917f4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac9190a361158c836000611caf565b6040518060400160405280600a815260200169302e342e33616c70686160b01b81525081565b60008115611d36576000808481526006602052604090206007015462010000900460ff166003811115611cde57fe5b14611ce857600080fd5b60008381526006602052604090206007015460ff1615156001148015611d245750600083815260066020526040902060070154610100900460ff165b15611d3157506001610afc565b611db0565b6000808481526005602052604090206006015462010000900460ff166003811115611d5d57fe5b14611d6757600080fd5b60008381526005602052604090206006015460ff1615156001148015611da35750600083815260056020526040902060060154610100900460ff165b15611db057506001610afc565b50600092915050565b600082821115611dc857600080fd5b5090039056feb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273a2646970667358221220b4cfb929094c0bc167b247e913fa0b06fd026b4998ebd4bf223e019178f2883564736f6c63430007040033"

// DeployPaymentchannel deploys a new Ethereum contract, binding an instance of Paymentchannel to it.
func DeployPaymentchannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Paymentchannel, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentchannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentchannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Paymentchannel{PaymentchannelCaller: PaymentchannelCaller{contract: contract}, PaymentchannelTransactor: PaymentchannelTransactor{contract: contract}, PaymentchannelFilterer: PaymentchannelFilterer{contract: contract}}, nil
}

// Paymentchannel is an auto generated Go binding around an Ethereum contract.
type Paymentchannel struct {
	PaymentchannelCaller     // Read-only binding to the contract
	PaymentchannelTransactor // Write-only binding to the contract
	PaymentchannelFilterer   // Log filterer for contract events
}

// PaymentchannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentchannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentchannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentchannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentchannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentchannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentchannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentchannelSession struct {
	Contract     *Paymentchannel   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentchannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentchannelCallerSession struct {
	Contract *PaymentchannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PaymentchannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentchannelTransactorSession struct {
	Contract     *PaymentchannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentchannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentchannelRaw struct {
	Contract *Paymentchannel // Generic contract binding to access the raw methods on
}

// PaymentchannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentchannelCallerRaw struct {
	Contract *PaymentchannelCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentchannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentchannelTransactorRaw struct {
	Contract *PaymentchannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentchannel creates a new instance of Paymentchannel, bound to a specific deployed contract.
func NewPaymentchannel(address common.Address, backend bind.ContractBackend) (*Paymentchannel, error) {
	contract, err := bindPaymentchannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paymentchannel{PaymentchannelCaller: PaymentchannelCaller{contract: contract}, PaymentchannelTransactor: PaymentchannelTransactor{contract: contract}, PaymentchannelFilterer: PaymentchannelFilterer{contract: contract}}, nil
}

// NewPaymentchannelCaller creates a new read-only instance of Paymentchannel, bound to a specific deployed contract.
func NewPaymentchannelCaller(address common.Address, caller bind.ContractCaller) (*PaymentchannelCaller, error) {
	contract, err := bindPaymentchannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelCaller{contract: contract}, nil
}

// NewPaymentchannelTransactor creates a new write-only instance of Paymentchannel, bound to a specific deployed contract.
func NewPaymentchannelTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentchannelTransactor, error) {
	contract, err := bindPaymentchannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelTransactor{contract: contract}, nil
}

// NewPaymentchannelFilterer creates a new log filterer instance of Paymentchannel, bound to a specific deployed contract.
func NewPaymentchannelFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentchannelFilterer, error) {
	contract, err := bindPaymentchannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelFilterer{contract: contract}, nil
}

// bindPaymentchannel binds a generic wrapper to an already deployed contract.
func bindPaymentchannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentchannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymentchannel *PaymentchannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymentchannel.Contract.PaymentchannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymentchannel *PaymentchannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentchannel.Contract.PaymentchannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymentchannel *PaymentchannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymentchannel.Contract.PaymentchannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymentchannel *PaymentchannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymentchannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymentchannel *PaymentchannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentchannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymentchannel *PaymentchannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymentchannel.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Paymentchannel *PaymentchannelCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Paymentchannel *PaymentchannelSession) VERSION() (string, error) {
	return _Paymentchannel.Contract.VERSION(&_Paymentchannel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Paymentchannel *PaymentchannelCallerSession) VERSION() (string, error) {
	return _Paymentchannel.Contract.VERSION(&_Paymentchannel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Paymentchannel *PaymentchannelCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Paymentchannel *PaymentchannelSession) Admin() (common.Address, error) {
	return _Paymentchannel.Contract.Admin(&_Paymentchannel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Paymentchannel *PaymentchannelCallerSession) Admin() (common.Address, error) {
	return _Paymentchannel.Contract.Admin(&_Paymentchannel.CallOpts)
}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Paymentchannel *PaymentchannelCaller) DefaultState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "defaultState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Paymentchannel *PaymentchannelSession) DefaultState() (uint8, error) {
	return _Paymentchannel.Contract.DefaultState(&_Paymentchannel.CallOpts)
}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Paymentchannel *PaymentchannelCallerSession) DefaultState() (uint8, error) {
	return _Paymentchannel.Contract.DefaultState(&_Paymentchannel.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Paymentchannel *PaymentchannelCaller) Dev(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Paymentchannel *PaymentchannelSession) Dev() (bool, error) {
	return _Paymentchannel.Contract.Dev(&_Paymentchannel.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Paymentchannel *PaymentchannelCallerSession) Dev() (bool, error) {
	return _Paymentchannel.Contract.Dev(&_Paymentchannel.CallOpts)
}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelCaller) ErcChannels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "ercChannels", arg0)

	outstruct := new(struct {
		Source                    common.Address
		Destination               common.Address
		TokenAddress              common.Address
		Value                     *big.Int
		ClosingDate               *big.Int
		OpenDate                  *big.Int
		ChannelId                 [32]byte
		SourceProofSubmitted      bool
		DestinationProofSubmitted bool
		State                     uint8
	})

	outstruct.Source = out[0].(common.Address)
	outstruct.Destination = out[1].(common.Address)
	outstruct.TokenAddress = out[2].(common.Address)
	outstruct.Value = out[3].(*big.Int)
	outstruct.ClosingDate = out[4].(*big.Int)
	outstruct.OpenDate = out[5].(*big.Int)
	outstruct.ChannelId = out[6].([32]byte)
	outstruct.SourceProofSubmitted = out[7].(bool)
	outstruct.DestinationProofSubmitted = out[8].(bool)
	outstruct.State = out[9].(uint8)

	return *outstruct, err

}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelSession) ErcChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Paymentchannel.Contract.ErcChannels(&_Paymentchannel.CallOpts, arg0)
}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelCallerSession) ErcChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Paymentchannel.Contract.ErcChannels(&_Paymentchannel.CallOpts, arg0)
}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelCaller) EthChannels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "ethChannels", arg0)

	outstruct := new(struct {
		Source                    common.Address
		Destination               common.Address
		Value                     *big.Int
		ClosingDate               *big.Int
		OpenDate                  *big.Int
		ChannelId                 [32]byte
		SourceProofSubmitted      bool
		DestinationProofSubmitted bool
		State                     uint8
	})

	outstruct.Source = out[0].(common.Address)
	outstruct.Destination = out[1].(common.Address)
	outstruct.Value = out[2].(*big.Int)
	outstruct.ClosingDate = out[3].(*big.Int)
	outstruct.OpenDate = out[4].(*big.Int)
	outstruct.ChannelId = out[5].([32]byte)
	outstruct.SourceProofSubmitted = out[6].(bool)
	outstruct.DestinationProofSubmitted = out[7].(bool)
	outstruct.State = out[8].(uint8)

	return *outstruct, err

}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelSession) EthChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Paymentchannel.Contract.EthChannels(&_Paymentchannel.CallOpts, arg0)
}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Paymentchannel *PaymentchannelCallerSession) EthChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Paymentchannel.Contract.EthChannels(&_Paymentchannel.CallOpts, arg0)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Paymentchannel *PaymentchannelCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "frozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Paymentchannel *PaymentchannelSession) Frozen() (bool, error) {
	return _Paymentchannel.Contract.Frozen(&_Paymentchannel.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Paymentchannel *PaymentchannelCallerSession) Frozen() (bool, error) {
	return _Paymentchannel.Contract.Frozen(&_Paymentchannel.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Paymentchannel *PaymentchannelCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "moderators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Paymentchannel *PaymentchannelSession) Moderators(arg0 common.Address) (bool, error) {
	return _Paymentchannel.Contract.Moderators(&_Paymentchannel.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Paymentchannel *PaymentchannelCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _Paymentchannel.Contract.Moderators(&_Paymentchannel.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymentchannel *PaymentchannelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymentchannel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymentchannel *PaymentchannelSession) Owner() (common.Address, error) {
	return _Paymentchannel.Contract.Owner(&_Paymentchannel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymentchannel *PaymentchannelCallerSession) Owner() (common.Address, error) {
	return _Paymentchannel.Contract.Owner(&_Paymentchannel.CallOpts)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) CloseErcChannel(opts *bind.TransactOpts, _channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "closeErcChannel", _channelId, _tokenAddress)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelSession) CloseErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.CloseErcChannel(&_Paymentchannel.TransactOpts, _channelId, _tokenAddress)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) CloseErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.CloseErcChannel(&_Paymentchannel.TransactOpts, _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) ExpireErcChannel(opts *bind.TransactOpts, _channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "expireErcChannel", _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelSession) ExpireErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.ExpireErcChannel(&_Paymentchannel.TransactOpts, _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) ExpireErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.ExpireErcChannel(&_Paymentchannel.TransactOpts, _channelId, _tokenAddress)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) OpenErcChannel(opts *bind.TransactOpts, _tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "openErcChannel", _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Paymentchannel *PaymentchannelSession) OpenErcChannel(_tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.OpenErcChannel(&_Paymentchannel.TransactOpts, _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) OpenErcChannel(_tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.OpenErcChannel(&_Paymentchannel.TransactOpts, _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) OpenEthChannel(opts *bind.TransactOpts, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "openEthChannel", _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Paymentchannel *PaymentchannelSession) OpenEthChannel(_destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.OpenEthChannel(&_Paymentchannel.TransactOpts, _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) OpenEthChannel(_destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.OpenEthChannel(&_Paymentchannel.TransactOpts, _destination, _channelValueInWei, _durationInDays)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SetAdmin(&_Paymentchannel.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SetAdmin(&_Paymentchannel.TransactOpts, _newAdmin)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SetModerator(opts *bind.TransactOpts, _moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "setModerator", _moderator)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SetModerator(_moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SetModerator(&_Paymentchannel.TransactOpts, _moderator)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SetModerator(_moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SetModerator(&_Paymentchannel.TransactOpts, _moderator)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SubmitErcDestinationProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "submitErcDestinationProof", _h, _v, _r, _s, _channelId)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SubmitErcDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcDestinationProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SubmitErcDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcDestinationProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SubmitErcMicroPaymentProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "submitErcMicroPaymentProof", _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SubmitErcMicroPaymentProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcMicroPaymentProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SubmitErcMicroPaymentProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcMicroPaymentProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SubmitErcSourceProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "submitErcSourceProof", _h, _v, _r, _s, _channelId)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SubmitErcSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcSourceProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SubmitErcSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitErcSourceProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SubmitEthDestinationProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "submitEthDestinationProof", _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SubmitEthDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitEthDestinationProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SubmitEthDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitEthDestinationProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) SubmitEthSourceProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "submitEthSourceProof", _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelSession) SubmitEthSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitEthSourceProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) SubmitEthSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.SubmitEthSourceProof(&_Paymentchannel.TransactOpts, _h, _v, _r, _s, _channelId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Paymentchannel *PaymentchannelSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.TransferOwnership(&_Paymentchannel.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.TransferOwnership(&_Paymentchannel.TransactOpts, _newOwner)
}

// UnsetModerator is a paid mutator transaction binding the contract method 0xed15652f.
//
// Solidity: function unsetModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) UnsetModerator(opts *bind.TransactOpts, _moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "unsetModerator", _moderator)
}

// UnsetModerator is a paid mutator transaction binding the contract method 0xed15652f.
//
// Solidity: function unsetModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelSession) UnsetModerator(_moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.UnsetModerator(&_Paymentchannel.TransactOpts, _moderator)
}

// UnsetModerator is a paid mutator transaction binding the contract method 0xed15652f.
//
// Solidity: function unsetModerator(address _moderator) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) UnsetModerator(_moderator common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.UnsetModerator(&_Paymentchannel.TransactOpts, _moderator)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Paymentchannel *PaymentchannelSession) WithdrawEth() (*types.Transaction, error) {
	return _Paymentchannel.Contract.WithdrawEth(&_Paymentchannel.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _Paymentchannel.Contract.WithdrawEth(&_Paymentchannel.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactor) WithdrawTokens(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.contract.Transact(opts, "withdrawTokens", _tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelSession) WithdrawTokens(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.WithdrawTokens(&_Paymentchannel.TransactOpts, _tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Paymentchannel *PaymentchannelTransactorSession) WithdrawTokens(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Paymentchannel.Contract.WithdrawTokens(&_Paymentchannel.TransactOpts, _tokenAddress)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Paymentchannel *PaymentchannelTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Paymentchannel.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Paymentchannel *PaymentchannelSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.Fallback(&_Paymentchannel.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Paymentchannel *PaymentchannelTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Paymentchannel.Contract.Fallback(&_Paymentchannel.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paymentchannel *PaymentchannelTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentchannel.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paymentchannel *PaymentchannelSession) Receive() (*types.Transaction, error) {
	return _Paymentchannel.Contract.Receive(&_Paymentchannel.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paymentchannel *PaymentchannelTransactorSession) Receive() (*types.Transaction, error) {
	return _Paymentchannel.Contract.Receive(&_Paymentchannel.TransactOpts)
}

// PaymentchannelAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Paymentchannel contract.
type PaymentchannelAdminSetIterator struct {
	Event *PaymentchannelAdminSet // Event containing the contract specifics and raw log

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
func (it *PaymentchannelAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelAdminSet)
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
		it.Event = new(PaymentchannelAdminSet)
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
func (it *PaymentchannelAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelAdminSet represents a AdminSet event raised by the Paymentchannel contract.
type PaymentchannelAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Paymentchannel *PaymentchannelFilterer) FilterAdminSet(opts *bind.FilterOpts) (*PaymentchannelAdminSetIterator, error) {

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &PaymentchannelAdminSetIterator{contract: _Paymentchannel.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Paymentchannel *PaymentchannelFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *PaymentchannelAdminSet) (event.Subscription, error) {

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelAdminSet)
				if err := _Paymentchannel.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// ParseAdminSet is a log parse operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Paymentchannel *PaymentchannelFilterer) ParseAdminSet(log types.Log) (*PaymentchannelAdminSet, error) {
	event := new(PaymentchannelAdminSet)
	if err := _Paymentchannel.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the Paymentchannel contract.
type PaymentchannelChannelClosedIterator struct {
	Event *PaymentchannelChannelClosed // Event containing the contract specifics and raw log

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
func (it *PaymentchannelChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelChannelClosed)
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
		it.Event = new(PaymentchannelChannelClosed)
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
func (it *PaymentchannelChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelChannelClosed represents a ChannelClosed event raised by the Paymentchannel contract.
type PaymentchannelChannelClosed struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) FilterChannelClosed(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentchannelChannelClosedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelChannelClosedIterator{contract: _Paymentchannel.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *PaymentchannelChannelClosed, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelChannelClosed)
				if err := _Paymentchannel.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// ParseChannelClosed is a log parse operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) ParseChannelClosed(log types.Log) (*PaymentchannelChannelClosed, error) {
	event := new(PaymentchannelChannelClosed)
	if err := _Paymentchannel.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelChannelExpiredIterator is returned from FilterChannelExpired and is used to iterate over the raw logs and unpacked data for ChannelExpired events raised by the Paymentchannel contract.
type PaymentchannelChannelExpiredIterator struct {
	Event *PaymentchannelChannelExpired // Event containing the contract specifics and raw log

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
func (it *PaymentchannelChannelExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelChannelExpired)
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
		it.Event = new(PaymentchannelChannelExpired)
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
func (it *PaymentchannelChannelExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelChannelExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelChannelExpired represents a ChannelExpired event raised by the Paymentchannel contract.
type PaymentchannelChannelExpired struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelExpired is a free log retrieval operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) FilterChannelExpired(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentchannelChannelExpiredIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "ChannelExpired", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelChannelExpiredIterator{contract: _Paymentchannel.contract, event: "ChannelExpired", logs: logs, sub: sub}, nil
}

// WatchChannelExpired is a free log subscription operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) WatchChannelExpired(opts *bind.WatchOpts, sink chan<- *PaymentchannelChannelExpired, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "ChannelExpired", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelChannelExpired)
				if err := _Paymentchannel.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
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

// ParseChannelExpired is a log parse operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) ParseChannelExpired(log types.Log) (*PaymentchannelChannelExpired, error) {
	event := new(PaymentchannelChannelExpired)
	if err := _Paymentchannel.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelDestinationProofSubmittedIterator is returned from FilterDestinationProofSubmitted and is used to iterate over the raw logs and unpacked data for DestinationProofSubmitted events raised by the Paymentchannel contract.
type PaymentchannelDestinationProofSubmittedIterator struct {
	Event *PaymentchannelDestinationProofSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentchannelDestinationProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelDestinationProofSubmitted)
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
		it.Event = new(PaymentchannelDestinationProofSubmitted)
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
func (it *PaymentchannelDestinationProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelDestinationProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelDestinationProofSubmitted represents a DestinationProofSubmitted event raised by the Paymentchannel contract.
type PaymentchannelDestinationProofSubmitted struct {
	ChannelId        [32]byte
	RecoveredAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDestinationProofSubmitted is a free log retrieval operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) FilterDestinationProofSubmitted(opts *bind.FilterOpts, _channelId [][32]byte, _recoveredAddress []common.Address) (*PaymentchannelDestinationProofSubmittedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "DestinationProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelDestinationProofSubmittedIterator{contract: _Paymentchannel.contract, event: "DestinationProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchDestinationProofSubmitted is a free log subscription operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) WatchDestinationProofSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentchannelDestinationProofSubmitted, _channelId [][32]byte, _recoveredAddress []common.Address) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "DestinationProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelDestinationProofSubmitted)
				if err := _Paymentchannel.contract.UnpackLog(event, "DestinationProofSubmitted", log); err != nil {
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

// ParseDestinationProofSubmitted is a log parse operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) ParseDestinationProofSubmitted(log types.Log) (*PaymentchannelDestinationProofSubmitted, error) {
	event := new(PaymentchannelDestinationProofSubmitted)
	if err := _Paymentchannel.contract.UnpackLog(event, "DestinationProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelErcChannelOpenedIterator is returned from FilterErcChannelOpened and is used to iterate over the raw logs and unpacked data for ErcChannelOpened events raised by the Paymentchannel contract.
type PaymentchannelErcChannelOpenedIterator struct {
	Event *PaymentchannelErcChannelOpened // Event containing the contract specifics and raw log

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
func (it *PaymentchannelErcChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelErcChannelOpened)
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
		it.Event = new(PaymentchannelErcChannelOpened)
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
func (it *PaymentchannelErcChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelErcChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelErcChannelOpened represents a ErcChannelOpened event raised by the Paymentchannel contract.
type PaymentchannelErcChannelOpened struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterErcChannelOpened is a free log retrieval operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) FilterErcChannelOpened(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentchannelErcChannelOpenedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "ErcChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelErcChannelOpenedIterator{contract: _Paymentchannel.contract, event: "ErcChannelOpened", logs: logs, sub: sub}, nil
}

// WatchErcChannelOpened is a free log subscription operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) WatchErcChannelOpened(opts *bind.WatchOpts, sink chan<- *PaymentchannelErcChannelOpened, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "ErcChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelErcChannelOpened)
				if err := _Paymentchannel.contract.UnpackLog(event, "ErcChannelOpened", log); err != nil {
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

// ParseErcChannelOpened is a log parse operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) ParseErcChannelOpened(log types.Log) (*PaymentchannelErcChannelOpened, error) {
	event := new(PaymentchannelErcChannelOpened)
	if err := _Paymentchannel.contract.UnpackLog(event, "ErcChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelEthChannelOpenedIterator is returned from FilterEthChannelOpened and is used to iterate over the raw logs and unpacked data for EthChannelOpened events raised by the Paymentchannel contract.
type PaymentchannelEthChannelOpenedIterator struct {
	Event *PaymentchannelEthChannelOpened // Event containing the contract specifics and raw log

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
func (it *PaymentchannelEthChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelEthChannelOpened)
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
		it.Event = new(PaymentchannelEthChannelOpened)
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
func (it *PaymentchannelEthChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelEthChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelEthChannelOpened represents a EthChannelOpened event raised by the Paymentchannel contract.
type PaymentchannelEthChannelOpened struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthChannelOpened is a free log retrieval operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) FilterEthChannelOpened(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentchannelEthChannelOpenedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "EthChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelEthChannelOpenedIterator{contract: _Paymentchannel.contract, event: "EthChannelOpened", logs: logs, sub: sub}, nil
}

// WatchEthChannelOpened is a free log subscription operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) WatchEthChannelOpened(opts *bind.WatchOpts, sink chan<- *PaymentchannelEthChannelOpened, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "EthChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelEthChannelOpened)
				if err := _Paymentchannel.contract.UnpackLog(event, "EthChannelOpened", log); err != nil {
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

// ParseEthChannelOpened is a log parse operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Paymentchannel *PaymentchannelFilterer) ParseEthChannelOpened(log types.Log) (*PaymentchannelEthChannelOpened, error) {
	event := new(PaymentchannelEthChannelOpened)
	if err := _Paymentchannel.contract.UnpackLog(event, "EthChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the Paymentchannel contract.
type PaymentchannelEthWithdrawnIterator struct {
	Event *PaymentchannelEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentchannelEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelEthWithdrawn)
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
		it.Event = new(PaymentchannelEthWithdrawn)
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
func (it *PaymentchannelEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelEthWithdrawn represents a EthWithdrawn event raised by the Paymentchannel contract.
type PaymentchannelEthWithdrawn struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Paymentchannel *PaymentchannelFilterer) FilterEthWithdrawn(opts *bind.FilterOpts) (*PaymentchannelEthWithdrawnIterator, error) {

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PaymentchannelEthWithdrawnIterator{contract: _Paymentchannel.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Paymentchannel *PaymentchannelFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentchannelEthWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelEthWithdrawn)
				if err := _Paymentchannel.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Paymentchannel *PaymentchannelFilterer) ParseEthWithdrawn(log types.Log) (*PaymentchannelEthWithdrawn, error) {
	event := new(PaymentchannelEthWithdrawn)
	if err := _Paymentchannel.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelMicroPaymentWithdrawnIterator is returned from FilterMicroPaymentWithdrawn and is used to iterate over the raw logs and unpacked data for MicroPaymentWithdrawn events raised by the Paymentchannel contract.
type PaymentchannelMicroPaymentWithdrawnIterator struct {
	Event *PaymentchannelMicroPaymentWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentchannelMicroPaymentWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelMicroPaymentWithdrawn)
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
		it.Event = new(PaymentchannelMicroPaymentWithdrawn)
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
func (it *PaymentchannelMicroPaymentWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelMicroPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelMicroPaymentWithdrawn represents a MicroPaymentWithdrawn event raised by the Paymentchannel contract.
type PaymentchannelMicroPaymentWithdrawn struct {
	ChannelId             [32]byte
	Amount                *big.Int
	RemainingChannelValue *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMicroPaymentWithdrawn is a free log retrieval operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Paymentchannel *PaymentchannelFilterer) FilterMicroPaymentWithdrawn(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentchannelMicroPaymentWithdrawnIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "MicroPaymentWithdrawn", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelMicroPaymentWithdrawnIterator{contract: _Paymentchannel.contract, event: "MicroPaymentWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMicroPaymentWithdrawn is a free log subscription operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Paymentchannel *PaymentchannelFilterer) WatchMicroPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentchannelMicroPaymentWithdrawn, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "MicroPaymentWithdrawn", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelMicroPaymentWithdrawn)
				if err := _Paymentchannel.contract.UnpackLog(event, "MicroPaymentWithdrawn", log); err != nil {
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

// ParseMicroPaymentWithdrawn is a log parse operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Paymentchannel *PaymentchannelFilterer) ParseMicroPaymentWithdrawn(log types.Log) (*PaymentchannelMicroPaymentWithdrawn, error) {
	event := new(PaymentchannelMicroPaymentWithdrawn)
	if err := _Paymentchannel.contract.UnpackLog(event, "MicroPaymentWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelModeratorSetIterator is returned from FilterModeratorSet and is used to iterate over the raw logs and unpacked data for ModeratorSet events raised by the Paymentchannel contract.
type PaymentchannelModeratorSetIterator struct {
	Event *PaymentchannelModeratorSet // Event containing the contract specifics and raw log

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
func (it *PaymentchannelModeratorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelModeratorSet)
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
		it.Event = new(PaymentchannelModeratorSet)
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
func (it *PaymentchannelModeratorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelModeratorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelModeratorSet represents a ModeratorSet event raised by the Paymentchannel contract.
type PaymentchannelModeratorSet struct {
	Moderator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModeratorSet is a free log retrieval operation binding the contract event 0x07abfa597e8dcfbbfa28a386fa0728caf3360f382d92356232125424bcaa6a53.
//
// Solidity: event ModeratorSet(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) FilterModeratorSet(opts *bind.FilterOpts) (*PaymentchannelModeratorSetIterator, error) {

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "ModeratorSet")
	if err != nil {
		return nil, err
	}
	return &PaymentchannelModeratorSetIterator{contract: _Paymentchannel.contract, event: "ModeratorSet", logs: logs, sub: sub}, nil
}

// WatchModeratorSet is a free log subscription operation binding the contract event 0x07abfa597e8dcfbbfa28a386fa0728caf3360f382d92356232125424bcaa6a53.
//
// Solidity: event ModeratorSet(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) WatchModeratorSet(opts *bind.WatchOpts, sink chan<- *PaymentchannelModeratorSet) (event.Subscription, error) {

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "ModeratorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelModeratorSet)
				if err := _Paymentchannel.contract.UnpackLog(event, "ModeratorSet", log); err != nil {
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

// ParseModeratorSet is a log parse operation binding the contract event 0x07abfa597e8dcfbbfa28a386fa0728caf3360f382d92356232125424bcaa6a53.
//
// Solidity: event ModeratorSet(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) ParseModeratorSet(log types.Log) (*PaymentchannelModeratorSet, error) {
	event := new(PaymentchannelModeratorSet)
	if err := _Paymentchannel.contract.UnpackLog(event, "ModeratorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelModeratorUnsetIterator is returned from FilterModeratorUnset and is used to iterate over the raw logs and unpacked data for ModeratorUnset events raised by the Paymentchannel contract.
type PaymentchannelModeratorUnsetIterator struct {
	Event *PaymentchannelModeratorUnset // Event containing the contract specifics and raw log

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
func (it *PaymentchannelModeratorUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelModeratorUnset)
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
		it.Event = new(PaymentchannelModeratorUnset)
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
func (it *PaymentchannelModeratorUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelModeratorUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelModeratorUnset represents a ModeratorUnset event raised by the Paymentchannel contract.
type PaymentchannelModeratorUnset struct {
	Moderator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModeratorUnset is a free log retrieval operation binding the contract event 0x6e69dd57dae2d984f9d5b55a45dac79a60c83270d517cfd720af1eda0dd2d09b.
//
// Solidity: event ModeratorUnset(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) FilterModeratorUnset(opts *bind.FilterOpts) (*PaymentchannelModeratorUnsetIterator, error) {

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "ModeratorUnset")
	if err != nil {
		return nil, err
	}
	return &PaymentchannelModeratorUnsetIterator{contract: _Paymentchannel.contract, event: "ModeratorUnset", logs: logs, sub: sub}, nil
}

// WatchModeratorUnset is a free log subscription operation binding the contract event 0x6e69dd57dae2d984f9d5b55a45dac79a60c83270d517cfd720af1eda0dd2d09b.
//
// Solidity: event ModeratorUnset(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) WatchModeratorUnset(opts *bind.WatchOpts, sink chan<- *PaymentchannelModeratorUnset) (event.Subscription, error) {

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "ModeratorUnset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelModeratorUnset)
				if err := _Paymentchannel.contract.UnpackLog(event, "ModeratorUnset", log); err != nil {
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

// ParseModeratorUnset is a log parse operation binding the contract event 0x6e69dd57dae2d984f9d5b55a45dac79a60c83270d517cfd720af1eda0dd2d09b.
//
// Solidity: event ModeratorUnset(address _moderator)
func (_Paymentchannel *PaymentchannelFilterer) ParseModeratorUnset(log types.Log) (*PaymentchannelModeratorUnset, error) {
	event := new(PaymentchannelModeratorUnset)
	if err := _Paymentchannel.contract.UnpackLog(event, "ModeratorUnset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paymentchannel contract.
type PaymentchannelOwnershipTransferredIterator struct {
	Event *PaymentchannelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentchannelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelOwnershipTransferred)
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
		it.Event = new(PaymentchannelOwnershipTransferred)
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
func (it *PaymentchannelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelOwnershipTransferred represents a OwnershipTransferred event raised by the Paymentchannel contract.
type PaymentchannelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Paymentchannel *PaymentchannelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*PaymentchannelOwnershipTransferredIterator, error) {

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &PaymentchannelOwnershipTransferredIterator{contract: _Paymentchannel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Paymentchannel *PaymentchannelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentchannelOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelOwnershipTransferred)
				if err := _Paymentchannel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Paymentchannel *PaymentchannelFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentchannelOwnershipTransferred, error) {
	event := new(PaymentchannelOwnershipTransferred)
	if err := _Paymentchannel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelSourceProofSubmittedIterator is returned from FilterSourceProofSubmitted and is used to iterate over the raw logs and unpacked data for SourceProofSubmitted events raised by the Paymentchannel contract.
type PaymentchannelSourceProofSubmittedIterator struct {
	Event *PaymentchannelSourceProofSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentchannelSourceProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelSourceProofSubmitted)
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
		it.Event = new(PaymentchannelSourceProofSubmitted)
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
func (it *PaymentchannelSourceProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelSourceProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelSourceProofSubmitted represents a SourceProofSubmitted event raised by the Paymentchannel contract.
type PaymentchannelSourceProofSubmitted struct {
	ChannelId        [32]byte
	RecoveredAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSourceProofSubmitted is a free log retrieval operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) FilterSourceProofSubmitted(opts *bind.FilterOpts, _channelId [][32]byte, _recoveredAddress []common.Address) (*PaymentchannelSourceProofSubmittedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "SourceProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelSourceProofSubmittedIterator{contract: _Paymentchannel.contract, event: "SourceProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchSourceProofSubmitted is a free log subscription operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) WatchSourceProofSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentchannelSourceProofSubmitted, _channelId [][32]byte, _recoveredAddress []common.Address) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "SourceProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelSourceProofSubmitted)
				if err := _Paymentchannel.contract.UnpackLog(event, "SourceProofSubmitted", log); err != nil {
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

// ParseSourceProofSubmitted is a log parse operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Paymentchannel *PaymentchannelFilterer) ParseSourceProofSubmitted(log types.Log) (*PaymentchannelSourceProofSubmitted, error) {
	event := new(PaymentchannelSourceProofSubmitted)
	if err := _Paymentchannel.contract.UnpackLog(event, "SourceProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentchannelTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the Paymentchannel contract.
type PaymentchannelTokensWithdrawnIterator struct {
	Event *PaymentchannelTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentchannelTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentchannelTokensWithdrawn)
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
		it.Event = new(PaymentchannelTokensWithdrawn)
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
func (it *PaymentchannelTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentchannelTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentchannelTokensWithdrawn represents a TokensWithdrawn event raised by the Paymentchannel contract.
type PaymentchannelTokensWithdrawn struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Paymentchannel *PaymentchannelFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, _tokenAddress []common.Address) (*PaymentchannelTokensWithdrawnIterator, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.FilterLogs(opts, "TokensWithdrawn", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentchannelTokensWithdrawnIterator{contract: _Paymentchannel.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Paymentchannel *PaymentchannelFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentchannelTokensWithdrawn, _tokenAddress []common.Address) (event.Subscription, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Paymentchannel.contract.WatchLogs(opts, "TokensWithdrawn", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentchannelTokensWithdrawn)
				if err := _Paymentchannel.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Paymentchannel *PaymentchannelFilterer) ParseTokensWithdrawn(log types.Log) (*PaymentchannelTokensWithdrawn, error) {
	event := new(PaymentchannelTokensWithdrawn)
	if err := _Paymentchannel.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
