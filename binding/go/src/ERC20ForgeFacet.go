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

// ERC20ForgeFacetMetaData contains all meta data concerning the ERC20ForgeFacet contract.
var ERC20ForgeFacetMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20From\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseForge__ECDSAInvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ERC20BurnForge__InvalidFromSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC20MintForge__InvalidRecipientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC20TransferForge__InvalidRecipientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ERC20ForgeFacet",
	Bin: "0x6080604052348015600e575f5ffd5b50611ee88061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100b9575f3560e01c80633e2f35091161007257806384b0196e1161005857806384b0196e146101c1578063a199c756146101dc578063d93e4044146101ef575f5ffd5b80633e2f35091461015a5780637ecebe001461016d575f5ffd5b80632e6526c8116100a25780632e6526c8146100e55780633644e515146100f85780633a5381b514610113575f5ffd5b80630588ec1d146100bd57806307c082f2146100d2575b5f5ffd5b6100d06100cb366004611b00565b610202565b005b6100d06100e0366004611bad565b6104c1565b6100d06100f3366004611b00565b61064b565b61010061095a565b6040519081526020015b60405180910390f35b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010a565b6100d0610168366004611bad565b610968565b61010061017b366004611c19565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090205490565b6101c9610b54565b60405161010a9796959493929190611c85565b6100d06101ea366004611bad565b610c4e565b6100d06101fd366004611b00565b610e4a565b8461020c81611149565b5f7f8b4141bfe385e9e19a60e5e8216bfb64e59e7e73d0b69318be9a1316defebec589896102828d73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f6102e382611189565b90505f6103258289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111d692505050565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146103a9576040517f3c725bab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024015b60405180910390fd5b5050505f7f7dbc318bd2848e28f56eca353ec5c77fc714e5969ab2b10d4e9d399f7679ceb18b8b88886040516020016103e6959493929190611d44565b6040516020818303038152906040528051906020012090505f61040882611189565b90506104498186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b5061046d905073ffffffffffffffffffffffffffffffffffffffff89168a89611304565b6040805173ffffffffffffffffffffffffffffffffffffffff8b1660208201529081018890526104b5905f908c908b906060015b604051602081830303815290604052611396565b50505050505050505050565b826104cb81611149565b335f7f972335ea83f5a60c2ef1f633495b4137e25728a171b2c24e15250fa65f8ca00e89838a8a6105448373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f6105b182611189565b90506105f28187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b61061373ffffffffffffffffffffffffffffffffffffffff8a16848a611304565b6040805173ffffffffffffffffffffffffffffffffffffffff851660208201529081018990526104b5905f908c908c906060016104a1565b8461065581611149565b5f7fc9d15b3c70a92ba5918c1c7dea408aa3264f3ea1b00c34a3f111934100bc86ab89896106cb8d73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f61072c82611189565b90505f61076e8289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111d692505050565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146107ed576040517f6fb1dc1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016103a0565b5050505f7f3279e9bd42ac832f9f552e3bcd46286a3a92d40be49821816a5804e107c0f9ff8b8b888860405160200161082a959493929190611d44565b6040516020818303038152906040528051906020012090505f61084c82611189565b905061088d8186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b50506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990528916906340c10f19906044015f604051808303815f87803b1580156108fc575f5ffd5b505af115801561090e573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8d1660208201529081018a90526104b592505f91508c908b906060015b604051602081830303815290604052611447565b5f6109636114c4565b905090565b8261097281611149565b335f7f0bbe2a47c9896a0d4ac145f429a0a95c7bf7909d4f4f14e9d9401eece4b9e87b89838a8a6109eb8373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f610a5882611189565b9050610a998187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018a90528a16906340c10f19906044015f604051808303815f87803b158015610b06575f5ffd5b505af1158015610b18573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff871660208201529081018b90526104b592505f91508c908c90606001610946565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610b9257506001810154155b610bf8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016103a0565b610c006114cd565b610c086115a0565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b82610c5881611149565b335f7fa69fbd2604b7b87c6eb41de807f7f56c03f95f311b26a469bbb87ee6335c666089838a8a610cd18373ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019790975286019490945273ffffffffffffffffffffffffffffffffffffffff92831660608601529116608084015260a083015260c082015260e08101879052610100016040516020818303038152906040528051906020012090505f610d3e82611189565b9050610d7f8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b6040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018a90528a16906379cc6790906044015f604051808303815f87803b158015610dec575f5ffd5b505af1158015610dfe573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff871660208201529081018b90526104b592505f91508c908c906060015b6040516020818303038152906040526115f1565b84610e5481611149565b5f7fa32315684bbb8a0a13dc1447ad39319ff7a5c1c390dd45e09a6bcb7190acd5b88989610eca8d73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019590955273ffffffffffffffffffffffffffffffffffffffff909316928401929092526060830152608082015260a0810188905260c0016040516020818303038152906040528051906020012090505f610f2b82611189565b90505f610f6d8289898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111d692505050565b90508b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610fec576040517ffe0003be00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d1660048201526024016103a0565b5050505f7fe1bc7f1d02e94d490b537fa8074a1b08e8208824df594b9ba30e905e6e2e7ce78b8b8888604051602001611029959493929190611d44565b6040516020818303038152906040528051906020012090505f61104b82611189565b905061108c8186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506111fe92505050565b50506040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018990528916906379cc6790906044015f604051808303815f87803b1580156110fb575f5ffd5b505af115801561110d573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8d1660208201529081018a90526104b592505f91508c908b90606001610e36565b80421115611186576040517fa564eef9000000000000000000000000000000000000000000000000000000008152600481018290526024016103a0565b50565b5f6111d06111956114c4565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b92915050565b5f5f5f5f6111e4868661166e565b9250925092506111f482826116b7565b5090949350505050565b5f61123d7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff811661128c576040517f95d18bbf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61129784846111d6565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146112fe576040517f1a126c3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526113919084906117be565b505050565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600546040517f0b7e40c600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190630b7e40c690611413908890889088908890600401611de8565b5f604051808303815f87803b15801561142a575f5ffd5b505af115801561143c573d5f5f3e3d5ffd5b505050505050505050565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600546040517fe98a578400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690819063e98a578490611413908890889088908890600401611de8565b5f61096361185d565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161151e90611e61565b80601f016020809104026020016040519081016040528092919081815260200182805461154a90611e61565b80156115955780601f1061156c57610100808354040283529160200191611595565b820191905f5260205f20905b81548152906001019060200180831161157857829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161151e90611e61565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600546040517f3ea113d000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116908190633ea113d090611413908890889088908890600401611de8565b5f5f5f83516041036116a5576020840151604085015160608601515f1a611697888285856118d0565b9550955095505050506116b0565b505081515f91506002905b9250925092565b5f8260038111156116ca576116ca611dbb565b036116d3575050565b60018260038111156116e7576116e7611dbb565b0361171e576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561173257611732611dbb565b0361176c576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016103a0565b600382600381111561178057611780611dbb565b036117ba576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016103a0565b5050565b5f5f60205f8451602086015f885af1806117dd576040513d5f823e3d81fd5b50505f513d915081156117f457806001141561180e565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156112fe576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016103a0565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6118876119c3565b61188f611a3e565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561190957505f915060039050826119b9565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561195a573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166119b057505f9250600191508290506119b9565b92505f91508190505b9450945094915050565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816119ee6114cd565b805190915015611a0657805160209091012092915050565b81548015611a15579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611a696115a0565b805190915015611a8157805160209091012092915050565b60018201548015611a15579392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611ab6575f5ffd5b919050565b5f5f83601f840112611acb575f5ffd5b50813567ffffffffffffffff811115611ae2575f5ffd5b602083019150836020828501011115611af9575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f60e08a8c031215611b18575f5ffd5b89359850611b2860208b01611a93565b9750611b3660408b01611a93565b965060608a0135955060808a0135945060a08a013567ffffffffffffffff811115611b5f575f5ffd5b611b6b8c828d01611abb565b90955093505060c08a013567ffffffffffffffff811115611b8a575f5ffd5b611b968c828d01611abb565b915080935050809150509295985092959850929598565b5f5f5f5f5f5f60a08789031215611bc2575f5ffd5b86359550611bd260208801611a93565b94506040870135935060608701359250608087013567ffffffffffffffff811115611bfb575f5ffd5b611c0789828a01611abb565b979a9699509497509295939492505050565b5f60208284031215611c29575f5ffd5b611c3282611a93565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f611cbf60e0830189611c39565b8281036040840152611cd18189611c39565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611d33578351835260209384019390920191600101611d15565b50909b9a5050505050505050505050565b85815284602082015273ffffffffffffffffffffffffffffffffffffffff8416604082015260806060820152816080820152818360a08301375f81830160a090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f60038610611e1e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b85825284602083015273ffffffffffffffffffffffffffffffffffffffff8416604083015260806060830152611e576080830184611c39565b9695505050505050565b600181811c90821680611e7557607f821691505b602082108103611eac577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fea26469706673582212209a3e4293f06f1f4ddc8c342dab2948e42788c371393b1b168dddd447db0c407f64736f6c634300081c0033",
}

