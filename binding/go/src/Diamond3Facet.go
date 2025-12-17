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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.

// Diamond3FacetMetaData contains all meta data concerning the Diamond3Facet contract.
var Diamond3FacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"DIAMOND3_FACET_FUNCTIONS\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDiamondFacetCut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializationContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"InitializationFunctionReverted\",\"type\":\"error\"}]",
	ID:  "Diamond3Facet",
	Bin: "0x6080346101715761010081016001600160401b0381118282101761015d5760409081526307e4c70760e21b8252637a0ed62760e01b60208301526356fe50af60e11b908201526314bbdacb60e21b60608201526366ffd66360e11b60808201526301ffc9a760e01b60a082015263f2fde38b60e01b60c0820152638da5cb5b60e01b60e08201525f805460089182905590811161010b575b505f8080525f5160206128185f395f51905f5291905b600181106100c4576040516126a290816101768239f35b5f5f5b600881106100dc5750838201556001016100ad565b600160208196938695965160e01c9063ffffffff8560051b92831b921b191617940191019093929194506100c7565b5f805260070160031c5f5160206128185f395f51905f52017f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5645b8181106101525750610097565b5f8155600101610145565b634e487b7160e01b5f52604160045260245ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a7146114ac575080631bd8f7f1146111715780631f931c1c1461080257806352ef6b2c146106e1578063549fc1a4146106595780637a0ed627146103965780638da5cb5b14610326578063adfca15e1461023b578063cdffacc6146101975763f2fde38b1461008a575f80fd5b346101935760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576100c161160a565b6100c96119cc565b73ffffffffffffffffffffffffffffffffffffffff807fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c13205416911690817fffffffffffffffffffffffff00000000000000000000000000000000000000007fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c13205416177fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c1320557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b5f80fd5b346101935760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193577fffffffff000000000000000000000000000000000000000000000000000000006101ef61153b565b165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b346101935760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576102c06102bb61027861160a565b73ffffffffffffffffffffffffffffffffffffffff165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f2090565b6116ab565b6040518091602082016020835281518091526020604084019201905f5b8181106102eb575050500390f35b82517fffffffff00000000000000000000000000000000000000000000000000000000168452859450602093840193909201916001016102dd565b34610193575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019357602073ffffffffffffffffffffffffffffffffffffffff7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c13205416604051908152f35b34610193575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193577fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e546103ef81611951565b906103fd604051928361166a565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061042a82611951565b015f5b8181106105f55750507fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e545f5b82811061050b57836040518091602082016020835281518091526040830190602060408260051b8601019301915f905b82821061049957505050500390f35b919360206104fb827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186526040838a5173ffffffffffffffffffffffffffffffffffffffff81511684520151918185820152019061156a565b960192019201859493919261048a565b818110156105c8576001907fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e5f526105a973ffffffffffffffffffffffffffffffffffffffff8260205f200154168061056484896119b8565b515273ffffffffffffffffffffffffffffffffffffffff165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f2090565b6105c060206105b884896119b8565b5101916116ab565b90520161045a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60405190604082019180831067ffffffffffffffff84111761062c576020926040525f81526060838201528282870101520161042d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b346101935760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576004355f54811015610193575f548110156105c8576020905f805260e0825f208260031c01549160051b161c60e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b34610193575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576040518060207fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e54928381520180927fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e5f5260205f20905f5b8181106107d6575050508161077e91038261166a565b604051918291602083019060208452518091526040830191905f5b8181106107a7575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610799565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201610768565b346101935760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101935760043567ffffffffffffffff811161019357366023820112156101935780600401359067ffffffffffffffff82116101935760248260051b8201013681116101935760243573ffffffffffffffffffffffffffffffffffffffff811692838203610193576044359267ffffffffffffffff841161019357366023850112156101935783600401359067ffffffffffffffff8211610193573660248387010111610193576108e29695949392966119cc565b6108eb86611951565b956108f9604051978861166a565b86526020860190819760248101925b82841061103357505050506020818060246109235f95611969565b96610931604051988961166a565b8288520183870137840101525f935b8051851015610f7257602061095586836119b8565b5101516003811015610f455780610b56575073ffffffffffffffffffffffffffffffffffffffff61098686836119b8565b51511693604061099687846119b8565b510151936109a685511515611ad3565b6109b1861515611b5e565b6bffffffffffffffffffffffff610a058773ffffffffffffffffffffffffffffffffffffffff165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f2090565b5416938415610b48575b5f945b8651861015610b30577fffffffff00000000000000000000000000000000000000000000000000000000610a4687896119b8565b5116805f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205273ffffffffffffffffffffffffffffffffffffffff60405f205416610aac5781610a9f8a600194610aa4946124bb565b611be9565b950194610a12565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f60448201527f6e207468617420616c72656164792065786973747300000000000000000000006064820152fd5b5094509450946001919692505b019392919094610940565b610b5187612387565b610a0f565b93959360018103610d52575073ffffffffffffffffffffffffffffffffffffffff610b8186836119b8565b515116956040610b9187846119b8565b51015193610ba185511515611ad3565b610bac881515611b5e565b6bffffffffffffffffffffffff610c008973ffffffffffffffffffffffffffffffffffffffff165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f2090565b5416938415610d44575b5f945b8651861015610d33577fffffffff00000000000000000000000000000000000000000000000000000000610c4187896119b8565b5116805f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205273ffffffffffffffffffffffffffffffffffffffff60405f205416918b8314610caf57610a9f8c8284610ca2610ca796600198611d66565b6124bb565b950194610c0d565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e00000000000000006064820152fd5b509450949250946001919650610b3d565b610d4d89612387565b610c0a565b93959193909290600203610ec15773ffffffffffffffffffffffffffffffffffffffff610d7f86856119b8565b515116936040610d8f87866119b8565b51015194610d9f86511515611ad3565b610e3d575f5b8551811015610e2e5780610e287fffffffff00000000000000000000000000000000000000000000000000000000610ddf6001948a6119b8565b5116805f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205273ffffffffffffffffffffffffffffffffffffffff60405f205416611d66565b01610da5565b50929591946001919450610b3d565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f7665206661636574206164647260448201527f657373206d7573742062652061646472657373283029000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f7272656374204661636574437560448201527f74416374696f6e000000000000000000000000000000000000000000000000006064820152fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9285906040519460608601906060875251809152608086019060808160051b88010193915f905b828210610fea57610fe888887f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb6738c80610fe08c8c6020840152828103604084015286611a90565b0390a1611c3d565b005b90919294602080611025837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808d6001960301865289516115bf565b970192019201909291610f99565b839995969798993567ffffffffffffffff811161019357820160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261019357604051906110858261164e565b602481013573ffffffffffffffffffffffffffffffffffffffff8116810361019357825260448101356003811015610193576020830152606481013567ffffffffffffffff811161019357602491010136601f820112156101935780356110eb81611951565b916110f9604051938461166a565b81835260208084019260051b8201019036821161019357602001915b818310611138575050506040820152815294989796959460209384019301610908565b82357fffffffff000000000000000000000000000000000000000000000000000000008116810361019357815260209283019201611115565b34610193575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576060604080516111ae8161164e565b5f81525f602082015201526040516111c58161164e565b3081525f60208083018290526040518254808252838052918101939290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b816007840110611425579461126292849261127b9754918181106113ef575b8181106113b9575b818110611383575b81811061134d575b818110611317575b8181106112e1575b8181106112ac575b1061127f575b50038261166a565b60408201526040519182916020835260208301906115bf565b0390f35b7fffffffff000000000000000000000000000000000000000000000000000000001681526020018661125a565b9260206001917fffffffff0000000000000000000000000000000000000000000000000000000085831b168152019301611254565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560401b16815201930161124c565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560601b168152019301611244565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560801b16815201930161123c565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560a01b168152019301611234565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560c01b16815201930161122c565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560e01b168152019301611224565b9460016101006008927fffffffff000000000000000000000000000000000000000000000000000000008954818160e01b168352818160c01b166020840152818160a01b166040840152818160801b166060840152818160601b166080840152818160401b1660a0840152818160201b1660c08401521660e0820152019601920191611205565b346101935760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610193576020907fffffffff0000000000000000000000000000000000000000000000000000000061150761153b565b165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f825260ff60405f20541615158152f35b600435907fffffffff000000000000000000000000000000000000000000000000000000008216820361019357565b90602080835192838152019201905f5b8181106115875750505090565b82517fffffffff000000000000000000000000000000000000000000000000000000001684526020938401939092019160010161157a565b9073ffffffffffffffffffffffffffffffffffffffff82511681526020820151916003831015610f45576040606091611607946020850152015191816040820152019061156a565b90565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019357565b91909180548310156105c8575f52601c60205f208360031c019260021b1690565b6060810190811067ffffffffffffffff82111761062c57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761062c57604052565b90604051918281549182825260208201905f5260205f20925f905b8060078301106118c45761171c94549181811061188e575b818110611858575b818110611822575b8181106117ec575b8181106117b6575b818110611780575b81811061174b575b1061171e575b50038361166a565b565b7fffffffff000000000000000000000000000000000000000000000000000000001681526020015f611714565b9260206001917fffffffff0000000000000000000000000000000000000000000000000000000085831b16815201930161170e565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560401b168152019301611706565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560601b1681520193016116fe565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560801b1681520193016116f6565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560a01b1681520193016116ee565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560c01b1681520193016116e6565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560e01b1681520193016116de565b9160089193506101006001917fffffffff000000000000000000000000000000000000000000000000000000008754818160e01b168352818160c01b166020840152818160a01b166040840152818160801b166060840152818160601b166080840152818160401b1660a0840152818160201b1660c08401521660e08201520194019201859293916116c6565b67ffffffffffffffff811161062c5760051b60200190565b67ffffffffffffffff811161062c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b80548210156105c8575f5260205f2001905f90565b80518210156105c85760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132054163303611a0c57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b15611ada57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660448201527f6163657420746f206375740000000000000000000000000000000000000000006064820152fd5b15611b6557565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201527f65206164647265737328302900000000000000000000000000000000000000006064820152fd5b6bffffffffffffffffffffffff166bffffffffffffffffffffffff8114611c105760010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9073ffffffffffffffffffffffffffffffffffffffff8216918215611d61575f8091611cc6604051611c7060608261166a565b602881527f4c69624469616d6f6e644375743a205f696e697420616464726573732068617360208201527f206e6f20636f6465000000000000000000000000000000000000000000000000604082015282612625565b83519060208501905af4913d15611d59573d92611ce284611969565b93611cf0604051958661166a565b84523d5f602086013e5b15611d0457505050565b825115611d1357825160208401fd5b611d556040519283927f192105d70000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190611a90565b0390fd5b606092611cfa565b505050565b73ffffffffffffffffffffffffffffffffffffffff169081156123035730821461227f577fffffffff0000000000000000000000000000000000000000000000000000000016805f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205260405f205460a01c90825f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f2054917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301928311611c1057828103612151575b50825f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f208054801561204e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190611e9c828261162d565b63ffffffff82549160031b1b19169055555f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6020525f604081205515611ee15750565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111611c1057815f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d602052600160405f2001549080820361207b575b50507fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e54801561204e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611fdb817fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e6119a3565b73ffffffffffffffffffffffffffffffffffffffff82549160031b1b191690557fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e555f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d6020525f6001604082200155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b6120ba73ffffffffffffffffffffffffffffffffffffffff917fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e6119a3565b90549060031b1c1661211c816120f0847fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e6119a3565b90919073ffffffffffffffffffffffffffffffffffffffff8084549260031b9316831b921b1916179055565b5f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d602052600160405f2001555f80611f65565b61227990845f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d6020527fffffffff000000000000000000000000000000000000000000000000000000006121aa8560405f2061162d565b90549060031b1c60e01b865f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d602052612209816121ec8560405f2061162d565b90919063ffffffff83549160031b9260e01c831b921b1916179055565b165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205260405f209073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b5f611e3b565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201527f7461626c652066756e6374696f6e0000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e27742065786973740000000000000000006064820152fd5b6123ee60405161239860608261166a565b602481527f4c69624469616d6f6e644375743a204e657720666163657420686173206e6f2060208201527f636f646500000000000000000000000000000000000000000000000000000000604082015282612625565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e549073ffffffffffffffffffffffffffffffffffffffff81165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205281600160405f2001556801000000000000000082101561062c576120f082600161171c94017fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e557fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e6119a3565b6125527fffffffff00000000000000000000000000000000000000000000000000000000821692835f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205260405f209073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b73ffffffffffffffffffffffffffffffffffffffff83165f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60205260405f20908154916801000000000000000083101561062c57826121ec9160016125bb9501815561162d565b5f527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c60205273ffffffffffffffffffffffffffffffffffffffff60405f2091167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b3b1561262e5750565b611d55906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611a9056fea2646970667358221220b112aca32aa3ae363004de91c10983359952c135a594a0a80d0625c38104d29f64736f6c634300081e0033290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
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
// the contract method with ID 0x549fc1a4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DIAMOND3_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (diamond3Facet *Diamond3Facet) PackDIAMOND3FACETFUNCTIONS(arg0 *big.Int) []byte {
	enc, err := diamond3Facet.abi.Pack("DIAMOND3_FACET_FUNCTIONS", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDIAMOND3FACETFUNCTIONS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x549fc1a4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DIAMOND3_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (diamond3Facet *Diamond3Facet) TryPackDIAMOND3FACETFUNCTIONS(arg0 *big.Int) ([]byte, error) {
	return diamond3Facet.abi.Pack("DIAMOND3_FACET_FUNCTIONS", arg0)
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
	return out0, nil
}

// PackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]))
func (diamond3Facet *Diamond3Facet) PackDefaultDiamondFacetCut() []byte {
	enc, err := diamond3Facet.abi.Pack("defaultDiamondFacetCut")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]))
