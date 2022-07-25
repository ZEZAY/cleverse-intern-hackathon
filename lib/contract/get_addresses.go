package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/pkg/errors"
)

// TODO: subscribe factory (?)

func GetVoteAddresses(client *ethclient.Client, callOpts *bind.CallOpts) ([]common.Address, error) {

	// TODO: sort No. from 1 to n/2
	addresses := []common.Address{
		common.HexToAddress("0x3c59bF5BE87ae9e3dD7d913DE200F3E1253e7c90"),
		common.HexToAddress("0xaF409beD0ae55037BfB2fB1Fae3Ca7698DD86B9B"),
	}

	return addresses, nil
}

func GetAskAddresses(client *ethclient.Client, callOpts *bind.CallOpts) ([]common.Address, error) {

	// TODO: sort No. from (n/2+1) to n
	addresses := []common.Address{
		common.HexToAddress("0x68389682d7855961B61128D8aEef37A4b19C4fC5"),
		common.HexToAddress("0xa775b226951Ac8e07D5D2bC89a75bcA6a6a49b3c"),
	}

	return addresses, nil
}
