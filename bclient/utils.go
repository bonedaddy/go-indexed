package bclient

import "errors"

func (bc *Client) GetIndexPool(name string) (IndexPool, error) {
	switch name {
	case "defi5":
		return bc.DEFI5()
	case "cc10":
		return bc.CC10()
	case "orcl5":
		return bc.ORCL5()
	default:
		return nil, errors.New("invalid pool name")
	}
}
