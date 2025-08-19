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

// ERC20MultiMintLimitedMetaData contains all meta data concerning the ERC20MultiMintLimited contract.
var ERC20MultiMintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacities\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriods\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfigs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTimes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"updateMintLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"oldLimits\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20MultiMintLimited__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20MultiMintLimited__InvalidInitialData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodsMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offsetSeconds\",\"type\":\"int256\"}],\"name\":\"PeriodManager__InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20MultiMintLimited",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051614a8e6100395f395f8181611e0101528181611e2a015261204a0152614a8e5ff3fe608060405260043610610290575f3560e01c80637ecebe0011610165578063b0193227116100c6578063d547741f1161007c578063dd62ed3e11610062578063dd62ed3e146107d7578063ec87621c14610847578063fe2df3e81461087a575f5ffd5b8063d547741f146107a4578063da0239a6146107c3575f5ffd5b8063ccf6d724116100ac578063ccf6d7241461074f578063cdd760b914610763578063d505accf14610785575f5ffd5b8063b01932271461071c578063c3ac4c2d1461073b575f5ffd5b80639ca92df91161011b578063a40283d511610101578063a40283d5146106a1578063a9059cbb146106b5578063ad3cb1cc146106d4575f5ffd5b80639ca92df91461064a578063a217fddf1461068e575f5ffd5b806387f453531161014b57806387f45353146105a757806391d14854146105c657806395d89b4114610636575f5ffd5b80637ecebe001461056157806384b0196e14610580575f5ffd5b8063355274ea1161020f5780634f1ef286116101c55780635c4e62c4116101ab5780635c4e62c41461050257806370a082311461052357806379cc679014610542575f5ffd5b80634f1ef286146104db57806352d1902d146104ee575f5ffd5b806336568abe116101f557806336568abe1461047e57806337e653ea1461049d57806340c10f19146104bc575f5ffd5b8063355274ea146104375780633644e5151461046a575f5ffd5b806318160ddd11610264578063248a9ca31161024a578063248a9ca3146103895780632f2ff15d146103d6578063313ce567146103f7575f5ffd5b806318160ddd1461032d57806323b872dd1461036a575f5ffd5b806286ced81461029457806301ffc9a7146102be57806306fdde03146102ed578063095ea7b31461030e575b5f5ffd5b34801561029f575f5ffd5b506102a8610899565b6040516102b59190613e9d565b60405180910390f35b3480156102c9575f5ffd5b506102dd6102d8366004613eaf565b610981565b60405190151581526020016102b5565b3480156102f8575f5ffd5b506103016109dc565b6040516102b59190613f3a565b348015610319575f5ffd5b506102dd610328366004613f6f565b610a94565b348015610338575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016102b5565b348015610375575f5ffd5b506102dd610384366004613f97565b610aab565b348015610394575f5ffd5b5061035c6103a3366004613fd1565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156103e1575f5ffd5b506103f56103f0366004613fe8565b610ac1565b005b348015610402575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff90911681526020016102b5565b348015610442575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005461035c565b348015610475575f5ffd5b5061035c610b0a565b348015610489575f5ffd5b506103f5610498366004613fe8565b610b18565b3480156104a8575f5ffd5b506103f56104b736600461405a565b610b76565b3480156104c7575f5ffd5b506103f56104d6366004613f6f565b610d89565b6103f56104e93660046141a5565b610d97565b3480156104f9575f5ffd5b5061035c610db2565b34801561050d575f5ffd5b50610516610de0565b6040516102b591906141f0565b34801561052e575f5ffd5b5061035c61053d366004614248565b610e0b565b34801561054d575f5ffd5b506103f561055c366004613f6f565b610e5b565b34801561056c575f5ffd5b5061035c61057b366004614248565b610e70565b34801561058b575f5ffd5b50610594610e7a565b6040516102b59796959493929190614261565b3480156105b2575f5ffd5b506102dd6105c1366004614248565b610f74565b3480156105d1575f5ffd5b506102dd6105e0366004613fe8565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b348015610641575f5ffd5b50610301610f9f565b348015610655575f5ffd5b50610669610664366004613fd1565b610ff0565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102b5565b348015610699575f5ffd5b5061035c5f81565b3480156106ac575f5ffd5b5061035c61101b565b3480156106c0575f5ffd5b506102dd6106cf366004613f6f565b611045565b3480156106df575f5ffd5b506103016040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610727575f5ffd5b506103f5610736366004614307565b611050565b348015610746575f5ffd5b506102a861127d565b34801561075a575f5ffd5b506102a8611422565b34801561076e575f5ffd5b5061077761149a565b6040516102b59291906143e0565b348015610790575f5ffd5b506103f561079f36600461442c565b611618565b3480156107af575f5ffd5b506103f56107be366004613fe8565b6117e0565b3480156107ce575f5ffd5b5061035c611823565b3480156107e2575f5ffd5b5061035c6107f1366004614492565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610852575f5ffd5b5061035c7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b348015610885575f5ffd5b506103f56108943660046144ba565b611892565b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d0180546060917f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00915f8167ffffffffffffffff8111156108fb576108fb614099565b604051908082528060200260200182016040528015610924578160200160208202803683370190505b5090505f5b828110156109785761095384828154811061094657610946614511565b905f5260205f2001611939565b82828151811061096557610965614511565b6020908102919091010152600101610929565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f89145aae0000000000000000000000000000000000000000000000000000000014806109d657506109d682611944565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b9050806003018054610a129061453e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3e9061453e565b8015610a895780601f10610a6057610100808354040283529160200191610a89565b820191905f5260205f20905b815481529060010190602001808311610a6c57829003601f168201915b505050505091505090565b5f33610aa181858561194e565b5060019392505050565b5f610ab784848461195b565b90505b9392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610afa8161197e565b610b04838361198b565b50505050565b5f610b13611aa9565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610b67576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b718282611ab2565b505050565b5f610b808161197e565b815f819003610be2576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d697473000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d0080548214610c3d576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805b83811015610d33575f878783818110610c5b57610c5b614511565b905060200201359050805f03610cbf576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69747300000000000000000000000000000000000000000000006004820152602401610bd9565b828111158015610cef57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114155b15610d29576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101839052602401610bd9565b9150600101610c40565b507f89be4a9eb7586334d719b1c87a0fd1e8306235335bf257d957c53aa7b22fcb74826002018787604051610d6a9392919061458f565b60405180910390a1610d80600283018787613dc9565b50505050505050565b610d938282611b8e565b5050565b610d9f611de9565b610da882611eef565b610d938282611ef9565b5f610dbb612032565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6060610b137f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474006120a1565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b610e668233836120ad565b610d938282612194565b5f6109d6826121ee565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610eb857506001810154155b610f1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610bd9565b610f26612216565b610f2e612267565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6109d67f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083612290565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610a129061453e565b5f6109d67f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400836122be565b5f610b137f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474006122c9565b5f610aba83836122d2565b5f6110596122df565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f811580156110855750825b90505f8267ffffffffffffffff1660011480156110a15750303b155b9050811580156110af575080155b156110e6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156111475784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6111568d8d8d8d8d8d8d612307565b6080865111611191576040517f94c8879f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f5f898060200190518101906111a991906146a1565b93509350935093508b8410156111f5576040517fce739feb00000000000000000000000000000000000000000000000000000000815260048101859052602481018d9052604401610bd9565b6112008383836123ef565b611209846126e1565b50505050831561126e5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b60607f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d005f6112a9610899565b90505f826003018054806020026020016040519081016040528092919081815260200182805480156112f857602002820191905f5260205f20905b8154815260200190600101908083116112e4575b505086549394505f92508391505067ffffffffffffffff81111561131e5761131e614099565b604051908082528060200260200182016040528015611347578160200160208202803683370190505b5090505f5b82811015611418575f84828151811061136757611367614511565b6020026020010151905085828151811061138357611383614511565b602002602001015181036113d2578660040182815481106113a6576113a6614511565b905f5260205f2001548383815181106113c1576113c1614511565b60200260200101818152505061140f565b8660020182815481106113e7576113e7614511565b905f5260205f20015483838151811061140257611402614511565b6020026020010181815250505b5060010161134c565b5095945050505050565b60607f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d0060020180548060200260200160405190810160405280929190818152602001828054801561149057602002820191905f5260205f20905b81548152602001906001019080831161147c575b5050505050905090565b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00805460609182915f8167ffffffffffffffff8111156114dc576114dc614099565b604051908082528060200260200182016040528015611505578160200160208202803683370190505b5090505f8267ffffffffffffffff81111561152257611522614099565b60405190808252806020026020018201604052801561154b578160200160208202803683370190505b509050600184015f5b8481101561160b5781818154811061156e5761156e614511565b5f9182526020909120015484516fffffffffffffffffffffffffffffffff909116908590839081106115a2576115a2614511565b6020026020010181815250508181815481106115c0576115c0614511565b5f918252602090912001548351700100000000000000000000000000000000909104600f0b908490839081106115f8576115f8614511565b6020908102919091010152600101611554565b5091969095509350505050565b83421115611655576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610bd9565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886116cc8c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61173382612768565b90505f611742828787876127af565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146117c9576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610bd9565b6117d48a8a8a61194e565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546118198161197e565b610b048383611ab2565b5f5f61184d7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f6118787f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b9050818111611887575f61188b565b8181035b9250505090565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6118bc8161197e565b613e0e826118cc576127db6118d0565b61282e5b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b85811015610d80576119318288888481811061191357611913614511565b90506020020160208101906119289190614248565b8563ffffffff16565b6001016118f5565b5f6109d682426128f0565b5f6109d682612a6c565b610b718383836001612b02565b5f336119688582856120ad565b611973858585612c6c565b506001949350505050565b6119888133612d15565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611aa0575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611a3c3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506109d6565b5f9150506109d6565b5f610b13612dbb565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611aa0575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506109d6565b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d007f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d037f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d045f611bfa610899565b84549091505f5b81811015611dde575f5f5f868481548110611c1e57611c1e614511565b905f5260205f200154888581548110611c3957611c39614511565b905f5260205f200154878681518110611c5457611c54614511565b6020026020010151925092509250808214611ce8575f896002018581548110611c7f57611c7f614511565b905f5260205f200154905081898681548110611c9d57611c9d614511565b905f5260205f200181905550809350817ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d182604051611cde91815260200190565b60405180910390a2505b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303611d325782878581548110611d2257611d22614511565b5f91825260209091200155611dd0565b89831015611daf57886001018481548110611d4f57611d4f614511565b5f918252602090912001546040517faace2cbe0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff9091166004820152602481018b905260448101849052606401610bd9565b898303878581548110611dc457611dc4614511565b5f918252602090912001555b505050806001019050611c01565b50610d808787612e2e565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611eb657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611e9d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611eed576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610d938161197e565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611f7e575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611f7b91810190614793565b60015b611fcc576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610bd9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612028576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610bd9565b610b718383612e79565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611eed576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f610aba83612edb565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610b045781811015612186576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610bd9565b610b0484848484035f612b02565b73ffffffffffffffffffffffffffffffffffffffff82166121e3576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b610d93825f83612f34565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610e2f565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10091610a129061453e565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610a01565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610aba565b5f610aba8383612f3f565b5f6109d6825490565b5f33610aa1818585612c6c565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006109d6565b61230f612f65565b84515f0361236b576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610bd9565b83515f036123c7576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610bd9565b6123d18787612fa3565b6123db858561310c565b6123e48561311e565b610d80838383613165565b6123f7612f65565b8051801580612407575083518114155b80612413575082518114155b1561244a576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d005f80805b8481101561260a575f5f89838151811061248b5761248b614511565b60200260200101518884815181106124a5576124a5614511565b602002602001015191509150815f14806124bd575080155b15612516576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6c696d697473206f72206475726174696f6e73000000000000000000000000006004820152602401610bd9565b84821115806125255750838111155b1561255f576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101849052602401610bd9565b85600101604051806040016040528061257785613211565b6fffffffffffffffffffffffffffffffff1681526020016125b08c87815181106125a3576125a3614511565b602002602001015161326a565b600f0b90528154600181810184555f938452602093849020835194909301516fffffffffffffffffffffffffffffffff9081167001000000000000000000000000000000000294169390931791015591945092500161246f565b5083835584516126239060028501906020880190613e16565b508367ffffffffffffffff81111561263d5761263d614099565b604051908082528060200260200182016040528015612666578160200160208202803683370190505b50805161267d916003860191602090910190613e16565b508367ffffffffffffffff81111561269757612697614099565b6040519080825280602002602001820160405280156126c0578160200160208202803683370190505b5080516126d7916004860191602090910190613e16565b5050505050505050565b6126e9612f65565b805f03612744576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f63617000000000000000000000000000000000000000000000000000000000006004820152602401610bd9565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b5f6109d6612774611aa9565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f6127bf888888886132b7565b9250925092506127cf82826133aa565b50909695505050505050565b6127e582826134ad565b15610d935760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff811661289d576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610bd9565b6128a782826134ce565b15610d935760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b81545f906fffffffffffffffffffffffffffffffff16810361293e576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff168281612960576129606147aa565b8454919006909203917001000000000000000000000000000000009004600f0b5f0361298d5750806109d6565b82545f700100000000000000000000000000000000909104600f0b13156129d85782546129d1907001000000000000000000000000000000009004600f0b83614804565b90506109d6565b82545f906129fc907001000000000000000000000000000000009004600f0b614817565b905080831015612a5a5783546040517ff41373a600000000000000000000000000000000000000000000000000000000815260048101859052700100000000000000000000000000000000909104600f0b6024820152604401610bd9565b612a64818461484d565b9150506109d6565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109d657507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146109d6565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516612b72576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b73ffffffffffffffffffffffffffffffffffffffff8416612bc1576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115612c65578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051612c5c91815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8316612cbb576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b73ffffffffffffffffffffffffffffffffffffffff8216612d0a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b610b71838383612f34565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610d93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610bd9565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612de56134ef565b612ded61356a565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b612e3733610f74565b612e6f576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610bd9565b610d9382826135bf565b612e8282613619565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612ed357610b7182826136e7565b610d93613766565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612f2857602002820191905f5260205f20905b815481526020019060010190808311612f14575b50505050509050919050565b610b7183838361379e565b5f825f018281548110612f5457612f54614511565b905f5260205f200154905092915050565b612f6d61385e565b611eed576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fab612f65565b612fb361387c565b612fbb61387c565b73ffffffffffffffffffffffffffffffffffffffff821661302a576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610bd9565b6130345f8361198b565b5061305f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361198b565b5073ffffffffffffffffffffffffffffffffffffffff8116158015906130b157508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156130e2576130e07faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8261198b565b505b610d937faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f613884565b613114612f65565b610d938282613925565b613126612f65565b611988816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250613988565b61316d612f65565b827f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055815f14610b715773ffffffffffffffffffffffffffffffffffffffff8116613206576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f696e697469616c526563697069656e74000000000000000000000000000000006004820152602401610bd9565b610b715f82846139fa565b5f6fffffffffffffffffffffffffffffffff821115613266576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526080600482015260248101839052604401610bd9565b5090565b80600f81900b81146132b2576040517f327269a70000000000000000000000000000000000000000000000000000000081526080600482015260248101839052604401610bd9565b919050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156132f057505f915060039050826133a0565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613341573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661339757505f9250600191508290506133a0565b92505f91508190505b9450945094915050565b5f8260038111156133bd576133bd614860565b036133c6575050565b60018260038111156133da576133da614860565b03613411576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561342557613425614860565b0361345f576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610bd9565b600382600381111561347357613473614860565b03610d93576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610bd9565b5f610aba8373ffffffffffffffffffffffffffffffffffffffff8416613bc7565b5f610aba8373ffffffffffffffffffffffffffffffffffffffff8416613ca1565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008161351a612216565b80519091501561353257805160209091012092915050565b81548015613541579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613595612267565b8051909150156135ad57805160209091012092915050565b60018201548015613541579392505050565b73ffffffffffffffffffffffffffffffffffffffff821661360e576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bd9565b610d935f8383612f34565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03613681576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610bd9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051613710919061488d565b5f60405180830381855af49150503d805f8114613748576040519150601f19603f3d011682016040523d82523d5f602084013e61374d565b606091505b509150915061375d858383613ced565b95945050505050565b3415611eed576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6137a9838383613d7c565b73ffffffffffffffffffffffffffffffffffffffff8316610b71575f6137ed7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f6138187f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081811115612c65576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610bd9565b5f6138676122df565b5468010000000000000000900460ff16919050565b611eed612f65565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f6138dd845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b61392d612f65565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0361397984826148e7565b5060048101610b0483826148e7565b613990612f65565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026139dc84826148e7565b50600381016139eb83826148e7565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416613a545781816002015f828254613a499190614804565b90915550613b049050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020829052604090205482811015613ad9576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610bd9565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316613b2f576002810180548390039055613b5a565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051613bb991815260200190565b60405180910390a350505050565b5f8181526001830160205260408120548015611aa0575f613be960018361484d565b85549091505f90613bfc9060019061484d565b9050808214613c5b575f865f018281548110613c1a57613c1a614511565b905f5260205f200154905080875f018481548110613c3a57613c3a614511565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080613c6c57613c6c6149fe565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109d6565b5f818152600183016020526040812054613ce657508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109d6565b505f6109d6565b606082613d0257613cfd82613d87565b610aba565b8151158015613d26575073ffffffffffffffffffffffffffffffffffffffff84163b155b15613d75576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610bd9565b5080610aba565b610b718383836139fa565b805115613d975780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828054828255905f5260205f20908101928215613e02579160200282015b82811115613e02578235825591602001919060010190613de7565b50613266929150613e4f565b611eed614a2b565b828054828255905f5260205f20908101928215613e02579160200282015b82811115613e02578251825591602001919060010190613e34565b5b80821115613266575f8155600101613e50565b5f8151808452602084019350602083015f5b82811015613e93578151865260209586019590910190600101613e75565b5093949350505050565b602081525f610aba6020830184613e63565b5f60208284031215613ebf575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610aba575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610aba6020830184613eee565b803573ffffffffffffffffffffffffffffffffffffffff811681146132b2575f5ffd5b5f5f60408385031215613f80575f5ffd5b613f8983613f4c565b946020939093013593505050565b5f5f5f60608486031215613fa9575f5ffd5b613fb284613f4c565b9250613fc060208501613f4c565b929592945050506040919091013590565b5f60208284031215613fe1575f5ffd5b5035919050565b5f5f60408385031215613ff9575f5ffd5b8235915061400960208401613f4c565b90509250929050565b5f5f83601f840112614022575f5ffd5b50813567ffffffffffffffff811115614039575f5ffd5b6020830191508360208260051b8501011115614053575f5ffd5b9250929050565b5f5f6020838503121561406b575f5ffd5b823567ffffffffffffffff811115614081575f5ffd5b61408d85828601614012565b90969095509350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561410d5761410d614099565b604052919050565b5f82601f830112614124575f5ffd5b8135602083015f5f67ffffffffffffffff84111561414457614144614099565b50601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001614177816140c6565b91505082815285838301111561418b575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f604083850312156141b6575f5ffd5b6141bf83613f4c565b9150602083013567ffffffffffffffff8111156141da575f5ffd5b6141e685828601614115565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561423d57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101614209565b509095945050505050565b5f60208284031215614258575f5ffd5b610aba82613f4c565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61429b60e0830189613eee565b82810360408401526142ad8189613eee565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c08401526142e98185613e63565b9a9950505050505050505050565b803560ff811681146132b2575f5ffd5b5f5f5f5f5f5f5f5f610100898b03121561431f575f5ffd5b61432889613f4c565b975061433660208a01613f4c565b9650604089013567ffffffffffffffff811115614351575f5ffd5b61435d8b828c01614115565b965050606089013567ffffffffffffffff811115614379575f5ffd5b6143858b828c01614115565b95505061439460808a016142f7565b935060a089013592506143a960c08a01613f4c565b915060e089013567ffffffffffffffff8111156143c4575f5ffd5b6143d08b828c01614115565b9150509295985092959890939650565b604081525f6143f26040830185613e63565b82810360208401528084518083526020830191506020860192505f5b818110156127cf57835183526020938401939092019160010161440e565b5f5f5f5f5f5f5f60e0888a031215614442575f5ffd5b61444b88613f4c565b965061445960208901613f4c565b95506040880135945060608801359350614475608089016142f7565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156144a3575f5ffd5b6144ac83613f4c565b915061400960208401613f4c565b5f5f5f604084860312156144cc575f5ffd5b833567ffffffffffffffff8111156144e2575f5ffd5b6144ee86828701614012565b90945092505060208401358015158114614506575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061455257607f821691505b602082108103614589577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b604080825284549082018190525f8581526020812090916060840190835b818110156145cb5783548352600193840193602090930192016145ad565b505083810360208501528481527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff851115614604575f5ffd5b8460051b9150818660208301370160200195945050505050565b5f67ffffffffffffffff82111561463757614637614099565b5060051b60200190565b5f82601f830112614650575f5ffd5b815161466361465e8261461e565b6140c6565b8082825260208201915060208360051b860101925085831115614684575f5ffd5b602085015b83811015611418578051835260209283019201614689565b5f5f5f5f608085870312156146b4575f5ffd5b8451602086015190945067ffffffffffffffff8111156146d2575f5ffd5b6146de87828801614641565b935050604085015167ffffffffffffffff8111156146fa575f5ffd5b8501601f8101871361470a575f5ffd5b805161471861465e8261461e565b8082825260208201915060208360051b850101925089831115614739575f5ffd5b6020840193505b8284101561475b578351825260209384019390910190614740565b80955050505050606085015167ffffffffffffffff81111561477b575f5ffd5b61478787828801614641565b91505092959194509250565b5f602082840312156147a3575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156109d6576109d66147d7565b5f7f80000000000000000000000000000000000000000000000000000000000000008203614847576148476147d7565b505f0390565b818103818111156109d6576109d66147d7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f821115610b7157805f5260205f20601f840160051c810160208510156148c85750805b601f840160051c820191505b81811015612c65575f81556001016148d4565b815167ffffffffffffffff81111561490157614901614099565b6149158161490f845461453e565b846148a3565b6020601f821160018114614966575f83156149305750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612c65565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156149b35787850151825560209485019460019092019101614993565b50848210156149ef57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212205f1b021d8a99dbf4995317c47732000dd7ceb18dc315cffe3f71bdccce183f1264736f6c634300081c0033",
}

