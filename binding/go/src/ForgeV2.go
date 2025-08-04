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
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidPermitSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"BaseForge__InvalidFeeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC1155ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC721ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ForgeV2",
	Bin: "0x6080604052348015600e575f5ffd5b50614b618061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610184575f3560e01c80633d3ae4e9116100dd578063978df96e11610088578063bc197c8111610063578063bc197c81146103f0578063e319b87c14610428578063f23a6e611461043b575f5ffd5b8063978df96e146103b7578063a84adb32146103ca578063b55ba9c4146103dd575f5ffd5b806382a9bb38116100b857806382a9bb381461037657806384b0196e14610389578063939cefc1146103a4575f5ffd5b80633d3ae4e9146102fc57806372033c8c1461030f5780637ecebe0014610322575f5ffd5b80632b0f65a51161013d5780633644e515116101185780633644e5151461028c57806338033a8b146102a25780633a5381b5146102b5575f5ffd5b80632b0f65a5146102535780632f2ddc1f146102665780633081f74614610279575f5ffd5b806311d99fb41161016d57806311d99fb4146101c5578063150b7a02146101d85780632af4c9d014610240575f5ffd5b806301ffc9a714610188578063088d63cb146101b0575b5f5ffd5b61019b610196366004613dae565b610473565b60405190151581526020015b60405180910390f35b6101c36101be366004613fa2565b61050b565b005b6101c36101d33660046140af565b6108bb565b61020f6101e636600461410e565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016101a7565b6101c361024e366004614172565b610bc0565b6101c3610261366004614224565b610be5565b6101c36102743660046142c6565b610f72565b6101c3610287366004614172565b6112c8565b610294611674565b6040519081526020016101a7565b6101c36102b0366004614224565b611682565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b6101c361030a366004614366565b6119fb565b6101c361031d366004614172565b611da2565b610294610330366004614430565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b6101c3610384366004614172565b612078565b610391612092565b6040516101a797969594939291906144cf565b6101c36103b2366004614565565b61218c565b6101c36103c53660046140af565b6121c9565b6101c36103d8366004614624565b612508565b6101c36103eb366004614565565b612855565b61020f6103fe3660046146c3565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b6101c3610436366004613fa2565b612881565b61020f610449366004614772565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e000000000000000000000000000000000000000000000000000000000148061050557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b8361051581612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7f68d51bf6ac4c700b2d9c53a4782b15424ee9bc072db19733548034d3de4258e18a8a60405160200161059791906147c6565b604051602081830303815290604052805190602001208a6040516020016105be91906147c6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019590955273ffffffffffffffffffffffffffffffffffffffff909316928101929092526060820152608081019190915260a0810183905260c0810188905260e0016040516020818303038152906040528051906020012090505f61065a82612c12565b90505f6106678289612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106eb576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024015b60405180910390fd5b50508551602080880191909120604080517f811c2af50c1ad6d46b121ba5a43024be13e428134866be726da388c2357856a58185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f61076582612c12565b90506107718187612c81565b50506040517f1f7fdffa00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1690631f7fdffa906107cb908d908c908c9089906004016147fb565b5f604051808303815f87803b1580156107e2575f5ffd5b505af11580156107f4573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f1682840152606080830187905283518084039091018152608090920190925280519101206108af9250600291508b60018e8d8d60405160200161085f9392919061485a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261089b929160200161489a565b604051602081830303815290604052612d87565b50505050505050505050565b836108c581612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f1ef7abfc69e2b8493691b5bc7ef0f6656ac6f2adc11c0ee1d02fee588cda7a1d602082015273ffffffffffffffffffffffffffffffffffffffff8c1691810191909152606081018a90526080810189905260a0810182905260c081018890529091505f9060e0016040516020818303038152906040528051906020012090505f61099f82612c12565b90505f6109ac8289612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a2b576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b50508551602080880191909120604080517ffb69794a96dc00d8da302bbdbf1b9af3b713bfeb066c705c1f55ba40dd7eabed8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f610aa582612c12565b9050610ab18187612c81565b50506040517f731133e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063731133e990610b0b908d908c908c9089906004016148bc565b5f604051808303815f87803b158015610b22575f5ffd5b505af1158015610b34573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f1682840152606080830187905283518084039091018152608090920190925280519101206108af9250600291506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c9052606081018b90528c905f9060800161085f565b82610bca81612bd2565b610bda8989898989898989612e2d565b505050505050505050565b84610bef81612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f4c892b37ef914c8b56eb6107f7952081e6f90e70811186b7de05971d0a563f07602082015273ffffffffffffffffffffffffffffffffffffffff8c1691810191909152606081018a90526080810182905260a081018990529091505f9060c0016040516020818303038152906040528051906020012090505f610cc282612c12565b90505f610d04828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c5992505050565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d83576040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b5050505f7f2dd43c6051e2af91d874ee86be1d3d1dabd6562aa4f73852cebc8779fcbfc6498b8888604051610db99291906148f6565b604051908190038120610df793929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f610e1982612c12565b9050610e5a8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c8192505050565b50506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018a90528a16906340c10f19906044016020604051808303815f875af1158015610ece573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef29190614905565b50604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8d1682840152606080830185905283518084039091018152608090920190925280519101206108af906001906040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c90528c9060600161089b565b82610f7c81612bd2565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f4c82fe77677b753d6e8e9d6fd479564b08b135757219f2044e54db1882de3aa6602082015273ffffffffffffffffffffffffffffffffffffffff8b1691810191909152606081018990526080810188905260a0810182905260c081018790529091505f9060e0016040516020818303038152906040528051906020012090505f61105682612c12565b90505f6110638288612c59565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110e2576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016106e2565b50508451602080870191909120604080517fe9f21759dec582500f25edee01bab124984bac6771cf3f3dcb232329555083478185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61115c82612c12565b90506111688186612c81565b50506040517f124d91e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990526044820188905289169063124d91e5906064015f604051808303815f87803b1580156111de575f5ffd5b505af11580156111f0573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e168284015260608083018790528351808403909101815260809092019092528051910120610bda9250600291506040805173ffffffffffffffffffffffffffffffffffffffff8e1660208201529081018b9052606081018a90528b905f906080015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526112b4929160200161489a565b60405160208183030381529060405261315f565b826112d281612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f9eff2bc106f16f11dab69de8c466346a0a369dd1d3a0878683947093830994fa60208083019190915273ffffffffffffffffffffffffffffffffffffffff8d811683850152606083018d90528b16608083015260a082018a905260c0820184905260e08083018a90528351808403909101815261010090920190925280519101209091505f6113b182612c12565b90505f6113be8288612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461143d576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b50508451602080870191909120604080517f340f6ac2b427cb47a7172eaa76879e09534311918b4cff11daea1028ca325aa88185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f6114b782612c12565b90506114c38186612c81565b50505f5f6114d289898c6131dc565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8f8116600483015260248201839052929450909250908c16906340c10f19906044015f604051808303815f87803b158015611546575f5ffd5b505af1158015611558573d5f5f3e3d5ffd5b50505050815f146115e7576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018490528c16906340c10f19906044015f604051808303815f87803b1580156115d0575f5ffd5b505af11580156115e2573d5f5f3e3d5ffd5b505050505b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018d9052918b1660e08201526101008101849052611666915f918e906101200161089b565b505050505050505050505050565b5f61167d61327e565b905090565b8461168c81612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f4e6479369d4c42395619be8224f096421cd6800a840e701655040f5a933c33cd602082015273ffffffffffffffffffffffffffffffffffffffff8c1691810191909152606081018a90526080810182905260a081018990529091505f9060c0016040516020818303038152906040528051906020012090505f61175f82612c12565b90505f6117a1828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c5992505050565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611820576040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b5050505f7f42dec6f4478d32028294c542bb3ee6cb16fb5523fd8f13922dca479fe76475138b88886040516118569291906148f6565b60405190819003812061189493929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f6118b682612c12565b90506118f78187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c8192505050565b50506040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018a90528a16906379cc6790906044015f604051808303815f87803b158015611966575f5ffd5b505af1158015611978573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f1682840152606080830187905283518084039091018152608090920190925280519101206108af9250600191506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c90528c906060016112b4565b85611a0581612bd2565b73ffffffffffffffffffffffffffffffffffffffff8a165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f898da923fc7c770744862aa43404766a802c312f63aae08d2416994e4e79be2c602082015273ffffffffffffffffffffffffffffffffffffffff8d1691810191909152606081018b90526080810182905260a081018a90529091505f9060c0016040516020818303038152906040528051906020012090505f611ad882612c12565b90505f611b1a828b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c5992505050565b90508d73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b99576040517ffb1b8f0f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8f1660048201526024016106e2565b5050505f7f7a638779cbc1a1dd3218e65032adb415412015122b2788070739a2ace3509a318c8989604051611bcf9291906148f6565b604051908190038120611c0d93929160200192835273ffffffffffffffffffffffffffffffffffffffff919091166020830152604082015260600190565b6040516020818303038152906040528051906020012090505f611c2f82612c12565b9050611c708188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612c8192505050565b50506040517fb88d4fde00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b169063b88d4fde90611cca9030908f908e90899060040161491c565b5f604051808303815f87803b158015611ce1575f5ffd5b505af1158015611cf3573d5f5f3e3d5ffd5b50505050611d956001611d4c8d84604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff9490941681830152606080820193909352815180820390930183526080019052805191012090565b8c8e8d604051602001611d8192919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b604051602081830303815290604052613287565b5050505050505050505050565b82611dac81612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517fab25e47ab9e13a956bdee490006a87ac9cbe78ae734a45c431fe3938cbd0858160208083019190915273ffffffffffffffffffffffffffffffffffffffff8d811683850152606083018d90528b16608083015260a082018a905260c0820184905260e08083018a90528351808403909101815261010090920190925280519101209091505f611e8b82612c12565b90505f611e988288612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611f17576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b50508451602080870191909120604080517f7ae019cd8a06bf0d82d3d0063c026c8cfe203e2686c8de51ae89d5645e022d618185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f611f9182612c12565b9050611f9d8186612c81565b50505f5f611fac89898c6131dc565b9092509050611fd273ffffffffffffffffffffffffffffffffffffffff8c168d83613304565b8115611ff957611ff973ffffffffffffffffffffffffffffffffffffffff8c168a84613304565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018d9052918b1660e08201526101008101849052611666915f918e9061012001611d81565b8261208281612bd2565b610bda898989898989898961338a565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10080549091501580156120d057506001810154155b612136576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016106e2565b61213e613663565b612146613736565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b8361219681612bd2565b89898987856121a88486858585613787565b6121b88f8f8f8f8f8f8f8f612e2d565b505050505050505050505050505050565b836121d381612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517fa0ad07bd3d5d0ba1e944135f6fba4a6c631fd50830c4c4f7212a557370901e58602082015273ffffffffffffffffffffffffffffffffffffffff8c1691810191909152606081018a90526080810189905260a0810182905260c081018890529091505f9060e0016040516020818303038152906040528051906020012090505f6122ad82612c12565b90505f6122ba8289612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612339576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b50508551602080880191909120604080517fadcb1a40c0ace98e86ce9038ac7bb2e74cec43bdaa9e5727e2ea08c07a1f36dd8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f6123b382612c12565b90506123bf8187612c81565b50506040517ff242432a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063f242432a9061241b9030908e908d908d908a9060040161496c565b5f604051808303815f87803b158015612432575f5ffd5b505af1158015612444573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f1682840152606080830187905283518084039091018152608090920190925280519101206108af9250600291506040805173ffffffffffffffffffffffffffffffffffffffff8f1660208201529081018c9052606081018b90528c905f906080015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611d81929160200161489a565b8261251281612bd2565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7fda6dbf9c08df397da92be8fda42638fad009adf2a495b08f3526089ad02e7dd1898960405160200161259491906147c6565b60405160208183030381529060405280519060200120896040516020016125bb91906147c6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019590955273ffffffffffffffffffffffffffffffffffffffff909316928101929092526060820152608081019190915260a0810183905260c0810187905260e0016040516020818303038152906040528051906020012090505f61265782612c12565b90505f6126648288612c59565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146126e3576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016106e2565b50508451602080870191909120604080517fa221f8b9ccdf4675afa9893d3a99032eb17f5b0c291026fd55d97bb9993fa4e38185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61275d82612c12565b90506127698186612c81565b50506040517fe08ba4bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89169063e08ba4bb906127c1908c908b908b9060040161485a565b5f604051808303815f87803b1580156127d8575f5ffd5b505af11580156127ea573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e168284015260608083018790528351808403909101815260809092019092528051910120610bda9250600291508a60018d8c8c6040516020016112789392919061485a565b8361285f81612bd2565b89898987856128718486858585613787565b6121b88f8f8f8f8f8f8f8f61338a565b8361288b81612bd2565b73ffffffffffffffffffffffffffffffffffffffff89165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040812080546001810190915590505f7f34004b82faef882e1e55da2bb82673c86f97d3cff37ec31e09fd2d60db0cdf1a8a8a60405160200161290d91906147c6565b604051602081830303815290604052805190602001208a60405160200161293491906147c6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019590955273ffffffffffffffffffffffffffffffffffffffff909316928101929092526060820152608081019190915260a0810183905260c0810188905260e0016040516020818303038152906040528051906020012090505f6129d082612c12565b90505f6129dd8289612c59565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612a5c576040517fa424a94100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e1660048201526024016106e2565b50508551602080880191909120604080517fbe51d3aafe3852e9cbe2ce5b4d184540bcb68868fc9307bf8b45d1f170d6739f8185015273ffffffffffffffffffffffffffffffffffffffff8f1681830152606080820193909352815180820390930183526080019052805191012090505f612ad682612c12565b9050612ae28187612c81565b50506040517f2eb2c2d600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a1690632eb2c2d690612b3e9030908e908d908d908a906004016149c2565b5f604051808303815f87803b158015612b55575f5ffd5b505af1158015612b67573d5f5f3e3d5ffd5b5050604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8f1682840152606080830187905283518084039091018152608090920190925280519101206108af9250600291508b60018e8d8d6040516020016124cc9392919061485a565b80421115612c0f576040517fa564eef9000000000000000000000000000000000000000000000000000000008152600481018290526024016106e2565b50565b5f610505612c1e61327e565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f612c6786866138c6565b925092509250612c77828261390f565b5090949350505050565b5f612cc07f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff8116612d0f576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f612d1a8484612c59565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614612d81576040517f1a126c3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517fe98a578400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690819063e98a578490612e04908890889088908890600401614a6b565b5f604051808303815f87803b158015612e1b575f5ffd5b505af1158015610bda573d5f5f3e3d5ffd5b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517f50c95a183fa9071740cc40b876855b5397200fe0dec9b35ebbeabeb42e6fe5ff60208083019190915273ffffffffffffffffffffffffffffffffffffffff8c811683850152606083018c90528a16608083015260a0820189905260c0820184905260e08083018990528351808403909101815261010090920190925280519101209091505f612f0c82612c12565b90505f612f198287612c59565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612f98576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016106e2565b50508351602080860191909120604080517f92b6ebbf8a8d1a4a46407e547752332e48ebdb6dbadef77e897e681f10a9c8458185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61301282612c12565b905061301e8185612c81565b50505f5f61302d88888b6131dc565b6040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e8116600483015260248201839052929450909250908b16906379cc6790906044015f604051808303815f87803b1580156130a1575f5ffd5b505af11580156130b3573d5f5f3e3d5ffd5b50505050815f146130e0576130e073ffffffffffffffffffffffffffffffffffffffff8b168c8a85613a16565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8e8116838501819052606080850189905285518086039091018152608085019095528451949092019390932060a083019190915260c082018c9052918a1660e08201526101008101849052611d95915f918d90610120016112b4565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f3ea113d000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633ea113d090612e04908890889088908890600401614a6b565b5f5f835f036131ef57505f905081613276565b73ffffffffffffffffffffffffffffffffffffffff85161580613213575061271084115b15613269576040517fc028e60c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602481018590526044016106e2565b5050612710828202048082035b935093915050565b5f61167d613a5c565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f0b7e40c600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190630b7e40c690612e04908890889088908890600401614a6b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261338591859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613acf565b505050565b73ffffffffffffffffffffffffffffffffffffffff88165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260408120805460018101909155604080517fe02fdfc6897b4a06000941baacdeb844de72778c2eef3e70571d3e5920463ec160208083019190915273ffffffffffffffffffffffffffffffffffffffff8c811683850152606083018c90528a16608083015260a0820189905260c0820184905260e08083018990528351808403909101815261010090920190925280519101209091505f61346982612c12565b90505f6134768287612c59565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146134f5576040517f9c0f933100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016106e2565b50508351602080860191909120604080517fcadc7cd2eeaec3e47314c811398d3aa8a9ecdf0c276677031bb365dc962e99e78185015273ffffffffffffffffffffffffffffffffffffffff8e1681830152606080820193909352815180820390930183526080019052805191012090505f61356f82612c12565b905061357b8185612c81565b50505f61358987878a6131dc565b5090506135ae73ffffffffffffffffffffffffffffffffffffffff8a168b308b613a16565b80156135d5576135d573ffffffffffffffffffffffffffffffffffffffff8a168883613304565b604080513060208083019190915273ffffffffffffffffffffffffffffffffffffffff8d8116838501819052606080850188905285518086039091018152608085019095528451949092019390932060a083019190915260c082018b905291891660e082015261010081018390526108af915f918c9061012001604051602081830303815290604052613b6e565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916136b490614ada565b80601f01602080910402602001604051908101604052809291908181526020018280546136e090614ada565b801561372b5780601f106137025761010080835404028352916020019161372b565b820191905f5260205f20905b81548152906001019060200180831161370e57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916136b490614ada565b73ffffffffffffffffffffffffffffffffffffffff85166137d4576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160411461380f576040517f41d3af3500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020810151604080830151606084015191517fd505accf00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff888116600483015230602483015260448201889052606482018790525f9390931a6084820181905260a4820185905260c48201839052919289169063d505accf9060e4015f604051808303815f87803b1580156138b4575f5ffd5b505af1158015611666573d5f5f3e3d5ffd5b5f5f5f83516041036138fd576020840151604085015160608601515f1a6138ef88828585613beb565b955095509550505050613908565b505081515f91506002905b9250925092565b5f82600381111561392257613922614a3e565b0361392b575050565b600182600381111561393f5761393f614a3e565b03613976576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561398a5761398a614a3e565b036139c4576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016106e2565b60038260038111156139d8576139d8614a3e565b03613a12576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016106e2565b5050565b60405173ffffffffffffffffffffffffffffffffffffffff8481166024830152838116604483015260648201839052612d819186918216906323b872dd9060840161333e565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613a86613cde565b613a8e613d59565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f5f60205f8451602086015f885af180613aee576040513d5f823e3d81fd5b50505f513d91508115613b05578060011415613b1f565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15612d81576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016106e2565b7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00546040517f3a15d07000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633a15d07090612e04908890889088908890600401614a6b565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613c2457505f91506003905082613cd4565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613c75573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613ccb57505f925060019150829050613cd4565b92505f91508190505b9450945094915050565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613d09613663565b805190915015613d2157805160209091012092915050565b81548015613d30579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613d84613736565b805190915015613d9c57805160209091012092915050565b60018201548015613d30579392505050565b5f60208284031215613dbe575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114613ded575f5ffd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613e17575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e9057613e90613e1c565b604052919050565b5f82601f830112613ea7575f5ffd5b813567ffffffffffffffff811115613ec157613ec1613e1c565b8060051b613ed160208201613e49565b91825260208185018101929081019086841115613eec575f5ffd5b6020860192505b83831015613f0e578235825260209283019290910190613ef3565b9695505050505050565b5f82601f830112613f27575f5ffd5b813567ffffffffffffffff811115613f4157613f41613e1c565b613f7260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613e49565b818152846020838601011115613f86575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f5f5f610100898b031215613fba575f5ffd5b613fc389613df4565b9750613fd160208a01613df4565b9650604089013567ffffffffffffffff811115613fec575f5ffd5b613ff88b828c01613e98565b965050606089013567ffffffffffffffff811115614014575f5ffd5b6140208b828c01613e98565b9550506080890135935060a089013567ffffffffffffffff811115614043575f5ffd5b61404f8b828c01613f18565b93505060c089013567ffffffffffffffff81111561406b575f5ffd5b6140778b828c01613f18565b92505060e089013567ffffffffffffffff811115614093575f5ffd5b61409f8b828c01613f18565b9150509295985092959890939650565b5f5f5f5f5f5f5f5f610100898b0312156140c7575f5ffd5b6140d089613df4565b97506140de60208a01613df4565b965060408901359550606089013594506080890135935060a089013567ffffffffffffffff811115614043575f5ffd5b5f5f5f5f60808587031215614121575f5ffd5b61412a85613df4565b935061413860208601613df4565b925060408501359150606085013567ffffffffffffffff81111561415a575f5ffd5b61416687828801613f18565b91505092959194509250565b5f5f5f5f5f5f5f5f610100898b03121561418a575f5ffd5b61419389613df4565b97506141a160208a01613df4565b9650604089013595506141b660608a01613df4565b94506080890135935060a0890135925060c089013567ffffffffffffffff81111561406b575f5ffd5b5f5f83601f8401126141ef575f5ffd5b50813567ffffffffffffffff811115614206575f5ffd5b60208301915083602082850101111561421d575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b03121561423b575f5ffd5b61424489613df4565b975061425260208a01613df4565b96506040890135955060608901359450608089013567ffffffffffffffff81111561427b575f5ffd5b6142878b828c016141df565b90955093505060a089013567ffffffffffffffff8111156142a6575f5ffd5b6142b28b828c016141df565b999c989b5096995094979396929594505050565b5f5f5f5f5f5f5f60e0888a0312156142dc575f5ffd5b6142e588613df4565b96506142f360208901613df4565b955060408801359450606088013593506080880135925060a088013567ffffffffffffffff811115614323575f5ffd5b61432f8a828b01613f18565b92505060c088013567ffffffffffffffff81111561434b575f5ffd5b6143578a828b01613f18565b91505092959891949750929550565b5f5f5f5f5f5f5f5f5f60e08a8c03121561437e575f5ffd5b6143878a613df4565b985061439560208b01613df4565b975060408a0135965060608a0135955060808a013567ffffffffffffffff8111156143be575f5ffd5b6143ca8c828d016141df565b90965094505060a08a013567ffffffffffffffff8111156143e9575f5ffd5b6143f58c828d016141df565b90945092505060c08a013567ffffffffffffffff811115614414575f5ffd5b6144208c828d01613f18565b9150509295985092959850929598565b5f60208284031215614440575f5ffd5b613ded82613df4565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f8151808452602084019350602083015f5b828110156144c55781518652602095860195909101906001016144a7565b5093949350505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61450960e0830189614449565b828103604084015261451b8189614449565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c08401526145578185614495565b9a9950505050505050505050565b5f5f5f5f5f5f5f5f5f6101208a8c03121561457e575f5ffd5b6145878a613df4565b985061459560208b01613df4565b975060408a013596506145aa60608b01613df4565b955060808a0135945060a08a0135935060c08a013567ffffffffffffffff8111156145d3575f5ffd5b6145df8c828d01613f18565b93505060e08a013567ffffffffffffffff8111156145fb575f5ffd5b6146078c828d01613f18565b9250506101008a013567ffffffffffffffff811115614414575f5ffd5b5f5f5f5f5f5f5f60e0888a03121561463a575f5ffd5b61464388613df4565b965061465160208901613df4565b9550604088013567ffffffffffffffff81111561466c575f5ffd5b6146788a828b01613e98565b955050606088013567ffffffffffffffff811115614694575f5ffd5b6146a08a828b01613e98565b9450506080880135925060a088013567ffffffffffffffff811115614323575f5ffd5b5f5f5f5f5f60a086880312156146d7575f5ffd5b6146e086613df4565b94506146ee60208701613df4565b9350604086013567ffffffffffffffff811115614709575f5ffd5b61471588828901613e98565b935050606086013567ffffffffffffffff811115614731575f5ffd5b61473d88828901613e98565b925050608086013567ffffffffffffffff811115614759575f5ffd5b61476588828901613f18565b9150509295509295909350565b5f5f5f5f5f60a08688031215614786575f5ffd5b61478f86613df4565b945061479d60208701613df4565b93506040860135925060608601359150608086013567ffffffffffffffff811115614759575f5ffd5b81515f90829060208501835b828110156147f05781518452602093840193909101906001016147d2565b509195945050505050565b73ffffffffffffffffffffffffffffffffffffffff85168152608060208201525f6148296080830186614495565b828103604084015261483b8186614495565b9050828103606084015261484f8185614449565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201525f6148886060830185614495565b8281036040840152613f0e8185614495565b8215158152604060208201525f6148b46040830184614449565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152826040820152608060608201525f613f0e6080830184614449565b818382375f9101908152919050565b5f60208284031215614915575f5ffd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f613f0e6080830184614449565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015282606082015260a060808201525f61484f60a0830184614449565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015260a060408201525f614a0c60a0830186614495565b8281036060840152614a1e8186614495565b90508281036080840152614a328185614449565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f60038610614aa1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b85825284602083015273ffffffffffffffffffffffffffffffffffffffff8416604083015260806060830152613f0e6080830184614449565b600181811c90821680614aee57607f821691505b602082108103614b25577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fea2646970667358221220cfe13e83f4481ad2e6f9b0a978c1cb1dc804e1e66bf9928a8a791b3d050efb4864736f6c634300081c0033",
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
// the contract method with ID 0x11d99fb4.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes recipientSig, bytes validatorSig, bytes data) returns()
func (forgeV2 *ForgeV2) PackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte, data []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155", recipient, token, tokenID, amount, deadline, recipientSig, validatorSig, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x088d63cb.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 deadline, bytes recipientSig, bytes validatorSig, bytes data) returns()
func (forgeV2 *ForgeV2) PackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte, data []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, deadline, recipientSig, validatorSig, data)
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
// the contract method with ID 0x2b0f65a5.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC721", recipient, token, tokenID, deadline, recipientSig, validatorSig)
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
// the contract method with ID 0x978df96e.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes recipientSig, bytes validatorSig, bytes data) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte, data []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155", recipient, token, tokenID, amount, deadline, recipientSig, validatorSig, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe319b87c.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 deadline, bytes recipientSig, bytes validatorSig, bytes data) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte, data []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, deadline, recipientSig, validatorSig, data)
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
// the contract method with ID 0x3d3ae4e9.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, uint256 deadline, bytes recipientSig, bytes validatorSig, bytes data) returns()
func (forgeV2 *ForgeV2) PackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte, data []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC721", recipient, token, tokenID, deadline, recipientSig, validatorSig, data)
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
