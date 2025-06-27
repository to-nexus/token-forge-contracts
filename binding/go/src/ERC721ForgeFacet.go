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

// ERC721ForgeFacetMetaData contains all meta data concerning the ERC721ForgeFacet contract.
var ERC721ForgeFacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC721From\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC721To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ERC721BurnForge__InvalidFromSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721MintForge__InvalidRecipientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721TransferForge__InvalidRecipientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"}]",
	ID:  "ERC721ForgeFacet",
	Bin: "0x6080604052348015600e575f5ffd5b5061220b8061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100c4575f3560e01c806384b0196e1161007d578063a8b8bdc511610058578063a8b8bdc51461023c578063bc2aed6c1461024f578063fd185ec514610262575f5ffd5b806384b0196e146101fb5780638c768fb8146102165780639daac2e614610229575f5ffd5b80633a5381b5116100ad5780633a5381b51461014b5780635d1a2de4146101925780637ecebe00146101a7575f5ffd5b8063150b7a02146100c85780633644e51514610135575b5f5ffd5b6100ff6100d6366004611b4f565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b61013d610275565b60405190815260200161012c565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012c565b6101a56101a0366004611bf8565b610283565b005b61013d6101b5366004611c64565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b61020361048c565b60405161012c9796959493929190611cd0565b6101a5610224366004611d8f565b61058b565b6101a5610237366004611bf8565b61088b565b6101a561024a366004611e3c565b610a88565b6101a561025d366004611d8f565b610d97565b6101a5610270366004611f15565b611097565b5f61027e611296565b905090565b8261028d8161129f565b335f7fe9dafd2dee8b38e04da5f914e54474a829a8dc3e85ae26496eb3099a3cce9f8a89838a8a6103068373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f610373826112df565b90506103b48187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018a90528a16906340c10f19906044015f604051808303815f87803b158015610421575f5ffd5b505af1158015610433573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff871660208201529081018b90526104809250600191508c908c906060015b604051602081830303815290604052611432565b50505050505050505050565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10080549091501580156104ca57506001810154155b610535576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b61053d6114e3565b6105456115b6565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b846105958161129f565b5f7ff6ee008a8b7b59fd4498eeaac762d59832615e3a5a1e650ca4aee37d1aa115df898961060b8d73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f61066c826112df565b90505f6106ae8289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061160792505050565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461072d576040517fcd3457ba00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260240161052c565b5050505f7f6743f8f4a310cd4e07c0b63b30c5ea8d7c609761294c9cc2c09f25153a561d2e8b8b888860405160200161076a959493929190611ff7565b6040516020818303038152906040528051906020012090505f61078c826112df565b90506107cd8186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b50506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990528916906340c10f19906044015f604051808303815f87803b15801561083c575f5ffd5b505af115801561084e573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8d1660208201529081018a90526104809250600191508c908b9060600161046c565b826108958161129f565b335f7f3aca9496191af1d8672735e28cb6e80e161763b9ce62e73efbd83dcb0b1fd49f89838a8a61090e8373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f61097b826112df565b90506109bc8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b6040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018a90528a16906379cc6790906044015f604051808303815f87803b158015610a29575f5ffd5b505af1158015610a3b573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff871660208201529081018b90526104809250600191508c908c906060015b60405160208183030381529060405261162f565b84610a928161129f565b5f7f82d117aa3c2f5377737ea027c8df4d0d2cc9d71bcb4f93eb60688b8810faaf588a8a610b088e73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f610b69826112df565b90505f610bab8289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061160792505050565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c2a576040517fdd93020600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e16600482015260240161052c565b5050505f7fd85eac2262fdc55c196c8602b9cd46035d340cd3e4ec58a849e5feea588219a58c8c8888604051602001610c67959493929190611ff7565b6040516020818303038152906040528051906020012090505f610c89826112df565b9050610cca8186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b50506040517fb88d4fde00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063b88d4fde90610d249030908e908d908d9060040161203d565b5f604051808303815f87803b158015610d3b575f5ffd5b505af1158015610d4d573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8e1660208201529081018b9052610d8a9250600191508d908c9060600161046c565b5050505050505050505050565b84610da18161129f565b5f7f02d2bd81b29d385e551058853a854cf81a16ed898911a385194b6458cd5061ce8989610e178d73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f610e78826112df565b90505f610eba8289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061160792505050565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f39576040517fde7ce5be00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260240161052c565b5050505f7fd2c6427412f1fa4eb9ed3183fa027c0f5226d5a85ea8a7af91dde94dc90f2dad8b8b8888604051602001610f76959493929190611ff7565b6040516020818303038152906040528051906020012090505f610f98826112df565b9050610fd98186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b50506040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990528916906379cc6790906044015f604051808303815f87803b158015611048575f5ffd5b505af115801561105a573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8d1660208201529081018a90526104809250600191508c908b90606001610a74565b826110a18161129f565b335f7f044eaaf6293a9f02fce75b10e66a99af83d0e2184c8ec0f67c393cc300f78ba58b838c8c61111a8373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f611187826112df565b90506111c88187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061132c92505050565b6040517fb88d4fde00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c169063b88d4fde9061122290309087908f908f908f90600401612097565b5f604051808303815f87803b158015611239575f5ffd5b505af115801561124b573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff871660208201529081018d90526112889250600191508e908e9060600161046c565b505050505050505050505050565b5f61027e6116ac565b804211156112dc576040517fa564eef90000000000000000000000000000000000000000000000000000000081526004810182905260240161052c565b50565b5f6113266112eb611296565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b92915050565b5f61136b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff81166113ba576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6113c58484611607565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461142c576040517f1a126c3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600546040517fe98a578400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690819063e98a5784906114af908890889088908890600401612115565b5f604051808303815f87803b1580156114c6575f5ffd5b505af11580156114d8573d5f5f3e3d5ffd5b505050505050505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161153490612184565b80601f016020809104026020016040519081016040528092919081815260200182805461156090612184565b80156115ab5780601f10611582576101008083540402835291602001916115ab565b820191905f5260205f20905b81548152906001019060200180831161158e57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161153490612184565b5f5f5f5f611615868661171f565b9250925092506116258282611768565b5090949350505050565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600546040517f3ea113d000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633ea113d0906114af908890889088908890600401612115565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6116d661186f565b6116de6118ea565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f5f5f8351604103611756576020840151604085015160608601515f1a6117488882858561193f565b955095509550505050611761565b505081515f91506002905b9250925092565b5f82600381111561177b5761177b6120e8565b03611784575050565b6001826003811115611798576117986120e8565b036117cf576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156117e3576117e36120e8565b0361181d576040517ffce698f70000000000000000000000000000000000000000000000000000000081526004810182905260240161052c565b6003826003811115611831576118316120e8565b0361186b576040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004810182905260240161052c565b5050565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008161189a6114e3565b8051909150156118b257805160209091012092915050565b815480156118c1579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816119156115b6565b80519091501561192d57805160209091012092915050565b600182015480156118c1579392505050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561197857505f91506003905082611a28565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156119c9573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611a1f57505f925060019150829050611a28565b92505f91508190505b9450945094915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a55575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611a96575f5ffd5b813567ffffffffffffffff811115611ab057611ab0611a5a565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611b1c57611b1c611a5a565b604052818152838201602001851015611b33575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f60808587031215611b62575f5ffd5b611b6b85611a32565b9350611b7960208601611a32565b925060408501359150606085013567ffffffffffffffff811115611b9b575f5ffd5b611ba787828801611a87565b91505092959194509250565b5f5f83601f840112611bc3575f5ffd5b50813567ffffffffffffffff811115611bda575f5ffd5b602083019150836020828501011115611bf1575f5ffd5b9250929050565b5f5f5f5f5f5f60a08789031215611c0d575f5ffd5b86359550611c1d60208801611a32565b94506040870135935060608701359250608087013567ffffffffffffffff811115611c46575f5ffd5b611c5289828a01611bb3565b979a9699509497509295939492505050565b5f60208284031215611c74575f5ffd5b611c7d82611a32565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f611d0a60e0830189611c84565b8281036040840152611d1c8189611c84565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611d7e578351835260209384019390920191600101611d60565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f5f5f60e08a8c031215611da7575f5ffd5b89359850611db760208b01611a32565b9750611dc560408b01611a32565b965060608a0135955060808a0135945060a08a013567ffffffffffffffff811115611dee575f5ffd5b611dfa8c828d01611bb3565b90955093505060c08a013567ffffffffffffffff811115611e19575f5ffd5b611e258c828d01611bb3565b915080935050809150509295985092959850929598565b5f5f5f5f5f5f5f5f5f5f6101008b8d031215611e56575f5ffd5b8a359950611e6660208c01611a32565b9850611e7460408c01611a32565b975060608b0135965060808b013567ffffffffffffffff811115611e96575f5ffd5b611ea28d828e01611a87565b96505060a08b0135945060c08b013567ffffffffffffffff811115611ec5575f5ffd5b611ed18d828e01611bb3565b90955093505060e08b013567ffffffffffffffff811115611ef0575f5ffd5b611efc8d828e01611bb3565b915080935050809150509295989b9194979a5092959850565b5f5f5f5f5f5f5f5f60c0898b031215611f2c575f5ffd5b88359750611f3c60208a01611a32565b965060408901359550606089013567ffffffffffffffff811115611f5e575f5ffd5b611f6a8b828c01611bb3565b9096509450506080890135925060a089013567ffffffffffffffff811115611f90575f5ffd5b611f9c8b828c01611bb3565b999c989b5096995094979396929594505050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b85815284602082015273ffffffffffffffffffffffffffffffffffffffff84166040820152608060608201525f612032608083018486611fb0565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f61208d6080830184611c84565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152608060608201525f612032608083018486611fb0565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f6003861061214b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b85825284602083015273ffffffffffffffffffffffffffffffffffffffff841660408301526080606083015261208d6080830184611c84565b600181811c9082168061219857607f821691505b6020821081036121cf577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fea2646970667358221220f772b917374eb9c210e4d2727a0b1cbd3549d32a6c940d7731c407c0f4b6268364736f6c634300081c0033",
}

