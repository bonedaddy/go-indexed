package bclient

import (
	"errors"
	"strings"
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
	default:
		return nil, errors.New("invalid pool name")
	}
}
