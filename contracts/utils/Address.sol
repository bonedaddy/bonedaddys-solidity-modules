pragma solidity >=0.4.24 <=0.8.0;

contract Address {

    // from https://ethereum.stackexchange.com/questions/24248/how-to-calculate-an-ethereum-contracts-address-during-its-creation-using-the-so
    function addressFromNonce(address _origin, uint _nonce) public pure returns (address _address) {
        bytes memory data;
        if(_nonce == 0x00) {
            data = abi.encodePacked(byte(0xd6), byte(0x94), _origin, byte(0x80));
        } else if(_nonce <= 0x7f) {
            data = abi.encodePacked(byte(0xd6), byte(0x94), _origin, uint8(_nonce));
        } else if(_nonce <= 0xff)  {
            data = abi.encodePacked(byte(0xd7), byte(0x94), _origin, byte(0x81), uint8(_nonce));
        } else if(_nonce <= 0xffff) {
            data = abi.encodePacked(byte(0xd8), byte(0x94), _origin, byte(0x82), uint16(_nonce));
        } else if(_nonce <= 0xffffff) {
            data = abi.encodePacked(byte(0xd9), byte(0x94), _origin, byte(0x83), uint24(_nonce));
        } else {
            data = abi.encodePacked(byte(0xda), byte(0x94), _origin, byte(0x84), uint32(_nonce));
        }
        bytes32 hash = keccak256(data);
        assembly {
            mstore(0, hash)
            _address := mload(0)
        }
    }

}
