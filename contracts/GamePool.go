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
)

// GamePoolPlayer is an auto generated low-level Go binding around an user-defined struct.
type GamePoolPlayer struct {
	Amount        *big.Int
	PlayerAddress common.Address
}

// GamePoolMetaData contains all meta data concerning the GamePool contract.
var GamePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"gameId\",\"type\":\"string\"}],\"name\":\"GenesisPlayer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"gameId\",\"type\":\"string\"}],\"name\":\"OtherPlayer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"gameId\",\"type\":\"string\"}],\"name\":\"ReturnAGame\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"}],\"internalType\":\"structGamePool.Player[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gameId\",\"type\":\"string\"}],\"name\":\"WithdrawToWin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611197806100606000396000f3fe60806040526004361061004a5760003560e01c8063263d4ef71461004f5780633e6ba9e01461006b5780634756b105146100875780638cca55e8146100c4578063e2d49bb0146100e0575b600080fd5b61006960048036036100649190810190610c4a565b61011e565b005b61008560048036036100809190810190610c4a565b61033d565b005b34801561009357600080fd5b506100ae60048036036100a99190810190610c05565b610543565b6040516100bb9190610f06565b60405180910390f35b6100de60048036036100d99190810190610bad565b610623565b005b3480156100ec57600080fd5b5061010760048036036101029190810190610c8b565b6109ab565b604051610115929190610f88565b60405180910390f35b3460018260405161012f9190610eef565b908152602001604051809103902060008154811061014957fe5b9060005260206000209060020201600001541461016557600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156101bc57600080fd5b6001816040516101cc9190610eef565b908152602001604051809103902060405180604001604052803481526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506002600081548092919060010191905055503373ffffffffffffffffffffffffffffffffffffffff167f0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e060405161033290610f48565b60405180910390a250565b600060018260405161034f9190610eef565b9081526020016040518091039020805490501461036b57600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156103c257600080fd5b6001816040516103d29190610eef565b908152602001604051809103902060405180604001604052803481526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506002600081548092919060010191905055503373ffffffffffffffffffffffffffffffffffffffff167f0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e060405161053890610f68565b60405180910390a250565b606060018383604051610557929190610ed6565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156106175783829060005260206000209060020201604051806040016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505081526020019060010190610585565b50505050905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461067c57600080fd5b600080905060008090505b60018484604051610699929190610ed6565b9081526020016040518091039020805490508160ff161015610869578473ffffffffffffffffffffffffffffffffffffffff16600185856040516106de929190610ed6565b90815260200160405180910390208260ff16815481106106fa57fe5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156107a557600191506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b600060036000600187876040516107bd929190610ed6565b90815260200160405180910390208460ff16815481106107d957fe5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050610687565b5080156109a5578373ffffffffffffffffffffffffffffffffffffffff166108fc6001858560405161089c929190610ed6565b908152602001604051809103902080549050600186866040516108c0929190610ed6565b90815260200160405180910390206000815481106108da57fe5b906000526020600020906002020160000154029081150290604051600060405180830381858888f19350505050158015610918573d6000803e3d6000fd5b506001838360405161092b929190610ed6565b908152602001604051809103902060006109459190610a1f565b600260008154809291906001900391905055503373ffffffffffffffffffffffffffffffffffffffff167f0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e060405161099c90610f28565b60405180910390a25b50505050565b60018280516020810182018051848252602083016020850120818352809550505050505081815481106109da57fe5b9060005260206000209060020201600091509150508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b5080546000825560020290600052602060002090810190610a409190610a43565b50565b610a8e91905b80821115610a8a576000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600201610a49565b5090565b90565b600081359050610aa081611126565b92915050565b60008083601f840112610ab857600080fd5b8235905067ffffffffffffffff811115610ad157600080fd5b602083019150836001820283011115610ae957600080fd5b9250929050565b600082601f830112610b0157600080fd5b8135610b14610b0f82610fde565b610fb1565b91508082526020830160208301858383011115610b3057600080fd5b610b3b8382846110e4565b50505092915050565b600082601f830112610b5557600080fd5b8135610b68610b638261100a565b610fb1565b91508082526020830160208301858383011115610b8457600080fd5b610b8f8382846110e4565b50505092915050565b600081359050610ba78161113d565b92915050565b600080600060408486031215610bc257600080fd5b6000610bd086828701610a91565b935050602084013567ffffffffffffffff811115610bed57600080fd5b610bf986828701610aa6565b92509250509250925092565b60008060208385031215610c1857600080fd5b600083013567ffffffffffffffff811115610c3257600080fd5b610c3e85828601610aa6565b92509250509250929050565b600060208284031215610c5c57600080fd5b600082013567ffffffffffffffff811115610c7657600080fd5b610c8284828501610b44565b91505092915050565b60008060408385031215610c9e57600080fd5b600083013567ffffffffffffffff811115610cb857600080fd5b610cc485828601610af0565b9250506020610cd585828601610b98565b9150509250929050565b6000610ceb8383610e89565b60408301905092915050565b610d0081611096565b82525050565b610d0f81611096565b82525050565b6000610d2082611046565b610d2a8185611069565b9350610d3583611036565b8060005b83811015610d66578151610d4d8882610cdf565b9750610d588361105c565b925050600181019050610d39565b5085935050505092915050565b6000610d7f838561108b565b9350610d8c8385846110e4565b82840190509392505050565b6000610da382611051565b610dad818561108b565b9350610dbd8185602086016110f3565b80840191505092915050565b6000610dd660168361107a565b91507f596f75206861766520776f6e207468652067616d6521000000000000000000006000830152602082019050919050565b6000610e1660138361107a565b91507f596f75206a6f696e6564207468652067616d65000000000000000000000000006000830152602082019050919050565b6000610e5660198361107a565b91507f47616d652063726561746564207375636365737366756c6c79000000000000006000830152602082019050919050565b604082016000820151610e9f6000850182610eb8565b506020820151610eb26020850182610cf7565b50505050565b610ec1816110da565b82525050565b610ed0816110da565b82525050565b6000610ee3828486610d73565b91508190509392505050565b6000610efb8284610d98565b915081905092915050565b60006020820190508181036000830152610f208184610d15565b905092915050565b60006020820190508181036000830152610f4181610dc9565b9050919050565b60006020820190508181036000830152610f6181610e09565b9050919050565b60006020820190508181036000830152610f8181610e49565b9050919050565b6000604082019050610f9d6000830185610ec7565b610faa6020830184610d06565b9392505050565b6000604051905081810181811067ffffffffffffffff82111715610fd457600080fd5b8060405250919050565b600067ffffffffffffffff821115610ff557600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111561102157600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006110a1826110ba565b9050919050565b60006110b3826110ba565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156111115780820151818401526020810190506110f6565b83811115611120576000848401525b50505050565b61112f816110a8565b811461113a57600080fd5b50565b611146816110da565b811461115157600080fd5b5056fea365627a7a72315820a2faea161c48bd7ca72f8c937f8f3163f338582dfda9a641b669bdf3978e37cc6c6578706572696d656e74616cf564736f6c63430005100040",
}

// GamePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use GamePoolMetaData.ABI instead.
var GamePoolABI = GamePoolMetaData.ABI

// GamePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GamePoolMetaData.Bin instead.
var GamePoolBin = GamePoolMetaData.Bin

// DeployGamePool deploys a new Ethereum contract, binding an instance of GamePool to it.
func DeployGamePool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GamePool, error) {
	parsed, err := GamePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GamePoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GamePool{GamePoolCaller: GamePoolCaller{contract: contract}, GamePoolTransactor: GamePoolTransactor{contract: contract}, GamePoolFilterer: GamePoolFilterer{contract: contract}}, nil
}

// GamePool is an auto generated Go binding around an Ethereum contract.
type GamePool struct {
	GamePoolCaller     // Read-only binding to the contract
	GamePoolTransactor // Write-only binding to the contract
	GamePoolFilterer   // Log filterer for contract events
}

// GamePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type GamePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GamePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GamePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GamePoolSession struct {
	Contract     *GamePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GamePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GamePoolCallerSession struct {
	Contract *GamePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GamePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GamePoolTransactorSession struct {
	Contract     *GamePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GamePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type GamePoolRaw struct {
	Contract *GamePool // Generic contract binding to access the raw methods on
}

// GamePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GamePoolCallerRaw struct {
	Contract *GamePoolCaller // Generic read-only contract binding to access the raw methods on
}

// GamePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GamePoolTransactorRaw struct {
	Contract *GamePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGamePool creates a new instance of GamePool, bound to a specific deployed contract.
func NewGamePool(address common.Address, backend bind.ContractBackend) (*GamePool, error) {
	contract, err := bindGamePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GamePool{GamePoolCaller: GamePoolCaller{contract: contract}, GamePoolTransactor: GamePoolTransactor{contract: contract}, GamePoolFilterer: GamePoolFilterer{contract: contract}}, nil
}

// NewGamePoolCaller creates a new read-only instance of GamePool, bound to a specific deployed contract.
func NewGamePoolCaller(address common.Address, caller bind.ContractCaller) (*GamePoolCaller, error) {
	contract, err := bindGamePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GamePoolCaller{contract: contract}, nil
}

// NewGamePoolTransactor creates a new write-only instance of GamePool, bound to a specific deployed contract.
func NewGamePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*GamePoolTransactor, error) {
	contract, err := bindGamePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GamePoolTransactor{contract: contract}, nil
}

