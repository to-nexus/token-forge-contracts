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

// TokenForgeFactoryMetaData contains all meta data concerning the TokenForgeFactory contract.
var TokenForgeFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"addCuts\",\"type\":\"tuple[]\"}],\"name\":\"addService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allForges\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"}],\"name\":\"forgeByService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forgeProxyCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_diamondImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isRunningForge\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lengthAllForges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"pauseService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"ServicePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"ServiceUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__AlreadyUsedService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsNotForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsPausedForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__InvalidData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__ServiceNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenForge__InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "TokenForgeFactory",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051613e466100395f395f81816120e70152818161211001526123300152613e465ff3fe60806040526004361061016d575f3560e01c80639ca92df9116100c6578063d547741f1161007c578063f4d8cab911610057578063f4d8cab9146104db578063f51acaea146104fa578063f8c8765e14610519575f5ffd5b8063d547741f1461047b578063e57997d21461049a578063e98a5784146104bc575f5ffd5b8063ad3cb1cc116100ac578063ad3cb1cc146103ae578063c7525cba14610403578063d3de1d7f14610430575f5ffd5b80639ca92df91461036d578063a217fddf1461039b575f5ffd5b80632f44ec09116101265780634f1ef286116101015780634f1ef286146102d657806352d1902d146102e957806391d14854146102fd575f5ffd5b80632f44ec091461028457806336568abe146102985780633ea113d0146102b7575f5ffd5b8063217edd3411610156578063217edd34146101c6578063248a9ca31461020a5780632f2ff15d14610265575f5ffd5b806301ffc9a7146101715780630b7e40c6146101a5575b5f5ffd5b34801561017c575f5ffd5b5061019061018b3660046130f5565b610538565b60405190151581526020015b60405180910390f35b3480156101b0575f5ffd5b506101c46101bf36600461317b565b6105d0565b005b3480156101d1575f5ffd5b506101e56101e03660046131e9565b610910565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b348015610215575f5ffd5b506102576102243660046132ae565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161019c565b348015610270575f5ffd5b506101c461027f3660046132c5565b610dbe565b34801561028f575f5ffd5b50610257610e07565b3480156102a3575f5ffd5b506101c46102b23660046132c5565b610e36565b3480156102c2575f5ffd5b506101c46102d136600461317b565b610e94565b6101c46102e4366004613431565b611170565b3480156102f4575f5ffd5b5061025761118f565b348015610308575f5ffd5b506101906103173660046132c5565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b348015610378575f5ffd5b5061038c6103873660046132ae565b6111bd565b60405161019c939291906134ca565b3480156103a6575f5ffd5b506102575f81565b3480156103b9575f5ffd5b506103f66040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161019c9190613509565b34801561040e575f5ffd5b5061042261041d36600461351b565b611238565b60405161019c929190613536565b34801561043b575f5ffd5b5061044f61044a366004613559565b6112d1565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835290151560208301520161019c565b348015610486575f5ffd5b506101c46104953660046132c5565b611353565b3480156104a5575f5ffd5b506104ae611396565b60405161019c92919061358b565b3480156104c7575f5ffd5b506101c46104d636600461317b565b61151d565b3480156104e6575f5ffd5b506101c46104f5366004613674565b6117f9565b348015610505575f5ffd5b506101c4610514366004613559565b61198e565b348015610524575f5ffd5b506101c46105333660046136bf565b611aea565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806105ca57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f6105fc7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06005b33611c69565b90505f86600281111561061157610611613718565b036106a2575f8061062484860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f86d733c2654f9db1a830cfbbf2180ee3329101d853428a7ead4e85ddb3230d60898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a45050610908565b60018660028111156106b6576106b6613718565b03610738575f806106c984860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847fa36201ae386aaacb89bc573b7c0ecc0913a93028f710f28ca29840454f446886898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600286600281111561074c5761074c613718565b036108d6575f8061075f8486018661376f565b91509150811561084c575f5f5f838060200190518101906107809190613811565b815192955090935091505f5b81811015610842578473ffffffffffffffffffffffffffffffffffffffff168c897fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c38e8886815181106107e1576107e1613888565b60200260200101518887815181106107fb576107fb613888565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a460010161078c565b50505050506108cf565b5f5f5f8380602001905181019061086391906138b5565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c3906060015b60405180910390a45050505b5050610908565b6040517f9570ddf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61093b81611d47565b73ffffffffffffffffffffffffffffffffffffffff88166109af576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8716610a1e576040517f435db3480000000000000000000000000000000000000000000000000000000081527f76616c696461746f72000000000000000000000000000000000000000000000060048201526024016109a6565b5f859003610a7a576040517f435db3480000000000000000000000000000000000000000000000000000000081527f736572766963650000000000000000000000000000000000000000000000000060048201526024016109a6565b5f610ab987878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611d5492505050565b90507f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600610b067f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060383611daa565b15610b40576040517f5502583a000000000000000000000000000000000000000000000000000000008152600481018390526024016109a6565b610ca95f83835f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015610baf573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610bf491908101906138e9565b8d8d8b8b8f8f8a6001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b6002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001610c59989796959493929190613a06565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610c959291602001613c3a565b604051602081830303815290604052611db5565b935073ffffffffffffffffffffffffffffffffffffffff8416610d1a576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f726765206465706c6f79000000000000000000000000000000000000000060048201526024016109a6565b610d28600382018386611ea1565b5073ffffffffffffffffffffffffffffffffffffffff84165f81815260068301602090815260408083208690558583526007850190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f0453e3ffce9b0d4f994d54652037d3e15d87ec29aab0aafba90ac322edc87c2591a35050509695505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610df781611d47565b610e018383611ecb565b50505050565b5f610e317f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0603611fe9565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610e85576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e8f8282611ff3565b505050565b5f610ebe7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06006105f6565b90505f866002811115610ed357610ed3613718565b03610f55575f80610ee684860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847fe5e0650048e213b8a668474e72c774ec69d8a6830270b7cc1c339fcd6a3b2e41898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6001866002811115610f6957610f69613718565b03610feb575f80610f7c84860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f28fcdc3e3355f548f5fe83373acf91a82e1e561ee22797f6e91f831d2b2b3e13898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6002866002811115610fff57610fff613718565b036108d6575f806110128486018661376f565b9150915081156110f5575f5f5f838060200190518101906110339190613811565b815192955090935091505f5b81811015610842578473ffffffffffffffffffffffffffffffffffffffff168c897f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b8e88868151811061109457611094613888565b60200260200101518887815181106110ae576110ae613888565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a460010161103f565b5f5f5f8380602001905181019061110c91906138b5565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b906060016108c3565b6111786120cf565b611181826121d5565b61118b82826121df565b5050565b5f611198612318565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60605f807f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06008161120d7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060387612387565b9450905061121a816123a4565b5f91825260079092016020526040902054909592945060ff16925050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06066020526040812054606091907f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06009080156112ca576112b1816123a4565b5f82815260078401602052604090205490945060ff1692505b5050915091565b5f5f5f6112dd84611d54565b90507f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06005f61132b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0603846123e1565b95509050801561134b575f83815260078301602052604090205460ff1693505b505050915091565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461138c81611d47565b610e018383611ff3565b6060807f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06005f6113e47f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0603611fe9565b90505f8167ffffffffffffffff811115611400576114006132f3565b60405190808252806020026020018201604052801561143357816020015b606081526020019060019003908161141e5790505b5090505f8267ffffffffffffffff811115611450576114506132f3565b604051908082528060200260200182016040528015611479578160200160208202803683370190505b5090505f5b83811015611511575f6114946003870183612387565b8484815181106114a6576114a6613888565b602002602001018173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525081925050506114eb816123a4565b8483815181106114fd576114fd613888565b60209081029190910101525060010161147e565b50909590945092505050565b5f6115477f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06006105f6565b90505f86600281111561155c5761155c613718565b036115de575f8061156f84860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f9612604afba70e4cf03261d7d86ca03d08911d9887aa41621dea929e34cfd7b1898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60018660028111156115f2576115f2613718565b03611674575f8061160584860186613745565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f556d832108ef7e0123726b3e037bc1cd39023e8d397b242d8b36426550d8c123898560405161069392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600286600281111561168857611688613718565b036108d6575f8061169b8486018661376f565b91509150811561177e575f5f5f838060200190518101906116bc9190613811565b815192955090935091505f5b81811015610842578473ffffffffffffffffffffffffffffffffffffffff168c897f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f618e88868151811061171d5761171d613888565b602002602001015188878151811061173757611737613888565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a46001016116c8565b5f5f5f8380602001905181019061179591906138b5565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f61906060016108c3565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61182381611d47565b5f61182d84611d54565b90507f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060061187a7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060383611daa565b6118b3576040517f2a188b7f000000000000000000000000000000000000000000000000000000008152600481018390526024016109a6565b5f82815260078201602052604090205484151560ff909116151503611987578315611912575f828152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905561194b565b5f828152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b817eb31832845d6567ecb46caeeee58acfcfddfcbba821c68dcbae9055ea4d79558560405161197e911515815260200190565b60405180910390a25b5050505050565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6119b881611d47565b5f6119c283611d54565b90507f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06005f80611a117f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0603856123e1565b9150915081611a4f576040517f2a188b7f000000000000000000000000000000000000000000000000000000008152600481018590526024016109a6565b611a5c60038401856123ef565b5073ffffffffffffffffffffffffffffffffffffffff81165f90815260068401602090815260408083208390558683526007860190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555185917f475a09e552c07ad6aa3d8b6f3fb547091aae08290004e78277eb2ffb854ea21e91a2505050505050565b5f611af36123fa565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015611b1f5750825b90505f8267ffffffffffffffff166001148015611b3b5750303b155b905081158015611b49575080155b15611b80576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611be15784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b611be9612422565b611bf1612422565b611bfd8989898961242a565b8315611c5e5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260068301602052604081205480611cdf576040517f4b73e4e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016109a6565b5f81815260078501602052604090205460ff16611d40576040517fae111e9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016109a6565b9392505050565b611d5181336129f2565b50565b5f5f829050601f81511115611d9757826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016109a69190613509565b8051611da282613c4e565b179392505050565b5f611d408383612a98565b5f83471015611df9576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018590526044016109a6565b81515f03611e33576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611e54576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611d40576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611ec3848473ffffffffffffffffffffffffffffffffffffffff8516612aa3565b949350505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611fe0575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611f7c3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105ca565b5f9150506105ca565b5f6105ca82612abf565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611fe0575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105ca565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061219c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166121837f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156121d3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61118b81611d47565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612264575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261226191810190613c93565b60015b6122b2576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016109a6565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461230e576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016109a6565b610e8f8383612ac9565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146121d3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8080806123958686612b2b565b909450925050505b9250929050565b60605f6123b083612b54565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f8080806123958686612b94565b5f611d408383612bcc565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006105ca565b6121d3612be8565b612432612be8565b73ffffffffffffffffffffffffffffffffffffffff84166124a1576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024016109a6565b73ffffffffffffffffffffffffffffffffffffffff8316612510576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f646500000000000000000000000000000000000060048201526024016109a6565b73ffffffffffffffffffffffffffffffffffffffff821661257f576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c00000000000000000000000000000000000000000060048201526024016109a6565b73ffffffffffffffffffffffffffffffffffffffff81166125ee576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c00000000000000000000000000000000000000000000000060048201526024016109a6565b8273ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015612636573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261267b91908101906138e9565b515f036126d6576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f646520636f64650000000000000000000000000060048201526024016109a6565b8173ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa15801561271e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526127639190810190613caa565b5173ffffffffffffffffffffffffffffffffffffffff8381169116146127d7576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c20666163657400000000000000000000000000000060048201526024016109a6565b8073ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa15801561281f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526128649190810190613caa565b5173ffffffffffffffffffffffffffffffffffffffff8281169116146128d8576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c20666163657400000000000000000000000000000000000060048201526024016109a6565b7f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600805473ffffffffffffffffffffffffffffffffffffffff8581167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161783557f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a060180548683169084161790557f5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a06028054918516919092161790556129bd7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f612c26565b6129c75f86611ecb565b506109087faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c86611ecb565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661118b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016109a6565b5f611d408383612cc7565b5f8281526002840160205260408120829055611ec38484612cde565b5f6105ca82612ce9565b612ad282612cf2565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612b2357610e8f8282612dc0565b61118b612e3f565b5f8080612b388585612e77565b5f81815260029690960160205260409095205494959350505050565b5f60ff8216601f8111156105ca576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600283016020526040812054819080612bc157612bb58585612a98565b92505f915061239d9050565b60019250905061239d565b5f8181526002830160205260408120819055611d408383612e82565b612bf0612e8d565b6121d3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f612c7f845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b5f8181526001830160205260408120541515611d40565b5f611d408383612eab565b5f6105ca825490565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612d5a576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109a6565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612de99190613da0565b5f60405180830381855af49150503d805f8114612e21576040519150601f19603f3d011682016040523d82523d5f602084013e612e26565b606091505b5091509150612e36858383612ef7565b95945050505050565b34156121d3576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611d408383612f86565b5f611d408383612fac565b5f612e966123fa565b5468010000000000000000900460ff16919050565b5f818152600183016020526040812054612ef057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556105ca565b505f6105ca565b606082612f0c57612f0782613086565b611d40565b8151158015612f30575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612f7f576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016109a6565b5080611d40565b5f825f018281548110612f9b57612f9b613888565b905f5260205f200154905092915050565b5f8181526001830160205260408120548015611fe0575f612fce600183613dab565b85549091505f90612fe190600190613dab565b9050808214613040575f865f018281548110612fff57612fff613888565b905f5260205f200154905080875f01848154811061301f5761301f613888565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061305157613051613de3565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506105ca565b8051156130965780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081168114611d51575f5ffd5b5f60208284031215613105575f5ffd5b8135611d40816130c8565b60038110611d51575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611d51575f5ffd5b5f5f83601f84011261314d575f5ffd5b50813567ffffffffffffffff811115613164575f5ffd5b60208301915083602082850101111561239d575f5ffd5b5f5f5f5f5f6080868803121561318f575f5ffd5b853561319a81613110565b94506020860135935060408601356131b18161311c565b9250606086013567ffffffffffffffff8111156131cc575f5ffd5b6131d88882890161313d565b969995985093965092949392505050565b5f5f5f5f5f5f608087890312156131fe575f5ffd5b86356132098161311c565b955060208701356132198161311c565b9450604087013567ffffffffffffffff811115613234575f5ffd5b61324089828a0161313d565b909550935050606087013567ffffffffffffffff81111561325f575f5ffd5b8701601f8101891361326f575f5ffd5b803567ffffffffffffffff811115613285575f5ffd5b8960208260051b8401011115613299575f5ffd5b60208201935080925050509295509295509295565b5f602082840312156132be575f5ffd5b5035919050565b5f5f604083850312156132d6575f5ffd5b8235915060208301356132e88161311c565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff81118282101715613343576133436132f3565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613390576133906132f3565b604052919050565b5f67ffffffffffffffff8211156133b1576133b16132f3565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126133ec575f5ffd5b8135602083015f6134046133ff84613398565b613349565b9050828152858383011115613417575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f60408385031215613442575f5ffd5b823561344d8161311c565b9150602083013567ffffffffffffffff811115613468575f5ffd5b613474858286016133dd565b9150509250929050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081525f6134dc606083018661347e565b73ffffffffffffffffffffffffffffffffffffffff94909416602083015250901515604090910152919050565b602081525f611d40602083018461347e565b5f6020828403121561352b575f5ffd5b8135611d408161311c565b604081525f613548604083018561347e565b905082151560208301529392505050565b5f60208284031215613569575f5ffd5b813567ffffffffffffffff81111561357f575f5ffd5b611ec3848285016133dd565b5f604082016040835280855180835260608501915060608160051b8601019250602087015f5b82811015613600577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08786030184526135eb85835161347e565b945060209384019391909101906001016135b1565b5050505082810360208401528084518083526020830191506020860192505f5b8181101561365457835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613620565b50909695505050505050565b8035801515811461366f575f5ffd5b919050565b5f5f60408385031215613685575f5ffd5b823567ffffffffffffffff81111561369b575f5ffd5b6136a7858286016133dd565b9250506136b660208401613660565b90509250929050565b5f5f5f5f608085870312156136d2575f5ffd5b84356136dd8161311c565b935060208501356136ed8161311c565b925060408501356136fd8161311c565b9150606085013561370d8161311c565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f5f60408385031215613756575f5ffd5b82356137618161311c565b946020939093013593505050565b5f5f60408385031215613780575f5ffd5b61344d83613660565b5f67ffffffffffffffff8211156137a2576137a26132f3565b5060051b60200190565b5f82601f8301126137bb575f5ffd5b81516137c96133ff82613789565b8082825260208201915060208360051b8601019250858311156137ea575f5ffd5b602085015b838110156138075780518352602092830192016137ef565b5095945050505050565b5f5f5f60608486031215613823575f5ffd5b835161382e8161311c565b602085015190935067ffffffffffffffff81111561384a575f5ffd5b613856868287016137ac565b925050604084015167ffffffffffffffff811115613872575f5ffd5b61387e868287016137ac565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f606084860312156138c7575f5ffd5b83516138d28161311c565b602085015160409095015190969495509392505050565b5f602082840312156138f9575f5ffd5b815167ffffffffffffffff81111561390f575f5ffd5b8201601f8101841361391f575f5ffd5b805161392d6133ff82613398565b818152856020838501011115613941575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b8183526020830192505f815f5b848110156139b557813561397e816130c8565b7fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195919091019060010161396b565b5093949350505050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f60c0820173ffffffffffffffffffffffffffffffffffffffff8b16835273ffffffffffffffffffffffffffffffffffffffff8a16602084015260c060408401528088825260e08401905060e08960051b8501019150895f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18c3603015b8b821015613bbe577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208786030184528235818112613ac0575f5ffd5b8d018035613acd8161311c565b73ffffffffffffffffffffffffffffffffffffffff1686526020810135613af381613110565b60038110613b28577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208701526040810135368290037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112613b62575f5ffd5b0160208101903567ffffffffffffffff811115613b7d575f5ffd5b8060051b3603821315613b8e575f5ffd5b60606040880152613ba360608801828461395e565b96505050602083019250602084019350600182019150613a84565b505050508281036060840152613bd58187896139bf565b915050613bfa608083018573ffffffffffffffffffffffffffffffffffffffff169052565b73ffffffffffffffffffffffffffffffffffffffff831660a08301529998505050505050505050565b5f81518060208401855e5f93019283525090919050565b5f611ec3613c488386613c23565b84613c23565b80516020808301519190811015613c8d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b5f60208284031215613ca3575f5ffd5b5051919050565b5f60208284031215613cba575f5ffd5b815167ffffffffffffffff811115613cd0575f5ffd5b820160608185031215613ce1575f5ffd5b613ce9613320565b8151613cf48161311c565b81526020820151613d0481613110565b6020820152604082015167ffffffffffffffff811115613d22575f5ffd5b80830192505084601f830112613d36575f5ffd5b8151613d446133ff82613789565b8082825260208201915060208360051b860101925087831115613d65575f5ffd5b6020850194505b82851015613d90578451613d7f816130c8565b825260209485019490910190613d6c565b6040840152509095945050505050565b5f611d408284613c23565b818103818111156105ca577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea264697066735822122038b854df30be349a23ce08ac3891c986adfec070631090b02607567b8620a00264736f6c634300081c0033",
}

