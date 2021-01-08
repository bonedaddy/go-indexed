# go-indexed


`go-indexed` is a Golang SDK for interacting with [Indexed Finance](https://indexed.finance) smart contracts. The main purpose of go-indexed is to facilitate the development of NDXBot which is a Discord bot designed to allow performing read-only queries to these smart contracts.


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
* `uniswap`
  * uniswap contract wrappers, essentially the `bclient` equivalent but for Uniswap V2 contracts
  * largely copy & pasted from [mysterium/payments](https://github.com/mysteriumnetwork/payments)
* `utils`
  * Utility functions from [goethereumbook](goethereumbook.org/)
# Installation

Requires you have a valid Go 1.15+ installation, as well as Docker Compose installed locally

```Shell
$> git clone https://github.com/bonedaddy/go-indexed
$> cd go-indexed
$> go mod download
$> make release # builds the docker image and CLI 
```

# Usage

## Library

At the moment there is no documentation for using go-indexed as a library, only tests. Please look at all `_test.go` files in the following directories to get a better understanding of how the library functions:

* bclient
* uniswap

## Discord 

There are a large number of available commands, and subcommands. For documentation on on that [please click here](./docs/DISCORD_BOT.md). Alternatively if you are on Discord simply open up a chat the bot is listening on and run `!ndx help`.

# Development

## ABI Changes

If you are doing something that requires ABI changes, you will want to make sure that all submodules are correctly updated, they are:

 * indexed-core
 * indexed-governance
 * indexed-js
 * indexed-uv2oracle
 * uniswap-v2-core

After ensuring the submodules are updated, you will need to copy ABIs from `indexed-js` with `make copy-abi`. After that you can genrate all other bindings (except the uniswap ones) with `make gen-bindings`.

Updating the uniswap bindings is a bit tedious because of the usage of waffle which outputs a combined json file that isnt properly generating the golang bindings so you must manually copy the ABI from the combined json output into its own file
These are contract bindings for uniswapv2 itself

# Support

Support development of go-indexed by sending tips to [`0x5a361a1dfd52538a158e352d21b5b622360a7c13`](https://etherscan.io/address/0x5a361a1dfd52538a158e352d21b5b622360a7c13), all funds received will be re-invested into Indexed Finance via governance token (NDX) purchases, or pool token (defi5, cc10, etc...) purchases.