package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/pkg/errors"
)

// TODO: subscribe factory (?)

func GetVoteAddresses(client *ethclient.Client, callOpts *bind.CallOpts) ([]common.Address, error) {

	// TODO: sort No. from 1 to ...
	addresses := []common.Address{
		common.HexToAddress("0x6abB57d2cdCCcEb423eB268Afe9B8796289D2dca"),
		common.HexToAddress("0x38cBb2Cb6f95f84A12a5D24d30d9AFeef0BCaC1D"),
		common.HexToAddress("0xaD1574e841eb93Dad0857EDC6b7F83a7686B3210"),
		common.HexToAddress("0xeb3C91751D2B6Eb52d99746cEEce51650Da68c9E"),
		common.HexToAddress("0x14d6Ffaf184150dd7d445E4995845ffe5DB4BA53"),
		common.HexToAddress("0x56B390CC1c54Bc1037454a11329f7B6721F6274e"),
	}

	return addresses, nil
}

func GetAskAddresses(client *ethclient.Client, callOpts *bind.CallOpts) ([]common.Address, error) {

	// TODO: sort No. from ... to n
	addresses := []common.Address{
		// common.HexToAddress("0x68389682d7855961B61128D8aEef37A4b19C4fC5"),
		// common.HexToAddress("0xa775b226951Ac8e07D5D2bC89a75bcA6a6a49b3c"),
	}

	return addresses, nil
}
