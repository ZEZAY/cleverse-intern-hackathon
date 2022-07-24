// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// QuestionContractMetaData contains all meta data concerning the QuestionContract contract.
var QuestionContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_choices\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_startBetAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_betDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startVoteAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_voteDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clickClose\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumQuestion.questionMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"forceCloseQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChoices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"}],\"name\":\"getBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"}],\"name\":\"getChoiceBetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getChoiceName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"}],\"name\":\"getChoiceVoteData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuestionData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserInformation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformCurrency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestampData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createQuestionAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBetAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startVoteAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBetAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endVoteAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"choiceName\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestionContractMetaData.ABI instead.
var QuestionContractABI = QuestionContractMetaData.ABI

// QuestionContract is an auto generated Go binding around an Ethereum contract.
type QuestionContract struct {
	QuestionContractCaller     // Read-only binding to the contract
	QuestionContractTransactor // Write-only binding to the contract
	QuestionContractFilterer   // Log filterer for contract events
}

// QuestionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuestionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestionContractSession struct {
	Contract     *QuestionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestionContractCallerSession struct {
	Contract *QuestionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// QuestionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestionContractTransactorSession struct {
	Contract     *QuestionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QuestionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuestionContractRaw struct {
	Contract *QuestionContract // Generic contract binding to access the raw methods on
}

// QuestionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestionContractCallerRaw struct {
	Contract *QuestionContractCaller // Generic read-only contract binding to access the raw methods on
}

// QuestionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestionContractTransactorRaw struct {
	Contract *QuestionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuestionContract creates a new instance of QuestionContract, bound to a specific deployed contract.
func NewQuestionContract(address common.Address, backend bind.ContractBackend) (*QuestionContract, error) {
	contract, err := bindQuestionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuestionContract{QuestionContractCaller: QuestionContractCaller{contract: contract}, QuestionContractTransactor: QuestionContractTransactor{contract: contract}, QuestionContractFilterer: QuestionContractFilterer{contract: contract}}, nil
}

// NewQuestionContractCaller creates a new read-only instance of QuestionContract, bound to a specific deployed contract.
func NewQuestionContractCaller(address common.Address, caller bind.ContractCaller) (*QuestionContractCaller, error) {
	contract, err := bindQuestionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestionContractCaller{contract: contract}, nil
}

// NewQuestionContractTransactor creates a new write-only instance of QuestionContract, bound to a specific deployed contract.
func NewQuestionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*QuestionContractTransactor, error) {
	contract, err := bindQuestionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestionContractTransactor{contract: contract}, nil
}

// NewQuestionContractFilterer creates a new log filterer instance of QuestionContract, bound to a specific deployed contract.
func NewQuestionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*QuestionContractFilterer, error) {
	contract, err := bindQuestionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestionContractFilterer{contract: contract}, nil
}

// bindQuestionContract binds a generic wrapper to an already deployed contract.
func bindQuestionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuestionContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestionContract *QuestionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestionContract.Contract.QuestionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestionContract *QuestionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestionContract.Contract.QuestionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestionContract *QuestionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestionContract.Contract.QuestionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestionContract *QuestionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestionContract *QuestionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestionContract *QuestionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestionContract.Contract.contract.Transact(opts, method, params...)
}

// ClickClose is a free data retrieval call binding the contract method 0x6cbae7a5.
//
// Solidity: function clickClose() view returns(bool)
func (_QuestionContract *QuestionContractCaller) ClickClose(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "clickClose")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClickClose is a free data retrieval call binding the contract method 0x6cbae7a5.
//
// Solidity: function clickClose() view returns(bool)
func (_QuestionContract *QuestionContractSession) ClickClose() (bool, error) {
	return _QuestionContract.Contract.ClickClose(&_QuestionContract.CallOpts)
}

