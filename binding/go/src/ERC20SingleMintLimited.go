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

// ERC20SingleMintLimitedMetaData contains all meta data concerning the ERC20SingleMintLimited contract.
var ERC20SingleMintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"updateMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimits\",\"type\":\"uint256\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20SingleMintLimited__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20SingleMintLimited__InvalidInitialData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20SingleMintLimited",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051613efb6100395f395f81816117f50152818161181e0152611a3e0152613efb5ff3fe608060405260043610610291575f3560e01c806379cc679011610165578063ad3cb1cc116100c6578063da0239a61161007c578063e01d55c511610062578063e01d55c51461087f578063ec87621c1461089e578063fe2df3e8146108d1575f5ffd5b8063da0239a6146107fb578063dd62ed3e1461080f575f5ffd5b8063bdf7acce116100ac578063bdf7acce146107a9578063d505accf146107bd578063d547741f146107dc575f5ffd5b8063ad3cb1cc14610742578063b01932271461078a575f5ffd5b806395d89b411161011b578063a217fddf11610101578063a217fddf146106fc578063a40283d51461070f578063a9059cbb14610723575f5ffd5b806395d89b41146106a45780639ca92df9146106b8575f5ffd5b806384b0196e1161014b57806384b0196e146105ee57806387f453531461061557806391d1485414610634575f5ffd5b806379cc6790146105b05780637ecebe00146105cf575f5ffd5b8063355274ea1161020f5780634f1ef286116101c557806352d1902d116101ab57806352d1902d1461055c5780635c4e62c41461057057806370a0823114610591575f5ffd5b80634f1ef286146104df5780635068b4de146104f2575f5ffd5b806336568abe116101f557806336568abe1461048d57806340c10f19146104ac578063436ba626146104cb575f5ffd5b8063355274ea146104465780633644e51514610479575f5ffd5b806323b872dd116102645780632d9ed80d1161024a5780632d9ed80d146103b25780632f2ff15d146103e5578063313ce56714610406575f5ffd5b806323b872dd14610346578063248a9ca314610365575f5ffd5b806301ffc9a71461029557806306fdde03146102c9578063095ea7b3146102ea57806318160ddd14610309575b5f5ffd5b3480156102a0575f5ffd5b506102b46102af366004613598565b6108f0565b60405190151581526020015b60405180910390f35b3480156102d4575f5ffd5b506102dd61094b565b6040516102c09190613623565b3480156102f5575f5ffd5b506102b4610304366004613658565b610a03565b348015610314575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016102c0565b348015610351575f5ffd5b506102b4610360366004613680565b610a1a565b348015610370575f5ffd5b5061033861037f3660046136ba565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156103bd575f5ffd5b507f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0254610338565b3480156103f0575f5ffd5b506104046103ff3660046136d1565b610a30565b005b348015610411575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff90911681526020016102c0565b348015610451575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30054610338565b348015610484575f5ffd5b50610338610a79565b348015610498575f5ffd5b506104046104a73660046136d1565b610a87565b3480156104b7575f5ffd5b506104046104c6366004613658565b610ae5565b3480156104d6575f5ffd5b50610338610af3565b6104046104ed3660046137d7565b610b5e565b3480156104fd575f5ffd5b507f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0054604080516fffffffffffffffffffffffffffffffff83168152700100000000000000000000000000000000909204600f0b6020830152016102c0565b348015610567575f5ffd5b50610338610b79565b34801561057b575f5ffd5b50610584610ba7565b6040516102c09190613822565b34801561059c575f5ffd5b506103386105ab36600461387a565b610bd2565b3480156105bb575f5ffd5b506104046105ca366004613658565b610c22565b3480156105da575f5ffd5b506103386105e936600461387a565b610c37565b3480156105f9575f5ffd5b50610602610c41565b6040516102c09796959493929190613893565b348015610620575f5ffd5b506102b461062f36600461387a565b610d40565b34801561063f575f5ffd5b506102b461064e3660046136d1565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156106af575f5ffd5b506102dd610d6b565b3480156106c3575f5ffd5b506106d76106d23660046136ba565b610dbc565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c0565b348015610707575f5ffd5b506103385f81565b34801561071a575f5ffd5b50610338610de7565b34801561072e575f5ffd5b506102b461073d366004613658565b610e11565b34801561074d575f5ffd5b506102dd6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610795575f5ffd5b506104046107a4366004613962565b610e1c565b3480156107b4575f5ffd5b50610338611049565b3480156107c8575f5ffd5b506104046107d7366004613a3b565b611073565b3480156107e7575f5ffd5b506104046107f63660046136d1565b61123b565b348015610806575f5ffd5b5061033861127e565b34801561081a575f5ffd5b50610338610829366004613aa1565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b34801561088a575f5ffd5b506104046108993660046136ba565b6112ed565b3480156108a9575f5ffd5b506103387faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b3480156108dc575f5ffd5b506104046108eb366004613ac9565b6113d9565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f89145aae000000000000000000000000000000000000000000000000000000001480610945575061094582611489565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461098190613b4f565b80601f01602080910402602001604051908101604052809291908181526020018280546109ad90613b4f565b80156109f85780601f106109cf576101008083540402835291602001916109f8565b820191905f5260205f20905b8154815290600101906020018083116109db57829003601f168201915b505050505091505090565b5f33610a10818585611493565b5060019392505050565b5f610a268484846114a0565b90505b9392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610a69816114c3565b610a7383836114d0565b50505050565b5f610a826115ee565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610ad6576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ae082826115f7565b505050565b610aef82826116d3565b5050565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de01545f907f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0090610b449082906117ca565b610b52578060030154610b58565b80600201545b91505090565b610b666117dd565b610b6f826118e3565b610aef82826118ed565b5f610b82611a26565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6060610a827f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611a95565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b610c2d823383611aa1565b610aef8282611b88565b5f61094582611be2565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610c7f57506001810154155b610cea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064015b60405180910390fd5b610cf2611c0a565b610cfa611c5b565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6109457f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611c84565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161098190613b4f565b5f6109457f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611cb2565b5f610a827f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611cbd565b5f610a298383611cc6565b5f610e25611cd3565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610e515750825b90505f8267ffffffffffffffff166001148015610e6d5750303b155b905081158015610e7b575080155b15610eb2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610f135784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610f228d8d8d8d8d8d8d611cfb565b8551608014610f5d576040517f5431945d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f5f89806020019051810190610f759190613ba0565b93509350935093508b841015610fc1576040517fc013187a00000000000000000000000000000000000000000000000000000000815260048101859052602481018d9052604401610ce1565b610fcc838383611de3565b610fd584611ee8565b50505050831561103a5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b5f610a827f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00611f6f565b834211156110b0576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610ce1565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886111278c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61118e82611f7a565b90505f61119d82878787611fc1565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611224576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610ce1565b61122f8a8a8a611493565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611274816114c3565b610a7383836115f7565b5f5f6112a87f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f6112d37f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90508181116112e2575f6112e6565b8181035b9250505090565b5f6112f7816114c3565b815f03611352576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69740000000000000000000000000000000000000000000000006004820152602401610ce1565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0254604080519182526020820184905280517f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00927f864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e48292908290030190a16002019190915550565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c611403816114c3565b6135908261141357611fed611417565b6120405b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b85811015611480576114788288888481811061145a5761145a613bd3565b905060200201602081019061146f919061387a565b8563ffffffff16565b60010161143c565b50505050505050565b5f61094582612102565b610ae08383836001612198565b5f336114ad858285611aa1565b6114b8858585612302565b506001949350505050565b6114cd81336123ab565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166115e5575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556115813390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610945565b5f915050610945565b5f610a82612451565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156115e5575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610945565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de03547f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00905f61172183611f6f565b9050826001015481146117735760018301819055600283015460405181815290925081907ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d19060200160405180910390a25b50828110156117b8576040517f156d32c50000000000000000000000000000000000000000000000000000000081526004810184905260248101829052604401610ce1565b8281036003830155610a7384846124c4565b5f816117d584611f6f565b119392505050565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806118aa57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166118917f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156118e1576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610aef816114c3565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611972575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261196f91810190613c00565b60015b6119c0576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610ce1565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611a1c576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610ce1565b610ae0838361250f565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146118e1576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f610a2983612571565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610a735781811015611b7a576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610ce1565b610a7384848484035f612198565b73ffffffffffffffffffffffffffffffffffffffff8216611bd7576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b610aef825f836125ca565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610bf6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161098190613b4f565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610970565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610a29565b5f610a2983836125d5565b5f610945825490565b5f33610a10818585612302565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610945565b611d036125fb565b84515f03611d5f576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610ce1565b83515f03611dbb576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610ce1565b611dc58787612639565b611dcf85856127a2565b611dd8856127b4565b6114808383836127fb565b611deb6125fb565b821580611df6575080155b15611e4f576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6c696d6974206f7220706572696f6400000000000000000000000000000000006004820152602401610ce1565b5f7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0090506040518060400160405280611e87866128a7565b6fffffffffffffffffffffffffffffffff168152602001611ea785612900565b600f0b905280516020909101516fffffffffffffffffffffffffffffffff908116700100000000000000000000000000000000029116178155600201555050565b611ef06125fb565b805f03611f4b576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f63617000000000000000000000000000000000000000000000000000000000006004820152602401610ce1565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b5f610945824261294d565b5f610945611f866115ee565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611fd188888888612a7e565b925092509250611fe18282612b71565b50909695505050505050565b611ff78282612c74565b15610aef5760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff81166120af576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610ce1565b6120b98282612c95565b15610aef5760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061094557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610945565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516612208576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b73ffffffffffffffffffffffffffffffffffffffff8416612257576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b73ffffffffffffffffffffffffffffffffffffffff8086165f908152600183016020908152604080832093881683529290522083905581156122fb578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516122f291815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8316612351576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b73ffffffffffffffffffffffffffffffffffffffff82166123a0576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b610ae08383836125ca565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610aef576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610ce1565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61247b612cb6565b612483612d31565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6124cd33610d40565b612505576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610ce1565b610aef8282612d86565b61251882612de0565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561256957610ae08282612eae565b610aef612f2d565b6060815f018054806020026020016040519081016040528092919081815260200182805480156125be57602002820191905f5260205f20905b8154815260200190600101908083116125aa575b50505050509050919050565b610ae0838383612f65565b5f825f0182815481106125ea576125ea613bd3565b905f5260205f200154905092915050565b612603613025565b6118e1576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6126416125fb565b612649613043565b612651613043565b73ffffffffffffffffffffffffffffffffffffffff82166126c0576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610ce1565b6126ca5f836114d0565b506126f57faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836114d0565b5073ffffffffffffffffffffffffffffffffffffffff81161580159061274757508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b15612778576127767faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c826114d0565b505b610aef7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f61304b565b6127aa6125fb565b610aef82826130ec565b6127bc6125fb565b6114cd816040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525061314f565b6128036125fb565b827f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055815f14610ae05773ffffffffffffffffffffffffffffffffffffffff811661289c576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f696e697469616c526563697069656e74000000000000000000000000000000006004820152602401610ce1565b610ae05f82846131c1565b5f6fffffffffffffffffffffffffffffffff8211156128fc576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526080600482015260248101839052604401610ce1565b5090565b80600f81900b8114612948576040517f327269a70000000000000000000000000000000000000000000000000000000081526080600482015260248101839052604401610ce1565b919050565b81545f906fffffffffffffffffffffffffffffffff16810361299b576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff1682816129bd576129bd613c17565b8454919006909203917001000000000000000000000000000000009004600f0b5f036129ea575080610945565b82545f700100000000000000000000000000000000909104600f0b1315612a35578254612a2e907001000000000000000000000000000000009004600f0b83613c71565b9050610945565b82545f90612a59907001000000000000000000000000000000009004600f0b613c84565b905080831015612a6c575f915050610945565b612a768184613cba565b915050610945565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115612ab757505f91506003905082612b67565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612b08573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612b5e57505f925060019150829050612b67565b92505f91508190505b9450945094915050565b5f826003811115612b8457612b84613ccd565b03612b8d575050565b6001826003811115612ba157612ba1613ccd565b03612bd8576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115612bec57612bec613ccd565b03612c26576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610ce1565b6003826003811115612c3a57612c3a613ccd565b03610aef576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610ce1565b5f610a298373ffffffffffffffffffffffffffffffffffffffff841661338e565b5f610a298373ffffffffffffffffffffffffffffffffffffffff8416613468565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612ce1611c0a565b805190915015612cf957805160209091012092915050565b81548015612d08579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612d5c611c5b565b805190915015612d7457805160209091012092915050565b60018201548015612d08579392505050565b73ffffffffffffffffffffffffffffffffffffffff8216612dd5576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610ce1565b610aef5f83836125ca565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612e48576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610ce1565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612ed79190613cfa565b5f60405180830381855af49150503d805f8114612f0f576040519150601f19603f3d011682016040523d82523d5f602084013e612f14565b606091505b5091509150612f248583836134b4565b95945050505050565b34156118e1576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f70838383613543565b73ffffffffffffffffffffffffffffffffffffffff8316610ae0575f612fb47f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f612fdf7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b9050818111156122fb576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610ce1565b5f61302e611cd3565b5468010000000000000000900460ff16919050565b6118e16125fb565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f6130a4845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b6130f46125fb565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036131408482613d54565b5060048101610a738382613d54565b6131576125fb565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026131a38482613d54565b50600381016131b28382613d54565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661321b5781816002015f8282546132109190613c71565b909155506132cb9050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156132a0576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610ce1565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff83166132f6576002810180548390039055613321565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161338091815260200190565b60405180910390a350505050565b5f81815260018301602052604081205480156115e5575f6133b0600183613cba565b85549091505f906133c390600190613cba565b9050808214613422575f865f0182815481106133e1576133e1613bd3565b905f5260205f200154905080875f01848154811061340157613401613bd3565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061343357613433613e6b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610945565b5f8181526001830160205260408120546134ad57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610945565b505f610945565b6060826134c9576134c48261354e565b610a29565b81511580156134ed575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561353c576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610ce1565b5080610a29565b610ae08383836131c1565b80511561355e5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118e1613e98565b5f602082840312156135a8575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a29575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610a2960208301846135d7565b803573ffffffffffffffffffffffffffffffffffffffff81168114612948575f5ffd5b5f5f60408385031215613669575f5ffd5b61367283613635565b946020939093013593505050565b5f5f5f60608486031215613692575f5ffd5b61369b84613635565b92506136a960208501613635565b929592945050506040919091013590565b5f602082840312156136ca575f5ffd5b5035919050565b5f5f604083850312156136e2575f5ffd5b823591506136f260208401613635565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112613737575f5ffd5b8135602083015f5f67ffffffffffffffff841115613757576137576136fb565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156137a4576137a46136fb565b6040528381529050808284018710156137bb575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156137e8575f5ffd5b6137f183613635565b9150602083013567ffffffffffffffff81111561380c575f5ffd5b61381885828601613728565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561386f57835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161383b565b509095945050505050565b5f6020828403121561388a575f5ffd5b610a2982613635565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6138cd60e08301896135d7565b82810360408401526138df81896135d7565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015613941578351835260209384019390920191600101613923565b50909b9a5050505050505050505050565b803560ff81168114612948575f5ffd5b5f5f5f5f5f5f5f5f610100898b03121561397a575f5ffd5b61398389613635565b975061399160208a01613635565b9650604089013567ffffffffffffffff8111156139ac575f5ffd5b6139b88b828c01613728565b965050606089013567ffffffffffffffff8111156139d4575f5ffd5b6139e08b828c01613728565b9550506139ef60808a01613952565b935060a08901359250613a0460c08a01613635565b915060e089013567ffffffffffffffff811115613a1f575f5ffd5b613a2b8b828c01613728565b9150509295985092959890939650565b5f5f5f5f5f5f5f60e0888a031215613a51575f5ffd5b613a5a88613635565b9650613a6860208901613635565b95506040880135945060608801359350613a8460808901613952565b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215613ab2575f5ffd5b613abb83613635565b91506136f260208401613635565b5f5f5f60408486031215613adb575f5ffd5b833567ffffffffffffffff811115613af1575f5ffd5b8401601f81018613613b01575f5ffd5b803567ffffffffffffffff811115613b17575f5ffd5b8660208260051b8401011115613b2b575f5ffd5b6020918201945092508401358015158114613b44575f5ffd5b809150509250925092565b600181811c90821680613b6357607f821691505b602082108103613b9a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f5f5f5f60808587031215613bb3575f5ffd5b505082516020840151604085015160609095015191969095509092509050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215613c10575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561094557610945613c44565b5f7f80000000000000000000000000000000000000000000000000000000000000008203613cb457613cb4613c44565b505f0390565b8181038181111561094557610945613c44565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f821115610ae057805f5260205f20601f840160051c81016020851015613d355750805b601f840160051c820191505b818110156122fb575f8155600101613d41565b815167ffffffffffffffff811115613d6e57613d6e6136fb565b613d8281613d7c8454613b4f565b84613d10565b6020601f821160018114613dd3575f8315613d9d5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556122fb565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015613e205787850151825560209485019460019092019101613e00565b5084821015613e5c57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212200ee2b211c4ba9553efd4d9b35a83bb4d66fd922bb808761ea726d798198ad67164736f6c634300081c0033",
}

