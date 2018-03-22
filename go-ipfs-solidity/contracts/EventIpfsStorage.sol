pragma solidity 0.4.21;

/*
	Test contract use to explore storing 
*/
contract EventIpfsStorage {

	event TextData(string _a, string _b);

	function EmitStringStorageEvent() public returns (bool) {
		emit TextData("abc", "def");
		return true;
	}

}