package datamodel

import "github.com/ethereum/go-ethereum/common"

// might have no usage

type Vote struct {
	User     string `json: "user"`
	Question string `json: "question"`
	Choice   int    `json: "choice"`
}

type Bet struct {
	User     string  `json: "user"`
	Question string  `json: "question"`
	Choice   int     `json: "choice"`
	Value    float64 `json: "value"`
}

type User struct {
	Address common.Address `json:"address"`
}
