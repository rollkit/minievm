// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_inference

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
	_ = abi.ConvertType
)

// IInferenceMetaData contains all meta data concerning the IInference contract.
var IInferenceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"model\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IInferenceABI is the input ABI used to generate the binding from.
// Deprecated: Use IInferenceMetaData.ABI instead.
var IInferenceABI = IInferenceMetaData.ABI

// IInference is an auto generated Go binding around an Ethereum contract.
type IInference struct {
	IInferenceCaller     // Read-only binding to the contract
	IInferenceTransactor // Write-only binding to the contract
	IInferenceFilterer   // Log filterer for contract events
}

// IInferenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInferenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInferenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInferenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInferenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInferenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInferenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInferenceSession struct {
	Contract     *IInference       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInferenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInferenceCallerSession struct {
	Contract *IInferenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IInferenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInferenceTransactorSession struct {
	Contract     *IInferenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IInferenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInferenceRaw struct {
	Contract *IInference // Generic contract binding to access the raw methods on
}

// IInferenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInferenceCallerRaw struct {
	Contract *IInferenceCaller // Generic read-only contract binding to access the raw methods on
}

// IInferenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInferenceTransactorRaw struct {
	Contract *IInferenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInference creates a new instance of IInference, bound to a specific deployed contract.
func NewIInference(address common.Address, backend bind.ContractBackend) (*IInference, error) {
	contract, err := bindIInference(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInference{IInferenceCaller: IInferenceCaller{contract: contract}, IInferenceTransactor: IInferenceTransactor{contract: contract}, IInferenceFilterer: IInferenceFilterer{contract: contract}}, nil
}

// NewIInferenceCaller creates a new read-only instance of IInference, bound to a specific deployed contract.
func NewIInferenceCaller(address common.Address, caller bind.ContractCaller) (*IInferenceCaller, error) {
	contract, err := bindIInference(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInferenceCaller{contract: contract}, nil
}

// NewIInferenceTransactor creates a new write-only instance of IInference, bound to a specific deployed contract.
func NewIInferenceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInferenceTransactor, error) {
	contract, err := bindIInference(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInferenceTransactor{contract: contract}, nil
}

// NewIInferenceFilterer creates a new log filterer instance of IInference, bound to a specific deployed contract.
func NewIInferenceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInferenceFilterer, error) {
	contract, err := bindIInference(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInferenceFilterer{contract: contract}, nil
}

// bindIInference binds a generic wrapper to an already deployed contract.
func bindIInference(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInferenceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInference *IInferenceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInference.Contract.IInferenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInference *IInferenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInference.Contract.IInferenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInference *IInferenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInference.Contract.IInferenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInference *IInferenceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInference.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInference *IInferenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInference.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInference *IInferenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInference.Contract.contract.Transact(opts, method, params...)
}

// Infer is a paid mutator transaction binding the contract method 0x5586dca6.
//
// Solidity: function infer(bytes model, bytes input) returns(bytes)
func (_IInference *IInferenceTransactor) Infer(opts *bind.TransactOpts, model []byte, input []byte) (*types.Transaction, error) {
	return _IInference.contract.Transact(opts, "infer", model, input)
}

// Infer is a paid mutator transaction binding the contract method 0x5586dca6.
//
// Solidity: function infer(bytes model, bytes input) returns(bytes)
func (_IInference *IInferenceSession) Infer(model []byte, input []byte) (*types.Transaction, error) {
	return _IInference.Contract.Infer(&_IInference.TransactOpts, model, input)
}

// Infer is a paid mutator transaction binding the contract method 0x5586dca6.
//
// Solidity: function infer(bytes model, bytes input) returns(bytes)
func (_IInference *IInferenceTransactorSession) Infer(model []byte, input []byte) (*types.Transaction, error) {
	return _IInference.Contract.Infer(&_IInference.TransactOpts, model, input)
}
