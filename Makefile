.PHONY: cli
cli:
	go build -o gondx ./cmd/gondx

.PHONY: all
all: copy-abi gen-bindings

.PHONY: copy-abi
copy-abi:
	if [ ! -d abi ]; then mkdir abi ; fi
	cp -r ./indexed-js/src/abi/*.json abi
	cp indexed-core/abi/MarketCapSqrtController.json abi
	cp uniswap-v2-core/abi/IUniswap*.json abi

.PHONY: gen-bindings
gen-bindings:
	if [ ! -d bindings ]; then mkdir bindings ; fi
	if [ ! -d bindings/pool ]; then mkdir bindings/pool ; fi
	if [ ! -d bindings/pair ]; then mkdir bindings/pair ; fi
	if [ ! -d bindings/staking_rewards ]; then mkdir bindings/staking_rewards ; fi
	if [ ! -d bindings/uniswapv2_oracle ]; then mkdir bindings/uniswapv2_oracle ; fi
	if [ ! -d bindings/marketcap_sqrt_controller ]; then mkdir bindings/marketcap_sqrt_controller ; fi
	if [ ! -d bindings/erc20 ]; then mkdir bindings/erc20 ; fi
	abigen --abi abi/IPool.json --pkg poolbindings --out bindings/pool/bindings.go
	abigen --abi abi/Pair.json --pkg pairbindings --out bindings/pair/bindings.go
	abigen --abi abi/StakingRewards.json --pkg stakingbindings --out bindings/staking_rewards/bindings.go
	abigen --abi abi/IIndexedUniswapV2Oracle.json --pkg uv2oraclebindings --out bindings/uniswapv2_oracle/bindings.go
	abigen --abi abi/MarketCapSqrtController.json --pkg mcapscontroller --out bindings/marketcap_sqrt_controller/bindings.go
	abigen --abi abi/IERC20.json --pkg erc20 --out bindings/erc20/bindings.go