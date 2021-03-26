package bclient

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// GetIndexPool is a helper function returning the underlying index pool by name
func (bc *Client) GetIndexPool(name string) (IndexPool, error) {
	switch strings.ToLower(name) {
	case "defi5":
		return bc.DEFI5()
	case "cc10":
		return bc.CC10()
	case "orcl5":
		return bc.ORCL5()
	case "degen10", "degen":
		return bc.DEGEN10()
	case "nftp":
		return bc.NFTP()
	default:
		return nil, errors.New("invalid pool name")
	}
}

// GetPoolAddress returns the index pool address for a given pool
func (bc *Client) GetPoolAddress(name string) (common.Address, error) {
	switch strings.ToLower(name) {
	case "defi5":
		return DEFI5TokenAddress, nil
	case "cc10":
		return CC10TokenAddress, nil
	case "orcl5":
		return ORCL5TokenAddress, nil
	case "degen10", "degen":
		return DEGEN10TokenAddress, nil
	case "nftp":
		return NFTPTokenAddress, nil
	default:
		return common.Address{}, errors.New("invalid pool name")
	}
}
