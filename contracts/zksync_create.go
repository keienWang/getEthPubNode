// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CreateWithdrawMetaData contains all meta data concerning the CreateWithdraw contract.
var CreateWithdrawMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CreateWithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use CreateWithdrawMetaData.ABI instead.
var CreateWithdrawABI = CreateWithdrawMetaData.ABI

// CreateWithdraw is an auto generated Go binding around an Ethereum contract.
type CreateWithdraw struct {
	CreateWithdrawCaller     // Read-only binding to the contract
	CreateWithdrawTransactor // Write-only binding to the contract
	CreateWithdrawFilterer   // Log filterer for contract events
}

// CreateWithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreateWithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateWithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreateWithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateWithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreateWithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateWithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreateWithdrawSession struct {
	Contract     *CreateWithdraw   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreateWithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreateWithdrawCallerSession struct {
	Contract *CreateWithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CreateWithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreateWithdrawTransactorSession struct {
	Contract     *CreateWithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CreateWithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreateWithdrawRaw struct {
	Contract *CreateWithdraw // Generic contract binding to access the raw methods on
}

// CreateWithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreateWithdrawCallerRaw struct {
	Contract *CreateWithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// CreateWithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreateWithdrawTransactorRaw struct {
	Contract *CreateWithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreateWithdraw creates a new instance of CreateWithdraw, bound to a specific deployed contract.
func NewCreateWithdraw(address common.Address, backend bind.ContractBackend) (*CreateWithdraw, error) {
	contract, err := bindCreateWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreateWithdraw{CreateWithdrawCaller: CreateWithdrawCaller{contract: contract}, CreateWithdrawTransactor: CreateWithdrawTransactor{contract: contract}, CreateWithdrawFilterer: CreateWithdrawFilterer{contract: contract}}, nil
}

// NewCreateWithdrawCaller creates a new read-only instance of CreateWithdraw, bound to a specific deployed contract.
func NewCreateWithdrawCaller(address common.Address, caller bind.ContractCaller) (*CreateWithdrawCaller, error) {
	contract, err := bindCreateWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreateWithdrawCaller{contract: contract}, nil
}

// NewCreateWithdrawTransactor creates a new write-only instance of CreateWithdraw, bound to a specific deployed contract.
func NewCreateWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*CreateWithdrawTransactor, error) {
	contract, err := bindCreateWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreateWithdrawTransactor{contract: contract}, nil
}

// NewCreateWithdrawFilterer creates a new log filterer instance of CreateWithdraw, bound to a specific deployed contract.
func NewCreateWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*CreateWithdrawFilterer, error) {
	contract, err := bindCreateWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreateWithdrawFilterer{contract: contract}, nil
}

// bindCreateWithdraw binds a generic wrapper to an already deployed contract.
func bindCreateWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreateWithdrawMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateWithdraw *CreateWithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreateWithdraw.Contract.CreateWithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateWithdraw *CreateWithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.CreateWithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateWithdraw *CreateWithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.CreateWithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateWithdraw *CreateWithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreateWithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateWithdraw *CreateWithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateWithdraw *CreateWithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_CreateWithdraw *CreateWithdrawTransactor) Create(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _CreateWithdraw.contract.Transact(opts, "create", _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_CreateWithdraw *CreateWithdrawSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.Create(&_CreateWithdraw.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_CreateWithdraw *CreateWithdrawTransactorSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.Create(&_CreateWithdraw.TransactOpts, _salt, _bytecodeHash, _input)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_CreateWithdraw *CreateWithdrawTransactor) Withdraw(opts *bind.TransactOpts, _l1Receiver common.Address) (*types.Transaction, error) {
	return _CreateWithdraw.contract.Transact(opts, "withdraw", _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_CreateWithdraw *CreateWithdrawSession) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.Withdraw(&_CreateWithdraw.TransactOpts, _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_CreateWithdraw *CreateWithdrawTransactorSession) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _CreateWithdraw.Contract.Withdraw(&_CreateWithdraw.TransactOpts, _l1Receiver)
}