// TokenForgeFactory is an auto generated Go binding around an Ethereum contract.
type TokenForgeFactory struct {
	abi abi.ABI
}

// NewTokenForgeFactory creates a new instance of TokenForgeFactory.
func NewTokenForgeFactory() *TokenForgeFactory {
	parsed, err := TokenForgeFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TokenForgeFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TokenForgeFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) PackDEFAULTADMINROLE() []byte {
	enc, err := tokenForgeFactory.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := tokenForgeFactory.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (tokenForgeFactory *TokenForgeFactory) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := tokenForgeFactory.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (tokenForgeFactory *TokenForgeFactory) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := tokenForgeFactory.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAddService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x217edd34.
//
// Solidity: function addService(address owner, address validator, string service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (tokenForgeFactory *TokenForgeFactory) PackAddService(owner common.Address, validator common.Address, service string, addCuts []IDiamondCutFacetCut) []byte {
	enc, err := tokenForgeFactory.abi.Pack("addService", owner, validator, service, addCuts)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAddService is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x217edd34.
//
// Solidity: function addService(address owner, address validator, string service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (tokenForgeFactory *TokenForgeFactory) UnpackAddService(data []byte) (common.Address, error) {
	out, err := tokenForgeFactory.abi.Unpack("addService", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackAlertBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ea113d0.
//
// Solidity: function alertBurn(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (tokenForgeFactory *TokenForgeFactory) PackAlertBurn(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("alertBurn", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAlertMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe98a5784.
//
// Solidity: function alertMint(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (tokenForgeFactory *TokenForgeFactory) PackAlertMint(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("alertMint", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAlertTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b7e40c6.
//
// Solidity: function alertTransfer(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (tokenForgeFactory *TokenForgeFactory) PackAlertTransfer(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("alertTransfer", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe57997d2.
//
// Solidity: function allForges() view returns(string[], address[])
func (tokenForgeFactory *TokenForgeFactory) PackAllForges() []byte {
	enc, err := tokenForgeFactory.abi.Pack("allForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// AllForgesOutput serves as a container for the return parameters of contract
// method AllForges.
type AllForgesOutput struct {
	Arg0 []string
	Arg1 []common.Address
}

// UnpackAllForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe57997d2.
//
// Solidity: function allForges() view returns(string[], address[])
func (tokenForgeFactory *TokenForgeFactory) UnpackAllForges(data []byte) (AllForgesOutput, error) {
	out, err := tokenForgeFactory.abi.Unpack("allForges", data)
	outstruct := new(AllForgesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, err

}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(string service, address forge, bool running)
func (tokenForgeFactory *TokenForgeFactory) PackForgeByIndex(index *big.Int) []byte {
	enc, err := tokenForgeFactory.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// ForgeByIndexOutput serves as a container for the return parameters of contract
// method ForgeByIndex.
type ForgeByIndexOutput struct {
	Service string
	Forge   common.Address
	Running bool
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(string service, address forge, bool running)
func (tokenForgeFactory *TokenForgeFactory) UnpackForgeByIndex(data []byte) (ForgeByIndexOutput, error) {
	out, err := tokenForgeFactory.abi.Unpack("forgeByIndex", data)
	outstruct := new(ForgeByIndexOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Service = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Forge = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Running = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackForgeByService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3de1d7f.
//
// Solidity: function forgeByService(string service) view returns(address forge, bool running)
func (tokenForgeFactory *TokenForgeFactory) PackForgeByService(service string) []byte {
	enc, err := tokenForgeFactory.abi.Pack("forgeByService", service)
	if err != nil {
		panic(err)
	}
	return enc
}

// ForgeByServiceOutput serves as a container for the return parameters of contract
// method ForgeByService.
type ForgeByServiceOutput struct {
	Forge   common.Address
	Running bool
}

// UnpackForgeByService is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3de1d7f.
//
// Solidity: function forgeByService(string service) view returns(address forge, bool running)
func (tokenForgeFactory *TokenForgeFactory) UnpackForgeByService(data []byte) (ForgeByServiceOutput, error) {
	out, err := tokenForgeFactory.abi.Unpack("forgeByService", data)
	outstruct := new(ForgeByServiceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Forge = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Running = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := tokenForgeFactory.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (tokenForgeFactory *TokenForgeFactory) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (tokenForgeFactory *TokenForgeFactory) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (tokenForgeFactory *TokenForgeFactory) UnpackHasRole(data []byte) (bool, error) {
	out, err := tokenForgeFactory.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _forgeProxyCode, address _diamondImpl, address _baseImpl) returns()
func (tokenForgeFactory *TokenForgeFactory) PackInitialize(owner common.Address, forgeProxyCode common.Address, diamondImpl common.Address, baseImpl common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("initialize", owner, forgeProxyCode, diamondImpl, baseImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsRunningForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7525cba.
//
// Solidity: function isRunningForge(address forge) view returns(string service, bool running)
func (tokenForgeFactory *TokenForgeFactory) PackIsRunningForge(forge common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("isRunningForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// IsRunningForgeOutput serves as a container for the return parameters of contract
// method IsRunningForge.
type IsRunningForgeOutput struct {
	Service string
	Running bool
}

// UnpackIsRunningForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc7525cba.
//
// Solidity: function isRunningForge(address forge) view returns(string service, bool running)
func (tokenForgeFactory *TokenForgeFactory) UnpackIsRunningForge(data []byte) (IsRunningForgeOutput, error) {
	out, err := tokenForgeFactory.abi.Unpack("isRunningForge", data)
	outstruct := new(IsRunningForgeOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Service = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Running = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, err

}

// PackLengthAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f44ec09.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (tokenForgeFactory *TokenForgeFactory) PackLengthAllForges() []byte {
	enc, err := tokenForgeFactory.abi.Pack("lengthAllForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackLengthAllForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f44ec09.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (tokenForgeFactory *TokenForgeFactory) UnpackLengthAllForges(data []byte) (*big.Int, error) {
	out, err := tokenForgeFactory.abi.Unpack("lengthAllForges", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPauseService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf4d8cab9.
//
// Solidity: function pauseService(string service, bool paused) returns()
func (tokenForgeFactory *TokenForgeFactory) PackPauseService(service string, paused bool) []byte {
	enc, err := tokenForgeFactory.abi.Pack("pauseService", service, paused)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) PackProxiableUUID() []byte {
	enc, err := tokenForgeFactory.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenForgeFactory *TokenForgeFactory) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := tokenForgeFactory.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRemoveService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf51acaea.
//
// Solidity: function removeService(string service) returns()
func (tokenForgeFactory *TokenForgeFactory) PackRemoveService(service string) []byte {
	enc, err := tokenForgeFactory.abi.Pack("removeService", service)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (tokenForgeFactory *TokenForgeFactory) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (tokenForgeFactory *TokenForgeFactory) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenForgeFactory.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenForgeFactory *TokenForgeFactory) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenForgeFactory *TokenForgeFactory) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := tokenForgeFactory.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (tokenForgeFactory *TokenForgeFactory) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := tokenForgeFactory.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TokenForgeFactoryERC1155Burned represents a ERC1155Burned event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC1155Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC1155BurnedEventName = "ERC1155Burned"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC1155Burned) ContractEventName() string {
	return TokenForgeFactoryERC1155BurnedEventName
}

// UnpackERC1155BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 tokenID, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC1155BurnedEvent(log *types.Log) (*TokenForgeFactoryERC1155Burned, error) {
	event := "ERC1155Burned"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC1155Burned)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC1155Minted represents a ERC1155Minted event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC1155Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC1155MintedEventName = "ERC1155Minted"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC1155Minted) ContractEventName() string {
	return TokenForgeFactoryERC1155MintedEventName
}

// UnpackERC1155MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC1155MintedEvent(log *types.Log) (*TokenForgeFactoryERC1155Minted, error) {
	event := "ERC1155Minted"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC1155Minted)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC1155Transferred represents a ERC1155Transferred event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC1155Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC1155TransferredEventName = "ERC1155Transferred"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC1155Transferred) ContractEventName() string {
	return TokenForgeFactoryERC1155TransferredEventName
}

// UnpackERC1155TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC1155TransferredEvent(log *types.Log) (*TokenForgeFactoryERC1155Transferred, error) {
	event := "ERC1155Transferred"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC1155Transferred)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC20Burned represents a ERC20Burned event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC20Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC20BurnedEventName = "ERC20Burned"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC20Burned) ContractEventName() string {
	return TokenForgeFactoryERC20BurnedEventName
}

// UnpackERC20BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC20BurnedEvent(log *types.Log) (*TokenForgeFactoryERC20Burned, error) {
	event := "ERC20Burned"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC20Burned)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC20Minted represents a ERC20Minted event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC20Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC20MintedEventName = "ERC20Minted"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC20Minted) ContractEventName() string {
	return TokenForgeFactoryERC20MintedEventName
}

// UnpackERC20MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC20MintedEvent(log *types.Log) (*TokenForgeFactoryERC20Minted, error) {
	event := "ERC20Minted"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC20Minted)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC20Transferred represents a ERC20Transferred event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC20Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC20TransferredEventName = "ERC20Transferred"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC20Transferred) ContractEventName() string {
	return TokenForgeFactoryERC20TransferredEventName
}

// UnpackERC20TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC20TransferredEvent(log *types.Log) (*TokenForgeFactoryERC20Transferred, error) {
	event := "ERC20Transferred"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC20Transferred)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC721Burned represents a ERC721Burned event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC721Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC721BurnedEventName = "ERC721Burned"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC721Burned) ContractEventName() string {
	return TokenForgeFactoryERC721BurnedEventName
}

// UnpackERC721BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 tokenID)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC721BurnedEvent(log *types.Log) (*TokenForgeFactoryERC721Burned, error) {
	event := "ERC721Burned"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC721Burned)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC721Minted represents a ERC721Minted event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC721Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC721MintedEventName = "ERC721Minted"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC721Minted) ContractEventName() string {
	return TokenForgeFactoryERC721MintedEventName
}

// UnpackERC721MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC721MintedEvent(log *types.Log) (*TokenForgeFactoryERC721Minted, error) {
	event := "ERC721Minted"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC721Minted)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryERC721Transferred represents a ERC721Transferred event raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC721Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryERC721TransferredEventName = "ERC721Transferred"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryERC721Transferred) ContractEventName() string {
	return TokenForgeFactoryERC721TransferredEventName
}