// ERC20MultiMintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20MultiMintLimited struct {
	abi abi.ABI
}

// NewERC20MultiMintLimited creates a new instance of ERC20MultiMintLimited.
func NewERC20MultiMintLimited() *ERC20MultiMintLimited {
	parsed, err := ERC20MultiMintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20MultiMintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20MultiMintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("allowance", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackAvailableMintCapacities is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3ac4c2d.
//
// Solidity: function availableMintCapacities() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackAvailableMintCapacities() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("availableMintCapacities")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacities is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3ac4c2d.
//
// Solidity: function availableMintCapacities() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAvailableMintCapacities(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("availableMintCapacities", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("balanceOf", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackCap() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("cap", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDecimals() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("decimals", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackEip712Domain() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("eip712Domain", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forgeByIndex", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForgeCount() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forgeCount", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForges() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forges", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("hasRole", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, initialRecipient common.Address, data []byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, initialRecipient, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMaxMintPerPeriods is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xccf6d724.
//
// Solidity: function maxMintPerPeriods() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMaxMintPerPeriods() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("maxMintPerPeriods")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriods is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xccf6d724.
//
// Solidity: function maxMintPerPeriods() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMaxMintPerPeriods(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("maxMintPerPeriods", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackName() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("name", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPeriodConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdd760b9.
//
// Solidity: function periodConfigs() view returns(uint256[], int256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodConfigs() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodConfigs")
	if err != nil {
		panic(err)
	}
	return enc
}

// PeriodConfigsOutput serves as a container for the return parameters of contract
// method PeriodConfigs.
type PeriodConfigsOutput struct {
	Arg0 []*big.Int
	Arg1 []*big.Int
}

// UnpackPeriodConfigs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdd760b9.
//
// Solidity: function periodConfigs() view returns(uint256[], int256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodConfigs(data []byte) (PeriodConfigsOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodConfigs", data)
	outstruct := new(PeriodConfigsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackPeriodStartTimes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0086ced8.
//
// Solidity: function periodStartTimes() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodStartTimes() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodStartTimes")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTimes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0086ced8.
//
// Solidity: function periodStartTimes() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodStartTimes(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodStartTimes", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackProxiableUUID() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("proxiableUUID", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRemainingSupply() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("remainingSupply", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("supportsInterface", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSymbol() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("symbol", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTotalSupply() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("totalSupply", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("transfer", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateMintLimits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37e653ea.
//
// Solidity: function updateMintLimits(uint256[] newLimits) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackUpdateMintLimits(newLimits []*big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("updateMintLimits", newLimits)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20MultiMintLimitedApproval represents a Approval event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedApproval) ContractEventName() string {
	return ERC20MultiMintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20MultiMintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20MultiMintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MultiMintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedForgeAdded) ContractEventName() string {
	return ERC20MultiMintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20MultiMintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedForgeRemoved) ContractEventName() string {
	return ERC20MultiMintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MultiMintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedInitialized represents a Initialized event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedInitialized) ContractEventName() string {
	return ERC20MultiMintLimitedInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackInitializedEvent(log *types.Log) (*ERC20MultiMintLimitedInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedInitialized)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedMintLimitUpdated struct {
	OldLimits []*big.Int
	NewLimits []*big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20MultiMintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20MultiMintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedPeriodStarted) ContractEventName() string {
	return ERC20MultiMintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20MultiMintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20MultiMintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleGranted represents a RoleGranted event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleGranted) ContractEventName() string {
	return ERC20MultiMintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleRevoked) ContractEventName() string {
	return ERC20MultiMintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedTransfer represents a Transfer event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedTransfer) ContractEventName() string {
	return ERC20MultiMintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransferEvent(log *types.Log) (*ERC20MultiMintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedUpgraded represents a Upgraded event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedUpgraded) ContractEventName() string {
	return ERC20MultiMintLimitedUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackUpgradedEvent(log *types.Log) (*ERC20MultiMintLimitedUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20MultiMintLimitedCapTooLow"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20MultiMintLimitedCapTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20MultiMintLimitedInvalidInitialData"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20MultiMintLimitedInvalidInitialDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["PeriodManagerInvalidOffset"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackPeriodManagerInvalidOffsetError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MultiMintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20MultiMintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20MultiMintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20MultiMintLimitedAccessControlBadConfirmation)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20MultiMintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20MultiMintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20MultiMintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20MultiMintLimitedAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAddressEmptyCodeError(raw []byte) (*ERC20MultiMintLimitedAddressEmptyCode, error) {
	out := new(ERC20MultiMintLimitedAddressEmptyCode)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MultiMintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignature)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MultiMintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignatureLength)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MultiMintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignatureS)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20MultiMintLimitedERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20MultiMintLimitedERC1967InvalidImplementation, error) {
	out := new(ERC20MultiMintLimitedERC1967InvalidImplementation)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20MultiMintLimitedERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC1967NonPayableError(raw []byte) (*ERC20MultiMintLimitedERC1967NonPayable, error) {
	out := new(ERC20MultiMintLimitedERC1967NonPayable)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20MultiMintLimitedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20MultiMintLimitedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20MultiMintLimitedERC20CapableERC20ExceededCap)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MultiMintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MultiMintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20MultiMintLimitedERC20InsufficientAllowance)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MultiMintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MultiMintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20MultiMintLimitedERC20InsufficientBalance)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MultiMintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidApprover, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidApprover)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MultiMintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidReceiver)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MultiMintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidSender, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidSender)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MultiMintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidSpender, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidSpender)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20MultiMintLimitedCapTooLow represents a ERC20MultiMintLimited__CapTooLow error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20MultiMintLimitedCapTooLow struct {
	Cap           *big.Int
	InitialSupply *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20MultiMintLimited__CapTooLow(uint256 cap, uint256 initialSupply)
func ERC20MultiMintLimitedERC20MultiMintLimitedCapTooLowErrorID() common.Hash {
	return common.HexToHash("0xce739febd26a983d3aa57569a49b7404b717ec163c232607a10f1ba54db04b6a")
}

// UnpackERC20MultiMintLimitedCapTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20MultiMintLimited__CapTooLow(uint256 cap, uint256 initialSupply)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20MultiMintLimitedCapTooLowError(raw []byte) (*ERC20MultiMintLimitedERC20MultiMintLimitedCapTooLow, error) {
	out := new(ERC20MultiMintLimitedERC20MultiMintLimitedCapTooLow)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20MultiMintLimitedCapTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20MultiMintLimitedInvalidInitialData represents a ERC20MultiMintLimited__InvalidInitialData error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20MultiMintLimitedInvalidInitialData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20MultiMintLimited__InvalidInitialData()
func ERC20MultiMintLimitedERC20MultiMintLimitedInvalidInitialDataErrorID() common.Hash {
	return common.HexToHash("0x94c8879faa01db54cc06d86bfae0defbb953f51dbb218204141f704e14c6a739")
}

// UnpackERC20MultiMintLimitedInvalidInitialDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20MultiMintLimited__InvalidInitialData()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20MultiMintLimitedInvalidInitialDataError(raw []byte) (*ERC20MultiMintLimitedERC20MultiMintLimitedInvalidInitialData, error) {
	out := new(ERC20MultiMintLimitedERC20MultiMintLimitedInvalidInitialData)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20MultiMintLimitedInvalidInitialData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit represents a ERC20PeriodsMintLimit__ExceedsPeriodLimit error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit struct {
	Period    *big.Int
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0xaace2cbe23069663ea57043f9018ae5639c94c4b7537987827b85a4a671f9687")
}

// UnpackERC20PeriodsMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength represents a ERC20PeriodsMintLimit__InvalidLength error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0x3fffd6c1aac13cb22e11380a50735d4cdde9620742e341a28fbb7c0113c01dbe")
}

