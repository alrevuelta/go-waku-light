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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"membershipDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DuplicateIdCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FullTree\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientContractBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientWithdrawalBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"InvalidIdCommitment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidPaginationQuery\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InvalidReceiverAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"MemberHasNoStake\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"MemberNotRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MemberRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MemberWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MEMBERSHIP_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Q\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployedBlockNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getCommitments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCommitmentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imtData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"maxIndex\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"numberOfLeaves\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexToCommitment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"isValidCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"index\",\"type\":\"uint40\"}],\"name\":\"merkleProofElements\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Contract *ContractCaller) DEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DEPTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Contract *ContractSession) DEPTH() (*big.Int, error) {
	return _Contract.Contract.DEPTH(&_Contract.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Contract *ContractCallerSession) DEPTH() (*big.Int, error) {
	return _Contract.Contract.DEPTH(&_Contract.CallOpts)
}

// MEMBERSHIPDEPOSIT is a free data retrieval call binding the contract method 0xf220b9ec.
//
// Solidity: function MEMBERSHIP_DEPOSIT() view returns(uint256)
func (_Contract *ContractCaller) MEMBERSHIPDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MEMBERSHIP_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MEMBERSHIPDEPOSIT is a free data retrieval call binding the contract method 0xf220b9ec.
//
// Solidity: function MEMBERSHIP_DEPOSIT() view returns(uint256)
func (_Contract *ContractSession) MEMBERSHIPDEPOSIT() (*big.Int, error) {
	return _Contract.Contract.MEMBERSHIPDEPOSIT(&_Contract.CallOpts)
}

// MEMBERSHIPDEPOSIT is a free data retrieval call binding the contract method 0xf220b9ec.
//
// Solidity: function MEMBERSHIP_DEPOSIT() view returns(uint256)
func (_Contract *ContractCallerSession) MEMBERSHIPDEPOSIT() (*big.Int, error) {
	return _Contract.Contract.MEMBERSHIPDEPOSIT(&_Contract.CallOpts)
}

// Q is a free data retrieval call binding the contract method 0xe493ef8c.
//
// Solidity: function Q() view returns(uint256)
func (_Contract *ContractCaller) Q(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Q")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Q is a free data retrieval call binding the contract method 0xe493ef8c.
//
// Solidity: function Q() view returns(uint256)
func (_Contract *ContractSession) Q() (*big.Int, error) {
	return _Contract.Contract.Q(&_Contract.CallOpts)
}

// Q is a free data retrieval call binding the contract method 0xe493ef8c.
//
// Solidity: function Q() view returns(uint256)
func (_Contract *ContractCallerSession) Q() (*big.Int, error) {
	return _Contract.Contract.Q(&_Contract.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Contract *ContractCaller) SETSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SET_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Contract *ContractSession) SETSIZE() (*big.Int, error) {
	return _Contract.Contract.SETSIZE(&_Contract.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Contract *ContractCallerSession) SETSIZE() (*big.Int, error) {
	return _Contract.Contract.SETSIZE(&_Contract.CallOpts)
}

// DeployedBlockNumber is a free data retrieval call binding the contract method 0x4add651e.
//
// Solidity: function deployedBlockNumber() view returns(uint32)
func (_Contract *ContractCaller) DeployedBlockNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "deployedBlockNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DeployedBlockNumber is a free data retrieval call binding the contract method 0x4add651e.
//
// Solidity: function deployedBlockNumber() view returns(uint32)
func (_Contract *ContractSession) DeployedBlockNumber() (uint32, error) {
	return _Contract.Contract.DeployedBlockNumber(&_Contract.CallOpts)
}

// DeployedBlockNumber is a free data retrieval call binding the contract method 0x4add651e.
//
// Solidity: function deployedBlockNumber() view returns(uint32)
func (_Contract *ContractCallerSession) DeployedBlockNumber() (uint32, error) {
	return _Contract.Contract.DeployedBlockNumber(&_Contract.CallOpts)
}

// GetCommitments is a free data retrieval call binding the contract method 0x933ebfdd.
//
// Solidity: function getCommitments(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Contract *ContractCaller) GetCommitments(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCommitments", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCommitments is a free data retrieval call binding the contract method 0x933ebfdd.
//
// Solidity: function getCommitments(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Contract *ContractSession) GetCommitments(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetCommitments(&_Contract.CallOpts, startIndex, endIndex)
}

// GetCommitments is a free data retrieval call binding the contract method 0x933ebfdd.
//
// Solidity: function getCommitments(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Contract *ContractCallerSession) GetCommitments(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetCommitments(&_Contract.CallOpts, startIndex, endIndex)
}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint256)
func (_Contract *ContractCaller) IdCommitmentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "idCommitmentIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint256)
func (_Contract *ContractSession) IdCommitmentIndex() (*big.Int, error) {
	return _Contract.Contract.IdCommitmentIndex(&_Contract.CallOpts)
}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint256)
func (_Contract *ContractCallerSession) IdCommitmentIndex() (*big.Int, error) {
	return _Contract.Contract.IdCommitmentIndex(&_Contract.CallOpts)
}

