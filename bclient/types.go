package bclient

import (
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
)

// to compute addresses for contracts go to https://etherscan.io/address/0xf00a38376c8668fc1f3cd3daeef42e0e44a7fcdb#readProxyContract

var (
	defi5Addr      = "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41"
	defi5StakeAddr = "0x11bf850D1B85eA02eF9F06Cf09488E443655b586"
	// ETH-DEFI5 stake contract
	defi5UNILPStakeAddr = "0x04834fE50BE6954088d979A84bE20913efCA9118"
	// ETH-CC10 stake contract
	cc10UNILPStakeAddr = "0x08F0d214d6A80520F18E6837e2cc40899E67B5c0"
	cc10Addr           = "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3"
	cc10StakeAddr      = "0xE1dE5Ebd607f1Da1c34D78b760B9C87901d0Ba35"
	orcl5Addr          = "0xD6cb2aDF47655B1bABdDc214d79257348CBC39A7"
	degen10Addr        = "0x126c121f99e1E211dF2e5f8De2d96Fa36647c855"
	nftpAddr           = "0x68bB81B3F67f7AAb5fd1390ECB0B8e1a806F2465"
	errorAddr          = "0xd3dEFf001ef67E39212F4973B617C2E684fa436C"
	fffAddr            = "0xaBAfA52D3d5A2c18A4C1Ae24480D22B831fC0413"
	// DEFI5TokenAddress is the address of the DEFI5 token/pool contract
	DEFI5TokenAddress = common.HexToAddress(defi5Addr)
	// DEFI5StakingAddress is the address of the DEFI5 staking contract
	DEFI5StakingAddress = common.HexToAddress(defi5StakeAddr)
	// CC10StakingAddress is the address of the CC10 staking contract
	CC10StakingAddress = common.HexToAddress(cc10StakeAddr)
	// DEFI5UNILPStakingAddress is the address of the ETH-DEFI5 Uniswap staking contract
	DEFI5UNILPStakingAddress = common.HexToAddress(defi5UNILPStakeAddr)
	// CC10UNILPStakingAddress is the address of the ETH-CC10 uniswap staking contract
	CC10UNILPStakingAddress = common.HexToAddress(cc10UNILPStakeAddr)
	// CC10TokenAddress is the address of the CC10 token/pool contract
	CC10TokenAddress = common.HexToAddress(cc10Addr)
	// ORCL5TokenAddress is the address of the ORCL5 token/pool contract
	ORCL5TokenAddress = common.HexToAddress(orcl5Addr)
	// DEGEN10TokenAddress is the address of the DEGEN10 token contract
	DEGEN10TokenAddress = common.HexToAddress(degen10Addr)
	// NFTPTokenAddress is the address of the nftp token contract
	NFTPTokenAddress = common.HexToAddress(nftpAddr)
	// ERRORTokenAddress is the address of the error token contract
	ERRORTokenAddress = common.HexToAddress(errorAddr)
	/// FFFTokenAddress is the address of the fff token contract
	FFFTokenAddress = common.HexToAddress(fffAddr)
	// WETHTokenAddress is the address of the WETH token contract
	WETHTokenAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	// DAITokenAddress is the address of the MCD (Multi Collateral DAI) contract
	DAITokenAddress = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	// NDXTokenAddress is the address of the NDX contract
	NDXTokenAddress = common.HexToAddress("0x86772b1409b61c639eaac9ba0acfbb6e238e5f83")
	// GovernorAlpha is the address of the GovernorALpha governance contract
	GovernorAlpha = common.HexToAddress("0x95129751769f99CC39824a0793eF4933DD8Bb74B")
	// InfuraWSURL is the URL for INFURA websockets access
	InfuraWSURL = "wss://mainnet.infura.io/ws/v3/"
	// InfuraHTTPURL is the URL for INFURA HTTP access
	InfuraHTTPURL = "https://mainnet.infura.io/v3/"
	// IndexPools contains all current available idnex pools
	// note that it may not be truly up-to-date
	IndexPools = map[string]common.Address{"defi5": DEFI5TokenAddress, "cc10": CC10TokenAddress}
)

// DEFI5 is the DEFI Top 5 IndexPool
type DEFI5 struct {
	*poolbindings.Poolbindings
}
