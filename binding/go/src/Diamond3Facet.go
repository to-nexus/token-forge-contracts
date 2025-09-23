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

// Diamond3FacetMetaData contains all meta data concerning the Diamond3Facet contract.
var Diamond3FacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DIAMOND3_FACET_FUNCTIONS\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDiamondFacetCut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializationContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"InitializationFunctionReverted\",\"type\":\"error\"}]",
	ID:  "Diamond3Facet",
	Bin: "0x608060405234801561000f575f5ffd5b5060408051610100810182526307e4c70760e21b8152637a0ed62760e01b60208201526356fe50af60e11b918101919091526314bbdacb60e21b60608201526366ffd66360e11b60808201526301ffc9a760e01b60a082015263f2fde38b60e01b60c0820152638da5cb5b60e01b60e082015261008f905f906008610095565b50610152565b828054828255905f5260205f209060070160089004810192821561012e579160200282015f5b838211156100fc57835183826101000a81548163ffffffff021916908360e01c021790555092602001926004016020816003010492830192600103026100bb565b801561012c5782816101000a81549063ffffffff02191690556004016020816003010492830192600103026100fc565b505b5061013a92915061013e565b5090565b5b8082111561013a575f815560010161013f565b6124898061015f5f395ff3fe608060405234801561000f575f5ffd5b50600436106100b9575f3560e01c80637a0ed62711610072578063adfca15e11610058578063adfca15e146101fa578063cdffacc61461021a578063f2fde38b14610290575f5ffd5b80637a0ed627146101b85780638da5cb5b146101cd575f5ffd5b80631f931c1c116100a25780631f931c1c1461014a57806352ef6b2c1461015f578063549fc1a414610174575f5ffd5b806301ffc9a7146100bd5780631bd8f7f114610135575b5f5ffd5b6101206100cb366004611c17565b7fffffffff00000000000000000000000000000000000000000000000000000000165f9081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f602052604090205460ff1690565b60405190151581526020015b60405180910390f35b61013d6102a3565b60405161012c9190611d3d565b61015d610158366004611db7565b610366565b005b6101676103be565b60405161012c9190611e67565b610187610182366004611ebf565b61044d565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161012c565b6101c0610480565b60405161012c9190611ed6565b6101d5610676565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012c565b61020d610208366004611f84565b6106ba565b60405161012c9190611f9d565b6101d5610228366004611c17565b7fffffffff00000000000000000000000000000000000000000000000000000000165f9081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b61015d61029e366004611f84565b6107b6565b60408051606080820183525f808352602083015291810191909152604080516060810182523081525f6020808301829052815484518183028101830186528181529394850193929183018282801561035957602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116103065790505b5050505050815250905090565b61036e6107ca565b6103b761037b85876120be565b8484848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061089a92505050565b5050505050565b60605f7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6002810180546040805160208084028201810190925282815293945083018282801561044257602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610417575b505050505091505090565b5f818154811061045b575f80fd5b905f5260205f209060089182820401919006600402915054906101000a900460e01b81565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e546060907fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c908067ffffffffffffffff8111156104e0576104e0611ff6565b60405190808252806020026020018201604052801561052557816020015b604080518082019091525f8152606060208201528152602001906001900390816104fe5790505b5092505f5b81811015610670575f836002018281548110610548576105486121e0565b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080858381518110610585576105856121e0565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff92831690529082165f9081526001860182526040908190208054825181850281018501909352808352919290919083018282801561064257602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116105ef5790505b5050505050858381518110610659576106596121e0565b60209081029190910181015101525060010161052a565b50505090565b5f6106b57fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c13205473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d602090815260409182902080548351818402810184019094528084526060937fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c93909291908301828280156107a957602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116107565790505b5050505050915050919050565b6107be6107ca565b6107c781610acb565b50565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6004015473ffffffffffffffffffffffffffffffffffffffff163314610898576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b565b5f5b8351811015610a80575f8482815181106108b8576108b86121e0565b60200260200101516020015190505f60028111156108d8576108d8611c37565b8160028111156108ea576108ea611c37565b0361093757610932858381518110610904576109046121e0565b60200260200101515f0151868481518110610921576109216121e0565b602002602001015160400151610b84565b610a77565b600181600281111561094b5761094b611c37565b0361099357610932858381518110610965576109656121e0565b60200260200101515f0151868481518110610982576109826121e0565b602002602001015160400151610f26565b60028160028111156109a7576109a7611c37565b036109ef576109328583815181106109c1576109c16121e0565b60200260200101515f01518684815181106109de576109de6121e0565b6020026020010151604001516112d8565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f7272656374204661636574437560448201527f74416374696f6e00000000000000000000000000000000000000000000000000606482015260840161088f565b5060010161089c565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673838383604051610ab493929190612259565b60405180910390a1610ac682826114bc565b505050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff8481169182179093556040517fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f815111610c14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660448201527f6163657420746f20637574000000000000000000000000000000000000000000606482015260840161088f565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c73ffffffffffffffffffffffffffffffffffffffff8316610cd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201527f6520616464726573732830290000000000000000000000000000000000000000606482015260840161088f565b73ffffffffffffffffffffffffffffffffffffffff83165f908152600182016020526040812054906bffffffffffffffffffffffff82169003610d1f57610d1f82856115b1565b5f5b83518110156103b7575f848281518110610d3d57610d3d6121e0565b6020908102919091018101517fffffffff0000000000000000000000000000000000000000000000000000000081165f9081529186905260409091205490915073ffffffffffffffffffffffffffffffffffffffff168015610e21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f60448201527f6e207468617420616c7265616479206578697374730000000000000000000000606482015260840161088f565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260208781526040808320805473ffffffffffffffffffffffffffffffffffffffff908116740100000000000000000000000000000000000000006bffffffffffffffffffffffff8c16021782558c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281547fffffffffffffffffffffffff00000000000000000000000000000000000000001617905583610f1581612337565b94505060019092019150610d219050565b5f815111610fb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660448201527f6163657420746f20637574000000000000000000000000000000000000000000606482015260840161088f565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c73ffffffffffffffffffffffffffffffffffffffff831661107a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201527f6520616464726573732830290000000000000000000000000000000000000000606482015260840161088f565b73ffffffffffffffffffffffffffffffffffffffff83165f908152600182016020526040812054906bffffffffffffffffffffffff821690036110c1576110c182856115b1565b5f5b83518110156103b7575f8482815181106110df576110df6121e0565b6020908102919091018101517fffffffff0000000000000000000000000000000000000000000000000000000081165f9081529186905260409091205490915073ffffffffffffffffffffffffffffffffffffffff90811690871681036111c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e0000000000000000606482015260840161088f565b6111d385828461163f565b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260208781526040808320805473ffffffffffffffffffffffffffffffffffffffff908116740100000000000000000000000000000000000000006bffffffffffffffffffffffff8c16021782558c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281547fffffffffffffffffffffffff000000000000000000000000000000000000000016179055836112c781612337565b945050600190920191506110c39050565b5f815111611368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660448201527f6163657420746f20637574000000000000000000000000000000000000000000606482015260840161088f565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c73ffffffffffffffffffffffffffffffffffffffff83161561142d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f7665206661636574206164647260448201527f657373206d757374206265206164647265737328302900000000000000000000606482015260840161088f565b5f5b82518110156114b6575f83828151811061144b5761144b6121e0565b6020908102919091018101517fffffffff0000000000000000000000000000000000000000000000000000000081165f9081529185905260409091205490915073ffffffffffffffffffffffffffffffffffffffff166114ac84828461163f565b505060010161142f565b50505050565b73ffffffffffffffffffffffffffffffffffffffff82166114db575050565b6114fd8260405180606001604052806028815260200161240860289139611ba8565b5f5f8373ffffffffffffffffffffffffffffffffffffffff1683604051611524919061236b565b5f60405180830381855af49150503d805f811461155c576040519150601f19603f3d011682016040523d82523d5f602084013e611561565b606091505b5091509150816114b65780511561157b5780518082602001fd5b83836040517f192105d700000000000000000000000000000000000000000000000000000000815260040161088f929190612381565b6115d38160405180606001604052806024815260200161243060249139611ba8565b60028201805473ffffffffffffffffffffffffffffffffffffffff9092165f8181526001948501602090815260408220860185905594840183559182529290200180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169091179055565b73ffffffffffffffffffffffffffffffffffffffff82166116e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e2774206578697374000000000000000000606482015260840161088f565b3073ffffffffffffffffffffffffffffffffffffffff831603611787576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201527f7461626c652066756e6374696f6e000000000000000000000000000000000000606482015260840161088f565b7fffffffff0000000000000000000000000000000000000000000000000000000081165f908152602084815260408083205473ffffffffffffffffffffffffffffffffffffffff86168452600180880190935290832054740100000000000000000000000000000000000000009091046bffffffffffffffffffffffff169291611810916123af565b90508082146119545773ffffffffffffffffffffffffffffffffffffffff84165f9081526001860160205260408120805483908110611851576118516121e0565b5f918252602080832060088304015473ffffffffffffffffffffffffffffffffffffffff8916845260018a019091526040909220805460079092166004026101000a90920460e01b9250829190859081106118ae576118ae6121e0565b5f91825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c929092029390931790557fffffffff0000000000000000000000000000000000000000000000000000000092909216825286905260409020805473ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000006bffffffffffffffffffffffff8516021790555b73ffffffffffffffffffffffffffffffffffffffff84165f9081526001860160205260409020805480611989576119896123c8565b5f828152602080822060087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90940193840401805463ffffffff600460078716026101000a0219169055919092557fffffffff0000000000000000000000000000000000000000000000000000000085168252869052604081208190558190036103b75760028501545f90611a20906001906123af565b73ffffffffffffffffffffffffffffffffffffffff86165f908152600180890160205260409091200154909150808214611b0a575f876002018381548110611a6a57611a6a6121e0565b5f9182526020909120015460028901805473ffffffffffffffffffffffffffffffffffffffff9092169250829184908110611aa757611aa76121e0565b5f91825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055929091168152600189810190925260409020018190555b86600201805480611b1d57611b1d6123c8565b5f828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff88168252600189810190915260408220015550505050505050565b813b81816114b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088f91906123f5565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114611c12575f5ffd5b919050565b5f60208284031215611c27575f5ffd5b611c3082611be3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f8151808452602084019350602083015f5b82811015611cb65781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101611c76565b5093949350505050565b73ffffffffffffffffffffffffffffffffffffffff81511682525f602082015160038110611d15577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b80602085015250604082015160606040850152611d356060850182611c64565b949350505050565b602081525f611c306020830184611cc0565b803573ffffffffffffffffffffffffffffffffffffffff81168114611c12575f5ffd5b5f5f83601f840112611d82575f5ffd5b50813567ffffffffffffffff811115611d99575f5ffd5b602083019150836020828501011115611db0575f5ffd5b9250929050565b5f5f5f5f5f60608688031215611dcb575f5ffd5b853567ffffffffffffffff811115611de1575f5ffd5b8601601f81018813611df1575f5ffd5b803567ffffffffffffffff811115611e07575f5ffd5b8860208260051b8401011115611e1b575f5ffd5b602091820196509450611e2f908701611d4f565b9250604086013567ffffffffffffffff811115611e4a575f5ffd5b611e5688828901611d72565b969995985093965092949392505050565b602080825282518282018190525f918401906040840190835b81811015611eb457835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101611e80565b509095945050505050565b5f60208284031215611ecf575f5ffd5b5035919050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611f78577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815173ffffffffffffffffffffffffffffffffffffffff81511686526020810151905060406020870152611f626040870182611c64565b9550506020938401939190910190600101611efc565b50929695505050505050565b5f60208284031215611f94575f5ffd5b611c3082611d4f565b602080825282518282018190525f918401906040840190835b81811015611eb45783517fffffffff0000000000000000000000000000000000000000000000000000000016835260209384019390920191600101611fb6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561204657612046611ff6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561209357612093611ff6565b604052919050565b5f67ffffffffffffffff8211156120b4576120b4611ff6565b5060051b60200190565b5f6120d06120cb8461209b565b61204c565b8381526020810190600585901b8401368111156120eb575f5ffd5b845b81811015611eb457803567ffffffffffffffff81111561210b575f5ffd5b8601606036829003121561211d575f5ffd5b612125612023565b61212e82611d4f565b8152602082013560038110612141575f5ffd5b6020820152604082013567ffffffffffffffff81111561215f575f5ffd5b919091019036601f830112612172575f5ffd5b81356121806120cb8261209b565b8082825260208201915060208360051b8601019250368311156121a1575f5ffd5b6020850194505b828510156121ca576121b985611be3565b8252602094850194909101906121a8565b60408401525050855250602093840193016120ed565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b828110156122ce577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808786030184526122b9858351611cc0565b9450602093840193919091019060010161227f565b5050505073ffffffffffffffffffffffffffffffffffffffff851660208401528281036040840152612300818561220d565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6bffffffffffffffffffffffff82166bffffffffffffffffffffffff81036123625761236261230a565b60010192915050565b5f82518060208501845e5f920191825250919050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f611d35604083018461220d565b818103818111156123c2576123c261230a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b602081525f611c30602083018461220d56fe4c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a204e657720666163657420686173206e6f20636f6465a2646970667358221220b0f64c1a86330a1519d68e3571e6116fbc20033c3103be2392fb781f46374d9364736f6c634300081c0033",
}

