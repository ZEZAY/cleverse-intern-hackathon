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

// AskQuestionMetaData contains all meta data concerning the AskQuestion contract.
var AskQuestionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_questionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_questionDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startQuestionAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userAnswer\",\"type\":\"string\"}],\"name\":\"answer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAnswer\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeStampData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUserAnswerData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"questionDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"questionName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AskQuestionABI is the input ABI used to generate the binding from.
// Deprecated: Use AskQuestionMetaData.ABI instead.
var AskQuestionABI = AskQuestionMetaData.ABI

// AskQuestion is an auto generated Go binding around an Ethereum contract.
type AskQuestion struct {
	AskQuestionCaller     // Read-only binding to the contract
	AskQuestionTransactor // Write-only binding to the contract
	AskQuestionFilterer   // Log filterer for contract events
}

// AskQuestionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AskQuestionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskQuestionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AskQuestionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskQuestionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AskQuestionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskQuestionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AskQuestionSession struct {
	Contract     *AskQuestion      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AskQuestionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AskQuestionCallerSession struct {
	Contract *AskQuestionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AskQuestionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AskQuestionTransactorSession struct {
	Contract     *AskQuestionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AskQuestionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AskQuestionRaw struct {
	Contract *AskQuestion // Generic contract binding to access the raw methods on
}

// AskQuestionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AskQuestionCallerRaw struct {
	Contract *AskQuestionCaller // Generic read-only contract binding to access the raw methods on
}

// AskQuestionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AskQuestionTransactorRaw struct {
	Contract *AskQuestionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAskQuestion creates a new instance of AskQuestion, bound to a specific deployed contract.
func NewAskQuestion(address common.Address, backend bind.ContractBackend) (*AskQuestion, error) {
	contract, err := bindAskQuestion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AskQuestion{AskQuestionCaller: AskQuestionCaller{contract: contract}, AskQuestionTransactor: AskQuestionTransactor{contract: contract}, AskQuestionFilterer: AskQuestionFilterer{contract: contract}}, nil
}

// NewAskQuestionCaller creates a new read-only instance of AskQuestion, bound to a specific deployed contract.
func NewAskQuestionCaller(address common.Address, caller bind.ContractCaller) (*AskQuestionCaller, error) {
	contract, err := bindAskQuestion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AskQuestionCaller{contract: contract}, nil
}

// NewAskQuestionTransactor creates a new write-only instance of AskQuestion, bound to a specific deployed contract.
func NewAskQuestionTransactor(address common.Address, transactor bind.ContractTransactor) (*AskQuestionTransactor, error) {
	contract, err := bindAskQuestion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AskQuestionTransactor{contract: contract}, nil
}

// NewAskQuestionFilterer creates a new log filterer instance of AskQuestion, bound to a specific deployed contract.
func NewAskQuestionFilterer(address common.Address, filterer bind.ContractFilterer) (*AskQuestionFilterer, error) {
	contract, err := bindAskQuestion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AskQuestionFilterer{contract: contract}, nil
}

// bindAskQuestion binds a generic wrapper to an already deployed contract.
func bindAskQuestion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AskQuestionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AskQuestion *AskQuestionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AskQuestion.Contract.AskQuestionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AskQuestion *AskQuestionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskQuestion.Contract.AskQuestionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AskQuestion *AskQuestionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AskQuestion.Contract.AskQuestionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AskQuestion *AskQuestionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AskQuestion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AskQuestion *AskQuestionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskQuestion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AskQuestion *AskQuestionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AskQuestion.Contract.contract.Transact(opts, method, params...)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_AskQuestion *AskQuestionCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "contractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_AskQuestion *AskQuestionSession) ContractAddress() (common.Address, error) {
	return _AskQuestion.Contract.ContractAddress(&_AskQuestion.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_AskQuestion *AskQuestionCallerSession) ContractAddress() (common.Address, error) {
	return _AskQuestion.Contract.ContractAddress(&_AskQuestion.CallOpts)
}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_AskQuestion *AskQuestionCaller) ContractCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "contractCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_AskQuestion *AskQuestionSession) ContractCreator() (common.Address, error) {
	return _AskQuestion.Contract.ContractCreator(&_AskQuestion.CallOpts)
}

// ContractCreator is a free data retrieval call binding the contract method 0x1e2f73b1.
//
// Solidity: function contractCreator() view returns(address)
func (_AskQuestion *AskQuestionCallerSession) ContractCreator() (common.Address, error) {
	return _AskQuestion.Contract.ContractCreator(&_AskQuestion.CallOpts)
}

// GetAllAnswer is a free data retrieval call binding the contract method 0xafd7eb21.
//
// Solidity: function getAllAnswer() view returns(string[])
func (_AskQuestion *AskQuestionCaller) GetAllAnswer(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "getAllAnswer")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllAnswer is a free data retrieval call binding the contract method 0xafd7eb21.
//
// Solidity: function getAllAnswer() view returns(string[])
func (_AskQuestion *AskQuestionSession) GetAllAnswer() ([]string, error) {
	return _AskQuestion.Contract.GetAllAnswer(&_AskQuestion.CallOpts)
}

