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
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// IMulticall3Call is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call struct {
	Target   common.Address
	CallData []byte
}

// IMulticall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// IMulticall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// IMulticall3Result is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Result struct {
	Success    bool
	ReturnData []byte
}

// BaseForgeFacetMetaData contains all meta data concerning the BaseForgeFacet contract.
var BaseForgeFacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BASEFORGE_FACET_FUNCTIONS\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDiamondFacetCut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut\",\"name\":\"facetCut\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validator_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"}]",
	ID:  "BaseForgeFacet",
	Bin: "0x608060405234801561000f575f5ffd5b50604080516080810182526342580cb760e11b8152623f675f60e91b6020820152633644e51560e01b91810191909152633a5381b560e01b6060820152610059905f90600461007e565b5061013b565b81548152906001019060200180831161005f5750939695505050505050565b828054828255905f5260205f2090600701600890048101928215610117579160200282015f5b838211156100e557835183826101000a81548163ffffffff021916908360e01c021790555092602001926004016020816003010492830192600103026100a4565b80156101155782816101000a81549063ffffffff02191690556004016020816003010492830192600103026100e5565b505b50610123929150610127565b5090565b5b80821115610123575f8155600101610128565b61117b806101485f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80633a5381b5116100585780633a5381b5146101165780637ab4339d1461015d5780637ecebe001461017057806384b0196e146101c4575f5ffd5b80631327d3d8146100895780631bd8f7f11461009e5780631c293885146100bc5780633644e51514610100575b5f5ffd5b61009c610097366004610c4e565b6101df565b005b6100a66101f3565b6040516100b39190610c6e565b60405180910390f35b6100cf6100ca366004610d41565b6102b6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016100b3565b6101086102e9565b6040519081526020016100b3565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b3565b61009c61016b366004610d85565b610332565b61010861017e366004610c4e565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b6101cc6104f8565b6040516100b39796959493929190610ed2565b6101e76105f7565b6101f0816106c2565b50565b60408051606080820183525f808352602083015291810191909152604080516060810182523081525f602080830182905281548451818302810183018652818152939485019392918301828280156102a957602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116102565790505b5050505050815250905090565b5f81815481106102c4575f80fd5b905f5260205f209060089182820401919006600402915054906101000a900460e01b81565b5f6102f26107c0565b905090565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015473ffffffffffffffffffffffffffffffffffffffff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561037c5750825b90505f8267ffffffffffffffff1660011480156103985750303b155b9050811580156103a6575080155b156103dd576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561043e5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b61044887876107c9565b83156104a95784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561053657506001810154155b6105a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b6105a9610825565b6105b16108f8565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6004015473ffffffffffffffffffffffffffffffffffffffff1633146106c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152608401610598565b565b73ffffffffffffffffffffffffffffffffffffffff811661070f576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060091907fb3a3a56265020415cf2f7ff198e2052a6e1d43d7eb127450af725829e40e08c2905f90a25050565b5f6102f2610949565b6107d16109bc565b610810826040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250610a23565b610818610a35565b61082181610a3d565b5050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161087690610f91565b80601f01602080910402602001604051908101604052809291908181526020018280546108a290610f91565b80156108ed5780601f106108c4576101008083540402835291602001916108ed565b820191905f5260205f20905b8154815290600101906020018083116108d057829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161087690610f91565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610973610ae4565b61097b610b5f565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166106c0576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a2b6109bc565b6108218282610bb4565b6106c06109bc565b610a456109bc565b73ffffffffffffffffffffffffffffffffffffffff8116610a92576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633178155610821826106c2565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081610b0f610825565b805190915015610b2757805160209091012092915050565b81548015610b36579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081610b8a6108f8565b805190915015610ba257805160209091012092915050565b60018201548015610b36579392505050565b610bbc6109bc565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102610c08848261102e565b5060038101610c17838261102e565b505f8082556001909101555050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c49575f5ffd5b919050565b5f60208284031215610c5e575f5ffd5b610c6782610c26565b9392505050565b602081525f6080820173ffffffffffffffffffffffffffffffffffffffff8451166020840152602084015160038110610cce577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b604084810191909152840151606080850152805191829052602001905f9060a08501905b80831015610d37577fffffffff000000000000000000000000000000000000000000000000000000008451168252602082019150602084019350600183019250610cf2565b5095945050505050565b5f60208284031215610d51575f5ffd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215610d96575f5ffd5b823567ffffffffffffffff811115610dac575f5ffd5b8301601f81018513610dbc575f5ffd5b803567ffffffffffffffff811115610dd657610dd6610d58565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715610e4257610e42610d58565b604052818152828201602001871015610e59575f5ffd5b816020840160208301375f60208383010152809450505050610e7d60208401610c26565b90509250929050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f610f0c60e0830189610e86565b8281036040840152610f1e8189610e86565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015610f80578351835260209384019390920191600101610f62565b50909b9a5050505050505050505050565b600181811c90821680610fa557607f821691505b602082108103610fdc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f82111561102957805f5260205f20601f840160051c810160208510156110075750805b601f840160051c820191505b81811015611026575f8155600101611013565b50505b505050565b815167ffffffffffffffff81111561104857611048610d58565b61105c816110568454610f91565b84610fe2565b6020601f8211600181146110ad575f83156110775750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611026565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156110fa57878501518255602094850194600190920191016110da565b508482101561113657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b0190555056fea26469706673582212201aa5dae2afa7ac6c588e747a7a1456d072bb78383bf2c13e997f1d47a2d6629d64736f6c634300081c0033",
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
// the contract method with ID 0x1c293885.
//
// Solidity: function BASEFORGE_FACET_FUNCTIONS(uint256 ) view returns(bytes4)
func (baseForgeFacet *BaseForgeFacet) PackBASEFORGEFACETFUNCTIONS(arg0 *big.Int) []byte {
	enc, err := baseForgeFacet.abi.Pack("BASEFORGE_FACET_FUNCTIONS", arg0)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (baseForgeFacet *BaseForgeFacet) PackDOMAINSEPARATOR() []byte {
	enc, err := baseForgeFacet.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackDefaultDiamondFacetCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bd8f7f1.
//
// Solidity: function defaultDiamondFacetCut() view returns((address,uint8,bytes4[]) facetCut)
func (baseForgeFacet *BaseForgeFacet) PackDefaultDiamondFacetCut() []byte {
	enc, err := baseForgeFacet.abi.Pack("defaultDiamondFacetCut")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (baseForgeFacet *BaseForgeFacet) PackEip712Domain() []byte {
	enc, err := baseForgeFacet.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

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
	return *outstruct, err

}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ab4339d.
//
// Solidity: function initialize(string service_, address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) PackInitialize(service string, validator common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("initialize", service, validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (baseForgeFacet *BaseForgeFacet) PackNonces(owner common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
}

// PackSetValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1327d3d8.
//
// Solidity: function setValidator(address validator_) returns()
func (baseForgeFacet *BaseForgeFacet) PackSetValidator(validator common.Address) []byte {
	enc, err := baseForgeFacet.abi.Pack("setValidator", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (baseForgeFacet *BaseForgeFacet) PackValidator() []byte {
	enc, err := baseForgeFacet.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
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
	return out0, err
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
	if log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
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
	if log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
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
	if log.Topics[0] != baseForgeFacet.abi.Events[event].ID {
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
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeECDSAInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], baseForgeFacet.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return baseForgeFacet.UnpackBaseForgeExpiredSignatureError(raw[4:])
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

// BaseForgeFacetBaseForgeECDSAInvalidValidatorSignature represents a BaseForge__ECDSAInvalidValidatorSignature error raised by the BaseForgeFacet contract.
type BaseForgeFacetBaseForgeECDSAInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func BaseForgeFacetBaseForgeECDSAInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0x1a126c3ca91a72268f221dd707e7d3ce39d63b8c31fd1956351a914fc1e8f138")
}

// UnpackBaseForgeECDSAInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func (baseForgeFacet *BaseForgeFacet) UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw []byte) (*BaseForgeFacetBaseForgeECDSAInvalidValidatorSignature, error) {
	out := new(BaseForgeFacetBaseForgeECDSAInvalidValidatorSignature)
	if err := baseForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeECDSAInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
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