// ERC721ForgeFacet is an auto generated Go binding around an Ethereum contract.
type ERC721ForgeFacet struct {
	abi abi.ABI
}

// NewERC721ForgeFacet creates a new instance of ERC721ForgeFacet.
func NewERC721ForgeFacet() *ERC721ForgeFacet {
	parsed, err := ERC721ForgeFacetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721ForgeFacet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721ForgeFacet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC721ForgeFacet *ERC721ForgeFacet) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC721ForgeFacet.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackBurnERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9daac2e6.
//
// Solidity: function burnERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackBurnERC721(uuid *big.Int, token common.Address, tokenID *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("burnERC721", uuid, token, tokenID, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC721From is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc2aed6c.
//
// Solidity: function burnERC721From(uint256 uuid, address from, address token, uint256 tokenID, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackBurnERC721From(uuid *big.Int, from common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("burnERC721From", uuid, from, token, tokenID, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC721ForgeFacet *ERC721ForgeFacet) PackEip712Domain() []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC721ForgeFacet.abi.Unpack("eip712Domain", data)
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

// PackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d1a2de4.
//
// Solidity: function mintERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackMintERC721(uuid *big.Int, token common.Address, tokenID *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("mintERC721", uuid, token, tokenID, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC721To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c768fb8.
//
// Solidity: function mintERC721To(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackMintERC721To(uuid *big.Int, recipient common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("mintERC721To", uuid, recipient, token, tokenID, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC721ForgeFacet *ERC721ForgeFacet) PackNonces(owner common.Address) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC721ForgeFacet.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (eRC721ForgeFacet *ERC721ForgeFacet) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := eRC721ForgeFacet.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd185ec5.
//
// Solidity: function transferERC721(uint256 uuid, address token, uint256 tokenID, bytes data, uint256 deadline, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackTransferERC721(uuid *big.Int, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("transferERC721", uuid, token, tokenID, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC721To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8b8bdc5.
//
// Solidity: function transferERC721To(uint256 uuid, address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (eRC721ForgeFacet *ERC721ForgeFacet) PackTransferERC721To(uuid *big.Int, recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("transferERC721To", uuid, recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (eRC721ForgeFacet *ERC721ForgeFacet) PackValidator() []byte {
	enc, err := eRC721ForgeFacet.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackValidator(data []byte) (common.Address, error) {
	out, err := eRC721ForgeFacet.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// ERC721ForgeFacetEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC721ForgeFacetEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC721ForgeFacetEIP712DomainChanged) ContractEventName() string {
	return ERC721ForgeFacetEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC721ForgeFacetEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC721ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721ForgeFacetEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721ForgeFacet.abi.Events[event].Inputs {
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

// ERC721ForgeFacetInitialized represents a Initialized event raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721ForgeFacetInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC721ForgeFacetInitialized) ContractEventName() string {
	return ERC721ForgeFacetInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackInitializedEvent(log *types.Log) (*ERC721ForgeFacetInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC721ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721ForgeFacetInitialized)
	if len(log.Data) > 0 {
		if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721ForgeFacet.abi.Events[event].Inputs {
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

// ERC721ForgeFacetValidatorUpdated represents a ValidatorUpdated event raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC721ForgeFacetValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (ERC721ForgeFacetValidatorUpdated) ContractEventName() string {
	return ERC721ForgeFacetValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackValidatorUpdatedEvent(log *types.Log) (*ERC721ForgeFacetValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if log.Topics[0] != eRC721ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721ForgeFacetValidatorUpdated)
	if len(log.Data) > 0 {
		if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721ForgeFacet.abi.Events[event].Inputs {
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
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["BaseForgeECDSAInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ERC721BurnForgeInvalidFromSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackERC721BurnForgeInvalidFromSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ERC721MintForgeInvalidRecipientSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackERC721MintForgeInvalidRecipientSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["ERC721TransferForgeInvalidRecipientSignature"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackERC721TransferForgeInvalidRecipientSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721ForgeFacet.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC721ForgeFacet.UnpackNotInitializingError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721ForgeFacetBaseForgeECDSAInvalidValidatorSignature represents a BaseForge__ECDSAInvalidValidatorSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetBaseForgeECDSAInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func ERC721ForgeFacetBaseForgeECDSAInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0x1a126c3ca91a72268f221dd707e7d3ce39d63b8c31fd1956351a914fc1e8f138")
}

// UnpackBaseForgeECDSAInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw []byte) (*ERC721ForgeFacetBaseForgeECDSAInvalidValidatorSignature, error) {
	out := new(ERC721ForgeFacetBaseForgeECDSAInvalidValidatorSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeECDSAInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetBaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetBaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func ERC721ForgeFacetBaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackBaseForgeExpiredSignatureError(raw []byte) (*ERC721ForgeFacetBaseForgeExpiredSignature, error) {
	out := new(ERC721ForgeFacetBaseForgeExpiredSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetBaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetBaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func ERC721ForgeFacetBaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackBaseForgeZeroAddressError(raw []byte) (*ERC721ForgeFacetBaseForgeZeroAddress, error) {
	out := new(ERC721ForgeFacetBaseForgeZeroAddress)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC721ForgeFacetECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC721ForgeFacetECDSAInvalidSignature, error) {
	out := new(ERC721ForgeFacetECDSAInvalidSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC721ForgeFacetECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC721ForgeFacetECDSAInvalidSignatureLength, error) {
	out := new(ERC721ForgeFacetECDSAInvalidSignatureLength)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC721ForgeFacetECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC721ForgeFacetECDSAInvalidSignatureS, error) {
	out := new(ERC721ForgeFacetECDSAInvalidSignatureS)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetERC721BurnForgeInvalidFromSignature represents a ERC721BurnForge__InvalidFromSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetERC721BurnForgeInvalidFromSignature struct {
	From common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721BurnForge__InvalidFromSignature(address from)
func ERC721ForgeFacetERC721BurnForgeInvalidFromSignatureErrorID() common.Hash {
	return common.HexToHash("0xde7ce5be00ae6217963deecce7891b5e9191ebbfe992ce4b03962044ca114bce")
}

// UnpackERC721BurnForgeInvalidFromSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721BurnForge__InvalidFromSignature(address from)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackERC721BurnForgeInvalidFromSignatureError(raw []byte) (*ERC721ForgeFacetERC721BurnForgeInvalidFromSignature, error) {
	out := new(ERC721ForgeFacetERC721BurnForgeInvalidFromSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ERC721BurnForgeInvalidFromSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetERC721MintForgeInvalidRecipientSignature represents a ERC721MintForge__InvalidRecipientSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetERC721MintForgeInvalidRecipientSignature struct {
	Recipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721MintForge__InvalidRecipientSignature(address recipient)
func ERC721ForgeFacetERC721MintForgeInvalidRecipientSignatureErrorID() common.Hash {
	return common.HexToHash("0xcd3457babac82634ab3742a6056df59537d2eba6a801e097bd1a5c85f54409e6")
}

// UnpackERC721MintForgeInvalidRecipientSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721MintForge__InvalidRecipientSignature(address recipient)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackERC721MintForgeInvalidRecipientSignatureError(raw []byte) (*ERC721ForgeFacetERC721MintForgeInvalidRecipientSignature, error) {
	out := new(ERC721ForgeFacetERC721MintForgeInvalidRecipientSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ERC721MintForgeInvalidRecipientSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetERC721TransferForgeInvalidRecipientSignature represents a ERC721TransferForge__InvalidRecipientSignature error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetERC721TransferForgeInvalidRecipientSignature struct {
	Recipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721TransferForge__InvalidRecipientSignature(address recipient)
func ERC721ForgeFacetERC721TransferForgeInvalidRecipientSignatureErrorID() common.Hash {
	return common.HexToHash("0xdd9302066a81bf9b311e581276749fdbe763a00f7d36ad9dcb5f6f51ab1fe041")
}

// UnpackERC721TransferForgeInvalidRecipientSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721TransferForge__InvalidRecipientSignature(address recipient)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackERC721TransferForgeInvalidRecipientSignatureError(raw []byte) (*ERC721ForgeFacetERC721TransferForgeInvalidRecipientSignature, error) {
	out := new(ERC721ForgeFacetERC721TransferForgeInvalidRecipientSignature)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "ERC721TransferForgeInvalidRecipientSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC721ForgeFacetInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackInvalidAccountNonceError(raw []byte) (*ERC721ForgeFacetInvalidAccountNonce, error) {
	out := new(ERC721ForgeFacetInvalidAccountNonce)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetInvalidInitialization represents a InvalidInitialization error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC721ForgeFacetInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackInvalidInitializationError(raw []byte) (*ERC721ForgeFacetInvalidInitialization, error) {
	out := new(ERC721ForgeFacetInvalidInitialization)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ForgeFacetNotInitializing represents a NotInitializing error raised by the ERC721ForgeFacet contract.
type ERC721ForgeFacetNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC721ForgeFacetNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC721ForgeFacet *ERC721ForgeFacet) UnpackNotInitializingError(raw []byte) (*ERC721ForgeFacetNotInitializing, error) {
	out := new(ERC721ForgeFacetNotInitializing)
	if err := eRC721ForgeFacet.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}
