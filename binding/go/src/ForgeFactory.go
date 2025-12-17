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

// ForgeFactoryMetaData contains all meta data concerning the ForgeFactory contract.
var ForgeFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"addCuts\",\"type\":\"tuple[]\"}],\"name\":\"addService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allForges\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"forgeByService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forgeProxyCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_diamondImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isRunningForge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lengthAllForges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"pauseService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ERC20FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20TransferredFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"ServicePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"ServiceUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__AlreadyUsedService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsNotForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsPausedForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__InvalidData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__ServiceNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenForge__InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ForgeFactory",
	Bin: "0x60a080604052346100c257306080525f5160206138b95f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b6040516137f290816100c78239608051818181611b9e0152611c7f0152f35b6001600160401b0319166001600160401b039081175f5160206138b95f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a7146125d7575080630b7e40c614612415578063248a9ca3146123a55780632f2ff15d1461232a5780632f44ec09146122d057806336568abe146122485780633a15d0701461215c5780633ea113d014611f9a5780634f1ef28614611c1657806352d1902d14611b595780637f765bfe146114b457806391d14854146114205780639ca92df91461138a578063a217fddf14611352578063ad3cb1cc146112a0578063c7525cba146111db578063d3884c3f14611059578063d547741f14610fd9578063e57997d214610e0a578063e98a578414610af1578063ec0b6a4d14610a52578063ecf5bcc1146108915763f8c8765e14610119575f80fd5b3461088d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d576101506126b6565b610158612693565b906044359073ffffffffffffffffffffffffffffffffffffffff821680920361088d5760643573ffffffffffffffffffffffffffffffffffffffff811680910361088d577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460ff8160401c16159467ffffffffffffffff821680159081610885575b600114908161087b575b159081610872575b5061084a57818660017fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000073ffffffffffffffffffffffffffffffffffffffff9516177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00556107f5575b5061025f6132dc565b6102676132dc565b16926102716132dc565b73ffffffffffffffffffffffffffffffffffffffff8316156107a957831561075d5780156107115781156106c5576040517f24c12bf60000000000000000000000000000000000000000000000000000000081525f81600481885afa9081156105e6575f916106a3575b505115610657576040517f1bd8f7f10000000000000000000000000000000000000000000000000000000081525f81600481855afa80156105e65773ffffffffffffffffffffffffffffffffffffffff915f9161063d575b50511681036105f1576040517f1bd8f7f10000000000000000000000000000000000000000000000000000000081525f81600481865afa80156105e65773ffffffffffffffffffffffffffffffffffffffff915f916105c4575b5051168203610578576104de937fffffffffffffffffffffffff00000000000000000000000000000000000000007f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005416177f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00557fffffffffffffffffffffffff00000000000000000000000000000000000000007f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e015416177f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e01557fffffffffffffffffffffffff00000000000000000000000000000000000000007f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e025416177f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e02556104d881612cd5565b50612dbc565b506104e557005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f435db348000000000000000000000000000000000000000000000000000000005f527f62617365496d706c20666163657400000000000000000000000000000000000060045260245ffd5b6105e091503d805f833e6105d8818361279a565b810190613151565b5f61038d565b6040513d5f823e3d90fd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f6469616d6f6e64496d706c20666163657400000000000000000000000000000060045260245ffd5b61065191503d805f833e6105d8818361279a565b5f610333565b7f435db348000000000000000000000000000000000000000000000000000000005f527f666f72676550726f7879436f646520636f64650000000000000000000000000060045260245ffd5b6106bf91503d805f833e6106b7818361279a565b810190612a8d565b5f6102db565b7f435db348000000000000000000000000000000000000000000000000000000005f527f62617365496d706c00000000000000000000000000000000000000000000000060045260245ffd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f6469616d6f6e64496d706c00000000000000000000000000000000000000000060045260245ffd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f666f72676550726f7879436f646500000000000000000000000000000000000060045260245ffd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f6f776e657200000000000000000000000000000000000000000000000000000060045260245ffd5b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00555f610256565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b9050155f6101ed565b303b1591506101e5565b8791506101db565b5f80fd5b3461088d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d576004356024358015159081810361088d576108d9612bc7565b61090d835f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f2054151590565b15610a2657825f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e076020528160ff60405f20541615151461094b57005b7eb31832845d6567ecb46caeeee58acfcfddfcbba821c68dcbae9055ea4d795591602091156109ce57835f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e07825260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690555b604051908152a2005b835f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e07825260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556109c5565b827f2a188b7f000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760406004355f610a908261326d565b9290610abc575b5073ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152f35b90505f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0760205260ff825f20541683610a97565b3461088d57610aff366126fa565b93929190610b0c33612af0565b93610b1681612888565b80610beb575082610b60827f9612604afba70e4cf03261d7d86ca03d08911d9887aa41621dea929e34cfd7b19773ffffffffffffffffffffffffffffffffffffffff940190612a63565b6040805173ffffffffffffffffffffffffffffffffffffffff89168152602081019490945291999098939095169491a482610b9d575b505050505b005b602073ffffffffffffffffffffffffffffffffffffffff807fa2b448c8495be4baf0f3195bf0d82d5ec29d24aa3c5dd949969149bd797fb3159360405196875216951693a480808080610b96565b919492939291610bfa81612888565b60018103610c805750610c7b610c49827f556d832108ef7e0123726b3e037bc1cd39023e8d397b242d8b36426550d8c1239473ffffffffffffffffffffffffffffffffffffffff940190612a46565b6040805173ffffffffffffffffffffffffffffffffffffffff9099168952602089019190915292169591829190820190565b0390a4005b600290610c8c81612888565b03610de257610c9d918101906128bf565b9015610d5d5791610ccd8360208073ffffffffffffffffffffffffffffffffffffffff95965183010191016129b1565b81519592909416935f5b868110610ce057005b808686867f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f61610d1160019689612a05565b518d610d54610d20888b612a05565b516040519384938460409194939273ffffffffffffffffffffffffffffffffffffffff606083019616825260208201520152565b0390a401610cd7565b73ffffffffffffffffffffffffffffffffffffffff610c7b610da9836020807f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f619651830101910161291a565b9060409392935194859416988460409194939273ffffffffffffffffffffffffffffffffffffffff606083019616825260208201520152565b7f9570ddf5000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461088d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d577f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0354610e638161293c565b610e70604051918261279a565b818152610e7c8261293c565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437610eb08161293c565b92610ebe604051948561279a565b818452610eca8261293c565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020860193013684375f5b818110610f8f575050604051938493604085019060408652518091526060850192905f5b818110610f76575050506020908483038286015251918281520191905f5b818110610f47575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610f39565b8251855287965060209485019490920191600101610f1b565b80610fa060019297949596976130da565b73ffffffffffffffffffffffffffffffffffffffff610fbf8488612a05565b91169052610fcd8289612a05565b52019493929194610ef7565b3461088d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57610b9b600435611016612693565b9061105461104f825f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b612c4f565b612fd2565b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57600435611093612bc7565b61109c8161326d565b90156111af5761113773ffffffffffffffffffffffffffffffffffffffff5f928484527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e056020528360408120556110f285613525565b501673ffffffffffffffffffffffffffffffffffffffff165f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0660205260405f2090565b55805f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0760205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f475a09e552c07ad6aa3d8b6f3fb547091aae08290004e78277eb2ffb854ea21e5f80a2005b507f2a188b7f000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760406112146126b6565b73ffffffffffffffffffffffffffffffffffffffff165f9081527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e066020526040812054908161126c575b825191825215156020820152f35b50805f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0760205260ff825f20541661125e565b3461088d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760408051906112dd818361279a565b6005825260208201917f352e302e3000000000000000000000000000000000000000000000000000000083527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8351948593602085525180918160208701528686015e5f85828601015201168101030190f35b3461088d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760206040515f8152f35b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760606113c66004356130da565b815f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0760205273ffffffffffffffffffffffffffffffffffffffff60ff60405f2054169160405193845216602083015215156040820152f35b3461088d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57611457612693565b6004355f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b3461088d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d576114eb6126b6565b6114f3612693565b604435916064359167ffffffffffffffff831161088d573660238401121561088d5782600401359267ffffffffffffffff841161088d578360051b92366024858401011161088d5773ffffffffffffffffffffffffffffffffffffffff90611559612bc7565b169182156107a95773ffffffffffffffffffffffffffffffffffffffff16938415611b0d578515611ac1576115b8865f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f2054151590565b611a95576004935f73ffffffffffffffffffffffffffffffffffffffff7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005416604051968780927f24c12bf60000000000000000000000000000000000000000000000000000000082525afa9485156105e6575f95611a79575b5073ffffffffffffffffffffffffffffffffffffffff7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0193929354169073ffffffffffffffffffffffffffffffffffffffff7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e025416928460405196602088019960e08901918b52604089015260c0606089015252610100808701928701019460248201925f927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7d813603015b83851061192a578c8c8c6117a48d61174f818f8f8f89608085015260a084015260c0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261279a565b60206040519485928280850196805191829101885e840190838201905f8252519283915e01015f8152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361279a565b81511561190257829151905ff5903d15198215166105e65773ffffffffffffffffffffffffffffffffffffffff821680156118da5781611859602094825f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0586528360405f2055611815836133e1565b5073ffffffffffffffffffffffffffffffffffffffff165f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0660205260405f2090565b55815f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e07835260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905580604051927f0453e3ffce9b0d4f994d54652037d3e15d87ec29aab0aafba90ac322edc87c255f80a38152f35b7fb06ebf3d000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4ca249dc000000000000000000000000000000000000000000000000000000005f5260045ffd5b9091929394977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008a820301845288358281121561088d578301606082019073ffffffffffffffffffffffffffffffffffffffff611989602483016126d9565b1683526044810135600381101561088d576119a381612888565b60208401526064810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd368290030182121561088d576024910101906020823592019267ffffffffffffffff831161088d578260051b3603841361088d578260809260606040840152520191905f905b808210611a35575050506020806001929a019401950193929491906116fe565b9091928335907fffffffff00000000000000000000000000000000000000000000000000000000821680920361088d57602081600193829352019401920190611a15565b611a8e9195503d805f833e6106b7818361279a565b9387611632565b857f5502583a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f736572766963650000000000000000000000000000000000000000000000000060045260245ffd5b7f435db348000000000000000000000000000000000000000000000000000000005f527f76616c696461746f72000000000000000000000000000000000000000000000060045260245ffd5b3461088d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003611bee5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57611c486126b6565b60243567ffffffffffffffff811161088d57611c68903690600401612842565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115611f58575b50611bee57335f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff1615611f285773ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181611ef4575b50611d6557837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc859203611ec95750813b15611e9e57807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115611e6d575f80836020610b9b95519101845af43d15611e65573d91611e4983612808565b92611e57604051948561279a565b83523d5f602085013e613333565b606091613333565b505034611e7657005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011611f20575b81611f106020938361279a565b8101031261088d57519085611d34565b3d9150611f03565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004525f60245260445ffd5b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583611caa565b3461088d57611fa8366126fa565b93929190611fb533612af0565b93611fbf81612888565b80612009575082610b60827fe5e0650048e213b8a668474e72c774ec69d8a6830270b7cc1c339fcd6a3b2e419773ffffffffffffffffffffffffffffffffffffffff940190612a63565b91949293929161201881612888565b600181036120675750610c7b610c49827f28fcdc3e3355f548f5fe83373acf91a82e1e561ee22797f6e91f831d2b2b3e139473ffffffffffffffffffffffffffffffffffffffff940190612a46565b60029061207381612888565b03610de257612084918101906128bf565b901561211057916120b48360208073ffffffffffffffffffffffffffffffffffffffff95965183010191016129b1565b81519592909416935f5b8681106120c757005b808686867f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b6120f860019689612a05565b518d612107610d20888b612a05565b0390a4016120be565b73ffffffffffffffffffffffffffffffffffffffff610c7b610da9836020807f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b9651830101910161291a565b3461088d5761216a366126fa565b612178949392919433612af0565b9361218281612888565b610de257826121ca8673ffffffffffffffffffffffffffffffffffffffff937f14bed88907ac31fe1ba4cf8981f80fd25e29877a08d1568c7210adc31ee9bc95980190612a63565b6040805173ffffffffffffffffffffffffffffffffffffffff89168152602081019490945291999098939095169491a48261220157005b602073ffffffffffffffffffffffffffffffffffffffff807fa2b448c8495be4baf0f3195bf0d82d5ec29d24aa3c5dd949969149bd797fb3159360405196875216951693a4005b3461088d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5761227f612693565b3373ffffffffffffffffffffffffffffffffffffffff8216036122a857610b9b90600435612fd2565b7f6697b232000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461088d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d5760207f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0354604051908152f35b3461088d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57610b9b600435612367612693565b906123a061104f825f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b612ec0565b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57602061240d6004355f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b604051908152f35b3461088d57612423366126fa565b9392919061243033612af0565b9361243a81612888565b80612484575082610b60827f86d733c2654f9db1a830cfbbf2180ee3329101d853428a7ead4e85ddb3230d609773ffffffffffffffffffffffffffffffffffffffff940190612a63565b91949293929161249381612888565b600181036124e25750610c7b610c49827fa36201ae386aaacb89bc573b7c0ecc0913a93028f710f28ca29840454f4468869473ffffffffffffffffffffffffffffffffffffffff940190612a46565b6002906124ee81612888565b03610de2576124ff918101906128bf565b901561258b579161252f8360208073ffffffffffffffffffffffffffffffffffffffff95965183010191016129b1565b81519592909416935f5b86811061254257005b808686867fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c361257360019689612a05565b518d612582610d20888b612a05565b0390a401612539565b73ffffffffffffffffffffffffffffffffffffffff610c7b610da9836020807fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c39651830101910161291a565b3461088d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261088d57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361088d57817f7965db0b0000000000000000000000000000000000000000000000000000000060209314908115612669575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612662565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361088d57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361088d57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361088d57565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261088d57600435600381101561088d57916024359160443573ffffffffffffffffffffffffffffffffffffffff8116810361088d579160643567ffffffffffffffff811161088d578260238201121561088d5780600401359267ffffffffffffffff841161088d576024848301011161088d576024019190565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176127db57604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b67ffffffffffffffff81116127db57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561088d5780359061285982612808565b92612867604051948561279a565b8284526020838301011161088d57815f926020809301838601378301015290565b6003111561289257565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91909160408184031261088d578035801515810361088d5792602082013567ffffffffffffffff811161088d576128f69201612842565b90565b519073ffffffffffffffffffffffffffffffffffffffff8216820361088d57565b9081606091031261088d5761292e816128f9565b916040602083015192015190565b67ffffffffffffffff81116127db5760051b60200190565b9080601f8301121561088d57815161296b8161293c565b92612979604051948561279a565b81845260208085019260051b82010192831161088d57602001905b8282106129a15750505090565b8151815260209182019101612994565b9160608383031261088d576129c5836128f9565b92602081015167ffffffffffffffff811161088d57836129e6918301612954565b92604082015167ffffffffffffffff811161088d576128f69201612954565b8051821015612a195760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b919082604091031261088d576020612a5d836126d9565b92013590565b919082608091031261088d57612a78826126d9565b916020810135916060612a5d604084016126d9565b60208183031261088d5780519067ffffffffffffffff821161088d570181601f8201121561088d57805190612ac182612808565b92612acf604051948561279a565b8284526020838301011161088d57815f9260208093018386015e8301015290565b73ffffffffffffffffffffffffffffffffffffffff16805f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0660205260405f2054908115612b9c57815f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0760205260ff60405f20541615612b71575090565b7fae111e92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f4b73e4e6000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b335f9081527f511d0314ee32589a99ab73eccd900c257efd9224711c0e5aadda466017c07f6c602052604090205460ff1615612bff57565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c60245260445ffd5b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f20541615612ca65750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b73ffffffffffffffffffffffffffffffffffffffff81165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff16612db75773ffffffffffffffffffffffffffffffffffffffff165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505f90565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f511d0314ee32589a99ab73eccd900c257efd9224711c0e5aadda466017c07f6c602052604090205460ff16612db75773ffffffffffffffffffffffffffffffffffffffff165f8181527f511d0314ee32589a99ab73eccd900c257efd9224711c0e5aadda466017c07f6c6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790553391907faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f205416155f14612fcc57805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f2054165f14612fcc57805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b613104907f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036133cc565b90549060031b1c90815f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0560205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b60208183031261088d5780519067ffffffffffffffff821161088d57019060608282031261088d57604051916060830183811067ffffffffffffffff8211176127db57604052805173ffffffffffffffffffffffffffffffffffffffff8116810361088d5783526020810151600381101561088d57602084015260408101519067ffffffffffffffff821161088d57019080601f8301121561088d578151906131f98261293c565b92613207604051948561279a565b82845260208085019360051b82010191821161088d57602001915b81831061323457505050604082015290565b82517fffffffff000000000000000000000000000000000000000000000000000000008116810361088d57815260209283019201613222565b805f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0560205260405f205480155f146132d457505f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f20541515905f90565b600192909150565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561330b57565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b90613370575080511561334857805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b815115806133c3575b613381575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15613379565b8054821015612a19575f5260205f2001905f90565b805f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f2054155f14612db7577f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0354680100000000000000008110156127db576134d061349b8260018594017f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036133cc565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0354905f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f2055600190565b5f8181527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260409020548015612fcc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161378f577f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161378f578181036136fa575b5050507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e035480156136cd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161364e817f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036133cc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03555f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e046020525f6040812055600190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b61375a61372a61349b937f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036133cc565b90549060031b1c9283927f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036133cc565b90555f527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0460205260405f20555f80806135d7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212200a5c8274b283697038a9cd64dbd268813cdae63f0184378a25718ff525ecc75664736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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
