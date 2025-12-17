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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.

// IDiamondCutMetaData contains all meta data concerning the IDiamondCut contract.
var IDiamondCutMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"}]",
	ID:  "IDiamondCut",
}

// IDiamondCut is an auto generated Go binding around an Ethereum contract.
type IDiamondCut struct {
	abi abi.ABI
}

// NewIDiamondCut creates a new instance of IDiamondCut.
func NewIDiamondCut() *IDiamondCut {
	parsed, err := IDiamondCutMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IDiamondCut{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IDiamondCut) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (iDiamondCut *IDiamondCut) PackDiamondCut(diamondCut []IDiamondCutFacetCut, init common.Address, calldata []byte) []byte {
	enc, err := iDiamondCut.abi.Pack("diamondCut", diamondCut, init, calldata)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (iDiamondCut *IDiamondCut) TryPackDiamondCut(diamondCut []IDiamondCutFacetCut, init common.Address, calldata []byte) ([]byte, error) {
	return iDiamondCut.abi.Pack("diamondCut", diamondCut, init, calldata)
}

// IDiamondCutDiamondCut represents a DiamondCut event raised by the IDiamondCut contract.
type IDiamondCutDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IDiamondCutDiamondCutEventName = "DiamondCut"

// ContractEventName returns the user-defined event name.
func (IDiamondCutDiamondCut) ContractEventName() string {
	return IDiamondCutDiamondCutEventName
}

// UnpackDiamondCutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (iDiamondCut *IDiamondCut) UnpackDiamondCutEvent(log *types.Log) (*IDiamondCutDiamondCut, error) {
	event := "DiamondCut"
	if len(log.Topics) == 0 || log.Topics[0] != iDiamondCut.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IDiamondCutDiamondCut)
	if len(log.Data) > 0 {
		if err := iDiamondCut.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iDiamondCut.abi.Events[event].Inputs {
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
