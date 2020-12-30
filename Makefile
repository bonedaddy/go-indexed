.PHONY: copy-abi
copy-abi:
	if [ ! -d abi ]; then mkdir abi ; fi
	cp -r ./indexed-js/src/abi/*.json abi

.PHONY: gen-bindings
gen-bindings:
	if [ ! -d bindings ]; then mkdir bindings ; fi
	if [ ! -d bindings/pool ]; then mkdir bindings/pool ; fi
	if [ ! -d bindings/pair ]; then mkdir bindings/pair ; fi
	if [ ! -d bindings/staking_rewards ]; then mkdir bindings/staking_rewards ; fi
	if [ ! -d bindings/uniswapv2_oracle ]; then mkdir bindings/uniswapv2_oracle ; fi
	abigen --pkg poolbindings --out bindings/pool/bindings.go
	abigen --pkg pairbindings --out bindings/pair/bindings.go
	abigen --pkg stakingbindings --out bindings/staking_rewards/bindings.go
	abigen --pkg uv2oraclebindings --out bindings/uniswapv2_oracle/bindings.go
