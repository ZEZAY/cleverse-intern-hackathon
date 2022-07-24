package datamodel

import (
	"github.com/ethereum/go-ethereum/common"
)

type TopicMassage struct {
	Address        common.Address `json:"address"`
	Question       string         `json:"title"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	TimeEndBet     string         `json:"timeEndBet"`
	TimeEndVote    string         `json:"timeEndVote"`
	TotalVoteCount int            `json:"totalVoteCount"`
	TotalBetCount  int            `json:"totalBetCount"`
	TotalBetValue  float64        `json:"prize"`
	Choices        []Choice       `json:"choices"`
	IsVote         bool           `json:"isVote"`
	IsBet          bool           `json:"isBet"`
	IsActive       bool           `json:"isActive"`
}