// ERC20ForgeFacet is an auto generated Go binding around an Ethereum contract.
type ERC20ForgeFacet struct {
	abi abi.ABI
}

// NewERC20ForgeFacet creates a new instance of ERC20ForgeFacet.
func NewERC20ForgeFacet() *ERC20ForgeFacet {
	parsed, err := ERC20ForgeFacetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20ForgeFacet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20ForgeFacet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20ForgeFacet *ERC20ForgeFacet) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20ForgeFacet.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackBurnERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa199c756.
//
// Solidity: function burnERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackBurnERC20(uuid *big.Int, token common.Address, amount *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("burnERC20", uuid, token, amount, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnERC20From is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd93e4044.
//
// Solidity: function burnERC20From(uint256 uuid, address from, address token, uint256 amount, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackBurnERC20From(uuid *big.Int, from common.Address, token common.Address, amount *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("burnERC20From", uuid, from, token, amount, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20ForgeFacet *ERC20ForgeFacet) PackEip712Domain() []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20ForgeFacet.abi.Unpack("eip712Domain", data)
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

// PackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e2f3509.
//
// Solidity: function mintERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackMintERC20(uuid *big.Int, token common.Address, amount *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("mintERC20", uuid, token, amount, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e6526c8.
//
// Solidity: function mintERC20To(uint256 uuid, address recipient, address token, uint256 amount, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackMintERC20To(uuid *big.Int, recipient common.Address, token common.Address, amount *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("mintERC20To", uuid, recipient, token, amount, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20ForgeFacet *ERC20ForgeFacet) PackNonces(owner common.Address) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20ForgeFacet.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07c082f2.
//
// Solidity: function transferERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackTransferERC20(uuid *big.Int, token common.Address, amount *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("transferERC20", uuid, token, amount, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0588ec1d.
//
// Solidity: function transferERC20To(uint256 uuid, address recipient, address token, uint256 amount, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (eRC20ForgeFacet *ERC20ForgeFacet) PackTransferERC20To(uuid *big.Int, recipient common.Address, token common.Address, amount *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("transferERC20To", uuid, recipient, token, amount, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (eRC20ForgeFacet *ERC20ForgeFacet) PackValidator() []byte {
	enc, err := eRC20ForgeFacet.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackValidator(data []byte) (common.Address, error) {
	out, err := eRC20ForgeFacet.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// ERC20ForgeFacetEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20ForgeFacetEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20ForgeFacetEIP712DomainChanged) ContractEventName() string {
	return ERC20ForgeFacetEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20ForgeFacetEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20ForgeFacetEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20ForgeFacet.abi.Events[event].Inputs {
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

// ERC20ForgeFacetInitialized represents a Initialized event raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20ForgeFacetInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20ForgeFacetInitialized) ContractEventName() string {
	return ERC20ForgeFacetInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackInitializedEvent(log *types.Log) (*ERC20ForgeFacetInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20ForgeFacetInitialized)
	if len(log.Data) > 0 {
		if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20ForgeFacet.abi.Events[event].Inputs {
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

// ERC20ForgeFacetValidatorUpdated represents a ValidatorUpdated event raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20ForgeFacetValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20ForgeFacetValidatorUpdated) ContractEventName() string {
	return ERC20ForgeFacetValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackValidatorUpdatedEvent(log *types.Log) (*ERC20ForgeFacetValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if log.Topics[0] != eRC20ForgeFacet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20ForgeFacetValidatorUpdated)
	if len(log.Data) > 0 {
		if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20ForgeFacet.abi.Events[event].Inputs {
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
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["BaseForgeECDSAInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ERC20BurnForgeInvalidFromSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackERC20BurnForgeInvalidFromSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ERC20MintForgeInvalidRecipientSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackERC20MintForgeInvalidRecipientSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["ERC20TransferForgeInvalidRecipientSignature"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackERC20TransferForgeInvalidRecipientSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20ForgeFacet.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return eRC20ForgeFacet.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20ForgeFacetBaseForgeECDSAInvalidValidatorSignature represents a BaseForge__ECDSAInvalidValidatorSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetBaseForgeECDSAInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func ERC20ForgeFacetBaseForgeECDSAInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0x1a126c3ca91a72268f221dd707e7d3ce39d63b8c31fd1956351a914fc1e8f138")
}

// UnpackBaseForgeECDSAInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ECDSAInvalidValidatorSignature()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackBaseForgeECDSAInvalidValidatorSignatureError(raw []byte) (*ERC20ForgeFacetBaseForgeECDSAInvalidValidatorSignature, error) {
	out := new(ERC20ForgeFacetBaseForgeECDSAInvalidValidatorSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeECDSAInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetBaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetBaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func ERC20ForgeFacetBaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackBaseForgeExpiredSignatureError(raw []byte) (*ERC20ForgeFacetBaseForgeExpiredSignature, error) {
	out := new(ERC20ForgeFacetBaseForgeExpiredSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetBaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetBaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func ERC20ForgeFacetBaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackBaseForgeZeroAddressError(raw []byte) (*ERC20ForgeFacetBaseForgeZeroAddress, error) {
	out := new(ERC20ForgeFacetBaseForgeZeroAddress)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20ForgeFacetECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20ForgeFacetECDSAInvalidSignature, error) {
	out := new(ERC20ForgeFacetECDSAInvalidSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20ForgeFacetECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20ForgeFacetECDSAInvalidSignatureLength, error) {
	out := new(ERC20ForgeFacetECDSAInvalidSignatureLength)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20ForgeFacetECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20ForgeFacetECDSAInvalidSignatureS, error) {
	out := new(ERC20ForgeFacetECDSAInvalidSignatureS)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetERC20BurnForgeInvalidFromSignature represents a ERC20BurnForge__InvalidFromSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetERC20BurnForgeInvalidFromSignature struct {
	From common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20BurnForge__InvalidFromSignature(address from)
func ERC20ForgeFacetERC20BurnForgeInvalidFromSignatureErrorID() common.Hash {
	return common.HexToHash("0xfe0003beb67304596404eed58fb6bdf19a0528d0e69efe77151172d81dac0715")
}

// UnpackERC20BurnForgeInvalidFromSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20BurnForge__InvalidFromSignature(address from)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackERC20BurnForgeInvalidFromSignatureError(raw []byte) (*ERC20ForgeFacetERC20BurnForgeInvalidFromSignature, error) {
	out := new(ERC20ForgeFacetERC20BurnForgeInvalidFromSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ERC20BurnForgeInvalidFromSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetERC20MintForgeInvalidRecipientSignature represents a ERC20MintForge__InvalidRecipientSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetERC20MintForgeInvalidRecipientSignature struct {
	Recipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20MintForge__InvalidRecipientSignature(address recipient)
func ERC20ForgeFacetERC20MintForgeInvalidRecipientSignatureErrorID() common.Hash {
	return common.HexToHash("0x6fb1dc192672bd664b7a247e31bab49b8bb7de43ff70b48bad270c32491d1974")
}

// UnpackERC20MintForgeInvalidRecipientSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20MintForge__InvalidRecipientSignature(address recipient)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackERC20MintForgeInvalidRecipientSignatureError(raw []byte) (*ERC20ForgeFacetERC20MintForgeInvalidRecipientSignature, error) {
	out := new(ERC20ForgeFacetERC20MintForgeInvalidRecipientSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ERC20MintForgeInvalidRecipientSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetERC20TransferForgeInvalidRecipientSignature represents a ERC20TransferForge__InvalidRecipientSignature error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetERC20TransferForgeInvalidRecipientSignature struct {
	Recipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20TransferForge__InvalidRecipientSignature(address recipient)
func ERC20ForgeFacetERC20TransferForgeInvalidRecipientSignatureErrorID() common.Hash {
	return common.HexToHash("0x3c725bab87a3a54c1298e1c05d52c85fcfa98ce673fdd74b242f59d0080a0c52")
}

// UnpackERC20TransferForgeInvalidRecipientSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20TransferForge__InvalidRecipientSignature(address recipient)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackERC20TransferForgeInvalidRecipientSignatureError(raw []byte) (*ERC20ForgeFacetERC20TransferForgeInvalidRecipientSignature, error) {
	out := new(ERC20ForgeFacetERC20TransferForgeInvalidRecipientSignature)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "ERC20TransferForgeInvalidRecipientSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20ForgeFacetInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackInvalidAccountNonceError(raw []byte) (*ERC20ForgeFacetInvalidAccountNonce, error) {
	out := new(ERC20ForgeFacetInvalidAccountNonce)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetInvalidInitialization represents a InvalidInitialization error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20ForgeFacetInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackInvalidInitializationError(raw []byte) (*ERC20ForgeFacetInvalidInitialization, error) {
	out := new(ERC20ForgeFacetInvalidInitialization)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetNotInitializing represents a NotInitializing error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20ForgeFacetNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackNotInitializingError(raw []byte) (*ERC20ForgeFacetNotInitializing, error) {
	out := new(ERC20ForgeFacetNotInitializing)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20ForgeFacetSafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ERC20ForgeFacet contract.
type ERC20ForgeFacetSafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ERC20ForgeFacetSafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (eRC20ForgeFacet *ERC20ForgeFacet) UnpackSafeERC20FailedOperationError(raw []byte) (*ERC20ForgeFacetSafeERC20FailedOperation, error) {
	out := new(ERC20ForgeFacetSafeERC20FailedOperation)
	if err := eRC20ForgeFacet.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
