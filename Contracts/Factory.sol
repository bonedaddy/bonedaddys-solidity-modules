pragma solidity 0.4.20;

contract Factory is Administration {

	// replace this with the compiled bytecode fro the contract you wish to produce
	// we mark as constant and private to drasticalyl reduce deployment gas costs
	bytes constant private code = hex"..";

	function deployToken() external {
		bytes memory _code = code;
		address a;
		assembly {
			a := create(0, add(_code, 0x20), mload(_code))
		}
		NewContract(a);
		tokens.push(a);
	}

}