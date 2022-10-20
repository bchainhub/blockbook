// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	core "github.com/core-coin/go-core"
	"github.com/core-coin/go-core/accounts/abi"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/core/types"
	"github.com/core-coin/go-core/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = core.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Bounty is an auto generated low-level Go binding around an user-defined struct.
type Bounty struct {
	Owner       common.Address
	Target      common.Address
	Data        []byte
	Reward      *big.Int
	Nonce       *big.Int
	Deadline    *big.Int
	EnergyLimit *big.Int
	Signature   []byte
}

// Cheque is an auto generated low-level Go binding around an user-defined struct.
type Cheque struct {
	Owner     common.Address
	Spender   common.Address
	Amount    *big.Int
	Nonce     *big.Int
	Deadline  *big.Int
	Signature []byte
}

// Field is an auto generated low-level Go binding around an user-defined struct.
type Field struct {
	Role     [32]byte
	Required bool
	Complete bool
}

// Info is an auto generated low-level Go binding around an user-defined struct.
type Info struct {
	Expiration   *big.Int
	User         common.Address
	Fingerprints [][32]byte
}

// Submission is an auto generated low-level Go binding around an user-defined struct.
type Submission struct {
	Role        [32]byte
	Expiration  *big.Int
	User        common.Address
	Removed     bool
	Invalidated bool
}

// AccessControlABI is the input ABI used to generate the binding from.
const AccessControlABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlBin is the compiled bytecode used for deploying new contracts.
var AccessControlBin = "0x608060405234801561001057600080fd5b50336000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff160217905550610e3d806100646000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806389a7fa411161005b57806389a7fa4114610113578063bd8f8ebe1461012f578063e05cab1214610139578063eb8325fb1461015557610088565b80631993cd0b1461008d57806325e8d282146100bd5780634e5a827c146100d9578063614b9887146100f5575b600080fd5b6100a760048036038101906100a2919061095d565b610173565b6040516100b49190610aa4565b60405180910390f35b6100d760048036038101906100d2919061095d565b610187565b005b6100f360048036038101906100ee919061095d565b610229565b005b6100fd6102cb565b60405161010a9190610aa4565b60405180910390f35b61012d6004803603810190610128919061090b565b6102e2565b005b610137610489565b005b610153600480360381019061014e9190610934565b6105b6565b005b61015d61060e565b60405161016a9190610a89565b60405180910390f35b600061017f8383610639565b905092915050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff161461021b576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161021290610b3f565b60405180910390fd5b61022582826106a5565b5050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff16146102bd576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016102b490610b3f565b60405180910390fd5b6102c782826107ef565b5050565b6000600260009054906101000a900460ff16905090565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610376576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161036d90610b3f565b60405180910390fd5b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168175ffffffffffffffffffffffffffffffffffffffffffff16141561040b576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161040290610b5f565b60405180910390fd5b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055507fc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be108160405161047e9190610a89565b60405180910390a150565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff161461051d576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161051490610b3f565b60405180910390fd5b600260009054906101000a900460ff161561056d576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161056490610aff565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba2560405160405180910390a1565b806105c18133610639565b610600576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016105f790610adf565b60405180910390fd5b61060a82336106a5565b5050565b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905090565b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6001600083815260200190815260200160002060008275ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610746576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161073d90610abf565b60405180910390fd5b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850826040516107e39190610a89565b60405180910390a25050565b6107f98282610639565b15610839576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161083090610b1f565b60405180910390fd5b600180600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb826040516108d59190610a89565b60405180910390a25050565b6000813590506108f081610db4565b92915050565b60008135905061090581610dcb565b92915050565b60006020828403121561091d57600080fd5b600061092b848285016108e1565b91505092915050565b60006020828403121561094657600080fd5b6000610954848285016108f6565b91505092915050565b6000806040838503121561097057600080fd5b600061097e858286016108f6565b925050602061098f858286016108e1565b9150509250929050565b6109a281610b90565b82525050565b6109b181610ba2565b82525050565b60006109c4602283610b7f565b91506109cf82610bda565b604082019050919050565b60006109e7602583610b7f565b91506109f282610c29565b604082019050919050565b6000610a0a602583610b7f565b9150610a1582610c78565b604082019050919050565b6000610a2d602683610b7f565b9150610a3882610cc7565b604082019050919050565b6000610a50602583610b7f565b9150610a5b82610d16565b604082019050919050565b6000610a73603583610b7f565b9150610a7e82610d65565b604082019050919050565b6000602082019050610a9e6000830184610999565b92915050565b6000602082019050610ab960008301846109a8565b92915050565b60006020820190508181036000830152610ad8816109b7565b9050919050565b60006020820190508181036000830152610af8816109da565b9050919050565b60006020820190508181036000830152610b18816109fd565b9050919050565b60006020820190508181036000830152610b3881610a20565b9050919050565b60006020820190508181036000830152610b5881610a43565b9050919050565b60006020820190508181036000830152610b7881610a66565b9050919050565b600082825260208201905092915050565b6000610b9b82610bb8565b9050919050565b60008115159050919050565b6000819050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b7f416363657373436f6e74726f6c3a20726f6c65206973206e6f74206772616e7460008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f7564206861766560008201527f20726f6c65000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20636f6e747261637420697320646570726560008201527f6361746564000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c6520697320616c7265616479206760008201527f72616e7465640000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f756c642062652060008201527f61646d696e000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a206e657720616e64206f6c642061646d696e60008201527f732073686f756c6420626520646966666572656e740000000000000000000000602082015250565b610dbd81610b90565b8114610dc857600080fd5b50565b610dd481610bae565b8114610ddf57600080fd5b5056fea2646970667358221220036b9cc948d5e4b30afe16cacb2d688bb1620240ec1e745a2c8450250af5643a64736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058"

