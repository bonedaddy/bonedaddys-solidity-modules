pragma solidity 0.4.20;


contract AddressGenerationCalculator {

	function calculateAddressForNonce(
		address _creator,
		uint8   _nonce)
		public
		pure
		returns (address)
	{
		require(_nonce < 127);
		return address(keccak256(0xd6, 0x94, _creator, _nonce));
	}
}
