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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.

// IDiamondLoupeMetaData contains all meta data concerning the IDiamondLoupe contract.
var IDiamondLoupeMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "IDiamondLoupe",
}

// IDiamondLoupe is an auto generated Go binding around an Ethereum contract.
type IDiamondLoupe struct {
	abi abi.ABI
}

// NewIDiamondLoupe creates a new instance of IDiamondLoupe.
func NewIDiamondLoupe() *IDiamondLoupe {
	parsed, err := IDiamondLoupeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IDiamondLoupe{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IDiamondLoupe) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (iDiamondLoupe *IDiamondLoupe) PackFacetAddress(functionSelector [4]byte) []byte {
	enc, err := iDiamondLoupe.abi.Pack("facetAddress", functionSelector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (iDiamondLoupe *IDiamondLoupe) TryPackFacetAddress(functionSelector [4]byte) ([]byte, error) {
	return iDiamondLoupe.abi.Pack("facetAddress", functionSelector)
}

// UnpackFacetAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (iDiamondLoupe *IDiamondLoupe) UnpackFacetAddress(data []byte) (common.Address, error) {
	out, err := iDiamondLoupe.abi.Unpack("facetAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (iDiamondLoupe *IDiamondLoupe) PackFacetAddresses() []byte {
	enc, err := iDiamondLoupe.abi.Pack("facetAddresses")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (iDiamondLoupe *IDiamondLoupe) TryPackFacetAddresses() ([]byte, error) {
	return iDiamondLoupe.abi.Pack("facetAddresses")
}

// UnpackFacetAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (iDiamondLoupe *IDiamondLoupe) UnpackFacetAddresses(data []byte) ([]common.Address, error) {
	out, err := iDiamondLoupe.abi.Unpack("facetAddresses", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (iDiamondLoupe *IDiamondLoupe) PackFacetFunctionSelectors(facet common.Address) []byte {
	enc, err := iDiamondLoupe.abi.Pack("facetFunctionSelectors", facet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (iDiamondLoupe *IDiamondLoupe) TryPackFacetFunctionSelectors(facet common.Address) ([]byte, error) {
	return iDiamondLoupe.abi.Pack("facetFunctionSelectors", facet)
}

// UnpackFacetFunctionSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (iDiamondLoupe *IDiamondLoupe) UnpackFacetFunctionSelectors(data []byte) ([][4]byte, error) {
	out, err := iDiamondLoupe.abi.Unpack("facetFunctionSelectors", data)
	if err != nil {
		return *new([][4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)
	return out0, nil
}

// PackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (iDiamondLoupe *IDiamondLoupe) PackFacets() []byte {
	enc, err := iDiamondLoupe.abi.Pack("facets")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (iDiamondLoupe *IDiamondLoupe) TryPackFacets() ([]byte, error) {
	return iDiamondLoupe.abi.Pack("facets")
}

// UnpackFacets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (iDiamondLoupe *IDiamondLoupe) UnpackFacets(data []byte) ([]IDiamondLoupeFacet, error) {
	out, err := iDiamondLoupe.abi.Unpack("facets", data)
	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}
	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)
	return out0, nil
}
