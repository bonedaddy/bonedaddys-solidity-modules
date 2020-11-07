pragma solidity >=0.4.24 <=0.8.0;

import "./utils/Bytes.sol";
import "./interfaces/ERC20I.sol";
import "./RuntimeFactory.sol";

contract ERC20Factory is ByteUtils, RuntimeFactory {

    bytes constant private TOKENCODE = hex"608060405234801561001057600080fd5b50604051610a5c380380610a5c8339810160409081528151602080840151928401516060850151608086015160008054600160a060020a031916600160a060020a03871617905560038690556004805460ff191660ff85161790559086018051949692949093919092019161008b9160019190850190610106565b50805161009f906002906020840190610106565b50600160a060020a038516600081815260056020908152604080832088905580519283529082019290925280820186905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a150505050506101a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014757805160ff1916838001178555610174565b82800160010185558215610174579182015b82811115610174578251825591602001919060010190610159565b50610180929150610184565b5090565b61019e91905b80821115610180576000815560010161018a565b90565b6108ac806101b06000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a8578063095ea7b31461013257806318160ddd1461016a57806323b872dd1461019157806327e235e3146101bb578063313ce567146101dc5780635c658165146102075780638da5cb5b1461022e57806395d89b411461025f578063a9059cbb14610274575b600080fd5b3480156100b457600080fd5b506100bd610298565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f75781810151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013e57600080fd5b50610156600160a060020a0360043516602435610325565b604080519115158252519081900360200190f35b34801561017657600080fd5b5061017f61041f565b60408051918252519081900360200190f35b34801561019d57600080fd5b50610156600160a060020a0360043581169060243516604435610425565b3480156101c757600080fd5b5061017f600160a060020a036004351661066e565b3480156101e857600080fd5b506101f1610680565b6040805160ff9092168252519081900360200190f35b34801561021357600080fd5b5061017f600160a060020a0360043581169060243516610689565b34801561023a57600080fd5b506102436106a6565b60408051600160a060020a039092168252519081900360200190f35b34801561026b57600080fd5b506100bd6106b5565b34801561028057600080fd5b50610156600160a060020a036004351660243561070d565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561031d5780601f106102f25761010080835404028352916020019161031d565b820191906000526020600020905b81548152906001019060200180831161030057829003601f168201915b505050505081565b600080821161037e576040805160e560020a62461bcd02815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e2030000000604482015290519081900360640190fd5b336000908152600660209081526040808320600160a060020a03871684529091529020546103b2908363ffffffff61085216565b336000818152600660209081526040808320600160a060020a0389168085529083529281902094909455835192835282015280820184905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a150600192915050565b60035481565b600160a060020a0383166000908152600560205260408120548211156104bb576040805160e560020a62461bcd02815260206004820152602160248201527f6f776e657220646f6573206e6f74206861766520656e6f75676820746f6b656e60448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a038416600090815260066020908152604080832033845290915290205482111561055c576040805160e560020a62461bcd02815260206004820152602560248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820616c6c6f60448201527f77616e6365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600660209081526040808320338452909152902054610590908363ffffffff61086b16565b600160a060020a0385166000818152600660209081526040808320338452825280832094909455918152600590915220546105d1908363ffffffff61086b16565b600160a060020a038086166000908152600560205260408082209390935590851681522054610606908363ffffffff61085216565b600160a060020a0380851660008181526005602090815260409182902094909455805192881683529282015280820184905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a15060019392505050565b60056020526000908152604090205481565b60045460ff1681565b600660209081526000928352604080842090915290825290205481565b600054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561031d5780601f106102f25761010080835404028352916020019161031d565b3360009081526005602052604081205482111561079a576040805160e560020a62461bcd02815260206004820152602260248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820746f6b6560448201527f6e73000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600560205260409020546107ba908363ffffffff61086b16565b3360009081526005602052604080822092909255600160a060020a038516815220546107ec908363ffffffff61085216565b600160a060020a03841660008181526005602090815260409182902093909355805133815292830191909152818101849052517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a150600192915050565b60008282018381101561086457600080fd5b9392505050565b60008282111561087a57600080fd5b509003905600a165627a7a7230582025bc031d1ee00a03299cef0075e167cb75cb0c8b6d02b38db025f337638628910029";

    address[] public generatedTokens;

    struct Tokens {
        bytes32 nameHash;
        bytes32 symbolHash;
        address owner;
        uint256 supply;
        uint8 decimals;
    }

    mapping (address => Tokens) public tokens;

    function deployToken(address _owner, uint256 _supply, uint8 _decimals, string memory _name, string memory _symbol) public returns (bool) {
        //bytes memory _code = TOKENCODE;
        bytes memory args = abi.encode(_owner, _supply, _decimals, _name, _symbol);
        bytes memory code = concat(TOKENCODE, args);
        address a = deployToken(code);
        require(a != address(0), "failed to deploy contract, likely ran out of gas");
        generatedTokens.push(a);
        Tokens memory t = Tokens({
            nameHash: keccak256(abi.encodePacked(_name)),
            symbolHash: keccak256(abi.encodePacked(_symbol)),
            owner: _owner,
            supply: _supply,
            decimals: _decimals
        });
        tokens[a] = t;
        return true;
    }

}