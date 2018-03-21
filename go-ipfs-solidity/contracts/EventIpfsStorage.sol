pragma solidity 0.4.21;

/*
	Test contract use to explore storing 
*/
contract EventIpfsStorage {

	event StringStorage(string _a);

	function EmitStringStorageEvent(string _a) public returns (bool) {
		emit StringStorage(_a);
		return true;
	}

}