// ClickClose is a free data retrieval call binding the contract method 0x6cbae7a5.
//
// Solidity: function clickClose() view returns(bool)
func (_QuestionContract *QuestionContractCallerSession) ClickClose() (bool, error) {
	return _QuestionContract.Contract.ClickClose(&_QuestionContract.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_QuestionContract *QuestionContractCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "contractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_QuestionContract *QuestionContractSession) ContractAddress() (common.Address, error) {
	return _QuestionContract.Contract.ContractAddress(&_QuestionContract.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_QuestionContract *QuestionContractCallerSession) ContractAddress() (common.Address, error) {
	return _QuestionContract.Contract.ContractAddress(&_QuestionContract.CallOpts)
}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[], uint256)
func (_QuestionContract *QuestionContractCaller) GetAllChoices(opts *bind.CallOpts) ([]string, *big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getAllChoices")

	if err != nil {
		return *new([]string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[], uint256)
func (_QuestionContract *QuestionContractSession) GetAllChoices() ([]string, *big.Int, error) {
	return _QuestionContract.Contract.GetAllChoices(&_QuestionContract.CallOpts)
}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[], uint256)
func (_QuestionContract *QuestionContractCallerSession) GetAllChoices() ([]string, *big.Int, error) {
	return _QuestionContract.Contract.GetAllChoices(&_QuestionContract.CallOpts)
}

// GetBet is a free data retrieval call binding the contract method 0x078d1600.
//
// Solidity: function getBet(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCaller) GetBet(opts *bind.CallOpts, choiceName string) (*big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getBet", choiceName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBet is a free data retrieval call binding the contract method 0x078d1600.
//
// Solidity: function getBet(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractSession) GetBet(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetBet(&_QuestionContract.CallOpts, choiceName)
}

// GetBet is a free data retrieval call binding the contract method 0x078d1600.
//
// Solidity: function getBet(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCallerSession) GetBet(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetBet(&_QuestionContract.CallOpts, choiceName)
}

// GetChoiceBetData is a free data retrieval call binding the contract method 0x362448ce.
//
// Solidity: function getChoiceBetData(string choiceName) view returns(uint256, uint256)
func (_QuestionContract *QuestionContractCaller) GetChoiceBetData(opts *bind.CallOpts, choiceName string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getChoiceBetData", choiceName)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetChoiceBetData is a free data retrieval call binding the contract method 0x362448ce.
//
// Solidity: function getChoiceBetData(string choiceName) view returns(uint256, uint256)
func (_QuestionContract *QuestionContractSession) GetChoiceBetData(choiceName string) (*big.Int, *big.Int, error) {
	return _QuestionContract.Contract.GetChoiceBetData(&_QuestionContract.CallOpts, choiceName)
}

// GetChoiceBetData is a free data retrieval call binding the contract method 0x362448ce.
//
// Solidity: function getChoiceBetData(string choiceName) view returns(uint256, uint256)
func (_QuestionContract *QuestionContractCallerSession) GetChoiceBetData(choiceName string) (*big.Int, *big.Int, error) {
	return _QuestionContract.Contract.GetChoiceBetData(&_QuestionContract.CallOpts, choiceName)
}

// GetChoiceName is a free data retrieval call binding the contract method 0xf3bd7907.
//
// Solidity: function getChoiceName(uint256 _index) view returns(string)
func (_QuestionContract *QuestionContractCaller) GetChoiceName(opts *bind.CallOpts, _index *big.Int) (string, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getChoiceName", _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetChoiceName is a free data retrieval call binding the contract method 0xf3bd7907.
//
// Solidity: function getChoiceName(uint256 _index) view returns(string)
func (_QuestionContract *QuestionContractSession) GetChoiceName(_index *big.Int) (string, error) {
	return _QuestionContract.Contract.GetChoiceName(&_QuestionContract.CallOpts, _index)
}

// GetChoiceName is a free data retrieval call binding the contract method 0xf3bd7907.
//
// Solidity: function getChoiceName(uint256 _index) view returns(string)
func (_QuestionContract *QuestionContractCallerSession) GetChoiceName(_index *big.Int) (string, error) {
	return _QuestionContract.Contract.GetChoiceName(&_QuestionContract.CallOpts, _index)
}

// GetChoiceVoteData is a free data retrieval call binding the contract method 0xc3567445.
//
// Solidity: function getChoiceVoteData(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCaller) GetChoiceVoteData(opts *bind.CallOpts, choiceName string) (*big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getChoiceVoteData", choiceName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChoiceVoteData is a free data retrieval call binding the contract method 0xc3567445.
//
// Solidity: function getChoiceVoteData(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractSession) GetChoiceVoteData(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetChoiceVoteData(&_QuestionContract.CallOpts, choiceName)
}

// GetChoiceVoteData is a free data retrieval call binding the contract method 0xc3567445.
//
// Solidity: function getChoiceVoteData(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCallerSession) GetChoiceVoteData(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetChoiceVoteData(&_QuestionContract.CallOpts, choiceName)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_QuestionContract *QuestionContractCaller) GetContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_QuestionContract *QuestionContractSession) GetContractAddress() (common.Address, error) {
	return _QuestionContract.Contract.GetContractAddress(&_QuestionContract.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_QuestionContract *QuestionContractCallerSession) GetContractAddress() (common.Address, error) {
	return _QuestionContract.Contract.GetContractAddress(&_QuestionContract.CallOpts)
}

// GetQuestionData is a free data retrieval call binding the contract method 0x0d8e9cc7.
//
// Solidity: function getQuestionData() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_QuestionContract *QuestionContractCaller) GetQuestionData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getQuestionData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetQuestionData is a free data retrieval call binding the contract method 0x0d8e9cc7.
//
// Solidity: function getQuestionData() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_QuestionContract *QuestionContractSession) GetQuestionData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _QuestionContract.Contract.GetQuestionData(&_QuestionContract.CallOpts)
}

