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

// BaseForgeFacetMetaData contains all meta data concerning the BaseForgeFacet contract.
var BaseForgeFacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BASEFORGE_FACET_FUNCTIONS\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDiamondFacetCut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut\",\"name\":\"facetCut\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validator_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"BaseForge__InvalidFeeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidPermitSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"}]",
	ID:  "BaseForgeFacet",
	Bin: "0x608080604052346101da575f5160206118245f395f51905f525460ff8160401c166101cb576002600160401b03196001600160401b03821601610178575b505060405160a081016001600160401b038111828210176101645760409081526342580cb760e11b8252623f675f60e91b6020830152633644e51560e01b90820152633a5381b560e01b6060820152630264fa7b60e31b60808201525f8054600591829055908111610102575b505f808052905f5160206118045f395f51905f5290825b600581106100db5750505560405161162590816101df8239f35b90926020600191855160e01c9063ffffffff8560051b92831b921b191617940191016100c1565b5f80525f5160206118045f395f51905f5280546001600160a01b0316815560079190910160031c017f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5645b81811061015957506100aa565b5f815560010161014c565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0319166001600160401b039081175f5160206118245f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8061003d565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c9081631327d3d814611042575080631bd8f7f114610c555780631c29388514610bb95780633644e51514610b1e5780633a5381b514610aae5780636910e3341461029f5780637ecebe001461021d576384b0196e14610074575f80fd5b34610219575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005415806101f0575b15610192576101366100db611259565b6100e361136c565b6020610144604051926100f683856111c7565b5f84525f3681376040519586957f0f00000000000000000000000000000000000000000000000000000000000000875260e08588015260e0870190611168565b908582036040870152611168565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b82811061017b57505050500390f35b83518552869550938101939281019260010161016c565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154156100cb565b5f80fd5b346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195773ffffffffffffffffffffffffffffffffffffffff610269611145565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b346102195760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760243573ffffffffffffffffffffffffffffffffffffffff811690818103610219577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549060ff8260401c16159167ffffffffffffffff811680159081610aa6575b6001149081610a9c575b159081610a93575b50610a6b578260017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055610a16575b5061039e6114e9565b6103a66114e9565b604051926004356020850152602084526103c16040856111c7565b604051936040850185811067ffffffffffffffff82111761087257604052600185527f310000000000000000000000000000000000000000000000000000000000000060208601526104116114e9565b6104196114e9565b80519067ffffffffffffffff82116108725781906104577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10254611208565b601f8111610989575b50602090601f83116001146108aa575f9261089f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102555b835167ffffffffffffffff8111610872576105037fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10354611208565b601f81116107f0575b50602094601f8211600114610712579481929394955f92610707575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103555b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100555f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101556105c86114e9565b6105d06114e9565b156106df5761064690337fffffffffffffffffffffffff00000000000000000000000000000000000000007f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416177f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005561143f565b61064c57005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f95d18bbf000000000000000000000000000000000000000000000000000000005f5260045ffd5b015190508580610528565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216957fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f52805f20915f5b8881106107d8575083600195969798106107a1575b505050811b017fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035561057a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055858080610774565b9192602060018192868501518155019401920161075f565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75601f830160051c81019160208410610868575b601f0160051c01905b81811061085d575061050c565b5f8155600101610850565b9091508190610847565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b015190508680610476565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f52815f20925f5b818110610971575090846001959493921061093a575b505050811b017fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102556104c8565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905586808061090d565b929360206001819287860151815501950193016108f7565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f529091507f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d601f840160051c81019160208510610a0c575b90601f859493920160051c01905b8181106109fe5750610460565b5f81558493506001016109f1565b90915081906109e3565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005583610395565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b90501585610342565b303b15915061033a565b849150610330565b34610219575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957602073ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015416604051908152f35b34610219575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219576020610b56611540565b610b5e6115aa565b60405190838201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152610bae60c0826111c7565b519020604051908152f35b346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219576004355f54811015610219576020905f80527fffffffff000000000000000000000000000000000000000000000000000000008160031c7f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563015460e06040519360051b161c60e01b168152f35b34610219575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957606060408051610c92816111ab565b5f81525f60208201520152604051610ca9816111ab565b30815260208101905f8252604051808160205f5492838152015f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563925f905b806007830110610fb557610d40945491818110610f7f575b818110610f49575b818110610f13575b818110610edd575b818110610ea7575b818110610e71575b818110610e3c575b10610e0f575b5003826111c7565b60408201908152604051926020845273ffffffffffffffffffffffffffffffffffffffff608085019351166020850152516003811015610de257604084015251606080840152805191829052829160a0830191602001905f5b818110610da7575050500390f35b82517fffffffff0000000000000000000000000000000000000000000000000000000016845285945060209384019390920191600101610d99565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7fffffffff0000000000000000000000000000000000000000000000000000000016815260200186610d38565b9260206001917fffffffff0000000000000000000000000000000000000000000000000000000085831b168152019301610d32565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560401b168152019301610d2a565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560601b168152019301610d22565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560801b168152019301610d1a565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560a01b168152019301610d12565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560c01b168152019301610d0a565b9260206001917fffffffff000000000000000000000000000000000000000000000000000000008560e01b168152019301610d02565b9160089193506101006001917fffffffff000000000000000000000000000000000000000000000000000000008754818160e01b168352818160c01b166020840152818160a01b166040840152818160801b166060840152818160601b166080840152818160401b1660a0840152818160201b1660c08401521660e0820152019401920184929391610cea565b346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957611079611145565b9073ffffffffffffffffffffffffffffffffffffffff7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c1320541633036110c3576110c18261143f565b005b807f08c379a0000000000000000000000000000000000000000000000000000000006084925260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361021957565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6060810190811067ffffffffffffffff82111761087257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761087257604052565b90600182811c9216801561124f575b602083101461122257565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691611217565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102549161128b83611208565b808352926001811690811561132f57506001146112b1575b6112af925003836111c7565b565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106113135750509060206112af928201016112a3565b60209193508060019154838589010152019101909184926112fb565b602092506112af9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016112a3565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103549161139e83611208565b808352926001811690811561132f57506001146113c1576112af925003836111c7565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b8183106114235750509060206112af928201016112a3565b602091935080600191548385890101520191019091849261140b565b73ffffffffffffffffffffffffffffffffffffffff1680156106df57807fffffffffffffffffffffffff00000000000000000000000000000000000000007f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015416177f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca01557fb3a3a56265020415cf2f7ff198e2052a6e1d43d7eb127450af725829e40e08c25f80a2565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561151857565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b611548611259565b8051908115611558576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005480156115855790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6115b261136c565b80519081156115c2576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015611585579056fea2646970667358221220a0d41183da3ce3363517ddac2d62ca7e4be3a9da5925a5c236cead4eaf50fc4664736f6c634300081e0033290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BaseForgeFacet is an auto generated Go binding around an Ethereum contract.
