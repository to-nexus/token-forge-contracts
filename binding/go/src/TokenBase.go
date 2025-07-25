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

// TokenBaseMetaData contains all meta data concerning the TokenBase contract.
var TokenBaseMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "TokenBase",
}

// TokenBase is an auto generated Go binding around an Ethereum contract.
type TokenBase struct {
	abi abi.ABI
}

// NewTokenBase creates a new instance of TokenBase.
func NewTokenBase() *TokenBase {
	parsed, err := TokenBaseMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TokenBase{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TokenBase) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (tokenBase *TokenBase) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := tokenBase.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (tokenBase *TokenBase) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := tokenBase.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (tokenBase *TokenBase) PackForgeByIndex(index *big.Int) []byte {
	enc, err := tokenBase.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (tokenBase *TokenBase) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := tokenBase.abi.Unpack("forgeByIndex", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackForgeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (tokenBase *TokenBase) PackForgeCount() []byte {
	enc, err := tokenBase.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (tokenBase *TokenBase) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := tokenBase.abi.Unpack("forgeCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (tokenBase *TokenBase) PackForges() []byte {
	enc, err := tokenBase.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (tokenBase *TokenBase) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := tokenBase.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (tokenBase *TokenBase) PackIsForge(forge common.Address) []byte {
	enc, err := tokenBase.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (tokenBase *TokenBase) UnpackIsForge(data []byte) (bool, error) {
	out, err := tokenBase.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (tokenBase *TokenBase) PackOwner() []byte {
	enc, err := tokenBase.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (tokenBase *TokenBase) UnpackOwner(data []byte) (common.Address, error) {
	out, err := tokenBase.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenBase *TokenBase) PackProxiableUUID() []byte {
	enc, err := tokenBase.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenBase *TokenBase) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := tokenBase.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (tokenBase *TokenBase) PackRenounceOwnership() []byte {
	enc, err := tokenBase.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (tokenBase *TokenBase) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := tokenBase.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenBase *TokenBase) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := tokenBase.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenBase *TokenBase) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := tokenBase.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (tokenBase *TokenBase) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := tokenBase.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (tokenBase *TokenBase) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := tokenBase.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TokenBaseForgeAdded represents a ForgeAdded event raised by the TokenBase contract.
type TokenBaseForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const TokenBaseForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (TokenBaseForgeAdded) ContractEventName() string {
	return TokenBaseForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (tokenBase *TokenBase) UnpackForgeAddedEvent(log *types.Log) (*TokenBaseForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != tokenBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenBaseForgeAdded)
	if len(log.Data) > 0 {
		if err := tokenBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenBase.abi.Events[event].Inputs {
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

// TokenBaseForgeRemoved represents a ForgeRemoved event raised by the TokenBase contract.
type TokenBaseForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const TokenBaseForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (TokenBaseForgeRemoved) ContractEventName() string {
	return TokenBaseForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (tokenBase *TokenBase) UnpackForgeRemovedEvent(log *types.Log) (*TokenBaseForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != tokenBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenBaseForgeRemoved)
	if len(log.Data) > 0 {
		if err := tokenBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenBase.abi.Events[event].Inputs {
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

// TokenBaseInitialized represents a Initialized event raised by the TokenBase contract.
type TokenBaseInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenBaseInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TokenBaseInitialized) ContractEventName() string {
	return TokenBaseInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (tokenBase *TokenBase) UnpackInitializedEvent(log *types.Log) (*TokenBaseInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != tokenBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenBaseInitialized)
	if len(log.Data) > 0 {
		if err := tokenBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenBase.abi.Events[event].Inputs {
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

// TokenBaseOwnershipTransferred represents a OwnershipTransferred event raised by the TokenBase contract.
type TokenBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const TokenBaseOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (TokenBaseOwnershipTransferred) ContractEventName() string {
	return TokenBaseOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (tokenBase *TokenBase) UnpackOwnershipTransferredEvent(log *types.Log) (*TokenBaseOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != tokenBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenBaseOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := tokenBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenBase.abi.Events[event].Inputs {
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

// TokenBaseUpgraded represents a Upgraded event raised by the TokenBase contract.
type TokenBaseUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TokenBaseUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TokenBaseUpgraded) ContractEventName() string {
	return TokenBaseUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (tokenBase *TokenBase) UnpackUpgradedEvent(log *types.Log) (*TokenBaseUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != tokenBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenBaseUpgraded)
	if len(log.Data) > 0 {
		if err := tokenBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenBase.abi.Events[event].Inputs {
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
func (tokenBase *TokenBase) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return tokenBase.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return tokenBase.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return tokenBase.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return tokenBase.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return tokenBase.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return tokenBase.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return tokenBase.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return tokenBase.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return tokenBase.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return tokenBase.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return tokenBase.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenBase.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return tokenBase.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TokenBaseAddressEmptyCode represents a AddressEmptyCode error raised by the TokenBase contract.
type TokenBaseAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TokenBaseAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (tokenBase *TokenBase) UnpackAddressEmptyCodeError(raw []byte) (*TokenBaseAddressEmptyCode, error) {
	out := new(TokenBaseAddressEmptyCode)
	if err := tokenBase.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TokenBase contract.
type TokenBaseERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TokenBaseERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (tokenBase *TokenBase) UnpackERC1967InvalidImplementationError(raw []byte) (*TokenBaseERC1967InvalidImplementation, error) {
	out := new(TokenBaseERC1967InvalidImplementation)
	if err := tokenBase.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseERC1967NonPayable represents a ERC1967NonPayable error raised by the TokenBase contract.
type TokenBaseERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TokenBaseERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (tokenBase *TokenBase) UnpackERC1967NonPayableError(raw []byte) (*TokenBaseERC1967NonPayable, error) {
	out := new(TokenBaseERC1967NonPayable)
	if err := tokenBase.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseFailedCall represents a FailedCall error raised by the TokenBase contract.
type TokenBaseFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TokenBaseFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (tokenBase *TokenBase) UnpackFailedCallError(raw []byte) (*TokenBaseFailedCall, error) {
	out := new(TokenBaseFailedCall)
	if err := tokenBase.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseInvalidInitialization represents a InvalidInitialization error raised by the TokenBase contract.
type TokenBaseInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TokenBaseInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (tokenBase *TokenBase) UnpackInvalidInitializationError(raw []byte) (*TokenBaseInvalidInitialization, error) {
	out := new(TokenBaseInvalidInitialization)
	if err := tokenBase.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseNotInitializing represents a NotInitializing error raised by the TokenBase contract.
type TokenBaseNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TokenBaseNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (tokenBase *TokenBase) UnpackNotInitializingError(raw []byte) (*TokenBaseNotInitializing, error) {
	out := new(TokenBaseNotInitializing)
	if err := tokenBase.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the TokenBase contract.
type TokenBaseOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func TokenBaseOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (tokenBase *TokenBase) UnpackOwnableInvalidOwnerError(raw []byte) (*TokenBaseOwnableInvalidOwner, error) {
	out := new(TokenBaseOwnableInvalidOwner)
	if err := tokenBase.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the TokenBase contract.
type TokenBaseOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func TokenBaseOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (tokenBase *TokenBase) UnpackOwnableUnauthorizedAccountError(raw []byte) (*TokenBaseOwnableUnauthorizedAccount, error) {
	out := new(TokenBaseOwnableUnauthorizedAccount)
	if err := tokenBase.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseTokenBaseNullInput represents a TokenBase__NullInput error raised by the TokenBase contract.
type TokenBaseTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func TokenBaseTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (tokenBase *TokenBase) UnpackTokenBaseNullInputError(raw []byte) (*TokenBaseTokenBaseNullInput, error) {
	out := new(TokenBaseTokenBaseNullInput)
	if err := tokenBase.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the TokenBase contract.
type TokenBaseTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func TokenBaseTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (tokenBase *TokenBase) UnpackTokenBaseOnlyForgeError(raw []byte) (*TokenBaseTokenBaseOnlyForge, error) {
	out := new(TokenBaseTokenBaseOnlyForge)
	if err := tokenBase.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TokenBase contract.
type TokenBaseUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TokenBaseUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (tokenBase *TokenBase) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TokenBaseUUPSUnauthorizedCallContext, error) {
	out := new(TokenBaseUUPSUnauthorizedCallContext)
	if err := tokenBase.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenBaseUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TokenBase contract.
type TokenBaseUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TokenBaseUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (tokenBase *TokenBase) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TokenBaseUUPSUnsupportedProxiableUUID, error) {
	out := new(TokenBaseUUPSUnsupportedProxiableUUID)
	if err := tokenBase.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}