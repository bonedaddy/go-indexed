package bclient

import (
	"math/big"

	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
)

var (
	defi5Addr      = "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41"
	defi5StakeAddr = "0x11bf850D1B85eA02eF9F06Cf09488E443655b586"
	// DEFI5TokenAddress is the address of the DEFI5 token contract
	DEFI5TokenAddress = common.HexToAddress(defi5Addr)
	// DEFI5StakingAddress is the address of the DEFI5 staking contract
	DEFI5StakingAddress = common.HexToAddress(defi5StakeAddr)
	INFURA_WS_URL       = "wss://mainnet.infura.io/ws/v3/"
	INFURA_HTTP_URL     = "https://mainnet.infura.io/v3/"
	OneEthInWei         = big.NewInt(1000000000000000000)
)

// DEFI5 is the DEFI Top 5 IndexPool
type DEFI5 struct {
	*poolbindings.Poolbindings
}
