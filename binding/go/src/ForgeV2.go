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

// ForgeV2MetaData contains all meta data concerning the ForgeV2 contract.
var ForgeV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidPermitSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"BaseForge__InvalidFeeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC1155ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC721ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ForgeV2",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b614ccb806100d65f395ff3fe608060405234801561000f575f5ffd5b5060043610610184575f3560e01c80637ecebe00116100dd578063b55ba9c411610088578063bf6bdd9711610063578063bf6bdd9714610415578063db686a7814610428578063f23a6e611461043b575f5ffd5b8063b55ba9c4146103b7578063ba32f859146103ca578063bc197c81146103dd575f5ffd5b80638abb7071116100b85780638abb70711461037e578063939cefc114610391578063a84adb32146103a4575f5ffd5b80637ecebe00146102fc57806382a9bb381461035057806384b0196e14610363575f5ffd5b80633644e5151161013d57806343dc7f781161011857806343dc7f78146102c357806360cfd501146102d657806372033c8c146102e9575f5ffd5b80633644e5151461025357806338033a8b146102695780633a5381b51461027c575f5ffd5b80632af4c9d01161016d5780632af4c9d0146102185780632f2ddc1f1461022d5780633081f74614610240575f5ffd5b806301ffc9a714610188578063150b7a02146101b0575b5f5ffd5b61019b610196366004613dbb565b610473565b60405190151581526020015b60405180910390f35b6101e76101be366004613f2f565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016101a7565b61022b610226366004613f93565b61050b565b005b61022b61023b366004614044565b610530565b61022b61024e366004613f93565b61088b565b61025b610c47565b6040519081526020016101a7565b61022b610277366004614129565b610c55565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b61022b6102d136600461424b565b610fda565b61022b6102e436600461424b565b611371565b61022b6102f7366004613f93565b61171e565b61025b61030a366004614314565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b61022b61035e366004613f93565b6119f4565b61036b611a0e565b6040516101a797969594939291906143b3565b61022b61038c366004614449565b611b08565b61022b61039f3660046144a1565b611e02565b61022b6103b236600461457c565b611e3f565b61022b6103c53660046144a1565b61218c565b61022b6103d836600461461b565b6121b8565b6101e76103eb3660046146e9565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b61022b610423366004614449565b61253d565b61022b610436366004614798565b612835565b6101e761044936600461486c565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e000000000000000000000000000000000000000000000000000000000148061050557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b8261051581612bdf565b6105258989898989898989612c1f565b505050505050505050565b8261053a81612bdf565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f4c82fe77677b753d6e8e9d6fd479564b08b135757219f2044e54db1882de3aa6602082015273ffffffffffffffffffffffffffffffffffffffff8b1691810191909152606081018990526080810188905260a0810182905260c081018790529091505f9060e0016040516020818303038152906040528051906020012090505f61061482612f51565b90505f6106218288612f98565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106a5576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024015b60405180910390fd5b50508451602080870191909120604080517fe9f21759dec582500f25edee01bab124984bac6771cf3f3dcb232329555083478185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61071f82612f51565b905061072b8186612fc0565b50506040517f124d91e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990526044820188905289169063124d91e5906064015f604051808303815f87803b1580156107a1575f5ffd5b505af11580156107b3573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e1682840152606080830187905283518084039091018152608090920190925280519101206105259250600291506040805173ffffffffffffffffffffffffffffffffffffffff8e1660208201529081018b9052606081018a90528b905f906080015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261087792916020016148c0565b6040516020818303038152906040526130c6565b8261089581612bdf565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f9eff2bc106f16f11dab69de8c466346a0a369dd1d3a0878683947093830994fa60208083019190915273ffffffffffffffffffffffffffffffffffffffff8d811683850152606083018d90528b16608083015260a082018a905260c0820184905260e08083018a90528351808403909101815261010090920190925280519101209091505f61097482612f51565b90505f6109818288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a00576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517f340f6ac2b427cb47a7172eaa76879e09534311918b4cff11daea1028ca325aa88185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f610a7a82612f51565b9050610a868186612fc0565b50505f5f610a9589898c61316c565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8f8116600483015260248201839052929450909250908c16906340c10f19906044015f604051808303815f87803b158015610b09575f5ffd5b505af1158015610b1b573d5f5f3e3d5ffd5b50505050815f14610baa576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018490528c16906340c10f19906044015f604051808303815f87803b158015610b93575f5ffd5b505af1158015610ba5573d5f5f3e3d5ffd5b505050505b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018d9052918b1660e08201526101008101849052610c39915f918e90610120015b60405160208183030381529060405261320e565b505050505050505050505050565b5f610c5061328b565b905090565b84610c5f81612bdf565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f4e6479369d4c42395619be8224f096421cd6800a840e701655040f5a933c33cd602082015273ffffffffffffffffffffffffffffffffffffffff8c1691810191909152606081018a90526080810182905260a081018990529091505f9060c0016040516020818303038152906040528051906020012090505f610d3282612f51565b90505f610d74828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612f9892505050565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610df3576040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b5050505f7f42dec6f4478d32028294c542bb3ee6cb16fb5523fd8f13922dca479fe76475138b8888604051610e299291906148e2565b604051908190038120610e6793929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f610e8982612f51565b9050610eca8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612fc092505050565b50506040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018a90528a16906379cc6790906044015f604051808303815f87803b158015610f39575f5ffd5b505af1158015610f4b573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f168284015260608083018790528351808403909101815260809092019092528051910120610fce9250600191506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c90528c90606001610877565b50505050505050505050565b82610fe481612bdf565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7fae67447a4abe000721c47402f895d79b71458082b87e1a69d5837b6241fed4858a8a60405160200161106691906148f1565b604051602081830303815290604052805190602001208a60405160200161108d91906148f1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201208d518e8301209184019690965273ffffffffffffffffffffffffffffffffffffffff909416908201526060810191909152608081019290925260a082015260c0810183905260e08101879052610100016040516020818303038152906040528051906020012090505f61113582612f51565b90505f6111428288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146111c1576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517f811c2af50c1ad6d46b121ba5a43024be13e428134866be726da388c2357856a58185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f61123b82612f51565b90506112478186612fc0565b50506040517f1f7fdffa00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1690631f7fdffa906112a1908d908c908c908c90600401614926565b5f604051808303815f87803b1580156112b8575f5ffd5b505af11580156112ca573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f168284015260608083018790528351808403909101815260809092019092528051910120610fce9250600291508b60018e8d8d60405160200161133593929190614985565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610c2592916020016148c0565b8261137b81612bdf565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7f53b86f7b5453d9de3449a5a75167b5d95aaf0968acc8570e67d9cb44e4c0a09d8a8a6040516020016113fd91906148f1565b604051602081830303815290604052805190602001208a60405160200161142491906148f1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201208d518e8301209184019690965273ffffffffffffffffffffffffffffffffffffffff909416908201526060810191909152608081019290925260a082015260c0810183905260e08101879052610100016040516020818303038152906040528051906020012090505f6114cc82612f51565b90505f6114d98288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611558576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517fbe51d3aafe3852e9cbe2ce5b4d184540bcb68868fc9307bf8b45d1f170d6739f8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f6115d282612f51565b90506115de8186612fc0565b50506040517f2eb2c2d600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1690632eb2c2d69061163a9030908e908d908d908d906004016149c5565b5f604051808303815f87803b158015611651575f5ffd5b505af1158015611663573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f168284015260608083018790528351808403909101815260809092019092528051910120610fce9250600291508b60018e8d8d6040516020016116ce93929190614985565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261170a92916020016148c0565b604051602081830303815290604052613294565b8261172881612bdf565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517fab25e47ab9e13a956bdee490006a87ac9cbe78ae734a45c431fe3938cbd0858160208083019190915273ffffffffffffffffffffffffffffffffffffffff8d811683850152606083018d90528b16608083015260a082018a905260c0820184905260e08083018a90528351808403909101815261010090920190925280519101209091505f61180782612f51565b90505f6118148288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611893576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517f7ae019cd8a06bf0d82d3d0063c026c8cfe203e2686c8de51ae89d5645e022d618185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f61190d82612f51565b90506119198186612fc0565b50505f5f61192889898c61316c565b909250905061194e73ffffffffffffffffffffffffffffffffffffffff8c168d83613311565b81156119755761197573ffffffffffffffffffffffffffffffffffffffff8c168a84613311565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018d9052918b1660e08201526101008101849052610c39915f918e906101200161170a565b826119fe81612bdf565b6105258989898989898989613397565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015611a4c57506001810154155b611ab2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161069c565b611aba613670565b611ac2613743565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b82611b1281612bdf565b73ffffffffffffffffffffffffffffffffffffffff8981165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020908152604080832080546001810190915589518a84012082517f3991a11da0d98ebde4668c7c2716088d93b9497b88d2b4cda98f19cfc97acd5281860152958e1686840152606086018d9052608086018c905260a086015260c0850181905260e08086018a9052825180870390910181526101009095019091528351939091019290922090611bdf82612f51565b90505f611bec8288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611c6b576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517fadcb1a40c0ace98e86ce9038ac7bb2e74cec43bdaa9e5727e2ea08c07a1f36dd8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f611ce582612f51565b9050611cf18186612fc0565b50506040517ff242432a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063f242432a90611d4d9030908e908d908d908d90600401614a41565b5f604051808303815f87803b158015611d64575f5ffd5b505af1158015611d76573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f168284015260608083018790528351808403909101815260809092019092528051910120610fce9250600291506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c9052606081018b90528c905f906080016116ce565b83611e0c81612bdf565b8989898785611e1e8486858585613794565b611e2e8f8f8f8f8f8f8f8f612c1f565b505050505050505050505050505050565b82611e4981612bdf565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7fda6dbf9c08df397da92be8fda42638fad009adf2a495b08f3526089ad02e7dd18989604051602001611ecb91906148f1565b6040516020818303038152906040528051906020012089604051602001611ef291906148f1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019590955273ffffffffffffffffffffffffffffffffffffffff909316928101929092526060820152608081019190915260a0810183905260c0810187905260e0016040516020818303038152906040528051906020012090505f611f8e82612f51565b90505f611f9b8288612f98565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461201a576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260240161069c565b50508451602080870191909120604080517fa221f8b9ccdf4675afa9893d3a99032eb17f5b0c291026fd55d97bb9993fa4e38185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61209482612f51565b90506120a08186612fc0565b50506040517fe08ba4bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89169063e08ba4bb906120f8908c908b908b90600401614985565b5f604051808303815f87803b15801561210f575f5ffd5b505af1158015612121573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e1682840152606080830187905283518084039091018152608090920190925280519101206105259250600291508a60018d8c8c60405160200161083b93929190614985565b8361219681612bdf565b89898987856121a88486858585613794565b611e2e8f8f8f8f8f8f8f8f613397565b846121c281612bdf565b73ffffffffffffffffffffffffffffffffffffffff8a81165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602090815260408083208054600181019091558b518c84012082517f3babb3dee300cb5c1296b0da361f7f00e8b4572bd64f6e0ecd782c87c264f83081860152958f1686840152606086018e9052608086015260a0850181905260c08086018c90528251808703909101815260e0909501909152835193909101929092209061228782612f51565b90505f6122c9828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612f9892505050565b90508d73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612348576040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8f16600482015260240161069c565b5050505f7f7a638779cbc1a1dd3218e65032adb415412015122b2788070739a2ace3509a318c888860405161237e9291906148e2565b6040519081900381206123bc93929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f6123de82612f51565b905061241f8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612fc092505050565b50506040517fb88d4fde00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b169063b88d4fde906124799030908f908e908e90600401614a97565b5f604051808303815f87803b158015612490575f5ffd5b505af11580156124a2573d5f5f3e3d5ffd5b5050505061253060016124fb8d84604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff9490941681830152606080820193909352815180820390930183526080019052805191012090565b8c8e8d60405160200161170a92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b5050505050505050505050565b8261254781612bdf565b73ffffffffffffffffffffffffffffffffffffffff8981165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020908152604080832080546001810190915589518a84012082517f71afff252ff99c81e2e331f598fa5b6726b8fede3187b1e76ccddde23814194a81860152958e1686840152606086018d9052608086018c905260a086015260c0850181905260e08086018a905282518087039091018152610100909501909152835193909101929092209061261482612f51565b90505f6126218288612f98565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146126a0576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161069c565b50508451602080870191909120604080517ffb69794a96dc00d8da302bbdbf1b9af3b713bfeb066c705c1f55ba40dd7eabed8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f61271a82612f51565b90506127268186612fc0565b50506040517f731133e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063731133e990612780908d908c908c908c90600401614ae7565b5f604051808303815f87803b158015612797575f5ffd5b505af11580156127a9573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f168284015260608083018790528351808403909101815260809092019092528051910120610fce9250600291506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c9052606081018b90528c905f90608001611335565b8461283f81612bdf565b73ffffffffffffffffffffffffffffffffffffffff8b165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7f8c9fe637b6999b95af5c3d94f9198143c67fc87f05c92635e2623e9d030902e48c8c8c8c6040516128c19291906148e2565b6040519081900381206129169493929187908e9060200195865273ffffffffffffffffffffffffffffffffffffffff94909416602086015260408501929092526060840152608083015260a082015260c00190565b6040516020818303038152906040528051906020012090505f61293882612f51565b90505f61297a828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612f9892505050565b90508e73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146129ff578e6040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815260040161069c919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b5050505f7f2dd43c6051e2af91d874ee86be1d3d1dabd6562aa4f73852cebc8779fcbfc6498d8888604051612a359291906148e2565b604051908190038120612a7393929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f612a9582612f51565b9050612ad68187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612fc092505050565b50506040517f94d008ef00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c16906394d008ef90612b30908f908e908e908e90600401614b21565b6020604051808303815f875af1158015612b4c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b709190614b91565b50604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f16828401819052606080840186905284518085039091018152608084019094528351939091019290922060a082019290925260c081018c9052610c39916001918e9060e001610c25565b80421115612c1c576040517fa564eef90000000000000000000000000000000000000000000000000000000081526004810182905260240161069c565b50565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f50c95a183fa9071740cc40b876855b5397200fe0dec9b35ebbeabeb42e6fe5ff60208083019190915273ffffffffffffffffffffffffffffffffffffffff8c811683850152606083018c90528a16608083015260a0820189905260c0820184905260e08083018990528351808403909101815261010090920190925280519101209091505f612cfe82612f51565b90505f612d0b8287612f98565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612d8a576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260240161069c565b50508351602080860191909120604080517f92b6ebbf8a8d1a4a46407e547752332e48ebdb6dbadef77e897e681f10a9c8458185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f612e0482612f51565b9050612e108185612fc0565b50505f5f612e1f88888b61316c565b6040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e8116600483015260248201839052929450909250908b16906379cc6790906044015f604051808303815f87803b158015612e93575f5ffd5b505af1158015612ea5573d5f5f3e3d5ffd5b50505050815f14612ed257612ed273ffffffffffffffffffffffffffffffffffffffff8b168c8a856138d3565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018c9052918a1660e08201526101008101849052612530915f918d9061012001610877565b5f610505612f5d61328b565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f612fa68686613919565b925092509250612fb68282613962565b5090949350505050565b5f612fff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff811661304e576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6130598484612f98565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146130c0576040517f1a126c3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f3ea113d000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633ea113d090613143908890889088908890600401614bd5565b5f604051808303815f87803b15801561315a575f5ffd5b505af1158015610525573d5f5f3e3d5ffd5b5f5f835f0361317f57505f905081613206565b73ffffffffffffffffffffffffffffffffffffffff851615806131a3575061271084115b156131f9576040517fc028e60c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810185905260440161069c565b5050612710828202048082035b935093915050565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517fe98a578400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690819063e98a578490613143908890889088908890600401614bd5565b5f610c50613a69565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f0b7e40c600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190630b7e40c690613143908890889088908890600401614bd5565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261339291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613adc565b505050565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517fe02fdfc6897b4a06000941baacdeb844de72778c2eef3e70571d3e5920463ec160208083019190915273ffffffffffffffffffffffffffffffffffffffff8c811683850152606083018c90528a16608083015260a0820189905260c0820184905260e08083018990528351808403909101815261010090920190925280519101209091505f61347682612f51565b90505f6134838287612f98565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613502576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260240161069c565b50508351602080860191909120604080517fcadc7cd2eeaec3e47314c811398d3aa8a9ecdf0c276677031bb365dc962e99e78185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61357c82612f51565b90506135888185612fc0565b50505f61359687878a61316c565b5090506135bb73ffffffffffffffffffffffffffffffffffffffff8a168b308b6138d3565b80156135e2576135e273ffffffffffffffffffffffffffffffffffffffff8a168883613311565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8d8116838501819052606080850188905285518086039091018152608085019095528451949092019390932060a083019190915260c082018b905291891660e08201526101008101839052610fce915f918c9061012001604051602081830303815290604052613b7b565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916136c190614c44565b80601f01602080910402602001604051908101604052809291908181526020018280546136ed90614c44565b80156137385780601f1061370f57610100808354040283529160200191613738565b820191905f5260205f20905b81548152906001019060200180831161371b57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916136c190614c44565b73ffffffffffffffffffffffffffffffffffffffff85166137e1576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160411461381c576040517f41d3af3500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020810151604080830151606084015191517fd505accf00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff888116600483015230602483015260448201889052606482018790525f9390931a6084820181905260a4820185905260c48201839052919289169063d505accf9060e4015f604051808303815f87803b1580156138c1575f5ffd5b505af1158015610c39573d5f5f3e3d5ffd5b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526130c09186918216906323b872dd9060840161334b565b5f5f5f8351604103613950576020840151604085015160608601515f1a61394288828585613bf8565b95509550955050505061395b565b505081515f91506002905b9250925092565b5f82600381111561397557613975614ba8565b0361397e575050565b600182600381111561399257613992614ba8565b036139c9576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156139dd576139dd614ba8565b03613a17576040517ffce698f70000000000000000000000000000000000000000000000000000000081526004810182905260240161069c565b6003826003811115613a2b57613a2b614ba8565b03613a65576040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004810182905260240161069c565b5050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613a93613ceb565b613a9b613d66565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f5f60205f8451602086015f885af180613afb576040513d5f823e3d81fd5b50505f513d91508115613b12578060011415613b2c565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156130c0576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161069c565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f3a15d07000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633a15d07090613143908890889088908890600401614bd5565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613c3157505f91506003905082613ce1565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613c82573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613cd857505f925060019150829050613ce1565b92505f91508190505b9450945094915050565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613d16613670565b805190915015613d2e57805160209091012092915050565b81548015613d3d579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613d91613743565b805190915015613da957805160209091012092915050565b60018201548015613d3d579392505050565b5f60208284031215613dcb575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114613dfa575f5ffd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613e24575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e9d57613e9d613e29565b604052919050565b5f82601f830112613eb4575f5ffd5b813567ffffffffffffffff811115613ece57613ece613e29565b613eff60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613e56565b818152846020838601011115613f13575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f60808587031215613f42575f5ffd5b613f4b85613e01565b9350613f5960208601613e01565b925060408501359150606085013567ffffffffffffffff811115613f7b575f5ffd5b613f8787828801613ea5565b91505092959194509250565b5f5f5f5f5f5f5f5f610100898b031215613fab575f5ffd5b613fb489613e01565b9750613fc260208a01613e01565b965060408901359550613fd760608a01613e01565b94506080890135935060a0890135925060c089013567ffffffffffffffff811115614000575f5ffd5b61400c8b828c01613ea5565b92505060e089013567ffffffffffffffff811115614028575f5ffd5b6140348b828c01613ea5565b9150509295985092959890939650565b5f5f5f5f5f5f5f60e0888a03121561405a575f5ffd5b61406388613e01565b965061407160208901613e01565b955060408801359450606088013593506080880135925060a088013567ffffffffffffffff8111156140a1575f5ffd5b6140ad8a828b01613ea5565b92505060c088013567ffffffffffffffff8111156140c9575f5ffd5b6140d58a828b01613ea5565b91505092959891949750929550565b5f5f83601f8401126140f4575f5ffd5b50813567ffffffffffffffff81111561410b575f5ffd5b602083019150836020828501011115614122575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b031215614140575f5ffd5b61414989613e01565b975061415760208a01613e01565b96506040890135955060608901359450608089013567ffffffffffffffff811115614180575f5ffd5b61418c8b828c016140e4565b90955093505060a089013567ffffffffffffffff8111156141ab575f5ffd5b6141b78b828c016140e4565b999c989b5096995094979396929594505050565b5f82601f8301126141da575f5ffd5b813567ffffffffffffffff8111156141f4576141f4613e29565b8060051b61420460208201613e56565b9182526020818501810192908101908684111561421f575f5ffd5b6020860192505b83831015614241578235825260209283019290910190614226565b9695505050505050565b5f5f5f5f5f5f5f5f610100898b031215614263575f5ffd5b61426c89613e01565b975061427a60208a01613e01565b9650604089013567ffffffffffffffff811115614295575f5ffd5b6142a18b828c016141cb565b965050606089013567ffffffffffffffff8111156142bd575f5ffd5b6142c98b828c016141cb565b955050608089013567ffffffffffffffff8111156142e5575f5ffd5b6142f18b828c01613ea5565b94505060a0890135925060c089013567ffffffffffffffff811115614000575f5ffd5b5f60208284031215614324575f5ffd5b613dfa82613e01565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f8151808452602084019350602083015f5b828110156143a957815186526020958601959091019060010161438b565b5093949350505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6143ed60e083018961432d565b82810360408401526143ff818961432d565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c084015261443b8185614379565b9a9950505050505050505050565b5f5f5f5f5f5f5f5f610100898b031215614461575f5ffd5b61446a89613e01565b975061447860208a01613e01565b96506040890135955060608901359450608089013567ffffffffffffffff8111156142e5575f5ffd5b5f5f5f5f5f5f5f5f5f6101208a8c0312156144ba575f5ffd5b6144c38a613e01565b98506144d160208b01613e01565b975060408a013596506144e660608b01613e01565b955060808a0135945060a08a0135935060c08a013567ffffffffffffffff81111561450f575f5ffd5b61451b8c828d01613ea5565b93505060e08a013567ffffffffffffffff811115614537575f5ffd5b6145438c828d01613ea5565b9250506101008a013567ffffffffffffffff811115614560575f5ffd5b61456c8c828d01613ea5565b9150509295985092959850929598565b5f5f5f5f5f5f5f60e0888a031215614592575f5ffd5b61459b88613e01565b96506145a960208901613e01565b9550604088013567ffffffffffffffff8111156145c4575f5ffd5b6145d08a828b016141cb565b955050606088013567ffffffffffffffff8111156145ec575f5ffd5b6145f88a828b016141cb565b9450506080880135925060a088013567ffffffffffffffff8111156140a1575f5ffd5b5f5f5f5f5f5f5f5f5f60e08a8c031215614633575f5ffd5b61463c8a613e01565b985061464a60208b01613e01565b975060408a0135965060608a013567ffffffffffffffff81111561466c575f5ffd5b6146788c828d01613ea5565b96505060808a0135945060a08a013567ffffffffffffffff81111561469b575f5ffd5b6146a78c828d016140e4565b90955093505060c08a013567ffffffffffffffff8111156146c6575f5ffd5b6146d28c828d016140e4565b915080935050809150509295985092959850929598565b5f5f5f5f5f60a086880312156146fd575f5ffd5b61470686613e01565b945061471460208701613e01565b9350604086013567ffffffffffffffff81111561472f575f5ffd5b61473b888289016141cb565b935050606086013567ffffffffffffffff811115614757575f5ffd5b614763888289016141cb565b925050608086013567ffffffffffffffff81111561477f575f5ffd5b61478b88828901613ea5565b9150509295509295909350565b5f5f5f5f5f5f5f5f5f5f60e08b8d0312156147b1575f5ffd5b6147ba8b613e01565b99506147c860208c01613e01565b985060408b0135975060608b013567ffffffffffffffff8111156147ea575f5ffd5b6147f68d828e016140e4565b90985096505060808b0135945060a08b013567ffffffffffffffff81111561481c575f5ffd5b6148288d828e016140e4565b90955093505060c08b013567ffffffffffffffff811115614847575f5ffd5b6148538d828e016140e4565b915080935050809150509295989b9194979a5092959850565b5f5f5f5f5f60a08688031215614880575f5ffd5b61488986613e01565b945061489760208701613e01565b93506040860135925060608601359150608086013567ffffffffffffffff81111561477f575f5ffd5b8215158152604060208201525f6148da604083018461432d565b949350505050565b818382375f9101908152919050565b81515f90829060208501835b8281101561491b5781518452602093840193909101906001016148fd565b509195945050505050565b73ffffffffffffffffffffffffffffffffffffffff85168152608060208201525f6149546080830186614379565b82810360408401526149668186614379565b9050828103606084015261497a818561432d565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201525f6149b36060830185614379565b82810360408401526142418185614379565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015260a060408201525f614a0f60a0830186614379565b8281036060840152614a218186614379565b90508281036080840152614a35818561432d565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015282606082015260a060808201525f61497a60a083018461432d565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f614241608083018461432d565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152826040820152608060608201525f614241608083018461432d565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f60208284031215614ba1575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f60038610614c0b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b85825284602083015273ffffffffffffffffffffffffffffffffffffffff8416604083015260806060830152614241608083018461432d565b600181811c90821680614c5857607f821691505b602082108103614c8f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fea2646970667358221220a3e4a97d7fdf4cd1bf04cd0196988b2185162cf2828b8092f52cb8156d8786b564736f6c634300081c0033",
}

