package bclient

import (
	"math/big"
	"os"
	"testing"

	"github.com/bonedaddy/go-indexed/utils"
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
			stake, err := client.StakingAt("defi5")
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
			orc, err := client.OracleAt(defi5)
			require.NoError(t, err)
			tokens, err := client.PoolTokensFor(defi5)
			require.NoError(t, err)
			for name, addr := range tokens {
				// get time weighted average price of token in terms of weth
				price, err := orc.ComputeAverageTokenPrice(nil, addr, big.NewInt(0), big.NewInt(36288000)) // price between 0 seconds and 7 days
				require.NoError(t, err)
				t.Logf("token %s (%s) average eth price %s", addr, name, price.X)
			}
		})
		t.Run("Uniswap", func(t *testing.T) {
			t.Log("NDX-ETH uniswap pair address: ", client.NdxEthPairAddress())
			t.Log("DEFI5-ETH uniswap pair address: ", client.Defi5EthPairAddress())
			t.Log("CC10-ETH uniswap pair address: ", client.Cc10EthPairAddress())
			var args = map[string]bool{
				"ndx-eth":    false,
				"eth-ndx":    false,
				"ndxx-eth":   true,
				"cc10-eth":   false,
				"eth-cc10":   false,
				"ethh-cc10":  true,
				"defi5-eth":  false,
				"eth-defi5":  false,
				"deffi5-eth": true,
			}
			for name, wantErr := range args {
				t.Run(name, func(t *testing.T) {
					amt, err := client.ExchangeAmount(utils.ToWei(10.0, 18), name) // get amount for swaping
					if (err != nil) != wantErr {
						t.Errorf("%s wantErr %v", name, wantErr)
					}
					if wantErr {
						return
					}
					t.Logf("swapping 10 %s costs: %s", name, amt)
				})
			}
		})
	})

}
