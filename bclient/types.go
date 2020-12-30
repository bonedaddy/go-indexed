package bclient

import (
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
)

var (
	defi5Addr      = "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41"
	defi5StakeAddr = "0x11bf850D1B85eA02eF9F06Cf09488E443655b586"
	cc10Addr       = "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3"
	// DEFI5TokenAddress is the address of the DEFI5 token/pool contract
	DEFI5TokenAddress = common.HexToAddress(defi5Addr)
	// DEFI5StakingAddress is the address of the DEFI5 staking contract
	DEFI5StakingAddress = common.HexToAddress(defi5StakeAddr)
	// CC10TokenAddress is the address of the CC10 token/pool contract
	CC10TokenAddress = common.HexToAddress(cc10Addr)
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
