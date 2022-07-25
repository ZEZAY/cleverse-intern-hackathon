package datamodel

import (
	"github.com/ethereum/go-ethereum/common"
)

type TopicMassage struct {
	No             int            `json:"no"`
	Address        common.Address `json:"address"`
	Question       string         `json:"title"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	TimeStartBet   string         `json:"timeStartBet"`
	TimeStartVote  string         `json:"timeStartVote"`
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

// sort.Interface

type TopicMsgByTitle []TopicMassage

func (t TopicMsgByTitle) Len() int           { return len(t) }
func (t TopicMsgByTitle) Less(i, j int) bool { return t[i].Question < t[j].Question }
func (t TopicMsgByTitle) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type TopicMsgByPrize []TopicMassage

func (t TopicMsgByPrize) Len() int           { return len(t) }
func (t TopicMsgByPrize) Less(i, j int) bool { return t[i].TotalBetValue < t[j].TotalBetValue }
func (t TopicMsgByPrize) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type TopicMsgByVote []TopicMassage

func (t TopicMsgByVote) Len() int           { return len(t) }
func (t TopicMsgByVote) Less(i, j int) bool { return t[i].TotalVoteCount < t[j].TotalVoteCount }
func (t TopicMsgByVote) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
