# go-indexed


`go-indexed` is a library for interacting [Indexed Finance](https://indexed.finance) smart contracts, and includes a Discord bot that can be used for polling various Indexed contracts via Discord messages.

# Overview

## Discord

Current capabilities:

* Return account balance for IndexPools contracts
* Return account staking balance and stake rewards earned for the staking contracts
* Return current tokens basketed into an IndexPool
* Notify about arbitrary conditions:
  * Currently only supports stake earned


# Development

To update the ABIs from the `indexed-js` submodule: `make copy-abi`

To update the generated golang code for the ABIs: `make gen-bindings`

Note: Updating the uniswap bindings is a bit tedious because of the usage of waffle which outputs a combined json file that isnt properly generating the golang bindings so you must manually copy the ABI from the combined json output into its own file

## Contract Bindings

#  bindings/uniswapv2_oracle

These are the contract bindings for the Indexed UniswapV2 Oracle

# bindings/uniswapv2

These are contract bindings for uniswapv2 itself