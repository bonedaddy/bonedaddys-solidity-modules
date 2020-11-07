pragma solidity >=0.4.24 <=0.8.0;

contract RuntimeFactory {

	function deployToken(bytes memory _code) public returns (address) {
		address a;
		assembly {
			a := create(0, add(_code, 0x20), mload(_code))
		}
        return a;
	}

}