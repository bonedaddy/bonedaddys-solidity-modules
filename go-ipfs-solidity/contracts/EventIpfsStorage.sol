pragma solidity 0.4.21;

/*
	Test contract use to explore storing 
*/
contract EventIpfsStorage {

	event TextData(string _name, string _contents);

	function StoreTextData(
		string _name,
		string _contents)
		public
		returns (bool) 
	{
		emit TextData(_name, _contents);
		return true;
	}

}