// GetQuestionData is a free data retrieval call binding the contract method 0x0d8e9cc7.
//
// Solidity: function getQuestionData() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_QuestionContract *QuestionContractCallerSession) GetQuestionData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _QuestionContract.Contract.GetQuestionData(&_QuestionContract.CallOpts)
}

// GetRewardPending is a free data retrieval call binding the contract method 0x6efb9b7d.
//
// Solidity: function getRewardPending() view returns(uint256)
func (_QuestionContract *QuestionContractCaller) GetRewardPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getRewardPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPending is a free data retrieval call binding the contract method 0x6efb9b7d.
//
// Solidity: function getRewardPending() view returns(uint256)
func (_QuestionContract *QuestionContractSession) GetRewardPending() (*big.Int, error) {
	return _QuestionContract.Contract.GetRewardPending(&_QuestionContract.CallOpts)
}

// GetRewardPending is a free data retrieval call binding the contract method 0x6efb9b7d.
//
// Solidity: function getRewardPending() view returns(uint256)
func (_QuestionContract *QuestionContractCallerSession) GetRewardPending() (*big.Int, error) {
	return _QuestionContract.Contract.GetRewardPending(&_QuestionContract.CallOpts)
}

// GetUserInformation is a free data retrieval call binding the contract method 0x8997bb6e.
//
// Solidity: function getUserInformation() view returns(uint256, bool, bool, uint256, string, string)
func (_QuestionContract *QuestionContractCaller) GetUserInformation(opts *bind.CallOpts) (*big.Int, bool, bool, *big.Int, string, string, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getUserInformation")

	if err != nil {
		return *new(*big.Int), *new(bool), *new(bool), *new(*big.Int), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, err

}

// GetUserInformation is a free data retrieval call binding the contract method 0x8997bb6e.
//
// Solidity: function getUserInformation() view returns(uint256, bool, bool, uint256, string, string)
func (_QuestionContract *QuestionContractSession) GetUserInformation() (*big.Int, bool, bool, *big.Int, string, string, error) {
	return _QuestionContract.Contract.GetUserInformation(&_QuestionContract.CallOpts)
}

// GetUserInformation is a free data retrieval call binding the contract method 0x8997bb6e.
//
// Solidity: function getUserInformation() view returns(uint256, bool, bool, uint256, string, string)
func (_QuestionContract *QuestionContractCallerSession) GetUserInformation() (*big.Int, bool, bool, *big.Int, string, string, error) {
	return _QuestionContract.Contract.GetUserInformation(&_QuestionContract.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCaller) GetVote(opts *bind.CallOpts, choiceName string) (*big.Int, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "getVote", choiceName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractSession) GetVote(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetVote(&_QuestionContract.CallOpts, choiceName)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string choiceName) view returns(uint256)
func (_QuestionContract *QuestionContractCallerSession) GetVote(choiceName string) (*big.Int, error) {
	return _QuestionContract.Contract.GetVote(&_QuestionContract.CallOpts, choiceName)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestionContract *QuestionContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestionContract *QuestionContractSession) Name() (string, error) {
	return _QuestionContract.Contract.Name(&_QuestionContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestionContract *QuestionContractCallerSession) Name() (string, error) {
	return _QuestionContract.Contract.Name(&_QuestionContract.CallOpts)
}

// PlatformCurrency is a free data retrieval call binding the contract method 0xd435e399.
//
// Solidity: function platformCurrency() view returns(string)
func (_QuestionContract *QuestionContractCaller) PlatformCurrency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "platformCurrency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PlatformCurrency is a free data retrieval call binding the contract method 0xd435e399.
//
// Solidity: function platformCurrency() view returns(string)
func (_QuestionContract *QuestionContractSession) PlatformCurrency() (string, error) {
	return _QuestionContract.Contract.PlatformCurrency(&_QuestionContract.CallOpts)
}

// PlatformCurrency is a free data retrieval call binding the contract method 0xd435e399.
//
// Solidity: function platformCurrency() view returns(string)
func (_QuestionContract *QuestionContractCallerSession) PlatformCurrency() (string, error) {
	return _QuestionContract.Contract.PlatformCurrency(&_QuestionContract.CallOpts)
}

// PlatformToken is a free data retrieval call binding the contract method 0xd1b812cd.
//
// Solidity: function platformToken() view returns(address)
func (_QuestionContract *QuestionContractCaller) PlatformToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "platformToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformToken is a free data retrieval call binding the contract method 0xd1b812cd.
//
// Solidity: function platformToken() view returns(address)
func (_QuestionContract *QuestionContractSession) PlatformToken() (common.Address, error) {
	return _QuestionContract.Contract.PlatformToken(&_QuestionContract.CallOpts)
}

// PlatformToken is a free data retrieval call binding the contract method 0xd1b812cd.
//
// Solidity: function platformToken() view returns(address)
func (_QuestionContract *QuestionContractCallerSession) PlatformToken() (common.Address, error) {
	return _QuestionContract.Contract.PlatformToken(&_QuestionContract.CallOpts)
}

// TimestampData is a free data retrieval call binding the contract method 0x46376c35.
//
// Solidity: function timestampData() view returns(uint256 createQuestionAt, uint256 startBetAt, uint256 startVoteAt, uint256 endBetAt, uint256 endVoteAt)
func (_QuestionContract *QuestionContractCaller) TimestampData(opts *bind.CallOpts) (struct {
	CreateQuestionAt *big.Int
	StartBetAt       *big.Int
	StartVoteAt      *big.Int
	EndBetAt         *big.Int
	EndVoteAt        *big.Int
}, error) {
	var out []interface{}
	err := _QuestionContract.contract.Call(opts, &out, "timestampData")

	outstruct := new(struct {
		CreateQuestionAt *big.Int
		StartBetAt       *big.Int
		StartVoteAt      *big.Int
		EndBetAt         *big.Int
		EndVoteAt        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreateQuestionAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBetAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartVoteAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndBetAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndVoteAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TimestampData is a free data retrieval call binding the contract method 0x46376c35.
//
// Solidity: function timestampData() view returns(uint256 createQuestionAt, uint256 startBetAt, uint256 startVoteAt, uint256 endBetAt, uint256 endVoteAt)
func (_QuestionContract *QuestionContractSession) TimestampData() (struct {
	CreateQuestionAt *big.Int
	StartBetAt       *big.Int
	StartVoteAt      *big.Int
	EndBetAt         *big.Int
	EndVoteAt        *big.Int
}, error) {
	return _QuestionContract.Contract.TimestampData(&_QuestionContract.CallOpts)
}

// TimestampData is a free data retrieval call binding the contract method 0x46376c35.
//
// Solidity: function timestampData() view returns(uint256 createQuestionAt, uint256 startBetAt, uint256 startVoteAt, uint256 endBetAt, uint256 endVoteAt)
func (_QuestionContract *QuestionContractCallerSession) TimestampData() (struct {
	CreateQuestionAt *big.Int
	StartBetAt       *big.Int
	StartVoteAt      *big.Int
	EndBetAt         *big.Int
	EndVoteAt        *big.Int
}, error) {
	return _QuestionContract.Contract.TimestampData(&_QuestionContract.CallOpts)
}

// Bet is a paid mutator transaction binding the contract method 0xe9c20cb9.
//
// Solidity: function bet(string choiceName, uint256 amount) returns()
func (_QuestionContract *QuestionContractTransactor) Bet(opts *bind.TransactOpts, choiceName string, amount *big.Int) (*types.Transaction, error) {
	return _QuestionContract.contract.Transact(opts, "bet", choiceName, amount)
}

// Bet is a paid mutator transaction binding the contract method 0xe9c20cb9.
//
// Solidity: function bet(string choiceName, uint256 amount) returns()
func (_QuestionContract *QuestionContractSession) Bet(choiceName string, amount *big.Int) (*types.Transaction, error) {
	return _QuestionContract.Contract.Bet(&_QuestionContract.TransactOpts, choiceName, amount)
}

// Bet is a paid mutator transaction binding the contract method 0xe9c20cb9.
//
// Solidity: function bet(string choiceName, uint256 amount) returns()
func (_QuestionContract *QuestionContractTransactorSession) Bet(choiceName string, amount *big.Int) (*types.Transaction, error) {
	return _QuestionContract.Contract.Bet(&_QuestionContract.TransactOpts, choiceName, amount)
}

// CloseQuestion is a paid mutator transaction binding the contract method 0x6143bb03.
//
// Solidity: function closeQuestion() returns()
func (_QuestionContract *QuestionContractTransactor) CloseQuestion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestionContract.contract.Transact(opts, "closeQuestion")
}

// CloseQuestion is a paid mutator transaction binding the contract method 0x6143bb03.
//
// Solidity: function closeQuestion() returns()
func (_QuestionContract *QuestionContractSession) CloseQuestion() (*types.Transaction, error) {
	return _QuestionContract.Contract.CloseQuestion(&_QuestionContract.TransactOpts)
}

// CloseQuestion is a paid mutator transaction binding the contract method 0x6143bb03.
//
// Solidity: function closeQuestion() returns()
func (_QuestionContract *QuestionContractTransactorSession) CloseQuestion() (*types.Transaction, error) {
	return _QuestionContract.Contract.CloseQuestion(&_QuestionContract.TransactOpts)
}

// ForceCloseQuestion is a paid mutator transaction binding the contract method 0xe0e812b5.
//
// Solidity: function forceCloseQuestion(uint8 mode) returns()
func (_QuestionContract *QuestionContractTransactor) ForceCloseQuestion(opts *bind.TransactOpts, mode uint8) (*types.Transaction, error) {
	return _QuestionContract.contract.Transact(opts, "forceCloseQuestion", mode)
}

// ForceCloseQuestion is a paid mutator transaction binding the contract method 0xe0e812b5.
//
// Solidity: function forceCloseQuestion(uint8 mode) returns()
func (_QuestionContract *QuestionContractSession) ForceCloseQuestion(mode uint8) (*types.Transaction, error) {
	return _QuestionContract.Contract.ForceCloseQuestion(&_QuestionContract.TransactOpts, mode)
}

// ForceCloseQuestion is a paid mutator transaction binding the contract method 0xe0e812b5.
//
// Solidity: function forceCloseQuestion(uint8 mode) returns()
func (_QuestionContract *QuestionContractTransactorSession) ForceCloseQuestion(mode uint8) (*types.Transaction, error) {
	return _QuestionContract.Contract.ForceCloseQuestion(&_QuestionContract.TransactOpts, mode)
}

// HarvestReward is a paid mutator transaction binding the contract method 0x5a34928e.
//
// Solidity: function harvestReward() returns()
func (_QuestionContract *QuestionContractTransactor) HarvestReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestionContract.contract.Transact(opts, "harvestReward")
}

// HarvestReward is a paid mutator transaction binding the contract method 0x5a34928e.
//
// Solidity: function harvestReward() returns()
func (_QuestionContract *QuestionContractSession) HarvestReward() (*types.Transaction, error) {
	return _QuestionContract.Contract.HarvestReward(&_QuestionContract.TransactOpts)
}

// HarvestReward is a paid mutator transaction binding the contract method 0x5a34928e.
//
// Solidity: function harvestReward() returns()
func (_QuestionContract *QuestionContractTransactorSession) HarvestReward() (*types.Transaction, error) {
	return _QuestionContract.Contract.HarvestReward(&_QuestionContract.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string choiceName) returns()
func (_QuestionContract *QuestionContractTransactor) Vote(opts *bind.TransactOpts, choiceName string) (*types.Transaction, error) {
	return _QuestionContract.contract.Transact(opts, "vote", choiceName)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string choiceName) returns()
func (_QuestionContract *QuestionContractSession) Vote(choiceName string) (*types.Transaction, error) {
	return _QuestionContract.Contract.Vote(&_QuestionContract.TransactOpts, choiceName)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string choiceName) returns()
func (_QuestionContract *QuestionContractTransactorSession) Vote(choiceName string) (*types.Transaction, error) {
	return _QuestionContract.Contract.Vote(&_QuestionContract.TransactOpts, choiceName)
}
