package bclient

import (
	"github.com/ethereum/go-ethereum/common"
)

// to compute addresses for contracts go to https://etherscan.io/address/0xf00a38376c8668fc1f3cd3daeef42e0e44a7fcdb#readProxyContract

var (

	// InfuraWSURL is the URL for INFURA websockets access
	InfuraWSURL = "wss://mainnet.infura.io/ws/v3/"
	// InfuraHTTPURL is the URL for INFURA HTTP access
	InfuraHTTPURL = "https://mainnet.infura.io/v3/"
)

// staking addresses
var (
	// DEFI5StakingAddress is the address of the DEFI5 staking contract
	DEFI5StakingAddress = common.HexToAddress("0x11bf850D1B85eA02eF9F06Cf09488E443655b586")
	// CC10StakingAddress is the address of the CC10 staking contract
	CC10StakingAddress = common.HexToAddress("0xE1dE5Ebd607f1Da1c34D78b760B9C87901d0Ba35")
	// DEFI5UNILPStakingAddress is the address of the ETH-DEFI5 Uniswap staking contract
	DEFI5UNILPStakingAddress = common.HexToAddress("0x04834fE50BE6954088d979A84bE20913efCA9118")
	// CC10UNILPStakingAddress is the address of the ETH-CC10 uniswap staking contract
	CC10UNILPStakingAddress = common.HexToAddress("0x08F0d214d6A80520F18E6837e2cc40899E67B5c0")
	// DEGENUNILPStakingAddress is  the address of the ETH-DEGEN uniswap staking contract
	DEGENUNILPStakingAddress = common.HexToAddress("0x649A94e2b3010338508CF50865BaafB1FA07A32c")
)

// token addresses
var (
	// DEFI5TokenAddress is the address of the DEFI5 token/pool contract
	DEFI5TokenAddress = common.HexToAddress("0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41")
	// CC10TokenAddress is the address of the CC10 token/pool contract
	CC10TokenAddress = common.HexToAddress("0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3")
	// ORCL5TokenAddress is the address of the ORCL5 token/pool contract
	ORCL5TokenAddress = common.HexToAddress("0xD6cb2aDF47655B1bABdDc214d79257348CBC39A7")
	// DEGEN10TokenAddress is the address of the DEGEN10 token contract
	DEGEN10TokenAddress = common.HexToAddress("0x126c121f99e1E211dF2e5f8De2d96Fa36647c855")
	// WETHTokenAddress is the address of the WETH token contract
	WETHTokenAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	// DAITokenAddress is the address of the MCD (Multi Collateral DAI) contract
	DAITokenAddress = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	// NDXTokenAddress is the address of the NDX contract
	NDXTokenAddress = common.HexToAddress("0x86772b1409b61c639eaac9ba0acfbb6e238e5f83")
)

// misc addresses
var (
	// GovernorAlpha is the address of the GovernorALpha governance contract
	GovernorAlpha = common.HexToAddress("0x95129751769f99CC39824a0793eF4933DD8Bb74B")
)

// uniswap lp tokens
var (
	// DEFI5LP is the address of the uniswap pair
	DEFI5LP = common.HexToAddress("0x8dCBa0B75c1038c4BaBBdc0Ff3bD9a8f6979Dd13")
	// CC10LP is the address of the uniswap pair
	CC10LP = common.HexToAddress("0x2701eA55b8B4f0FE46C15a0F560e9cf0C430f833")
	// DEGENLP is the address of the uniswap pair
	DEGENLP = common.HexToAddress("0xFaAD1072E259B5ED342D3f16277477B46D379ABC")
	// ORCL5 LP is the address of the uniswap pair
	ORCL5LP = common.HexToAddress("0xf58d2bacbc68c587730ea0ce5131f6ae7c622a5d")
	// NDXLP is the address of the uniswap pair
	NDXLP = common.HexToAddress("0x46af8ac1b82f73db6aacc1645d40c56191ab787b")
)
