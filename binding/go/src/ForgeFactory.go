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

// ForgeFactoryMetaData contains all meta data concerning the ForgeFactory contract.
var ForgeFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"addCuts\",\"type\":\"tuple[]\"}],\"name\":\"addService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allForges\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"forgeByService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forgeProxyCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_diamondImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isRunningForge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lengthAllForges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"pauseService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ERC20FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20TransferredFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"ServicePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"ServiceUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__AlreadyUsedService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsNotForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsPausedForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__InvalidData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__ServiceNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenForge__InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ForgeFactory",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613bdc6100f95f395f81816120640152818161208d01526122ad0152613bdc5ff3fe608060405260043610610178575f3560e01c806391d14854116100d1578063d547741f1161007c578063ec0b6a4d11610057578063ec0b6a4d14610503578063ecf5bcc11461054e578063f8c8765e1461056d575f5ffd5b8063d547741f146104a3578063e57997d2146104c2578063e98a5784146104e4575f5ffd5b8063ad3cb1cc116100ac578063ad3cb1cc146103fb578063c7525cba14610450578063d3884c3f14610484575f5ffd5b806391d14854146103275780639ca92df914610397578063a217fddf146103e8575f5ffd5b806336568abe116101315780634f1ef2861161010c5780634f1ef286146102bc57806352d1902d146102cf5780637f765bfe146102e3575f5ffd5b806336568abe1461025f5780633a15d0701461027e5780633ea113d01461029d575f5ffd5b8063248a9ca311610161578063248a9ca3146101d15780632f2ff15d1461022c5780632f44ec091461024b575f5ffd5b806301ffc9a71461017c5780630b7e40c6146101b0575b5f5ffd5b348015610187575f5ffd5b5061019b61019636600461304b565b61058c565b60405190151581526020015b60405180910390f35b3480156101bb575f5ffd5b506101cf6101ca366004613093565b610624565b005b3480156101dc575f5ffd5b5061021e6101eb36600461312c565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101a7565b348015610237575f5ffd5b506101cf610246366004613143565b6109ea565b348015610256575f5ffd5b5061021e610a33565b34801561026a575f5ffd5b506101cf610279366004613143565b610a62565b348015610289575f5ffd5b506101cf610298366004613093565b610ac0565b3480156102a8575f5ffd5b506101cf6102b7366004613093565b610b87565b6101cf6102ca3660046132ad565b610e69565b3480156102da575f5ffd5b5061021e610e88565b3480156102ee575f5ffd5b506103026102fd3660046132fa565b610eb6565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b348015610332575f5ffd5b5061019b610341366004613143565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156103a2575f5ffd5b506103b66103b136600461312c565b61131c565b6040805193845273ffffffffffffffffffffffffffffffffffffffff90921660208401521515908201526060016101a7565b3480156103f3575f5ffd5b5061021e5f81565b348015610406575f5ffd5b506104436040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a79190613385565b34801561045b575f5ffd5b5061046f61046a3660046133d8565b61138b565b604080519283529015156020830152016101a7565b34801561048f575f5ffd5b506101cf61049e36600461312c565b611413565b3480156104ae575f5ffd5b506101cf6104bd366004613143565b611562565b3480156104cd575f5ffd5b506104d66115a5565b6040516101a79291906133f3565b3480156104ef575f5ffd5b506101cf6104fe366004613093565b611700565b34801561050e575f5ffd5b5061052261051d36600461312c565b6119e2565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683529015156020830152016101a7565b348015610559575f5ffd5b506101cf61056836600461349d565b611a57565b348015610578575f5ffd5b506101cf6105873660046134c7565b611bde565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061061e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f6106507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005b33611d5d565b90505f86600281111561066557610665613520565b0361076d575f80808061067a8688018861354d565b93509350935093508373ffffffffffffffffffffffffffffffffffffffff1689867f86d733c2654f9db1a830cfbbf2180ee3329101d853428a7ead4e85ddb3230d608b876040516106ed92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a48015610764578773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168a7fa2b448c8495be4baf0f3195bf0d82d5ec29d24aa3c5dd949969149bd797fb3158460405161075b91815260200190565b60405180910390a45b505050506109e2565b600186600281111561078157610781613520565b03610812575f8061079484860186613592565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847fa36201ae386aaacb89bc573b7c0ecc0913a93028f710f28ca29840454f446886898560405161080392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a450506109e2565b600286600281111561082657610826613520565b036109b0575f80610839848601866135bc565b915091508115610926575f5f5f8380602001905181019061085a919061365e565b815192955090935091505f5b8181101561091c578473ffffffffffffffffffffffffffffffffffffffff168c897fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c38e8886815181106108bb576108bb6136d5565b60200260200101518887815181106108d5576108d56136d5565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a4600101610866565b50505050506109a9565b5f5f5f8380602001905181019061093d9190613702565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c3906060015b60405180910390a45050505b50506109e2565b6040517f9570ddf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610a2381611e3b565b610a2d8383611e48565b50505050565b5f610a5d7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03611f66565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610ab1576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610abb8282611f70565b505050565b5f610aea7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0061064a565b90505f866002811115610aff57610aff613520565b036109b0575f808080610b148688018861354d565b93509350935093508373ffffffffffffffffffffffffffffffffffffffff1689867f14bed88907ac31fe1ba4cf8981f80fd25e29877a08d1568c7210adc31ee9bc958b876040516106ed92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b5f610bb17f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0061064a565b90505f866002811115610bc657610bc6613520565b03610c4e575f808080610bdb8688018861354d565b93509350935093508373ffffffffffffffffffffffffffffffffffffffff1689867fe5e0650048e213b8a668474e72c774ec69d8a6830270b7cc1c339fcd6a3b2e418b876040516106ed92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6001866002811115610c6257610c62613520565b03610ce4575f80610c7584860186613592565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f28fcdc3e3355f548f5fe83373acf91a82e1e561ee22797f6e91f831d2b2b3e13898560405161080392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6002866002811115610cf857610cf8613520565b036109b0575f80610d0b848601866135bc565b915091508115610dee575f5f5f83806020019051810190610d2c919061365e565b815192955090935091505f5b8181101561091c578473ffffffffffffffffffffffffffffffffffffffff168c897f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b8e888681518110610d8d57610d8d6136d5565b6020026020010151888781518110610da757610da76136d5565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a4600101610d38565b5f5f5f83806020019051810190610e059190613702565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b9060600161099d565b610e7161204c565b610e7a82612152565b610e84828261215c565b5050565b5f610e91612295565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610ee181611e3b565b73ffffffffffffffffffffffffffffffffffffffff8716610f55576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8616610fc4576040517f435db3480000000000000000000000000000000000000000000000000000000081527f76616c696461746f7200000000000000000000000000000000000000000000006004820152602401610f4c565b8461101d576040517f435db3480000000000000000000000000000000000000000000000000000000081527f73657276696365000000000000000000000000000000000000000000000000006004820152602401610f4c565b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e006110687f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0387612304565b156110a2576040517f5502583a00000000000000000000000000000000000000000000000000000000815260048101879052602401610f4c565b6112095f87835f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611111573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111569190810190613736565b8b8b8a8a8d896001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a6002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516020016111b9979695949392919061380c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526111f59291602001613a15565b60405160208183030381529060405261230f565b925073ffffffffffffffffffffffffffffffffffffffff831661127a576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f726765206465706c6f7900000000000000000000000000000000000000006004820152602401610f4c565b6112886003820187856123fb565b5073ffffffffffffffffffffffffffffffffffffffff83165f81815260068301602090815260408083208a90558983526007850190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555188917f0453e3ffce9b0d4f994d54652037d3e15d87ec29aab0aafba90ac322edc87c2591a3505095945050505050565b5f80807f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0061136a7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0386612425565b5f8281526007909301602052604090922054909691955060ff169350915050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e066020526040812054907f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00821561140d575f83815260078201602052604090205460ff1691505b50915091565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61143d81611e3b565b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005f8061148a7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0386612442565b91509150816114c8576040517f2a188b7f00000000000000000000000000000000000000000000000000000000815260048101869052602401610f4c565b6114d56003840186612450565b5073ffffffffffffffffffffffffffffffffffffffff81165f90815260068401602090815260408083208390558783526007860190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555186917f475a09e552c07ad6aa3d8b6f3fb547091aae08290004e78277eb2ffb854ea21e91a25050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461159b81611e3b565b610a2d8383611f70565b6060807f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005f6115f37f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03611f66565b90505f8167ffffffffffffffff81111561160f5761160f613171565b604051908082528060200260200182016040528015611638578160200160208202803683370190505b5090505f8267ffffffffffffffff81111561165557611655613171565b60405190808252806020026020018201604052801561167e578160200160208202803683370190505b5090505f5b838110156116f4576116986003860182612425565b8483815181106116aa576116aa6136d5565b602002602001018484815181106116c3576116c36136d5565b73ffffffffffffffffffffffffffffffffffffffff9093166020938402919091019092019190915252600101611683565b50909590945092505050565b5f61172a7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0061064a565b90505f86600281111561173f5761173f613520565b036117c7575f8080806117548688018861354d565b93509350935093508373ffffffffffffffffffffffffffffffffffffffff1689867f9612604afba70e4cf03261d7d86ca03d08911d9887aa41621dea929e34cfd7b18b876040516106ed92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60018660028111156117db576117db613520565b0361185d575f806117ee84860186613592565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f556d832108ef7e0123726b3e037bc1cd39023e8d397b242d8b36426550d8c123898560405161080392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600286600281111561187157611871613520565b036109b0575f80611884848601866135bc565b915091508115611967575f5f5f838060200190518101906118a5919061365e565b815192955090935091505f5b8181101561091c578473ffffffffffffffffffffffffffffffffffffffff168c897f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f618e888681518110611906576119066136d5565b6020026020010151888781518110611920576119206136d5565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a46001016118b1565b5f5f5f8380602001905181019061197e9190613702565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f619060600161099d565b5f807f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0081611a307f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0386612442565b945090508015611a50575f85815260078301602052604090205460ff1692505b5050915091565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c611a8181611e3b565b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00611acc7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0385612304565b611b05576040517f2a188b7f00000000000000000000000000000000000000000000000000000000815260048101859052602401610f4c565b5f84815260078201602052604090205483151560ff909116151503610a2d578215611b64575f848152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055611b9d565b5f848152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b837eb31832845d6567ecb46caeeee58acfcfddfcbba821c68dcbae9055ea4d795584604051611bd0911515815260200190565b60405180910390a250505050565b5f611be761245b565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015611c135750825b90505f8267ffffffffffffffff166001148015611c2f5750303b155b905081158015611c3d575080155b15611c74576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611cd55784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b611cdd612483565b611ce5612483565b611cf18989898961248b565b8315611d525784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260068301602052604081205480611dd3576040517f4b73e4e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610f4c565b5f81815260078501602052604090205460ff16611e34576040517fae111e9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610f4c565b9392505050565b611e458133612a29565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611f5d575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611ef93390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061061e565b5f91505061061e565b5f61061e82612acf565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611f5d575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061061e565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061211957507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166121007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612150576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610e8481611e3b565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156121e1575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526121de91810190613a29565b60015b61222f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f4c565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461228b576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610f4c565b610abb8383612ad9565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614612150576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611e348383612b3b565b5f83471015612353576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101859052604401610f4c565b81515f0361238d576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d1519811516156123ae576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611e34576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61241d848473ffffffffffffffffffffffffffffffffffffffff8516612b46565b949350505050565b5f8080806124338686612b62565b909450925050505b9250929050565b5f8080806124338686612b8b565b5f611e348383612bc3565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061061e565b612150612bdf565b612493612bdf565b73ffffffffffffffffffffffffffffffffffffffff8416612502576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610f4c565b73ffffffffffffffffffffffffffffffffffffffff8316612571576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f64650000000000000000000000000000000000006004820152602401610f4c565b73ffffffffffffffffffffffffffffffffffffffff82166125e0576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c0000000000000000000000000000000000000000006004820152602401610f4c565b73ffffffffffffffffffffffffffffffffffffffff811661264f576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c0000000000000000000000000000000000000000000000006004820152602401610f4c565b8273ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015612697573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526126dc9190810190613736565b515f03612737576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f646520636f6465000000000000000000000000006004820152602401610f4c565b8173ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa15801561277f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526127c49190810190613a40565b5173ffffffffffffffffffffffffffffffffffffffff838116911614612838576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c2066616365740000000000000000000000000000006004820152602401610f4c565b8073ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa158015612880573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526128c59190810190613a40565b5173ffffffffffffffffffffffffffffffffffffffff828116911614612939576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c2066616365740000000000000000000000000000000000006004820152602401610f4c565b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00805473ffffffffffffffffffffffffffffffffffffffff8581167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161783557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0180548683169084161790557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e028054918516919092161790556129fe5f86611e48565b506109e27faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c86611e48565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610e84576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610f4c565b5f61061e82612c1d565b612ae282612c26565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612b3357610abb8282612cf4565b610e84612d73565b5f611e348383612dab565b5f828152600284016020526040812082905561241d8484612dc2565b5f8080612b6f8585612dcd565b5f81815260029690960160205260409095205494959350505050565b5f818152600283016020526040812054819080612bb857612bac8585612b3b565b92505f915061243b9050565b60019250905061243b565b5f8181526002830160205260408120819055611e348383612dd8565b612be7612de3565b612150576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61061e825490565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612c8e576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610f4c565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612d1d9190613b36565b5f60405180830381855af49150503d805f8114612d55576040519150601f19603f3d011682016040523d82523d5f602084013e612d5a565b606091505b5091509150612d6a858383612e01565b95945050505050565b3415612150576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526001830160205260408120541515611e34565b5f611e348383612e90565b5f611e348383612edc565b5f611e348383612f02565b5f612dec61245b565b5468010000000000000000900460ff16919050565b606082612e1657612e1182612fdc565b611e34565b8151158015612e3a575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612e89576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610f4c565b5080611e34565b5f818152600183016020526040812054612ed557508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561061e565b505f61061e565b5f825f018281548110612ef157612ef16136d5565b905f5260205f200154905092915050565b5f8181526001830160205260408120548015611f5d575f612f24600183613b41565b85549091505f90612f3790600190613b41565b9050808214612f96575f865f018281548110612f5557612f556136d5565b905f5260205f200154905080875f018481548110612f7557612f756136d5565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612fa757612fa7613b79565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061061e565b805115612fec5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081168114611e45575f5ffd5b5f6020828403121561305b575f5ffd5b8135611e348161301e565b60038110611e45575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611e45575f5ffd5b5f5f5f5f5f608086880312156130a7575f5ffd5b85356130b281613066565b94506020860135935060408601356130c981613072565b9250606086013567ffffffffffffffff8111156130e4575f5ffd5b8601601f810188136130f4575f5ffd5b803567ffffffffffffffff81111561310a575f5ffd5b88602082840101111561311b575f5ffd5b959894975092955050506020019190565b5f6020828403121561313c575f5ffd5b5035919050565b5f5f60408385031215613154575f5ffd5b82359150602083013561316681613072565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156131c1576131c1613171565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561320e5761320e613171565b604052919050565b5f67ffffffffffffffff82111561322f5761322f613171565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f83011261326a575f5ffd5b813561327d61327882613216565b6131c7565b818152846020838601011115613291575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f604083850312156132be575f5ffd5b82356132c981613072565b9150602083013567ffffffffffffffff8111156132e4575f5ffd5b6132f08582860161325b565b9150509250929050565b5f5f5f5f5f6080868803121561330e575f5ffd5b853561331981613072565b9450602086013561332981613072565b935060408601359250606086013567ffffffffffffffff81111561334b575f5ffd5b8601601f8101881361335b575f5ffd5b803567ffffffffffffffff811115613371575f5ffd5b8860208260051b840101111561311b575f5ffd5b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f602082840312156133e8575f5ffd5b8135611e3481613072565b604080825283519082018190525f9060208501906060840190835b8181101561342c57835183526020938401939092019160010161340e565b5050838103602080860191909152855180835291810192508501905f5b8181101561347d57825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613449565b50919695505050505050565b80358015158114613498575f5ffd5b919050565b5f5f604083850312156134ae575f5ffd5b823591506134be60208401613489565b90509250929050565b5f5f5f5f608085870312156134da575f5ffd5b84356134e581613072565b935060208501356134f581613072565b9250604085013561350581613072565b9150606085013561351581613072565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f5f5f5f60808587031215613560575f5ffd5b843561356b81613072565b935060208501359250604085013561358281613072565b9396929550929360600135925050565b5f5f604083850312156135a3575f5ffd5b82356135ae81613072565b946020939093013593505050565b5f5f604083850312156135cd575f5ffd5b6132c983613489565b5f67ffffffffffffffff8211156135ef576135ef613171565b5060051b60200190565b5f82601f830112613608575f5ffd5b8151613616613278826135d6565b8082825260208201915060208360051b860101925085831115613637575f5ffd5b602085015b8381101561365457805183526020928301920161363c565b5095945050505050565b5f5f5f60608486031215613670575f5ffd5b835161367b81613072565b602085015190935067ffffffffffffffff811115613697575f5ffd5b6136a3868287016135f9565b925050604084015167ffffffffffffffff8111156136bf575f5ffd5b6136cb868287016135f9565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f60608486031215613714575f5ffd5b835161371f81613072565b602085015160409095015190969495509392505050565b5f60208284031215613746575f5ffd5b815167ffffffffffffffff81111561375c575f5ffd5b8201601f8101841361376c575f5ffd5b805161377a61327882613216565b81815285602083850101111561378e575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b8183526020830192505f815f5b848110156138025781356137cb8161301e565b7fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019591909101906001016137b8565b5093949350505050565b5f60c0820173ffffffffffffffffffffffffffffffffffffffff8a16835273ffffffffffffffffffffffffffffffffffffffff8916602084015260c060408401528087825260e08401905060e08860051b8501019150885f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18b3603015b8a8210156139c4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2087860301845282358181126138c6575f5ffd5b8c0180356138d381613072565b73ffffffffffffffffffffffffffffffffffffffff16865260208101356138f981613066565b6003811061392e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208701526040810135368290037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112613968575f5ffd5b0160208101903567ffffffffffffffff811115613983575f5ffd5b8060051b3603821315613994575f5ffd5b606060408801526139a96060880182846137ab565b9650505060208301925060208401935060018201915061388a565b5050505060608301959095525073ffffffffffffffffffffffffffffffffffffffff9283166080820152911660a090910152949350505050565b5f81518060208401855e5f93019283525090919050565b5f61241d613a2383866139fe565b846139fe565b5f60208284031215613a39575f5ffd5b5051919050565b5f60208284031215613a50575f5ffd5b815167ffffffffffffffff811115613a66575f5ffd5b820160608185031215613a77575f5ffd5b613a7f61319e565b8151613a8a81613072565b81526020820151613a9a81613066565b6020820152604082015167ffffffffffffffff811115613ab8575f5ffd5b80830192505084601f830112613acc575f5ffd5b8151613ada613278826135d6565b8082825260208201915060208360051b860101925087831115613afb575f5ffd5b6020850194505b82851015613b26578451613b158161301e565b825260209485019490910190613b02565b6040840152509095945050505050565b5f611e3482846139fe565b8181038181111561061e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220d78c6c24d11b97a3b96c198df916bf0d72f652660e88b07601354903f22b162b64736f6c634300081c0033",
}

