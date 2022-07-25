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

// VoteQuestionMetaData contains all meta data concerning the VoteQuestion contract.
var VoteQuestionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_questionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_questionDescription\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_choices\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_startQuestionAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChoices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"choiceId\",\"type\":\"uint256\"}],\"name\":\"getChoiceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeStampData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUserVoteData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"questionDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"questionName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"choiceId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoteQuestionABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteQuestionMetaData.ABI instead.
var VoteQuestionABI = VoteQuestionMetaData.ABI

// VoteQuestion is an auto generated Go binding around an Ethereum contract.
type VoteQuestion struct {
	VoteQuestionCaller     // Read-only binding to the contract
	VoteQuestionTransactor // Write-only binding to the contract
	VoteQuestionFilterer   // Log filterer for contract events
}

// VoteQuestionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteQuestionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteQuestionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteQuestionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteQuestionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteQuestionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteQuestionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteQuestionSession struct {
	Contract     *VoteQuestion     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteQuestionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteQuestionCallerSession struct {
	Contract *VoteQuestionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoteQuestionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteQuestionTransactorSession struct {
	Contract     *VoteQuestionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoteQuestionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteQuestionRaw struct {
	Contract *VoteQuestion // Generic contract binding to access the raw methods on
}

// VoteQuestionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteQuestionCallerRaw struct {
	Contract *VoteQuestionCaller // Generic read-only contract binding to access the raw methods on
}

// VoteQuestionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteQuestionTransactorRaw struct {
	Contract *VoteQuestionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoteQuestion creates a new instance of VoteQuestion, bound to a specific deployed contract.
func NewVoteQuestion(address common.Address, backend bind.ContractBackend) (*VoteQuestion, error) {
	contract, err := bindVoteQuestion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoteQuestion{VoteQuestionCaller: VoteQuestionCaller{contract: contract}, VoteQuestionTransactor: VoteQuestionTransactor{contract: contract}, VoteQuestionFilterer: VoteQuestionFilterer{contract: contract}}, nil
}

// NewVoteQuestionCaller creates a new read-only instance of VoteQuestion, bound to a specific deployed contract.
func NewVoteQuestionCaller(address common.Address, caller bind.ContractCaller) (*VoteQuestionCaller, error) {
	contract, err := bindVoteQuestion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteQuestionCaller{contract: contract}, nil
}

// NewVoteQuestionTransactor creates a new write-only instance of VoteQuestion, bound to a specific deployed contract.
func NewVoteQuestionTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteQuestionTransactor, error) {
	contract, err := bindVoteQuestion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteQuestionTransactor{contract: contract}, nil
}

// NewVoteQuestionFilterer creates a new log filterer instance of VoteQuestion, bound to a specific deployed contract.
func NewVoteQuestionFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteQuestionFilterer, error) {
	contract, err := bindVoteQuestion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteQuestionFilterer{contract: contract}, nil
}

// bindVoteQuestion binds a generic wrapper to an already deployed contract.
func bindVoteQuestion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteQuestionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteQuestion *VoteQuestionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteQuestion.Contract.VoteQuestionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteQuestion *VoteQuestionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteQuestion.Contract.VoteQuestionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteQuestion *VoteQuestionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteQuestion.Contract.VoteQuestionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteQuestion *VoteQuestionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteQuestion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteQuestion *VoteQuestionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteQuestion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteQuestion *VoteQuestionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteQuestion.Contract.contract.Transact(opts, method, params...)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_VoteQuestion *VoteQuestionCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "contractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_VoteQuestion *VoteQuestionSession) ContractAddress() (common.Address, error) {
	return _VoteQuestion.Contract.ContractAddress(&_VoteQuestion.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_VoteQuestion *VoteQuestionCallerSession) ContractAddress() (common.Address, error) {
	return _VoteQuestion.Contract.ContractAddress(&_VoteQuestion.CallOpts)
}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_VoteQuestion *VoteQuestionCaller) ContractCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "contractCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_VoteQuestion *VoteQuestionSession) ContractCreator() (common.Address, error) {
	return _VoteQuestion.Contract.ContractCreator(&_VoteQuestion.CallOpts)
}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_VoteQuestion *VoteQuestionCallerSession) ContractCreator() (common.Address, error) {
	return _VoteQuestion.Contract.ContractCreator(&_VoteQuestion.CallOpts)
}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[])
func (_VoteQuestion *VoteQuestionCaller) GetAllChoices(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "getAllChoices")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[])
func (_VoteQuestion *VoteQuestionSession) GetAllChoices() ([]string, error) {
	return _VoteQuestion.Contract.GetAllChoices(&_VoteQuestion.CallOpts)
}

// GetAllChoices is a free data retrieval call binding the contract method 0xe97995dc.
//
// Solidity: function getAllChoices() view returns(string[])
func (_VoteQuestion *VoteQuestionCallerSession) GetAllChoices() ([]string, error) {
	return _VoteQuestion.Contract.GetAllChoices(&_VoteQuestion.CallOpts)
}

