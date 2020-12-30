package bclient

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	// my personal address
	myAddress = common.HexToAddress("0x5a361A1dfd52538A158e352d21B5b622360a7C13")
)

func TestBClient(t *testing.T) {
	infuraAPIKey := os.Getenv("INFURA_API_KEY")
	client, err := NewInfuraClient(infuraAPIKey, false)
	require.NoError(t, err)
	t.Run("DEFI5", func(t *testing.T) {
		defi5, err := client.DEFI5()
		require.NoError(t, err)

		max, err := defi5.GetMaxPoolTokens(nil)
		require.NoError(t, err)
		t.Log("max pool tokens: ", max)

		bal, err := defi5.BalanceOf(nil, myAddress)
		require.NoError(t, err)
		require.Greater(t, bal.Int64(), int64(0))
		t.Log("balance of: ", bal)

		dec, err := defi5.Decimals(nil)
		require.NoError(t, err)
		t.Log("decimals: ", dec)

		cont, err := defi5.GetController(nil)
		require.NoError(t, err)
		t.Log("controller: ", cont)

		l, _ := BalanceOfDecimal(defi5, myAddress)
		t.Log("balance decimal: ", l)

		t.Run("DEFI5_Staking", func(t *testing.T) {
			stake, err := client.StakingAt(DEFI5StakingAddress)
			require.NoError(t, err)
			bal, err := StakeBalanceOf(stake, defi5, myAddress)
			require.NoError(t, err)
			t.Log("tokens staked: ", bal)
			earned, err := StakeEarned(stake, defi5, myAddress)
			require.NoError(t, err)
			t.Log("tokens earned: ", earned)
		})

		t.Run("DEFI5_MCAP_Controller", func(t *testing.T) {
			_, err := client.MCAPControllerAt(defi5)
			require.NoError(t, err)
		})
		t.Run("DEFI5_UNISWAP_ORACLE", func(t *testing.T) {
			addr, err := client.OracleFor(defi5)
			require.NoError(t, err)
			t.Log("defi5 uniswap oracle address: ", addr)
		})
	})

}