// DeployAccessControl deploys a new Core contract, binding an instance of AccessControl to it.
func DeployAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// AccessControl is an auto generated Go binding around an Core contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Core contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Core contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Core contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Core contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_AccessControl *AccessControlCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_AccessControl *AccessControlSession) Admin() (common.Address, error) {
	return _AccessControl.Contract.Admin(&_AccessControl.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_AccessControl *AccessControlCallerSession) Admin() (common.Address, error) {
	return _AccessControl.Contract.Admin(&_AccessControl.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_AccessControl *AccessControlCaller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "isDeprecated")
	return *ret0, err
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_AccessControl *AccessControlSession) IsDeprecated() (bool, error) {
	return _AccessControl.Contract.IsDeprecated(&_AccessControl.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_AccessControl *AccessControlCallerSession) IsDeprecated() (bool, error) {
	return _AccessControl.Contract.IsDeprecated(&_AccessControl.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_AccessControl *AccessControlTransactor) ChangeAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "changeAdmin", account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_AccessControl *AccessControlSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.ChangeAdmin(&_AccessControl.TransactOpts, account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_AccessControl *AccessControlTransactorSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.ChangeAdmin(&_AccessControl.TransactOpts, account)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_AccessControl *AccessControlTransactor) DeprecateContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "deprecateContract")
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_AccessControl *AccessControlSession) DeprecateContract() (*types.Transaction, error) {
	return _AccessControl.Contract.DeprecateContract(&_AccessControl.TransactOpts)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_AccessControl *AccessControlTransactorSession) DeprecateContract() (*types.Transaction, error) {
	return _AccessControl.Contract.DeprecateContract(&_AccessControl.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AccessControl contract.
type AccessControlAdminChangedIterator struct {
	Event *AccessControlAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlAdminChanged)
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
		it.Event = new(AccessControlAdminChanged)
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
func (it *AccessControlAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlAdminChanged represents a AdminChanged event raised by the AccessControl contract.
type AccessControlAdminChanged struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_AccessControl *AccessControlFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AccessControlAdminChangedIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AccessControlAdminChangedIterator{contract: _AccessControl.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_AccessControl *AccessControlFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_AccessControl *AccessControlFilterer) ParseAdminChanged(log types.Log) (*AccessControlAdminChanged, error) {
	event := new(AccessControlAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlContractDeprecatedIterator is returned from FilterContractDeprecated and is used to iterate over the raw logs and unpacked data for ContractDeprecated events raised by the AccessControl contract.
type AccessControlContractDeprecatedIterator struct {
	Event *AccessControlContractDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlContractDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlContractDeprecated)
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
		it.Event = new(AccessControlContractDeprecated)
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
func (it *AccessControlContractDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlContractDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlContractDeprecated represents a ContractDeprecated event raised by the AccessControl contract.
type AccessControlContractDeprecated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractDeprecated is a free log retrieval operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_AccessControl *AccessControlFilterer) FilterContractDeprecated(opts *bind.FilterOpts) (*AccessControlContractDeprecatedIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return &AccessControlContractDeprecatedIterator{contract: _AccessControl.contract, event: "ContractDeprecated", logs: logs, sub: sub}, nil
}

// WatchContractDeprecated is a free log subscription operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_AccessControl *AccessControlFilterer) WatchContractDeprecated(opts *bind.WatchOpts, sink chan<- *AccessControlContractDeprecated) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlContractDeprecated)
				if err := _AccessControl.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
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

// ParseContractDeprecated is a log parse operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_AccessControl *AccessControlFilterer) ParseContractDeprecated(log types.Log) (*AccessControlContractDeprecated, error) {
	event := new(AccessControlContractDeprecated)
	if err := _AccessControl.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseTokenABI is the input ABI used to generate the binding from.
const BaseTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BaseToken is an auto generated Go binding around an Core contract.
type BaseToken struct {
	BaseTokenCaller     // Read-only binding to the contract
	BaseTokenTransactor // Write-only binding to the contract
	BaseTokenFilterer   // Log filterer for contract events
}

// BaseTokenCaller is an auto generated read-only Go binding around an Core contract.
type BaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTokenTransactor is an auto generated write-only Go binding around an Core contract.
type BaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type BaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type BaseTokenSession struct {
	Contract     *BaseToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type BaseTokenCallerSession struct {
	Contract *BaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BaseTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type BaseTokenTransactorSession struct {
	Contract     *BaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BaseTokenRaw is an auto generated low-level Go binding around an Core contract.
type BaseTokenRaw struct {
	Contract *BaseToken // Generic contract binding to access the raw methods on
}

// BaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type BaseTokenCallerRaw struct {
	Contract *BaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type BaseTokenTransactorRaw struct {
	Contract *BaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseToken creates a new instance of BaseToken, bound to a specific deployed contract.
func NewBaseToken(address common.Address, backend bind.ContractBackend) (*BaseToken, error) {
	contract, err := bindBaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseToken{BaseTokenCaller: BaseTokenCaller{contract: contract}, BaseTokenTransactor: BaseTokenTransactor{contract: contract}, BaseTokenFilterer: BaseTokenFilterer{contract: contract}}, nil
}

// NewBaseTokenCaller creates a new read-only instance of BaseToken, bound to a specific deployed contract.
func NewBaseTokenCaller(address common.Address, caller bind.ContractCaller) (*BaseTokenCaller, error) {
	contract, err := bindBaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTokenCaller{contract: contract}, nil
}

// NewBaseTokenTransactor creates a new write-only instance of BaseToken, bound to a specific deployed contract.
func NewBaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseTokenTransactor, error) {
	contract, err := bindBaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTokenTransactor{contract: contract}, nil
}

// NewBaseTokenFilterer creates a new log filterer instance of BaseToken, bound to a specific deployed contract.
func NewBaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseTokenFilterer, error) {
	contract, err := bindBaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseTokenFilterer{contract: contract}, nil
}

// bindBaseToken binds a generic wrapper to an already deployed contract.
func bindBaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseToken *BaseTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseToken.Contract.BaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseToken *BaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseToken.Contract.BaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseToken *BaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseToken.Contract.BaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseToken *BaseTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseToken *BaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseToken *BaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseToken *BaseTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseToken *BaseTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseToken.Contract.Allowance(&_BaseToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseToken *BaseTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseToken.Contract.Allowance(&_BaseToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseToken *BaseTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseToken *BaseTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseToken.Contract.BalanceOf(&_BaseToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseToken *BaseTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseToken.Contract.BalanceOf(&_BaseToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseToken *BaseTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseToken *BaseTokenSession) Decimals() (uint8, error) {
	return _BaseToken.Contract.Decimals(&_BaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseToken *BaseTokenCallerSession) Decimals() (uint8, error) {
	return _BaseToken.Contract.Decimals(&_BaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BaseToken *BaseTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BaseToken *BaseTokenSession) Name() (string, error) {
	return _BaseToken.Contract.Name(&_BaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BaseToken *BaseTokenCallerSession) Name() (string, error) {
	return _BaseToken.Contract.Name(&_BaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BaseToken *BaseTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BaseToken *BaseTokenSession) Symbol() (string, error) {
	return _BaseToken.Contract.Symbol(&_BaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BaseToken *BaseTokenCallerSession) Symbol() (string, error) {
	return _BaseToken.Contract.Symbol(&_BaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseToken *BaseTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseToken *BaseTokenSession) TotalSupply() (*big.Int, error) {
	return _BaseToken.Contract.TotalSupply(&_BaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseToken *BaseTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BaseToken.Contract.TotalSupply(&_BaseToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.Approve(&_BaseToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.Approve(&_BaseToken.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.Transfer(&_BaseToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.Transfer(&_BaseToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.TransferFrom(&_BaseToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseToken *BaseTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseToken.Contract.TransferFrom(&_BaseToken.TransactOpts, sender, recipient, amount)
}

// BaseTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BaseToken contract.
type BaseTokenApprovalIterator struct {
	Event *BaseTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTokenApproval)
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
		it.Event = new(BaseTokenApproval)
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
func (it *BaseTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTokenApproval represents a Approval event raised by the BaseToken contract.
type BaseTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseToken *BaseTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BaseTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BaseTokenApprovalIterator{contract: _BaseToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseToken *BaseTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BaseTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTokenApproval)
				if err := _BaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseToken *BaseTokenFilterer) ParseApproval(log types.Log) (*BaseTokenApproval, error) {
	event := new(BaseTokenApproval)
	if err := _BaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BaseToken contract.
type BaseTokenBurnIterator struct {
	Event *BaseTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTokenBurn)
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
		it.Event = new(BaseTokenBurn)
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
func (it *BaseTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTokenBurn represents a Burn event raised by the BaseToken contract.
type BaseTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BaseToken *BaseTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*BaseTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseTokenBurnIterator{contract: _BaseToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BaseToken *BaseTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BaseTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTokenBurn)
				if err := _BaseToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BaseToken *BaseTokenFilterer) ParseBurn(log types.Log) (*BaseTokenBurn, error) {
	event := new(BaseTokenBurn)
	if err := _BaseToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BaseToken contract.
type BaseTokenMintIterator struct {
	Event *BaseTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTokenMint)
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
		it.Event = new(BaseTokenMint)
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
func (it *BaseTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTokenMint represents a Mint event raised by the BaseToken contract.
type BaseTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BaseTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BaseTokenMintIterator{contract: _BaseToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BaseTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTokenMint)
				if err := _BaseToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) ParseMint(log types.Log) (*BaseTokenMint, error) {
	event := new(BaseTokenMint)
	if err := _BaseToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BaseTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BaseToken contract.
type BaseTokenTransferIterator struct {
	Event *BaseTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTokenTransfer)
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
		it.Event = new(BaseTokenTransfer)
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
func (it *BaseTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTokenTransfer represents a Transfer event raised by the BaseToken contract.
type BaseTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseTokenTransferIterator{contract: _BaseToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaseTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTokenTransfer)
				if err := _BaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseToken *BaseTokenFilterer) ParseTransfer(log types.Log) (*BaseTokenTransfer, error) {
	event := new(BaseTokenTransfer)
	if err := _BaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BountiableTokenABI is the input ABI used to generate the binding from.
const BountiableTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"BountyCashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"bountyNonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBounty[]\",\"name\":\"bounties\",\"type\":\"tuple[]\"}],\"name\":\"cashBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BountiableToken is an auto generated Go binding around an Core contract.
type BountiableToken struct {
	BountiableTokenCaller     // Read-only binding to the contract
	BountiableTokenTransactor // Write-only binding to the contract
	BountiableTokenFilterer   // Log filterer for contract events
}

// BountiableTokenCaller is an auto generated read-only Go binding around an Core contract.
type BountiableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountiableTokenTransactor is an auto generated write-only Go binding around an Core contract.
type BountiableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountiableTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type BountiableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountiableTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type BountiableTokenSession struct {
	Contract     *BountiableToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BountiableTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type BountiableTokenCallerSession struct {
	Contract *BountiableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BountiableTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type BountiableTokenTransactorSession struct {
	Contract     *BountiableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BountiableTokenRaw is an auto generated low-level Go binding around an Core contract.
type BountiableTokenRaw struct {
	Contract *BountiableToken // Generic contract binding to access the raw methods on
}

// BountiableTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type BountiableTokenCallerRaw struct {
	Contract *BountiableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BountiableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type BountiableTokenTransactorRaw struct {
	Contract *BountiableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBountiableToken creates a new instance of BountiableToken, bound to a specific deployed contract.
func NewBountiableToken(address common.Address, backend bind.ContractBackend) (*BountiableToken, error) {
	contract, err := bindBountiableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BountiableToken{BountiableTokenCaller: BountiableTokenCaller{contract: contract}, BountiableTokenTransactor: BountiableTokenTransactor{contract: contract}, BountiableTokenFilterer: BountiableTokenFilterer{contract: contract}}, nil
}

// NewBountiableTokenCaller creates a new read-only instance of BountiableToken, bound to a specific deployed contract.
func NewBountiableTokenCaller(address common.Address, caller bind.ContractCaller) (*BountiableTokenCaller, error) {
	contract, err := bindBountiableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenCaller{contract: contract}, nil
}

// NewBountiableTokenTransactor creates a new write-only instance of BountiableToken, bound to a specific deployed contract.
func NewBountiableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BountiableTokenTransactor, error) {
	contract, err := bindBountiableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenTransactor{contract: contract}, nil
}

// NewBountiableTokenFilterer creates a new log filterer instance of BountiableToken, bound to a specific deployed contract.
func NewBountiableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BountiableTokenFilterer, error) {
	contract, err := bindBountiableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenFilterer{contract: contract}, nil
}

// bindBountiableToken binds a generic wrapper to an already deployed contract.
func bindBountiableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BountiableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountiableToken *BountiableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BountiableToken.Contract.BountiableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountiableToken *BountiableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountiableToken.Contract.BountiableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountiableToken *BountiableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountiableToken.Contract.BountiableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountiableToken *BountiableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BountiableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountiableToken *BountiableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountiableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountiableToken *BountiableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountiableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BountiableToken *BountiableTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BountiableToken *BountiableTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.Allowance(&_BountiableToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BountiableToken *BountiableTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.Allowance(&_BountiableToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.BalanceOf(&_BountiableToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.BalanceOf(&_BountiableToken.CallOpts, account)
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenCaller) BountyNonceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "bountyNonceOf", account)
	return *ret0, err
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenSession) BountyNonceOf(account common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.BountyNonceOf(&_BountiableToken.CallOpts, account)
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_BountiableToken *BountiableTokenCallerSession) BountyNonceOf(account common.Address) (*big.Int, error) {
	return _BountiableToken.Contract.BountyNonceOf(&_BountiableToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BountiableToken *BountiableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BountiableToken *BountiableTokenSession) Decimals() (uint8, error) {
	return _BountiableToken.Contract.Decimals(&_BountiableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_BountiableToken *BountiableTokenCallerSession) Decimals() (uint8, error) {
	return _BountiableToken.Contract.Decimals(&_BountiableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BountiableToken *BountiableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BountiableToken *BountiableTokenSession) Name() (string, error) {
	return _BountiableToken.Contract.Name(&_BountiableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_BountiableToken *BountiableTokenCallerSession) Name() (string, error) {
	return _BountiableToken.Contract.Name(&_BountiableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BountiableToken *BountiableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BountiableToken *BountiableTokenSession) Symbol() (string, error) {
	return _BountiableToken.Contract.Symbol(&_BountiableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_BountiableToken *BountiableTokenCallerSession) Symbol() (string, error) {
	return _BountiableToken.Contract.Symbol(&_BountiableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BountiableToken *BountiableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountiableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BountiableToken *BountiableTokenSession) TotalSupply() (*big.Int, error) {
	return _BountiableToken.Contract.TotalSupply(&_BountiableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BountiableToken *BountiableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BountiableToken.Contract.TotalSupply(&_BountiableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.Approve(&_BountiableToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.Approve(&_BountiableToken.TransactOpts, spender, amount)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_BountiableToken *BountiableTokenTransactor) CashBounty(opts *bind.TransactOpts, bounties []Bounty) (*types.Transaction, error) {
	return _BountiableToken.contract.Transact(opts, "cashBounty", bounties)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_BountiableToken *BountiableTokenSession) CashBounty(bounties []Bounty) (*types.Transaction, error) {
	return _BountiableToken.Contract.CashBounty(&_BountiableToken.TransactOpts, bounties)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_BountiableToken *BountiableTokenTransactorSession) CashBounty(bounties []Bounty) (*types.Transaction, error) {
	return _BountiableToken.Contract.CashBounty(&_BountiableToken.TransactOpts, bounties)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.Transfer(&_BountiableToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.Transfer(&_BountiableToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.TransferFrom(&_BountiableToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BountiableToken *BountiableTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BountiableToken.Contract.TransferFrom(&_BountiableToken.TransactOpts, sender, recipient, amount)
}

// BountiableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BountiableToken contract.
type BountiableTokenApprovalIterator struct {
	Event *BountiableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BountiableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountiableTokenApproval)
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
		it.Event = new(BountiableTokenApproval)
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
func (it *BountiableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountiableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountiableTokenApproval represents a Approval event raised by the BountiableToken contract.
type BountiableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BountiableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BountiableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenApprovalIterator{contract: _BountiableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BountiableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BountiableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountiableTokenApproval)
				if err := _BountiableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) ParseApproval(log types.Log) (*BountiableTokenApproval, error) {
	event := new(BountiableTokenApproval)
	if err := _BountiableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BountiableTokenBountyCashedIterator is returned from FilterBountyCashed and is used to iterate over the raw logs and unpacked data for BountyCashed events raised by the BountiableToken contract.
type BountiableTokenBountyCashedIterator struct {
	Event *BountiableTokenBountyCashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BountiableTokenBountyCashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountiableTokenBountyCashed)
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
		it.Event = new(BountiableTokenBountyCashed)
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
func (it *BountiableTokenBountyCashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountiableTokenBountyCashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountiableTokenBountyCashed represents a BountyCashed event raised by the BountiableToken contract.
type BountiableTokenBountyCashed struct {
	Owner       common.Address
	Target      common.Address
	Data        []byte
	Reward      *big.Int
	Nonce       *big.Int
	Deadline    *big.Int
	EnergyLimit *big.Int
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBountyCashed is a free log retrieval operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_BountiableToken *BountiableTokenFilterer) FilterBountyCashed(opts *bind.FilterOpts, owner []common.Address, target []common.Address) (*BountiableTokenBountyCashedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BountiableToken.contract.FilterLogs(opts, "BountyCashed", ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenBountyCashedIterator{contract: _BountiableToken.contract, event: "BountyCashed", logs: logs, sub: sub}, nil
}

// WatchBountyCashed is a free log subscription operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_BountiableToken *BountiableTokenFilterer) WatchBountyCashed(opts *bind.WatchOpts, sink chan<- *BountiableTokenBountyCashed, owner []common.Address, target []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BountiableToken.contract.WatchLogs(opts, "BountyCashed", ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountiableTokenBountyCashed)
				if err := _BountiableToken.contract.UnpackLog(event, "BountyCashed", log); err != nil {
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

// ParseBountyCashed is a log parse operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_BountiableToken *BountiableTokenFilterer) ParseBountyCashed(log types.Log) (*BountiableTokenBountyCashed, error) {
	event := new(BountiableTokenBountyCashed)
	if err := _BountiableToken.contract.UnpackLog(event, "BountyCashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BountiableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BountiableToken contract.
type BountiableTokenBurnIterator struct {
	Event *BountiableTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BountiableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountiableTokenBurn)
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
		it.Event = new(BountiableTokenBurn)
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
func (it *BountiableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountiableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountiableTokenBurn represents a Burn event raised by the BountiableToken contract.
type BountiableTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*BountiableTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BountiableToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenBurnIterator{contract: _BountiableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BountiableTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BountiableToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountiableTokenBurn)
				if err := _BountiableToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) ParseBurn(log types.Log) (*BountiableTokenBurn, error) {
	event := new(BountiableTokenBurn)
	if err := _BountiableToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BountiableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BountiableToken contract.
type BountiableTokenMintIterator struct {
	Event *BountiableTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BountiableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountiableTokenMint)
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
		it.Event = new(BountiableTokenMint)
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
func (it *BountiableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountiableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountiableTokenMint represents a Mint event raised by the BountiableToken contract.
type BountiableTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BountiableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BountiableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenMintIterator{contract: _BountiableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BountiableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BountiableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountiableTokenMint)
				if err := _BountiableToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) ParseMint(log types.Log) (*BountiableTokenMint, error) {
	event := new(BountiableTokenMint)
	if err := _BountiableToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BountiableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BountiableToken contract.
type BountiableTokenTransferIterator struct {
	Event *BountiableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BountiableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountiableTokenTransfer)
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
		it.Event = new(BountiableTokenTransfer)
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
func (it *BountiableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountiableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountiableTokenTransfer represents a Transfer event raised by the BountiableToken contract.
type BountiableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BountiableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BountiableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BountiableTokenTransferIterator{contract: _BountiableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BountiableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BountiableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountiableTokenTransfer)
				if err := _BountiableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BountiableToken *BountiableTokenFilterer) ParseTransfer(log types.Log) (*BountiableTokenTransfer, error) {
	event := new(BountiableTokenTransfer)
	if err := _BountiableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CallerABI is the input ABI used to generate the binding from.
const CallerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"energyLimit\",\"type\":\"uint256\"}],\"name\":\"bountyCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CallerBin is the compiled bytecode used for deploying new contracts.
var CallerBin = "0x608060405234801561001057600080fd5b506040516105ae3803806105ae83398181016040528101906100329190610091565b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555050610105565b60008151905061008b816100ee565b92915050565b6000602082840312156100a357600080fd5b60006100b18482850161007c565b91505092915050565b60006100c5826100cc565b9050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f7816100ba565b811461010257600080fd5b50565b61049a806101146000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806334c198a714610030575b600080fd5b61004a600480360381019061004591906101cf565b610060565b604051610057919061028d565b60405180910390f35b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff16146100c157600080fd5b60008475ffffffffffffffffffffffffffffffffffffffffffff1683856040516100eb9190610276565b60006040518083038160008787f1925050503d8060008114610129576040519150601f19603f3d011682016040523d82523d6000602084013e61012e565b606091505b50509050809150509392505050565b600061015061014b846102cd565b6102a8565b90508281526020810184848401111561016857600080fd5b61017384828561035e565b509392505050565b60008135905061018a81610411565b92915050565b600082601f8301126101a157600080fd5b81356101b184826020860161013d565b91505092915050565b6000813590506101c981610428565b92915050565b6000806000606084860312156101e457600080fd5b60006101f28682870161017b565b935050602084013567ffffffffffffffff81111561020f57600080fd5b61021b86828701610190565b925050604061022c868287016101ba565b9150509250925092565b61023f81610326565b82525050565b6000610250826102fe565b61025a8185610309565b935061026a81856020860161036d565b80840191505092915050565b60006102828284610245565b915081905092915050565b60006020820190506102a26000830184610236565b92915050565b60006102b26102c3565b90506102be82826103a0565b919050565b6000604051905090565b600067ffffffffffffffff8211156102e8576102e76103d1565b5b6102f182610400565b9050602081019050919050565b600081519050919050565b600081905092915050565b600061031f82610332565b9050919050565b60008115159050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561038b578082015181840152602081019050610370565b8381111561039a576000848401525b50505050565b6103a982610400565b810181811067ffffffffffffffff821117156103c8576103c76103d1565b5b80604052505050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b61041a81610314565b811461042557600080fd5b50565b61043181610354565b811461043c57600080fd5b5056fea264697066735822122001197287c47c80ea6a122ccf5c7c40dd82323e1b6de88a9d8fb221c68714aed764736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058"

// DeployCaller deploys a new Core contract, binding an instance of Caller to it.
func DeployCaller(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress_ common.Address) (common.Address, *types.Transaction, *Caller, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallerBin), backend, tokenAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// Caller is an auto generated Go binding around an Core contract.
type Caller struct {
	CallerCaller     // Read-only binding to the contract
	CallerTransactor // Write-only binding to the contract
	CallerFilterer   // Log filterer for contract events
}

// CallerCaller is an auto generated read-only Go binding around an Core contract.
type CallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerTransactor is an auto generated write-only Go binding around an Core contract.
type CallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFilterer is an auto generated log filtering Go binding around an Core contract events.
type CallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type CallerSession struct {
	Contract     *Caller           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type CallerCallerSession struct {
	Contract *CallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CallerTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type CallerTransactorSession struct {
	Contract     *CallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerRaw is an auto generated low-level Go binding around an Core contract.
type CallerRaw struct {
	Contract *Caller // Generic contract binding to access the raw methods on
}

// CallerCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type CallerCallerRaw struct {
	Contract *CallerCaller // Generic read-only contract binding to access the raw methods on
}

// CallerTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type CallerTransactorRaw struct {
	Contract *CallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaller creates a new instance of Caller, bound to a specific deployed contract.
func NewCaller(address common.Address, backend bind.ContractBackend) (*Caller, error) {
	contract, err := bindCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// NewCallerCaller creates a new read-only instance of Caller, bound to a specific deployed contract.
func NewCallerCaller(address common.Address, caller bind.ContractCaller) (*CallerCaller, error) {
	contract, err := bindCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerCaller{contract: contract}, nil
}

// NewCallerTransactor creates a new write-only instance of Caller, bound to a specific deployed contract.
func NewCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerTransactor, error) {
	contract, err := bindCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerTransactor{contract: contract}, nil
}

// NewCallerFilterer creates a new log filterer instance of Caller, bound to a specific deployed contract.
func NewCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFilterer, error) {
	contract, err := bindCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFilterer{contract: contract}, nil
}

// bindCaller binds a generic wrapper to an already deployed contract.
func bindCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.CallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transact(opts, method, params...)
}

// BountyCall is a paid mutator transaction binding the contract method 0x34c198a7.
//
// Solidity: function bountyCall(address target, bytes data, uint256 energyLimit) returns(bool)
func (_Caller *CallerTransactor) BountyCall(opts *bind.TransactOpts, target common.Address, data []byte, energyLimit *big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "bountyCall", target, data, energyLimit)
}

// BountyCall is a paid mutator transaction binding the contract method 0x34c198a7.
//
// Solidity: function bountyCall(address target, bytes data, uint256 energyLimit) returns(bool)
func (_Caller *CallerSession) BountyCall(target common.Address, data []byte, energyLimit *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.BountyCall(&_Caller.TransactOpts, target, data, energyLimit)
}

// BountyCall is a paid mutator transaction binding the contract method 0x34c198a7.
//
// Solidity: function bountyCall(address target, bytes data, uint256 energyLimit) returns(bool)
func (_Caller *CallerTransactorSession) BountyCall(target common.Address, data []byte, energyLimit *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.BountyCall(&_Caller.TransactOpts, target, data, energyLimit)
}

// ChequableTokenABI is the input ABI used to generate the binding from.
const ChequableTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ChequeCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"}],\"name\":\"cashCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"chequeNonceOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"firstChequeNonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChequableToken is an auto generated Go binding around an Core contract.
type ChequableToken struct {
	ChequableTokenCaller     // Read-only binding to the contract
	ChequableTokenTransactor // Write-only binding to the contract
	ChequableTokenFilterer   // Log filterer for contract events
}

// ChequableTokenCaller is an auto generated read-only Go binding around an Core contract.
type ChequableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequableTokenTransactor is an auto generated write-only Go binding around an Core contract.
type ChequableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequableTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type ChequableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequableTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type ChequableTokenSession struct {
	Contract     *ChequableToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChequableTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type ChequableTokenCallerSession struct {
	Contract *ChequableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChequableTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type ChequableTokenTransactorSession struct {
	Contract     *ChequableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChequableTokenRaw is an auto generated low-level Go binding around an Core contract.
type ChequableTokenRaw struct {
	Contract *ChequableToken // Generic contract binding to access the raw methods on
}

// ChequableTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type ChequableTokenCallerRaw struct {
	Contract *ChequableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ChequableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type ChequableTokenTransactorRaw struct {
	Contract *ChequableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChequableToken creates a new instance of ChequableToken, bound to a specific deployed contract.
func NewChequableToken(address common.Address, backend bind.ContractBackend) (*ChequableToken, error) {
	contract, err := bindChequableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChequableToken{ChequableTokenCaller: ChequableTokenCaller{contract: contract}, ChequableTokenTransactor: ChequableTokenTransactor{contract: contract}, ChequableTokenFilterer: ChequableTokenFilterer{contract: contract}}, nil
}

// NewChequableTokenCaller creates a new read-only instance of ChequableToken, bound to a specific deployed contract.
func NewChequableTokenCaller(address common.Address, caller bind.ContractCaller) (*ChequableTokenCaller, error) {
	contract, err := bindChequableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenCaller{contract: contract}, nil
}

// NewChequableTokenTransactor creates a new write-only instance of ChequableToken, bound to a specific deployed contract.
func NewChequableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ChequableTokenTransactor, error) {
	contract, err := bindChequableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenTransactor{contract: contract}, nil
}

// NewChequableTokenFilterer creates a new log filterer instance of ChequableToken, bound to a specific deployed contract.
func NewChequableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ChequableTokenFilterer, error) {
	contract, err := bindChequableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenFilterer{contract: contract}, nil
}

// bindChequableToken binds a generic wrapper to an already deployed contract.
func bindChequableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChequableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequableToken *ChequableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequableToken.Contract.ChequableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequableToken *ChequableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequableToken.Contract.ChequableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequableToken *ChequableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequableToken.Contract.ChequableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequableToken *ChequableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequableToken *ChequableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequableToken *ChequableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChequableToken *ChequableTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChequableToken *ChequableTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.Allowance(&_ChequableToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChequableToken *ChequableTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.Allowance(&_ChequableToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.BalanceOf(&_ChequableToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.BalanceOf(&_ChequableToken.CallOpts, account)
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_ChequableToken *ChequableTokenCaller) ChequeNonceOf(opts *bind.CallOpts, account common.Address, nonce *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "chequeNonceOf", account, nonce)
	return *ret0, err
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_ChequableToken *ChequableTokenSession) ChequeNonceOf(account common.Address, nonce *big.Int) (bool, error) {
	return _ChequableToken.Contract.ChequeNonceOf(&_ChequableToken.CallOpts, account, nonce)
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_ChequableToken *ChequableTokenCallerSession) ChequeNonceOf(account common.Address, nonce *big.Int) (bool, error) {
	return _ChequableToken.Contract.ChequeNonceOf(&_ChequableToken.CallOpts, account, nonce)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_ChequableToken *ChequableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_ChequableToken *ChequableTokenSession) Decimals() (uint8, error) {
	return _ChequableToken.Contract.Decimals(&_ChequableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_ChequableToken *ChequableTokenCallerSession) Decimals() (uint8, error) {
	return _ChequableToken.Contract.Decimals(&_ChequableToken.CallOpts)
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenCaller) FirstChequeNonceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "firstChequeNonceOf", account)
	return *ret0, err
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenSession) FirstChequeNonceOf(account common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.FirstChequeNonceOf(&_ChequableToken.CallOpts, account)
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_ChequableToken *ChequableTokenCallerSession) FirstChequeNonceOf(account common.Address) (*big.Int, error) {
	return _ChequableToken.Contract.FirstChequeNonceOf(&_ChequableToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_ChequableToken *ChequableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_ChequableToken *ChequableTokenSession) Name() (string, error) {
	return _ChequableToken.Contract.Name(&_ChequableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_ChequableToken *ChequableTokenCallerSession) Name() (string, error) {
	return _ChequableToken.Contract.Name(&_ChequableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_ChequableToken *ChequableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_ChequableToken *ChequableTokenSession) Symbol() (string, error) {
	return _ChequableToken.Contract.Symbol(&_ChequableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_ChequableToken *ChequableTokenCallerSession) Symbol() (string, error) {
	return _ChequableToken.Contract.Symbol(&_ChequableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChequableToken *ChequableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChequableToken *ChequableTokenSession) TotalSupply() (*big.Int, error) {
	return _ChequableToken.Contract.TotalSupply(&_ChequableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChequableToken *ChequableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ChequableToken.Contract.TotalSupply(&_ChequableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.Approve(&_ChequableToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.Approve(&_ChequableToken.TransactOpts, spender, amount)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_ChequableToken *ChequableTokenTransactor) CashCheque(opts *bind.TransactOpts, cheque Cheque) (*types.Transaction, error) {
	return _ChequableToken.contract.Transact(opts, "cashCheque", cheque)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_ChequableToken *ChequableTokenSession) CashCheque(cheque Cheque) (*types.Transaction, error) {
	return _ChequableToken.Contract.CashCheque(&_ChequableToken.TransactOpts, cheque)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_ChequableToken *ChequableTokenTransactorSession) CashCheque(cheque Cheque) (*types.Transaction, error) {
	return _ChequableToken.Contract.CashCheque(&_ChequableToken.TransactOpts, cheque)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.Transfer(&_ChequableToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.Transfer(&_ChequableToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.TransferFrom(&_ChequableToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChequableToken *ChequableTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChequableToken.Contract.TransferFrom(&_ChequableToken.TransactOpts, sender, recipient, amount)
}

// ChequableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ChequableToken contract.
type ChequableTokenApprovalIterator struct {
	Event *ChequableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChequableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequableTokenApproval)
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
		it.Event = new(ChequableTokenApproval)
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
func (it *ChequableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequableTokenApproval represents a Approval event raised by the ChequableToken contract.
type ChequableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChequableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChequableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenApprovalIterator{contract: _ChequableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChequableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChequableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequableTokenApproval)
				if err := _ChequableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) ParseApproval(log types.Log) (*ChequableTokenApproval, error) {
	event := new(ChequableTokenApproval)
	if err := _ChequableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ChequableToken contract.
type ChequableTokenBurnIterator struct {
	Event *ChequableTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChequableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequableTokenBurn)
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
		it.Event = new(ChequableTokenBurn)
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
func (it *ChequableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequableTokenBurn represents a Burn event raised by the ChequableToken contract.
type ChequableTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*ChequableTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ChequableToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenBurnIterator{contract: _ChequableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ChequableTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ChequableToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequableTokenBurn)
				if err := _ChequableToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) ParseBurn(log types.Log) (*ChequableTokenBurn, error) {
	event := new(ChequableTokenBurn)
	if err := _ChequableToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequableTokenChequeCashIterator is returned from FilterChequeCash and is used to iterate over the raw logs and unpacked data for ChequeCash events raised by the ChequableToken contract.
type ChequableTokenChequeCashIterator struct {
	Event *ChequableTokenChequeCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChequableTokenChequeCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequableTokenChequeCash)
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
		it.Event = new(ChequableTokenChequeCash)
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
func (it *ChequableTokenChequeCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequableTokenChequeCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequableTokenChequeCash represents a ChequeCash event raised by the ChequableToken contract.
type ChequableTokenChequeCash struct {
	Owner    common.Address
	Spender  common.Address
	Amount   *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChequeCash is a free log retrieval operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_ChequableToken *ChequableTokenFilterer) FilterChequeCash(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChequableTokenChequeCashIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChequableToken.contract.FilterLogs(opts, "ChequeCash", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenChequeCashIterator{contract: _ChequableToken.contract, event: "ChequeCash", logs: logs, sub: sub}, nil
}

// WatchChequeCash is a free log subscription operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_ChequableToken *ChequableTokenFilterer) WatchChequeCash(opts *bind.WatchOpts, sink chan<- *ChequableTokenChequeCash, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChequableToken.contract.WatchLogs(opts, "ChequeCash", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequableTokenChequeCash)
				if err := _ChequableToken.contract.UnpackLog(event, "ChequeCash", log); err != nil {
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

// ParseChequeCash is a log parse operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_ChequableToken *ChequableTokenFilterer) ParseChequeCash(log types.Log) (*ChequableTokenChequeCash, error) {
	event := new(ChequableTokenChequeCash)
	if err := _ChequableToken.contract.UnpackLog(event, "ChequeCash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ChequableToken contract.
type ChequableTokenMintIterator struct {
	Event *ChequableTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChequableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequableTokenMint)
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
		it.Event = new(ChequableTokenMint)
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
func (it *ChequableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequableTokenMint represents a Mint event raised by the ChequableToken contract.
type ChequableTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*ChequableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChequableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenMintIterator{contract: _ChequableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ChequableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChequableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequableTokenMint)
				if err := _ChequableToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) ParseMint(log types.Log) (*ChequableTokenMint, error) {
	event := new(ChequableTokenMint)
	if err := _ChequableToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ChequableToken contract.
type ChequableTokenTransferIterator struct {
	Event *ChequableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChequableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequableTokenTransfer)
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
		it.Event = new(ChequableTokenTransfer)
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
func (it *ChequableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequableTokenTransfer represents a Transfer event raised by the ChequableToken contract.
type ChequableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChequableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChequableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChequableTokenTransferIterator{contract: _ChequableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChequableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChequableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequableTokenTransfer)
				if err := _ChequableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChequableToken *ChequableTokenFilterer) ParseTransfer(log types.Log) (*ChequableTokenTransfer, error) {
	event := new(ChequableTokenTransfer)
	if err := _ChequableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenABI is the input ABI used to generate the binding from.
const CoreTokenABI = "[{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"contractPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"BountyCashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ChequeCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"bountyNonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBounty[]\",\"name\":\"bounties\",\"type\":\"tuple[]\"}],\"name\":\"cashBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"}],\"name\":\"cashCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"chequeNonceOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"equivalentValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"firstChequeNonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFromEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CoreTokenBin is the compiled bytecode used for deploying new contracts.
var CoreTokenBin = "0x6101c06040523480156200001257600080fd5b5060405162003cb138038062003cb1833981810160405281019062000038919062000451565b80826040518060400160405280600981526020017f436f7265546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f434e54000000000000000000000000000000000000000000000000000000000081525060126040518060400160405280600981526020017f436f7265546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250600060405180608001604052806054815260200162003bfa6054913980519060200120905080608081815250506000838051906020012090508060a081815250506000838051906020012090508060c081815250504660e08181525050620001848383836200032960201b60201c565b610100818152505050505050508260009080519060200190620001a992919062000365565b508160019080519060200190620001c292919062000365565b508060ff166101208160ff1660f81b815250505050508075ffffffffffffffffffffffffffffffffffffffffffff166101408175ffffffffffffffffffffffffffffffffffffffffffff1660501b81525050508075ffffffffffffffffffffffffffffffffffffffffffff166101608175ffffffffffffffffffffffffffffffffffffffffffff1660501b815250505060405180608001604052806045815260200162003bb5604591398051906020012061018081815250506040518060a001604052806063815260200162003c4e60639139805190602001206101a0818152505030604051620002b390620003f6565b620002bf9190620004c5565b604051809103906000f080158015620002dc573d6000803e3d6000fd5b50600760006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555050506200064a565b6000838383463060405160200162000346959493929190620004e2565b6040516020818303038152906040528051906020012090509392505050565b8280546200037390620005b1565b90600052602060002090601f016020900481019282620003975760008555620003e3565b82601f10620003b257805160ff1916838001178555620003e3565b82800160010185558215620003e3579182015b82811115620003e2578251825591602001919060010190620003c5565b5b509050620003f2919062000404565b5090565b6105ae806200360783390190565b5b808211156200041f57600081600090555060010162000405565b5090565b600081519050620004348162000616565b92915050565b6000815190506200044b8162000630565b92915050565b600080604083850312156200046557600080fd5b6000620004758582860162000423565b925050602062000488858286016200043a565b9150509250929050565b6200049d816200053f565b82525050565b620004ae8162000553565b82525050565b620004bf81620005a7565b82525050565b6000602082019050620004dc600083018462000492565b92915050565b600060a082019050620004f96000830188620004a3565b620005086020830187620004a3565b620005176040830186620004a3565b620005266060830185620004b4565b62000535608083018462000492565b9695505050505050565b60006200054c8262000585565b9050919050565b6000819050919050565b60006200056a826200053f565b9050919050565b60006200057e826200053f565b9050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006002820490506001821680620005ca57607f821691505b60208210811415620005e157620005e0620005e7565b5b50919050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b62000621816200055d565b81146200062d57600080fd5b50565b6200063b8162000571565b81146200064757600080fd5b50565b60805160a05160c05160e051610100516101205160f81c6101405160501c6101605160501c610180516101a051612f23620006e4600039600061188f01526000611c8e015260008181610b8b0152611603015260008181610ad801528181610ffa015261108e01526000610fd2015260006119d40152600061193e015260006119a901526000611988015260006119670152612f236000f3fe6080604052600436106101355760003560e01c80634b40e901116100ab57806370c808411161006f57806370c808411461047657806389b9ce32146104b35780638b782835146104f0578063a613914d1461050c578063bcc28adb14610549578063c64a4ca01461057257610135565b80634b40e9011461037d5780634cb53686146103ba578063542f48f3146103e35780635d1fb5f914610420578063659954791461044b57610135565b80632227fd8b116100fd5780632227fd8b14610247578063231782d81461028457806331f2e679146102af578063380544b3146102ec5780633f378b7614610329578063449207ec1461035257610135565b806307ba2a171461013a5780630bf3a456146101655780631832bdea146101a25780631d7976f3146101df5780631f1881f81461021c575b600080fd5b34801561014657600080fd5b5061014f61059d565b60405161015c9190612504565b60405180910390f35b34801561017157600080fd5b5061018c60048036038101906101879190611e16565b61062f565b6040516101999190612626565b60405180910390f35b3480156101ae57600080fd5b506101c960048036038101906101c49190611ea1565b6106be565b6040516101d69190612306565b60405180910390f35b3480156101eb57600080fd5b5061020660048036038101906102019190611ded565b610774565b6040516102139190612626565b60405180910390f35b34801561022857600080fd5b506102316107c1565b60405161023e9190612626565b60405180910390f35b34801561025357600080fd5b5061026e60048036038101906102699190611ded565b6107cb565b60405161027b9190612626565b60405180910390f35b34801561029057600080fd5b50610299610920565b6040516102a69190612504565b60405180910390f35b3480156102bb57600080fd5b506102d660048036038101906102d19190611e52565b6109b2565b6040516102e39190612306565b60405180910390f35b3480156102f857600080fd5b50610313600480360381019061030e9190611e52565b610aa9565b6040516103209190612306565b60405180910390f35b34801561033557600080fd5b50610350600480360381019061034b9190611f8c565b610acc565b005b34801561035e57600080fd5b50610367610b89565b60405161037491906124e9565b60405180910390f35b34801561038957600080fd5b506103a4600480360381019061039f9190611ea1565b610bad565b6040516103b19190612306565b60405180910390f35b3480156103c657600080fd5b506103e160048036038101906103dc9190611edd565b610bc4565b005b3480156103ef57600080fd5b5061040a60048036038101906104059190611ea1565b610fad565b6040516104179190612306565b60405180910390f35b34801561042c57600080fd5b50610435610fce565b6040516104429190612678565b60405180910390f35b34801561045757600080fd5b50610460610ff6565b60405161046d91906124ce565b60405180910390f35b34801561048257600080fd5b5061049d60048036038101906104989190611ded565b61101e565b6040516104aa9190612626565b60405180910390f35b3480156104bf57600080fd5b506104da60048036038101906104d59190611ea1565b61106b565b6040516104e79190612306565b60405180910390f35b61050a60048036038101906105059190611f8c565b61108c565b005b34801561051857600080fd5b50610533600480360381019061052e9190611ea1565b61114b565b6040516105409190612306565b60405180910390f35b34801561055557600080fd5b50610570600480360381019061056b9190611f4b565b611162565b005b34801561057e57600080fd5b50610587611356565b6040516105949190612626565b60405180910390f35b6060600080546105ac90612a90565b80601f01602080910402602001604051908101604052809291908181526020018280546105d890612a90565b80156106255780601f106105fa57610100808354040283529160200191610625565b820191906000526020600020905b81548152906001019060200180831161060857829003601f168201915b5050505050905090565b6000600460008475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600080610100836106cf91906127a8565b90506000610100846106e19190612b15565b905060008160026106f2919061282c565b90506000600560008875ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002054905060006002838361075d91906127a8565b6107679190612b15565b1494505050505092915050565b6000600360008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600254905090565b600080600090505b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600560008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054141561086257808061085a90612ac2565b9150506107d3565b6000600560008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905060005b600160028260026108cf919061282c565b846108da91906127a8565b6108e49190612b15565b14156108fd5780806108f590612ac2565b9150506108be565b806101008461090c919061294a565b6109169190612752565b9350505050919050565b60606001805461092f90612a90565b80601f016020809104026020016040519081016040528092919081815260200182805461095b90612a90565b80156109a85780601f1061097d576101008083540402835291602001916109a8565b820191906000526020600020905b81548152906001019060200180831161098b57829003601f168201915b5050505050905090565b600080600460008675ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a7f576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610a7690612526565b60405180910390fd5b60008382039050610a91868686611367565b610a9c863383611505565b6001925050509392505050565b600080610ab5836115fc565b9050610ac28585836109b2565b9150509392505050565b610ad63382611704565b7f000000000000000000000000000000000000000000000000000000000000000075ffffffffffffffffffffffffffffffffffffffffffff16634b40e90133836040518363ffffffff1660e01b8152600401610b339291906122dd565b602060405180830381600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b859190611f22565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000610bba338484611367565b6001905092915050565b60005b82829050811015610fa85736838383818110610c0c577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002810190610c1e91906126ea565b90506000816000016020810190610c359190611ded565b905060008260600135905080610c4a8361183e565b1015610c5857505050610f95565b600083608001359050600660008475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548114610cb45750505050610f95565b60008460a00135905042811015610ccf575050505050610f95565b6000856020016020810190610ce49190611ded565b90506000868060400190610cf89190612693565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060008760c0013590506000610d5684848989898761188b565b9050610db4818a8060e00190610d6c9190612693565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506118f3565b75ffffffffffffffffffffffffffffffffffffffffffff168875ffffffffffffffffffffffffffffffffffffffffffff1614610df857505050505050505050610f95565b6001600660008a75ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e4c9190612752565b92505081905550610e5e883389611367565b6000600760009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff166334c198a78686866040518463ffffffff1660e01b8152600401610ec39392919061229f565b602060405180830381600087803b158015610edd57600080fd5b505af1158015610ef1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f159190611f22565b90508475ffffffffffffffffffffffffffffffffffffffffffff168975ffffffffffffffffffffffffffffffffffffffffffff167f43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98868b8b8b8988604051610f8296959493929190612466565b60405180910390a3505050505050505050505b8080610fa090612ac2565b915050610bc7565b505050565b600080610fb9836115fc565b9050610fc5848261114b565b91505092915050565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b6000600660008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600080611077836115fc565b90506110838482610bad565b91505092915050565b7f000000000000000000000000000000000000000000000000000000000000000075ffffffffffffffffffffffffffffffffffffffffffff166331f2e6793330846040518463ffffffff1660e01b81526004016110eb93929190612268565b602060405180830381600087803b15801561110557600080fd5b505af1158015611119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113d9190611f22565b506111483382611a78565b50565b6000611158338484611505565b6001905092915050565b60008160000160208101906111779190611ded565b9050600082602001602081019061118e9190611ded565b9050600083604001359050600084606001359050600085608001359050428110156111ee576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016111e590612606565b60405180910390fd5b6111f88583611b3f565b600061120685858585611c8a565b905061126481888060a0019061121c9190612693565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506118f3565b75ffffffffffffffffffffffffffffffffffffffffffff168675ffffffffffffffffffffffffffffffffffffffffffff16146112d5576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016112cc906125c6565b60405180910390fd5b6112e0868686611505565b8475ffffffffffffffffffffffffffffffffffffffffffff168675ffffffffffffffffffffffffffffffffffffffffffff167f313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b086868660405161134593929190612641565b60405180910390a350505050505050565b600061136260016115fc565b905090565b6000600360008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156113f2576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016113e9906125e6565b60405180910390fd5b818103600360008675ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600360008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461148f9190612752565b925050819055508275ffffffffffffffffffffffffffffffffffffffffffff168475ffffffffffffffffffffffffffffffffffffffffffff167fc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1846040516114f79190612626565b60405180910390a350505050565b80600460008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508175ffffffffffffffffffffffffffffffffffffffffffff168375ffffffffffffffffffffffffffffffffffffffffffff167fafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d836040516115ef9190612626565b60405180910390a3505050565b60008060007f000000000000000000000000000000000000000000000000000000000000000075ffffffffffffffffffffffffffffffffffffffffffff1663d9c1c1c96040518163ffffffff1660e01b815260040160606040518083038186803b15801561166957600080fd5b505afa15801561167d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a19190611fb5565b63ffffffff169250925050600081116116ef576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016116e690612546565b60405180910390fd5b83826116fb919061294a565b92505050919050565b6000600360008475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561178f576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611786906125a6565b60405180910390fd5b818103600360008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816002600082825403925050819055508275ffffffffffffffffffffffffffffffffffffffffffff167fe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e836040516118319190612626565b60405180910390a2505050565b6000600360008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60007f0000000000000000000000000000000000000000000000000000000000000000878780519060200120878787876040516020016118d19796959493929190612321565b6040516020818303038152906040528051906020012090509695505050505050565b600060ab825114611939576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161193090612586565b60405180910390fd5b6000467f0000000000000000000000000000000000000000000000000000000000000000146119d2576119cd7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611ce5565b6119f4565b7f00000000000000000000000000000000000000000000000000000000000000005b905060008185604051602001611a0b929190612231565b6040516020818303038152906040528051906020012090506001818560405160008152602001604052604051611a42929190612436565b6020604051602081039080840390855afa158015611a64573d6000803e3d6000fd5b505050602060405103519250505092915050565b80600360008475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611acb9190612752565b925050819055508060026000828254611ae49190612752565b925050819055508175ffffffffffffffffffffffffffffffffffffffffffff167f058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb167882604051611b339190612626565b60405180910390a25050565b600061010082611b4f91906127a8565b9050600061010083611b619190612b15565b90506000816002611b72919061282c565b90506000600560008775ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020549050600060028383611bdd91906127a8565b611be79190612b15565b14611c27576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611c1e90612566565b60405180910390fd5b818117600560008875ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002081905550505050505050565b60007f000000000000000000000000000000000000000000000000000000000000000085858585604051602001611cc5959493929190612390565b604051602081830303815290604052805190602001209050949350505050565b60008383834630604051602001611d009594939291906123e3565b6040516020818303038152906040528051906020012090509392505050565b600081359050611d2e81612e6c565b92915050565b60008083601f840112611d4657600080fd5b8235905067ffffffffffffffff811115611d5f57600080fd5b602083019150836020820283011115611d7757600080fd5b9250929050565b600081519050611d8d81612e83565b92915050565b600060c08284031215611da557600080fd5b81905092915050565b600081359050611dbd81612e9a565b92915050565b600081519050611dd281612e9a565b92915050565b600081519050611de781612eb1565b92915050565b600060208284031215611dff57600080fd5b6000611e0d84828501611d1f565b91505092915050565b60008060408385031215611e2957600080fd5b6000611e3785828601611d1f565b9250506020611e4885828601611d1f565b9150509250929050565b600080600060608486031215611e6757600080fd5b6000611e7586828701611d1f565b9350506020611e8686828701611d1f565b9250506040611e9786828701611dae565b9150509250925092565b60008060408385031215611eb457600080fd5b6000611ec285828601611d1f565b9250506020611ed385828601611dae565b9150509250929050565b60008060208385031215611ef057600080fd5b600083013567ffffffffffffffff811115611f0a57600080fd5b611f1685828601611d34565b92509250509250929050565b600060208284031215611f3457600080fd5b6000611f4284828501611d7e565b91505092915050565b600060208284031215611f5d57600080fd5b600082013567ffffffffffffffff811115611f7757600080fd5b611f8384828501611d93565b91505092915050565b600060208284031215611f9e57600080fd5b6000611fac84828501611dae565b91505092915050565b600080600060608486031215611fca57600080fd5b6000611fd886828701611dc3565b9350506020611fe986828701611dc3565b9250506040611ffa86828701611dd8565b9150509250925092565b61200d816129a4565b82525050565b61201c816129b6565b82525050565b61202b816129c2565b82525050565b61204261203d826129c2565b612b0b565b82525050565b60006120538261270f565b61205d8185612725565b935061206d818560208601612a5d565b61207681612bd3565b840191505092915050565b61208a81612a15565b82525050565b61209981612a39565b82525050565b60006120aa8261271a565b6120b48185612736565b93506120c4818560208601612a5d565b6120cd81612bd3565b840191505092915050565b60006120e5602c83612736565b91506120f082612bf1565b604082019050919050565b6000612108602983612736565b915061211382612c40565b604082019050919050565b600061212b601d83612736565b915061213682612c8f565b602082019050919050565b600061214e602583612736565b915061215982612cb8565b604082019050919050565b6000612171602683612736565b915061217c82612d07565b604082019050919050565b6000612194602483612736565b915061219f82612d56565b604082019050919050565b60006121b7602a83612736565b91506121c282612da5565b604082019050919050565b60006121da602283612736565b91506121e582612df4565b604082019050919050565b60006121fd600283612747565b915061220882612e43565b600282019050919050565b61221c816129ee565b82525050565b61222b81612a08565b82525050565b600061223c826121f0565b91506122488285612031565b6020820191506122588284612031565b6020820191508190509392505050565b600060608201905061227d6000830186612004565b61228a6020830185612004565b6122976040830184612213565b949350505050565b60006060820190506122b46000830186612004565b81810360208301526122c68185612048565b90506122d56040830184612213565b949350505050565b60006040820190506122f26000830185612004565b6122ff6020830184612213565b9392505050565b600060208201905061231b6000830184612013565b92915050565b600060e082019050612336600083018a612022565b6123436020830189612004565b6123506040830188612022565b61235d6060830187612213565b61236a6080830186612213565b61237760a0830185612213565b61238460c0830184612213565b98975050505050505050565b600060a0820190506123a56000830188612022565b6123b26020830187612004565b6123bf6040830186612213565b6123cc6060830185612213565b6123d96080830184612213565b9695505050505050565b600060a0820190506123f86000830188612022565b6124056020830187612022565b6124126040830186612022565b61241f6060830185612213565b61242c6080830184612004565b9695505050505050565b600060408201905061244b6000830185612022565b818103602083015261245d8184612048565b90509392505050565b600060c08201905081810360008301526124808189612048565b905061248f6020830188612213565b61249c6040830187612213565b6124a96060830186612213565b6124b66080830185612213565b6124c360a0830184612013565b979650505050505050565b60006020820190506124e36000830184612081565b92915050565b60006020820190506124fe6000830184612090565b92915050565b6000602082019050818103600083015261251e818461209f565b905092915050565b6000602082019050818103600083015261253f816120d8565b9050919050565b6000602082019050818103600083015261255f816120fb565b9050919050565b6000602082019050818103600083015261257f8161211e565b9050919050565b6000602082019050818103600083015261259f81612141565b9050919050565b600060208201905081810360008301526125bf81612164565b9050919050565b600060208201905081810360008301526125df81612187565b9050919050565b600060208201905081810360008301526125ff816121aa565b9050919050565b6000602082019050818103600083015261261f816121cd565b9050919050565b600060208201905061263b6000830184612213565b92915050565b60006060820190506126566000830186612213565b6126636020830185612213565b6126706040830184612213565b949350505050565b600060208201905061268d6000830184612222565b92915050565b600080833560016020038436030381126126ac57600080fd5b80840192508235915067ffffffffffffffff8211156126ca57600080fd5b6020830192506001820236038313156126e257600080fd5b509250929050565b6000823560016101000383360303811261270357600080fd5b80830191505092915050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061275d826129ee565b9150612768836129ee565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561279d5761279c612b46565b5b828201905092915050565b60006127b3826129ee565b91506127be836129ee565b9250826127ce576127cd612b75565b5b828204905092915050565b6000808291508390505b6001851115612823578086048111156127ff576127fe612b46565b5b600185161561280e5780820291505b808102905061281c85612be4565b94506127e3565b94509492505050565b6000612837826129ee565b9150612842836129ee565b925061286f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612877565b905092915050565b6000826128875760019050612943565b816128955760009050612943565b81600181146128ab57600281146128b5576128e4565b6001915050612943565b60ff8411156128c7576128c6612b46565b5b8360020a9150848211156128de576128dd612b46565b5b50612943565b5060208310610133831016604e8410600b84101617156129195782820a90508381111561291457612913612b46565b5b612943565b61292684848460016127d9565b9250905081840481111561293d5761293c612b46565b5b81810290505b9392505050565b6000612955826129ee565b9150612960836129ee565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561299957612998612b46565b5b828202905092915050565b60006129af826129cc565b9050919050565b60008115159050919050565b6000819050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600060ff82169050919050565b6000612a2082612a27565b9050919050565b6000612a32826129cc565b9050919050565b6000612a4482612a4b565b9050919050565b6000612a56826129cc565b9050919050565b60005b83811015612a7b578082015181840152602081019050612a60565b83811115612a8a576000848401525b50505050565b60006002820490506001821680612aa857607f821691505b60208210811415612abc57612abb612ba4565b5b50919050565b6000612acd826129ee565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415612b0057612aff612b46565b5b600182019050919050565b6000819050919050565b6000612b20826129ee565b9150612b2b836129ee565b925082612b3b57612b3a612b75565b5b828206905092915050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b7f42617365546f6b656e3a207472616e7366657220616d6f756e7420657863656560008201527f647320616c6c6f77616e63650000000000000000000000000000000000000000602082015250565b7f4571756976616c656e74546f6b656e3a20756e6b6e6f776e206571756976616c60008201527f656e742076616c75650000000000000000000000000000000000000000000000602082015250565b7f436865717561626c65546f6b656e3a206e6f6e63652069732075736564000000600082015250565b7f4549503731323a207369676e6174757265206c656e67746820697320696e636f60008201527f7272656374000000000000000000000000000000000000000000000000000000602082015250565b7f42617365546f6b656e3a206275726e20616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f436865717561626c65546f6b656e3a207369676e617475726520697320696e7660008201527f616c696400000000000000000000000000000000000000000000000000000000602082015250565b7f42617365546f6b656e3a207472616e7366657220616d6f756e7420657863656560008201527f64732062616c616e636500000000000000000000000000000000000000000000602082015250565b7f436865717561626c65546f6b656e3a20646561646c696e65206973207061737360008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b612e75816129a4565b8114612e8057600080fd5b50565b612e8c816129b6565b8114612e9757600080fd5b50565b612ea3816129ee565b8114612eae57600080fd5b50565b612eba816129f8565b8114612ec557600080fd5b5056fea264697066735822122014f5c33b0e2d778bce3498ec3b7c069a530e822010cf02404f6332f1470e0e7464736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058608060405234801561001057600080fd5b506040516105ae3803806105ae83398181016040528101906100329190610091565b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555050610105565b60008151905061008b816100ee565b92915050565b6000602082840312156100a357600080fd5b60006100b18482850161007c565b91505092915050565b60006100c5826100cc565b9050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f7816100ba565b811461010257600080fd5b50565b61049a806101146000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806334c198a714610030575b600080fd5b61004a600480360381019061004591906101cf565b610060565b604051610057919061028d565b60405180910390f35b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff16146100c157600080fd5b60008475ffffffffffffffffffffffffffffffffffffffffffff1683856040516100eb9190610276565b60006040518083038160008787f1925050503d8060008114610129576040519150601f19603f3d011682016040523d82523d6000602084013e61012e565b606091505b50509050809150509392505050565b600061015061014b846102cd565b6102a8565b90508281526020810184848401111561016857600080fd5b61017384828561035e565b509392505050565b60008135905061018a81610411565b92915050565b600082601f8301126101a157600080fd5b81356101b184826020860161013d565b91505092915050565b6000813590506101c981610428565b92915050565b6000806000606084860312156101e457600080fd5b60006101f28682870161017b565b935050602084013567ffffffffffffffff81111561020f57600080fd5b61021b86828701610190565b925050604061022c868287016101ba565b9150509250925092565b61023f81610326565b82525050565b6000610250826102fe565b61025a8185610309565b935061026a81856020860161036d565b80840191505092915050565b60006102828284610245565b915081905092915050565b60006020820190506102a26000830184610236565b92915050565b60006102b26102c3565b90506102be82826103a0565b919050565b6000604051905090565b600067ffffffffffffffff8211156102e8576102e76103d1565b5b6102f182610400565b9050602081019050919050565b600081519050919050565b600081905092915050565b600061031f82610332565b9050919050565b60008115159050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561038b578082015181840152602081019050610370565b8381111561039a576000848401525b50505050565b6103a982610400565b810181811067ffffffffffffffff821117156103c8576103c76103d1565b5b80604052505050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b61041a81610314565b811461042557600080fd5b50565b61043181610354565b811461043c57600080fd5b5056fea264697066735822122001197287c47c80ea6a122ccf5c7c40dd82323e1b6de88a9d8fb221c68714aed764736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e373839613539656500584368657175652861646472657373207370656e6465722c75696e7432353620616d6f756e742c75696e74323536206e6f6e63652c75696e7432353620646561646c696e6529454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e74323536206e6574776f726b49642c6164647265737320766572696679696e67436f6e747261637429426f756e74792861646472657373207461726765742c627974657320646174612c75696e74323536207265776172642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c75696e7432353620656e657267794c696d697429"

// DeployCoreToken deploys a new Core contract, binding an instance of CoreToken to it.
func DeployCoreToken(auth *bind.TransactOpts, backend bind.ContractBackend, wrappedToken common.Address, priceFeed common.Address) (common.Address, *types.Transaction, *CoreToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CoreTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoreTokenBin), backend, wrappedToken, priceFeed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoreToken{CoreTokenCaller: CoreTokenCaller{contract: contract}, CoreTokenTransactor: CoreTokenTransactor{contract: contract}, CoreTokenFilterer: CoreTokenFilterer{contract: contract}}, nil
}

// CoreToken is an auto generated Go binding around an Core contract.
type CoreToken struct {
	CoreTokenCaller     // Read-only binding to the contract
	CoreTokenTransactor // Write-only binding to the contract
	CoreTokenFilterer   // Log filterer for contract events
}

// CoreTokenCaller is an auto generated read-only Go binding around an Core contract.
type CoreTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTokenTransactor is an auto generated write-only Go binding around an Core contract.
type CoreTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type CoreTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type CoreTokenSession struct {
	Contract     *CoreToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type CoreTokenCallerSession struct {
	Contract *CoreTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CoreTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type CoreTokenTransactorSession struct {
	Contract     *CoreTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CoreTokenRaw is an auto generated low-level Go binding around an Core contract.
type CoreTokenRaw struct {
	Contract *CoreToken // Generic contract binding to access the raw methods on
}

// CoreTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type CoreTokenCallerRaw struct {
	Contract *CoreTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type CoreTokenTransactorRaw struct {
	Contract *CoreTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoreToken creates a new instance of CoreToken, bound to a specific deployed contract.
func NewCoreToken(address common.Address, backend bind.ContractBackend) (*CoreToken, error) {
	contract, err := bindCoreToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoreToken{CoreTokenCaller: CoreTokenCaller{contract: contract}, CoreTokenTransactor: CoreTokenTransactor{contract: contract}, CoreTokenFilterer: CoreTokenFilterer{contract: contract}}, nil
}

// NewCoreTokenCaller creates a new read-only instance of CoreToken, bound to a specific deployed contract.
func NewCoreTokenCaller(address common.Address, caller bind.ContractCaller) (*CoreTokenCaller, error) {
	contract, err := bindCoreToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTokenCaller{contract: contract}, nil
}

// NewCoreTokenTransactor creates a new write-only instance of CoreToken, bound to a specific deployed contract.
func NewCoreTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTokenTransactor, error) {
	contract, err := bindCoreToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTokenTransactor{contract: contract}, nil
}

// NewCoreTokenFilterer creates a new log filterer instance of CoreToken, bound to a specific deployed contract.
func NewCoreTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreTokenFilterer, error) {
	contract, err := bindCoreToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreTokenFilterer{contract: contract}, nil
}

// bindCoreToken binds a generic wrapper to an already deployed contract.
func bindCoreToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoreTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreToken *CoreTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoreToken.Contract.CoreTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreToken *CoreTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreToken.Contract.CoreTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreToken *CoreTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreToken.Contract.CoreTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreToken *CoreTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoreToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreToken *CoreTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreToken *CoreTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CoreToken *CoreTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CoreToken *CoreTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CoreToken.Contract.Allowance(&_CoreToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CoreToken.Contract.Allowance(&_CoreToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.BalanceOf(&_CoreToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.BalanceOf(&_CoreToken.CallOpts, account)
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCaller) BountyNonceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "bountyNonceOf", account)
	return *ret0, err
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenSession) BountyNonceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.BountyNonceOf(&_CoreToken.CallOpts, account)
}

// BountyNonceOf is a free data retrieval call binding the contract method 0x70c80841.
//
// Solidity: function bountyNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) BountyNonceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.BountyNonceOf(&_CoreToken.CallOpts, account)
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_CoreToken *CoreTokenCaller) ChequeNonceOf(opts *bind.CallOpts, account common.Address, nonce *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "chequeNonceOf", account, nonce)
	return *ret0, err
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_CoreToken *CoreTokenSession) ChequeNonceOf(account common.Address, nonce *big.Int) (bool, error) {
	return _CoreToken.Contract.ChequeNonceOf(&_CoreToken.CallOpts, account, nonce)
}

// ChequeNonceOf is a free data retrieval call binding the contract method 0x1832bdea.
//
// Solidity: function chequeNonceOf(address account, uint256 nonce) view returns(bool)
func (_CoreToken *CoreTokenCallerSession) ChequeNonceOf(account common.Address, nonce *big.Int) (bool, error) {
	return _CoreToken.Contract.ChequeNonceOf(&_CoreToken.CallOpts, account, nonce)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_CoreToken *CoreTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_CoreToken *CoreTokenSession) Decimals() (uint8, error) {
	return _CoreToken.Contract.Decimals(&_CoreToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_CoreToken *CoreTokenCallerSession) Decimals() (uint8, error) {
	return _CoreToken.Contract.Decimals(&_CoreToken.CallOpts)
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_CoreToken *CoreTokenCaller) EquivalentValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "equivalentValue")
	return *ret0, err
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_CoreToken *CoreTokenSession) EquivalentValue() (*big.Int, error) {
	return _CoreToken.Contract.EquivalentValue(&_CoreToken.CallOpts)
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) EquivalentValue() (*big.Int, error) {
	return _CoreToken.Contract.EquivalentValue(&_CoreToken.CallOpts)
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCaller) FirstChequeNonceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "firstChequeNonceOf", account)
	return *ret0, err
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenSession) FirstChequeNonceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.FirstChequeNonceOf(&_CoreToken.CallOpts, account)
}

// FirstChequeNonceOf is a free data retrieval call binding the contract method 0x2227fd8b.
//
// Solidity: function firstChequeNonceOf(address account) view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) FirstChequeNonceOf(account common.Address) (*big.Int, error) {
	return _CoreToken.Contract.FirstChequeNonceOf(&_CoreToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_CoreToken *CoreTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_CoreToken *CoreTokenSession) Name() (string, error) {
	return _CoreToken.Contract.Name(&_CoreToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_CoreToken *CoreTokenCallerSession) Name() (string, error) {
	return _CoreToken.Contract.Name(&_CoreToken.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_CoreToken *CoreTokenCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "priceFeed")
	return *ret0, err
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_CoreToken *CoreTokenSession) PriceFeed() (common.Address, error) {
	return _CoreToken.Contract.PriceFeed(&_CoreToken.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_CoreToken *CoreTokenCallerSession) PriceFeed() (common.Address, error) {
	return _CoreToken.Contract.PriceFeed(&_CoreToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_CoreToken *CoreTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_CoreToken *CoreTokenSession) Symbol() (string, error) {
	return _CoreToken.Contract.Symbol(&_CoreToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_CoreToken *CoreTokenCallerSession) Symbol() (string, error) {
	return _CoreToken.Contract.Symbol(&_CoreToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CoreToken *CoreTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CoreToken *CoreTokenSession) TotalSupply() (*big.Int, error) {
	return _CoreToken.Contract.TotalSupply(&_CoreToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CoreToken *CoreTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CoreToken.Contract.TotalSupply(&_CoreToken.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_CoreToken *CoreTokenCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CoreToken.contract.Call(opts, out, "wrappedToken")
	return *ret0, err
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_CoreToken *CoreTokenSession) WrappedToken() (common.Address, error) {
	return _CoreToken.Contract.WrappedToken(&_CoreToken.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_CoreToken *CoreTokenCallerSession) WrappedToken() (common.Address, error) {
	return _CoreToken.Contract.WrappedToken(&_CoreToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Approve(&_CoreToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Approve(&_CoreToken.TransactOpts, spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) ApproveEquivalent(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "approveEquivalent", spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) ApproveEquivalent(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.ApproveEquivalent(&_CoreToken.TransactOpts, spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) ApproveEquivalent(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.ApproveEquivalent(&_CoreToken.TransactOpts, spender, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_CoreToken *CoreTokenTransactor) Buy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "buy", amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_CoreToken *CoreTokenSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Buy(&_CoreToken.TransactOpts, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_CoreToken *CoreTokenTransactorSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Buy(&_CoreToken.TransactOpts, amount)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_CoreToken *CoreTokenTransactor) CashBounty(opts *bind.TransactOpts, bounties []Bounty) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "cashBounty", bounties)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_CoreToken *CoreTokenSession) CashBounty(bounties []Bounty) (*types.Transaction, error) {
	return _CoreToken.Contract.CashBounty(&_CoreToken.TransactOpts, bounties)
}

// CashBounty is a paid mutator transaction binding the contract method 0x4cb53686.
//
// Solidity: function cashBounty([]Bounty bounties) returns()
func (_CoreToken *CoreTokenTransactorSession) CashBounty(bounties []Bounty) (*types.Transaction, error) {
	return _CoreToken.Contract.CashBounty(&_CoreToken.TransactOpts, bounties)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_CoreToken *CoreTokenTransactor) CashCheque(opts *bind.TransactOpts, cheque Cheque) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "cashCheque", cheque)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_CoreToken *CoreTokenSession) CashCheque(cheque Cheque) (*types.Transaction, error) {
	return _CoreToken.Contract.CashCheque(&_CoreToken.TransactOpts, cheque)
}

// CashCheque is a paid mutator transaction binding the contract method 0xbcc28adb.
//
// Solidity: function cashCheque(Cheque cheque) returns()
func (_CoreToken *CoreTokenTransactorSession) CashCheque(cheque Cheque) (*types.Transaction, error) {
	return _CoreToken.Contract.CashCheque(&_CoreToken.TransactOpts, cheque)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_CoreToken *CoreTokenTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_CoreToken *CoreTokenSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Sell(&_CoreToken.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_CoreToken *CoreTokenTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Sell(&_CoreToken.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Transfer(&_CoreToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.Transfer(&_CoreToken.TransactOpts, recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) TransferEquivalent(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "transferEquivalent", recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) TransferEquivalent(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferEquivalent(&_CoreToken.TransactOpts, recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) TransferEquivalent(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferEquivalent(&_CoreToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferFrom(&_CoreToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferFrom(&_CoreToken.TransactOpts, sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactor) TransferFromEquivalent(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.contract.Transact(opts, "transferFromEquivalent", sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenSession) TransferFromEquivalent(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferFromEquivalent(&_CoreToken.TransactOpts, sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_CoreToken *CoreTokenTransactorSession) TransferFromEquivalent(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreToken.Contract.TransferFromEquivalent(&_CoreToken.TransactOpts, sender, recipient, amount)
}

// CoreTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CoreToken contract.
type CoreTokenApprovalIterator struct {
	Event *CoreTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenApproval)
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
		it.Event = new(CoreTokenApproval)
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
func (it *CoreTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenApproval represents a Approval event raised by the CoreToken contract.
type CoreTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CoreToken *CoreTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CoreTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenApprovalIterator{contract: _CoreToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CoreToken *CoreTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CoreTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenApproval)
				if err := _CoreToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CoreToken *CoreTokenFilterer) ParseApproval(log types.Log) (*CoreTokenApproval, error) {
	event := new(CoreTokenApproval)
	if err := _CoreToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenBountyCashedIterator is returned from FilterBountyCashed and is used to iterate over the raw logs and unpacked data for BountyCashed events raised by the CoreToken contract.
type CoreTokenBountyCashedIterator struct {
	Event *CoreTokenBountyCashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenBountyCashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenBountyCashed)
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
		it.Event = new(CoreTokenBountyCashed)
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
func (it *CoreTokenBountyCashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenBountyCashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenBountyCashed represents a BountyCashed event raised by the CoreToken contract.
type CoreTokenBountyCashed struct {
	Owner       common.Address
	Target      common.Address
	Data        []byte
	Reward      *big.Int
	Nonce       *big.Int
	Deadline    *big.Int
	EnergyLimit *big.Int
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBountyCashed is a free log retrieval operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_CoreToken *CoreTokenFilterer) FilterBountyCashed(opts *bind.FilterOpts, owner []common.Address, target []common.Address) (*CoreTokenBountyCashedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "BountyCashed", ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenBountyCashedIterator{contract: _CoreToken.contract, event: "BountyCashed", logs: logs, sub: sub}, nil
}

// WatchBountyCashed is a free log subscription operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_CoreToken *CoreTokenFilterer) WatchBountyCashed(opts *bind.WatchOpts, sink chan<- *CoreTokenBountyCashed, owner []common.Address, target []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "BountyCashed", ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenBountyCashed)
				if err := _CoreToken.contract.UnpackLog(event, "BountyCashed", log); err != nil {
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

// ParseBountyCashed is a log parse operation binding the contract event 0x43300dc0f903c4632fd4d7a18a0f87a73747ad0543474a60d731df03aef16b98.
//
// Solidity: event BountyCashed(address indexed owner, address indexed target, bytes data, uint256 reward, uint256 nonce, uint256 deadline, uint256 energyLimit, bool success)
func (_CoreToken *CoreTokenFilterer) ParseBountyCashed(log types.Log) (*CoreTokenBountyCashed, error) {
	event := new(CoreTokenBountyCashed)
	if err := _CoreToken.contract.UnpackLog(event, "BountyCashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the CoreToken contract.
type CoreTokenBurnIterator struct {
	Event *CoreTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenBurn)
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
		it.Event = new(CoreTokenBurn)
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
func (it *CoreTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenBurn represents a Burn event raised by the CoreToken contract.
type CoreTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_CoreToken *CoreTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*CoreTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenBurnIterator{contract: _CoreToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_CoreToken *CoreTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *CoreTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenBurn)
				if err := _CoreToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_CoreToken *CoreTokenFilterer) ParseBurn(log types.Log) (*CoreTokenBurn, error) {
	event := new(CoreTokenBurn)
	if err := _CoreToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenChequeCashIterator is returned from FilterChequeCash and is used to iterate over the raw logs and unpacked data for ChequeCash events raised by the CoreToken contract.
type CoreTokenChequeCashIterator struct {
	Event *CoreTokenChequeCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenChequeCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenChequeCash)
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
		it.Event = new(CoreTokenChequeCash)
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
func (it *CoreTokenChequeCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenChequeCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenChequeCash represents a ChequeCash event raised by the CoreToken contract.
type CoreTokenChequeCash struct {
	Owner    common.Address
	Spender  common.Address
	Amount   *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChequeCash is a free log retrieval operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_CoreToken *CoreTokenFilterer) FilterChequeCash(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CoreTokenChequeCashIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "ChequeCash", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenChequeCashIterator{contract: _CoreToken.contract, event: "ChequeCash", logs: logs, sub: sub}, nil
}

// WatchChequeCash is a free log subscription operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_CoreToken *CoreTokenFilterer) WatchChequeCash(opts *bind.WatchOpts, sink chan<- *CoreTokenChequeCash, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "ChequeCash", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenChequeCash)
				if err := _CoreToken.contract.UnpackLog(event, "ChequeCash", log); err != nil {
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

// ParseChequeCash is a log parse operation binding the contract event 0x313836ac2698246e2377f8d8005d6e55a59b9acee376fdaafda7653ab33988b0.
//
// Solidity: event ChequeCash(address indexed owner, address indexed spender, uint256 amount, uint256 nonce, uint256 deadline)
func (_CoreToken *CoreTokenFilterer) ParseChequeCash(log types.Log) (*CoreTokenChequeCash, error) {
	event := new(CoreTokenChequeCash)
	if err := _CoreToken.contract.UnpackLog(event, "ChequeCash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CoreToken contract.
type CoreTokenMintIterator struct {
	Event *CoreTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenMint)
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
		it.Event = new(CoreTokenMint)
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
func (it *CoreTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenMint represents a Mint event raised by the CoreToken contract.
type CoreTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*CoreTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenMintIterator{contract: _CoreToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CoreTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenMint)
				if err := _CoreToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) ParseMint(log types.Log) (*CoreTokenMint, error) {
	event := new(CoreTokenMint)
	if err := _CoreToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoreTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CoreToken contract.
type CoreTokenTransferIterator struct {
	Event *CoreTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTokenTransfer)
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
		it.Event = new(CoreTokenTransfer)
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
func (it *CoreTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTokenTransfer represents a Transfer event raised by the CoreToken contract.
type CoreTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CoreTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CoreToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CoreTokenTransferIterator{contract: _CoreToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CoreTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CoreToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTokenTransfer)
				if err := _CoreToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CoreToken *CoreTokenFilterer) ParseTransfer(log types.Log) (*CoreTokenTransfer, error) {
	event := new(CoreTokenTransfer)
	if err := _CoreToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EIP712ABI is the input ABI used to generate the binding from.
const EIP712ABI = "[]"

// EIP712 is an auto generated Go binding around an Core contract.
type EIP712 struct {
	EIP712Caller     // Read-only binding to the contract
	EIP712Transactor // Write-only binding to the contract
	EIP712Filterer   // Log filterer for contract events
}

// EIP712Caller is an auto generated read-only Go binding around an Core contract.
type EIP712Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Transactor is an auto generated write-only Go binding around an Core contract.
type EIP712Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Filterer is an auto generated log filtering Go binding around an Core contract events.
type EIP712Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Session is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type EIP712Session struct {
	Contract     *EIP712           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712CallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type EIP712CallerSession struct {
	Contract *EIP712Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP712TransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type EIP712TransactorSession struct {
	Contract     *EIP712Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712Raw is an auto generated low-level Go binding around an Core contract.
type EIP712Raw struct {
	Contract *EIP712 // Generic contract binding to access the raw methods on
}

// EIP712CallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type EIP712CallerRaw struct {
	Contract *EIP712Caller // Generic read-only contract binding to access the raw methods on
}

// EIP712TransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type EIP712TransactorRaw struct {
	Contract *EIP712Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712 creates a new instance of EIP712, bound to a specific deployed contract.
func NewEIP712(address common.Address, backend bind.ContractBackend) (*EIP712, error) {
	contract, err := bindEIP712(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// NewEIP712Caller creates a new read-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Caller(address common.Address, caller bind.ContractCaller) (*EIP712Caller, error) {
	contract, err := bindEIP712(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Caller{contract: contract}, nil
}

// NewEIP712Transactor creates a new write-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP712Transactor, error) {
	contract, err := bindEIP712(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Transactor{contract: contract}, nil
}

// NewEIP712Filterer creates a new log filterer instance of EIP712, bound to a specific deployed contract.
func NewEIP712Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP712Filterer, error) {
	contract, err := bindEIP712(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712Filterer{contract: contract}, nil
}

// bindEIP712 binds a generic wrapper to an already deployed contract.
func bindEIP712(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP712ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.EIP712Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20 is an auto generated Go binding around an Core contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Core contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Core contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Core contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Core contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EquivalentTokenABI is the input ABI used to generate the binding from.
const EquivalentTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"equivalentValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFromEquivalent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EquivalentToken is an auto generated Go binding around an Core contract.
type EquivalentToken struct {
	EquivalentTokenCaller     // Read-only binding to the contract
	EquivalentTokenTransactor // Write-only binding to the contract
	EquivalentTokenFilterer   // Log filterer for contract events
}

// EquivalentTokenCaller is an auto generated read-only Go binding around an Core contract.
type EquivalentTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EquivalentTokenTransactor is an auto generated write-only Go binding around an Core contract.
type EquivalentTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EquivalentTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type EquivalentTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EquivalentTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type EquivalentTokenSession struct {
	Contract     *EquivalentToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EquivalentTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type EquivalentTokenCallerSession struct {
	Contract *EquivalentTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EquivalentTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type EquivalentTokenTransactorSession struct {
	Contract     *EquivalentTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EquivalentTokenRaw is an auto generated low-level Go binding around an Core contract.
type EquivalentTokenRaw struct {
	Contract *EquivalentToken // Generic contract binding to access the raw methods on
}

// EquivalentTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type EquivalentTokenCallerRaw struct {
	Contract *EquivalentTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EquivalentTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type EquivalentTokenTransactorRaw struct {
	Contract *EquivalentTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEquivalentToken creates a new instance of EquivalentToken, bound to a specific deployed contract.
func NewEquivalentToken(address common.Address, backend bind.ContractBackend) (*EquivalentToken, error) {
	contract, err := bindEquivalentToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EquivalentToken{EquivalentTokenCaller: EquivalentTokenCaller{contract: contract}, EquivalentTokenTransactor: EquivalentTokenTransactor{contract: contract}, EquivalentTokenFilterer: EquivalentTokenFilterer{contract: contract}}, nil
}

// NewEquivalentTokenCaller creates a new read-only instance of EquivalentToken, bound to a specific deployed contract.
func NewEquivalentTokenCaller(address common.Address, caller bind.ContractCaller) (*EquivalentTokenCaller, error) {
	contract, err := bindEquivalentToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenCaller{contract: contract}, nil
}

// NewEquivalentTokenTransactor creates a new write-only instance of EquivalentToken, bound to a specific deployed contract.
func NewEquivalentTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EquivalentTokenTransactor, error) {
	contract, err := bindEquivalentToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenTransactor{contract: contract}, nil
}

// NewEquivalentTokenFilterer creates a new log filterer instance of EquivalentToken, bound to a specific deployed contract.
func NewEquivalentTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EquivalentTokenFilterer, error) {
	contract, err := bindEquivalentToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenFilterer{contract: contract}, nil
}

// bindEquivalentToken binds a generic wrapper to an already deployed contract.
func bindEquivalentToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EquivalentTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EquivalentToken *EquivalentTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EquivalentToken.Contract.EquivalentTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EquivalentToken *EquivalentTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EquivalentToken.Contract.EquivalentTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EquivalentToken *EquivalentTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EquivalentToken.Contract.EquivalentTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EquivalentToken *EquivalentTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EquivalentToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EquivalentToken *EquivalentTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EquivalentToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EquivalentToken *EquivalentTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EquivalentToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EquivalentToken *EquivalentTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EquivalentToken *EquivalentTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EquivalentToken.Contract.Allowance(&_EquivalentToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EquivalentToken *EquivalentTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EquivalentToken.Contract.Allowance(&_EquivalentToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EquivalentToken *EquivalentTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EquivalentToken *EquivalentTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EquivalentToken.Contract.BalanceOf(&_EquivalentToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EquivalentToken *EquivalentTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EquivalentToken.Contract.BalanceOf(&_EquivalentToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_EquivalentToken *EquivalentTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_EquivalentToken *EquivalentTokenSession) Decimals() (uint8, error) {
	return _EquivalentToken.Contract.Decimals(&_EquivalentToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_EquivalentToken *EquivalentTokenCallerSession) Decimals() (uint8, error) {
	return _EquivalentToken.Contract.Decimals(&_EquivalentToken.CallOpts)
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_EquivalentToken *EquivalentTokenCaller) EquivalentValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "equivalentValue")
	return *ret0, err
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_EquivalentToken *EquivalentTokenSession) EquivalentValue() (*big.Int, error) {
	return _EquivalentToken.Contract.EquivalentValue(&_EquivalentToken.CallOpts)
}

// EquivalentValue is a free data retrieval call binding the contract method 0xc64a4ca0.
//
// Solidity: function equivalentValue() view returns(uint256)
func (_EquivalentToken *EquivalentTokenCallerSession) EquivalentValue() (*big.Int, error) {
	return _EquivalentToken.Contract.EquivalentValue(&_EquivalentToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_EquivalentToken *EquivalentTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_EquivalentToken *EquivalentTokenSession) Name() (string, error) {
	return _EquivalentToken.Contract.Name(&_EquivalentToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_EquivalentToken *EquivalentTokenCallerSession) Name() (string, error) {
	return _EquivalentToken.Contract.Name(&_EquivalentToken.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_EquivalentToken *EquivalentTokenCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "priceFeed")
	return *ret0, err
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_EquivalentToken *EquivalentTokenSession) PriceFeed() (common.Address, error) {
	return _EquivalentToken.Contract.PriceFeed(&_EquivalentToken.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x449207ec.
//
// Solidity: function priceFeed() view returns(address)
func (_EquivalentToken *EquivalentTokenCallerSession) PriceFeed() (common.Address, error) {
	return _EquivalentToken.Contract.PriceFeed(&_EquivalentToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_EquivalentToken *EquivalentTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_EquivalentToken *EquivalentTokenSession) Symbol() (string, error) {
	return _EquivalentToken.Contract.Symbol(&_EquivalentToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_EquivalentToken *EquivalentTokenCallerSession) Symbol() (string, error) {
	return _EquivalentToken.Contract.Symbol(&_EquivalentToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EquivalentToken *EquivalentTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EquivalentToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EquivalentToken *EquivalentTokenSession) TotalSupply() (*big.Int, error) {
	return _EquivalentToken.Contract.TotalSupply(&_EquivalentToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EquivalentToken *EquivalentTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _EquivalentToken.Contract.TotalSupply(&_EquivalentToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.Approve(&_EquivalentToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.Approve(&_EquivalentToken.TransactOpts, spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) ApproveEquivalent(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "approveEquivalent", spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) ApproveEquivalent(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.ApproveEquivalent(&_EquivalentToken.TransactOpts, spender, amount)
}

// ApproveEquivalent is a paid mutator transaction binding the contract method 0x542f48f3.
//
// Solidity: function approveEquivalent(address spender, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) ApproveEquivalent(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.ApproveEquivalent(&_EquivalentToken.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.Transfer(&_EquivalentToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.Transfer(&_EquivalentToken.TransactOpts, recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) TransferEquivalent(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "transferEquivalent", recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) TransferEquivalent(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferEquivalent(&_EquivalentToken.TransactOpts, recipient, amount)
}

// TransferEquivalent is a paid mutator transaction binding the contract method 0x89b9ce32.
//
// Solidity: function transferEquivalent(address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) TransferEquivalent(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferEquivalent(&_EquivalentToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferFrom(&_EquivalentToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferFrom(&_EquivalentToken.TransactOpts, sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactor) TransferFromEquivalent(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.contract.Transact(opts, "transferFromEquivalent", sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenSession) TransferFromEquivalent(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferFromEquivalent(&_EquivalentToken.TransactOpts, sender, recipient, amount)
}

// TransferFromEquivalent is a paid mutator transaction binding the contract method 0x380544b3.
//
// Solidity: function transferFromEquivalent(address sender, address recipient, uint256 amount) returns(bool)
func (_EquivalentToken *EquivalentTokenTransactorSession) TransferFromEquivalent(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EquivalentToken.Contract.TransferFromEquivalent(&_EquivalentToken.TransactOpts, sender, recipient, amount)
}

// EquivalentTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EquivalentToken contract.
type EquivalentTokenApprovalIterator struct {
	Event *EquivalentTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EquivalentTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EquivalentTokenApproval)
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
		it.Event = new(EquivalentTokenApproval)
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
func (it *EquivalentTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EquivalentTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EquivalentTokenApproval represents a Approval event raised by the EquivalentToken contract.
type EquivalentTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EquivalentTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EquivalentToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenApprovalIterator{contract: _EquivalentToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EquivalentTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EquivalentToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EquivalentTokenApproval)
				if err := _EquivalentToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) ParseApproval(log types.Log) (*EquivalentTokenApproval, error) {
	event := new(EquivalentTokenApproval)
	if err := _EquivalentToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EquivalentTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the EquivalentToken contract.
type EquivalentTokenBurnIterator struct {
	Event *EquivalentTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EquivalentTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EquivalentTokenBurn)
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
		it.Event = new(EquivalentTokenBurn)
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
func (it *EquivalentTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EquivalentTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EquivalentTokenBurn represents a Burn event raised by the EquivalentToken contract.
type EquivalentTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*EquivalentTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _EquivalentToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenBurnIterator{contract: _EquivalentToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *EquivalentTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _EquivalentToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EquivalentTokenBurn)
				if err := _EquivalentToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) ParseBurn(log types.Log) (*EquivalentTokenBurn, error) {
	event := new(EquivalentTokenBurn)
	if err := _EquivalentToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EquivalentTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the EquivalentToken contract.
type EquivalentTokenMintIterator struct {
	Event *EquivalentTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EquivalentTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EquivalentTokenMint)
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
		it.Event = new(EquivalentTokenMint)
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
func (it *EquivalentTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EquivalentTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EquivalentTokenMint represents a Mint event raised by the EquivalentToken contract.
type EquivalentTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*EquivalentTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EquivalentToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenMintIterator{contract: _EquivalentToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *EquivalentTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EquivalentToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EquivalentTokenMint)
				if err := _EquivalentToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) ParseMint(log types.Log) (*EquivalentTokenMint, error) {
	event := new(EquivalentTokenMint)
	if err := _EquivalentToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EquivalentTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EquivalentToken contract.
type EquivalentTokenTransferIterator struct {
	Event *EquivalentTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EquivalentTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EquivalentTokenTransfer)
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
		it.Event = new(EquivalentTokenTransfer)
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
func (it *EquivalentTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EquivalentTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EquivalentTokenTransfer represents a Transfer event raised by the EquivalentToken contract.
type EquivalentTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EquivalentTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EquivalentToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EquivalentTokenTransferIterator{contract: _EquivalentToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EquivalentTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EquivalentToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EquivalentTokenTransfer)
				if err := _EquivalentToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EquivalentToken *EquivalentTokenFilterer) ParseTransfer(log types.Log) (*EquivalentTokenTransfer, error) {
	event := new(EquivalentTokenTransfer)
	if err := _EquivalentToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterABI is the input ABI used to generate the binding from.
const KYCTransmitterABI = "[{\"inputs\":[{\"internalType\":\"contractKYCVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractCoreToken\",\"name\":\"priceToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transmission\",\"type\":\"bytes32\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transmission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"fields\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transmission\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Invalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PriceAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCoreToken\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"changePriceAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCoreToken\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"changePriceToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transmission\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structCheque\",\"name\":\"providerCheque\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structCheque\",\"name\":\"requesterCheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"providerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"requesterSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"fields\",\"type\":\"bytes32[]\"}],\"name\":\"initiate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transmission\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"fields\",\"type\":\"bytes32[]\"}],\"name\":\"invalidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"priceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceToken\",\"outputs\":[{\"internalType\":\"contractCoreToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KYCTransmitterBin is the compiled bytecode used for deploying new contracts.
var KYCTransmitterBin = "0x6101206040523480156200001257600080fd5b5060405162003bac38038062003bac833981810160405281019062000038919062000261565b6040518060400160405280600e81526020017f4b59435472616e736d69747465720000000000000000000000000000000000008152506040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250600060405180608001604052806054815260200162003b586054913980519060200120905080608081815250506000838051906020012090508060a081815250506000838051906020012090508060c081815250504660e0818152505062000114838383620001f760201b60201c565b61010081815250505050505050336000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555081600260016101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055505050620003d8565b6000838383463060405160200162000214959493929190620002d5565b6040516020818303038152906040528051906020012090509392505050565b6000815190506200024481620003a4565b92915050565b6000815190506200025b81620003be565b92915050565b600080604083850312156200027557600080fd5b600062000285858286016200024a565b9250506020620002988582860162000233565b9150509250929050565b620002ad8162000332565b82525050565b620002be8162000346565b82525050565b620002cf816200039a565b82525050565b600060a082019050620002ec6000830188620002b3565b620002fb6020830187620002b3565b6200030a6040830186620002b3565b620003196060830185620002c4565b620003286080830184620002a2565b9695505050505050565b60006200033f8262000378565b9050919050565b6000819050919050565b60006200035d8262000332565b9050919050565b6000620003718262000332565b9050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b620003af8162000350565b8114620003bb57600080fd5b50565b620003c98162000364565b8114620003d557600080fd5b50565b60805160a05160c05160e0516101005161373b6200041d6000396000611b1301526000611a7d01526000611ae801526000611ac701526000611aa6015261373b6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063614b988711610097578063b9d319b611610066578063b9d319b614610258578063bd8f8ebe14610276578063e05cab1214610280578063eb8325fb1461029c576100f5565b8063614b9887146101e657806384e479661461020457806389a7fa4114610220578063946aba7b1461023c576100f5565b80632b610a34116100d35780632b610a3414610162578063344dba961461019257806347d4c036146101ae5780634e5a827c146101ca576100f5565b80630d06de54146100fa5780631993cd0b1461011657806325e8d28214610146575b600080fd5b610114600480360381019061010f9190612379565b6102ba565b005b610130600480360381019061012b919061221d565b610bd6565b60405161013d9190612ac2565b60405180910390f35b610160600480360381019061015b919061221d565b610bea565b005b61017c600480360381019061017791906121f4565b610c8c565b6040516101899190612dd9565b60405180910390f35b6101ac60048036038101906101a791906122ad565b610ca9565b005b6101c860048036038101906101c39190612350565b610d91565b005b6101e460048036038101906101df919061221d565b610ea4565b005b6101ee610f46565b6040516101fb9190612ac2565b60405180910390f35b61021e600480360381019061021991906122e9565b610f5d565b005b61023a600480360381019061023591906121a2565b61139a565b005b61025660048036038101906102519190612259565b611541565b005b610260611856565b60405161026d9190612bdc565b60405180910390f35b61027e611882565b005b61029a600480360381019061029591906121f4565b6119af565b005b6102a4611a07565b6040516102b191906129e7565b60405180910390f35b600260009054906101000a900460ff161561030a576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161030190612cf7565b60405180910390fd5b6000815190506000600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1663c64a4ca06040518163ffffffff1660e01b815260040160206040518083038186803b15801561037e57600080fd5b505afa158015610392573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b69190612468565b905060008267ffffffffffffffff8111156103fa577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156104285781602001602082028036833780820191505090505b5090506000805b8481101561054957600060046000888481518110610476577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151815260200190815260200160002054905060008114156104d3576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016104ca90612c97565b60405180910390fd5b84816104df9190612f92565b90508084838151811061051b577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505080836105339190612f3c565b9250508080610541906130e9565b91505061042f565b5060008560405160200161055d9190612999565b60405160208183030381529060405280519060200120905060006040518060800160405280604a815260200161364e604a9139805190602001208a600001518c606001518d60800151856040516020016105bb959493929190612add565b6040516020818303038152906040528051906020012090506105dd818a611a32565b75ffffffffffffffffffffffffffffffffffffffffffff168b6000015175ffffffffffffffffffffffffffffffffffffffffffff1614610652576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161064990612cb7565b60405180910390fd5b600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1663bcc28adb8c6040518263ffffffff1660e01b81526004016106b19190612db7565b600060405180830381600087803b1580156106cb57600080fd5b505af11580156106df573d6000803e3d6000fd5b50505050600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff166331f2e6798c6000015130866040518463ffffffff1660e01b815260040161074a93929190612a02565b602060405180830381600087803b15801561076457600080fd5b505af1158015610778573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079c91906121cb565b50600060405180608001604052806049815260200161369860499139805190602001208c600001518c606001518d60800151866040516020016107e3959493929190612add565b604051602081830303815290604052805190602001209050610805818a611a32565b75ffffffffffffffffffffffffffffffffffffffffffff168b6000015175ffffffffffffffffffffffffffffffffffffffffffff161461087a576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161087190612c37565b60405180910390fd5b600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1663bcc28adb8c6040518263ffffffff1660e01b81526004016108d99190612db7565b600060405180830381600087803b1580156108f357600080fd5b505af1158015610907573d6000803e3d6000fd5b50505050600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff166331f2e6798c60000151308760026109609190612f92565b6040518463ffffffff1660e01b815260040161097e93929190612a02565b602060405180830381600087803b15801561099857600080fd5b505af11580156109ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d091906121cb565b50600060055460001b90506001600560008282546109ee9190612f3c565b925050819055508c600001516006600083815260200190815260200160002060006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055508b600001516007600083815260200190815260200160002060006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff166008600083815260200190815260200160002060006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555084600960008381526020019081526020016000208190555083600a6000838152602001908152602001600020819055508b6000015175ffffffffffffffffffffffffffffffffffffffffffff168d6000015175ffffffffffffffffffffffffffffffffffffffffffff16827fd5e9cd7af10b71d0052610468eba244b812851c4e582629a8e52faa2484c56aa8c8a604051610bbf929190612a8b565b60405180910390a450505050505050505050505050565b6000610be28383611bb7565b905092915050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610c7e576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610c7590612d57565b60405180910390fd5b610c888282611c23565b5050565b600060046000838152602001908152602001600020549050919050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610d3d576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610d3490612d57565b60405180910390fd5b806004600084815260200190815260200160002081905550817fe0ef2e1c942634b7a75983cbbf5697092369f04f3ad3786c2466dd7f88e848f882604051610d859190612dd9565b60405180910390a25050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610e25576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610e1c90612d57565b60405180910390fd5b80600360006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055507f1716a3f72bcf68266196e4932576023ececd691daf05d5eb758f69e181adc82881604051610e999190612bdc565b60405180910390a150565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610f38576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610f2f90612d57565b60405180910390fd5b610f428282611d6d565b5050565b6000600260009054906101000a900460ff16905090565b600081604051602001610f709190612999565b60405160208183030381529060405280519060200120905080600a60008681526020019081526020016000205414610fdd576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610fd490612c57565b60405180910390fd5b81518310611020576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161101790612c17565b60405180910390fd5b600082848151811061105b577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151905060006006600087815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1690506000600260019054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1663c9e1424883856040518363ffffffff1660e01b8152600401611102929190612a39565b60206040518083038186803b15801561111a57600080fd5b505afa15801561112e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115291906121cb565b90508015611195576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161118c90612d37565b60405180910390fd5b60006007600089815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1690506000600860008a815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1690506000600960008b81526020019081526020016000205490508175ffffffffffffffffffffffffffffffffffffffffffff16634b40e90186836040518363ffffffff1660e01b815260040161125e929190612a62565b602060405180830381600087803b15801561127857600080fd5b505af115801561128c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b091906121cb565b508175ffffffffffffffffffffffffffffffffffffffffffff16634b40e901848360026112dd9190612f92565b6040518363ffffffff1660e01b81526004016112fa929190612a62565b602060405180830381600087803b15801561131457600080fd5b505af1158015611328573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134c91906121cb565b506113568a611e5f565b897f99905deb70a82b96b7ce3822e300fb0846eac837005061d54a943f56e1df6bb08a6040516113869190612dd9565b60405180910390a250505050505050505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff161461142e576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161142590612d57565b60405180910390fd5b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168175ffffffffffffffffffffffffffffffffffffffffffff1614156114c3576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016114ba90612d77565b60405180910390fd5b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055507fc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be108160405161153691906129e7565b60405180910390a150565b60006040518060400160405280601d81526020017f436f6e6669726d2862797465733332207472616e736d697373696f6e290000008152508051906020012083604051602001611592929190612b30565b60405160208183030381529060405280519060200120905060006007600085815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1690506115ee8284611a32565b75ffffffffffffffffffffffffffffffffffffffffffff168175ffffffffffffffffffffffffffffffffffffffffffff161461165f576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161165690612d97565b60405180910390fd5b60006006600086815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905060006008600087815260200190815260200160002060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1690506000600960008881526020019081526020016000205490508175ffffffffffffffffffffffffffffffffffffffffffff16634b40e901848360026117179190612f92565b6040518363ffffffff1660e01b8152600401611734929190612a62565b602060405180830381600087803b15801561174e57600080fd5b505af1158015611762573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061178691906121cb565b508175ffffffffffffffffffffffffffffffffffffffffffff16634b40e90185836040518363ffffffff1660e01b81526004016117c4929190612a62565b602060405180830381600087803b1580156117de57600080fd5b505af11580156117f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181691906121cb565b5061182087611e5f565b867ffe8a98fc20cafa18a1d9c00074771ee084073c363aabb2c34c1b2905f0b1202060405160405180910390a250505050505050565b6000600360009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614611916576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161190d90612d57565b60405180910390fd5b600260009054906101000a900460ff1615611966576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161195d90612cf7565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba2560405160405180910390a1565b806119ba8133611bb7565b6119f9576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016119f090612cd7565b60405180910390fd5b611a038233611c23565b5050565b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905090565b600060ab825114611a78576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611a6f90612c77565b60405180910390fd5b6000467f000000000000000000000000000000000000000000000000000000000000000014611b1157611b0c7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611f38565b611b33565b7f00000000000000000000000000000000000000000000000000000000000000005b905060008185604051602001611b4a9291906129b0565b6040516020818303038152906040528051906020012090506001818560405160008152602001604052604051611b81929190612bac565b6020604051602081039080840390855afa158015611ba3573d6000803e3d6000fd5b505050602060405103519250505092915050565b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6001600083815260200190815260200160002060008275ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611cc4576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611cbb90612bf7565b60405180910390fd5b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf46685082604051611d6191906129e7565b60405180910390a25050565b611d778282611bb7565b15611db7576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611dae90612d17565b60405180910390fd5b600180600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb82604051611e5391906129e7565b60405180910390a25050565b6006600082815260200190815260200160002060006101000a81549075ffffffffffffffffffffffffffffffffffffffffffff02191690556007600082815260200190815260200160002060006101000a81549075ffffffffffffffffffffffffffffffffffffffffffff02191690556008600082815260200190815260200160002060006101000a81549075ffffffffffffffffffffffffffffffffffffffffffff02191690556009600082815260200190815260200160002060009055600a60008281526020019081526020016000206000905550565b60008383834630604051602001611f53959493929190612b59565b6040516020818303038152906040528051906020012090509392505050565b6000611f85611f8084612e19565b612df4565b90508083825260208201905082856020860282011115611fa457600080fd5b60005b85811015611fd45781611fba8882612070565b845260208401935060208301925050600181019050611fa7565b5050509392505050565b6000611ff1611fec84612e45565b612df4565b90508281526020810184848401111561200957600080fd5b612014848285613076565b509392505050565b60008135905061202b816135da565b92915050565b600082601f83011261204257600080fd5b8135612052848260208601611f72565b91505092915050565b60008151905061206a816135f1565b92915050565b60008135905061207f81613608565b92915050565b600082601f83011261209657600080fd5b81356120a6848260208601611fde565b91505092915050565b6000813590506120be8161361f565b92915050565b600060c082840312156120d657600080fd5b6120e060c0612df4565b905060006120f08482850161201c565b60008301525060206121048482850161201c565b602083015250604061211884828501612178565b604083015250606061212c84828501612178565b606083015250608061214084828501612178565b60808301525060a082013567ffffffffffffffff81111561216057600080fd5b61216c84828501612085565b60a08301525092915050565b60008135905061218781613636565b92915050565b60008151905061219c81613636565b92915050565b6000602082840312156121b457600080fd5b60006121c28482850161201c565b91505092915050565b6000602082840312156121dd57600080fd5b60006121eb8482850161205b565b91505092915050565b60006020828403121561220657600080fd5b600061221484828501612070565b91505092915050565b6000806040838503121561223057600080fd5b600061223e85828601612070565b925050602061224f8582860161201c565b9150509250929050565b6000806040838503121561226c57600080fd5b600061227a85828601612070565b925050602083013567ffffffffffffffff81111561229757600080fd5b6122a385828601612085565b9150509250929050565b600080604083850312156122c057600080fd5b60006122ce85828601612070565b92505060206122df85828601612178565b9150509250929050565b6000806000606084860312156122fe57600080fd5b600061230c86828701612070565b935050602061231d86828701612178565b925050604084013567ffffffffffffffff81111561233a57600080fd5b61234686828701612031565b9150509250925092565b60006020828403121561236257600080fd5b6000612370848285016120af565b91505092915050565b600080600080600060a0868803121561239157600080fd5b600086013567ffffffffffffffff8111156123ab57600080fd5b6123b7888289016120c4565b955050602086013567ffffffffffffffff8111156123d457600080fd5b6123e0888289016120c4565b945050604086013567ffffffffffffffff8111156123fd57600080fd5b61240988828901612085565b935050606086013567ffffffffffffffff81111561242657600080fd5b61243288828901612085565b925050608086013567ffffffffffffffff81111561244f57600080fd5b61245b88828901612031565b9150509295509295909350565b60006020828403121561247a57600080fd5b60006124888482850161218d565b91505092915050565b600061249d8383612620565b60208301905092915050565b60006124b5838361263e565b60208301905092915050565b60006124cd838361297b565b60208301905092915050565b6124e281612fec565b82525050565b6124f181612fec565b82525050565b600061250282612e96565b61250c8185612ed1565b935061251783612e76565b8060005b8381101561254857815161252f8882612491565b975061253a83612eb7565b92505060018101905061251b565b5085935050505092915050565b600061256082612e96565b61256a8185612ee2565b935061257583612e76565b8060005b838110156125a657815161258d88826124a9565b975061259883612eb7565b925050600181019050612579565b5085935050505092915050565b60006125be82612ea1565b6125c88185612eed565b93506125d383612e86565b8060005b838110156126045781516125eb88826124c1565b97506125f683612ec4565b9250506001810190506125d7565b5085935050505092915050565b61261a81612ffe565b82525050565b6126298161300a565b82525050565b6126388161300a565b82525050565b6126478161300a565b82525050565b61265e6126598261300a565b613132565b82525050565b600061266f82612eac565b6126798185612efe565b9350612689818560208601613085565b6126928161319a565b840191505092915050565b60006126a882612eac565b6126b28185612f0f565b93506126c2818560208601613085565b6126cb8161319a565b840191505092915050565b6126df81613052565b82525050565b60006126f2602283612f20565b91506126fd826131ab565b604082019050919050565b6000612715602083612f20565b9150612720826131fa565b602082019050919050565b6000612738602d83612f20565b915061274382613223565b604082019050919050565b600061275b602483612f20565b915061276682613272565b604082019050919050565b600061277e602583612f20565b9150612789826132c1565b604082019050919050565b60006127a1601d83612f20565b91506127ac82613310565b602082019050919050565b60006127c4602c83612f20565b91506127cf82613339565b604082019050919050565b60006127e7602583612f20565b91506127f282613388565b604082019050919050565b600061280a602583612f20565b9150612815826133d7565b604082019050919050565b600061282d602683612f20565b915061283882613426565b604082019050919050565b6000612850602783612f20565b915061285b82613475565b604082019050919050565b6000612873602583612f20565b915061287e826134c4565b604082019050919050565b6000612896600283612f31565b91506128a182613513565b600282019050919050565b60006128b9603583612f20565b91506128c48261353c565b604082019050919050565b60006128dc602483612f20565b91506128e78261358b565b604082019050919050565b600060c08301600083015161290a60008601826124d9565b50602083015161291d60208601826124d9565b506040830151612930604086018261297b565b506060830151612943606086018261297b565b506080830151612956608086018261297b565b5060a083015184820360a086015261296e8282612664565b9150508091505092915050565b61298481613048565b82525050565b61299381613048565b82525050565b60006129a58284612555565b915081905092915050565b60006129bb82612889565b91506129c7828561264d565b6020820191506129d7828461264d565b6020820191508190509392505050565b60006020820190506129fc60008301846124e8565b92915050565b6000606082019050612a1760008301866124e8565b612a2460208301856124e8565b612a31604083018461298a565b949350505050565b6000604082019050612a4e60008301856124e8565b612a5b602083018461262f565b9392505050565b6000604082019050612a7760008301856124e8565b612a84602083018461298a565b9392505050565b60006040820190508181036000830152612aa581856124f7565b90508181036020830152612ab981846125b3565b90509392505050565b6000602082019050612ad76000830184612611565b92915050565b600060a082019050612af2600083018861262f565b612aff60208301876124e8565b612b0c604083018661298a565b612b19606083018561298a565b612b26608083018461262f565b9695505050505050565b6000604082019050612b45600083018561262f565b612b52602083018461262f565b9392505050565b600060a082019050612b6e600083018861262f565b612b7b602083018761262f565b612b88604083018661262f565b612b95606083018561298a565b612ba260808301846124e8565b9695505050505050565b6000604082019050612bc1600083018561262f565b8181036020830152612bd3818461269d565b90509392505050565b6000602082019050612bf160008301846126d6565b92915050565b60006020820190508181036000830152612c10816126e5565b9050919050565b60006020820190508181036000830152612c3081612708565b9050919050565b60006020820190508181036000830152612c508161272b565b9050919050565b60006020820190508181036000830152612c708161274e565b9050919050565b60006020820190508181036000830152612c9081612771565b9050919050565b60006020820190508181036000830152612cb081612794565b9050919050565b60006020820190508181036000830152612cd0816127b7565b9050919050565b60006020820190508181036000830152612cf0816127da565b9050919050565b60006020820190508181036000830152612d10816127fd565b9050919050565b60006020820190508181036000830152612d3081612820565b9050919050565b60006020820190508181036000830152612d5081612843565b9050919050565b60006020820190508181036000830152612d7081612866565b9050919050565b60006020820190508181036000830152612d90816128ac565b9050919050565b60006020820190508181036000830152612db0816128cf565b9050919050565b60006020820190508181036000830152612dd181846128f2565b905092915050565b6000602082019050612dee600083018461298a565b92915050565b6000612dfe612e0f565b9050612e0a82826130b8565b919050565b6000604051905090565b600067ffffffffffffffff821115612e3457612e3361316b565b5b602082029050602081019050919050565b600067ffffffffffffffff821115612e6057612e5f61316b565b5b612e698261319a565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000612f4782613048565b9150612f5283613048565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612f8757612f8661313c565b5b828201905092915050565b6000612f9d82613048565b9150612fa883613048565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612fe157612fe061313c565b5b828202905092915050565b6000612ff782613026565b9050919050565b60008115159050919050565b6000819050919050565b600061301f82612fec565b9050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061305d82613064565b9050919050565b600061306f82613026565b9050919050565b82818337600083830152505050565b60005b838110156130a3578082015181840152602081019050613088565b838111156130b2576000848401525b50505050565b6130c18261319a565b810181811067ffffffffffffffff821117156130e0576130df61316b565b5b80604052505050565b60006130f482613048565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156131275761312661313c565b5b600182019050919050565b6000819050919050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f416363657373436f6e74726f6c3a20726f6c65206973206e6f74206772616e7460008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b7f4b59435472616e736d69747465723a20696e64657820697320696e76616c6964600082015250565b7f4b59435472616e736d69747465723a207265717565737465725369676e61747560008201527f726520697320696e76616c696400000000000000000000000000000000000000602082015250565b7f4b59435472616e736d69747465723a206669656c64732061726520696e636f7260008201527f7265637400000000000000000000000000000000000000000000000000000000602082015250565b7f4549503731323a207369676e6174757265206c656e67746820697320696e636f60008201527f7272656374000000000000000000000000000000000000000000000000000000602082015250565b7f4b59435472616e736d69747465723a20696e76616c6964206669656c64000000600082015250565b7f4b59435472616e736d69747465723a2070726f76696465725369676e6174757260008201527f6520697320696e76616c69640000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f7564206861766560008201527f20726f6c65000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20636f6e747261637420697320646570726560008201527f6361746564000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c6520697320616c7265616479206760008201527f72616e7465640000000000000000000000000000000000000000000000000000602082015250565b7f4b59435472616e736d69747465723a20696e666f726d6174696f6e206973207660008201527f6572696669656400000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f756c642062652060008201527f61646d696e000000000000000000000000000000000000000000000000000000602082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a206e657720616e64206f6c642061646d696e60008201527f732073686f756c6420626520646966666572656e740000000000000000000000602082015250565b7f4b59435472616e736d69747465723a207369676e617475726520697320696e7660008201527f616c696400000000000000000000000000000000000000000000000000000000602082015250565b6135e381612fec565b81146135ee57600080fd5b50565b6135fa81612ffe565b811461360557600080fd5b50565b6136118161300a565b811461361c57600080fd5b50565b61362881613014565b811461363357600080fd5b50565b61363f81613048565b811461364a57600080fd5b5056fe50726f766964652861646472657373207265717565737465722c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c627974657333325b5d206669656c6473295265717565737428616464726573732070726f76696465722c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c627974657333325b5d206669656c647329a26469706673582212201186fc72b32c621b971937a0ba4cdb02ea1b40662c2681d163c946d63c6ff59864736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e74323536206e6574776f726b49642c6164647265737320766572696679696e67436f6e747261637429"

// DeployKYCTransmitter deploys a new Core contract, binding an instance of KYCTransmitter to it.
func DeployKYCTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend, vault common.Address, priceToken_ common.Address) (common.Address, *types.Transaction, *KYCTransmitter, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCTransmitterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KYCTransmitterBin), backend, vault, priceToken_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KYCTransmitter{KYCTransmitterCaller: KYCTransmitterCaller{contract: contract}, KYCTransmitterTransactor: KYCTransmitterTransactor{contract: contract}, KYCTransmitterFilterer: KYCTransmitterFilterer{contract: contract}}, nil
}

// KYCTransmitter is an auto generated Go binding around an Core contract.
type KYCTransmitter struct {
	KYCTransmitterCaller     // Read-only binding to the contract
	KYCTransmitterTransactor // Write-only binding to the contract
	KYCTransmitterFilterer   // Log filterer for contract events
}

// KYCTransmitterCaller is an auto generated read-only Go binding around an Core contract.
type KYCTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCTransmitterTransactor is an auto generated write-only Go binding around an Core contract.
type KYCTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCTransmitterFilterer is an auto generated log filtering Go binding around an Core contract events.
type KYCTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCTransmitterSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type KYCTransmitterSession struct {
	Contract     *KYCTransmitter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KYCTransmitterCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type KYCTransmitterCallerSession struct {
	Contract *KYCTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KYCTransmitterTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type KYCTransmitterTransactorSession struct {
	Contract     *KYCTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KYCTransmitterRaw is an auto generated low-level Go binding around an Core contract.
type KYCTransmitterRaw struct {
	Contract *KYCTransmitter // Generic contract binding to access the raw methods on
}

// KYCTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type KYCTransmitterCallerRaw struct {
	Contract *KYCTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// KYCTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type KYCTransmitterTransactorRaw struct {
	Contract *KYCTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKYCTransmitter creates a new instance of KYCTransmitter, bound to a specific deployed contract.
func NewKYCTransmitter(address common.Address, backend bind.ContractBackend) (*KYCTransmitter, error) {
	contract, err := bindKYCTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitter{KYCTransmitterCaller: KYCTransmitterCaller{contract: contract}, KYCTransmitterTransactor: KYCTransmitterTransactor{contract: contract}, KYCTransmitterFilterer: KYCTransmitterFilterer{contract: contract}}, nil
}

// NewKYCTransmitterCaller creates a new read-only instance of KYCTransmitter, bound to a specific deployed contract.
func NewKYCTransmitterCaller(address common.Address, caller bind.ContractCaller) (*KYCTransmitterCaller, error) {
	contract, err := bindKYCTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterCaller{contract: contract}, nil
}

// NewKYCTransmitterTransactor creates a new write-only instance of KYCTransmitter, bound to a specific deployed contract.
func NewKYCTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*KYCTransmitterTransactor, error) {
	contract, err := bindKYCTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterTransactor{contract: contract}, nil
}

// NewKYCTransmitterFilterer creates a new log filterer instance of KYCTransmitter, bound to a specific deployed contract.
func NewKYCTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*KYCTransmitterFilterer, error) {
	contract, err := bindKYCTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterFilterer{contract: contract}, nil
}

// bindKYCTransmitter binds a generic wrapper to an already deployed contract.
func bindKYCTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCTransmitter *KYCTransmitterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYCTransmitter.Contract.KYCTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCTransmitter *KYCTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.KYCTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCTransmitter *KYCTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.KYCTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCTransmitter *KYCTransmitterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYCTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCTransmitter *KYCTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCTransmitter *KYCTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCTransmitter *KYCTransmitterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KYCTransmitter.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCTransmitter *KYCTransmitterSession) Admin() (common.Address, error) {
	return _KYCTransmitter.Contract.Admin(&_KYCTransmitter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCTransmitter *KYCTransmitterCallerSession) Admin() (common.Address, error) {
	return _KYCTransmitter.Contract.Admin(&_KYCTransmitter.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCTransmitter *KYCTransmitterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCTransmitter.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCTransmitter *KYCTransmitterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KYCTransmitter.Contract.HasRole(&_KYCTransmitter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCTransmitter *KYCTransmitterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KYCTransmitter.Contract.HasRole(&_KYCTransmitter.CallOpts, role, account)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCTransmitter *KYCTransmitterCaller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCTransmitter.contract.Call(opts, out, "isDeprecated")
	return *ret0, err
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCTransmitter *KYCTransmitterSession) IsDeprecated() (bool, error) {
	return _KYCTransmitter.Contract.IsDeprecated(&_KYCTransmitter.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCTransmitter *KYCTransmitterCallerSession) IsDeprecated() (bool, error) {
	return _KYCTransmitter.Contract.IsDeprecated(&_KYCTransmitter.CallOpts)
}

// PriceAmount is a free data retrieval call binding the contract method 0x2b610a34.
//
// Solidity: function priceAmount(bytes32 field) view returns(uint256)
func (_KYCTransmitter *KYCTransmitterCaller) PriceAmount(opts *bind.CallOpts, field [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KYCTransmitter.contract.Call(opts, out, "priceAmount", field)
	return *ret0, err
}

// PriceAmount is a free data retrieval call binding the contract method 0x2b610a34.
//
// Solidity: function priceAmount(bytes32 field) view returns(uint256)
func (_KYCTransmitter *KYCTransmitterSession) PriceAmount(field [32]byte) (*big.Int, error) {
	return _KYCTransmitter.Contract.PriceAmount(&_KYCTransmitter.CallOpts, field)
}

// PriceAmount is a free data retrieval call binding the contract method 0x2b610a34.
//
// Solidity: function priceAmount(bytes32 field) view returns(uint256)
func (_KYCTransmitter *KYCTransmitterCallerSession) PriceAmount(field [32]byte) (*big.Int, error) {
	return _KYCTransmitter.Contract.PriceAmount(&_KYCTransmitter.CallOpts, field)
}

// PriceToken is a free data retrieval call binding the contract method 0xb9d319b6.
//
// Solidity: function priceToken() view returns(address)
func (_KYCTransmitter *KYCTransmitterCaller) PriceToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KYCTransmitter.contract.Call(opts, out, "priceToken")
	return *ret0, err
}

// PriceToken is a free data retrieval call binding the contract method 0xb9d319b6.
//
// Solidity: function priceToken() view returns(address)
func (_KYCTransmitter *KYCTransmitterSession) PriceToken() (common.Address, error) {
	return _KYCTransmitter.Contract.PriceToken(&_KYCTransmitter.CallOpts)
}

// PriceToken is a free data retrieval call binding the contract method 0xb9d319b6.
//
// Solidity: function priceToken() view returns(address)
func (_KYCTransmitter *KYCTransmitterCallerSession) PriceToken() (common.Address, error) {
	return _KYCTransmitter.Contract.PriceToken(&_KYCTransmitter.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) ChangeAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "changeAdmin", account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCTransmitter *KYCTransmitterSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangeAdmin(&_KYCTransmitter.TransactOpts, account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangeAdmin(&_KYCTransmitter.TransactOpts, account)
}

// ChangePriceAmount is a paid mutator transaction binding the contract method 0x344dba96.
//
// Solidity: function changePriceAmount(bytes32 field, uint256 amount) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) ChangePriceAmount(opts *bind.TransactOpts, field [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "changePriceAmount", field, amount)
}

// ChangePriceAmount is a paid mutator transaction binding the contract method 0x344dba96.
//
// Solidity: function changePriceAmount(bytes32 field, uint256 amount) returns()
func (_KYCTransmitter *KYCTransmitterSession) ChangePriceAmount(field [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangePriceAmount(&_KYCTransmitter.TransactOpts, field, amount)
}

// ChangePriceAmount is a paid mutator transaction binding the contract method 0x344dba96.
//
// Solidity: function changePriceAmount(bytes32 field, uint256 amount) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) ChangePriceAmount(field [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangePriceAmount(&_KYCTransmitter.TransactOpts, field, amount)
}

// ChangePriceToken is a paid mutator transaction binding the contract method 0x47d4c036.
//
// Solidity: function changePriceToken(address token) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) ChangePriceToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "changePriceToken", token)
}

// ChangePriceToken is a paid mutator transaction binding the contract method 0x47d4c036.
//
// Solidity: function changePriceToken(address token) returns()
func (_KYCTransmitter *KYCTransmitterSession) ChangePriceToken(token common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangePriceToken(&_KYCTransmitter.TransactOpts, token)
}

// ChangePriceToken is a paid mutator transaction binding the contract method 0x47d4c036.
//
// Solidity: function changePriceToken(address token) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) ChangePriceToken(token common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.ChangePriceToken(&_KYCTransmitter.TransactOpts, token)
}

// Confirm is a paid mutator transaction binding the contract method 0x946aba7b.
//
// Solidity: function confirm(bytes32 transmission, bytes signature) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) Confirm(opts *bind.TransactOpts, transmission [32]byte, signature []byte) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "confirm", transmission, signature)
}

// Confirm is a paid mutator transaction binding the contract method 0x946aba7b.
//
// Solidity: function confirm(bytes32 transmission, bytes signature) returns()
func (_KYCTransmitter *KYCTransmitterSession) Confirm(transmission [32]byte, signature []byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Confirm(&_KYCTransmitter.TransactOpts, transmission, signature)
}

// Confirm is a paid mutator transaction binding the contract method 0x946aba7b.
//
// Solidity: function confirm(bytes32 transmission, bytes signature) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) Confirm(transmission [32]byte, signature []byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Confirm(&_KYCTransmitter.TransactOpts, transmission, signature)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCTransmitter *KYCTransmitterTransactor) DeprecateContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "deprecateContract")
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCTransmitter *KYCTransmitterSession) DeprecateContract() (*types.Transaction, error) {
	return _KYCTransmitter.Contract.DeprecateContract(&_KYCTransmitter.TransactOpts)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) DeprecateContract() (*types.Transaction, error) {
	return _KYCTransmitter.Contract.DeprecateContract(&_KYCTransmitter.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.GrantRole(&_KYCTransmitter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.GrantRole(&_KYCTransmitter.TransactOpts, role, account)
}

// Initiate is a paid mutator transaction binding the contract method 0x0d06de54.
//
// Solidity: function initiate(Cheque providerCheque, Cheque requesterCheque, bytes providerSignature, bytes requesterSignature, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) Initiate(opts *bind.TransactOpts, providerCheque Cheque, requesterCheque Cheque, providerSignature []byte, requesterSignature []byte, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "initiate", providerCheque, requesterCheque, providerSignature, requesterSignature, fields)
}

// Initiate is a paid mutator transaction binding the contract method 0x0d06de54.
//
// Solidity: function initiate(Cheque providerCheque, Cheque requesterCheque, bytes providerSignature, bytes requesterSignature, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterSession) Initiate(providerCheque Cheque, requesterCheque Cheque, providerSignature []byte, requesterSignature []byte, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Initiate(&_KYCTransmitter.TransactOpts, providerCheque, requesterCheque, providerSignature, requesterSignature, fields)
}

// Initiate is a paid mutator transaction binding the contract method 0x0d06de54.
//
// Solidity: function initiate(Cheque providerCheque, Cheque requesterCheque, bytes providerSignature, bytes requesterSignature, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) Initiate(providerCheque Cheque, requesterCheque Cheque, providerSignature []byte, requesterSignature []byte, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Initiate(&_KYCTransmitter.TransactOpts, providerCheque, requesterCheque, providerSignature, requesterSignature, fields)
}

// Invalidate is a paid mutator transaction binding the contract method 0x84e47966.
//
// Solidity: function invalidate(bytes32 transmission, uint256 index, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) Invalidate(opts *bind.TransactOpts, transmission [32]byte, index *big.Int, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "invalidate", transmission, index, fields)
}

// Invalidate is a paid mutator transaction binding the contract method 0x84e47966.
//
// Solidity: function invalidate(bytes32 transmission, uint256 index, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterSession) Invalidate(transmission [32]byte, index *big.Int, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Invalidate(&_KYCTransmitter.TransactOpts, transmission, index, fields)
}

// Invalidate is a paid mutator transaction binding the contract method 0x84e47966.
//
// Solidity: function invalidate(bytes32 transmission, uint256 index, bytes32[] fields) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) Invalidate(transmission [32]byte, index *big.Int, fields [][32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.Invalidate(&_KYCTransmitter.TransactOpts, transmission, index, fields)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCTransmitter *KYCTransmitterSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.RenounceRole(&_KYCTransmitter.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.RenounceRole(&_KYCTransmitter.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.RevokeRole(&_KYCTransmitter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCTransmitter *KYCTransmitterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCTransmitter.Contract.RevokeRole(&_KYCTransmitter.TransactOpts, role, account)
}

// KYCTransmitterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the KYCTransmitter contract.
type KYCTransmitterAdminChangedIterator struct {
	Event *KYCTransmitterAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterAdminChanged)
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
		it.Event = new(KYCTransmitterAdminChanged)
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
func (it *KYCTransmitterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterAdminChanged represents a AdminChanged event raised by the KYCTransmitter contract.
type KYCTransmitterAdminChanged struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*KYCTransmitterAdminChangedIterator, error) {

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterAdminChangedIterator{contract: _KYCTransmitter.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *KYCTransmitterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterAdminChanged)
				if err := _KYCTransmitter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseAdminChanged(log types.Log) (*KYCTransmitterAdminChanged, error) {
	event := new(KYCTransmitterAdminChanged)
	if err := _KYCTransmitter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the KYCTransmitter contract.
type KYCTransmitterConfirmedIterator struct {
	Event *KYCTransmitterConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterConfirmed)
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
		it.Event = new(KYCTransmitterConfirmed)
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
func (it *KYCTransmitterConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterConfirmed represents a Confirmed event raised by the KYCTransmitter contract.
type KYCTransmitterConfirmed struct {
	Transmission [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xfe8a98fc20cafa18a1d9c00074771ee084073c363aabb2c34c1b2905f0b12020.
//
// Solidity: event Confirmed(bytes32 indexed transmission)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterConfirmed(opts *bind.FilterOpts, transmission [][32]byte) (*KYCTransmitterConfirmedIterator, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "Confirmed", transmissionRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterConfirmedIterator{contract: _KYCTransmitter.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xfe8a98fc20cafa18a1d9c00074771ee084073c363aabb2c34c1b2905f0b12020.
//
// Solidity: event Confirmed(bytes32 indexed transmission)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *KYCTransmitterConfirmed, transmission [][32]byte) (event.Subscription, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "Confirmed", transmissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterConfirmed)
				if err := _KYCTransmitter.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0xfe8a98fc20cafa18a1d9c00074771ee084073c363aabb2c34c1b2905f0b12020.
//
// Solidity: event Confirmed(bytes32 indexed transmission)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseConfirmed(log types.Log) (*KYCTransmitterConfirmed, error) {
	event := new(KYCTransmitterConfirmed)
	if err := _KYCTransmitter.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterContractDeprecatedIterator is returned from FilterContractDeprecated and is used to iterate over the raw logs and unpacked data for ContractDeprecated events raised by the KYCTransmitter contract.
type KYCTransmitterContractDeprecatedIterator struct {
	Event *KYCTransmitterContractDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterContractDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterContractDeprecated)
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
		it.Event = new(KYCTransmitterContractDeprecated)
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
func (it *KYCTransmitterContractDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterContractDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterContractDeprecated represents a ContractDeprecated event raised by the KYCTransmitter contract.
type KYCTransmitterContractDeprecated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractDeprecated is a free log retrieval operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCTransmitter *KYCTransmitterFilterer) FilterContractDeprecated(opts *bind.FilterOpts) (*KYCTransmitterContractDeprecatedIterator, error) {

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterContractDeprecatedIterator{contract: _KYCTransmitter.contract, event: "ContractDeprecated", logs: logs, sub: sub}, nil
}

// WatchContractDeprecated is a free log subscription operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCTransmitter *KYCTransmitterFilterer) WatchContractDeprecated(opts *bind.WatchOpts, sink chan<- *KYCTransmitterContractDeprecated) (event.Subscription, error) {

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterContractDeprecated)
				if err := _KYCTransmitter.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
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

// ParseContractDeprecated is a log parse operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCTransmitter *KYCTransmitterFilterer) ParseContractDeprecated(log types.Log) (*KYCTransmitterContractDeprecated, error) {
	event := new(KYCTransmitterContractDeprecated)
	if err := _KYCTransmitter.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the KYCTransmitter contract.
type KYCTransmitterInitiatedIterator struct {
	Event *KYCTransmitterInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterInitiated)
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
		it.Event = new(KYCTransmitterInitiated)
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
func (it *KYCTransmitterInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterInitiated represents a Initiated event raised by the KYCTransmitter contract.
type KYCTransmitterInitiated struct {
	Transmission [32]byte
	Provider     common.Address
	Requester    common.Address
	Fields       [][32]byte
	Amounts      []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0xd5e9cd7af10b71d0052610468eba244b812851c4e582629a8e52faa2484c56aa.
//
// Solidity: event Initiated(bytes32 indexed transmission, address indexed provider, address indexed requester, bytes32[] fields, uint256[] amounts)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterInitiated(opts *bind.FilterOpts, transmission [][32]byte, provider []common.Address, requester []common.Address) (*KYCTransmitterInitiatedIterator, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "Initiated", transmissionRule, providerRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterInitiatedIterator{contract: _KYCTransmitter.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0xd5e9cd7af10b71d0052610468eba244b812851c4e582629a8e52faa2484c56aa.
//
// Solidity: event Initiated(bytes32 indexed transmission, address indexed provider, address indexed requester, bytes32[] fields, uint256[] amounts)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *KYCTransmitterInitiated, transmission [][32]byte, provider []common.Address, requester []common.Address) (event.Subscription, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "Initiated", transmissionRule, providerRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterInitiated)
				if err := _KYCTransmitter.contract.UnpackLog(event, "Initiated", log); err != nil {
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

// ParseInitiated is a log parse operation binding the contract event 0xd5e9cd7af10b71d0052610468eba244b812851c4e582629a8e52faa2484c56aa.
//
// Solidity: event Initiated(bytes32 indexed transmission, address indexed provider, address indexed requester, bytes32[] fields, uint256[] amounts)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseInitiated(log types.Log) (*KYCTransmitterInitiated, error) {
	event := new(KYCTransmitterInitiated)
	if err := _KYCTransmitter.contract.UnpackLog(event, "Initiated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterInvalidatedIterator is returned from FilterInvalidated and is used to iterate over the raw logs and unpacked data for Invalidated events raised by the KYCTransmitter contract.
type KYCTransmitterInvalidatedIterator struct {
	Event *KYCTransmitterInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterInvalidated)
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
		it.Event = new(KYCTransmitterInvalidated)
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
func (it *KYCTransmitterInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterInvalidated represents a Invalidated event raised by the KYCTransmitter contract.
type KYCTransmitterInvalidated struct {
	Transmission [32]byte
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidated is a free log retrieval operation binding the contract event 0x99905deb70a82b96b7ce3822e300fb0846eac837005061d54a943f56e1df6bb0.
//
// Solidity: event Invalidated(bytes32 indexed transmission, uint256 index)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterInvalidated(opts *bind.FilterOpts, transmission [][32]byte) (*KYCTransmitterInvalidatedIterator, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "Invalidated", transmissionRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterInvalidatedIterator{contract: _KYCTransmitter.contract, event: "Invalidated", logs: logs, sub: sub}, nil
}

// WatchInvalidated is a free log subscription operation binding the contract event 0x99905deb70a82b96b7ce3822e300fb0846eac837005061d54a943f56e1df6bb0.
//
// Solidity: event Invalidated(bytes32 indexed transmission, uint256 index)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchInvalidated(opts *bind.WatchOpts, sink chan<- *KYCTransmitterInvalidated, transmission [][32]byte) (event.Subscription, error) {

	var transmissionRule []interface{}
	for _, transmissionItem := range transmission {
		transmissionRule = append(transmissionRule, transmissionItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "Invalidated", transmissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterInvalidated)
				if err := _KYCTransmitter.contract.UnpackLog(event, "Invalidated", log); err != nil {
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

// ParseInvalidated is a log parse operation binding the contract event 0x99905deb70a82b96b7ce3822e300fb0846eac837005061d54a943f56e1df6bb0.
//
// Solidity: event Invalidated(bytes32 indexed transmission, uint256 index)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseInvalidated(log types.Log) (*KYCTransmitterInvalidated, error) {
	event := new(KYCTransmitterInvalidated)
	if err := _KYCTransmitter.contract.UnpackLog(event, "Invalidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterPriceAmountChangedIterator is returned from FilterPriceAmountChanged and is used to iterate over the raw logs and unpacked data for PriceAmountChanged events raised by the KYCTransmitter contract.
type KYCTransmitterPriceAmountChangedIterator struct {
	Event *KYCTransmitterPriceAmountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterPriceAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterPriceAmountChanged)
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
		it.Event = new(KYCTransmitterPriceAmountChanged)
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
func (it *KYCTransmitterPriceAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterPriceAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterPriceAmountChanged represents a PriceAmountChanged event raised by the KYCTransmitter contract.
type KYCTransmitterPriceAmountChanged struct {
	Field  [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPriceAmountChanged is a free log retrieval operation binding the contract event 0xe0ef2e1c942634b7a75983cbbf5697092369f04f3ad3786c2466dd7f88e848f8.
//
// Solidity: event PriceAmountChanged(bytes32 indexed field, uint256 amount)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterPriceAmountChanged(opts *bind.FilterOpts, field [][32]byte) (*KYCTransmitterPriceAmountChangedIterator, error) {

	var fieldRule []interface{}
	for _, fieldItem := range field {
		fieldRule = append(fieldRule, fieldItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "PriceAmountChanged", fieldRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterPriceAmountChangedIterator{contract: _KYCTransmitter.contract, event: "PriceAmountChanged", logs: logs, sub: sub}, nil
}

// WatchPriceAmountChanged is a free log subscription operation binding the contract event 0xe0ef2e1c942634b7a75983cbbf5697092369f04f3ad3786c2466dd7f88e848f8.
//
// Solidity: event PriceAmountChanged(bytes32 indexed field, uint256 amount)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchPriceAmountChanged(opts *bind.WatchOpts, sink chan<- *KYCTransmitterPriceAmountChanged, field [][32]byte) (event.Subscription, error) {

	var fieldRule []interface{}
	for _, fieldItem := range field {
		fieldRule = append(fieldRule, fieldItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "PriceAmountChanged", fieldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterPriceAmountChanged)
				if err := _KYCTransmitter.contract.UnpackLog(event, "PriceAmountChanged", log); err != nil {
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

// ParsePriceAmountChanged is a log parse operation binding the contract event 0xe0ef2e1c942634b7a75983cbbf5697092369f04f3ad3786c2466dd7f88e848f8.
//
// Solidity: event PriceAmountChanged(bytes32 indexed field, uint256 amount)
func (_KYCTransmitter *KYCTransmitterFilterer) ParsePriceAmountChanged(log types.Log) (*KYCTransmitterPriceAmountChanged, error) {
	event := new(KYCTransmitterPriceAmountChanged)
	if err := _KYCTransmitter.contract.UnpackLog(event, "PriceAmountChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterPriceTokenChangedIterator is returned from FilterPriceTokenChanged and is used to iterate over the raw logs and unpacked data for PriceTokenChanged events raised by the KYCTransmitter contract.
type KYCTransmitterPriceTokenChangedIterator struct {
	Event *KYCTransmitterPriceTokenChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterPriceTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterPriceTokenChanged)
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
		it.Event = new(KYCTransmitterPriceTokenChanged)
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
func (it *KYCTransmitterPriceTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterPriceTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterPriceTokenChanged represents a PriceTokenChanged event raised by the KYCTransmitter contract.
type KYCTransmitterPriceTokenChanged struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceTokenChanged is a free log retrieval operation binding the contract event 0x1716a3f72bcf68266196e4932576023ececd691daf05d5eb758f69e181adc828.
//
// Solidity: event PriceTokenChanged(address token)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterPriceTokenChanged(opts *bind.FilterOpts) (*KYCTransmitterPriceTokenChangedIterator, error) {

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "PriceTokenChanged")
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterPriceTokenChangedIterator{contract: _KYCTransmitter.contract, event: "PriceTokenChanged", logs: logs, sub: sub}, nil
}

// WatchPriceTokenChanged is a free log subscription operation binding the contract event 0x1716a3f72bcf68266196e4932576023ececd691daf05d5eb758f69e181adc828.
//
// Solidity: event PriceTokenChanged(address token)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchPriceTokenChanged(opts *bind.WatchOpts, sink chan<- *KYCTransmitterPriceTokenChanged) (event.Subscription, error) {

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "PriceTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterPriceTokenChanged)
				if err := _KYCTransmitter.contract.UnpackLog(event, "PriceTokenChanged", log); err != nil {
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

// ParsePriceTokenChanged is a log parse operation binding the contract event 0x1716a3f72bcf68266196e4932576023ececd691daf05d5eb758f69e181adc828.
//
// Solidity: event PriceTokenChanged(address token)
func (_KYCTransmitter *KYCTransmitterFilterer) ParsePriceTokenChanged(log types.Log) (*KYCTransmitterPriceTokenChanged, error) {
	event := new(KYCTransmitterPriceTokenChanged)
	if err := _KYCTransmitter.contract.UnpackLog(event, "PriceTokenChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KYCTransmitter contract.
type KYCTransmitterRoleGrantedIterator struct {
	Event *KYCTransmitterRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterRoleGranted)
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
		it.Event = new(KYCTransmitterRoleGranted)
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
func (it *KYCTransmitterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterRoleGranted represents a RoleGranted event raised by the KYCTransmitter contract.
type KYCTransmitterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte) (*KYCTransmitterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterRoleGrantedIterator{contract: _KYCTransmitter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KYCTransmitterRoleGranted, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterRoleGranted)
				if err := _KYCTransmitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseRoleGranted(log types.Log) (*KYCTransmitterRoleGranted, error) {
	event := new(KYCTransmitterRoleGranted)
	if err := _KYCTransmitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCTransmitterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KYCTransmitter contract.
type KYCTransmitterRoleRevokedIterator struct {
	Event *KYCTransmitterRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCTransmitterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCTransmitterRoleRevoked)
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
		it.Event = new(KYCTransmitterRoleRevoked)
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
func (it *KYCTransmitterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCTransmitterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCTransmitterRoleRevoked represents a RoleRevoked event raised by the KYCTransmitter contract.
type KYCTransmitterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte) (*KYCTransmitterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCTransmitter.contract.FilterLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return &KYCTransmitterRoleRevokedIterator{contract: _KYCTransmitter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KYCTransmitterRoleRevoked, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCTransmitter.contract.WatchLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCTransmitterRoleRevoked)
				if err := _KYCTransmitter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCTransmitter *KYCTransmitterFilterer) ParseRoleRevoked(log types.Log) (*KYCTransmitterRoleRevoked, error) {
	event := new(KYCTransmitterRoleRevoked)
	if err := _KYCTransmitter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultABI is the input ABI used to generate the binding from.
const KYCVaultABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"submission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Invalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"submission\",\"type\":\"uint256\"}],\"name\":\"Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"submission\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field_\",\"type\":\"bytes32\"}],\"name\":\"field\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"required\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"}],\"internalType\":\"structField\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"field_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"submission_\",\"type\":\"uint256\"}],\"name\":\"fingerprint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"submissions_\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"}],\"name\":\"invalidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"field_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fingerprint_\",\"type\":\"bytes32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"field_\",\"type\":\"bytes32\"}],\"name\":\"isVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"submission_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"fields_\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"requireds\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"completes\",\"type\":\"bool[]\"}],\"name\":\"setFields\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"fields_\",\"type\":\"bytes32[]\"}],\"name\":\"setRoleFields\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"submission_\",\"type\":\"uint256\"}],\"name\":\"submission\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"removed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"invalidated\",\"type\":\"bool\"}],\"internalType\":\"structSubmission\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"fingerprints\",\"type\":\"bytes32[]\"}],\"internalType\":\"structInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KYCVaultBin is the compiled bytecode used for deploying new contracts.
var KYCVaultBin = "0x6101206040523480156200001257600080fd5b506040518060400160405280600881526020017f4b59435661756c740000000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250336000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff160217905550600060405180608001604052806054815260200162003cff6054913980519060200120905080608081815250506000838051906020012090508060a081815250506000838051906020012090508060c081815250504660e08181525050620001338383836200014660201b60201c565b610100818152505050505050506200025c565b6000838383463060405160200162000163959493929190620001b5565b6040516020818303038152906040528051906020012090509392505050565b6200018d8162000212565b82525050565b6200019e8162000226565b82525050565b620001af8162000252565b82525050565b600060a082019050620001cc600083018862000193565b620001db602083018762000193565b620001ea604083018662000193565b620001f96060830185620001a4565b62000208608083018462000182565b9695505050505050565b60006200021f8262000230565b9050919050565b6000819050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60805160a05160c05160e05161010051613a5e620002a1600039600061221d01526000612187015260006121f2015260006121d1015260006121b00152613a5e6000f3fe608060405234801561001057600080fd5b50600436106101155760003560e01c806399b4af83116100a2578063c9e1424811610071578063c9e14248146102da578063d569bc1c1461030a578063e05cab121461033a578063eadd855514610356578063eb8325fb1461037257610115565b806399b4af8314610254578063a3c6e14414610284578063a49b6f17146102b4578063bd8f8ebe146102d057610115565b80633e3c789e116100e95780633e3c789e146101b25780634e5a827c146101ce578063614b9887146101ea57806389a7fa411461020857806398662a031461022457610115565b80623ffbd41461011a5780631993cd0b146101365780631e5849281461016657806325e8d28214610196575b600080fd5b610134600480360381019061012f91906127fe565b610390565b005b610150600480360381019061014b9190612969565b610652565b60405161015d9190612f45565b60405180910390f35b610180600480360381019061017b91906129fd565b610666565b60405161018d9190612f23565b60405180910390f35b6101b060048036038101906101ab9190612969565b610c74565b005b6101cc60048036038101906101c79190612a7e565b610d16565b005b6101e860048036038101906101e39190612969565b610e99565b005b6101f2610f3b565b6040516101ff9190612f45565b60405180910390f35b610222600480360381019061021d91906126fb565b610f52565b005b61023e60048036038101906102399190612a55565b6110f9565b60405161024b9190613206565b60405180910390f35b61026e600480360381019061026991906127af565b6111cc565b60405161027b9190612f60565b60405180910390f35b61029e60048036038101906102999190612760565b61123d565b6040516102ab9190612f45565b60405180910390f35b6102ce60048036038101906102c991906128a2565b6115a4565b005b6102d861174a565b005b6102f460048036038101906102ef9190612724565b611877565b6040516103019190612f45565b60405180910390f35b610324600480360381019061031f9190612940565b611bdf565b60405161033191906131eb565b60405180910390f35b610354600480360381019061034f9190612940565b611c4c565b005b610370600480360381019061036b91906129a5565b611ca4565b005b61037a611f5b565b6040516103879190612f08565b60405180910390f35b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610424576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161041b9061312b565b60405180910390fd5b83839050868690501461046c576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610463906130eb565b60405180910390fd5b8181905086869050146104b4576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016104ab9061316b565b60405180910390fd5b60005b86869050811015610649578484828181106104fb577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020160208101906105109190612917565b6004600089898581811061054d577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135815260200190815260200160002060010160006101000a81548160ff0219169083151502179055508282828181106105b5577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020160208101906105ca9190612917565b60046000898985818110610607577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135815260200190815260200160002060010160016101000a81548160ff0219169083151502179055508080610641906134df565b9150506104b7565b50505050505050565b600061065e8383611f86565b905092915050565b6060836106738133611f86565b6106b2576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016106a9906130ab565b60405180910390fd5b60006007600087815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561071157602002820191906000526020600020905b8154815260200190600101908083116106fd575b5050505050905060008585905067ffffffffffffffff81111561075d577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561078b5781602001602082028036833780820191505090505b50905060005b86869050811015610c6657368787838181106107d6577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020028101906107e891906132cf565b90508080604001906107fa9190613221565b905084511461083e576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016108359061318b565b60405180910390fd5b600060036000815480929190610853906134df565b919050559050896005600083815260200190815260200160002060000181905550816000013560056000838152602001908152602001600020600101819055508160200160208101906108a691906126fb565b6005600083815260200190815260200160002060020160006101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff16021790555080848481518110610938577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050600860008b8152602001908152602001600020600083602001602081019061096c91906126fb565b75ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020016000909190919091505560005b8551811015610bf85760046000878381518110610a1a577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151815260200190815260200160002060010160009054906101000a900460ff161580610a9d57506000801b838060400190610a5c9190613221565b83818110610a93577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013514155b610adc576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610ad39061304b565b60405180910390fd5b828060400190610aec9190613221565b82818110610b23577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013560066000856020016020810190610b4191906126fb565b75ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000888481518110610bcb577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518152602001908152602001600020819055508080610bf0906134df565b9150506109d4565b50816020016020810190610c0c91906126fb565b75ffffffffffffffffffffffffffffffffffffffffffff16817f20e27110193713bb34c2cfe6c02bea3f232772da27307cd172608b93f120c0f760405160405180910390a350508080610c5e906134df565b915050610791565b508093505050509392505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610d08576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610cff9061312b565b60405180910390fd5b610d128282611ff2565b5050565b60006040518060400160405280601a81526020017f52656d6f76652875696e74323536207375626d697373696f6e290000000000008152508051906020012083604051602001610d67929190612ffe565b6040516020818303038152906040528051906020012090506000610d8b828461213c565b90508075ffffffffffffffffffffffffffffffffffffffffffff166005600086815260200190815260200160002060020160009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1614610e37576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610e2e906131cb565b60405180910390fd5b60016005600086815260200190815260200160002060020160166101000a81548160ff021916908315150217905550837f88b48749e543509091cdea1cd21b2eb5f5b0701da6d96d5ebdebcb8e7e2ba2d360405160405180910390a250505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610f2d576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610f249061312b565b60405180910390fd5b610f3782826122c1565b5050565b6000600260009054906101000a900460ff16905090565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610fe6576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610fdd9061312b565b60405180910390fd5b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168175ffffffffffffffffffffffffffffffffffffffffffff16141561107b576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611072906131ab565b60405180910390fd5b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055507fc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10816040516110ee9190612f08565b60405180910390a150565b6111016123ed565b600560008381526020019081526020016000206040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020016002820160169054906101000a900460ff161515151581526020016002820160179054906101000a900460ff1615151515815250509050919050565b6000600660008575ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060008481526020019081526020016000205490509392505050565b6000806004600085815260200190815260200160002060010160019054906101000a900460ff16905060006004600086815260200190815260200160002060000154905060006008600083815260200190815260200160002060008875ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561132157602002820191906000526020600020905b81548152602001906001019080831161130d575b505050505090506000815190505b6000811115611594576000600560008460018561134c91906133ba565b81518110611383577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015181526020019081526020016000206040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020016002820160169054906101000a900460ff161515151581526020016002820160179054906101000a900460ff16151515158152505090506000816020015114158015611465575042816020015111155b80611471575080606001515b8061147d575080608001515b1561149c5784156114965760009550505050505061159d565b50611581565b6000600660008b75ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000856001866114f191906133ba565b81518110611528577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151815260200190815260200160002060008a815260200190815260200160002054905087811415611569576001965050505050505061159d565b851561157e576000965050505050505061159d565b50505b808061158c90613484565b91505061132f565b50600093505050505b9392505050565b6000801b6115b28133611f86565b6115f1576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016115e8906130ab565b60405180910390fd5b60005b858590508110156117425760016005600088888581811061163e577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135815260200190815260200160002060020160176101000a81548160ff0219169083151502179055508585828181106116a6577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201357f155802ffbae5e440b691153667756a9bc180d55051a11ae48557ceaf634d04a6858584818110611707577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020028101906117199190613278565b604051611727929190613027565b60405180910390a2808061173a906134df565b9150506115f4565b505050505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff16146117de576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016117d59061312b565b60405180910390fd5b600260009054906101000a900460ff161561182e576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611825906130cb565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba2560405160405180910390a1565b6000806004600084815260200190815260200160002060010160019054906101000a900460ff16905060006004600085815260200190815260200160002060000154905060006008600083815260200190815260200160002060008775ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561195b57602002820191906000526020600020905b815481526020019060010190808311611947575b505050505090506000815190505b6000811115611bd0576000600560008460018561198691906133ba565b815181106119bd577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015181526020019081526020016000206040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff1681526020016002820160169054906101000a900460ff161515151581526020016002820160179054906101000a900460ff16151515158152505090506000816020015114158015611a9f575042816020015111155b80611aab575080606001515b80611ab7575080608001515b15611ad6578415611ad057600095505050505050611bd9565b50611bbd565b6000600660008a75ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085600186611b2b91906133ba565b81518110611b62577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151815260200190815260200160002060008981526020019081526020016000205490506000801b8114611ba55760019650505050505050611bd9565b8515611bba5760009650505050505050611bd9565b50505b8080611bc890613484565b915050611969565b50600093505050505b92915050565b611be761243b565b60046000838152602001908152602001600020604051806060016040529081600082015481526020016001820160009054906101000a900460ff161515151581526020016001820160019054906101000a900460ff1615151515815250509050919050565b80611c578133611f86565b611c96576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611c8d906130ab565b60405180910390fd5b611ca08233611ff2565b5050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614611d38576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611d2f9061312b565b60405180910390fd5b600060076000858152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015611d9757602002820191906000526020600020905b815481526020019060010190808311611d83575b5050505050905060005b8151811015611e1b576000801b60046000848481518110611deb577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518152602001908152602001600020600001819055508080611e13906134df565b915050611da1565b5060005b83839050811015611f31576000801b60046000868685818110611e6b577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013581526020019081526020016000206000015414611ec4576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401611ebb9061314b565b60405180910390fd5b8460046000868685818110611f02577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201358152602001908152602001600020600001819055508080611f29906134df565b915050611e1f565b508282600760008781526020019081526020016000209190611f54929190612463565b5050505050565b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905090565b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6001600083815260200190815260200160002060008275ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612093576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161208a9061306b565b60405180910390fd5b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850826040516121309190612f08565b60405180910390a25050565b600060ab825114612182576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016121799061308b565b60405180910390fd5b6000467f00000000000000000000000000000000000000000000000000000000000000001461221b576122167f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006123b3565b61223d565b7f00000000000000000000000000000000000000000000000000000000000000005b905060008185604051602001612254929190612ed1565b604051602081830303815290604052805190602001209050600181856040516000815260200160405260405161228b929190612fce565b6020604051602081039080840390855afa1580156122ad573d6000803e3d6000fd5b505050602060405103519250505092915050565b6122cb8282611f86565b1561230b576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016123029061310b565b60405180910390fd5b600180600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb826040516123a79190612f08565b60405180910390a25050565b600083838346306040516020016123ce959493929190612f7b565b6040516020818303038152906040528051906020012090509392505050565b6040518060a001604052806000801916815260200160008152602001600075ffffffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016000151581525090565b6040518060600160405280600080191681526020016000151581526020016000151581525090565b82805482825590600052602060002090810192821561249f579160200282015b8281111561249e578235825591602001919060010190612483565b5b5090506124ac91906124b0565b5090565b5b808211156124c95760008160009055506001016124b1565b5090565b60006124e06124db84613318565b6132f3565b9050828152602081018484840111156124f857600080fd5b612503848285613442565b509392505050565b60008135905061251a816139a7565b92915050565b60008083601f84011261253257600080fd5b8235905067ffffffffffffffff81111561254b57600080fd5b60208301915083602082028301111561256357600080fd5b9250929050565b60008083601f84011261257c57600080fd5b8235905067ffffffffffffffff81111561259557600080fd5b6020830191508360208202830111156125ad57600080fd5b9250929050565b60008083601f8401126125c657600080fd5b8235905067ffffffffffffffff8111156125df57600080fd5b6020830191508360208202830111156125f757600080fd5b9250929050565b60008083601f84011261261057600080fd5b8235905067ffffffffffffffff81111561262957600080fd5b60208301915083602082028301111561264157600080fd5b9250929050565b60008083601f84011261265a57600080fd5b8235905067ffffffffffffffff81111561267357600080fd5b60208301915083602082028301111561268b57600080fd5b9250929050565b6000813590506126a1816139be565b92915050565b6000813590506126b6816139d5565b92915050565b600082601f8301126126cd57600080fd5b81356126dd8482602086016124cd565b91505092915050565b6000813590506126f5816139ec565b92915050565b60006020828403121561270d57600080fd5b600061271b8482850161250b565b91505092915050565b6000806040838503121561273757600080fd5b60006127458582860161250b565b9250506020612756858286016126a7565b9150509250929050565b60008060006060848603121561277557600080fd5b60006127838682870161250b565b9350506020612794868287016126a7565b92505060406127a5868287016126a7565b9150509250925092565b6000806000606084860312156127c457600080fd5b60006127d28682870161250b565b93505060206127e3868287016126a7565b92505060406127f4868287016126e6565b9150509250925092565b6000806000806000806060878903121561281757600080fd5b600087013567ffffffffffffffff81111561283157600080fd5b61283d89828a0161256a565b9650965050602087013567ffffffffffffffff81111561285c57600080fd5b61286889828a01612520565b9450945050604087013567ffffffffffffffff81111561288757600080fd5b61289389828a01612520565b92509250509295509295509295565b600080600080604085870312156128b857600080fd5b600085013567ffffffffffffffff8111156128d257600080fd5b6128de87828801612648565b9450945050602085013567ffffffffffffffff8111156128fd57600080fd5b612909878288016125b4565b925092505092959194509250565b60006020828403121561292957600080fd5b600061293784828501612692565b91505092915050565b60006020828403121561295257600080fd5b6000612960848285016126a7565b91505092915050565b6000806040838503121561297c57600080fd5b600061298a858286016126a7565b925050602061299b8582860161250b565b9150509250929050565b6000806000604084860312156129ba57600080fd5b60006129c8868287016126a7565b935050602084013567ffffffffffffffff8111156129e557600080fd5b6129f18682870161256a565b92509250509250925092565b600080600060408486031215612a1257600080fd5b6000612a20868287016126a7565b935050602084013567ffffffffffffffff811115612a3d57600080fd5b612a49868287016125fe565b92509250509250925092565b600060208284031215612a6757600080fd5b6000612a75848285016126e6565b91505092915050565b60008060408385031215612a9157600080fd5b6000612a9f858286016126e6565b925050602083013567ffffffffffffffff811115612abc57600080fd5b612ac8858286016126bc565b9150509250929050565b6000612ade8383612eb3565b60208301905092915050565b612af3816133ee565b82525050565b612b02816133ee565b82525050565b6000612b1382613359565b612b1d818561337c565b9350612b2883613349565b8060005b83811015612b59578151612b408882612ad2565b9750612b4b8361336f565b925050600181019050612b2c565b5085935050505092915050565b612b6f81613400565b82525050565b612b7e81613400565b82525050565b612b8d8161340c565b82525050565b612b9c8161340c565b82525050565b612bb3612bae8261340c565b613528565b82525050565b6000612bc482613364565b612bce818561338d565b9350612bde818560208601613451565b612be781613590565b840191505092915050565b6000612bfe838561339e565b9350612c0b838584613442565b612c1483613590565b840190509392505050565b6000612c2c60288361339e565b9150612c37826135a1565b604082019050919050565b6000612c4f60228361339e565b9150612c5a826135f0565b604082019050919050565b6000612c7260258361339e565b9150612c7d8261363f565b604082019050919050565b6000612c9560258361339e565b9150612ca08261368e565b604082019050919050565b6000612cb860258361339e565b9150612cc3826136dd565b604082019050919050565b6000612cdb60378361339e565b9150612ce68261372c565b604082019050919050565b6000612cfe60268361339e565b9150612d098261377b565b604082019050919050565b6000612d2160258361339e565b9150612d2c826137ca565b604082019050919050565b6000612d4460318361339e565b9150612d4f82613819565b604082019050919050565b6000612d676002836133af565b9150612d7282613868565b600282019050919050565b6000612d8a60378361339e565b9150612d9582613891565b604082019050919050565b6000612dad603e8361339e565b9150612db8826138e0565b604082019050919050565b6000612dd060358361339e565b9150612ddb8261392f565b604082019050919050565b6000612df3601e8361339e565b9150612dfe8261397e565b602082019050919050565b606082016000820151612e1f6000850182612b84565b506020820151612e326020850182612b66565b506040820151612e456040850182612b66565b50505050565b60a082016000820151612e616000850182612b84565b506020820151612e746020850182612eb3565b506040820151612e876040850182612aea565b506060820151612e9a6060850182612b66565b506080820151612ead6080850182612b66565b50505050565b612ebc81613438565b82525050565b612ecb81613438565b82525050565b6000612edc82612d5a565b9150612ee88285612ba2565b602082019150612ef88284612ba2565b6020820191508190509392505050565b6000602082019050612f1d6000830184612af9565b92915050565b60006020820190508181036000830152612f3d8184612b08565b905092915050565b6000602082019050612f5a6000830184612b75565b92915050565b6000602082019050612f756000830184612b93565b92915050565b600060a082019050612f906000830188612b93565b612f9d6020830187612b93565b612faa6040830186612b93565b612fb76060830185612ec2565b612fc46080830184612af9565b9695505050505050565b6000604082019050612fe36000830185612b93565b8181036020830152612ff58184612bb9565b90509392505050565b60006040820190506130136000830185612b93565b6130206020830184612ec2565b9392505050565b60006020820190508181036000830152613042818486612bf2565b90509392505050565b6000602082019050818103600083015261306481612c1f565b9050919050565b6000602082019050818103600083015261308481612c42565b9050919050565b600060208201905081810360008301526130a481612c65565b9050919050565b600060208201905081810360008301526130c481612c88565b9050919050565b600060208201905081810360008301526130e481612cab565b9050919050565b6000602082019050818103600083015261310481612cce565b9050919050565b6000602082019050818103600083015261312481612cf1565b9050919050565b6000602082019050818103600083015261314481612d14565b9050919050565b6000602082019050818103600083015261316481612d37565b9050919050565b6000602082019050818103600083015261318481612d7d565b9050919050565b600060208201905081810360008301526131a481612da0565b9050919050565b600060208201905081810360008301526131c481612dc3565b9050919050565b600060208201905081810360008301526131e481612de6565b9050919050565b60006060820190506132006000830184612e09565b92915050565b600060a08201905061321b6000830184612e4b565b92915050565b6000808335600160200384360303811261323a57600080fd5b80840192508235915067ffffffffffffffff82111561325857600080fd5b60208301925060208202360383131561327057600080fd5b509250929050565b6000808335600160200384360303811261329157600080fd5b80840192508235915067ffffffffffffffff8211156132af57600080fd5b6020830192506001820236038313156132c757600080fd5b509250929050565b6000823560016060038336030381126132e757600080fd5b80830191505092915050565b60006132fd61330e565b905061330982826134ae565b919050565b6000604051905090565b600067ffffffffffffffff82111561333357613332613561565b5b61333c82613590565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006133c582613438565b91506133d083613438565b9250828210156133e3576133e2613532565b5b828203905092915050565b60006133f982613416565b9050919050565b60008115159050919050565b6000819050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561346f578082015181840152602081019050613454565b8381111561347e576000848401525b50505050565b600061348f82613438565b915060008214156134a3576134a2613532565b5b600182039050919050565b6134b782613590565b810181811067ffffffffffffffff821117156134d6576134d5613561565b5b80604052505050565b60006134ea82613438565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561351d5761351c613532565b5b600182019050919050565b6000819050919050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4b59435661756c743a207265717569726564206669656c642063616e206e6f7460008201527f206265207a65726f000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c65206973206e6f74206772616e7460008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b7f4549503731323a207369676e6174757265206c656e67746820697320696e636f60008201527f7272656374000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f7564206861766560008201527f20726f6c65000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20636f6e747261637420697320646570726560008201527f6361746564000000000000000000000000000000000000000000000000000000602082015250565b7f4b59435661756c743a206669656c64735f20616e64207265717569726564732060008201527f73686f756c6420686176652073616d65206c656e677468000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c6520697320616c7265616479206760008201527f72616e7465640000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f756c642062652060008201527f61646d696e000000000000000000000000000000000000000000000000000000602082015250565b7f4b59435661756c743a206669656c64735f2073686f756c64206e6f742068617660008201527f652061206469666672656e7420726f6c65000000000000000000000000000000602082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f4b59435661756c743a206669656c64735f20616e6420636f6d706c657465732060008201527f73686f756c6420686176652073616d65206c656e677468000000000000000000602082015250565b7f4b59435661756c743a20726f6c65206669656c647320616e642066696e67657260008201527f7072696e74732073686f756c6420686176652073616d65206c656e6774680000602082015250565b7f416363657373436f6e74726f6c3a206e657720616e64206f6c642061646d696e60008201527f732073686f756c6420626520646966666572656e740000000000000000000000602082015250565b7f4b59435661756c743a207369676e617475726520697320696e76616c69640000600082015250565b6139b0816133ee565b81146139bb57600080fd5b50565b6139c781613400565b81146139d257600080fd5b50565b6139de8161340c565b81146139e957600080fd5b50565b6139f581613438565b8114613a0057600080fd5b5056fea26469706673582212204338f65e3d96630e615a8938bae7281d9c8b1f3e1758118a211ab780ff91a56d64736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e74323536206e6574776f726b49642c6164647265737320766572696679696e67436f6e747261637429"

// DeployKYCVault deploys a new Core contract, binding an instance of KYCVault to it.
func DeployKYCVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KYCVault, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCVaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KYCVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KYCVault{KYCVaultCaller: KYCVaultCaller{contract: contract}, KYCVaultTransactor: KYCVaultTransactor{contract: contract}, KYCVaultFilterer: KYCVaultFilterer{contract: contract}}, nil
}

// KYCVault is an auto generated Go binding around an Core contract.
type KYCVault struct {
	KYCVaultCaller     // Read-only binding to the contract
	KYCVaultTransactor // Write-only binding to the contract
	KYCVaultFilterer   // Log filterer for contract events
}

// KYCVaultCaller is an auto generated read-only Go binding around an Core contract.
type KYCVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCVaultTransactor is an auto generated write-only Go binding around an Core contract.
type KYCVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCVaultFilterer is an auto generated log filtering Go binding around an Core contract events.
type KYCVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCVaultSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type KYCVaultSession struct {
	Contract     *KYCVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KYCVaultCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type KYCVaultCallerSession struct {
	Contract *KYCVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KYCVaultTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type KYCVaultTransactorSession struct {
	Contract     *KYCVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KYCVaultRaw is an auto generated low-level Go binding around an Core contract.
type KYCVaultRaw struct {
	Contract *KYCVault // Generic contract binding to access the raw methods on
}

// KYCVaultCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type KYCVaultCallerRaw struct {
	Contract *KYCVaultCaller // Generic read-only contract binding to access the raw methods on
}

// KYCVaultTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type KYCVaultTransactorRaw struct {
	Contract *KYCVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKYCVault creates a new instance of KYCVault, bound to a specific deployed contract.
func NewKYCVault(address common.Address, backend bind.ContractBackend) (*KYCVault, error) {
	contract, err := bindKYCVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KYCVault{KYCVaultCaller: KYCVaultCaller{contract: contract}, KYCVaultTransactor: KYCVaultTransactor{contract: contract}, KYCVaultFilterer: KYCVaultFilterer{contract: contract}}, nil
}

// NewKYCVaultCaller creates a new read-only instance of KYCVault, bound to a specific deployed contract.
func NewKYCVaultCaller(address common.Address, caller bind.ContractCaller) (*KYCVaultCaller, error) {
	contract, err := bindKYCVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KYCVaultCaller{contract: contract}, nil
}

// NewKYCVaultTransactor creates a new write-only instance of KYCVault, bound to a specific deployed contract.
func NewKYCVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*KYCVaultTransactor, error) {
	contract, err := bindKYCVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KYCVaultTransactor{contract: contract}, nil
}

// NewKYCVaultFilterer creates a new log filterer instance of KYCVault, bound to a specific deployed contract.
func NewKYCVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*KYCVaultFilterer, error) {
	contract, err := bindKYCVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KYCVaultFilterer{contract: contract}, nil
}

// bindKYCVault binds a generic wrapper to an already deployed contract.
func bindKYCVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCVault *KYCVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYCVault.Contract.KYCVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCVault *KYCVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCVault.Contract.KYCVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCVault *KYCVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCVault.Contract.KYCVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCVault *KYCVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYCVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCVault *KYCVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCVault *KYCVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCVault.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCVault *KYCVaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCVault *KYCVaultSession) Admin() (common.Address, error) {
	return _KYCVault.Contract.Admin(&_KYCVault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_KYCVault *KYCVaultCallerSession) Admin() (common.Address, error) {
	return _KYCVault.Contract.Admin(&_KYCVault.CallOpts)
}

// Field is a free data retrieval call binding the contract method 0xd569bc1c.
//
// Solidity: function field(bytes32 field_) view returns(Field)
func (_KYCVault *KYCVaultCaller) Field(opts *bind.CallOpts, field_ [32]byte) (Field, error) {
	var (
		ret0 = new(Field)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "field", field_)
	return *ret0, err
}

// Field is a free data retrieval call binding the contract method 0xd569bc1c.
//
// Solidity: function field(bytes32 field_) view returns(Field)
func (_KYCVault *KYCVaultSession) Field(field_ [32]byte) (Field, error) {
	return _KYCVault.Contract.Field(&_KYCVault.CallOpts, field_)
}

// Field is a free data retrieval call binding the contract method 0xd569bc1c.
//
// Solidity: function field(bytes32 field_) view returns(Field)
func (_KYCVault *KYCVaultCallerSession) Field(field_ [32]byte) (Field, error) {
	return _KYCVault.Contract.Field(&_KYCVault.CallOpts, field_)
}

// Fingerprint is a free data retrieval call binding the contract method 0x99b4af83.
//
// Solidity: function fingerprint(address user, bytes32 field_, uint256 submission_) view returns(bytes32)
func (_KYCVault *KYCVaultCaller) Fingerprint(opts *bind.CallOpts, user common.Address, field_ [32]byte, submission_ *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "fingerprint", user, field_, submission_)
	return *ret0, err
}

// Fingerprint is a free data retrieval call binding the contract method 0x99b4af83.
//
// Solidity: function fingerprint(address user, bytes32 field_, uint256 submission_) view returns(bytes32)
func (_KYCVault *KYCVaultSession) Fingerprint(user common.Address, field_ [32]byte, submission_ *big.Int) ([32]byte, error) {
	return _KYCVault.Contract.Fingerprint(&_KYCVault.CallOpts, user, field_, submission_)
}

// Fingerprint is a free data retrieval call binding the contract method 0x99b4af83.
//
// Solidity: function fingerprint(address user, bytes32 field_, uint256 submission_) view returns(bytes32)
func (_KYCVault *KYCVaultCallerSession) Fingerprint(user common.Address, field_ [32]byte, submission_ *big.Int) ([32]byte, error) {
	return _KYCVault.Contract.Fingerprint(&_KYCVault.CallOpts, user, field_, submission_)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCVault *KYCVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCVault *KYCVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KYCVault.Contract.HasRole(&_KYCVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KYCVault *KYCVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KYCVault.Contract.HasRole(&_KYCVault.CallOpts, role, account)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCVault *KYCVaultCaller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "isDeprecated")
	return *ret0, err
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCVault *KYCVaultSession) IsDeprecated() (bool, error) {
	return _KYCVault.Contract.IsDeprecated(&_KYCVault.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_KYCVault *KYCVaultCallerSession) IsDeprecated() (bool, error) {
	return _KYCVault.Contract.IsDeprecated(&_KYCVault.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xa3c6e144.
//
// Solidity: function isValid(address user, bytes32 field_, bytes32 fingerprint_) view returns(bool)
func (_KYCVault *KYCVaultCaller) IsValid(opts *bind.CallOpts, user common.Address, field_ [32]byte, fingerprint_ [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "isValid", user, field_, fingerprint_)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0xa3c6e144.
//
// Solidity: function isValid(address user, bytes32 field_, bytes32 fingerprint_) view returns(bool)
func (_KYCVault *KYCVaultSession) IsValid(user common.Address, field_ [32]byte, fingerprint_ [32]byte) (bool, error) {
	return _KYCVault.Contract.IsValid(&_KYCVault.CallOpts, user, field_, fingerprint_)
}

// IsValid is a free data retrieval call binding the contract method 0xa3c6e144.
//
// Solidity: function isValid(address user, bytes32 field_, bytes32 fingerprint_) view returns(bool)
func (_KYCVault *KYCVaultCallerSession) IsValid(user common.Address, field_ [32]byte, fingerprint_ [32]byte) (bool, error) {
	return _KYCVault.Contract.IsValid(&_KYCVault.CallOpts, user, field_, fingerprint_)
}

// IsVerified is a free data retrieval call binding the contract method 0xc9e14248.
//
// Solidity: function isVerified(address user, bytes32 field_) view returns(bool)
func (_KYCVault *KYCVaultCaller) IsVerified(opts *bind.CallOpts, user common.Address, field_ [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "isVerified", user, field_)
	return *ret0, err
}

// IsVerified is a free data retrieval call binding the contract method 0xc9e14248.
//
// Solidity: function isVerified(address user, bytes32 field_) view returns(bool)
func (_KYCVault *KYCVaultSession) IsVerified(user common.Address, field_ [32]byte) (bool, error) {
	return _KYCVault.Contract.IsVerified(&_KYCVault.CallOpts, user, field_)
}

// IsVerified is a free data retrieval call binding the contract method 0xc9e14248.
//
// Solidity: function isVerified(address user, bytes32 field_) view returns(bool)
func (_KYCVault *KYCVaultCallerSession) IsVerified(user common.Address, field_ [32]byte) (bool, error) {
	return _KYCVault.Contract.IsVerified(&_KYCVault.CallOpts, user, field_)
}

// Submission is a free data retrieval call binding the contract method 0x98662a03.
//
// Solidity: function submission(uint256 submission_) view returns(Submission)
func (_KYCVault *KYCVaultCaller) Submission(opts *bind.CallOpts, submission_ *big.Int) (Submission, error) {
	var (
		ret0 = new(Submission)
	)
	out := ret0
	err := _KYCVault.contract.Call(opts, out, "submission", submission_)
	return *ret0, err
}

// Submission is a free data retrieval call binding the contract method 0x98662a03.
//
// Solidity: function submission(uint256 submission_) view returns(Submission)
func (_KYCVault *KYCVaultSession) Submission(submission_ *big.Int) (Submission, error) {
	return _KYCVault.Contract.Submission(&_KYCVault.CallOpts, submission_)
}

// Submission is a free data retrieval call binding the contract method 0x98662a03.
//
// Solidity: function submission(uint256 submission_) view returns(Submission)
func (_KYCVault *KYCVaultCallerSession) Submission(submission_ *big.Int) (Submission, error) {
	return _KYCVault.Contract.Submission(&_KYCVault.CallOpts, submission_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCVault *KYCVaultTransactor) ChangeAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "changeAdmin", account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCVault *KYCVaultSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.ChangeAdmin(&_KYCVault.TransactOpts, account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_KYCVault *KYCVaultTransactorSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.ChangeAdmin(&_KYCVault.TransactOpts, account)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCVault *KYCVaultTransactor) DeprecateContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "deprecateContract")
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCVault *KYCVaultSession) DeprecateContract() (*types.Transaction, error) {
	return _KYCVault.Contract.DeprecateContract(&_KYCVault.TransactOpts)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_KYCVault *KYCVaultTransactorSession) DeprecateContract() (*types.Transaction, error) {
	return _KYCVault.Contract.DeprecateContract(&_KYCVault.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.GrantRole(&_KYCVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.GrantRole(&_KYCVault.TransactOpts, role, account)
}

// Invalidate is a paid mutator transaction binding the contract method 0xa49b6f17.
//
// Solidity: function invalidate(uint256[] submissions_, string[] reasons) returns()
func (_KYCVault *KYCVaultTransactor) Invalidate(opts *bind.TransactOpts, submissions_ []*big.Int, reasons []string) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "invalidate", submissions_, reasons)
}

// Invalidate is a paid mutator transaction binding the contract method 0xa49b6f17.
//
// Solidity: function invalidate(uint256[] submissions_, string[] reasons) returns()
func (_KYCVault *KYCVaultSession) Invalidate(submissions_ []*big.Int, reasons []string) (*types.Transaction, error) {
	return _KYCVault.Contract.Invalidate(&_KYCVault.TransactOpts, submissions_, reasons)
}

// Invalidate is a paid mutator transaction binding the contract method 0xa49b6f17.
//
// Solidity: function invalidate(uint256[] submissions_, string[] reasons) returns()
func (_KYCVault *KYCVaultTransactorSession) Invalidate(submissions_ []*big.Int, reasons []string) (*types.Transaction, error) {
	return _KYCVault.Contract.Invalidate(&_KYCVault.TransactOpts, submissions_, reasons)
}

// Remove is a paid mutator transaction binding the contract method 0x3e3c789e.
//
// Solidity: function remove(uint256 submission_, bytes signature) returns()
func (_KYCVault *KYCVaultTransactor) Remove(opts *bind.TransactOpts, submission_ *big.Int, signature []byte) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "remove", submission_, signature)
}

// Remove is a paid mutator transaction binding the contract method 0x3e3c789e.
//
// Solidity: function remove(uint256 submission_, bytes signature) returns()
func (_KYCVault *KYCVaultSession) Remove(submission_ *big.Int, signature []byte) (*types.Transaction, error) {
	return _KYCVault.Contract.Remove(&_KYCVault.TransactOpts, submission_, signature)
}

// Remove is a paid mutator transaction binding the contract method 0x3e3c789e.
//
// Solidity: function remove(uint256 submission_, bytes signature) returns()
func (_KYCVault *KYCVaultTransactorSession) Remove(submission_ *big.Int, signature []byte) (*types.Transaction, error) {
	return _KYCVault.Contract.Remove(&_KYCVault.TransactOpts, submission_, signature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCVault *KYCVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCVault *KYCVaultSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _KYCVault.Contract.RenounceRole(&_KYCVault.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_KYCVault *KYCVaultTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _KYCVault.Contract.RenounceRole(&_KYCVault.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.RevokeRole(&_KYCVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KYCVault *KYCVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KYCVault.Contract.RevokeRole(&_KYCVault.TransactOpts, role, account)
}

// SetFields is a paid mutator transaction binding the contract method 0x003ffbd4.
//
// Solidity: function setFields(bytes32[] fields_, bool[] requireds, bool[] completes) returns()
func (_KYCVault *KYCVaultTransactor) SetFields(opts *bind.TransactOpts, fields_ [][32]byte, requireds []bool, completes []bool) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "setFields", fields_, requireds, completes)
}

// SetFields is a paid mutator transaction binding the contract method 0x003ffbd4.
//
// Solidity: function setFields(bytes32[] fields_, bool[] requireds, bool[] completes) returns()
func (_KYCVault *KYCVaultSession) SetFields(fields_ [][32]byte, requireds []bool, completes []bool) (*types.Transaction, error) {
	return _KYCVault.Contract.SetFields(&_KYCVault.TransactOpts, fields_, requireds, completes)
}

// SetFields is a paid mutator transaction binding the contract method 0x003ffbd4.
//
// Solidity: function setFields(bytes32[] fields_, bool[] requireds, bool[] completes) returns()
func (_KYCVault *KYCVaultTransactorSession) SetFields(fields_ [][32]byte, requireds []bool, completes []bool) (*types.Transaction, error) {
	return _KYCVault.Contract.SetFields(&_KYCVault.TransactOpts, fields_, requireds, completes)
}

// SetRoleFields is a paid mutator transaction binding the contract method 0xeadd8555.
//
// Solidity: function setRoleFields(bytes32 role, bytes32[] fields_) returns()
func (_KYCVault *KYCVaultTransactor) SetRoleFields(opts *bind.TransactOpts, role [32]byte, fields_ [][32]byte) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "setRoleFields", role, fields_)
}

// SetRoleFields is a paid mutator transaction binding the contract method 0xeadd8555.
//
// Solidity: function setRoleFields(bytes32 role, bytes32[] fields_) returns()
func (_KYCVault *KYCVaultSession) SetRoleFields(role [32]byte, fields_ [][32]byte) (*types.Transaction, error) {
	return _KYCVault.Contract.SetRoleFields(&_KYCVault.TransactOpts, role, fields_)
}

// SetRoleFields is a paid mutator transaction binding the contract method 0xeadd8555.
//
// Solidity: function setRoleFields(bytes32 role, bytes32[] fields_) returns()
func (_KYCVault *KYCVaultTransactorSession) SetRoleFields(role [32]byte, fields_ [][32]byte) (*types.Transaction, error) {
	return _KYCVault.Contract.SetRoleFields(&_KYCVault.TransactOpts, role, fields_)
}

// Submit is a paid mutator transaction binding the contract method 0x1e584928.
//
// Solidity: function submit(bytes32 role, []Info infos) returns(uint256[])
func (_KYCVault *KYCVaultTransactor) Submit(opts *bind.TransactOpts, role [32]byte, infos []Info) (*types.Transaction, error) {
	return _KYCVault.contract.Transact(opts, "submit", role, infos)
}

// Submit is a paid mutator transaction binding the contract method 0x1e584928.
//
// Solidity: function submit(bytes32 role, []Info infos) returns(uint256[])
func (_KYCVault *KYCVaultSession) Submit(role [32]byte, infos []Info) (*types.Transaction, error) {
	return _KYCVault.Contract.Submit(&_KYCVault.TransactOpts, role, infos)
}

// Submit is a paid mutator transaction binding the contract method 0x1e584928.
//
// Solidity: function submit(bytes32 role, []Info infos) returns(uint256[])
func (_KYCVault *KYCVaultTransactorSession) Submit(role [32]byte, infos []Info) (*types.Transaction, error) {
	return _KYCVault.Contract.Submit(&_KYCVault.TransactOpts, role, infos)
}

// KYCVaultAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the KYCVault contract.
type KYCVaultAdminChangedIterator struct {
	Event *KYCVaultAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultAdminChanged)
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
		it.Event = new(KYCVaultAdminChanged)
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
func (it *KYCVaultAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultAdminChanged represents a AdminChanged event raised by the KYCVault contract.
type KYCVaultAdminChanged struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCVault *KYCVaultFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*KYCVaultAdminChangedIterator, error) {

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &KYCVaultAdminChangedIterator{contract: _KYCVault.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCVault *KYCVaultFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *KYCVaultAdminChanged) (event.Subscription, error) {

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultAdminChanged)
				if err := _KYCVault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_KYCVault *KYCVaultFilterer) ParseAdminChanged(log types.Log) (*KYCVaultAdminChanged, error) {
	event := new(KYCVaultAdminChanged)
	if err := _KYCVault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultContractDeprecatedIterator is returned from FilterContractDeprecated and is used to iterate over the raw logs and unpacked data for ContractDeprecated events raised by the KYCVault contract.
type KYCVaultContractDeprecatedIterator struct {
	Event *KYCVaultContractDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultContractDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultContractDeprecated)
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
		it.Event = new(KYCVaultContractDeprecated)
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
func (it *KYCVaultContractDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultContractDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultContractDeprecated represents a ContractDeprecated event raised by the KYCVault contract.
type KYCVaultContractDeprecated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractDeprecated is a free log retrieval operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCVault *KYCVaultFilterer) FilterContractDeprecated(opts *bind.FilterOpts) (*KYCVaultContractDeprecatedIterator, error) {

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return &KYCVaultContractDeprecatedIterator{contract: _KYCVault.contract, event: "ContractDeprecated", logs: logs, sub: sub}, nil
}

// WatchContractDeprecated is a free log subscription operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCVault *KYCVaultFilterer) WatchContractDeprecated(opts *bind.WatchOpts, sink chan<- *KYCVaultContractDeprecated) (event.Subscription, error) {

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultContractDeprecated)
				if err := _KYCVault.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
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

// ParseContractDeprecated is a log parse operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_KYCVault *KYCVaultFilterer) ParseContractDeprecated(log types.Log) (*KYCVaultContractDeprecated, error) {
	event := new(KYCVaultContractDeprecated)
	if err := _KYCVault.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultInvalidatedIterator is returned from FilterInvalidated and is used to iterate over the raw logs and unpacked data for Invalidated events raised by the KYCVault contract.
type KYCVaultInvalidatedIterator struct {
	Event *KYCVaultInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultInvalidated)
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
		it.Event = new(KYCVaultInvalidated)
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
func (it *KYCVaultInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultInvalidated represents a Invalidated event raised by the KYCVault contract.
type KYCVaultInvalidated struct {
	Submission *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidated is a free log retrieval operation binding the contract event 0x155802ffbae5e440b691153667756a9bc180d55051a11ae48557ceaf634d04a6.
//
// Solidity: event Invalidated(uint256 indexed submission, string reason)
func (_KYCVault *KYCVaultFilterer) FilterInvalidated(opts *bind.FilterOpts, submission []*big.Int) (*KYCVaultInvalidatedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "Invalidated", submissionRule)
	if err != nil {
		return nil, err
	}
	return &KYCVaultInvalidatedIterator{contract: _KYCVault.contract, event: "Invalidated", logs: logs, sub: sub}, nil
}

// WatchInvalidated is a free log subscription operation binding the contract event 0x155802ffbae5e440b691153667756a9bc180d55051a11ae48557ceaf634d04a6.
//
// Solidity: event Invalidated(uint256 indexed submission, string reason)
func (_KYCVault *KYCVaultFilterer) WatchInvalidated(opts *bind.WatchOpts, sink chan<- *KYCVaultInvalidated, submission []*big.Int) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "Invalidated", submissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultInvalidated)
				if err := _KYCVault.contract.UnpackLog(event, "Invalidated", log); err != nil {
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

// ParseInvalidated is a log parse operation binding the contract event 0x155802ffbae5e440b691153667756a9bc180d55051a11ae48557ceaf634d04a6.
//
// Solidity: event Invalidated(uint256 indexed submission, string reason)
func (_KYCVault *KYCVaultFilterer) ParseInvalidated(log types.Log) (*KYCVaultInvalidated, error) {
	event := new(KYCVaultInvalidated)
	if err := _KYCVault.contract.UnpackLog(event, "Invalidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultRemovedIterator is returned from FilterRemoved and is used to iterate over the raw logs and unpacked data for Removed events raised by the KYCVault contract.
type KYCVaultRemovedIterator struct {
	Event *KYCVaultRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultRemoved)
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
		it.Event = new(KYCVaultRemoved)
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
func (it *KYCVaultRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultRemoved represents a Removed event raised by the KYCVault contract.
type KYCVaultRemoved struct {
	Submission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoved is a free log retrieval operation binding the contract event 0x88b48749e543509091cdea1cd21b2eb5f5b0701da6d96d5ebdebcb8e7e2ba2d3.
//
// Solidity: event Removed(uint256 indexed submission)
func (_KYCVault *KYCVaultFilterer) FilterRemoved(opts *bind.FilterOpts, submission []*big.Int) (*KYCVaultRemovedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "Removed", submissionRule)
	if err != nil {
		return nil, err
	}
	return &KYCVaultRemovedIterator{contract: _KYCVault.contract, event: "Removed", logs: logs, sub: sub}, nil
}

// WatchRemoved is a free log subscription operation binding the contract event 0x88b48749e543509091cdea1cd21b2eb5f5b0701da6d96d5ebdebcb8e7e2ba2d3.
//
// Solidity: event Removed(uint256 indexed submission)
func (_KYCVault *KYCVaultFilterer) WatchRemoved(opts *bind.WatchOpts, sink chan<- *KYCVaultRemoved, submission []*big.Int) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "Removed", submissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultRemoved)
				if err := _KYCVault.contract.UnpackLog(event, "Removed", log); err != nil {
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

// ParseRemoved is a log parse operation binding the contract event 0x88b48749e543509091cdea1cd21b2eb5f5b0701da6d96d5ebdebcb8e7e2ba2d3.
//
// Solidity: event Removed(uint256 indexed submission)
func (_KYCVault *KYCVaultFilterer) ParseRemoved(log types.Log) (*KYCVaultRemoved, error) {
	event := new(KYCVaultRemoved)
	if err := _KYCVault.contract.UnpackLog(event, "Removed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KYCVault contract.
type KYCVaultRoleGrantedIterator struct {
	Event *KYCVaultRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultRoleGranted)
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
		it.Event = new(KYCVaultRoleGranted)
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
func (it *KYCVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultRoleGranted represents a RoleGranted event raised by the KYCVault contract.
type KYCVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte) (*KYCVaultRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return &KYCVaultRoleGrantedIterator{contract: _KYCVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KYCVaultRoleGranted, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultRoleGranted)
				if err := _KYCVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) ParseRoleGranted(log types.Log) (*KYCVaultRoleGranted, error) {
	event := new(KYCVaultRoleGranted)
	if err := _KYCVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KYCVault contract.
type KYCVaultRoleRevokedIterator struct {
	Event *KYCVaultRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultRoleRevoked)
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
		it.Event = new(KYCVaultRoleRevoked)
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
func (it *KYCVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultRoleRevoked represents a RoleRevoked event raised by the KYCVault contract.
type KYCVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte) (*KYCVaultRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return &KYCVaultRoleRevokedIterator{contract: _KYCVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KYCVaultRoleRevoked, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultRoleRevoked)
				if err := _KYCVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_KYCVault *KYCVaultFilterer) ParseRoleRevoked(log types.Log) (*KYCVaultRoleRevoked, error) {
	event := new(KYCVaultRoleRevoked)
	if err := _KYCVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KYCVaultSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the KYCVault contract.
type KYCVaultSubmittedIterator struct {
	Event *KYCVaultSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KYCVaultSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCVaultSubmitted)
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
		it.Event = new(KYCVaultSubmitted)
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
func (it *KYCVaultSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCVaultSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCVaultSubmitted represents a Submitted event raised by the KYCVault contract.
type KYCVaultSubmitted struct {
	Submission *big.Int
	User       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x20e27110193713bb34c2cfe6c02bea3f232772da27307cd172608b93f120c0f7.
//
// Solidity: event Submitted(uint256 indexed submission, address indexed user)
func (_KYCVault *KYCVaultFilterer) FilterSubmitted(opts *bind.FilterOpts, submission []*big.Int, user []common.Address) (*KYCVaultSubmittedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KYCVault.contract.FilterLogs(opts, "Submitted", submissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &KYCVaultSubmittedIterator{contract: _KYCVault.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x20e27110193713bb34c2cfe6c02bea3f232772da27307cd172608b93f120c0f7.
//
// Solidity: event Submitted(uint256 indexed submission, address indexed user)
func (_KYCVault *KYCVaultFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *KYCVaultSubmitted, submission []*big.Int, user []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KYCVault.contract.WatchLogs(opts, "Submitted", submissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCVaultSubmitted)
				if err := _KYCVault.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x20e27110193713bb34c2cfe6c02bea3f232772da27307cd172608b93f120c0f7.
//
// Solidity: event Submitted(uint256 indexed submission, address indexed user)
func (_KYCVault *KYCVaultFilterer) ParseSubmitted(log types.Log) (*KYCVaultSubmitted, error) {
	event := new(KYCVaultSubmitted)
	if err := _KYCVault.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedABI is the input ABI used to generate the binding from.
const PriceFeedABI = "[{\"inputs\":[],\"name\":\"getAggregatedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceFeed is an auto generated Go binding around an Core contract.
type PriceFeed struct {
	PriceFeedCaller     // Read-only binding to the contract
	PriceFeedTransactor // Write-only binding to the contract
	PriceFeedFilterer   // Log filterer for contract events
}

// PriceFeedCaller is an auto generated read-only Go binding around an Core contract.
type PriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTransactor is an auto generated write-only Go binding around an Core contract.
type PriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedFilterer is an auto generated log filtering Go binding around an Core contract events.
type PriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type PriceFeedSession struct {
	Contract     *PriceFeed        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type PriceFeedCallerSession struct {
	Contract *PriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PriceFeedTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type PriceFeedTransactorSession struct {
	Contract     *PriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PriceFeedRaw is an auto generated low-level Go binding around an Core contract.
type PriceFeedRaw struct {
	Contract *PriceFeed // Generic contract binding to access the raw methods on
}

// PriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type PriceFeedCallerRaw struct {
	Contract *PriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type PriceFeedTransactorRaw struct {
	Contract *PriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeed creates a new instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeed(address common.Address, backend bind.ContractBackend) (*PriceFeed, error) {
	contract, err := bindPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeed{PriceFeedCaller: PriceFeedCaller{contract: contract}, PriceFeedTransactor: PriceFeedTransactor{contract: contract}, PriceFeedFilterer: PriceFeedFilterer{contract: contract}}, nil
}

// NewPriceFeedCaller creates a new read-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedCaller, error) {
	contract, err := bindPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCaller{contract: contract}, nil
}

// NewPriceFeedTransactor creates a new write-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedTransactor, error) {
	contract, err := bindPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTransactor{contract: contract}, nil
}

// NewPriceFeedFilterer creates a new log filterer instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedFilterer, error) {
	contract, err := bindPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedFilterer{contract: contract}, nil
}

// bindPriceFeed binds a generic wrapper to an already deployed contract.
func bindPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.PriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transact(opts, method, params...)
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() view returns(uint256, uint256, uint32)
func (_PriceFeed *PriceFeedCaller) GetAggregatedPrice(opts *bind.CallOpts) (*big.Int, *big.Int, uint32, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PriceFeed.contract.Call(opts, out, "getAggregatedPrice")
	return *ret0, *ret1, *ret2, err
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() view returns(uint256, uint256, uint32)
func (_PriceFeed *PriceFeedSession) GetAggregatedPrice() (*big.Int, *big.Int, uint32, error) {
	return _PriceFeed.Contract.GetAggregatedPrice(&_PriceFeed.CallOpts)
}

// GetAggregatedPrice is a free data retrieval call binding the contract method 0xd9c1c1c9.
//
// Solidity: function getAggregatedPrice() view returns(uint256, uint256, uint32)
func (_PriceFeed *PriceFeedCallerSession) GetAggregatedPrice() (*big.Int, *big.Int, uint32, error) {
	return _PriceFeed.Contract.GetAggregatedPrice(&_PriceFeed.CallOpts)
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"FieldSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"keys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"values\",\"type\":\"bytes[]\"}],\"name\":\"setBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405234801561001057600080fd5b50336000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff160217905550611650806100646000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806389a7fa411161007157806389a7fa41146101505780639ed01eec1461016c578063bd8f8ebe14610188578063e05cab1214610192578063eb8325fb146101ae578063ee1effc4146101cc576100a9565b80631993cd0b146100ae57806325e8d282146100de5780634e5a827c146100fa578063614b98871461011657806373041daf14610134575b600080fd5b6100c860048036038101906100c39190610e31565b6101fc565b6040516100d59190611059565b60405180910390f35b6100f860048036038101906100f39190610e31565b610210565b005b610114600480360381019061010f9190610e31565b6102b2565b005b61011e610354565b60405161012b9190611059565b60405180910390f35b61014e60048036038101906101499190610d93565b61036b565b005b61016a60048036038101906101659190610d6a565b61046e565b005b61018660048036038101906101819190610e6d565b610615565b005b6101906106c2565b005b6101ac60048036038101906101a79190610e08565b6107ef565b005b6101b6610847565b6040516101c3919061103e565b60405180910390f35b6101e660048036038101906101e19190610e08565b610872565b6040516101f39190611098565b60405180910390f35b60006102088383610917565b905092915050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff16146102a4576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161029b9061115a565b60405180910390fd5b6102ae8282610983565b5050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610346576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161033d9061115a565b60405180910390fd5b6103508282610acd565b5050565b6000600260009054906101000a900460ff16905090565b8181905084849050146103b3576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016103aa906110ba565b60405180910390fd5b60005b84849050811015610467576104548585838181106103fd577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013584848481811061043d577f4b1f2ce300000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200281019061044f919061119a565b610615565b808061045f906112e6565b9150506103b6565b5050505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610502576040517f4e401cbe0000000000000000000000000000000000000000000000000000000081526004016104f99061115a565b60405180910390fd5b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff168175ffffffffffffffffffffffffffffffffffffffffffff161415610597576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161058e9061117a565b60405180910390fd5b806000806101000a81548175ffffffffffffffffffffffffffffffffffffffffffff021916908375ffffffffffffffffffffffffffffffffffffffffffff1602179055507fc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be108160405161060a919061103e565b60405180910390a150565b826106208133610917565b61065f576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610656906110fa565b60405180910390fd5b8282600360008781526020019081526020016000209190610681929190610bbf565b50837f2589a5cc18e856f8e7707822800bcc52307e608c3bd90828975209b7bf07307e84846040516106b4929190611074565b60405180910390a250505050565b60008054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff163375ffffffffffffffffffffffffffffffffffffffffffff1614610756576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161074d9061115a565b60405180910390fd5b600260009054906101000a900460ff16156107a6576040517f4e401cbe00000000000000000000000000000000000000000000000000000000815260040161079d9061111a565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba2560405160405180910390a1565b806107fa8133610917565b610839576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610830906110fa565b60405180910390fd5b6108438233610983565b5050565b60008060009054906101000a900475ffffffffffffffffffffffffffffffffffffffffffff16905090565b6060600360008381526020019081526020016000208054610892906112b4565b80601f01602080910402602001604051908101604052809291908181526020018280546108be906112b4565b801561090b5780601f106108e05761010080835404028352916020019161090b565b820191906000526020600020905b8154815290600101906020018083116108ee57829003601f168201915b50505050509050919050565b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6001600083815260200190815260200160002060008275ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a24576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610a1b906110da565b60405180910390fd5b60006001600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf46685082604051610ac1919061103e565b60405180910390a25050565b610ad78282610917565b15610b17576040517f4e401cbe000000000000000000000000000000000000000000000000000000008152600401610b0e9061113a565b60405180910390fd5b600180600084815260200190815260200160002060008375ffffffffffffffffffffffffffffffffffffffffffff1675ffffffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817f136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb82604051610bb3919061103e565b60405180910390a25050565b828054610bcb906112b4565b90600052602060002090601f016020900481019282610bed5760008555610c34565b82601f10610c0657803560ff1916838001178555610c34565b82800160010185558215610c34579182015b82811115610c33578235825591602001919060010190610c18565b5b509050610c419190610c45565b5090565b5b80821115610c5e576000816000905550600101610c46565b5090565b600081359050610c71816115c7565b92915050565b60008083601f840112610c8957600080fd5b8235905067ffffffffffffffff811115610ca257600080fd5b602083019150836020820283011115610cba57600080fd5b9250929050565b60008083601f840112610cd357600080fd5b8235905067ffffffffffffffff811115610cec57600080fd5b602083019150836020820283011115610d0457600080fd5b9250929050565b600081359050610d1a816115de565b92915050565b60008083601f840112610d3257600080fd5b8235905067ffffffffffffffff811115610d4b57600080fd5b602083019150836001820283011115610d6357600080fd5b9250929050565b600060208284031215610d7c57600080fd5b6000610d8a84828501610c62565b91505092915050565b60008060008060408587031215610da957600080fd5b600085013567ffffffffffffffff811115610dc357600080fd5b610dcf87828801610c77565b9450945050602085013567ffffffffffffffff811115610dee57600080fd5b610dfa87828801610cc1565b925092505092959194509250565b600060208284031215610e1a57600080fd5b6000610e2884828501610d0b565b91505092915050565b60008060408385031215610e4457600080fd5b6000610e5285828601610d0b565b9250506020610e6385828601610c62565b9150509250929050565b600080600060408486031215610e8257600080fd5b6000610e9086828701610d0b565b935050602084013567ffffffffffffffff811115610ead57600080fd5b610eb986828701610d20565b92509250509250925092565b610ece8161121e565b82525050565b610edd81611230565b82525050565b6000610eef83856111fc565b9350610efc838584611272565b610f058361138d565b840190509392505050565b6000610f1b826111f1565b610f2581856111fc565b9350610f35818560208601611281565b610f3e8161138d565b840191505092915050565b6000610f5660318361120d565b9150610f618261139e565b604082019050919050565b6000610f7960228361120d565b9150610f84826113ed565b604082019050919050565b6000610f9c60258361120d565b9150610fa78261143c565b604082019050919050565b6000610fbf60258361120d565b9150610fca8261148b565b604082019050919050565b6000610fe260268361120d565b9150610fed826114da565b604082019050919050565b600061100560258361120d565b915061101082611529565b604082019050919050565b600061102860358361120d565b915061103382611578565b604082019050919050565b60006020820190506110536000830184610ec5565b92915050565b600060208201905061106e6000830184610ed4565b92915050565b6000602082019050818103600083015261108f818486610ee3565b90509392505050565b600060208201905081810360008301526110b28184610f10565b905092915050565b600060208201905081810360008301526110d381610f49565b9050919050565b600060208201905081810360008301526110f381610f6c565b9050919050565b6000602082019050818103600083015261111381610f8f565b9050919050565b6000602082019050818103600083015261113381610fb2565b9050919050565b6000602082019050818103600083015261115381610fd5565b9050919050565b6000602082019050818103600083015261117381610ff8565b9050919050565b600060208201905081810360008301526111938161101b565b9050919050565b600080833560016020038436030381126111b357600080fd5b80840192508235915067ffffffffffffffff8211156111d157600080fd5b6020830192506001820236038313156111e957600080fd5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061122982611246565b9050919050565b60008115159050919050565b6000819050919050565b600075ffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561129f578082015181840152602081019050611284565b838111156112ae576000848401525b50505050565b600060028204905060018216806112cc57607f821691505b602082108114156112e0576112df61135e565b5b50919050565b60006112f182611268565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156113245761132361132f565b5b600182019050919050565b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4b1f2ce300000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f52656769737472793a206b65797320616e642076616c7565732073686f756c6460008201527f20686176652073616d65206c656e677468000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c65206973206e6f74206772616e7460008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f7564206861766560008201527f20726f6c65000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20636f6e747261637420697320646570726560008201527f6361746564000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a20726f6c6520697320616c7265616479206760008201527f72616e7465640000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a2063616c6c65722073686f756c642062652060008201527f61646d696e000000000000000000000000000000000000000000000000000000602082015250565b7f416363657373436f6e74726f6c3a206e657720616e64206f6c642061646d696e60008201527f732073686f756c6420626520646966666572656e740000000000000000000000602082015250565b6115d08161121e565b81146115db57600080fd5b50565b6115e78161123c565b81146115f257600080fd5b5056fea264697066735822122070293eb03e13a9b8b7e4320c95475a33bceab8ab3e94c4c338d6ae14849cd7bb64736f6c637827302e382e342d646576656c6f702e323032322e382e32322b636f6d6d69742e37383961353965650058"

// DeployRegistry deploys a new Core contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Core contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Core contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Core contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Core contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Core contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistrySession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xeb8325fb.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCallerSession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xee1effc4.
//
// Solidity: function get(bytes32 key) view returns(bytes)
func (_Registry *RegistryCaller) Get(opts *bind.CallOpts, key [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "get", key)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xee1effc4.
//
// Solidity: function get(bytes32 key) view returns(bytes)
func (_Registry *RegistrySession) Get(key [32]byte) ([]byte, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0xee1effc4.
//
// Solidity: function get(bytes32 key) view returns(bytes)
func (_Registry *RegistryCallerSession) Get(key [32]byte) ([]byte, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, key)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x1993cd0b.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_Registry *RegistryCaller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isDeprecated")
	return *ret0, err
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_Registry *RegistrySession) IsDeprecated() (bool, error) {
	return _Registry.Contract.IsDeprecated(&_Registry.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0x614b9887.
//
// Solidity: function isDeprecated() view returns(bool)
func (_Registry *RegistryCallerSession) IsDeprecated() (bool, error) {
	return _Registry.Contract.IsDeprecated(&_Registry.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_Registry *RegistryTransactor) ChangeAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeAdmin", account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_Registry *RegistrySession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeAdmin(&_Registry.TransactOpts, account)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x89a7fa41.
//
// Solidity: function changeAdmin(address account) returns()
func (_Registry *RegistryTransactorSession) ChangeAdmin(account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeAdmin(&_Registry.TransactOpts, account)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_Registry *RegistryTransactor) DeprecateContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deprecateContract")
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_Registry *RegistrySession) DeprecateContract() (*types.Transaction, error) {
	return _Registry.Contract.DeprecateContract(&_Registry.TransactOpts)
}

// DeprecateContract is a paid mutator transaction binding the contract method 0xbd8f8ebe.
//
// Solidity: function deprecateContract() returns()
func (_Registry *RegistryTransactorSession) DeprecateContract() (*types.Transaction, error) {
	return _Registry.Contract.DeprecateContract(&_Registry.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x4e5a827c.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_Registry *RegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_Registry *RegistrySession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0xe05cab12.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_Registry *RegistryTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x25e8d282.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// Set is a paid mutator transaction binding the contract method 0x9ed01eec.
//
// Solidity: function set(bytes32 key, bytes value) returns()
func (_Registry *RegistryTransactor) Set(opts *bind.TransactOpts, key [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0x9ed01eec.
//
// Solidity: function set(bytes32 key, bytes value) returns()
func (_Registry *RegistrySession) Set(key [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.Set(&_Registry.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0x9ed01eec.
//
// Solidity: function set(bytes32 key, bytes value) returns()
func (_Registry *RegistryTransactorSession) Set(key [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.Set(&_Registry.TransactOpts, key, value)
}

// SetBatch is a paid mutator transaction binding the contract method 0x73041daf.
//
// Solidity: function setBatch(bytes32[] keys, bytes[] values) returns()
func (_Registry *RegistryTransactor) SetBatch(opts *bind.TransactOpts, keys [][32]byte, values [][]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setBatch", keys, values)
}

// SetBatch is a paid mutator transaction binding the contract method 0x73041daf.
//
// Solidity: function setBatch(bytes32[] keys, bytes[] values) returns()
func (_Registry *RegistrySession) SetBatch(keys [][32]byte, values [][]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetBatch(&_Registry.TransactOpts, keys, values)
}

// SetBatch is a paid mutator transaction binding the contract method 0x73041daf.
//
// Solidity: function setBatch(bytes32[] keys, bytes[] values) returns()
func (_Registry *RegistryTransactorSession) SetBatch(keys [][32]byte, values [][]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetBatch(&_Registry.TransactOpts, keys, values)
}

// RegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Registry contract.
type RegistryAdminChangedIterator struct {
	Event *RegistryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAdminChanged)
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
		it.Event = new(RegistryAdminChanged)
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
func (it *RegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAdminChanged represents a AdminChanged event raised by the Registry contract.
type RegistryAdminChanged struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_Registry *RegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RegistryAdminChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryAdminChangedIterator{contract: _Registry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_Registry *RegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xc7c7b915837e4a31027f768078230ec8983ef038333f50a3d18ec1321b80be10.
//
// Solidity: event AdminChanged(address account)
func (_Registry *RegistryFilterer) ParseAdminChanged(log types.Log) (*RegistryAdminChanged, error) {
	event := new(RegistryAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryContractDeprecatedIterator is returned from FilterContractDeprecated and is used to iterate over the raw logs and unpacked data for ContractDeprecated events raised by the Registry contract.
type RegistryContractDeprecatedIterator struct {
	Event *RegistryContractDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryContractDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractDeprecated)
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
		it.Event = new(RegistryContractDeprecated)
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
func (it *RegistryContractDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractDeprecated represents a ContractDeprecated event raised by the Registry contract.
type RegistryContractDeprecated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractDeprecated is a free log retrieval operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_Registry *RegistryFilterer) FilterContractDeprecated(opts *bind.FilterOpts) (*RegistryContractDeprecatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return &RegistryContractDeprecatedIterator{contract: _Registry.contract, event: "ContractDeprecated", logs: logs, sub: sub}, nil
}

// WatchContractDeprecated is a free log subscription operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_Registry *RegistryFilterer) WatchContractDeprecated(opts *bind.WatchOpts, sink chan<- *RegistryContractDeprecated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractDeprecated)
				if err := _Registry.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
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

// ParseContractDeprecated is a log parse operation binding the contract event 0x3754e377228db45261033bde7298b9852db8f9e1efaba4bf5de1ecaaeaedba25.
//
// Solidity: event ContractDeprecated()
func (_Registry *RegistryFilterer) ParseContractDeprecated(log types.Log) (*RegistryContractDeprecated, error) {
	event := new(RegistryContractDeprecated)
	if err := _Registry.contract.UnpackLog(event, "ContractDeprecated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryFieldSetIterator is returned from FilterFieldSet and is used to iterate over the raw logs and unpacked data for FieldSet events raised by the Registry contract.
type RegistryFieldSetIterator struct {
	Event *RegistryFieldSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryFieldSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryFieldSet)
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
		it.Event = new(RegistryFieldSet)
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
func (it *RegistryFieldSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryFieldSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryFieldSet represents a FieldSet event raised by the Registry contract.
type RegistryFieldSet struct {
	Key   [32]byte
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFieldSet is a free log retrieval operation binding the contract event 0x2589a5cc18e856f8e7707822800bcc52307e608c3bd90828975209b7bf07307e.
//
// Solidity: event FieldSet(bytes32 indexed key, bytes value)
func (_Registry *RegistryFilterer) FilterFieldSet(opts *bind.FilterOpts, key [][32]byte) (*RegistryFieldSetIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "FieldSet", keyRule)
	if err != nil {
		return nil, err
	}
	return &RegistryFieldSetIterator{contract: _Registry.contract, event: "FieldSet", logs: logs, sub: sub}, nil
}

// WatchFieldSet is a free log subscription operation binding the contract event 0x2589a5cc18e856f8e7707822800bcc52307e608c3bd90828975209b7bf07307e.
//
// Solidity: event FieldSet(bytes32 indexed key, bytes value)
func (_Registry *RegistryFilterer) WatchFieldSet(opts *bind.WatchOpts, sink chan<- *RegistryFieldSet, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "FieldSet", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryFieldSet)
				if err := _Registry.contract.UnpackLog(event, "FieldSet", log); err != nil {
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

// ParseFieldSet is a log parse operation binding the contract event 0x2589a5cc18e856f8e7707822800bcc52307e608c3bd90828975209b7bf07307e.
//
// Solidity: event FieldSet(bytes32 indexed key, bytes value)
func (_Registry *RegistryFilterer) ParseFieldSet(log types.Log) (*RegistryFieldSet, error) {
	event := new(RegistryFieldSet)
	if err := _Registry.contract.UnpackLog(event, "FieldSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registry contract.
type RegistryRoleGrantedIterator struct {
	Event *RegistryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleGranted)
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
		it.Event = new(RegistryRoleGranted)
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
func (it *RegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleGranted represents a RoleGranted event raised by the Registry contract.
type RegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte) (*RegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleGrantedIterator{contract: _Registry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryRoleGranted, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleGranted", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleGranted)
				if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x136e75292551fe1957c5963049430eafc1ac5fab4ae16525221d819421175cbb.
//
// Solidity: event RoleGranted(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) ParseRoleGranted(log types.Log) (*RegistryRoleGranted, error) {
	event := new(RegistryRoleGranted)
	if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registry contract.
type RegistryRoleRevokedIterator struct {
	Event *RegistryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleRevoked)
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
		it.Event = new(RegistryRoleRevoked)
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
func (it *RegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleRevoked represents a RoleRevoked event raised by the Registry contract.
type RegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte) (*RegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleRevokedIterator{contract: _Registry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryRoleRevoked, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleRevoked", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleRevoked)
				if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x99d44913f63483d074e21a111bb455a1935af29d46dc3d20cf47e8fbbf466850.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address account)
func (_Registry *RegistryFilterer) ParseRoleRevoked(log types.Log) (*RegistryRoleRevoked, error) {
	event := new(RegistryRoleRevoked)
	if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrapperTokenABI is the input ABI used to generate the binding from.
const WrapperTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WrapperToken is an auto generated Go binding around an Core contract.
type WrapperToken struct {
	WrapperTokenCaller     // Read-only binding to the contract
	WrapperTokenTransactor // Write-only binding to the contract
	WrapperTokenFilterer   // Log filterer for contract events
}

// WrapperTokenCaller is an auto generated read-only Go binding around an Core contract.
type WrapperTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperTokenTransactor is an auto generated write-only Go binding around an Core contract.
type WrapperTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperTokenFilterer is an auto generated log filtering Go binding around an Core contract events.
type WrapperTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperTokenSession is an auto generated Go binding around an Core contract,
// with pre-set call and transact options.
type WrapperTokenSession struct {
	Contract     *WrapperToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapperTokenCallerSession is an auto generated read-only Go binding around an Core contract,
// with pre-set call options.
type WrapperTokenCallerSession struct {
	Contract *WrapperTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WrapperTokenTransactorSession is an auto generated write-only Go binding around an Core contract,
// with pre-set transact options.
type WrapperTokenTransactorSession struct {
	Contract     *WrapperTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WrapperTokenRaw is an auto generated low-level Go binding around an Core contract.
type WrapperTokenRaw struct {
	Contract *WrapperToken // Generic contract binding to access the raw methods on
}

// WrapperTokenCallerRaw is an auto generated low-level read-only Go binding around an Core contract.
type WrapperTokenCallerRaw struct {
	Contract *WrapperTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WrapperTokenTransactorRaw is an auto generated low-level write-only Go binding around an Core contract.
type WrapperTokenTransactorRaw struct {
	Contract *WrapperTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapperToken creates a new instance of WrapperToken, bound to a specific deployed contract.
func NewWrapperToken(address common.Address, backend bind.ContractBackend) (*WrapperToken, error) {
	contract, err := bindWrapperToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapperToken{WrapperTokenCaller: WrapperTokenCaller{contract: contract}, WrapperTokenTransactor: WrapperTokenTransactor{contract: contract}, WrapperTokenFilterer: WrapperTokenFilterer{contract: contract}}, nil
}

// NewWrapperTokenCaller creates a new read-only instance of WrapperToken, bound to a specific deployed contract.
func NewWrapperTokenCaller(address common.Address, caller bind.ContractCaller) (*WrapperTokenCaller, error) {
	contract, err := bindWrapperToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenCaller{contract: contract}, nil
}

// NewWrapperTokenTransactor creates a new write-only instance of WrapperToken, bound to a specific deployed contract.
func NewWrapperTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapperTokenTransactor, error) {
	contract, err := bindWrapperToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenTransactor{contract: contract}, nil
}

// NewWrapperTokenFilterer creates a new log filterer instance of WrapperToken, bound to a specific deployed contract.
func NewWrapperTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapperTokenFilterer, error) {
	contract, err := bindWrapperToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenFilterer{contract: contract}, nil
}

// bindWrapperToken binds a generic wrapper to an already deployed contract.
func bindWrapperToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapperToken *WrapperTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WrapperToken.Contract.WrapperTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapperToken *WrapperTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperToken.Contract.WrapperTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapperToken *WrapperTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperToken.Contract.WrapperTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapperToken *WrapperTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WrapperToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapperToken *WrapperTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapperToken *WrapperTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrapperToken *WrapperTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrapperToken *WrapperTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrapperToken.Contract.Allowance(&_WrapperToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x0bf3a456.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrapperToken *WrapperTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrapperToken.Contract.Allowance(&_WrapperToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrapperToken *WrapperTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrapperToken *WrapperTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrapperToken.Contract.BalanceOf(&_WrapperToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x1d7976f3.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrapperToken *WrapperTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrapperToken.Contract.BalanceOf(&_WrapperToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_WrapperToken *WrapperTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_WrapperToken *WrapperTokenSession) Decimals() (uint8, error) {
	return _WrapperToken.Contract.Decimals(&_WrapperToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5d1fb5f9.
//
// Solidity: function decimals() view returns(uint8)
func (_WrapperToken *WrapperTokenCallerSession) Decimals() (uint8, error) {
	return _WrapperToken.Contract.Decimals(&_WrapperToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_WrapperToken *WrapperTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_WrapperToken *WrapperTokenSession) Name() (string, error) {
	return _WrapperToken.Contract.Name(&_WrapperToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x07ba2a17.
//
// Solidity: function name() view returns(string)
func (_WrapperToken *WrapperTokenCallerSession) Name() (string, error) {
	return _WrapperToken.Contract.Name(&_WrapperToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_WrapperToken *WrapperTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_WrapperToken *WrapperTokenSession) Symbol() (string, error) {
	return _WrapperToken.Contract.Symbol(&_WrapperToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x231782d8.
//
// Solidity: function symbol() view returns(string)
func (_WrapperToken *WrapperTokenCallerSession) Symbol() (string, error) {
	return _WrapperToken.Contract.Symbol(&_WrapperToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrapperToken *WrapperTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrapperToken *WrapperTokenSession) TotalSupply() (*big.Int, error) {
	return _WrapperToken.Contract.TotalSupply(&_WrapperToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x1f1881f8.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrapperToken *WrapperTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WrapperToken.Contract.TotalSupply(&_WrapperToken.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_WrapperToken *WrapperTokenCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WrapperToken.contract.Call(opts, out, "wrappedToken")
	return *ret0, err
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_WrapperToken *WrapperTokenSession) WrappedToken() (common.Address, error) {
	return _WrapperToken.Contract.WrappedToken(&_WrapperToken.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x65995479.
//
// Solidity: function wrappedToken() view returns(address)
func (_WrapperToken *WrapperTokenCallerSession) WrappedToken() (common.Address, error) {
	return _WrapperToken.Contract.WrappedToken(&_WrapperToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Approve(&_WrapperToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xa613914d.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Approve(&_WrapperToken.TransactOpts, spender, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_WrapperToken *WrapperTokenTransactor) Buy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.contract.Transact(opts, "buy", amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_WrapperToken *WrapperTokenSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Buy(&_WrapperToken.TransactOpts, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x8b782835.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_WrapperToken *WrapperTokenTransactorSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Buy(&_WrapperToken.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_WrapperToken *WrapperTokenTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_WrapperToken *WrapperTokenSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Sell(&_WrapperToken.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x3f378b76.
//
// Solidity: function sell(uint256 amount) returns()
func (_WrapperToken *WrapperTokenTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Sell(&_WrapperToken.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Transfer(&_WrapperToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x4b40e901.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.Transfer(&_WrapperToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.TransferFrom(&_WrapperToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x31f2e679.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrapperToken *WrapperTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperToken.Contract.TransferFrom(&_WrapperToken.TransactOpts, sender, recipient, amount)
}

// WrapperTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrapperToken contract.
type WrapperTokenApprovalIterator struct {
	Event *WrapperTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WrapperTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperTokenApproval)
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
		it.Event = new(WrapperTokenApproval)
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
func (it *WrapperTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperTokenApproval represents a Approval event raised by the WrapperToken contract.
type WrapperTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrapperTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrapperToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenApprovalIterator{contract: _WrapperToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrapperTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrapperToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperTokenApproval)
				if err := _WrapperToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xafa504e0962ad93dec232a2c88581b4028671c11f4571f9edec54fb75bd7293d.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) ParseApproval(log types.Log) (*WrapperTokenApproval, error) {
	event := new(WrapperTokenApproval)
	if err := _WrapperToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrapperTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the WrapperToken contract.
type WrapperTokenBurnIterator struct {
	Event *WrapperTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WrapperTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperTokenBurn)
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
		it.Event = new(WrapperTokenBurn)
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
func (it *WrapperTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperTokenBurn represents a Burn event raised by the WrapperToken contract.
type WrapperTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*WrapperTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrapperToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenBurnIterator{contract: _WrapperToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *WrapperTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrapperToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperTokenBurn)
				if err := _WrapperToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xe8b8d39c36d62a24ad15be1765997bc47ac67488aa4cf22adb74dcc66e1cd39e.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) ParseBurn(log types.Log) (*WrapperTokenBurn, error) {
	event := new(WrapperTokenBurn)
	if err := _WrapperToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrapperTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the WrapperToken contract.
type WrapperTokenMintIterator struct {
	Event *WrapperTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WrapperTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperTokenMint)
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
		it.Event = new(WrapperTokenMint)
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
func (it *WrapperTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperTokenMint represents a Mint event raised by the WrapperToken contract.
type WrapperTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*WrapperTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenMintIterator{contract: _WrapperToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *WrapperTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperTokenMint)
				if err := _WrapperToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x058574c413d7a0f77f8a2502ca21e13cb5b407c9816ebb3ead4354fb50fb1678.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) ParseMint(log types.Log) (*WrapperTokenMint, error) {
	event := new(WrapperTokenMint)
	if err := _WrapperToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrapperTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrapperToken contract.
type WrapperTokenTransferIterator struct {
	Event *WrapperTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  core.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WrapperTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperTokenTransfer)
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
		it.Event = new(WrapperTokenTransfer)
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
func (it *WrapperTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperTokenTransfer represents a Transfer event raised by the WrapperToken contract.
type WrapperTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrapperTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperTokenTransferIterator{contract: _WrapperToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrapperTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperTokenTransfer)
				if err := _WrapperToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrapperToken *WrapperTokenFilterer) ParseTransfer(log types.Log) (*WrapperTokenTransfer, error) {
	event := new(WrapperTokenTransfer)
	if err := _WrapperToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
