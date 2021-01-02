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

func doSetup(t *testing.T) *Client {
	infuraAPIKey := os.Getenv("INFURA_API_KEY")
	if infuraAPIKey == "" {
		t.Fatal("INFURA_API_KEY env var is empty")
	}
	client, err := NewInfuraClient(infuraAPIKey, false)
	require.NoError(t, err)
	return client

}

func TestBClient(t *testing.T) {
	client := doSetup(t)
	t.Cleanup(func() {
		client.Close()
	})
	t.Run("Misc", func(t *testing.T) {
		_, err := client.CurrentBlock()
		require.NoError(t, err)
		require.NotNil(t, client.Uniswap())
		require.Equal(t, "Maker", guessTokenName("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"))
		require.Equal(t, "MKR", guessTokenSymbol("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"))
	})
	t.Run("CC10", func(t *testing.T) {
		_, err := client.CC10()
		require.NoError(t, err)
	})
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

		l, err := BalanceOfDecimal(defi5, myAddress)
		require.NoError(t, err)
		t.Log("balance decimal: ", l)

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
	})

}

func TestStaking(t *testing.T) {
	client := doSetup(t)
	t.Cleanup(func() {
		client.Close()
	})
	t.Run("DEFI5_Staking", func(t *testing.T) {
		defi5, err := client.DEFI5()
		require.NoError(t, err)
		stake, err := client.StakingAt("defi5")
		require.NoError(t, err)
		bal, err := StakeBalanceOf(stake, defi5, myAddress)
		require.NoError(t, err)
		t.Log("tokens staked: ", bal)
		earned, err := StakeEarned(stake, defi5, myAddress)
		require.NoError(t, err)
		t.Log("tokens earned: ", earned)
		_, err = client.StakingAt("univ2-eth-defi5")
		require.NoError(t, err)
	})
}

func TestUniswap(t *testing.T) {
	client := doSetup(t)
	t.Cleanup(func() {
		client.Close()
	})
	t.Run("NdxEthPairAddress", func(t *testing.T) {
		addr := client.NdxEthPairAddress()
		require.Equal(t, "0x46af8AC1b82F73dB6Aacc1645D40c56191ab787b", addr.String())
	})
	t.Run("Defi5EthPairAddress", func(t *testing.T) {
		addr := client.Defi5EthPairAddress()
		require.Equal(t, "0x8dCBa0B75c1038c4BaBBdc0Ff3bD9a8f6979Dd13", addr.String())
	})
	t.Run("Cc10EthPairAddress", func(t *testing.T) {
		addr := client.Cc10EthPairAddress()
		require.Equal(t, "0x2701eA55b8B4f0FE46C15a0F560e9cf0C430f833", addr.String())
	})
	t.Run("NdxDaiPrice", func(t *testing.T) {
		price, err := client.NdxDaiPrice()
		require.NoError(t, err)
		require.Greater(t, price, float64(0))
	})
	t.Run("Cc10DaiPrice", func(t *testing.T) {
		price, err := client.Cc10DaiPrice()
		require.NoError(t, err)
		require.Greater(t, price, float64(0))
	})
	t.Run("Defi5DaiPrice", func(t *testing.T) {
		price, err := client.Defi5DaiPrice()
		require.NoError(t, err)
		require.Greater(t, price, float64(0))
	})
	t.Run("EthDaiPrice", func(t *testing.T) {
		cost, err := client.EthDaiPrice()
		require.NoError(t, err)
		require.Greater(t, cost.Int64(), big.NewInt(0).Int64())
	})
	t.Run("ExchangeAmount", func(t *testing.T) {
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
	t.Run("Reserves", func(t *testing.T) {
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
				reserves, err := client.Reserves(name) // get amount for swaping
				if (err != nil) != wantErr {
					t.Errorf("%s wantErr %v", name, wantErr)
				}
				if wantErr {
					return
				}
				t.Logf("%s reserve: %v", name, reserves)
			})
		}
	})
	t.Run("PairDecimals", func(t *testing.T) {
		var args = map[string]int{
			"ndx-eth":   18,
			"eth-ndx":   18,
			"cc10-eth":  18,
			"eth-cc10":  18,
			"defi5-eth": 18,
			"eth-defi5": 18,
		}
		for name, wantDec := range args {
			dec := client.PairDecimals(name)
			require.Equal(t, dec, wantDec)
		}

	})
}
