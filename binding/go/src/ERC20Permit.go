// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ERC20PermitMetaData contains all meta data concerning the ERC20Permit contract.
var ERC20PermitMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"}]",
	ID:  "ERC20Permit",
}

// ERC20Permit is an auto generated Go binding around an Ethereum contract.
type ERC20Permit struct {
	abi abi.ABI
}

// NewERC20Permit creates a new instance of ERC20Permit.
func NewERC20Permit() *ERC20Permit {
	parsed, err := ERC20PermitMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Permit{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Permit) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Permit *ERC20Permit) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Permit.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Permit *ERC20Permit) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return eRC20Permit.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Permit *ERC20Permit) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Permit.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Permit *ERC20Permit) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Permit.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Permit *ERC20Permit) TryPackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	return eRC20Permit.abi.Pack("allowance", owner, spender)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Permit *ERC20Permit) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Permit.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Permit.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) TryPackApprove(spender common.Address, value *big.Int) ([]byte, error) {
	return eRC20Permit.abi.Pack("approve", spender, value)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Permit.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Permit *ERC20Permit) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Permit.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Permit *ERC20Permit) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return eRC20Permit.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Permit *ERC20Permit) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Permit.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Permit *ERC20Permit) PackDecimals() []byte {
	enc, err := eRC20Permit.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Permit *ERC20Permit) TryPackDecimals() ([]byte, error) {
	return eRC20Permit.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Permit *ERC20Permit) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Permit.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Permit *ERC20Permit) PackEip712Domain() []byte {
	enc, err := eRC20Permit.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Permit *ERC20Permit) TryPackEip712Domain() ([]byte, error) {
	return eRC20Permit.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Permit *ERC20Permit) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Permit.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (eRC20Permit *ERC20Permit) PackName() []byte {
	enc, err := eRC20Permit.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (eRC20Permit *ERC20Permit) TryPackName() ([]byte, error) {
	return eRC20Permit.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Permit *ERC20Permit) UnpackName(data []byte) (string, error) {
	out, err := eRC20Permit.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Permit *ERC20Permit) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Permit.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Permit *ERC20Permit) TryPackNonces(owner common.Address) ([]byte, error) {
	return eRC20Permit.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Permit *ERC20Permit) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Permit.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20Permit *ERC20Permit) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Permit.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20Permit *ERC20Permit) TryPackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return eRC20Permit.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (eRC20Permit *ERC20Permit) PackSymbol() []byte {
	enc, err := eRC20Permit.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (eRC20Permit *ERC20Permit) TryPackSymbol() ([]byte, error) {
	return eRC20Permit.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Permit *ERC20Permit) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Permit.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Permit *ERC20Permit) PackTotalSupply() []byte {
	enc, err := eRC20Permit.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Permit *ERC20Permit) TryPackTotalSupply() ([]byte, error) {
	return eRC20Permit.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Permit *ERC20Permit) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Permit.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Permit.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) TryPackTransfer(to common.Address, value *big.Int) ([]byte, error) {
	return eRC20Permit.abi.Pack("transfer", to, value)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Permit.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Permit.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) TryPackTransferFrom(from common.Address, to common.Address, value *big.Int) ([]byte, error) {
	return eRC20Permit.abi.Pack("transferFrom", from, to, value)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Permit *ERC20Permit) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Permit.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// ERC20PermitApproval represents a Approval event raised by the ERC20Permit contract.
type ERC20PermitApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20PermitApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20PermitApproval) ContractEventName() string {
	return ERC20PermitApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Permit *ERC20Permit) UnpackApprovalEvent(log *types.Log) (*ERC20PermitApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != eRC20Permit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20PermitApproval)
	if len(log.Data) > 0 {
		if err := eRC20Permit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Permit.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20PermitEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Permit contract.
type ERC20PermitEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20PermitEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20PermitEIP712DomainChanged) ContractEventName() string {
	return ERC20PermitEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Permit *ERC20Permit) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20PermitEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if len(log.Topics) == 0 || log.Topics[0] != eRC20Permit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20PermitEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Permit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Permit.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20PermitTransfer represents a Transfer event raised by the ERC20Permit contract.