// ForgeFactory is an auto generated Go binding around an Ethereum contract.
type ForgeFactory struct {
	abi abi.ABI
}

// NewForgeFactory creates a new instance of ForgeFactory.
func NewForgeFactory() *ForgeFactory {
	parsed, err := ForgeFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ForgeFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ForgeFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (forgeFactory *ForgeFactory) PackDEFAULTADMINROLE() []byte {
	enc, err := forgeFactory.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (forgeFactory *ForgeFactory) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := forgeFactory.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (forgeFactory *ForgeFactory) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := forgeFactory.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (forgeFactory *ForgeFactory) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := forgeFactory.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAddService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f765bfe.
//
// Solidity: function addService(address owner, address validator, bytes32 service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (forgeFactory *ForgeFactory) PackAddService(owner common.Address, validator common.Address, service [32]byte, addCuts []IDiamondCutFacetCut) []byte {
	enc, err := forgeFactory.abi.Pack("addService", owner, validator, service, addCuts)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAddService is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f765bfe.
//
// Solidity: function addService(address owner, address validator, bytes32 service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (forgeFactory *ForgeFactory) UnpackAddService(data []byte) (common.Address, error) {
	out, err := forgeFactory.abi.Unpack("addService", data)
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
func (forgeFactory *ForgeFactory) PackAlertBurn(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertBurn", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAlertMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe98a5784.
//
// Solidity: function alertMint(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertMint(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertMint", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAlertTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b7e40c6.
//
// Solidity: function alertTransfer(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertTransfer(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertTransfer", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAlertTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a15d070.
//
// Solidity: function alertTransferFrom(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertTransferFrom(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertTransferFrom", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe57997d2.
//
// Solidity: function allForges() view returns(bytes32[], address[])
func (forgeFactory *ForgeFactory) PackAllForges() []byte {
	enc, err := forgeFactory.abi.Pack("allForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// AllForgesOutput serves as a container for the return parameters of contract
// method AllForges.
type AllForgesOutput struct {
	Arg0 [][32]byte
	Arg1 []common.Address
}

// UnpackAllForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe57997d2.
//
// Solidity: function allForges() view returns(bytes32[], address[])
func (forgeFactory *ForgeFactory) UnpackAllForges(data []byte) (AllForgesOutput, error) {
	out, err := forgeFactory.abi.Unpack("allForges", data)
	outstruct := new(AllForgesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, err

}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(bytes32 service, address forge, bool running)
func (forgeFactory *ForgeFactory) PackForgeByIndex(index *big.Int) []byte {
	enc, err := forgeFactory.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// ForgeByIndexOutput serves as a container for the return parameters of contract
// method ForgeByIndex.
type ForgeByIndexOutput struct {
	Service [32]byte
	Forge   common.Address
	Running bool
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(bytes32 service, address forge, bool running)
func (forgeFactory *ForgeFactory) UnpackForgeByIndex(data []byte) (ForgeByIndexOutput, error) {
	out, err := forgeFactory.abi.Unpack("forgeByIndex", data)
	outstruct := new(ForgeByIndexOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Service = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Forge = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Running = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackForgeByService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec0b6a4d.
//
// Solidity: function forgeByService(bytes32 service) view returns(address forge, bool running)
func (forgeFactory *ForgeFactory) PackForgeByService(service [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("forgeByService", service)
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
// from invoking the contract method with ID 0xec0b6a4d.
//
// Solidity: function forgeByService(bytes32 service) view returns(address forge, bool running)
func (forgeFactory *ForgeFactory) UnpackForgeByService(data []byte) (ForgeByServiceOutput, error) {
	out, err := forgeFactory.abi.Unpack("forgeByService", data)
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
func (forgeFactory *ForgeFactory) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (forgeFactory *ForgeFactory) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := forgeFactory.abi.Unpack("getRoleAdmin", data)
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
func (forgeFactory *ForgeFactory) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (forgeFactory *ForgeFactory) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (forgeFactory *ForgeFactory) UnpackHasRole(data []byte) (bool, error) {
	out, err := forgeFactory.abi.Unpack("hasRole", data)
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
func (forgeFactory *ForgeFactory) PackInitialize(owner common.Address, forgeProxyCode common.Address, diamondImpl common.Address, baseImpl common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("initialize", owner, forgeProxyCode, diamondImpl, baseImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsRunningForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7525cba.
//
// Solidity: function isRunningForge(address forge) view returns(bytes32 service, bool running)
func (forgeFactory *ForgeFactory) PackIsRunningForge(forge common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("isRunningForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// IsRunningForgeOutput serves as a container for the return parameters of contract
// method IsRunningForge.
type IsRunningForgeOutput struct {
	Service [32]byte
	Running bool
}

// UnpackIsRunningForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc7525cba.
//
// Solidity: function isRunningForge(address forge) view returns(bytes32 service, bool running)
func (forgeFactory *ForgeFactory) UnpackIsRunningForge(data []byte) (IsRunningForgeOutput, error) {
	out, err := forgeFactory.abi.Unpack("isRunningForge", data)
	outstruct := new(IsRunningForgeOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Service = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Running = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, err

}

// PackLengthAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f44ec09.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (forgeFactory *ForgeFactory) PackLengthAllForges() []byte {
	enc, err := forgeFactory.abi.Pack("lengthAllForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackLengthAllForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f44ec09.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (forgeFactory *ForgeFactory) UnpackLengthAllForges(data []byte) (*big.Int, error) {
	out, err := forgeFactory.abi.Unpack("lengthAllForges", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPauseService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecf5bcc1.
//
// Solidity: function pauseService(bytes32 service, bool paused) returns()
func (forgeFactory *ForgeFactory) PackPauseService(service [32]byte, paused bool) []byte {
	enc, err := forgeFactory.abi.Pack("pauseService", service, paused)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (forgeFactory *ForgeFactory) PackProxiableUUID() []byte {
	enc, err := forgeFactory.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (forgeFactory *ForgeFactory) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := forgeFactory.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRemoveService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3884c3f.
//
// Solidity: function removeService(bytes32 service) returns()
func (forgeFactory *ForgeFactory) PackRemoveService(service [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("removeService", service)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (forgeFactory *ForgeFactory) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (forgeFactory *ForgeFactory) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeFactory *ForgeFactory) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := forgeFactory.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeFactory *ForgeFactory) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := forgeFactory.abi.Unpack("supportsInterface", data)
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
func (forgeFactory *ForgeFactory) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ForgeFactoryERC1155Burned represents a ERC1155Burned event raised by the ForgeFactory contract.
type ForgeFactoryERC1155Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC1155BurnedEventName = "ERC1155Burned"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC1155Burned) ContractEventName() string {
	return ForgeFactoryERC1155BurnedEventName
}

// UnpackERC1155BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 tokenID, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC1155BurnedEvent(log *types.Log) (*ForgeFactoryERC1155Burned, error) {
	event := "ERC1155Burned"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC1155Burned)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC1155Minted represents a ERC1155Minted event raised by the ForgeFactory contract.
type ForgeFactoryERC1155Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC1155MintedEventName = "ERC1155Minted"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC1155Minted) ContractEventName() string {
	return ForgeFactoryERC1155MintedEventName
}

// UnpackERC1155MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC1155MintedEvent(log *types.Log) (*ForgeFactoryERC1155Minted, error) {
	event := "ERC1155Minted"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC1155Minted)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC1155Transferred represents a ERC1155Transferred event raised by the ForgeFactory contract.
type ForgeFactoryERC1155Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC1155TransferredEventName = "ERC1155Transferred"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC1155Transferred) ContractEventName() string {
	return ForgeFactoryERC1155TransferredEventName
}

// UnpackERC1155TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC1155TransferredEvent(log *types.Log) (*ForgeFactoryERC1155Transferred, error) {
	event := "ERC1155Transferred"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC1155Transferred)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC20Burned represents a ERC20Burned event raised by the ForgeFactory contract.
type ForgeFactoryERC20Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC20BurnedEventName = "ERC20Burned"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC20Burned) ContractEventName() string {
	return ForgeFactoryERC20BurnedEventName
}

// UnpackERC20BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC20BurnedEvent(log *types.Log) (*ForgeFactoryERC20Burned, error) {
	event := "ERC20Burned"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC20Burned)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC20FeeCollected represents a ERC20FeeCollected event raised by the ForgeFactory contract.
type ForgeFactoryERC20FeeCollected struct {
	Uuid         *big.Int
	FeeRecipient common.Address
	Token        common.Address
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC20FeeCollectedEventName = "ERC20FeeCollected"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC20FeeCollected) ContractEventName() string {
	return ForgeFactoryERC20FeeCollectedEventName
}

// UnpackERC20FeeCollectedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20FeeCollected(uint256 indexed uuid, address indexed feeRecipient, address indexed token, uint256 fee)
func (forgeFactory *ForgeFactory) UnpackERC20FeeCollectedEvent(log *types.Log) (*ForgeFactoryERC20FeeCollected, error) {
	event := "ERC20FeeCollected"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC20FeeCollected)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC20Minted represents a ERC20Minted event raised by the ForgeFactory contract.
type ForgeFactoryERC20Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC20MintedEventName = "ERC20Minted"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC20Minted) ContractEventName() string {
	return ForgeFactoryERC20MintedEventName
}

// UnpackERC20MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC20MintedEvent(log *types.Log) (*ForgeFactoryERC20Minted, error) {
	event := "ERC20Minted"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC20Minted)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC20Transferred represents a ERC20Transferred event raised by the ForgeFactory contract.
type ForgeFactoryERC20Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC20TransferredEventName = "ERC20Transferred"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC20Transferred) ContractEventName() string {
	return ForgeFactoryERC20TransferredEventName
}

// UnpackERC20TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC20TransferredEvent(log *types.Log) (*ForgeFactoryERC20Transferred, error) {
	event := "ERC20Transferred"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC20Transferred)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC20TransferredFrom represents a ERC20TransferredFrom event raised by the ForgeFactory contract.
type ForgeFactoryERC20TransferredFrom struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC20TransferredFromEventName = "ERC20TransferredFrom"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC20TransferredFrom) ContractEventName() string {
	return ForgeFactoryERC20TransferredFromEventName
}

// UnpackERC20TransferredFromEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20TransferredFrom(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 amount)
func (forgeFactory *ForgeFactory) UnpackERC20TransferredFromEvent(log *types.Log) (*ForgeFactoryERC20TransferredFrom, error) {
	event := "ERC20TransferredFrom"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC20TransferredFrom)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC721Burned represents a ERC721Burned event raised by the ForgeFactory contract.
type ForgeFactoryERC721Burned struct {
	Service [32]byte
	Uuid    *big.Int
	From    common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC721BurnedEventName = "ERC721Burned"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC721Burned) ContractEventName() string {
	return ForgeFactoryERC721BurnedEventName
}

// UnpackERC721BurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Burned(bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 tokenID)
func (forgeFactory *ForgeFactory) UnpackERC721BurnedEvent(log *types.Log) (*ForgeFactoryERC721Burned, error) {
	event := "ERC721Burned"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC721Burned)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC721Minted represents a ERC721Minted event raised by the ForgeFactory contract.
type ForgeFactoryERC721Minted struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC721MintedEventName = "ERC721Minted"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC721Minted) ContractEventName() string {
	return ForgeFactoryERC721MintedEventName
}

// UnpackERC721MintedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID)
func (forgeFactory *ForgeFactory) UnpackERC721MintedEvent(log *types.Log) (*ForgeFactoryERC721Minted, error) {
	event := "ERC721Minted"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC721Minted)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryERC721Transferred represents a ERC721Transferred event raised by the ForgeFactory contract.
type ForgeFactoryERC721Transferred struct {
	Service [32]byte
	Uuid    *big.Int
	To      common.Address
	Token   common.Address
	TokenID *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryERC721TransferredEventName = "ERC721Transferred"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryERC721Transferred) ContractEventName() string {
	return ForgeFactoryERC721TransferredEventName
}

// UnpackERC721TransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721Transferred(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID)
func (forgeFactory *ForgeFactory) UnpackERC721TransferredEvent(log *types.Log) (*ForgeFactoryERC721Transferred, error) {
	event := "ERC721Transferred"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryERC721Transferred)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryInitialized represents a Initialized event raised by the ForgeFactory contract.
type ForgeFactoryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryInitialized) ContractEventName() string {
	return ForgeFactoryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (forgeFactory *ForgeFactory) UnpackInitializedEvent(log *types.Log) (*ForgeFactoryInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryInitialized)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the ForgeFactory contract.
type ForgeFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryRoleAdminChanged) ContractEventName() string {
	return ForgeFactoryRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (forgeFactory *ForgeFactory) UnpackRoleAdminChangedEvent(log *types.Log) (*ForgeFactoryRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryRoleGranted represents a RoleGranted event raised by the ForgeFactory contract.
type ForgeFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryRoleGranted) ContractEventName() string {
	return ForgeFactoryRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (forgeFactory *ForgeFactory) UnpackRoleGrantedEvent(log *types.Log) (*ForgeFactoryRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryRoleGranted)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryRoleRevoked represents a RoleRevoked event raised by the ForgeFactory contract.
type ForgeFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryRoleRevoked) ContractEventName() string {
	return ForgeFactoryRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (forgeFactory *ForgeFactory) UnpackRoleRevokedEvent(log *types.Log) (*ForgeFactoryRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryRoleRevoked)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryServicePaused represents a ServicePaused event raised by the ForgeFactory contract.
type ForgeFactoryServicePaused struct {
	Service [32]byte
	Paused  bool
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryServicePausedEventName = "ServicePaused"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryServicePaused) ContractEventName() string {
	return ForgeFactoryServicePausedEventName
}

// UnpackServicePausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServicePaused(bytes32 indexed service, bool paused)
func (forgeFactory *ForgeFactory) UnpackServicePausedEvent(log *types.Log) (*ForgeFactoryServicePaused, error) {
	event := "ServicePaused"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryServicePaused)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryServiceRegistered represents a ServiceRegistered event raised by the ForgeFactory contract.
type ForgeFactoryServiceRegistered struct {
	Service [32]byte
	Forge   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryServiceRegisteredEventName = "ServiceRegistered"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryServiceRegistered) ContractEventName() string {
	return ForgeFactoryServiceRegisteredEventName
}

// UnpackServiceRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServiceRegistered(bytes32 indexed service, address indexed forge)
func (forgeFactory *ForgeFactory) UnpackServiceRegisteredEvent(log *types.Log) (*ForgeFactoryServiceRegistered, error) {
	event := "ServiceRegistered"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryServiceRegistered)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryServiceUnregistered represents a ServiceUnregistered event raised by the ForgeFactory contract.
type ForgeFactoryServiceUnregistered struct {
	Service [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryServiceUnregisteredEventName = "ServiceUnregistered"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryServiceUnregistered) ContractEventName() string {
	return ForgeFactoryServiceUnregisteredEventName
}

// UnpackServiceUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ServiceUnregistered(bytes32 indexed service)
func (forgeFactory *ForgeFactory) UnpackServiceUnregisteredEvent(log *types.Log) (*ForgeFactoryServiceUnregistered, error) {
	event := "ServiceUnregistered"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryServiceUnregistered)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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

// ForgeFactoryUpgraded represents a Upgraded event raised by the ForgeFactory contract.
type ForgeFactoryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ForgeFactoryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ForgeFactoryUpgraded) ContractEventName() string {
	return ForgeFactoryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (forgeFactory *ForgeFactory) UnpackUpgradedEvent(log *types.Log) (*ForgeFactoryUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != forgeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeFactoryUpgraded)
	if len(log.Data) > 0 {
		if err := forgeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeFactory.abi.Events[event].Inputs {
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
func (forgeFactory *ForgeFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeFactoryAlreadyUsedService"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeFactoryAlreadyUsedServiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeFactoryCallerIsNotForge"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeFactoryCallerIsNotForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeFactoryCallerIsPausedForge"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeFactoryCallerIsPausedForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeFactoryInvalidData"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeFactoryInvalidDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeFactoryServiceNotFound"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeFactoryServiceNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["TokenForgeInvalidTokenType"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackTokenForgeInvalidTokenTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ForgeFactoryAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ForgeFactory contract.
type ForgeFactoryAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ForgeFactoryAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (forgeFactory *ForgeFactory) UnpackAccessControlBadConfirmationError(raw []byte) (*ForgeFactoryAccessControlBadConfirmation, error) {
	out := new(ForgeFactoryAccessControlBadConfirmation)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ForgeFactory contract.
type ForgeFactoryAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ForgeFactoryAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (forgeFactory *ForgeFactory) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ForgeFactoryAccessControlUnauthorizedAccount, error) {
	out := new(ForgeFactoryAccessControlUnauthorizedAccount)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryAddressEmptyCode represents a AddressEmptyCode error raised by the ForgeFactory contract.
type ForgeFactoryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ForgeFactoryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (forgeFactory *ForgeFactory) UnpackAddressEmptyCodeError(raw []byte) (*ForgeFactoryAddressEmptyCode, error) {
	out := new(ForgeFactoryAddressEmptyCode)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the ForgeFactory contract.
type ForgeFactoryCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func ForgeFactoryCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (forgeFactory *ForgeFactory) UnpackCreate2EmptyBytecodeError(raw []byte) (*ForgeFactoryCreate2EmptyBytecode, error) {
	out := new(ForgeFactoryCreate2EmptyBytecode)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ForgeFactory contract.
type ForgeFactoryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ForgeFactoryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (forgeFactory *ForgeFactory) UnpackERC1967InvalidImplementationError(raw []byte) (*ForgeFactoryERC1967InvalidImplementation, error) {
	out := new(ForgeFactoryERC1967InvalidImplementation)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryERC1967NonPayable represents a ERC1967NonPayable error raised by the ForgeFactory contract.
type ForgeFactoryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ForgeFactoryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (forgeFactory *ForgeFactory) UnpackERC1967NonPayableError(raw []byte) (*ForgeFactoryERC1967NonPayable, error) {
	out := new(ForgeFactoryERC1967NonPayable)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryFailedCall represents a FailedCall error raised by the ForgeFactory contract.
type ForgeFactoryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ForgeFactoryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (forgeFactory *ForgeFactory) UnpackFailedCallError(raw []byte) (*ForgeFactoryFailedCall, error) {
	out := new(ForgeFactoryFailedCall)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryFailedDeployment represents a FailedDeployment error raised by the ForgeFactory contract.
type ForgeFactoryFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func ForgeFactoryFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (forgeFactory *ForgeFactory) UnpackFailedDeploymentError(raw []byte) (*ForgeFactoryFailedDeployment, error) {
	out := new(ForgeFactoryFailedDeployment)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryInsufficientBalance represents a InsufficientBalance error raised by the ForgeFactory contract.
type ForgeFactoryInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func ForgeFactoryInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (forgeFactory *ForgeFactory) UnpackInsufficientBalanceError(raw []byte) (*ForgeFactoryInsufficientBalance, error) {
	out := new(ForgeFactoryInsufficientBalance)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryInvalidInitialization represents a InvalidInitialization error raised by the ForgeFactory contract.
type ForgeFactoryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ForgeFactoryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (forgeFactory *ForgeFactory) UnpackInvalidInitializationError(raw []byte) (*ForgeFactoryInvalidInitialization, error) {
	out := new(ForgeFactoryInvalidInitialization)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryNotInitializing represents a NotInitializing error raised by the ForgeFactory contract.
type ForgeFactoryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ForgeFactoryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (forgeFactory *ForgeFactory) UnpackNotInitializingError(raw []byte) (*ForgeFactoryNotInitializing, error) {
	out := new(ForgeFactoryNotInitializing)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeFactoryAlreadyUsedService represents a TokenForgeFactory__AlreadyUsedService error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeFactoryAlreadyUsedService struct {
	Service [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__AlreadyUsedService(bytes32 service)
func ForgeFactoryTokenForgeFactoryAlreadyUsedServiceErrorID() common.Hash {
	return common.HexToHash("0x5502583a4e5df3d611d0359980fb46a75c8ce25278a66f5253b4cd467c300887")
}

// UnpackTokenForgeFactoryAlreadyUsedServiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__AlreadyUsedService(bytes32 service)
func (forgeFactory *ForgeFactory) UnpackTokenForgeFactoryAlreadyUsedServiceError(raw []byte) (*ForgeFactoryTokenForgeFactoryAlreadyUsedService, error) {
	out := new(ForgeFactoryTokenForgeFactoryAlreadyUsedService)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryAlreadyUsedService", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeFactoryCallerIsNotForge represents a TokenForgeFactory__CallerIsNotForge error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeFactoryCallerIsNotForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__CallerIsNotForge(address caller)
func ForgeFactoryTokenForgeFactoryCallerIsNotForgeErrorID() common.Hash {
	return common.HexToHash("0x4b73e4e6b8076eaaeddcbe95bc8fdf6efa1a2606fc078882372b3a4389ba7e82")
}

// UnpackTokenForgeFactoryCallerIsNotForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__CallerIsNotForge(address caller)
func (forgeFactory *ForgeFactory) UnpackTokenForgeFactoryCallerIsNotForgeError(raw []byte) (*ForgeFactoryTokenForgeFactoryCallerIsNotForge, error) {
	out := new(ForgeFactoryTokenForgeFactoryCallerIsNotForge)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryCallerIsNotForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeFactoryCallerIsPausedForge represents a TokenForgeFactory__CallerIsPausedForge error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeFactoryCallerIsPausedForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__CallerIsPausedForge(address caller)
func ForgeFactoryTokenForgeFactoryCallerIsPausedForgeErrorID() common.Hash {
	return common.HexToHash("0xae111e923268b59d275e8c9ba5e0305328e43fc8b38f3601919e1045c4eacf8c")
}

// UnpackTokenForgeFactoryCallerIsPausedForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__CallerIsPausedForge(address caller)
func (forgeFactory *ForgeFactory) UnpackTokenForgeFactoryCallerIsPausedForgeError(raw []byte) (*ForgeFactoryTokenForgeFactoryCallerIsPausedForge, error) {
	out := new(ForgeFactoryTokenForgeFactoryCallerIsPausedForge)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryCallerIsPausedForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeFactoryInvalidData represents a TokenForgeFactory__InvalidData error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeFactoryInvalidData struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__InvalidData(bytes32 field)
func ForgeFactoryTokenForgeFactoryInvalidDataErrorID() common.Hash {
	return common.HexToHash("0x435db3481d8eb115a5d9fe457169c1d3dca88f1c1a6c4e6ee3b67d659593d3ca")
}

// UnpackTokenForgeFactoryInvalidDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__InvalidData(bytes32 field)
func (forgeFactory *ForgeFactory) UnpackTokenForgeFactoryInvalidDataError(raw []byte) (*ForgeFactoryTokenForgeFactoryInvalidData, error) {
	out := new(ForgeFactoryTokenForgeFactoryInvalidData)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryInvalidData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeFactoryServiceNotFound represents a TokenForgeFactory__ServiceNotFound error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeFactoryServiceNotFound struct {
	Service [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForgeFactory__ServiceNotFound(bytes32 service)
func ForgeFactoryTokenForgeFactoryServiceNotFoundErrorID() common.Hash {
	return common.HexToHash("0x2a188b7fac945d1d82f2b68ea1d938569eb745b8636b1c8a173274cce308b9cc")
}

// UnpackTokenForgeFactoryServiceNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForgeFactory__ServiceNotFound(bytes32 service)
func (forgeFactory *ForgeFactory) UnpackTokenForgeFactoryServiceNotFoundError(raw []byte) (*ForgeFactoryTokenForgeFactoryServiceNotFound, error) {
	out := new(ForgeFactoryTokenForgeFactoryServiceNotFound)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeFactoryServiceNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryTokenForgeInvalidTokenType represents a TokenForge__InvalidTokenType error raised by the ForgeFactory contract.
type ForgeFactoryTokenForgeInvalidTokenType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenForge__InvalidTokenType()
func ForgeFactoryTokenForgeInvalidTokenTypeErrorID() common.Hash {
	return common.HexToHash("0x9570ddf58a1418a686536d6c200028d56382c32ca6637e8defd204d5722e94e9")
}

// UnpackTokenForgeInvalidTokenTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenForge__InvalidTokenType()
func (forgeFactory *ForgeFactory) UnpackTokenForgeInvalidTokenTypeError(raw []byte) (*ForgeFactoryTokenForgeInvalidTokenType, error) {
	out := new(ForgeFactoryTokenForgeInvalidTokenType)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "TokenForgeInvalidTokenType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ForgeFactory contract.
type ForgeFactoryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ForgeFactoryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (forgeFactory *ForgeFactory) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ForgeFactoryUUPSUnauthorizedCallContext, error) {
	out := new(ForgeFactoryUUPSUnauthorizedCallContext)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeFactoryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ForgeFactory contract.
type ForgeFactoryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ForgeFactoryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (forgeFactory *ForgeFactory) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ForgeFactoryUUPSUnsupportedProxiableUUID, error) {
	out := new(ForgeFactoryUUPSUnsupportedProxiableUUID)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