func (diamond3Facet *Diamond3Facet) TryPackDefaultDiamondFacetCut() ([]byte, error) {
	return diamond3Facet.abi.Pack("defaultDiamondFacetCut")
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
	return out0, nil
}

// PackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (diamond3Facet *Diamond3Facet) PackDiamondCut(diamondCut []IDiamondCutFacetCut, init common.Address, calldata []byte) []byte {
	enc, err := diamond3Facet.abi.Pack("diamondCut", diamondCut, init, calldata)
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
func (diamond3Facet *Diamond3Facet) TryPackDiamondCut(diamondCut []IDiamondCutFacetCut, init common.Address, calldata []byte) ([]byte, error) {
	return diamond3Facet.abi.Pack("diamondCut", diamondCut, init, calldata)
}

// PackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (diamond3Facet *Diamond3Facet) PackFacetAddress(functionSelector [4]byte) []byte {
	enc, err := diamond3Facet.abi.Pack("facetAddress", functionSelector)
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
func (diamond3Facet *Diamond3Facet) TryPackFacetAddress(functionSelector [4]byte) ([]byte, error) {
	return diamond3Facet.abi.Pack("facetAddress", functionSelector)
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
	return out0, nil
}

// PackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (diamond3Facet *Diamond3Facet) PackFacetAddresses() []byte {
	enc, err := diamond3Facet.abi.Pack("facetAddresses")
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
func (diamond3Facet *Diamond3Facet) TryPackFacetAddresses() ([]byte, error) {
	return diamond3Facet.abi.Pack("facetAddresses")
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
	return out0, nil
}

// PackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (diamond3Facet *Diamond3Facet) PackFacetFunctionSelectors(facet common.Address) []byte {
	enc, err := diamond3Facet.abi.Pack("facetFunctionSelectors", facet)
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
func (diamond3Facet *Diamond3Facet) TryPackFacetFunctionSelectors(facet common.Address) ([]byte, error) {
	return diamond3Facet.abi.Pack("facetFunctionSelectors", facet)
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
	return out0, nil
}

// PackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (diamond3Facet *Diamond3Facet) PackFacets() []byte {
	enc, err := diamond3Facet.abi.Pack("facets")
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
func (diamond3Facet *Diamond3Facet) TryPackFacets() ([]byte, error) {
	return diamond3Facet.abi.Pack("facets")
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
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address owner_)
func (diamond3Facet *Diamond3Facet) PackOwner() []byte {
	enc, err := diamond3Facet.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address owner_)
func (diamond3Facet *Diamond3Facet) TryPackOwner() ([]byte, error) {
	return diamond3Facet.abi.Pack("owner")
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
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (diamond3Facet *Diamond3Facet) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := diamond3Facet.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (diamond3Facet *Diamond3Facet) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return diamond3Facet.abi.Pack("supportsInterface", interfaceId)
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
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (diamond3Facet *Diamond3Facet) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := diamond3Facet.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (diamond3Facet *Diamond3Facet) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return diamond3Facet.abi.Pack("transferOwnership", newOwner)
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
	if len(log.Topics) == 0 || log.Topics[0] != diamond3Facet.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != diamond3Facet.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != diamond3Facet.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != diamond3Facet.abi.Events[event].ID {
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
