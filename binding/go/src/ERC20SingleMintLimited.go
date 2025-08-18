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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"updateMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimits\",\"type\":\"uint256\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20SingleMintLimited__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20SingleMintLimited__InvalidInitialData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20SingleMintLimited",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051613bb56100395f395f8181611791015281816117ba01526119da0152613bb55ff3fe608060405260043610610291575f3560e01c806370a0823111610165578063a40283d5116100c6578063da0239a61161007c578063e01d55c511610062578063e01d55c514610848578063ec87621c14610867578063fe2df3e81461089a575f5ffd5b8063da0239a6146107c4578063dd62ed3e146107d8575f5ffd5b8063ad3cb1cc116100ac578063ad3cb1cc1461073e578063d505accf14610786578063d547741f146107a5575f5ffd5b8063a40283d51461070b578063a9059cbb1461071f575f5ffd5b806387f453531161011b57806395d89b411161010157806395d89b41146106a05780639ca92df9146106b4578063a217fddf146106f8575f5ffd5b806387f453531461061157806391d1485414610630575f5ffd5b806379cc67901161014b57806379cc6790146105ac5780637ecebe00146105cb57806384b0196e146105ea575f5ffd5b806370a082311461056e57806379320f781461058d575f5ffd5b8063355274ea1161020f5780634834c649116101c55780634f1ef286116101ab5780634f1ef2861461052657806352d1902d146105395780635c4e62c41461054d575f5ffd5b80634834c649146104df5780634836695b14610512575f5ffd5b806336568abe116101f557806336568abe1461048d57806340c10f19146104ac578063436ba626146104cb575f5ffd5b8063355274ea146104465780633644e51514610479575f5ffd5b806323b872dd116102645780632d9ed80d1161024a5780632d9ed80d146103b25780632f2ff15d146103e5578063313ce56714610406575f5ffd5b806323b872dd14610346578063248a9ca314610365575f5ffd5b806301ffc9a71461029557806306fdde03146102c9578063095ea7b3146102ea57806318160ddd14610309575b5f5ffd5b3480156102a0575f5ffd5b506102b46102af36600461329c565b6108b9565b60405190151581526020015b60405180910390f35b3480156102d4575f5ffd5b506102dd610914565b6040516102c09190613327565b3480156102f5575f5ffd5b506102b4610304366004613361565b6109cc565b348015610314575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016102c0565b348015610351575f5ffd5b506102b4610360366004613389565b6109e3565b348015610370575f5ffd5b5061033861037f3660046133c3565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156103bd575f5ffd5b507f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0254610338565b3480156103f0575f5ffd5b506104046103ff3660046133da565b6109f9565b005b348015610411575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff90911681526020016102c0565b348015610451575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30054610338565b348015610484575f5ffd5b50610338610a42565b348015610498575f5ffd5b506104046104a73660046133da565b610a50565b3480156104b7575f5ffd5b506104046104c6366004613361565b610aae565b3480156104d6575f5ffd5b50610338610abc565b3480156104ea575f5ffd5b507f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0054610338565b34801561051d575f5ffd5b50610338610b02565b6104046105343660046134e0565b610b3f565b348015610544575f5ffd5b50610338610b5a565b348015610558575f5ffd5b50610561610b88565b6040516102c0919061352b565b348015610579575f5ffd5b50610338610588366004613583565b610bb3565b348015610598575f5ffd5b506104046105a73660046135ac565b610c03565b3480156105b7575f5ffd5b506104046105c6366004613361565b610e2e565b3480156105d6575f5ffd5b506103386105e5366004613583565b610e43565b3480156105f5575f5ffd5b506105fe610e4d565b6040516102c09796959493929190613674565b34801561061c575f5ffd5b506102b461062b366004613583565b610f47565b34801561063b575f5ffd5b506102b461064a3660046133da565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156106ab575f5ffd5b506102dd610f72565b3480156106bf575f5ffd5b506106d36106ce3660046133c3565b610fc3565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c0565b348015610703575f5ffd5b506103385f81565b348015610716575f5ffd5b50610338610fee565b34801561072a575f5ffd5b506102b4610739366004613361565b611018565b348015610749575f5ffd5b506102dd6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610791575f5ffd5b506104046107a0366004613733565b611023565b3480156107b0575f5ffd5b506104046107bf3660046133da565b6111eb565b3480156107cf575f5ffd5b5061033861122e565b3480156107e3575f5ffd5b506103386107f2366004613799565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610853575f5ffd5b506104046108623660046133c3565b61129d565b348015610872575f5ffd5b506103387faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b3480156108a5575f5ffd5b506104046108b43660046137c1565b611389565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f403f67f100000000000000000000000000000000000000000000000000000000148061090e575061090e82611439565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461094a90613847565b80601f016020809104026020016040519081016040528092919081815260200182805461097690613847565b80156109c15780601f10610998576101008083540402835291602001916109c1565b820191905f5260205f20905b8154815290600101906020018083116109a457829003601f168201915b505050505091505090565b5f336109d9818585611443565b5060019392505050565b5f6109ef848484611450565b90505b9392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610a3281611473565b610a3c8383611480565b50505050565b5f610a4b61159e565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610a9f576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aa982826115a7565b505050565b610ab88282611683565b5050565b5f7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00610ae6610b02565b816001015403610af95760030154919050565b60020154919050565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00545f9043818181610b3657610b36613898565b06900392915050565b610b47611779565b610b508261187f565b610ab88282611889565b5f610b636119c2565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6060610a4b7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611a31565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f610c0c611a3d565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610c385750825b90505f8267ffffffffffffffff166001148015610c545750303b155b905081158015610c62575080155b15610c99576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610cfa5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b8551606014610d35576040517f5431945d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f88806020019051810190610d4c91906138c5565b92509250925089831015610d9b576040517fc013187a00000000000000000000000000000000000000000000000000000000815260048101849052602481018b90526044015b60405180910390fd5b610da58282611a65565b610dae83611b1a565b610dbc8f8f8f8f8f8f611ba1565b5050508315610e205784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b610e39823383611bdb565b610ab88282611cc2565b5f61090e82611d1c565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610e8b57506001810154155b610ef1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610d92565b610ef9611d44565b610f01611d95565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f61090e7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611dbe565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161094a90613847565b5f61090e7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611dec565b5f610a4b7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611df7565b5f6109f28383611e00565b83421115611060576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610d92565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886110d78c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61113e82611e0d565b90505f61114d82878787611e54565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146111d4576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610d92565b6111df8a8a8a611443565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461122481611473565b610a3c83836115a7565b5f5f6112587f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f6112837f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b9050818111611292575f611296565b8181035b9250505090565b5f6112a781611473565b815f03611302576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69740000000000000000000000000000000000000000000000006004820152602401610d92565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0254604080519182526020820184905280517f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00927f864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e48292908290030190a16002019190915550565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6113b381611473565b613294826113c357611e806113c7565b611ed35b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b85811015611430576114288288888481811061140a5761140a6138f0565b905060200201602081019061141f9190613583565b8563ffffffff16565b6001016113ec565b50505050505050565b5f61090e82611f95565b610aa9838383600161202b565b5f3361145d858285611bdb565b611468858585612195565b506001949350505050565b61147d813361223e565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611595575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556115313390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061090e565b5f91505061090e565b5f610a4b6122e4565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611595575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061090e565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de03547f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00905f6116d0610b02565b9050808360010154146117225760018301819055600283015460405181815290925081907ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d19060200160405180910390a25b5082811015611767576040517f156d32c50000000000000000000000000000000000000000000000000000000081526004810184905260248101829052604401610d92565b8281036003830155610a3c8484612357565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061184657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661182d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561187d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610ab881611473565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561190e575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261190b9181019061391d565b60015b61195c576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610d92565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146119b8576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610d92565b610aa983836123a2565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461187d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f6109f283612404565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061090e565b611a6d61245d565b811580611a78575080155b15611ad1576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6c696d6974206f7220706572696f6400000000000000000000000000000000006004820152602401610d92565b7f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00919091557f02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de0255565b611b2261245d565b805f03611b7d576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f63617000000000000000000000000000000000000000000000000000000000006004820152602401610d92565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b611ba961245d565b611bb3868661249b565b611bc08685858585612604565b611bca84846126f8565b611bd38461270a565b505050505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610a3c5781811015611cb4576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610d92565b610a3c84848484035f61202b565b73ffffffffffffffffffffffffffffffffffffffff8216611d11576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b610ab8825f83612751565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610bd7565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161094a90613847565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610939565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156109f2565b5f6109f2838361275c565b5f61090e825490565b5f336109d9818585612195565b5f61090e611e1961159e565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611e6488888888612782565b925092509250611e748282612875565b50909695505050505050565b611e8a8282612978565b15610ab85760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611f42576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610d92565b611f4c8282612999565b15610ab85760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061090e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461090e565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff851661209b576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b73ffffffffffffffffffffffffffffffffffffffff84166120ea576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b73ffffffffffffffffffffffffffffffffffffffff8086165f9081526001830160209081526040808320938816835292905220839055811561218e578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161218591815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff83166121e4576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b73ffffffffffffffffffffffffffffffffffffffff8216612233576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b610aa9838383612751565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610ab8576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610d92565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61230e6129ba565b612316612a35565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b61236033610f47565b612398576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610d92565b610ab88282612a8a565b6123ab82612ae4565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156123fc57610aa98282612bb2565b610ab8612c31565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561245157602002820191905f5260205f20905b81548152602001906001019080831161243d575b50505050509050919050565b612465612c69565b61187d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124a361245d565b6124ab612c87565b6124b3612c87565b73ffffffffffffffffffffffffffffffffffffffff8216612522576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610d92565b61252c5f83611480565b506125577faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c83611480565b5073ffffffffffffffffffffffffffffffffffffffff8116158015906125a957508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156125da576125d87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c82611480565b505b610ab87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f612c8f565b61260c61245d565b83515f03612668576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610d92565b82515f036126c4576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610d92565b817f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055805f1461218e5761218e8582612a8a565b61270061245d565b610ab88282612d30565b61271261245d565b61147d816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250612d93565b610aa9838383612e05565b5f825f018281548110612771576127716138f0565b905f5260205f200154905092915050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156127bb57505f9150600390508261286b565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561280c573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661286257505f92506001915082905061286b565b92505f91508190505b9450945094915050565b5f82600381111561288857612888613934565b03612891575050565b60018260038111156128a5576128a5613934565b036128dc576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156128f0576128f0613934565b0361292a576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610d92565b600382600381111561293e5761293e613934565b03610ab8576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610d92565b5f6109f28373ffffffffffffffffffffffffffffffffffffffff8416612ec5565b5f6109f28373ffffffffffffffffffffffffffffffffffffffff8416612f9f565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816129e5611d44565b8051909150156129fd57805160209091012092915050565b81548015612a0c579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612a60611d95565b805190915015612a7857805160209091012092915050565b60018201548015612a0c579392505050565b73ffffffffffffffffffffffffffffffffffffffff8216612ad9576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d92565b610ab85f8383612751565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612b4c576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610d92565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612bdb9190613961565b5f60405180830381855af49150503d805f8114612c13576040519150601f19603f3d011682016040523d82523d5f602084013e612c18565b606091505b5091509150612c28858383612feb565b95945050505050565b341561187d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f612c72611a3d565b5468010000000000000000900460ff16919050565b61187d61245d565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f612ce8845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b612d3861245d565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612d8484826139bb565b5060048101610a3c83826139bb565b612d9b61245d565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102612de784826139bb565b5060038101612df683826139bb565b505f8082556001909101555050565b612e1083838361307a565b73ffffffffffffffffffffffffffffffffffffffff8316610aa9575f612e547f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f612e7f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90508181111561218e576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610d92565b5f8181526001830160205260408120548015611595575f612ee7600183613aff565b85549091505f90612efa90600190613aff565b9050808214612f59575f865f018281548110612f1857612f186138f0565b905f5260205f200154905080875f018481548110612f3857612f386138f0565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612f6a57612f6a613b12565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061090e565b5f818152600183016020526040812054612fe457508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561090e565b505f61090e565b60608261300057612ffb82613085565b6109f2565b8151158015613024575073ffffffffffffffffffffffffffffffffffffffff84163b155b15613073576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610d92565b50806109f2565b610aa98383836130c7565b8051156130955780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff84166131215781816002015f8282546131169190613b3f565b909155506131d19050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156131a6576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610d92565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff83166131fc576002810180548390039055613227565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161328691815260200190565b60405180910390a350505050565b61187d613b52565b5f602082840312156132ac575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109f2575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109f260208301846132db565b803573ffffffffffffffffffffffffffffffffffffffff8116811461335c575f5ffd5b919050565b5f5f60408385031215613372575f5ffd5b61337b83613339565b946020939093013593505050565b5f5f5f6060848603121561339b575f5ffd5b6133a484613339565b92506133b260208501613339565b929592945050506040919091013590565b5f602082840312156133d3575f5ffd5b5035919050565b5f5f604083850312156133eb575f5ffd5b823591506133fb60208401613339565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112613440575f5ffd5b8135602083015f5f67ffffffffffffffff84111561346057613460613404565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156134ad576134ad613404565b6040528381529050808284018710156134c4575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156134f1575f5ffd5b6134fa83613339565b9150602083013567ffffffffffffffff811115613515575f5ffd5b61352185828601613431565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561357857835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613544565b509095945050505050565b5f60208284031215613593575f5ffd5b6109f282613339565b803560ff8116811461335c575f5ffd5b5f5f5f5f5f5f5f60e0888a0312156135c2575f5ffd5b6135cb88613339565b96506135d960208901613339565b9550604088013567ffffffffffffffff8111156135f4575f5ffd5b6136008a828b01613431565b955050606088013567ffffffffffffffff81111561361c575f5ffd5b6136288a828b01613431565b9450506136376080890161359c565b925060a0880135915060c088013567ffffffffffffffff811115613659575f5ffd5b6136658a828b01613431565b91505092959891949750929550565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6136ae60e08301896132db565b82810360408401526136c081896132db565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015613722578351835260209384019390920191600101613704565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215613749575f5ffd5b61375288613339565b965061376060208901613339565b9550604088013594506060880135935061377c6080890161359c565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156137aa575f5ffd5b6137b383613339565b91506133fb60208401613339565b5f5f5f604084860312156137d3575f5ffd5b833567ffffffffffffffff8111156137e9575f5ffd5b8401601f810186136137f9575f5ffd5b803567ffffffffffffffff81111561380f575f5ffd5b8660208260051b8401011115613823575f5ffd5b602091820194509250840135801515811461383c575f5ffd5b809150509250925092565b600181811c9082168061385b57607f821691505b602082108103613892577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f5f5f606084860312156138d7575f5ffd5b5050815160208301516040909301519094929350919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020828403121561392d575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f821115610aa957805f5260205f20601f840160051c8101602085101561399c5750805b601f840160051c820191505b8181101561218e575f81556001016139a8565b815167ffffffffffffffff8111156139d5576139d5613404565b6139e9816139e38454613847565b84613977565b6020601f821160018114613a3a575f8315613a045750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561218e565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015613a875787850151825560209485019460019092019101613a67565b5084821015613ac357868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561090e5761090e613ad2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b8082018082111561090e5761090e613ad2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212204573a1d2e662e072ae6f9aa73c12acd7ef67cca307c5b61a1da74ab885a412e664736f6c634300081c0033",
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
// the contract method with ID 0x79320f78.
//
// Solidity: function initialize(address _owner, address _manager, string _name, string _symbol, uint8 _decimals, uint256 _initialSupply, bytes _data) returns()
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, data []byte) []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, data)
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

// PackPeriodBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4834c649.
//
// Solidity: function periodBlock() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackPeriodBlock() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("periodBlock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodBlock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4834c649.
//
// Solidity: function periodBlock() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodBlock(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("periodBlock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPeriodStartBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4836695b.
//
// Solidity: function periodStartBlock() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) PackPeriodStartBlock() []byte {
	enc, err := eRC20SingleMintLimited.abi.Pack("periodStartBlock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartBlock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4836695b.
//
// Solidity: function periodStartBlock() view returns(uint256)
func (eRC20SingleMintLimited *ERC20SingleMintLimited) UnpackPeriodStartBlock(data []byte) (*big.Int, error) {
	out, err := eRC20SingleMintLimited.abi.Unpack("periodStartBlock", data)
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
	PeriodStartBlock  *big.Int
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
// Solidity: event PeriodStarted(uint256 indexed periodStartBlock, uint256 availableCapacity)
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
