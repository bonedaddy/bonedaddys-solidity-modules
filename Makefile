.PHONY: compile
compile:
	./scripts/solc_compile.sh \
		contracts/ERC20Factory.sol \
		bin/erc20_factory
	./scripts/solc_compile.sh \
		contracts/RuntimeFactory.sol \
			bin/runtime_factory
	./scripts/solc_compile.sh \
		contracts/EthErc20PaymentChannel.sol \
		bin/paymentchannel

.PHONY: abigen
abigen:
	./scripts/abigen.sh \
		bin/paymentchannel/PaymentChannels.bin \
		bin/paymentchannel/PaymentChannels.abi \
		paymentchannel \
		bindings/paymentchannel/bindings.go
	./scripts/abigen.sh \
		bin/runtime_factory/RuntimeFactory.bin \
		bin/runtime_factory/RuntimeFactory.abi \
		runtimefactory \
		bindings/runtimefactory/bindings.go
	./scripts/abigen.sh \
		bin/erc20_factory/ERC20Factory.bin \
		bin/erc20_factory/ERC20Factory.abi \
		erc20factory \
		bindings/erc20factory/bindings.go