// NewGamePoolFilterer creates a new log filterer instance of GamePool, bound to a specific deployed contract.
func NewGamePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*GamePoolFilterer, error) {
	contract, err := bindGamePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GamePoolFilterer{contract: contract}, nil
}

// bindGamePool binds a generic wrapper to an already deployed contract.
func bindGamePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GamePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamePool *GamePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GamePool.Contract.GamePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamePool *GamePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.Contract.GamePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamePool *GamePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamePool.Contract.GamePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamePool *GamePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GamePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamePool *GamePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamePool *GamePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamePool.Contract.contract.Transact(opts, method, params...)
}

// ReturnAGame is a free data retrieval call binding the contract method 0x4756b105.
//
// Solidity: function ReturnAGame(string gameId) view returns((uint256,address)[])
func (_GamePool *GamePoolCaller) ReturnAGame(opts *bind.CallOpts, gameId string) ([]GamePoolPlayer, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "ReturnAGame", gameId)

	if err != nil {
		return *new([]GamePoolPlayer), err
	}

	out0 := *abi.ConvertType(out[0], new([]GamePoolPlayer)).(*[]GamePoolPlayer)

	return out0, err

}

// ReturnAGame is a free data retrieval call binding the contract method 0x4756b105.
//
// Solidity: function ReturnAGame(string gameId) view returns((uint256,address)[])
func (_GamePool *GamePoolSession) ReturnAGame(gameId string) ([]GamePoolPlayer, error) {
	return _GamePool.Contract.ReturnAGame(&_GamePool.CallOpts, gameId)
}

// ReturnAGame is a free data retrieval call binding the contract method 0x4756b105.
//
// Solidity: function ReturnAGame(string gameId) view returns((uint256,address)[])
func (_GamePool *GamePoolCallerSession) ReturnAGame(gameId string) ([]GamePoolPlayer, error) {
	return _GamePool.Contract.ReturnAGame(&_GamePool.CallOpts, gameId)
}

// Games is a free data retrieval call binding the contract method 0xe2d49bb0.
//
// Solidity: function games(string , uint256 ) view returns(uint256 amount, address playerAddress)
func (_GamePool *GamePoolCaller) Games(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Amount        *big.Int
	PlayerAddress common.Address
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "games", arg0, arg1)

	outstruct := new(struct {
		Amount        *big.Int
		PlayerAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PlayerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xe2d49bb0.
//
// Solidity: function games(string , uint256 ) view returns(uint256 amount, address playerAddress)
func (_GamePool *GamePoolSession) Games(arg0 string, arg1 *big.Int) (struct {
	Amount        *big.Int
	PlayerAddress common.Address
}, error) {
	return _GamePool.Contract.Games(&_GamePool.CallOpts, arg0, arg1)
}

// Games is a free data retrieval call binding the contract method 0xe2d49bb0.
//
// Solidity: function games(string , uint256 ) view returns(uint256 amount, address playerAddress)
func (_GamePool *GamePoolCallerSession) Games(arg0 string, arg1 *big.Int) (struct {
	Amount        *big.Int
	PlayerAddress common.Address
}, error) {
	return _GamePool.Contract.Games(&_GamePool.CallOpts, arg0, arg1)
}

// GenesisPlayer is a paid mutator transaction binding the contract method 0x3e6ba9e0.
//
// Solidity: function GenesisPlayer(string gameId) payable returns()
func (_GamePool *GamePoolTransactor) GenesisPlayer(opts *bind.TransactOpts, gameId string) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "GenesisPlayer", gameId)
}

