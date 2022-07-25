package datamodel

import (
	"github.com/ethereum/go-ethereum/common"
)

type TopicMassage struct {
	No            int            `json:"no"`
	Address       common.Address `json:"address"`
	Question      string         `json:"title"`
	Description   string         `json:"description"`
	Category      string         `json:"category"`
	TimeStart     string         `json:"timeStart"`
	TimeEnd       string         `json:"timeEnd"`
	ResponseCount int            `json:"responseCount"`
	Responses     []interface{}  `json:"responses"`
	IsActive      bool           `json:"isActive"`
}

// sort.Interface

type TopicMsgByTitle []TopicMassage

func (t TopicMsgByTitle) Len() int           { return len(t) }
func (t TopicMsgByTitle) Less(i, j int) bool { return t[i].Question < t[j].Question }
func (t TopicMsgByTitle) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type TopicMsgByResponseCount []TopicMassage

func (t TopicMsgByResponseCount) Len() int { return len(t) }
func (t TopicMsgByResponseCount) Less(i, j int) bool {
	return t[i].ResponseCount > t[j].ResponseCount
}
func (t TopicMsgByResponseCount) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
