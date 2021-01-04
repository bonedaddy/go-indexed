package bclient

import (
	"math/big"

	governoralpha "github.com/bonedaddy/go-indexed/bindings/governor_alpha"
)

/*
  enum ProposalState {
    Pending,
    Active,
    Canceled,
    Defeated,
    Succeeded,
    Queued,
    Expired,
    Executed
  }
*/

// ProposalState indicates a state that a proposal may be in
type ProposalState string

func (ps ProposalState) String() string {
	return string(ps)
}

var (
	Pending   = ProposalState("Pending")
	Active    = ProposalState("Active")
	Canceled  = ProposalState("Canceled")
	Defeated  = ProposalState("Defeated")
	Succeeded = ProposalState("Succeeded")
	Queued    = ProposalState("Queued")
	Expired   = ProposalState("Expired")
	Executed  = ProposalState("Executed")
)

// ProposalFromUint returns a proposal
func ProposalFromUint(state uint8) ProposalState {
	switch state {
	case 0:
		return Pending
	case 1:
		return Active
	case 2:
		return Canceled
	case 3:
		return Defeated
	case 4:
		return Succeeded
	case 5:
		return Queued
	case 6:
		return Expired
	case 7:
		return Executed
	default:
		return ProposalState("INVALID")
	}
}

// GetProposalState returns the state of the proposal
func GetProposalState(
	gov *governoralpha.Governoralpha,
	number *big.Int,
) (ProposalState, error) {
	state, err := gov.State(nil, number)
	if err != nil {
		return ProposalState(""), err
	}
	return ProposalFromUint(state), nil
}
