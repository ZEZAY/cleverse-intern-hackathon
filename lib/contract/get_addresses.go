package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/pkg/errors"
)

// TODO:

func GetAddresses(client *ethclient.Client, callOpts *bind.CallOpts) ([]common.Address, error) {
	// factoryAddress := common.HexToAddress("0x611764dA2046cdBef0Bd826FD0e889457e35b61D")
	// factory, err := NewFactoryContract(factoryAddress, client)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "GetAddresses get factory failed")
	// }

	// addresses, err := factory.GetQuestionAddress(callOpts)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "GetAddresses get addresses failed")
	// }

	// TODO: sort No. from 1 - n
	addresses := []common.Address{
		common.HexToAddress("0xc3a4EAaCB866146138B494368c6A60Cf41A95CB8"),
		common.HexToAddress("0x9A04E4Cff8b5BaAbDA11d252FC342Df8B2872226"),
	}

	return addresses, nil
}