// ImtData is a free data retrieval call binding the contract method 0x3c979b5f.
//
// Solidity: function imtData() view returns(uint40 maxIndex, uint40 numberOfLeaves)
func (_Contract *ContractCaller) ImtData(opts *bind.CallOpts) (struct {
	MaxIndex       *big.Int
	NumberOfLeaves *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "imtData")

	outstruct := new(struct {
		MaxIndex       *big.Int
		NumberOfLeaves *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NumberOfLeaves = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ImtData is a free data retrieval call binding the contract method 0x3c979b5f.
//
// Solidity: function imtData() view returns(uint40 maxIndex, uint40 numberOfLeaves)
func (_Contract *ContractSession) ImtData() (struct {
	MaxIndex       *big.Int
	NumberOfLeaves *big.Int
}, error) {
	return _Contract.Contract.ImtData(&_Contract.CallOpts)
}

// ImtData is a free data retrieval call binding the contract method 0x3c979b5f.
//
// Solidity: function imtData() view returns(uint40 maxIndex, uint40 numberOfLeaves)
func (_Contract *ContractCallerSession) ImtData() (struct {
	MaxIndex       *big.Int
	NumberOfLeaves *big.Int
}, error) {
	return _Contract.Contract.ImtData(&_Contract.CallOpts)
}

// IndexToCommitment is a free data retrieval call binding the contract method 0x7671ac05.
//
// Solidity: function indexToCommitment(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) IndexToCommitment(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "indexToCommitment", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexToCommitment is a free data retrieval call binding the contract method 0x7671ac05.
//
// Solidity: function indexToCommitment(uint256 ) view returns(uint256)
func (_Contract *ContractSession) IndexToCommitment(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.IndexToCommitment(&_Contract.CallOpts, arg0)
}

// IndexToCommitment is a free data retrieval call binding the contract method 0x7671ac05.
//
// Solidity: function indexToCommitment(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) IndexToCommitment(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.IndexToCommitment(&_Contract.CallOpts, arg0)
}

// IsValidCommitment is a free data retrieval call binding the contract method 0x22d9730c.
//
// Solidity: function isValidCommitment(uint256 idCommitment) pure returns(bool)
func (_Contract *ContractCaller) IsValidCommitment(opts *bind.CallOpts, idCommitment *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isValidCommitment", idCommitment)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidCommitment is a free data retrieval call binding the contract method 0x22d9730c.
//
// Solidity: function isValidCommitment(uint256 idCommitment) pure returns(bool)
func (_Contract *ContractSession) IsValidCommitment(idCommitment *big.Int) (bool, error) {
	return _Contract.Contract.IsValidCommitment(&_Contract.CallOpts, idCommitment)
}

// IsValidCommitment is a free data retrieval call binding the contract method 0x22d9730c.
//
// Solidity: function isValidCommitment(uint256 idCommitment) pure returns(bool)
func (_Contract *ContractCallerSession) IsValidCommitment(idCommitment *big.Int) (bool, error) {
	return _Contract.Contract.IsValidCommitment(&_Contract.CallOpts, idCommitment)
}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 ) view returns(bool)
func (_Contract *ContractCaller) MemberExists(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "memberExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 ) view returns(bool)
func (_Contract *ContractSession) MemberExists(arg0 *big.Int) (bool, error) {
	return _Contract.Contract.MemberExists(&_Contract.CallOpts, arg0)
}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 ) view returns(bool)
func (_Contract *ContractCallerSession) MemberExists(arg0 *big.Int) (bool, error) {
	return _Contract.Contract.MemberExists(&_Contract.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(uint256)
func (_Contract *ContractSession) Members(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Members(&_Contract.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) Members(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Members(&_Contract.CallOpts, arg0)
}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[])
func (_Contract *ContractCaller) MerkleProofElements(opts *bind.CallOpts, index *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "merkleProofElements", index)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[])
func (_Contract *ContractSession) MerkleProofElements(index *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.MerkleProofElements(&_Contract.CallOpts, index)
}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[])
func (_Contract *ContractCallerSession) MerkleProofElements(index *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.MerkleProofElements(&_Contract.CallOpts, index)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(uint256)
func (_Contract *ContractCaller) Root(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(uint256)
func (_Contract *ContractSession) Root() (*big.Int, error) {
	return _Contract.Contract.Root(&_Contract.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(uint256)
func (_Contract *ContractCallerSession) Root() (*big.Int, error) {
	return _Contract.Contract.Root(&_Contract.CallOpts)
}

// StakedAmounts is a free data retrieval call binding the contract method 0xbc499128.
//
// Solidity: function stakedAmounts(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) StakedAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0xbc499128.
//
// Solidity: function stakedAmounts(uint256 ) view returns(uint256)
func (_Contract *ContractSession) StakedAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StakedAmounts(&_Contract.CallOpts, arg0)
}

// StakedAmounts is a free data retrieval call binding the contract method 0xbc499128.
//
// Solidity: function stakedAmounts(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) StakedAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StakedAmounts(&_Contract.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Contract *ContractCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Contract *ContractSession) Verifier() (common.Address, error) {
	return _Contract.Contract.Verifier(&_Contract.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Contract *ContractCallerSession) Verifier() (common.Address, error) {
	return _Contract.Contract.Verifier(&_Contract.CallOpts)
}

// WithdrawalBalance is a free data retrieval call binding the contract method 0xc5b208ff.
//
// Solidity: function withdrawalBalance(address ) view returns(uint256)
func (_Contract *ContractCaller) WithdrawalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawalBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalBalance is a free data retrieval call binding the contract method 0xc5b208ff.
//
// Solidity: function withdrawalBalance(address ) view returns(uint256)
func (_Contract *ContractSession) WithdrawalBalance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.WithdrawalBalance(&_Contract.CallOpts, arg0)
}

// WithdrawalBalance is a free data retrieval call binding the contract method 0xc5b208ff.
//
// Solidity: function withdrawalBalance(address ) view returns(uint256)
func (_Contract *ContractCallerSession) WithdrawalBalance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.WithdrawalBalance(&_Contract.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(uint256 idCommitment) payable returns()
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, idCommitment *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", idCommitment)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(uint256 idCommitment) payable returns()
func (_Contract *ContractSession) Register(idCommitment *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, idCommitment)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(uint256 idCommitment) payable returns()
func (_Contract *ContractTransactorSession) Register(idCommitment *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, idCommitment)
}

// Slash is a paid mutator transaction binding the contract method 0x8be9b119.
//
// Solidity: function slash(uint256 idCommitment, address receiver, uint256[8] proof) returns()
func (_Contract *ContractTransactor) Slash(opts *bind.TransactOpts, idCommitment *big.Int, receiver common.Address, proof [8]*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "slash", idCommitment, receiver, proof)
}

// Slash is a paid mutator transaction binding the contract method 0x8be9b119.
//
// Solidity: function slash(uint256 idCommitment, address receiver, uint256[8] proof) returns()
func (_Contract *ContractSession) Slash(idCommitment *big.Int, receiver common.Address, proof [8]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Slash(&_Contract.TransactOpts, idCommitment, receiver, proof)
}

// Slash is a paid mutator transaction binding the contract method 0x8be9b119.
//
// Solidity: function slash(uint256 idCommitment, address receiver, uint256[8] proof) returns()
func (_Contract *ContractTransactorSession) Slash(idCommitment *big.Int, receiver common.Address, proof [8]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Slash(&_Contract.TransactOpts, idCommitment, receiver, proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// ContractMemberRegisteredIterator is returned from FilterMemberRegistered and is used to iterate over the raw logs and unpacked data for MemberRegistered events raised by the Contract contract.
type ContractMemberRegisteredIterator struct {
	Event *ContractMemberRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMemberRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMemberRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMemberRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMemberRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMemberRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMemberRegistered represents a MemberRegistered event raised by the Contract contract.
type ContractMemberRegistered struct {
	IdCommitment *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMemberRegistered is a free log retrieval operation binding the contract event 0x5a92c2530f207992057b9c3e544108ffce3beda4a63719f316967c49bf6159d2.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) FilterMemberRegistered(opts *bind.FilterOpts) (*ContractMemberRegisteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractMemberRegisteredIterator{contract: _Contract.contract, event: "MemberRegistered", logs: logs, sub: sub}, nil
}

// WatchMemberRegistered is a free log subscription operation binding the contract event 0x5a92c2530f207992057b9c3e544108ffce3beda4a63719f316967c49bf6159d2.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) WatchMemberRegistered(opts *bind.WatchOpts, sink chan<- *ContractMemberRegistered) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMemberRegistered)
				if err := _Contract.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMemberRegistered is a log parse operation binding the contract event 0x5a92c2530f207992057b9c3e544108ffce3beda4a63719f316967c49bf6159d2.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) ParseMemberRegistered(log types.Log) (*ContractMemberRegistered, error) {
	event := new(ContractMemberRegistered)
	if err := _Contract.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMemberWithdrawnIterator is returned from FilterMemberWithdrawn and is used to iterate over the raw logs and unpacked data for MemberWithdrawn events raised by the Contract contract.
type ContractMemberWithdrawnIterator struct {
	Event *ContractMemberWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMemberWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMemberWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMemberWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMemberWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMemberWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMemberWithdrawn represents a MemberWithdrawn event raised by the Contract contract.
type ContractMemberWithdrawn struct {
	IdCommitment *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMemberWithdrawn is a free log retrieval operation binding the contract event 0x62ec3a516d22a993ce5cb4e7593e878c74f4d799dde522a88dc27a994fd5a943.
//
// Solidity: event MemberWithdrawn(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) FilterMemberWithdrawn(opts *bind.FilterOpts) (*ContractMemberWithdrawnIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MemberWithdrawn")
	if err != nil {
		return nil, err
	}
	return &ContractMemberWithdrawnIterator{contract: _Contract.contract, event: "MemberWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMemberWithdrawn is a free log subscription operation binding the contract event 0x62ec3a516d22a993ce5cb4e7593e878c74f4d799dde522a88dc27a994fd5a943.
//
// Solidity: event MemberWithdrawn(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) WatchMemberWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractMemberWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MemberWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMemberWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "MemberWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMemberWithdrawn is a log parse operation binding the contract event 0x62ec3a516d22a993ce5cb4e7593e878c74f4d799dde522a88dc27a994fd5a943.
//
// Solidity: event MemberWithdrawn(uint256 idCommitment, uint256 index)
func (_Contract *ContractFilterer) ParseMemberWithdrawn(log types.Log) (*ContractMemberWithdrawn, error) {
	event := new(ContractMemberWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "MemberWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