// GetChoiceData is a free data retrieval call binding the contract method 0x4eb1c4d1.
//
// Solidity: function getChoiceData(uint256 choiceId) view returns(uint256, string)
func (_VoteQuestion *VoteQuestionCaller) GetChoiceData(opts *bind.CallOpts, choiceId *big.Int) (*big.Int, string, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "getChoiceData", choiceId)

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetChoiceData is a free data retrieval call binding the contract method 0x4eb1c4d1.
//
// Solidity: function getChoiceData(uint256 choiceId) view returns(uint256, string)
func (_VoteQuestion *VoteQuestionSession) GetChoiceData(choiceId *big.Int) (*big.Int, string, error) {
	return _VoteQuestion.Contract.GetChoiceData(&_VoteQuestion.CallOpts, choiceId)
}

// GetChoiceData is a free data retrieval call binding the contract method 0x4eb1c4d1.
//
// Solidity: function getChoiceData(uint256 choiceId) view returns(uint256, string)
func (_VoteQuestion *VoteQuestionCallerSession) GetChoiceData(choiceId *big.Int) (*big.Int, string, error) {
	return _VoteQuestion.Contract.GetChoiceData(&_VoteQuestion.CallOpts, choiceId)
}

// GetTimeStampData is a free data retrieval call binding the contract method 0x95bb742f.
//
// Solidity: function getTimeStampData() view returns(uint256, uint256)
func (_VoteQuestion *VoteQuestionCaller) GetTimeStampData(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "getTimeStampData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTimeStampData is a free data retrieval call binding the contract method 0x95bb742f.
//
// Solidity: function getTimeStampData() view returns(uint256, uint256)
func (_VoteQuestion *VoteQuestionSession) GetTimeStampData() (*big.Int, *big.Int, error) {
	return _VoteQuestion.Contract.GetTimeStampData(&_VoteQuestion.CallOpts)
}

// GetTimeStampData is a free data retrieval call binding the contract method 0x95bb742f.
//
// Solidity: function getTimeStampData() view returns(uint256, uint256)
func (_VoteQuestion *VoteQuestionCallerSession) GetTimeStampData() (*big.Int, *big.Int, error) {
	return _VoteQuestion.Contract.GetTimeStampData(&_VoteQuestion.CallOpts)
}

// GetUserVoteData is a free data retrieval call binding the contract method 0x5ba91950.
//
// Solidity: function getUserVoteData(address userAddress) view returns(string)
func (_VoteQuestion *VoteQuestionCaller) GetUserVoteData(opts *bind.CallOpts, userAddress common.Address) (string, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "getUserVoteData", userAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserVoteData is a free data retrieval call binding the contract method 0x5ba91950.
//
// Solidity: function getUserVoteData(address userAddress) view returns(string)
func (_VoteQuestion *VoteQuestionSession) GetUserVoteData(userAddress common.Address) (string, error) {
	return _VoteQuestion.Contract.GetUserVoteData(&_VoteQuestion.CallOpts, userAddress)
}

// GetUserVoteData is a free data retrieval call binding the contract method 0x5ba91950.
//
// Solidity: function getUserVoteData(address userAddress) view returns(string)
func (_VoteQuestion *VoteQuestionCallerSession) GetUserVoteData(userAddress common.Address) (string, error) {
	return _VoteQuestion.Contract.GetUserVoteData(&_VoteQuestion.CallOpts, userAddress)
}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_VoteQuestion *VoteQuestionCaller) QuestionDescription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "questionDescription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_VoteQuestion *VoteQuestionSession) QuestionDescription() (string, error) {
	return _VoteQuestion.Contract.QuestionDescription(&_VoteQuestion.CallOpts)
}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_VoteQuestion *VoteQuestionCallerSession) QuestionDescription() (string, error) {
	return _VoteQuestion.Contract.QuestionDescription(&_VoteQuestion.CallOpts)
}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_VoteQuestion *VoteQuestionCaller) QuestionName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VoteQuestion.contract.Call(opts, &out, "questionName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_VoteQuestion *VoteQuestionSession) QuestionName() (string, error) {
	return _VoteQuestion.Contract.QuestionName(&_VoteQuestion.CallOpts)
}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_VoteQuestion *VoteQuestionCallerSession) QuestionName() (string, error) {
	return _VoteQuestion.Contract.QuestionName(&_VoteQuestion.CallOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 choiceId) returns()
func (_VoteQuestion *VoteQuestionTransactor) Vote(opts *bind.TransactOpts, choiceId *big.Int) (*types.Transaction, error) {
	return _VoteQuestion.contract.Transact(opts, "vote", choiceId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 choiceId) returns()
func (_VoteQuestion *VoteQuestionSession) Vote(choiceId *big.Int) (*types.Transaction, error) {
	return _VoteQuestion.Contract.Vote(&_VoteQuestion.TransactOpts, choiceId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 choiceId) returns()
func (_VoteQuestion *VoteQuestionTransactorSession) Vote(choiceId *big.Int) (*types.Transaction, error) {
	return _VoteQuestion.Contract.Vote(&_VoteQuestion.TransactOpts, choiceId)
}