// ERC20SingleMintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20SingleMintLimited struct {
	abi abi.ABI
}

// NewERC20SingleMintLimited creates a new instance of ERC20SingleMintLimited.
func NewERC20SingleMintLimited() *ERC20SingleMintLimited {
	parsed, err := ERC20SingleMintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20SingleMintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20SingleMintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackMANAGERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackAvailableMintCapacity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x436ba626.
//
// Solidity: function availableMintCapacity() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackAvailableMintCapacity() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("availableMintCapacity")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x436ba626.
//
// Solidity: function availableMintCapacity() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackAvailableMintCapacity(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("availableMintCapacity", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackCap() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("cap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackDecimals() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackEip712Domain() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("eip712Domain", data)
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

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("forgeByIndex", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackForgeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackForgeCount() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("forgeCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackForges() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0193227.
//
// Solidity: function initialize(address owner, address manager, string name, string symbol, uint8 decimals, uint256 initialSupply, address initialRecipient, bytes data) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, initialRecipient common.Address, data []byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, initialRecipient, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMaxMintPerPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackMaxMintPerPeriod() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("maxMintPerPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackMaxMintPerPeriod(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("maxMintPerPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackName() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPeriodConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5068b4de.
//
// Solidity: function periodConfig() view returns(uint256, int256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackPeriodConfig() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("periodConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// PeriodConfigOutput serves as a container for the return parameters of contract
// method PeriodConfig.
type PeriodConfigOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
}

// UnpackPeriodConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5068b4de.
//
// Solidity: function periodConfig() view returns(uint256, int256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodConfig(data []byte) (PeriodConfigOutput, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("periodConfig", data)
	outstruct := new(PeriodConfigOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPeriodStartTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbdf7acce.
//
// Solidity: function periodStartTime() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackPeriodStartTime() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("periodStartTime")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbdf7acce.
//
// Solidity: function periodStartTime() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodStartTime(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("periodStartTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackProxiableUUID() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRemainingSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackRemainingSupply() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("remainingSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackSymbol() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackTotalSupply() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateMintLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe01d55c5.
//
// Solidity: function updateMintLimit(uint256 newLimit) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackUpdateMintLimit(newLimit *big.Int) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("updateMintLimit", newLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20SingleMintLimitedApproval represents a Approval event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedApproval) ContractEventName() string {
	return ERC20SingleMintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20SingleMintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20SingleMintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20SingleMintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedForgeAdded) ContractEventName() string {
	return ERC20SingleMintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20SingleMintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedForgeRemoved) ContractEventName() string {
	return ERC20SingleMintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20SingleMintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedInitialized represents a Initialized event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedInitialized) ContractEventName() string {
	return ERC20SingleMintLimitedInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackInitializedEvent(log *types.Log) (*ERC20SingleMintLimitedInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedInitialized)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedMintLimitUpdated struct {
	OldLimits *big.Int
	NewLimits *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20SingleMintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256 oldLimits, uint256 newLimits)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20SingleMintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedPeriodStarted) ContractEventName() string {
	return ERC20SingleMintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20SingleMintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20SingleMintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20SingleMintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedRoleGranted represents a RoleGranted event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedRoleGranted) ContractEventName() string {
	return ERC20SingleMintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20SingleMintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedRoleRevoked) ContractEventName() string {
	return ERC20SingleMintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20SingleMintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedTransfer represents a Transfer event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedTransfer) ContractEventName() string {
	return ERC20SingleMintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTransferEvent(log *types.Log) (*ERC20SingleMintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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

// ERC20SingleMintLimitedUpgraded represents a Upgraded event raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20SingleMintLimitedUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20SingleMintLimitedUpgraded) ContractEventName() string {
	return ERC20SingleMintLimitedUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackUpgradedEvent(log *types.Log) (*ERC20SingleMintLimitedUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20SingleMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20SingleMintLimitedUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20SingleMintLimited.abi.Events[event].Inputs {
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
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20PeriodMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20PeriodMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20PeriodMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20SingleMintLimitedCapTooLow"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20SingleMintLimitedCapTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC20SingleMintLimitedInvalidInitialData"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC20SingleMintLimitedInvalidInitialDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20SingleMintLimited.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20SingleMintLimited.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20SingleMintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20SingleMintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20SingleMintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20SingleMintLimitedAccessControlBadConfirmation)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20SingleMintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20SingleMintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20SingleMintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20SingleMintLimitedAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackAddressEmptyCodeError(raw []byte) (*ERC20SingleMintLimitedAddressEmptyCode, error) {
	out := new(ERC20SingleMintLimitedAddressEmptyCode)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20SingleMintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20SingleMintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20SingleMintLimitedECDSAInvalidSignature)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20SingleMintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20SingleMintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20SingleMintLimitedECDSAInvalidSignatureLength)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20SingleMintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20SingleMintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20SingleMintLimitedECDSAInvalidSignatureS)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20SingleMintLimitedERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20SingleMintLimitedERC1967InvalidImplementation, error) {
	out := new(ERC20SingleMintLimitedERC1967InvalidImplementation)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20SingleMintLimitedERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC1967NonPayableError(raw []byte) (*ERC20SingleMintLimitedERC1967NonPayable, error) {
	out := new(ERC20SingleMintLimitedERC1967NonPayable)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20SingleMintLimitedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20SingleMintLimitedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20SingleMintLimitedERC20CapableERC20ExceededCap)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20SingleMintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20SingleMintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20SingleMintLimitedERC20InsufficientAllowance)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20SingleMintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20SingleMintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20SingleMintLimitedERC20InsufficientBalance)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20SingleMintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20SingleMintLimitedERC20InvalidApprover, error) {
	out := new(ERC20SingleMintLimitedERC20InvalidApprover)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20SingleMintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20SingleMintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20SingleMintLimitedERC20InvalidReceiver)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20SingleMintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20SingleMintLimitedERC20InvalidSender, error) {
	out := new(ERC20SingleMintLimitedERC20InvalidSender)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20SingleMintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20SingleMintLimitedERC20InvalidSpender, error) {
	out := new(ERC20SingleMintLimitedERC20InvalidSpender)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20PeriodMintLimitExceedsPeriodLimit represents a ERC20PeriodMintLimit__ExceedsPeriodLimit error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20PeriodMintLimitExceedsPeriodLimit struct {
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func ERC20SingleMintLimitedERC20PeriodMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0x156d32c58749d9c638f5b3f8e359a7ee3d1aa2d1c7d31a6301ce10260bfd45b4")
}

// UnpackERC20PeriodMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20SingleMintLimitedERC20PeriodMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20SingleMintLimitedERC20PeriodMintLimitExceedsPeriodLimit)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLength represents a ERC20PeriodMintLimit__InvalidLength error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0xf38583ae0b0acd19cabaa28c7b388b3f7e6260d4a4adbf38b03fb876a11643f0")
}

