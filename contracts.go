// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"}]"

// TestBin is the compiled bytecode used for deploying new contracts.
const TestBin = `0x6060604052341561000c57fe5b5b6101048061001c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c55699c81146043578063a56dfe4a146062575bfe5b3415604a57fe5b60506081565b60408051918252519081900360200190f35b3415606957fe5b605060af565b60408051918252519081900360200190f35b600073af15d30a2742e90a929bb6370e5862880f62674169d3b62ef245430363e1a180825b0492505b505090565b69d3b62ef245430363e1a173af15d30a2742e90a929bb6370e5862880f62674181805b505050905600a165627a7a723058209ae49f55befe6a46f20209f4f8b92d70a122f959d6994acdcdd4f2a82194303c0029`

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_Test *TestCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Test.contract.Call(opts, out, "x")
	return *ret0, err
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_Test *TestSession) X() (*big.Int, error) {
	return _Test.Contract.X(&_Test.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_Test *TestCallerSession) X() (*big.Int, error) {
	return _Test.Contract.X(&_Test.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_Test *TestCaller) Y(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Test.contract.Call(opts, out, "y")
	return *ret0, err
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_Test *TestSession) Y() (*big.Int, error) {
	return _Test.Contract.Y(&_Test.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_Test *TestCallerSession) Y() (*big.Int, error) {
	return _Test.Contract.Y(&_Test.CallOpts)
}
