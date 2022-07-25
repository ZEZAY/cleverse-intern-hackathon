package contract

import (
	"math/big"
	"time"

	"hackathon/lib/contract/contract"
	"hackathon/lib/datamodel"
	"hackathon/lib/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	dummyDescription = "What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
)

func getTopicData(question contract.Question, callOpts *bind.CallOpts) (
	questionName string,
	timeStart, timeEnd time.Time,
	err error,
) {
	questionName, err = question.QuestionName(callOpts)
	if err != nil {
		err = errors.Wrap(err, "getTopicData get questionName failed")
		return
	}

	startVoteAt, endVoteAt, err := question.GetTimeStampData(callOpts)
	if err != nil {
		err = errors.Wrap(err, "getTopicData get questionData failed")
		return
	}

	timeStart, err = utils.StringToTime(startVoteAt.String())
	if err != nil {
		err = errors.Wrap(err, "getTopicData get timeStart failed")
		return
	}

	timeEnd, err = utils.StringToTime(endVoteAt.String())
	if err != nil {
		err = errors.Wrap(err, "getTopicData get timeEnd failed")
		return
	}
	return
}

func GetVoteTopic(client *ethclient.Client, callOpts *bind.CallOpts, address common.Address, no int, clientChan chan *ethclient.Client) (*datamodel.Topic, error) {
	question, err := contract.NewVoteQuestion(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "GetVoteTopic get question failed")
	}

	questionName, timeStart, timeEnd, err := getTopicData(question, callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetVoteTopic get getTopicData failed")
	}

	choicesName, err := question.GetAllChoices(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetVoteTopic get choicesName failed")
	}

	type choice struct {
		Index  int    `json:"index"`
		Choice string `json:"choice"`
		Vote   int    `json:"vote"`
	}

	responses := []interface{}{}
	totalVoteCount := 0

	for i := 0; i < len(choicesName); i++ {
		nVote, name, err := question.GetChoiceData(callOpts, big.NewInt(int64(i)))
		if err != nil {
			return nil, errors.Wrap(err, "GetVoteTopic get nVote failed")
		}

		totalVoteCount += int(nVote.Int64())
		choice := choice{
			Index:  i,
			Choice: name,
			Vote:   int(nVote.Int64()),
		}
		responses = append(responses, choice)
	}

	topic := datamodel.Topic{
		No:            no,
		Address:       address,
		Question:      questionName,
		Description:   dummyDescription,
		Category:      "vote",
		TimeStart:     timeStart,
		TimeEnd:       timeEnd,
		ResponseCount: totalVoteCount,
		Responses:     responses,
	}

	clientChan <- client
	return &topic, nil
}

func GetAskTopic(client *ethclient.Client, callOpts *bind.CallOpts, address common.Address, no int, clientChan chan *ethclient.Client) (*datamodel.Topic, error) {
	question, err := contract.NewAskQuestion(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "GetAskTopic get question failed")
	}

	questionName, timeStart, timeEnd, err := getTopicData(question, callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetAskTopic get getTopicData failed")
	}

	answers, err := question.GetAllAnswer(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "GetAskTopic get choicesName failed")
	}

	responses := []interface{}{}
	for _, answer := range answers {
		responses = append(responses, answer)
	}

	topic := datamodel.Topic{
		No:            no,
		Address:       address,
		Question:      questionName,
		Description:   dummyDescription,
		Category:      "ask",
		TimeStart:     timeStart,
		TimeEnd:       timeEnd,
		ResponseCount: len(responses),
		Responses:     responses,
	}

	clientChan <- client
	return &topic, nil
}