// UnpackERC20PeriodMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20PeriodMintLimitInvalidLengthError(raw []byte) (*ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLength, error) {
	out := new(ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLength)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLimitData represents a ERC20PeriodMintLimit__InvalidLimitData error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLimitData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0x436572d08cf45a9b922f4897fbf93bb0e84720cc3edcbc7328ffcb075921f9a8")
}

// UnpackERC20PeriodMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20PeriodMintLimitInvalidLimitDataError(raw []byte) (*ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLimitData, error) {
	out := new(ERC20SingleMintLimitedERC20PeriodMintLimitInvalidLimitData)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20SingleMintLimitedCapTooLow represents a ERC20SingleMintLimited__CapTooLow error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20SingleMintLimitedCapTooLow struct {
	Cap           *big.Int
	InitialSupply *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20SingleMintLimited__CapTooLow(uint256 cap, uint256 initialSupply)
func ERC20SingleMintLimitedERC20SingleMintLimitedCapTooLowErrorID() common.Hash {
	return common.HexToHash("0xc013187a3eb3b849cd2890f2bb70ccd911b8c1a271c56b4edf02470231698fb3")
}

// UnpackERC20SingleMintLimitedCapTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20SingleMintLimited__CapTooLow(uint256 cap, uint256 initialSupply)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20SingleMintLimitedCapTooLowError(raw []byte) (*ERC20SingleMintLimitedERC20SingleMintLimitedCapTooLow, error) {
	out := new(ERC20SingleMintLimitedERC20SingleMintLimitedCapTooLow)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20SingleMintLimitedCapTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC20SingleMintLimitedInvalidInitialData represents a ERC20SingleMintLimited__InvalidInitialData error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC20SingleMintLimitedInvalidInitialData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20SingleMintLimited__InvalidInitialData()
func ERC20SingleMintLimitedERC20SingleMintLimitedInvalidInitialDataErrorID() common.Hash {
	return common.HexToHash("0x5431945d26964ffd19ec54333e6d27b8a6612e38bb7349da2dc2e9b656b3db7f")
}

// UnpackERC20SingleMintLimitedInvalidInitialDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20SingleMintLimited__InvalidInitialData()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC20SingleMintLimitedInvalidInitialDataError(raw []byte) (*ERC20SingleMintLimitedERC20SingleMintLimitedInvalidInitialData, error) {
	out := new(ERC20SingleMintLimitedERC20SingleMintLimitedInvalidInitialData)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC20SingleMintLimitedInvalidInitialData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20SingleMintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20SingleMintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20SingleMintLimitedERC2612ExpiredSignature)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20SingleMintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20SingleMintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20SingleMintLimitedERC2612InvalidSigner)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedFailedCall represents a FailedCall error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20SingleMintLimitedFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackFailedCallError(raw []byte) (*ERC20SingleMintLimitedFailedCall, error) {
	out := new(ERC20SingleMintLimitedFailedCall)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20SingleMintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20SingleMintLimitedInvalidAccountNonce, error) {
	out := new(ERC20SingleMintLimitedInvalidAccountNonce)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedInvalidInitialization represents a InvalidInitialization error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20SingleMintLimitedInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackInvalidInitializationError(raw []byte) (*ERC20SingleMintLimitedInvalidInitialization, error) {
	out := new(ERC20SingleMintLimitedInvalidInitialization)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedNotInitializing represents a NotInitializing error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20SingleMintLimitedNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackNotInitializingError(raw []byte) (*ERC20SingleMintLimitedNotInitializing, error) {
	out := new(ERC20SingleMintLimitedNotInitializing)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20SingleMintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20SingleMintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20SingleMintLimitedPeriodManagerInvalidDuration)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20SingleMintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20SingleMintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20SingleMintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20SingleMintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20SingleMintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20SingleMintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20SingleMintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20SingleMintLimitedTokenBaseNullInput, error) {
	out := new(ERC20SingleMintLimitedTokenBaseNullInput)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20SingleMintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20SingleMintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20SingleMintLimitedTokenBaseOnlyForge)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20SingleMintLimitedUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20SingleMintLimitedUUPSUnauthorizedCallContext, error) {
	out := new(ERC20SingleMintLimitedUUPSUnauthorizedCallContext)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20SingleMintLimitedUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20SingleMintLimited contract.
type ERC20SingleMintLimitedUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20SingleMintLimitedUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20SingleMintLimitedUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20SingleMintLimitedUUPSUnsupportedProxiableUUID)
	if err := eRC20SingleMintLimited.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
