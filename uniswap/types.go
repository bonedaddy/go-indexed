package uniswap

import "math/big"

// Reserve represents a given uniswap pair reserve
type Reserve struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}