type BaseForgeFacet struct {
	abi abi.ABI
}

// NewBaseForgeFacet creates a new instance of BaseForgeFacet.
func NewBaseForgeFacet() *BaseForgeFacet {
	parsed, err := BaseForgeFacetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BaseForgeFacet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BaseForgeFacet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBASEFORGEFACETFUNCTIONS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c293885.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BASEFORGE_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (baseForgeFacet *BaseForgeFacet) PackBASEFORGEFACETFUNCTIONS(arg0 *big.Int) []byte {
	enc, err := baseForgeFacet.abi.Pack("BASEFORGE_FACET_FUNCTIONS", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBASEFORGEFACETFUNCTIONS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c293885.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BASEFORGE_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (baseForgeFacet *BaseForgeFacet) TryPackBASEFORGEFACETFUNCTIONS(arg0 *big.Int) ([]byte, error) {
	return baseForgeFacet.abi.Pack("BASEFORGE_FACET_FUNCTIONS", arg0)
}

// UnpackBASEFORGEFACETFUNCTIONS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c293885.
//
// Solidity: function BASEFORGE_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (baseForgeFacet *BaseForgeFacet) UnpackBASEFORGEFACETFUNCTIONS(data []byte) ([4]byte, error) {
	out, err := baseForgeFacet.abi.Unpack("BASEFORGE_FACET_FUNCTIONS", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (baseForgeFacet *BaseForgeFacet) PackDOMAINSEPARATOR() []byte {
	enc, err := baseForgeFacet.abi.Pack("DOMAIN_SEPARATOR")
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
func (baseForgeFacet *BaseForgeFacet) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return baseForgeFacet.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (baseForgeFacet *BaseForgeFacet) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := baseForgeFacet.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]) facetCut)
func (baseForgeFacet *BaseForgeFacet) PackDefaultDiamondFacetCut() []byte {
	enc, err := baseForgeFacet.abi.Pack("defaultDiamondFacetCut")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]) facetCut)
