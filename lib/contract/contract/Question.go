package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Question interface {
	QuestionName(*bind.CallOpts) (string, error)
	GetTimeStampData(*bind.CallOpts) (*big.Int, *big.Int, error)
}
