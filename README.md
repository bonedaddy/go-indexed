# go-indexed


`go-indexed` is an SDK for interacting [Indexed Finance](https://indexed.finance) smart contracts using the Golang programming language. It includes a library that leverages the Ethereum RPC client, a CLI for querying the smart contracts from the CLI, and a Discord bot that essentially provides a CLI/Library like usage experience through the Discord platform.

# Overview

## Discord

A Discord chat bot is included that allows for making read-only calls to Indexed Finance contracts Discord message. It has the current capabilities

* Return account balance for IndexPools contracts
* Return account staking balance and stake rewards earned for the staking contracts
* Return current tokens basketed into an IndexPool
* Notify about arbitrary conditions:
  * Currently only supports stake earned

## CLI

There is a basic CLI that can be used to make simple queries about the smart contracts.

## Library

TODO

# Installation

```Shell
$> git clone https://github.com/bonedaddy/go-indexed
$> cd go-indexed
$> go mod download
$> make # builds the CLI and creates an executable in the current directory named gondx
```

# Usage

## CLI

At the moment the only CLI command available is `gondx pool` please see its output for more information.

## Discord

First you'll need to create a Discord bot user that will be used, and get an appropriate access token. The following instructions star the Discord bot conencted to infura

```shell
$> ./gondx --infura.api_key supersecretinfurakey discord-bot --discord.token "supersecretdiscordtoken"
```

To view the general help menu send the following message from a channel the bot can read from and send messages to:
```
!ndx
```

To view help menu about the notify command:
```
!ndx notify help
```

# Development

To update the ABIs from the `indexed-js` submodule: `make copy-abi`

To update the generated golang code for the ABIs: `make gen-bindings`

Note: Updating the uniswap bindings is a bit tedious because of the usage of waffle which outputs a combined json file that isnt properly generating the golang bindings so you must manually copy the ABI from the combined json output into its own file

## Contract Bindings

#  bindings/uniswapv2_oracle

These are the contract bindings for the Indexed UniswapV2 Oracle

# bindings/uniswapv2

These are contract bindings for uniswapv2 itself