// Diamond3Facet is an auto generated Go binding around an Ethereum contract.
type Diamond3Facet struct {
	abi abi.ABI
}

// NewDiamond3Facet creates a new instance of Diamond3Facet.
func NewDiamond3Facet() *Diamond3Facet {
	parsed, err := Diamond3FacetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Diamond3Facet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Diamond3Facet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDIAMOND3FACETFUNCTIONS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x549fc1a4.
//
// Solidity: function DIAMOND3_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (diamond3Facet *Diamond3Facet) PackDIAMOND3FACETFUNCTIONS(arg0 *big.Int) []byte {
	enc, err := diamond3Facet.abi.Pack("DIAMOND3_FACET_FUNCTIONS", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDIAMOND3FACETFUNCTIONS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x549fc1a4.
//
// Solidity: function DIAMOND3_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (diamond3Facet *Diamond3Facet) UnpackDIAMOND3FACETFUNCTIONS(data []byte) ([4]byte, error) {
	out, err := diamond3Facet.abi.Unpack("DIAMOND3_FACET_FUNCTIONS", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]))
func (diamond3Facet *Diamond3Facet) PackDefaultDiamondFacetCut() []byte {
	enc, err := diamond3Facet.abi.Pack("defaultDiamondFacetCut")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultDiamondFacetCut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1bd8f7f1.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]))