type ERC20PermitTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20PermitTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20PermitTransfer) ContractEventName() string {
	return ERC20PermitTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Permit *ERC20Permit) UnpackTransferEvent(log *types.Log) (*ERC20PermitTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != eRC20Permit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20PermitTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Permit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Permit.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eRC20Permit *ERC20Permit) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Permit.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20Permit.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20PermitECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Permit contract.
type ERC20PermitECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20PermitECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Permit *ERC20Permit) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20PermitECDSAInvalidSignature, error) {
	out := new(ERC20PermitECDSAInvalidSignature)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Permit contract.
type ERC20PermitECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20PermitECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Permit *ERC20Permit) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20PermitECDSAInvalidSignatureLength, error) {
	out := new(ERC20PermitECDSAInvalidSignatureLength)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Permit contract.
type ERC20PermitECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20PermitECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Permit *ERC20Permit) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20PermitECDSAInvalidSignatureS, error) {
	out := new(ERC20PermitECDSAInvalidSignatureS)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Permit contract.
type ERC20PermitERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20PermitERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Permit *ERC20Permit) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20PermitERC20InsufficientAllowance, error) {
	out := new(ERC20PermitERC20InsufficientAllowance)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Permit contract.
type ERC20PermitERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20PermitERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Permit *ERC20Permit) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20PermitERC20InsufficientBalance, error) {
	out := new(ERC20PermitERC20InsufficientBalance)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Permit contract.
type ERC20PermitERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20PermitERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Permit *ERC20Permit) UnpackERC20InvalidApproverError(raw []byte) (*ERC20PermitERC20InvalidApprover, error) {
	out := new(ERC20PermitERC20InvalidApprover)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Permit contract.
type ERC20PermitERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20PermitERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Permit *ERC20Permit) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20PermitERC20InvalidReceiver, error) {
	out := new(ERC20PermitERC20InvalidReceiver)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Permit contract.
type ERC20PermitERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20PermitERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Permit *ERC20Permit) UnpackERC20InvalidSenderError(raw []byte) (*ERC20PermitERC20InvalidSender, error) {
	out := new(ERC20PermitERC20InvalidSender)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Permit contract.
type ERC20PermitERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20PermitERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Permit *ERC20Permit) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20PermitERC20InvalidSpender, error) {
	out := new(ERC20PermitERC20InvalidSpender)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Permit contract.
type ERC20PermitERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20PermitERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Permit *ERC20Permit) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20PermitERC2612ExpiredSignature, error) {
	out := new(ERC20PermitERC2612ExpiredSignature)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Permit contract.
type ERC20PermitERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20PermitERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Permit *ERC20Permit) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20PermitERC2612InvalidSigner, error) {
	out := new(ERC20PermitERC2612InvalidSigner)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Permit contract.
type ERC20PermitInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20PermitInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Permit *ERC20Permit) UnpackInvalidAccountNonceError(raw []byte) (*ERC20PermitInvalidAccountNonce, error) {
	out := new(ERC20PermitInvalidAccountNonce)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitInvalidShortString represents a InvalidShortString error raised by the ERC20Permit contract.
type ERC20PermitInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20PermitInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20Permit *ERC20Permit) UnpackInvalidShortStringError(raw []byte) (*ERC20PermitInvalidShortString, error) {
	out := new(ERC20PermitInvalidShortString)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PermitStringTooLong represents a StringTooLong error raised by the ERC20Permit contract.
type ERC20PermitStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20PermitStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20Permit *ERC20Permit) UnpackStringTooLongError(raw []byte) (*ERC20PermitStringTooLong, error) {
	out := new(ERC20PermitStringTooLong)
	if err := eRC20Permit.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
