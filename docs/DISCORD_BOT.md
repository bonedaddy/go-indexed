# NDXBot Documentation

This is the documentation for all NDXBot discord commands. 

# Usage

NDXBot commands start with the identifier `!ndx`. Calling `!ndx` will show the help menu for all top level commands. 

The level commands are:

* notify (disabled)
* uniswap
* pool-tokens
* pool-balance
* staked-earned

# Commands

## notify

Calling `!ndx notify` allows you to schedule notifications via PMs about arbitrary conditions. It is currently disabled.

## uniswap

Calling `!ndx uniswap` allows you to query various Indexed related uniswap contracts.

Command syntax: `!ndx uniswap <command> <args...>`

Existing `<command>`s:

* `reserves`
  * Returns available uniswap pair reserves
* `exchange-amount`
  * Returns the amount of tokens returned in exchange for the specified amount

Every uniswap related commands requires you to specify a "pair", each pair has two ways of being specified, with each method specifying a swap direction. For example `eth-defi5` is for the ETH/DEFI5 uniswap pair. When the previous is written as `defi5-eth` this also refers to the ETH/DEFI5 uniswap pair.  The difference is that in commands such as `exchange-amount` when writing `eth-defi5` you are specifying a swap in the direction of `ETH -> DEFI5`, and when writing `defi5-eth` you are a specifying a swap in the direction of `DEFI5 ->ETH`.

One way to think of this is that the left side of the dash refers to `token0` and the right side of the dash refers to `token1`. This means that in `defi5-eth`, `defi5` is `token0` and `eth` is `token1`.

Current accepted uniswap pairs are:

* `eth-defi5`, `defi5-eth`
* `eth-cc10`, `cc10-eth`
* `eh-ndx`, `ndx-eth`

### reserves

TODO

### exchange-amount

Command syntax: `!ndx uniswap exchange-amount <pair> <amount-float>`

`<pair>` Indicates a given pair, or more specifically the swap direction. If you were to specify `eth-defi5` you would be looking up the amount of DEFI5 tokens you would get in exchange for `<amount-float> of ETH`.
* `<amount-float>` is the amount of tokens to swap in decimal/float
  * NDXBot will do its best to appropriately handle conversion from decimal/float into whatever decimal value the tokens use. If you find erroneous values please submit an issue.


For example if you wanted to get the amount of DEFI5 tokens you would receive in exchange for 0.1 ETH you would write `!ndx uniswap exchange-amount eth-defi5 0.1`.

## pool-tokens

Calling `!ndx pool-tokens` allows you to return the current tokens basketed into a pool.

Command syntax: `!ndx pool-tokens <pool>`

Currently the only available pools are CC10 and DEFI5 although this will change as time goes on.

## pool-balance

Calling `!ndx pool-balance` will allow you to return the amount of pool tokens held by an account.

Command syntax: `!ndx pool-balance <pool> <account>`

* `<pool>` is the name of the pool
* `<account>` is the ethereum address to check

## stake-earned

Calling `!ndx stake-earned` will allow you to return the amount of staking rewards an account has earned.

Command syntax: `!ndx staking-earned <stake-type> <account>`

* `<stake-type.` is the type of staking rewards contract which currently is `defi5` and `univ2-eth-defi5` for the uniswap LP stake contract
* `<account>` is the ethereum address to check