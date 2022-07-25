package datamodel

import (
	"encoding/json"
	"time"

	"hackathon/lib/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Choice struct {
	Index     int     `json:"index"`
	Choice    string  `json:"choice"`
	VoteCount int     `json:"voteCount"`
	BetCount  int     `json:"betCount"`
	BetValue  float64 `json:"betValue"`
}

type Topic struct {
	No             int            `json:"no"`
	Address        common.Address `json:"address"`
	Question       string         `json:"question"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	TimeStartVote  time.Time      `json:"timeStartVote"`
	TimeEndVote    time.Time      `json:"timeEndVote"`
	TimeStartBet   time.Time      `json:"timeStartBet"`
	TimeEndBet     time.Time      `json:"timeEndBet"`
	TotalVoteCount int            `json:"totalVoteCount"`
	TotalBetCount  int            `json:"totalBetCount"`
	TotalBetValue  float64        `json:"totalBetValue"`
	Choices        []Choice       `json:"choices"`
}

func (t Topic) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t Topic) Unmarshal(data []byte) (*Topic, error) {
	topic := &Topic{}
	err := json.Unmarshal(data, topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (t Topic) ToMassage() TopicMassage {
	now := time.Now()
	isVote := now.After(t.TimeStartVote) && now.Before(t.TimeEndVote)
	isBet := now.After(t.TimeStartBet) && now.Before(t.TimeEndBet)
	isActive := isVote || isBet

	msg := TopicMassage{
		No:             t.No,
		Address:        t.Address,
		Question:       t.Question,
		Description:    t.Description,
		Category:       t.Category,
		TimeStartBet:   utils.TimeToString(t.TimeStartBet),
		TimeStartVote:  utils.TimeToString(t.TimeStartVote),
		TimeEndBet:     utils.TimeToString(t.TimeEndBet),
		TimeEndVote:    utils.TimeToString(t.TimeEndVote),
		TotalVoteCount: t.TotalVoteCount,
		TotalBetCount:  t.TotalBetCount,
		TotalBetValue:  t.TotalBetValue,
		Choices:        t.Choices,
		IsVote:         isVote,
		IsBet:          isBet,
		IsActive:       isActive,
	}
	return msg
}