// GetAllAnswer is a free data retrieval call binding the contract method 0xafd7eb21.
//
// Solidity: function getAllAnswer() view returns(string[])
func (_AskQuestion *AskQuestionCallerSession) GetAllAnswer() ([]string, error) {
	return _AskQuestion.Contract.GetAllAnswer(&_AskQuestion.CallOpts)
}

// GetTimeStampData is a free data retrieval call binding the contract method 0x95bb742f.
//
// Solidity: function getTimeStampData() view returns(uint256, uint256)
func (_AskQuestion *AskQuestionCaller) GetTimeStampData(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "getTimeStampData")

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
func (_AskQuestion *AskQuestionSession) GetTimeStampData() (*big.Int, *big.Int, error) {
	return _AskQuestion.Contract.GetTimeStampData(&_AskQuestion.CallOpts)
}

// GetTimeStampData is a free data retrieval call binding the contract method 0x95bb742f.
//
// Solidity: function getTimeStampData() view returns(uint256, uint256)
func (_AskQuestion *AskQuestionCallerSession) GetTimeStampData() (*big.Int, *big.Int, error) {
	return _AskQuestion.Contract.GetTimeStampData(&_AskQuestion.CallOpts)
}

// GetUserAnswerData is a free data retrieval call binding the contract method 0x56ad4523.
//
// Solidity: function getUserAnswerData(address userAddress) view returns(string)
func (_AskQuestion *AskQuestionCaller) GetUserAnswerData(opts *bind.CallOpts, userAddress common.Address) (string, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "getUserAnswerData", userAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserAnswerData is a free data retrieval call binding the contract method 0x56ad4523.
//
// Solidity: function getUserAnswerData(address userAddress) view returns(string)
func (_AskQuestion *AskQuestionSession) GetUserAnswerData(userAddress common.Address) (string, error) {
	return _AskQuestion.Contract.GetUserAnswerData(&_AskQuestion.CallOpts, userAddress)
}

// GetUserAnswerData is a free data retrieval call binding the contract method 0x56ad4523.
//
// Solidity: function getUserAnswerData(address userAddress) view returns(string)
func (_AskQuestion *AskQuestionCallerSession) GetUserAnswerData(userAddress common.Address) (string, error) {
	return _AskQuestion.Contract.GetUserAnswerData(&_AskQuestion.CallOpts, userAddress)
}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_AskQuestion *AskQuestionCaller) QuestionDescription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "questionDescription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_AskQuestion *AskQuestionSession) QuestionDescription() (string, error) {
	return _AskQuestion.Contract.QuestionDescription(&_AskQuestion.CallOpts)
}

// QuestionDescription is a free data retrieval call binding the contract method 0x8e227d61.
//
// Solidity: function questionDescription() view returns(string)
func (_AskQuestion *AskQuestionCallerSession) QuestionDescription() (string, error) {
	return _AskQuestion.Contract.QuestionDescription(&_AskQuestion.CallOpts)
}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_AskQuestion *AskQuestionCaller) QuestionName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AskQuestion.contract.Call(opts, &out, "questionName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_AskQuestion *AskQuestionSession) QuestionName() (string, error) {
	return _AskQuestion.Contract.QuestionName(&_AskQuestion.CallOpts)
}

// QuestionName is a free data retrieval call binding the contract method 0x2847920e.
//
// Solidity: function questionName() view returns(string)
func (_AskQuestion *AskQuestionCallerSession) QuestionName() (string, error) {
	return _AskQuestion.Contract.QuestionName(&_AskQuestion.CallOpts)
}

// Answer is a paid mutator transaction binding the contract method 0xb2a4a50e.
//
// Solidity: function answer(string userAnswer) returns()
func (_AskQuestion *AskQuestionTransactor) Answer(opts *bind.TransactOpts, userAnswer string) (*types.Transaction, error) {
	return _AskQuestion.contract.Transact(opts, "answer", userAnswer)
}

// Answer is a paid mutator transaction binding the contract method 0xb2a4a50e.
//
// Solidity: function answer(string userAnswer) returns()
func (_AskQuestion *AskQuestionSession) Answer(userAnswer string) (*types.Transaction, error) {
	return _AskQuestion.Contract.Answer(&_AskQuestion.TransactOpts, userAnswer)
}

// Answer is a paid mutator transaction binding the contract method 0xb2a4a50e.
//
// Solidity: function answer(string userAnswer) returns()
func (_AskQuestion *AskQuestionTransactorSession) Answer(userAnswer string) (*types.Transaction, error) {
	return _AskQuestion.Contract.Answer(&_AskQuestion.TransactOpts, userAnswer)
}
