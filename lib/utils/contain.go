package utils

import "github.com/ethereum/go-ethereum/common"

func AddressesContains(addresses []common.Address, address common.Address) bool {
	for _, a := range addresses {
		if address == a {
			return true
		}
	}
	return false
}
