package datamodel

import (
	"encoding/json"
	"strings"
	"time"

	"hackathon/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Topic struct {
	No            int            `json:"no"`
	Address       common.Address `json:"address"`
	Question      string         `json:"question"`
	Description   string         `json:"description"`
	Category      string         `json:"category"`
	TimeStart     time.Time      `json:"timeStart"`
	TimeEnd       time.Time      `json:"timeEnd"`
	ResponseCount int            `json:"responseCount"`
	Responses     []interface{}  `json:"responses"`
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
	isActive := now.After(t.TimeStart) && now.Before(t.TimeEnd)

	msg := TopicMassage{
		No:            t.No,
		Address:       t.Address,
		Question:      t.Question,
		Description:   t.Description,
		Category:      t.Category,
		TimeStart:     utils.TimeToString(t.TimeStart),
		TimeEnd:       utils.TimeToString(t.TimeEnd),
		ResponseCount: t.ResponseCount,
		Responses:     t.Responses,
		IsActive:      isActive,
	}
	return msg
}

func (t Topic) GetKey(topicType string) string {
	return topicType + "-" + strings.ToLower(t.Address.String())
}