// ForgeV2 is an auto generated Go binding around an Ethereum contract.
type ForgeV2 struct {
	abi abi.ABI
}

// NewForgeV2 creates a new instance of ForgeV2.
func NewForgeV2() *ForgeV2 {
	parsed, err := ForgeV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ForgeV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ForgeV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV2 *ForgeV2) PackDOMAINSEPARATOR() []byte {
	enc, err := forgeV2.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV2 *ForgeV2) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := forgeV2.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackBurnERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ddc1f.
//
// Solidity: function burnERC1155(address from, address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC1155(from common.Address, token common.Address, tokenID *big.Int, amount *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC1155", from, token, tokenID, amount, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa84adb32.
//
// Solidity: function burnERC1155Batch(address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC1155Batch(from common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC1155Batch", from, token, tokenIDs, amounts, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2af4c9d0.
//
// Solidity: function burnERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x939cefc1.
//
// Solidity: function burnERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38033a8b.
//
// Solidity: function burnERC721(address from, address token, uint256 tokenID, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC721(from common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC721", from, token, tokenID, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV2 *ForgeV2) PackEip712Domain() []byte {
	enc, err := forgeV2.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV2 *ForgeV2) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := forgeV2.abi.Unpack("eip712Domain", data)
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
	return *outstruct, err

}

// PackMintERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf6bdd97.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43dc7f78.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3081f746.
//
// Solidity: function mintERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb686a78.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV2 *ForgeV2) PackNonces(owner common.Address) []byte {
	enc, err := forgeV2.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV2 *ForgeV2) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := forgeV2.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV2 *ForgeV2) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := forgeV2.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV2 *ForgeV2) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := forgeV2.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8abb7071.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60cfd501.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72033c8c.
//
// Solidity: function transferERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba32f859.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferFromERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a9bb38.
//
// Solidity: function transferFromERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferFromERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferFromERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferFromERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb55ba9c4.
//
// Solidity: function transferFromERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) PackTransferFromERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferFromERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (forgeV2 *ForgeV2) PackValidator() []byte {
	enc, err := forgeV2.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (forgeV2 *ForgeV2) UnpackValidator(data []byte) (common.Address, error) {
	out, err := forgeV2.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// ForgeV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the ForgeV2 contract.
type ForgeV2EIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ForgeV2EIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ForgeV2EIP712DomainChanged) ContractEventName() string {
	return ForgeV2EIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (forgeV2 *ForgeV2) UnpackEIP712DomainChangedEvent(log *types.Log) (*ForgeV2EIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2EIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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

// ForgeV2Initialized represents a Initialized event raised by the ForgeV2 contract.
type ForgeV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ForgeV2Initialized) ContractEventName() string {
	return ForgeV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (forgeV2 *ForgeV2) UnpackInitializedEvent(log *types.Log) (*ForgeV2Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2Initialized)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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

// ForgeV2ValidatorUpdated represents a ValidatorUpdated event raised by the ForgeV2 contract.
type ForgeV2ValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const ForgeV2ValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (ForgeV2ValidatorUpdated) ContractEventName() string {
	return ForgeV2ValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (forgeV2 *ForgeV2) UnpackValidatorUpdatedEvent(log *types.Log) (*ForgeV2ValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2ValidatorUpdated)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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
func (forgeV2 *ForgeV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeECDSAInvalidPermitSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeECDSAInvalidPermitSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeECDSAInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeInvalidFeeData"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeInvalidFeeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return forgeV2.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return forgeV2.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC1155ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC1155ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC20ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC20ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC721ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC721ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return forgeV2.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return forgeV2.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return forgeV2.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return forgeV2.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ForgeV2BaseForgeECDSAInvalidPermitSignature represents a BaseForge__ECDSAInvalidPermitSignature error raised by the ForgeV2 contract.
type ForgeV2BaseForgeECDSAInvalidPermitSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ECDSAInvalidPermitSignature()
func ForgeV2BaseForgeECDSAInvalidPermitSignatureErrorID() common.Hash {
	return common.HexToHash("0x41d3af35d0c283f192ca1e883510e833504b1a251d5c41e80b88d29b340747f3")
}

// UnpackBaseForgeECDSAInvalidPermitSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ECDSAInvalidPermitSignature()
func (forgeV2 *ForgeV2) UnpackBaseForgeECDSAInvalidPermitSignatureError(raw []byte) (*ForgeV2BaseForgeECDSAInvalidPermitSignature, error) {
	out := new(ForgeV2BaseForgeECDSAInvalidPermitSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeECDSAInvalidPermitSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeECDSAInvalidValidatorSignature represents a BaseForge__ECDSAInvalidValidatorSignature error raised by the ForgeV2 contract.
type ForgeV2BaseForgeECDSAInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func ForgeV2BaseForgeECDSAInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0x1a126c3ca91a72268f221dd707e7d3ce39d63b8c31fd1956351a914fc1e8f138")
}

// UnpackBaseForgeECDSAInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func (forgeV2 *ForgeV2) UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw []byte) (*ForgeV2BaseForgeECDSAInvalidValidatorSignature, error) {
	out := new(ForgeV2BaseForgeECDSAInvalidValidatorSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeECDSAInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the ForgeV2 contract.
type ForgeV2BaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func ForgeV2BaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (forgeV2 *ForgeV2) UnpackBaseForgeExpiredSignatureError(raw []byte) (*ForgeV2BaseForgeExpiredSignature, error) {
	out := new(ForgeV2BaseForgeExpiredSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeInvalidFeeData represents a BaseForge__InvalidFeeData error raised by the ForgeV2 contract.
type ForgeV2BaseForgeInvalidFeeData struct {
	FeeRecipient common.Address
	FeeBPS       *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func ForgeV2BaseForgeInvalidFeeDataErrorID() common.Hash {
	return common.HexToHash("0xc028e60c3834e172f648e56b5588dbdbb2caa156ad1931ff196c489c710a7363")
}

// UnpackBaseForgeInvalidFeeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func (forgeV2 *ForgeV2) UnpackBaseForgeInvalidFeeDataError(raw []byte) (*ForgeV2BaseForgeInvalidFeeData, error) {
	out := new(ForgeV2BaseForgeInvalidFeeData)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeInvalidFeeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the ForgeV2 contract.
type ForgeV2BaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func ForgeV2BaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (forgeV2 *ForgeV2) UnpackBaseForgeZeroAddressError(raw []byte) (*ForgeV2BaseForgeZeroAddress, error) {
	out := new(ForgeV2BaseForgeZeroAddress)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ForgeV2 contract.
type ForgeV2ECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ForgeV2ECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (forgeV2 *ForgeV2) UnpackECDSAInvalidSignatureError(raw []byte) (*ForgeV2ECDSAInvalidSignature, error) {
	out := new(ForgeV2ECDSAInvalidSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ForgeV2 contract.
type ForgeV2ECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ForgeV2ECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (forgeV2 *ForgeV2) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ForgeV2ECDSAInvalidSignatureLength, error) {
	out := new(ForgeV2ECDSAInvalidSignatureLength)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ForgeV2 contract.
type ForgeV2ECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ForgeV2ECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (forgeV2 *ForgeV2) UnpackECDSAInvalidSignatureSError(raw []byte) (*ForgeV2ECDSAInvalidSignatureS, error) {
	out := new(ForgeV2ECDSAInvalidSignatureS)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC1155ForgeV2InvalidAccountSignature represents a ERC1155ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC1155ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC1155ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0xa424a9419d285008b7909b6400faa2d8734c2da12fd92cf72aa0ba5a815c12e6")
}

// UnpackERC1155ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC1155ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC1155ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC1155ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC1155ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC20ForgeV2InvalidAccountSignature represents a ERC20ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC20ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC20ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0x9c0f9331d31a14bc40f7b3648f1dd167e28f4c2ce6f3452830ca126ff2afbd53")
}

// UnpackERC20ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC20ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC20ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC20ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC20ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC721ForgeV2InvalidAccountSignature represents a ERC721ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC721ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC721ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0xfb1b8f0f5fbb9ddf1a7ea566d6daa430cb3b1d4f882ecfb03a31df2c5b0c1008")
}

// UnpackERC721ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC721ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC721ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC721ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC721ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2InvalidAccountNonce represents a InvalidAccountNonce error raised by the ForgeV2 contract.
type ForgeV2InvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ForgeV2InvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (forgeV2 *ForgeV2) UnpackInvalidAccountNonceError(raw []byte) (*ForgeV2InvalidAccountNonce, error) {
	out := new(ForgeV2InvalidAccountNonce)
	if err := forgeV2.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2InvalidInitialization represents a InvalidInitialization error raised by the ForgeV2 contract.
type ForgeV2InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ForgeV2InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (forgeV2 *ForgeV2) UnpackInvalidInitializationError(raw []byte) (*ForgeV2InvalidInitialization, error) {
	out := new(ForgeV2InvalidInitialization)
	if err := forgeV2.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2NotInitializing represents a NotInitializing error raised by the ForgeV2 contract.
type ForgeV2NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ForgeV2NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (forgeV2 *ForgeV2) UnpackNotInitializingError(raw []byte) (*ForgeV2NotInitializing, error) {
	out := new(ForgeV2NotInitializing)
	if err := forgeV2.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2SafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ForgeV2 contract.
type ForgeV2SafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ForgeV2SafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (forgeV2 *ForgeV2) UnpackSafeERC20FailedOperationError(raw []byte) (*ForgeV2SafeERC20FailedOperation, error) {
	out := new(ForgeV2SafeERC20FailedOperation)
	if err := forgeV2.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
