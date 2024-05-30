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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateIdCommitment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FullTree\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"InvalidIdCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidPaginationQuery\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageLimit\",\"type\":\"uint32\"}],\"name\":\"InvalidUserMessageLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"userMessageLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"MemberRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Q\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_SIZE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployedBlockNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"startIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endIndex\",\"type\":\"uint32\"}],\"name\":\"getCommitments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCommitmentIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"idCommitmentToMetadata\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imtData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"maxIndex\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"numberOfLeaves\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxMessageLimit\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"isValidCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"userMessageLimit\",\"type\":\"uint32\"}],\"name\":\"isValidUserMessageLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"}],\"name\":\"memberExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"userMessageLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"index\",\"type\":\"uint40\"}],\"name\":\"merkleProofElements\",\"outputs\":[{\"internalType\":\"uint256[20]\",\"name\":\"\",\"type\":\"uint256[20]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idCommitment\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"userMessageLimit\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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
// Solidity: function DEPTH() view returns(uint8)
func (_Contract *ContractCaller) DEPTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DEPTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint8)
func (_Contract *ContractSession) DEPTH() (uint8, error) {
	return _Contract.Contract.DEPTH(&_Contract.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint8)
func (_Contract *ContractCallerSession) DEPTH() (uint8, error) {
	return _Contract.Contract.DEPTH(&_Contract.CallOpts)
}

// MAXMESSAGELIMIT is a free data retrieval call binding the contract method 0x09aeb04c.
//
// Solidity: function MAX_MESSAGE_LIMIT() view returns(uint32)
func (_Contract *ContractCaller) MAXMESSAGELIMIT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_MESSAGE_LIMIT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXMESSAGELIMIT is a free data retrieval call binding the contract method 0x09aeb04c.
//
// Solidity: function MAX_MESSAGE_LIMIT() view returns(uint32)
func (_Contract *ContractSession) MAXMESSAGELIMIT() (uint32, error) {
	return _Contract.Contract.MAXMESSAGELIMIT(&_Contract.CallOpts)
}

// MAXMESSAGELIMIT is a free data retrieval call binding the contract method 0x09aeb04c.
//
// Solidity: function MAX_MESSAGE_LIMIT() view returns(uint32)
func (_Contract *ContractCallerSession) MAXMESSAGELIMIT() (uint32, error) {
	return _Contract.Contract.MAXMESSAGELIMIT(&_Contract.CallOpts)
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
// Solidity: function SET_SIZE() view returns(uint32)
func (_Contract *ContractCaller) SETSIZE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SET_SIZE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint32)
func (_Contract *ContractSession) SETSIZE() (uint32, error) {
	return _Contract.Contract.SETSIZE(&_Contract.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint32)
func (_Contract *ContractCallerSession) SETSIZE() (uint32, error) {
	return _Contract.Contract.SETSIZE(&_Contract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
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

// GetCommitments is a free data retrieval call binding the contract method 0x679537f9.
//
// Solidity: function getCommitments(uint32 startIndex, uint32 endIndex) view returns(uint256[])
func (_Contract *ContractCaller) GetCommitments(opts *bind.CallOpts, startIndex uint32, endIndex uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCommitments", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCommitments is a free data retrieval call binding the contract method 0x679537f9.
//
// Solidity: function getCommitments(uint32 startIndex, uint32 endIndex) view returns(uint256[])
func (_Contract *ContractSession) GetCommitments(startIndex uint32, endIndex uint32) ([]*big.Int, error) {
	return _Contract.Contract.GetCommitments(&_Contract.CallOpts, startIndex, endIndex)
}

// GetCommitments is a free data retrieval call binding the contract method 0x679537f9.
//
// Solidity: function getCommitments(uint32 startIndex, uint32 endIndex) view returns(uint256[])
func (_Contract *ContractCallerSession) GetCommitments(startIndex uint32, endIndex uint32) ([]*big.Int, error) {
	return _Contract.Contract.GetCommitments(&_Contract.CallOpts, startIndex, endIndex)
}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint32)
func (_Contract *ContractCaller) IdCommitmentIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "idCommitmentIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint32)
func (_Contract *ContractSession) IdCommitmentIndex() (uint32, error) {
	return _Contract.Contract.IdCommitmentIndex(&_Contract.CallOpts)
}

// IdCommitmentIndex is a free data retrieval call binding the contract method 0xae74552a.
//
// Solidity: function idCommitmentIndex() view returns(uint32)
func (_Contract *ContractCallerSession) IdCommitmentIndex() (uint32, error) {
	return _Contract.Contract.IdCommitmentIndex(&_Contract.CallOpts)
}

// IdCommitmentToMetadata is a free data retrieval call binding the contract method 0x9ac21345.
//
// Solidity: function idCommitmentToMetadata(uint256 idCommitment) view returns(uint32, uint32, uint256)
func (_Contract *ContractCaller) IdCommitmentToMetadata(opts *bind.CallOpts, idCommitment *big.Int) (uint32, uint32, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "idCommitmentToMetadata", idCommitment)

	if err != nil {
		return *new(uint32), *new(uint32), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// IdCommitmentToMetadata is a free data retrieval call binding the contract method 0x9ac21345.
//
// Solidity: function idCommitmentToMetadata(uint256 idCommitment) view returns(uint32, uint32, uint256)
func (_Contract *ContractSession) IdCommitmentToMetadata(idCommitment *big.Int) (uint32, uint32, *big.Int, error) {
	return _Contract.Contract.IdCommitmentToMetadata(&_Contract.CallOpts, idCommitment)
}

// IdCommitmentToMetadata is a free data retrieval call binding the contract method 0x9ac21345.
//
// Solidity: function idCommitmentToMetadata(uint256 idCommitment) view returns(uint32, uint32, uint256)
func (_Contract *ContractCallerSession) IdCommitmentToMetadata(idCommitment *big.Int) (uint32, uint32, *big.Int, error) {
	return _Contract.Contract.IdCommitmentToMetadata(&_Contract.CallOpts, idCommitment)
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

// IsValidUserMessageLimit is a free data retrieval call binding the contract method 0xa45d5e59.
//
// Solidity: function isValidUserMessageLimit(uint32 userMessageLimit) view returns(bool)
func (_Contract *ContractCaller) IsValidUserMessageLimit(opts *bind.CallOpts, userMessageLimit uint32) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isValidUserMessageLimit", userMessageLimit)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidUserMessageLimit is a free data retrieval call binding the contract method 0xa45d5e59.
//
// Solidity: function isValidUserMessageLimit(uint32 userMessageLimit) view returns(bool)
func (_Contract *ContractSession) IsValidUserMessageLimit(userMessageLimit uint32) (bool, error) {
	return _Contract.Contract.IsValidUserMessageLimit(&_Contract.CallOpts, userMessageLimit)
}

// IsValidUserMessageLimit is a free data retrieval call binding the contract method 0xa45d5e59.
//
// Solidity: function isValidUserMessageLimit(uint32 userMessageLimit) view returns(bool)
func (_Contract *ContractCallerSession) IsValidUserMessageLimit(userMessageLimit uint32) (bool, error) {
	return _Contract.Contract.IsValidUserMessageLimit(&_Contract.CallOpts, userMessageLimit)
}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 idCommitment) view returns(bool)
func (_Contract *ContractCaller) MemberExists(opts *bind.CallOpts, idCommitment *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "memberExists", idCommitment)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 idCommitment) view returns(bool)
func (_Contract *ContractSession) MemberExists(idCommitment *big.Int) (bool, error) {
	return _Contract.Contract.MemberExists(&_Contract.CallOpts, idCommitment)
}

// MemberExists is a free data retrieval call binding the contract method 0x6bdcc8ab.
//
// Solidity: function memberExists(uint256 idCommitment) view returns(bool)
func (_Contract *ContractCallerSession) MemberExists(idCommitment *big.Int) (bool, error) {
	return _Contract.Contract.MemberExists(&_Contract.CallOpts, idCommitment)
}

// MemberInfo is a free data retrieval call binding the contract method 0xd90d0ee6.
//
// Solidity: function memberInfo(uint256 ) view returns(uint32 userMessageLimit, uint32 index)
func (_Contract *ContractCaller) MemberInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserMessageLimit uint32
	Index            uint32
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "memberInfo", arg0)

	outstruct := new(struct {
		UserMessageLimit uint32
		Index            uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserMessageLimit = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Index = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// MemberInfo is a free data retrieval call binding the contract method 0xd90d0ee6.
//
// Solidity: function memberInfo(uint256 ) view returns(uint32 userMessageLimit, uint32 index)
func (_Contract *ContractSession) MemberInfo(arg0 *big.Int) (struct {
	UserMessageLimit uint32
	Index            uint32
}, error) {
	return _Contract.Contract.MemberInfo(&_Contract.CallOpts, arg0)
}

// MemberInfo is a free data retrieval call binding the contract method 0xd90d0ee6.
//
// Solidity: function memberInfo(uint256 ) view returns(uint32 userMessageLimit, uint32 index)
func (_Contract *ContractCallerSession) MemberInfo(arg0 *big.Int) (struct {
	UserMessageLimit uint32
	Index            uint32
}, error) {
	return _Contract.Contract.MemberInfo(&_Contract.CallOpts, arg0)
}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[20])
func (_Contract *ContractCaller) MerkleProofElements(opts *bind.CallOpts, index *big.Int) ([20]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "merkleProofElements", index)

	if err != nil {
		return *new([20]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([20]*big.Int)).(*[20]*big.Int)

	return out0, err

}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[20])
func (_Contract *ContractSession) MerkleProofElements(index *big.Int) ([20]*big.Int, error) {
	return _Contract.Contract.MerkleProofElements(&_Contract.CallOpts, index)
}

// MerkleProofElements is a free data retrieval call binding the contract method 0x74e942fa.
//
// Solidity: function merkleProofElements(uint40 index) view returns(uint256[20])
func (_Contract *ContractCallerSession) MerkleProofElements(index *big.Int) ([20]*big.Int, error) {
	return _Contract.Contract.MerkleProofElements(&_Contract.CallOpts, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address initialOwner, uint32 maxMessageLimit) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, maxMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", initialOwner, maxMessageLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address initialOwner, uint32 maxMessageLimit) returns()
func (_Contract *ContractSession) Initialize(initialOwner common.Address, maxMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner, maxMessageLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address initialOwner, uint32 maxMessageLimit) returns()
func (_Contract *ContractTransactorSession) Initialize(initialOwner common.Address, maxMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner, maxMessageLimit)
}

// Register is a paid mutator transaction binding the contract method 0xaf7b4210.
//
// Solidity: function register(uint256 idCommitment, uint32 userMessageLimit) returns()
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, idCommitment *big.Int, userMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", idCommitment, userMessageLimit)
}

// Register is a paid mutator transaction binding the contract method 0xaf7b4210.
//
// Solidity: function register(uint256 idCommitment, uint32 userMessageLimit) returns()
func (_Contract *ContractSession) Register(idCommitment *big.Int, userMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, idCommitment, userMessageLimit)
}

// Register is a paid mutator transaction binding the contract method 0xaf7b4210.
//
// Solidity: function register(uint256 idCommitment, uint32 userMessageLimit) returns()
func (_Contract *ContractTransactorSession) Register(idCommitment *big.Int, userMessageLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, idCommitment, userMessageLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	IdCommitment     *big.Int
	UserMessageLimit uint32
	Index            uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMemberRegistered is a free log retrieval operation binding the contract event 0x51ddaace4c9a58c725a5debea8eb4f4afc5085bd6c14c45c97e728d8f54c16df.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint32 userMessageLimit, uint32 index)
func (_Contract *ContractFilterer) FilterMemberRegistered(opts *bind.FilterOpts) (*ContractMemberRegisteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractMemberRegisteredIterator{contract: _Contract.contract, event: "MemberRegistered", logs: logs, sub: sub}, nil
}

// WatchMemberRegistered is a free log subscription operation binding the contract event 0x51ddaace4c9a58c725a5debea8eb4f4afc5085bd6c14c45c97e728d8f54c16df.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint32 userMessageLimit, uint32 index)
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

// ParseMemberRegistered is a log parse operation binding the contract event 0x51ddaace4c9a58c725a5debea8eb4f4afc5085bd6c14c45c97e728d8f54c16df.
//
// Solidity: event MemberRegistered(uint256 idCommitment, uint32 userMessageLimit, uint32 index)
func (_Contract *ContractFilterer) ParseMemberRegistered(log types.Log) (*ContractMemberRegistered, error) {
	event := new(ContractMemberRegistered)
	if err := _Contract.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Contract contract.
type ContractUpgradedIterator struct {
	Event *ContractUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpgraded)
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
		it.Event = new(ContractUpgraded)
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
func (it *ContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpgraded represents a Upgraded event raised by the Contract contract.
type ContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpgradedIterator{contract: _Contract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpgraded)
				if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) ParseUpgraded(log types.Log) (*ContractUpgraded, error) {
	event := new(ContractUpgraded)
	if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
