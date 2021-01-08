# go-indexed


`go-indexed` is a Golang SDK for interacting with [Indexed Finance](https://indexed.finance) smart contracts. The main purpose of go-indexed is to facilitate the development of NDXBot which is a Discord bot designed to allow performing read-only queries to these smart contracts. It can be used as a stand-alone SDK for general Indexed Finance usage.


# Packages

The library is broken down into a number of different packages:

* `bclient`
  * Wrapper around the Go-Ethereum `ethclient` package that includes helper functions for talking with Indexed contracts.
* `bindings`
  * Stores all generated code from `abigen`
* `db`
  * Small database package primarily designed for storing the value of various IndexedPool assets in terms of DAI
* `discord`
  * Provides the Discord bot functionality
# Installation

```Shell
$> git clone https://github.com/bonedaddy/go-indexed
$> cd go-indexed
$> go mod download
$> make # builds the CLI and creates an executable in the current directory named gondx
```
# Development

To update the ABIs from the `indexed-js` submodule: `make copy-abi`

To update the generated golang code for the ABIs: `make gen-bindings`

Note: Updating the uniswap bindings is a bit tedious because of the usage of waffle which outputs a combined json file that isnt properly generating the golang bindings so you must manually copy the ABI from the combined json output into its own file
These are contract bindings for uniswapv2 itself

# Support

Support development of go-indexed by sending tips to [`0x5a361a1dfd52538a158e352d21b5b622360a7c13`](https://etherscan.io/address/0x5a361a1dfd52538a158e352d21b5b622360a7c13), all funds received will be re-invested into Indexed Finance via governance token (NDX) purchases, or pool token (defi5, cc10, etc...) purchases.