// UnpackERC721TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC721TransferredEvent(log *types.Log) (*TokenForgeFactoryERC721Transferred, error) {
	event := "ERC721Transferred"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryERC721Transferred)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryInitialized represents a Initialized event raised by the TokenForgeFactory contract.
type TokenForgeFactoryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryInitialized) ContractEventName() string {
	return TokenForgeFactoryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (tokenForgeFactory *TokenForgeFactory) UnpackInitializedEvent(log *types.Log) (*TokenForgeFactoryInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryInitialized)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the TokenForgeFactory contract.
type TokenForgeFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryRoleAdminChanged) ContractEventName() string {
	return TokenForgeFactoryRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (tokenForgeFactory *TokenForgeFactory) UnpackRoleAdminChangedEvent(log *types.Log) (*TokenForgeFactoryRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryRoleGranted represents a RoleGranted event raised by the TokenForgeFactory contract.
type TokenForgeFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryRoleGranted) ContractEventName() string {
	return TokenForgeFactoryRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (tokenForgeFactory *TokenForgeFactory) UnpackRoleGrantedEvent(log *types.Log) (*TokenForgeFactoryRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryRoleGranted)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryRoleRevoked represents a RoleRevoked event raised by the TokenForgeFactory contract.
type TokenForgeFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryRoleRevoked) ContractEventName() string {
	return TokenForgeFactoryRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (tokenForgeFactory *TokenForgeFactory) UnpackRoleRevokedEvent(log *types.Log) (*TokenForgeFactoryRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryRoleRevoked)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryServicePaused represents a ServicePaused event raised by the TokenForgeFactory contract.
type TokenForgeFactoryServicePaused struct {
	Service [32]byte
	Paused  bool
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryServicePausedEventName = "ServicePaused"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryServicePaused) ContractEventName() string {
	return TokenForgeFactoryServicePausedEventName
}

// UnpackServicePausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServicePaused(bytes32 indexed service, bool paused)
func (tokenForgeFactory *TokenForgeFactory) UnpackServicePausedEvent(log *types.Log) (*TokenForgeFactoryServicePaused, error) {
	event := "ServicePaused"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryServicePaused)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryServiceRegistered represents a ServiceRegistered event raised by the TokenForgeFactory contract.
type TokenForgeFactoryServiceRegistered struct {
	Service [32]byte
	Forge   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryServiceRegisteredEventName = "ServiceRegistered"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryServiceRegistered) ContractEventName() string {
	return TokenForgeFactoryServiceRegisteredEventName
}

// UnpackServiceRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServiceRegistered(bytes32 indexed service, address indexed forge)
func (tokenForgeFactory *TokenForgeFactory) UnpackServiceRegisteredEvent(log *types.Log) (*TokenForgeFactoryServiceRegistered, error) {
	event := "ServiceRegistered"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryServiceRegistered)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryServiceUnregistered represents a ServiceUnregistered event raised by the TokenForgeFactory contract.
type TokenForgeFactoryServiceUnregistered struct {
	Service [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryServiceUnregisteredEventName = "ServiceUnregistered"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryServiceUnregistered) ContractEventName() string {
	return TokenForgeFactoryServiceUnregisteredEventName
}

// UnpackServiceUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServiceUnregistered(bytes32 indexed service)
func (tokenForgeFactory *TokenForgeFactory) UnpackServiceUnregisteredEvent(log *types.Log) (*TokenForgeFactoryServiceUnregistered, error) {
	event := "ServiceUnregistered"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryServiceUnregistered)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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

// TokenForgeFactoryUpgraded represents a Upgraded event raised by the TokenForgeFactory contract.
type TokenForgeFactoryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TokenForgeFactoryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TokenForgeFactoryUpgraded) ContractEventName() string {
	return TokenForgeFactoryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (tokenForgeFactory *TokenForgeFactory) UnpackUpgradedEvent(log *types.Log) (*TokenForgeFactoryUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != tokenForgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenForgeFactoryUpgraded)
	if len(log.Data) > 0 {
		if err := tokenForgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenForgeFactory.abi.Events[event].Inputs {
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
func (tokenForgeFactory *TokenForgeFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeFactoryAlreadyUsedService"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeFactoryAlreadyUsedServiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeFactoryCallerIsNotForge"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeFactoryCallerIsNotForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeFactoryCallerIsPausedForge"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeFactoryCallerIsPausedForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeFactoryInvalidData"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeFactoryInvalidDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeFactoryServiceNotFound"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeFactoryServiceNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["TokenForgeInvalidTokenType"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackTokenForgeInvalidTokenTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenForgeFactory.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return tokenForgeFactory.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TokenForgeFactoryAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the TokenForgeFactory contract.
type TokenForgeFactoryAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func TokenForgeFactoryAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (tokenForgeFactory *TokenForgeFactory) UnpackAccessControlBadConfirmationError(raw []byte) (*TokenForgeFactoryAccessControlBadConfirmation, error) {
	out := new(TokenForgeFactoryAccessControlBadConfirmation)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the TokenForgeFactory contract.
type TokenForgeFactoryAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func TokenForgeFactoryAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (tokenForgeFactory *TokenForgeFactory) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*TokenForgeFactoryAccessControlUnauthorizedAccount, error) {
	out := new(TokenForgeFactoryAccessControlUnauthorizedAccount)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryAddressEmptyCode represents a AddressEmptyCode error raised by the TokenForgeFactory contract.
type TokenForgeFactoryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TokenForgeFactoryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (tokenForgeFactory *TokenForgeFactory) UnpackAddressEmptyCodeError(raw []byte) (*TokenForgeFactoryAddressEmptyCode, error) {
	out := new(TokenForgeFactoryAddressEmptyCode)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the TokenForgeFactory contract.
type TokenForgeFactoryCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func TokenForgeFactoryCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (tokenForgeFactory *TokenForgeFactory) UnpackCreate2EmptyBytecodeError(raw []byte) (*TokenForgeFactoryCreate2EmptyBytecode, error) {
	out := new(TokenForgeFactoryCreate2EmptyBytecode)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TokenForgeFactoryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (tokenForgeFactory *TokenForgeFactory) UnpackERC1967InvalidImplementationError(raw []byte) (*TokenForgeFactoryERC1967InvalidImplementation, error) {
	out := new(TokenForgeFactoryERC1967InvalidImplementation)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryERC1967NonPayable represents a ERC1967NonPayable error raised by the TokenForgeFactory contract.
type TokenForgeFactoryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TokenForgeFactoryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (tokenForgeFactory *TokenForgeFactory) UnpackERC1967NonPayableError(raw []byte) (*TokenForgeFactoryERC1967NonPayable, error) {
	out := new(TokenForgeFactoryERC1967NonPayable)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryFailedCall represents a FailedCall error raised by the TokenForgeFactory contract.
type TokenForgeFactoryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TokenForgeFactoryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (tokenForgeFactory *TokenForgeFactory) UnpackFailedCallError(raw []byte) (*TokenForgeFactoryFailedCall, error) {
	out := new(TokenForgeFactoryFailedCall)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryFailedDeployment represents a FailedDeployment error raised by the TokenForgeFactory contract.
type TokenForgeFactoryFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func TokenForgeFactoryFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (tokenForgeFactory *TokenForgeFactory) UnpackFailedDeploymentError(raw []byte) (*TokenForgeFactoryFailedDeployment, error) {
	out := new(TokenForgeFactoryFailedDeployment)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryInsufficientBalance represents a InsufficientBalance error raised by the TokenForgeFactory contract.
type TokenForgeFactoryInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func TokenForgeFactoryInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (tokenForgeFactory *TokenForgeFactory) UnpackInsufficientBalanceError(raw []byte) (*TokenForgeFactoryInsufficientBalance, error) {
	out := new(TokenForgeFactoryInsufficientBalance)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryInvalidInitialization represents a InvalidInitialization error raised by the TokenForgeFactory contract.
type TokenForgeFactoryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TokenForgeFactoryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (tokenForgeFactory *TokenForgeFactory) UnpackInvalidInitializationError(raw []byte) (*TokenForgeFactoryInvalidInitialization, error) {
	out := new(TokenForgeFactoryInvalidInitialization)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryInvalidShortString represents a InvalidShortString error raised by the TokenForgeFactory contract.
type TokenForgeFactoryInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func TokenForgeFactoryInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (tokenForgeFactory *TokenForgeFactory) UnpackInvalidShortStringError(raw []byte) (*TokenForgeFactoryInvalidShortString, error) {
	out := new(TokenForgeFactoryInvalidShortString)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryNotInitializing represents a NotInitializing error raised by the TokenForgeFactory contract.
type TokenForgeFactoryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TokenForgeFactoryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (tokenForgeFactory *TokenForgeFactory) UnpackNotInitializingError(raw []byte) (*TokenForgeFactoryNotInitializing, error) {
	out := new(TokenForgeFactoryNotInitializing)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryStringTooLong represents a StringTooLong error raised by the TokenForgeFactory contract.
type TokenForgeFactoryStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func TokenForgeFactoryStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (tokenForgeFactory *TokenForgeFactory) UnpackStringTooLongError(raw []byte) (*TokenForgeFactoryStringTooLong, error) {
	out := new(TokenForgeFactoryStringTooLong)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeFactoryAlreadyUsedService represents a TokenForgeFactory__AlreadyUsedService error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeFactoryAlreadyUsedService struct {
	Service [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__AlreadyUsedService(bytes32 service)
func TokenForgeFactoryTokenForgeFactoryAlreadyUsedServiceErrorID() common.Hash {
	return common.HexToHash("0x5502583a4e5df3d611d0359980fb46a75c8ce25278a66f5253b4cd467c300887")
}

// UnpackTokenForgeFactoryAlreadyUsedServiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__AlreadyUsedService(bytes32 service)
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeFactoryAlreadyUsedServiceError(raw []byte) (*TokenForgeFactoryTokenForgeFactoryAlreadyUsedService, error) {
	out := new(TokenForgeFactoryTokenForgeFactoryAlreadyUsedService)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryAlreadyUsedService", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeFactoryCallerIsNotForge represents a TokenForgeFactory__CallerIsNotForge error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeFactoryCallerIsNotForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__CallerIsNotForge(address caller)
func TokenForgeFactoryTokenForgeFactoryCallerIsNotForgeErrorID() common.Hash {
	return common.HexToHash("0x4b73e4e6b8076eaaeddcbe95bc8fdf6efa1a2606fc078882372b3a4389ba7e82")
}

// UnpackTokenForgeFactoryCallerIsNotForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__CallerIsNotForge(address caller)
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeFactoryCallerIsNotForgeError(raw []byte) (*TokenForgeFactoryTokenForgeFactoryCallerIsNotForge, error) {
	out := new(TokenForgeFactoryTokenForgeFactoryCallerIsNotForge)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryCallerIsNotForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeFactoryCallerIsPausedForge represents a TokenForgeFactory__CallerIsPausedForge error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeFactoryCallerIsPausedForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__CallerIsPausedForge(address caller)
func TokenForgeFactoryTokenForgeFactoryCallerIsPausedForgeErrorID() common.Hash {
	return common.HexToHash("0xae111e923268b59d275e8c9ba5e0305328e43fc8b38f3601919e1045c4eacf8c")
}

// UnpackTokenForgeFactoryCallerIsPausedForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__CallerIsPausedForge(address caller)
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeFactoryCallerIsPausedForgeError(raw []byte) (*TokenForgeFactoryTokenForgeFactoryCallerIsPausedForge, error) {
	out := new(TokenForgeFactoryTokenForgeFactoryCallerIsPausedForge)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryCallerIsPausedForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeFactoryInvalidData represents a TokenForgeFactory__InvalidData error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeFactoryInvalidData struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__InvalidData(bytes32 field)
func TokenForgeFactoryTokenForgeFactoryInvalidDataErrorID() common.Hash {
	return common.HexToHash("0x435db3481d8eb115a5d9fe457169c1d3dca88f1c1a6c4e6ee3b67d659593d3ca")
}

// UnpackTokenForgeFactoryInvalidDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__InvalidData(bytes32 field)
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeFactoryInvalidDataError(raw []byte) (*TokenForgeFactoryTokenForgeFactoryInvalidData, error) {
	out := new(TokenForgeFactoryTokenForgeFactoryInvalidData)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryInvalidData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeFactoryServiceNotFound represents a TokenForgeFactory__ServiceNotFound error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeFactoryServiceNotFound struct {
	Service [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__ServiceNotFound(bytes32 service)
func TokenForgeFactoryTokenForgeFactoryServiceNotFoundErrorID() common.Hash {
	return common.HexToHash("0x2a188b7fac945d1d82f2b68ea1d938569eb745b8636b1c8a173274cce308b9cc")
}

// UnpackTokenForgeFactoryServiceNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__ServiceNotFound(bytes32 service)
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeFactoryServiceNotFoundError(raw []byte) (*TokenForgeFactoryTokenForgeFactoryServiceNotFound, error) {
	out := new(TokenForgeFactoryTokenForgeFactoryServiceNotFound)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryServiceNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryTokenForgeInvalidTokenType represents a TokenForge__InvalidTokenType error raised by the TokenForgeFactory contract.
type TokenForgeFactoryTokenForgeInvalidTokenType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForge__InvalidTokenType()
func TokenForgeFactoryTokenForgeInvalidTokenTypeErrorID() common.Hash {
	return common.HexToHash("0x9570ddf58a1418a686536d6c200028d56382c32ca6637e8defd204d5722e94e9")
}

// UnpackTokenForgeInvalidTokenTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForge__InvalidTokenType()
func (tokenForgeFactory *TokenForgeFactory) UnpackTokenForgeInvalidTokenTypeError(raw []byte) (*TokenForgeFactoryTokenForgeInvalidTokenType, error) {
	out := new(TokenForgeFactoryTokenForgeInvalidTokenType)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "TokenForgeInvalidTokenType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TokenForgeFactory contract.
type TokenForgeFactoryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TokenForgeFactoryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (tokenForgeFactory *TokenForgeFactory) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TokenForgeFactoryUUPSUnauthorizedCallContext, error) {
	out := new(TokenForgeFactoryUUPSUnauthorizedCallContext)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenForgeFactoryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TokenForgeFactory contract.
type TokenForgeFactoryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TokenForgeFactoryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (tokenForgeFactory *TokenForgeFactory) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TokenForgeFactoryUUPSUnsupportedProxiableUUID, error) {
	out := new(TokenForgeFactoryUUPSUnsupportedProxiableUUID)
	if err := tokenForgeFactory.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
