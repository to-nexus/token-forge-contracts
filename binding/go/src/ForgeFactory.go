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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"addCuts\",\"type\":\"tuple[]\"}],\"name\":\"addService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"alertTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allForges\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"}],\"name\":\"forgeByService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forgeProxyCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_diamondImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isRunningForge\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lengthAllForges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"pauseService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20TransferredFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"ERC721Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"ServicePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"ServiceUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__AlreadyUsedService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsNotForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenForgeFactory__CallerIsPausedForge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__InvalidData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"}],\"name\":\"TokenForgeFactory__ServiceNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenForge__InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ForgeFactory",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051613f316100395f395f81816121d2015281816121fb015261241b0152613f315ff3fe608060405260043610610178575f3560e01c806391d14854116100d1578063d547741f1161007c578063f4d8cab911610057578063f4d8cab914610505578063f51acaea14610524578063f8c8765e14610543575f5ffd5b8063d547741f146104a5578063e57997d2146104c4578063e98a5784146104e6575f5ffd5b8063ad3cb1cc116100ac578063ad3cb1cc146103d8578063c7525cba1461042d578063d3de1d7f1461045a575f5ffd5b806391d14854146103275780639ca92df914610397578063a217fddf146103c5575f5ffd5b80632f44ec09116101315780633ea113d01161010c5780633ea113d0146102e15780634f1ef2861461030057806352d1902d14610313575f5ffd5b80632f44ec091461028f57806336568abe146102a35780633a15d070146102c2575f5ffd5b8063217edd3411610161578063217edd34146101d1578063248a9ca3146102155780632f2ff15d14610270575f5ffd5b806301ffc9a71461017c5780630b7e40c6146101b0575b5f5ffd5b348015610187575f5ffd5b5061019b6101963660046131e0565b610562565b60405190151581526020015b60405180910390f35b3480156101bb575f5ffd5b506101cf6101ca366004613266565b6105fa565b005b3480156101dc575f5ffd5b506101f06101eb3660046132d4565b61093a565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b348015610220575f5ffd5b5061026261022f366004613399565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101a7565b34801561027b575f5ffd5b506101cf61028a3660046133b0565b610de8565b34801561029a575f5ffd5b50610262610e31565b3480156102ae575f5ffd5b506101cf6102bd3660046133b0565b610e60565b3480156102cd575f5ffd5b506101cf6102dc366004613266565b610ebe565b3480156102ec575f5ffd5b506101cf6102fb366004613266565b610f7f565b6101cf61030e36600461351c565b61125b565b34801561031e575f5ffd5b5061026261127a565b348015610332575f5ffd5b5061019b6103413660046133b0565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156103a2575f5ffd5b506103b66103b1366004613399565b6112a8565b6040516101a7939291906135b5565b3480156103d0575f5ffd5b506102625f81565b3480156103e3575f5ffd5b506104206040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a791906135f4565b348015610438575f5ffd5b5061044c610447366004613606565b611323565b6040516101a7929190613621565b348015610465575f5ffd5b50610479610474366004613644565b6113bc565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683529015156020830152016101a7565b3480156104b0575f5ffd5b506101cf6104bf3660046133b0565b61143e565b3480156104cf575f5ffd5b506104d8611481565b6040516101a7929190613676565b3480156104f1575f5ffd5b506101cf610500366004613266565b611608565b348015610510575f5ffd5b506101cf61051f36600461375f565b6118e4565b34801561052f575f5ffd5b506101cf61053e366004613644565b611a79565b34801561054e575f5ffd5b506101cf61055d3660046137aa565b611bd5565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806105f457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f6106267f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005b33611d54565b90505f86600281111561063b5761063b613803565b036106cc575f8061064e84860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f86d733c2654f9db1a830cfbbf2180ee3329101d853428a7ead4e85ddb3230d6089856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a45050610932565b60018660028111156106e0576106e0613803565b03610762575f806106f384860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847fa36201ae386aaacb89bc573b7c0ecc0913a93028f710f28ca29840454f44688689856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600286600281111561077657610776613803565b03610900575f806107898486018661385a565b915091508115610876575f5f5f838060200190518101906107aa91906138fc565b815192955090935091505f5b8181101561086c578473ffffffffffffffffffffffffffffffffffffffff168c897fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c38e88868151811061080b5761080b613973565b602002602001015188878151811061082557610825613973565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a46001016107b6565b50505050506108f9565b5f5f5f8380602001905181019061088d91906139a0565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907fa34d0136fb3d65c2424efbe91bfe903716015137d5dd5e965f59e7f0820eb2c3906060015b60405180910390a45050505b5050610932565b6040517f9570ddf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61096581611e32565b73ffffffffffffffffffffffffffffffffffffffff88166109d9576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8716610a48576040517f435db3480000000000000000000000000000000000000000000000000000000081527f76616c696461746f72000000000000000000000000000000000000000000000060048201526024016109d0565b5f859003610aa4576040517f435db3480000000000000000000000000000000000000000000000000000000081527f736572766963650000000000000000000000000000000000000000000000000060048201526024016109d0565b5f610ae387878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611e3f92505050565b90507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00610b307f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0383611e95565b15610b6a576040517f5502583a000000000000000000000000000000000000000000000000000000008152600481018390526024016109d0565b610cd35f83835f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015610bd9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610c1e91908101906139d4565b8d8d8b8b8f8f8a6001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b6002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001610c83989796959493929190613af1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610cbf9291602001613d25565b604051602081830303815290604052611ea0565b935073ffffffffffffffffffffffffffffffffffffffff8416610d44576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f726765206465706c6f79000000000000000000000000000000000000000060048201526024016109d0565b610d52600382018386611f8c565b5073ffffffffffffffffffffffffffffffffffffffff84165f81815260068301602090815260408083208690558583526007850190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f0453e3ffce9b0d4f994d54652037d3e15d87ec29aab0aafba90ac322edc87c2591a35050509695505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610e2181611e32565b610e2b8383611fb6565b50505050565b5f610e5b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036120d4565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610eaf576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb982826120de565b505050565b5f610ee87f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00610620565b90505f866002811115610efd57610efd613803565b03610900575f80610f1084860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f14bed88907ac31fe1ba4cf8981f80fd25e29877a08d1568c7210adc31ee9bc9589856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b5f610fa97f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00610620565b90505f866002811115610fbe57610fbe613803565b03611040575f80610fd184860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847fe5e0650048e213b8a668474e72c774ec69d8a6830270b7cc1c339fcd6a3b2e4189856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600186600281111561105457611054613803565b036110d6575f8061106784860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f28fcdc3e3355f548f5fe83373acf91a82e1e561ee22797f6e91f831d2b2b3e1389856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60028660028111156110ea576110ea613803565b03610900575f806110fd8486018661385a565b9150915081156111e0575f5f5f8380602001905181019061111e91906138fc565b815192955090935091505f5b8181101561086c578473ffffffffffffffffffffffffffffffffffffffff168c897f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b8e88868151811061117f5761117f613973565b602002602001015188878151811061119957611199613973565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a460010161112a565b5f5f5f838060200190518101906111f791906139a0565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f147c7aeaea1a7b56f7269db0eaea2435d8f957090d1bee114b6a6df6e556c42b906060016108ed565b6112636121ba565b61126c826122c0565b61127682826122ca565b5050565b5f611283612403565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60605f807f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00816112f87f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0387612472565b945090506113058161248f565b5f91825260079092016020526040902054909592945060ff16925050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081527f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e066020526040812054606091907f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e009080156113b55761139c8161248f565b5f82815260078401602052604090205490945060ff1692505b5050915091565b5f5f5f6113c884611e3f565b90507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005f6114167f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03846124cc565b955090508015611436575f83815260078301602052604090205460ff1693505b505050915091565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461147781611e32565b610e2b83836120de565b6060807f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005f6114cf7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e036120d4565b90505f8167ffffffffffffffff8111156114eb576114eb6133de565b60405190808252806020026020018201604052801561151e57816020015b60608152602001906001900390816115095790505b5090505f8267ffffffffffffffff81111561153b5761153b6133de565b604051908082528060200260200182016040528015611564578160200160208202803683370190505b5090505f5b838110156115fc575f61157f6003870183612472565b84848151811061159157611591613973565b602002602001018173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525081925050506115d68161248f565b8483815181106115e8576115e8613973565b602090810291909101015250600101611569565b50909590945092505050565b5f6116327f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00610620565b90505f86600281111561164757611647613803565b036116c9575f8061165a84860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f9612604afba70e4cf03261d7d86ca03d08911d9887aa41621dea929e34cfd7b189856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60018660028111156116dd576116dd613803565b0361175f575f806116f084860186613830565b915091508173ffffffffffffffffffffffffffffffffffffffff1687847f556d832108ef7e0123726b3e037bc1cd39023e8d397b242d8b36426550d8c12389856040516106bd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600286600281111561177357611773613803565b03610900575f806117868486018661385a565b915091508115611869575f5f5f838060200190518101906117a791906138fc565b815192955090935091505f5b8181101561086c578473ffffffffffffffffffffffffffffffffffffffff168c897f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f618e88868151811061180857611808613973565b602002602001015188878151811061182257611822613973565b6020908102919091018101516040805173ffffffffffffffffffffffffffffffffffffffff90951685529184019290925282015260600160405180910390a46001016117b3565b5f5f5f8380602001905181019061188091906139a0565b6040805173ffffffffffffffffffffffffffffffffffffffff8e811682526020820185905291810183905293965091945092508416908b9088907f3e4e312643f14a5cca74681c381ee1ed5862b2037f509b5c3d745522e8dc2f61906060016108ed565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61190e81611e32565b5f61191884611e3f565b90507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e006119657f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0383611e95565b61199e576040517f2a188b7f000000000000000000000000000000000000000000000000000000008152600481018390526024016109d0565b5f82815260078201602052604090205484151560ff909116151503611a725783156119fd575f828152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055611a36565b5f828152600782016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b817eb31832845d6567ecb46caeeee58acfcfddfcbba821c68dcbae9055ea4d795585604051611a69911515815260200190565b60405180910390a25b5050505050565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c611aa381611e32565b5f611aad83611e3f565b90507f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e005f80611afc7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e03856124cc565b9150915081611b3a576040517f2a188b7f000000000000000000000000000000000000000000000000000000008152600481018590526024016109d0565b611b4760038401856124da565b5073ffffffffffffffffffffffffffffffffffffffff81165f90815260068401602090815260408083208390558683526007860190915280822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555185917f475a09e552c07ad6aa3d8b6f3fb547091aae08290004e78277eb2ffb854ea21e91a2505050505050565b5f611bde6124e5565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015611c0a5750825b90505f8267ffffffffffffffff166001148015611c265750303b155b905081158015611c34575080155b15611c6b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611ccc5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b611cd461250d565b611cdc61250d565b611ce889898989612515565b8315611d495784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260068301602052604081205480611dca576040517f4b73e4e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016109d0565b5f81815260078501602052604090205460ff16611e2b576040517fae111e9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016109d0565b9392505050565b611e3c8133612add565b50565b5f5f829050601f81511115611e8257826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016109d091906135f4565b8051611e8d82613d39565b179392505050565b5f611e2b8383612b83565b5f83471015611ee4576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018590526044016109d0565b81515f03611f1e576040517f4ca249dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282516020840186f590503d151981151615611f3f576040513d5f823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff8116611e2b576040517fb06ebf3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611fae848473ffffffffffffffffffffffffffffffffffffffff8516612b8e565b949350505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166120cb575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556120673390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105f4565b5f9150506105f4565b5f6105f482612baa565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156120cb575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105f4565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061228757507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661226e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156122be576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61127681611e32565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561234f575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261234c91810190613d7e565b60015b61239d576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016109d0565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146123f9576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016109d0565b610eb98383612bb4565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146122be576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8080806124808686612c16565b909450925050505b9250929050565b60605f61249b83612c3f565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f8080806124808686612c7f565b5f611e2b8383612cb7565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006105f4565b6122be612cd3565b61251d612cd3565b73ffffffffffffffffffffffffffffffffffffffff841661258c576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024016109d0565b73ffffffffffffffffffffffffffffffffffffffff83166125fb576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f646500000000000000000000000000000000000060048201526024016109d0565b73ffffffffffffffffffffffffffffffffffffffff821661266a576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c00000000000000000000000000000000000000000060048201526024016109d0565b73ffffffffffffffffffffffffffffffffffffffff81166126d9576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c00000000000000000000000000000000000000000000000060048201526024016109d0565b8273ffffffffffffffffffffffffffffffffffffffff166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015612721573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261276691908101906139d4565b515f036127c1576040517f435db3480000000000000000000000000000000000000000000000000000000081527f666f72676550726f7879436f646520636f64650000000000000000000000000060048201526024016109d0565b8173ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa158015612809573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261284e9190810190613d95565b5173ffffffffffffffffffffffffffffffffffffffff8381169116146128c2576040517f435db3480000000000000000000000000000000000000000000000000000000081527f6469616d6f6e64496d706c20666163657400000000000000000000000000000060048201526024016109d0565b8073ffffffffffffffffffffffffffffffffffffffff16631bd8f7f16040518163ffffffff1660e01b81526004015f60405180830381865afa15801561290a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261294f9190810190613d95565b5173ffffffffffffffffffffffffffffffffffffffff8281169116146129c3576040517f435db3480000000000000000000000000000000000000000000000000000000081527f62617365496d706c20666163657400000000000000000000000000000000000060048201526024016109d0565b7f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00805473ffffffffffffffffffffffffffffffffffffffff8581167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161783557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e0180548683169084161790557f11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e02805491851691909216179055612aa87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f612d11565b612ab25f86611fb6565b506109327faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c86611fb6565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611276576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016109d0565b5f611e2b8383612db2565b5f8281526002840160205260408120829055611fae8484612dc9565b5f6105f482612dd4565b612bbd82612ddd565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612c0e57610eb98282612eab565b611276612f2a565b5f8080612c238585612f62565b5f81815260029690960160205260409095205494959350505050565b5f60ff8216601f8111156105f4576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600283016020526040812054819080612cac57612ca08585612b83565b92505f91506124889050565b600192509050612488565b5f8181526002830160205260408120819055611e2b8383612f6d565b612cdb612f78565b6122be576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f612d6a845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b5f8181526001830160205260408120541515611e2b565b5f611e2b8383612f96565b5f6105f4825490565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612e45576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109d0565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612ed49190613e8b565b5f60405180830381855af49150503d805f8114612f0c576040519150601f19603f3d011682016040523d82523d5f602084013e612f11565b606091505b5091509150612f21858383612fe2565b95945050505050565b34156122be576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611e2b8383613071565b5f611e2b8383613097565b5f612f816124e5565b5468010000000000000000900460ff16919050565b5f818152600183016020526040812054612fdb57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556105f4565b505f6105f4565b606082612ff757612ff282613171565b611e2b565b815115801561301b575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561306a576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016109d0565b5080611e2b565b5f825f01828154811061308657613086613973565b905f5260205f200154905092915050565b5f81815260018301602052604081205480156120cb575f6130b9600183613e96565b85549091505f906130cc90600190613e96565b905080821461312b575f865f0182815481106130ea576130ea613973565b905f5260205f200154905080875f01848154811061310a5761310a613973565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061313c5761313c613ece565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506105f4565b8051156131815780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081168114611e3c575f5ffd5b5f602082840312156131f0575f5ffd5b8135611e2b816131b3565b60038110611e3c575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114611e3c575f5ffd5b5f5f83601f840112613238575f5ffd5b50813567ffffffffffffffff81111561324f575f5ffd5b602083019150836020828501011115612488575f5ffd5b5f5f5f5f5f6080868803121561327a575f5ffd5b8535613285816131fb565b945060208601359350604086013561329c81613207565b9250606086013567ffffffffffffffff8111156132b7575f5ffd5b6132c388828901613228565b969995985093965092949392505050565b5f5f5f5f5f5f608087890312156132e9575f5ffd5b86356132f481613207565b9550602087013561330481613207565b9450604087013567ffffffffffffffff81111561331f575f5ffd5b61332b89828a01613228565b909550935050606087013567ffffffffffffffff81111561334a575f5ffd5b8701601f8101891361335a575f5ffd5b803567ffffffffffffffff811115613370575f5ffd5b8960208260051b8401011115613384575f5ffd5b60208201935080925050509295509295509295565b5f602082840312156133a9575f5ffd5b5035919050565b5f5f604083850312156133c1575f5ffd5b8235915060208301356133d381613207565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561342e5761342e6133de565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561347b5761347b6133de565b604052919050565b5f67ffffffffffffffff82111561349c5761349c6133de565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126134d7575f5ffd5b8135602083015f6134ef6134ea84613483565b613434565b9050828152858383011115613502575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f6040838503121561352d575f5ffd5b823561353881613207565b9150602083013567ffffffffffffffff811115613553575f5ffd5b61355f858286016134c8565b9150509250929050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081525f6135c76060830186613569565b73ffffffffffffffffffffffffffffffffffffffff94909416602083015250901515604090910152919050565b602081525f611e2b6020830184613569565b5f60208284031215613616575f5ffd5b8135611e2b81613207565b604081525f6136336040830185613569565b905082151560208301529392505050565b5f60208284031215613654575f5ffd5b813567ffffffffffffffff81111561366a575f5ffd5b611fae848285016134c8565b5f604082016040835280855180835260608501915060608160051b8601019250602087015f5b828110156136eb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08786030184526136d6858351613569565b9450602093840193919091019060010161369c565b5050505082810360208401528084518083526020830191506020860192505f5b8181101561373f57835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161370b565b50909695505050505050565b8035801515811461375a575f5ffd5b919050565b5f5f60408385031215613770575f5ffd5b823567ffffffffffffffff811115613786575f5ffd5b613792858286016134c8565b9250506137a16020840161374b565b90509250929050565b5f5f5f5f608085870312156137bd575f5ffd5b84356137c881613207565b935060208501356137d881613207565b925060408501356137e881613207565b915060608501356137f881613207565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f5f60408385031215613841575f5ffd5b823561384c81613207565b946020939093013593505050565b5f5f6040838503121561386b575f5ffd5b6135388361374b565b5f67ffffffffffffffff82111561388d5761388d6133de565b5060051b60200190565b5f82601f8301126138a6575f5ffd5b81516138b46134ea82613874565b8082825260208201915060208360051b8601019250858311156138d5575f5ffd5b602085015b838110156138f25780518352602092830192016138da565b5095945050505050565b5f5f5f6060848603121561390e575f5ffd5b835161391981613207565b602085015190935067ffffffffffffffff811115613935575f5ffd5b61394186828701613897565b925050604084015167ffffffffffffffff81111561395d575f5ffd5b61396986828701613897565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f606084860312156139b2575f5ffd5b83516139bd81613207565b602085015160409095015190969495509392505050565b5f602082840312156139e4575f5ffd5b815167ffffffffffffffff8111156139fa575f5ffd5b8201601f81018413613a0a575f5ffd5b8051613a186134ea82613483565b818152856020838501011115613a2c575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b8183526020830192505f815f5b84811015613aa0578135613a69816131b3565b7fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959190910190600101613a56565b5093949350505050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f60c0820173ffffffffffffffffffffffffffffffffffffffff8b16835273ffffffffffffffffffffffffffffffffffffffff8a16602084015260c060408401528088825260e08401905060e08960051b8501019150895f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18c3603015b8b821015613ca9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208786030184528235818112613bab575f5ffd5b8d018035613bb881613207565b73ffffffffffffffffffffffffffffffffffffffff1686526020810135613bde816131fb565b60038110613c13577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208701526040810135368290037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112613c4d575f5ffd5b0160208101903567ffffffffffffffff811115613c68575f5ffd5b8060051b3603821315613c79575f5ffd5b60606040880152613c8e606088018284613a49565b96505050602083019250602084019350600182019150613b6f565b505050508281036060840152613cc0818789613aaa565b915050613ce5608083018573ffffffffffffffffffffffffffffffffffffffff169052565b73ffffffffffffffffffffffffffffffffffffffff831660a08301529998505050505050505050565b5f81518060208401855e5f93019283525090919050565b5f611fae613d338386613d0e565b84613d0e565b80516020808301519190811015613d78577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b5f60208284031215613d8e575f5ffd5b5051919050565b5f60208284031215613da5575f5ffd5b815167ffffffffffffffff811115613dbb575f5ffd5b820160608185031215613dcc575f5ffd5b613dd461340b565b8151613ddf81613207565b81526020820151613def816131fb565b6020820152604082015167ffffffffffffffff811115613e0d575f5ffd5b80830192505084601f830112613e21575f5ffd5b8151613e2f6134ea82613874565b8082825260208201915060208360051b860101925087831115613e50575f5ffd5b6020850194505b82851015613e7b578451613e6a816131b3565b825260209485019490910190613e57565b6040840152509095945050505050565b5f611e2b8284613d0e565b818103818111156105f4577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220951a4e7ccca382207eb108b7f13153870f4640d76177e1e47b538aa17202ba3b64736f6c634300081c0033",
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
// the contract method with ID 0x217edd34.
//
// Solidity: function addService(address owner, address validator, string service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
func (forgeFactory *ForgeFactory) PackAddService(owner common.Address, validator common.Address, service string, addCuts []IDiamondCutFacetCut) []byte {
	enc, err := forgeFactory.abi.Pack("addService", owner, validator, service, addCuts)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAddService is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x217edd34.
//
// Solidity: function addService(address owner, address validator, string service, (address,uint8,bytes4[])[] addCuts) returns(address forge)
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
// Solidity: function allForges() view returns(string[], address[])
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
	Arg0 []string
	Arg1 []common.Address
}

// UnpackAllForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe57997d2.
//
// Solidity: function allForges() view returns(string[], address[])
func (forgeFactory *ForgeFactory) UnpackAllForges(data []byte) (AllForgesOutput, error) {
	out, err := forgeFactory.abi.Unpack("allForges", data)
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
	Service string
	Forge   common.Address
	Running bool
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(string service, address forge, bool running)
func (forgeFactory *ForgeFactory) UnpackForgeByIndex(data []byte) (ForgeByIndexOutput, error) {
	out, err := forgeFactory.abi.Unpack("forgeByIndex", data)
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
func (forgeFactory *ForgeFactory) PackForgeByService(service string) []byte {
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
// from invoking the contract method with ID 0xd3de1d7f.
//
// Solidity: function forgeByService(string service) view returns(address forge, bool running)
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
// Solidity: function isRunningForge(address forge) view returns(string service, bool running)
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
	Service string
	Running bool
}

// UnpackIsRunningForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc7525cba.
//
// Solidity: function isRunningForge(address forge) view returns(string service, bool running)
func (forgeFactory *ForgeFactory) UnpackIsRunningForge(data []byte) (IsRunningForgeOutput, error) {
	out, err := forgeFactory.abi.Unpack("isRunningForge", data)
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
// the contract method with ID 0xf4d8cab9.
//
// Solidity: function pauseService(string service, bool paused) returns()
func (forgeFactory *ForgeFactory) PackPauseService(service string, paused bool) []byte {
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
// the contract method with ID 0xf51acaea.
//
// Solidity: function removeService(string service) returns()
func (forgeFactory *ForgeFactory) PackRemoveService(service string) []byte {
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
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeFactory.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return forgeFactory.UnpackStringTooLongError(raw[4:])
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

// ForgeFactoryInvalidShortString represents a InvalidShortString error raised by the ForgeFactory contract.
type ForgeFactoryInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ForgeFactoryInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (forgeFactory *ForgeFactory) UnpackInvalidShortStringError(raw []byte) (*ForgeFactoryInvalidShortString, error) {
	out := new(ForgeFactoryInvalidShortString)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
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

// ForgeFactoryStringTooLong represents a StringTooLong error raised by the ForgeFactory contract.
type ForgeFactoryStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ForgeFactoryStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (forgeFactory *ForgeFactory) UnpackStringTooLongError(raw []byte) (*ForgeFactoryStringTooLong, error) {
	out := new(ForgeFactoryStringTooLong)
	if err := forgeFactory.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
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