// UnpackERC20PeriodsMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLengthError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData represents a ERC20PeriodsMintLimit__InvalidLimitData error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0xd82e5a5493b676ee930d3b301ed5da6339364f86d1828cf9de0b2e2baa8068b3")
}

// UnpackERC20PeriodsMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MultiMintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MultiMintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20MultiMintLimitedERC2612ExpiredSignature)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MultiMintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MultiMintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20MultiMintLimitedERC2612InvalidSigner)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedFailedCall represents a FailedCall error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20MultiMintLimitedFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackFailedCallError(raw []byte) (*ERC20MultiMintLimitedFailedCall, error) {
	out := new(ERC20MultiMintLimitedFailedCall)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MultiMintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MultiMintLimitedInvalidAccountNonce, error) {
	out := new(ERC20MultiMintLimitedInvalidAccountNonce)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedInvalidInitialization represents a InvalidInitialization error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20MultiMintLimitedInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackInvalidInitializationError(raw []byte) (*ERC20MultiMintLimitedInvalidInitialization, error) {
	out := new(ERC20MultiMintLimitedInvalidInitialization)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedNotInitializing represents a NotInitializing error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20MultiMintLimitedNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackNotInitializingError(raw []byte) (*ERC20MultiMintLimitedNotInitializing, error) {
	out := new(ERC20MultiMintLimitedNotInitializing)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20MultiMintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20MultiMintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20MultiMintLimitedPeriodManagerInvalidDuration)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedPeriodManagerInvalidOffset represents a PeriodManager__InvalidOffset error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodManagerInvalidOffset struct {
	Timestamp     *big.Int
	OffsetSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func ERC20MultiMintLimitedPeriodManagerInvalidOffsetErrorID() common.Hash {
	return common.HexToHash("0xf41373a6f64e5d167ce7bcc91a19d341958fd006a060f85330cfc8aa087401d6")
}

// UnpackPeriodManagerInvalidOffsetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodManagerInvalidOffsetError(raw []byte) (*ERC20MultiMintLimitedPeriodManagerInvalidOffset, error) {
	out := new(ERC20MultiMintLimitedPeriodManagerInvalidOffset)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidOffset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20MultiMintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20MultiMintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20MultiMintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20MultiMintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20MultiMintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20MultiMintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MultiMintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MultiMintLimitedTokenBaseNullInput, error) {
	out := new(ERC20MultiMintLimitedTokenBaseNullInput)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MultiMintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MultiMintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20MultiMintLimitedTokenBaseOnlyForge)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20MultiMintLimitedUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20MultiMintLimitedUUPSUnauthorizedCallContext, error) {
	out := new(ERC20MultiMintLimitedUUPSUnauthorizedCallContext)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20MultiMintLimitedUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20MultiMintLimitedUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20MultiMintLimitedUUPSUnsupportedProxiableUUID)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