// the contract method with ID 0xa217fddf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (forgeFactory *ForgeFactory) PackDEFAULTADMINROLE() []byte {
	enc, err := forgeFactory.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (forgeFactory *ForgeFactory) TryPackDEFAULTADMINROLE() ([]byte, error) {
	return forgeFactory.abi.Pack("DEFAULT_ADMIN_ROLE")
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
	return out0, nil
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (forgeFactory *ForgeFactory) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := forgeFactory.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (forgeFactory *ForgeFactory) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return forgeFactory.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
	return out0, nil
}

// PackAddService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f765bfe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addService(address owner, address validator, bytes32 service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (forgeFactory *ForgeFactory) PackAddService(owner common.Address, validator common.Address, service [32]byte, addCuts []IDiamondCutFacetCut) []byte {
	enc, err := forgeFactory.abi.Pack("addService", owner, validator, service, addCuts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f765bfe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addService(address owner, address validator, bytes32 service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (forgeFactory *ForgeFactory) TryPackAddService(owner common.Address, validator common.Address, service [32]byte, addCuts []IDiamondCutFacetCut) ([]byte, error) {
	return forgeFactory.abi.Pack("addService", owner, validator, service, addCuts)
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
	return out0, nil
}

// PackAlertBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ea113d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function alertBurn(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertBurn(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertBurn", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlertBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ea113d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function alertBurn(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) TryPackAlertBurn(tokenType uint8, uuid *big.Int, token common.Address, data []byte) ([]byte, error) {
	return forgeFactory.abi.Pack("alertBurn", tokenType, uuid, token, data)
}

// PackAlertMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe98a5784.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function alertMint(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertMint(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertMint", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlertMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe98a5784.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function alertMint(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) TryPackAlertMint(tokenType uint8, uuid *big.Int, token common.Address, data []byte) ([]byte, error) {
	return forgeFactory.abi.Pack("alertMint", tokenType, uuid, token, data)
}

// PackAlertTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b7e40c6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function alertTransfer(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertTransfer(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertTransfer", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlertTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b7e40c6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function alertTransfer(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) TryPackAlertTransfer(tokenType uint8, uuid *big.Int, token common.Address, data []byte) ([]byte, error) {
	return forgeFactory.abi.Pack("alertTransfer", tokenType, uuid, token, data)
}

// PackAlertTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a15d070.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function alertTransferFrom(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) PackAlertTransferFrom(tokenType uint8, uuid *big.Int, token common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("alertTransferFrom", tokenType, uuid, token, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlertTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a15d070.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function alertTransferFrom(uint8 tokenType, uint256 uuid, address token, bytes data) returns()
func (forgeFactory *ForgeFactory) TryPackAlertTransferFrom(tokenType uint8, uuid *big.Int, token common.Address, data []byte) ([]byte, error) {
	return forgeFactory.abi.Pack("alertTransferFrom", tokenType, uuid, token, data)
}

// PackAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe57997d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allForges() view returns(bytes32[], address[])
func (forgeFactory *ForgeFactory) PackAllForges() []byte {
	enc, err := forgeFactory.abi.Pack("allForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe57997d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allForges() view returns(bytes32[], address[])
func (forgeFactory *ForgeFactory) TryPackAllForges() ([]byte, error) {
	return forgeFactory.abi.Pack("allForges")
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
	return *outstruct, nil
}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function forgeByIndex(uint256 index) view returns(bytes32 service, address forge, bool running)
func (forgeFactory *ForgeFactory) PackForgeByIndex(index *big.Int) []byte {
	enc, err := forgeFactory.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function forgeByIndex(uint256 index) view returns(bytes32 service, address forge, bool running)
func (forgeFactory *ForgeFactory) TryPackForgeByIndex(index *big.Int) ([]byte, error) {
	return forgeFactory.abi.Pack("forgeByIndex", index)
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
	return *outstruct, nil
}

// PackForgeByService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec0b6a4d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function forgeByService(bytes32 service) view returns(address forge, bool running)
func (forgeFactory *ForgeFactory) PackForgeByService(service [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("forgeByService", service)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackForgeByService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec0b6a4d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function forgeByService(bytes32 service) view returns(address forge, bool running)
func (forgeFactory *ForgeFactory) TryPackForgeByService(service [32]byte) ([]byte, error) {
	return forgeFactory.abi.Pack("forgeByService", service)
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
	return *outstruct, nil
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (forgeFactory *ForgeFactory) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (forgeFactory *ForgeFactory) TryPackGetRoleAdmin(role [32]byte) ([]byte, error) {
	return forgeFactory.abi.Pack("getRoleAdmin", role)
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
	return out0, nil
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (forgeFactory *ForgeFactory) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (forgeFactory *ForgeFactory) TryPackGrantRole(role [32]byte, account common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("grantRole", role, account)
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (forgeFactory *ForgeFactory) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (forgeFactory *ForgeFactory) TryPackHasRole(role [32]byte, account common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("hasRole", role, account)
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
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c8765e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _owner, address _forgeProxyCode, address _diamondImpl, address _baseImpl) returns()
func (forgeFactory *ForgeFactory) PackInitialize(owner common.Address, forgeProxyCode common.Address, diamondImpl common.Address, baseImpl common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("initialize", owner, forgeProxyCode, diamondImpl, baseImpl)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c8765e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _owner, address _forgeProxyCode, address _diamondImpl, address _baseImpl) returns()
func (forgeFactory *ForgeFactory) TryPackInitialize(owner common.Address, forgeProxyCode common.Address, diamondImpl common.Address, baseImpl common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("initialize", owner, forgeProxyCode, diamondImpl, baseImpl)
}

// PackIsRunningForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7525cba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isRunningForge(address forge) view returns(bytes32 service, bool running)
func (forgeFactory *ForgeFactory) PackIsRunningForge(forge common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("isRunningForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsRunningForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7525cba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isRunningForge(address forge) view returns(bytes32 service, bool running)
func (forgeFactory *ForgeFactory) TryPackIsRunningForge(forge common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("isRunningForge", forge)
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
	return *outstruct, nil
}

// PackLengthAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f44ec09.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (forgeFactory *ForgeFactory) PackLengthAllForges() []byte {
	enc, err := forgeFactory.abi.Pack("lengthAllForges")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLengthAllForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f44ec09.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lengthAllForges() view returns(uint256)
func (forgeFactory *ForgeFactory) TryPackLengthAllForges() ([]byte, error) {
	return forgeFactory.abi.Pack("lengthAllForges")
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
	return out0, nil
}

// PackPauseService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecf5bcc1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauseService(bytes32 service, bool paused) returns()
func (forgeFactory *ForgeFactory) PackPauseService(service [32]byte, paused bool) []byte {
	enc, err := forgeFactory.abi.Pack("pauseService", service, paused)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauseService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecf5bcc1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauseService(bytes32 service, bool paused) returns()
func (forgeFactory *ForgeFactory) TryPackPauseService(service [32]byte, paused bool) ([]byte, error) {
	return forgeFactory.abi.Pack("pauseService", service, paused)
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (forgeFactory *ForgeFactory) PackProxiableUUID() []byte {
	enc, err := forgeFactory.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (forgeFactory *ForgeFactory) TryPackProxiableUUID() ([]byte, error) {
	return forgeFactory.abi.Pack("proxiableUUID")
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
	return out0, nil
}

// PackRemoveService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3884c3f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeService(bytes32 service) returns()
func (forgeFactory *ForgeFactory) PackRemoveService(service [32]byte) []byte {
	enc, err := forgeFactory.abi.Pack("removeService", service)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveService is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3884c3f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeService(bytes32 service) returns()
func (forgeFactory *ForgeFactory) TryPackRemoveService(service [32]byte) ([]byte, error) {
	return forgeFactory.abi.Pack("removeService", service)
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (forgeFactory *ForgeFactory) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (forgeFactory *ForgeFactory) TryPackRenounceRole(role [32]byte, callerConfirmation common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("renounceRole", role, callerConfirmation)
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (forgeFactory *ForgeFactory) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := forgeFactory.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (forgeFactory *ForgeFactory) TryPackRevokeRole(role [32]byte, account common.Address) ([]byte, error) {
	return forgeFactory.abi.Pack("revokeRole", role, account)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeFactory *ForgeFactory) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := forgeFactory.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeFactory *ForgeFactory) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return forgeFactory.abi.Pack("supportsInterface", interfaceId)
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
	return out0, nil
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (forgeFactory *ForgeFactory) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := forgeFactory.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (forgeFactory *ForgeFactory) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return forgeFactory.abi.Pack("upgradeToAndCall", newImplementation, data)
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
	if len(log.Topics) == 0 || log.Topics[0] != forgeFactory.abi.Events[event].ID {
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
