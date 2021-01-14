GIT_VERSION=`git describe --tags`

.PHONY: cli
cli:
	go build -o gondx -ldflags "-X main.Version=$(GIT_VERSION)" ./cmd/gondx

.PHONY: all
all: copy-abi gen-bindings

.PHONY: copy-abi
copy-abi:
	if [ ! -d abi ]; then mkdir abi ; fi
	cp -r ./indexed-js/src/abi/*.json abi
	cp indexed-core/abi/MarketCapSqrtController.json abi
	#cp uniswap-v2-core/abi/IUniswap*.json abi
	cp indexed-governance/abi/GovernorAlpha.json abi

.PHONY: gen-bindings
gen-bindings:
	if [ ! -d bindings ]; then mkdir bindings ; fi
	if [ ! -d bindings/pool ]; then mkdir bindings/pool ; fi
	if [ ! -d bindings/pair ]; then mkdir bindings/pair ; fi
	if [ ! -d bindings/staking_rewards ]; then mkdir bindings/staking_rewards ; fi
	if [ ! -d bindings/uniswapv2_oracle ]; then mkdir bindings/uniswapv2_oracle ; fi
	if [ ! -d bindings/marketcap_sqrt_controller ]; then mkdir bindings/marketcap_sqrt_controller ; fi
	if [ ! -d bindings/erc20 ]; then mkdir bindings/erc20 ; fi
	if [ ! -d bindings/uniswapv2 ]; then mkdir bindings/uniswapv2 ; fi
	if [ ! -d bindings/governor_alpha ]; then mkdir bindings/governor_alpha ; fi
	abigen --abi abi/IPool.json --pkg poolbindings --out bindings/pool/bindings.go
	abigen --abi abi/Pair.json --pkg pairbindings --out bindings/pair/bindings.go
	abigen --abi abi/StakingRewards.json --pkg stakingbindings --out bindings/staking_rewards/bindings.go
	abigen --abi abi/IIndexedUniswapV2Oracle.json --pkg uv2oraclebindings --out bindings/uniswapv2_oracle/bindings.go
	abigen --abi abi/MarketCapSqrtController.json --pkg mcapscontroller --out bindings/marketcap_sqrt_controller/bindings.go
	abigen --abi abi/IERC20.json --pkg erc20 --out bindings/erc20/bindings.go
	abigen --abi abi/IUniswapV2Pair.json --pkg uniswapv2pair --out bindings/uniswapv2/pair/v2pair.go
	abigen --abi abi/IUniswapV2ERC20.json --pkg uniswapv2erc20 --out bindings/uniswapv2/erc20/v2erc20.go
	abigen --abi abi/IUniswapV2Factory.json --pkg uniswapv2factory --out bindings/uniswapv2/factory/v2factory.go
	abigen --abi abi/IUniswapV2Callee.json --pkg uniswapv2callee --out bindings/uniswapv2/callee/v2callee.go
	abigen --abi abi/GovernorAlpha.json --pkg governoralpha --out bindings/governor_alpha/bindings.go

.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(GIT_VERSION) -t bonedaddy/gondx:$(GIT_VERSION) .
	docker image tag bonedaddy/gondx:$(GIT_VERSION) bonedaddy/gondx:latest

.PHONY: release
release:
	.scripts/release.sh

.PHONY: testenv
testenv:
	docker-compose up -d

.PHONY: testenv-clean
testenv-clean:
	docker-compose up -d --force-recreate


.PHONY: db-backup
db-backup:
	.scripts/db_backup.sh