func (diamond3Facet *Diamond3Facet) UnpackDefaultDiamondFacetCut(data []byte) (IDiamondCutFacetCut, error) {
	out, err := diamond3Facet.abi.Unpack("defaultDiamondFacetCut", data)
	if err != nil {
		return *new(IDiamondCutFacetCut), err
	}
	out0 := *abi.ConvertType(out[0], new(IDiamondCutFacetCut)).(*IDiamondCutFacetCut)
	return out0, err
}

// PackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (diamond3Facet *Diamond3Facet) PackDiamondCut(diamondCut []IDiamondCutFacetCut, init common.Address, calldata []byte) []byte {
	enc, err := diamond3Facet.abi.Pack("diamondCut", diamondCut, init, calldata)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (diamond3Facet *Diamond3Facet) PackFacetAddress(functionSelector [4]byte) []byte {
	enc, err := diamond3Facet.abi.Pack("facetAddress", functionSelector)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFacetAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (diamond3Facet *Diamond3Facet) UnpackFacetAddress(data []byte) (common.Address, error) {
	out, err := diamond3Facet.abi.Unpack("facetAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (diamond3Facet *Diamond3Facet) PackFacetAddresses() []byte {
	enc, err := diamond3Facet.abi.Pack("facetAddresses")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFacetAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (diamond3Facet *Diamond3Facet) UnpackFacetAddresses(data []byte) ([]common.Address, error) {
	out, err := diamond3Facet.abi.Unpack("facetAddresses", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (diamond3Facet *Diamond3Facet) PackFacetFunctionSelectors(facet common.Address) []byte {
	enc, err := diamond3Facet.abi.Pack("facetFunctionSelectors", facet)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFacetFunctionSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (diamond3Facet *Diamond3Facet) UnpackFacetFunctionSelectors(data []byte) ([][4]byte, error) {
	out, err := diamond3Facet.abi.Unpack("facetFunctionSelectors", data)
	if err != nil {
		return *new([][4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)
	return out0, err
}

// PackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (diamond3Facet *Diamond3Facet) PackFacets() []byte {
	enc, err := diamond3Facet.abi.Pack("facets")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFacets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (diamond3Facet *Diamond3Facet) UnpackFacets(data []byte) ([]IDiamondLoupeFacet, error) {
	out, err := diamond3Facet.abi.Unpack("facets", data)
	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}
	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (diamond3Facet *Diamond3Facet) PackOwner() []byte {
	enc, err := diamond3Facet.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (diamond3Facet *Diamond3Facet) UnpackOwner(data []byte) (common.Address, error) {
	out, err := diamond3Facet.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (diamond3Facet *Diamond3Facet) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := diamond3Facet.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (diamond3Facet *Diamond3Facet) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := diamond3Facet.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (diamond3Facet *Diamond3Facet) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := diamond3Facet.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// Diamond3FacetDiamondCut represents a DiamondCut event raised by the Diamond3Facet contract.
type Diamond3FacetDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const Diamond3FacetDiamondCutEventName = "DiamondCut"

// ContractEventName returns the user-defined event name.
func (Diamond3FacetDiamondCut) ContractEventName() string {
	return Diamond3FacetDiamondCutEventName
}

// UnpackDiamondCutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (diamond3Facet *Diamond3Facet) UnpackDiamondCutEvent(log *types.Log) (*Diamond3FacetDiamondCut, error) {
	event := "DiamondCut"
	if log.Topics[0] != diamond3Facet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Diamond3FacetDiamondCut)
	if len(log.Data) > 0 {
		if err := diamond3Facet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamond3Facet.abi.Events[event].Inputs {
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

// Diamond3FacetDiamondCut0 represents a DiamondCut0 event raised by the Diamond3Facet contract.
type Diamond3FacetDiamondCut0 struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const Diamond3FacetDiamondCut0EventName = "DiamondCut0"

// ContractEventName returns the user-defined event name.
func (Diamond3FacetDiamondCut0) ContractEventName() string {
	return Diamond3FacetDiamondCut0EventName
}

// UnpackDiamondCut0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (diamond3Facet *Diamond3Facet) UnpackDiamondCut0Event(log *types.Log) (*Diamond3FacetDiamondCut0, error) {
	event := "DiamondCut0"
	if log.Topics[0] != diamond3Facet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Diamond3FacetDiamondCut0)
	if len(log.Data) > 0 {
		if err := diamond3Facet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamond3Facet.abi.Events[event].Inputs {
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

// Diamond3FacetOwnershipTransferred represents a OwnershipTransferred event raised by the Diamond3Facet contract.
type Diamond3FacetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Diamond3FacetOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (Diamond3FacetOwnershipTransferred) ContractEventName() string {
	return Diamond3FacetOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (diamond3Facet *Diamond3Facet) UnpackOwnershipTransferredEvent(log *types.Log) (*Diamond3FacetOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != diamond3Facet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Diamond3FacetOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := diamond3Facet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamond3Facet.abi.Events[event].Inputs {
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

// Diamond3FacetOwnershipTransferred0 represents a OwnershipTransferred0 event raised by the Diamond3Facet contract.
type Diamond3FacetOwnershipTransferred0 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Diamond3FacetOwnershipTransferred0EventName = "OwnershipTransferred0"

// ContractEventName returns the user-defined event name.
func (Diamond3FacetOwnershipTransferred0) ContractEventName() string {
	return Diamond3FacetOwnershipTransferred0EventName
}

// UnpackOwnershipTransferred0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (diamond3Facet *Diamond3Facet) UnpackOwnershipTransferred0Event(log *types.Log) (*Diamond3FacetOwnershipTransferred0, error) {
	event := "OwnershipTransferred0"
	if log.Topics[0] != diamond3Facet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Diamond3FacetOwnershipTransferred0)
	if len(log.Data) > 0 {
		if err := diamond3Facet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamond3Facet.abi.Events[event].Inputs {
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
func (diamond3Facet *Diamond3Facet) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], diamond3Facet.abi.Errors["InitializationFunctionReverted"].ID.Bytes()[:4]) {
		return diamond3Facet.UnpackInitializationFunctionRevertedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Diamond3FacetInitializationFunctionReverted represents a InitializationFunctionReverted error raised by the Diamond3Facet contract.
type Diamond3FacetInitializationFunctionReverted struct {
	InitializationContractAddress common.Address
	Calldata                      []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InitializationFunctionReverted(address _initializationContractAddress, bytes _calldata)
func Diamond3FacetInitializationFunctionRevertedErrorID() common.Hash {
	return common.HexToHash("0x192105d7c9705d7758a80e15689daed7a9a10e21cf07f39c489e10b763120253")
}

// UnpackInitializationFunctionRevertedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InitializationFunctionReverted(address _initializationContractAddress, bytes _calldata)
func (diamond3Facet *Diamond3Facet) UnpackInitializationFunctionRevertedError(raw []byte) (*Diamond3FacetInitializationFunctionReverted, error) {
	out := new(Diamond3FacetInitializationFunctionReverted)
	if err := diamond3Facet.abi.UnpackIntoInterface(out, "InitializationFunctionReverted", raw); err != nil {
		return nil, err
	}
	return out, nil
}
