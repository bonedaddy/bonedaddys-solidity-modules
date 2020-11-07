.PHONY: abigen
abigen:
	./scripts/solc_compile.sh contracts/ERC20Factory.sol bin/erc20_factory
	./scripts/solc_compile.sh contracts/RuntimeFactory.sol bin/runtime_factory
	./scripts/solc_compile.sh contracts/EthErc20PaymentChannel.sol bin/eth_erc20_payment_channel