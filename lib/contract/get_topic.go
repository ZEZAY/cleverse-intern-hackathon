package contract

import (
	"fmt"
	"math/big"

	"hackathon/lib/datamodel"
	"hackathon/lib/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	dummyDescription = "What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	dummyCategory    = "Lorem"
)

func GetTopic(client *ethclient.Client, callOpts *bind.CallOpts, address common.Address, no int, clientChan chan *ethclient.Client) (*datamodel.Topic, error) {
	question, err := NewQuestionContract(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get question failed")
	}

	questionName, err := question.Name(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get questionName failed")
	}

	totalVoteCount, totalBetCount, totalBetValue, startBetAt, endBetAt, startVoteAt, endVoteAt, err := question.GetQuestionData(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get questionData failed")
	}

	timeStartVote, err := utils.StringToTime(startVoteAt.String())
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get timeStartVote failed")
	}

	timeStartBet, err := utils.StringToTime(startBetAt.String())
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get timeStartBet failed")
	}

	timeEndVote, err := utils.StringToTime(endVoteAt.String())
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get timeEndVote failed")
	}

	timeEndBet, err := utils.StringToTime(endBetAt.String())
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get timeEndBet failed")
	}

	choicesName, _, err := question.GetAllChoices(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get choicesName failed")
	}

	choices := []datamodel.Choice{}
	for i, name := range choicesName {

		nVote, err := question.GetChoiceVoteData(callOpts, name)
		if err != nil {
			return nil, errors.Wrap(err, "GetTopic get nVote failed")
		}

		nBet, vBet, err := question.GetChoiceBetData(callOpts, name)
		if err != nil {
			// return nil, errors.Wrap(err, "GetTopic get nBet, vBet failed")
			fmt.Println(err, "GetTopic get nBet, vBet failed")
			nBet, vBet = big.NewInt(0), big.NewInt(0)
		}

		choice := datamodel.Choice{
			Index:     i,
			Choice:    name,
			VoteCount: int(nVote.Int64()),
			BetCount:  int(nBet.Int64()),
			BetValue:  utils.BigIntToFloat64(vBet, 18),
		}
		choices = append(choices, choice)
	}

	topic := datamodel.Topic{
		No:             no,
		Address:        address,
		Question:       questionName,
		Description:    dummyDescription,
		Category:       dummyCategory,
		TimeStartBet:   timeStartBet,
		TimeStartVote:  timeStartVote,
		TimeEndBet:     timeEndBet,
		TimeEndVote:    timeEndVote,
		TotalVoteCount: int(totalVoteCount.Int64()),
		TotalBetCount:  int(totalBetCount.Int64()),
		TotalBetValue:  utils.BigIntToFloat64(totalBetValue, 18),
		Choices:        choices,
	}

	clientChan <- client
	return &topic, nil
}