// GenesisPlayer is a paid mutator transaction binding the contract method 0x3e6ba9e0.
//
// Solidity: function GenesisPlayer(string gameId) payable returns()
func (_GamePool *GamePoolSession) GenesisPlayer(gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.GenesisPlayer(&_GamePool.TransactOpts, gameId)
}

// GenesisPlayer is a paid mutator transaction binding the contract method 0x3e6ba9e0.
//
// Solidity: function GenesisPlayer(string gameId) payable returns()
func (_GamePool *GamePoolTransactorSession) GenesisPlayer(gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.GenesisPlayer(&_GamePool.TransactOpts, gameId)
}

// OtherPlayer is a paid mutator transaction binding the contract method 0x263d4ef7.
//
// Solidity: function OtherPlayer(string gameId) payable returns()
func (_GamePool *GamePoolTransactor) OtherPlayer(opts *bind.TransactOpts, gameId string) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "OtherPlayer", gameId)
}

// OtherPlayer is a paid mutator transaction binding the contract method 0x263d4ef7.
//
// Solidity: function OtherPlayer(string gameId) payable returns()
func (_GamePool *GamePoolSession) OtherPlayer(gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.OtherPlayer(&_GamePool.TransactOpts, gameId)
}

// OtherPlayer is a paid mutator transaction binding the contract method 0x263d4ef7.
//
// Solidity: function OtherPlayer(string gameId) payable returns()
func (_GamePool *GamePoolTransactorSession) OtherPlayer(gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.OtherPlayer(&_GamePool.TransactOpts, gameId)
}

// WithdrawToWin is a paid mutator transaction binding the contract method 0x8cca55e8.
//
// Solidity: function WithdrawToWin(address winnerAddress, string gameId) payable returns()
func (_GamePool *GamePoolTransactor) WithdrawToWin(opts *bind.TransactOpts, winnerAddress common.Address, gameId string) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "WithdrawToWin", winnerAddress, gameId)
}

// WithdrawToWin is a paid mutator transaction binding the contract method 0x8cca55e8.
//
// Solidity: function WithdrawToWin(address winnerAddress, string gameId) payable returns()
func (_GamePool *GamePoolSession) WithdrawToWin(winnerAddress common.Address, gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.WithdrawToWin(&_GamePool.TransactOpts, winnerAddress, gameId)
}

// WithdrawToWin is a paid mutator transaction binding the contract method 0x8cca55e8.
//
// Solidity: function WithdrawToWin(address winnerAddress, string gameId) payable returns()
func (_GamePool *GamePoolTransactorSession) WithdrawToWin(winnerAddress common.Address, gameId string) (*types.Transaction, error) {
	return _GamePool.Contract.WithdrawToWin(&_GamePool.TransactOpts, winnerAddress, gameId)
}

// GamePoolLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GamePool contract.
type GamePoolLogIterator struct {
	Event *GamePoolLog // Event containing the contract specifics and raw log

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
func (it *GamePoolLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolLog)
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
		it.Event = new(GamePoolLog)
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
func (it *GamePoolLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolLog represents a Log event raised by the GamePool contract.
type GamePoolLog struct {
	Sender  common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string message)
func (_GamePool *GamePoolFilterer) FilterLog(opts *bind.FilterOpts, sender []common.Address) (*GamePoolLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Log", senderRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolLogIterator{contract: _GamePool.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string message)
func (_GamePool *GamePoolFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GamePoolLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Log", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolLog)
				if err := _GamePool.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string message)
func (_GamePool *GamePoolFilterer) ParseLog(log types.Log) (*GamePoolLog, error) {
	event := new(GamePoolLog)
	if err := _GamePool.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