func (baseForgeFacet *BaseForgeFacet) TryPackDefaultDiamondFacetCut() ([]byte, error) {
	return baseForgeFacet.abi.Pack("defaultDiamondFacetCut")
}

// UnpackDefaultDiamondFacetCut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1bd8f7f1.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]) facetCut)
func (baseForgeFacet *BaseForgeFacet) UnpackDefaultDiamondFacetCut(data []byte) (IDiamondCutFacetCut, error) {
	out, err := baseForgeFacet.abi.Unpack("defaultDiamondFacetCut", data)
	if err != nil {
		return *new(IDiamondCutFacetCut), err
	}
	out0 := *abi.ConvertType(out[0], new(IDiamondCutFacetCut)).(*IDiamondCutFacetCut)
	return out0, nil
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (baseForgeFacet *BaseForgeFacet) PackEip712Domain() []byte {
	enc, err := baseForgeFacet.abi.Pack("eip712Domain")
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
func (baseForgeFacet *BaseForgeFacet) TryPackEip712Domain() ([]byte, error) {
	return baseForgeFacet.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (baseForgeFacet *BaseForgeFacet) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := baseForgeFacet.abi.Unpack("eip712Domain", data)
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

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6910e334.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(bytes32 service_, address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) PackInitialize(service [32]byte, validator common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("initialize", service, validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6910e334.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(bytes32 service_, address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) TryPackInitialize(service [32]byte, validator common.Address) ([]byte, error) {
	return baseForgeFacet.abi.Pack("initialize", service, validator)
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (baseForgeFacet *BaseForgeFacet) PackNonces(owner common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("nonces", owner)
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
func (baseForgeFacet *BaseForgeFacet) TryPackNonces(owner common.Address) ([]byte, error) {
	return baseForgeFacet.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (baseForgeFacet *BaseForgeFacet) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := baseForgeFacet.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSetValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1327d3d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setValidator(address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) PackSetValidator(validator common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("setValidator", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1327d3d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setValidator(address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) TryPackSetValidator(validator common.Address) ([]byte, error) {
	return baseForgeFacet.abi.Pack("setValidator", validator)
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validator() view returns(address)
func (baseForgeFacet *BaseForgeFacet) PackValidator() []byte {
	enc, err := baseForgeFacet.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validator() view returns(address)
func (baseForgeFacet *BaseForgeFacet) TryPackValidator() ([]byte, error) {
	return baseForgeFacet.abi.Pack("validator")
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (baseForgeFacet *BaseForgeFacet) UnpackValidator(data []byte) (common.Address, error) {
	out, err := baseForgeFacet.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// BaseForgeFacetEIP712DomainChanged represents a EIP712DomainChanged event raised by the BaseForgeFacet contract.
type BaseForgeFacetEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const BaseForgeFacetEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (BaseForgeFacetEIP712DomainChanged) ContractEventName() string {
	return BaseForgeFacetEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (baseForgeFacet *BaseForgeFacet) UnpackEIP712DomainChangedEvent(log *types.Log) (*BaseForgeFacetEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if len(log.Topics) == 0 || log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseForgeFacetEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := baseForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseForgeFacet.abi.Events[event].Inputs {
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

// BaseForgeFacetInitialized represents a Initialized event raised by the BaseForgeFacet contract.
type BaseForgeFacetInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const BaseForgeFacetInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (BaseForgeFacetInitialized) ContractEventName() string {
	return BaseForgeFacetInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (baseForgeFacet *BaseForgeFacet) UnpackInitializedEvent(log *types.Log) (*BaseForgeFacetInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseForgeFacetInitialized)
	if len(log.Data) > 0 {
		if err := baseForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseForgeFacet.abi.Events[event].Inputs {
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

// BaseForgeFacetValidatorUpdated represents a ValidatorUpdated event raised by the BaseForgeFacet contract.
type BaseForgeFacetValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const BaseForgeFacetValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (BaseForgeFacetValidatorUpdated) ContractEventName() string {
	return BaseForgeFacetValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (baseForgeFacet *BaseForgeFacet) UnpackValidatorUpdatedEvent(log *types.Log) (*BaseForgeFacetValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseForgeFacetValidatorUpdated)
	if len(log.Data) > 0 {
		if err := baseForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseForgeFacet.abi.Events[event].Inputs {
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
func (baseForgeFacet *BaseForgeFacet) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeInvalidFeeData"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeInvalidFeeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeInvalidPermitSignatureLength"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeInvalidPermitSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackNotInitializingError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// BaseForgeFacetBaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func BaseForgeFacetBaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeExpiredSignatureError(raw []byte) (*BaseForgeFacetBaseForgeExpiredSignature, error) {
	out := new(BaseForgeFacetBaseForgeExpiredSignature)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetBaseForgeInvalidFeeData represents a BaseForge__InvalidFeeData error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeInvalidFeeData struct {
	FeeRecipient common.Address
	FeeBPS       *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func BaseForgeFacetBaseForgeInvalidFeeDataErrorID() common.Hash {
	return common.HexToHash("0xc028e60c3834e172f648e56b5588dbdbb2caa156ad1931ff196c489c710a7363")
}

// UnpackBaseForgeInvalidFeeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeInvalidFeeDataError(raw []byte) (*BaseForgeFacetBaseForgeInvalidFeeData, error) {
	out := new(BaseForgeFacetBaseForgeInvalidFeeData)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeInvalidFeeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetBaseForgeInvalidPermitSignatureLength represents a BaseForge__InvalidPermitSignatureLength error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeInvalidPermitSignatureLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func BaseForgeFacetBaseForgeInvalidPermitSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xd114b2291dcfec2182b0d6c24f580b96bcae4149c53fa4278ac88775ec7bb032")
}

// UnpackBaseForgeInvalidPermitSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeInvalidPermitSignatureLengthError(raw []byte) (*BaseForgeFacetBaseForgeInvalidPermitSignatureLength, error) {
	out := new(BaseForgeFacetBaseForgeInvalidPermitSignatureLength)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeInvalidPermitSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetBaseForgeInvalidValidatorSignature represents a BaseForge__InvalidValidatorSignature error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func BaseForgeFacetBaseForgeInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0xc0299d906fef6fff02cc5dc92034a345928d480697ea46384b87962e5a4e2b11")
}

// UnpackBaseForgeInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeInvalidValidatorSignatureError(raw []byte) (*BaseForgeFacetBaseForgeInvalidValidatorSignature, error) {
	out := new(BaseForgeFacetBaseForgeInvalidValidatorSignature)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetBaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func BaseForgeFacetBaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeZeroAddressError(raw []byte) (*BaseForgeFacetBaseForgeZeroAddress, error) {
	out := new(BaseForgeFacetBaseForgeZeroAddress)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetInvalidAccountNonce represents a InvalidAccountNonce error raised by the BaseForgeFacet contract.
type BaseForgeFacetInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func BaseForgeFacetInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (baseForgeFacet *BaseForgeFacet) UnpackInvalidAccountNonceError(raw []byte) (*BaseForgeFacetInvalidAccountNonce, error) {
	out := new(BaseForgeFacetInvalidAccountNonce)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetInvalidInitialization represents a InvalidInitialization error raised by the BaseForgeFacet contract.
type BaseForgeFacetInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func BaseForgeFacetInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (baseForgeFacet *BaseForgeFacet) UnpackInvalidInitializationError(raw []byte) (*BaseForgeFacetInvalidInitialization, error) {
	out := new(BaseForgeFacetInvalidInitialization)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BaseForgeFacetNotInitializing represents a NotInitializing error raised by the BaseForgeFacet contract.
type BaseForgeFacetNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func BaseForgeFacetNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (baseForgeFacet *BaseForgeFacet) UnpackNotInitializingError(raw []byte) (*BaseForgeFacetNotInitializing, error) {
	out := new(BaseForgeFacetNotInitializing)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}
