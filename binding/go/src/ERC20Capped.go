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

// ERC20CappedMetaData contains all meta data concerning the ERC20Capped contract.
var ERC20CappedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20Capped__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Capped__InvalidCapData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20Capped",
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516137486100395f395f81816114e10152818161150a015261172a01526137485ff3fe60806040526004361061020f575f3560e01c806379cc679011610117578063a40283d5116100ac578063d547741f1161007c578063dd62ed3e11610062578063dd62ed3e146106c8578063ec87621c14610738578063fe2df3e81461076b575f5ffd5b8063d547741f14610695578063da0239a6146106b4575f5ffd5b8063a40283d5146105fb578063a9059cbb1461060f578063ad3cb1cc1461062e578063d505accf14610676575f5ffd5b806391d14854116100e757806391d148541461052057806395d89b41146105905780639ca92df9146105a4578063a217fddf146105e8575f5ffd5b806379cc67901461049c5780637ecebe00146104bb57806384b0196e146104da57806387f4535314610501575f5ffd5b8063355274ea116101a75780634f1ef286116101775780635c4e62c41161015d5780635c4e62c41461043d57806370a082311461045e57806379320f781461047d575f5ffd5b80634f1ef2861461041657806352d1902d14610429575f5ffd5b8063355274ea146103915780633644e515146103c457806336568abe146103d857806340c10f19146103f7575f5ffd5b806323b872dd116101e257806323b872dd146102c4578063248a9ca3146102e35780632f2ff15d14610330578063313ce56714610351575f5ffd5b806301ffc9a71461021357806306fdde0314610247578063095ea7b31461026857806318160ddd14610287575b5f5ffd5b34801561021e575f5ffd5b5061023261022d366004612e87565b61078a565b60405190151581526020015b60405180910390f35b348015610252575f5ffd5b5061025b6107e5565b60405161023e9190612f12565b348015610273575f5ffd5b50610232610282366004612f4c565b61089d565b348015610292575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161023e565b3480156102cf575f5ffd5b506102326102de366004612f74565b6108b4565b3480156102ee575f5ffd5b506102b66102fd366004612fae565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561033b575f5ffd5b5061034f61034a366004612fc5565b6108ca565b005b34801561035c575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff909116815260200161023e565b34801561039c575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b300546102b6565b3480156103cf575f5ffd5b506102b6610913565b3480156103e3575f5ffd5b5061034f6103f2366004612fc5565b610921565b348015610402575f5ffd5b5061034f610411366004612f4c565b61097f565b61034f6104243660046130cb565b6109d3565b348015610434575f5ffd5b506102b66109ee565b348015610448575f5ffd5b50610451610a1c565b60405161023e9190613116565b348015610469575f5ffd5b506102b661047836600461316e565b610a47565b348015610488575f5ffd5b5061034f610497366004613197565b610a97565b3480156104a7575f5ffd5b5061034f6104b6366004612f4c565b610d06565b3480156104c6575f5ffd5b506102b66104d536600461316e565b610d1b565b3480156104e5575f5ffd5b506104ee610d25565b60405161023e979695949392919061325f565b34801561050c575f5ffd5b5061023261051b36600461316e565b610e1f565b34801561052b575f5ffd5b5061023261053a366004612fc5565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561059b575f5ffd5b5061025b610e4a565b3480156105af575f5ffd5b506105c36105be366004612fae565b610e9b565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161023e565b3480156105f3575f5ffd5b506102b65f81565b348015610606575f5ffd5b506102b6610ec6565b34801561061a575f5ffd5b50610232610629366004612f4c565b610ef0565b348015610639575f5ffd5b5061025b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610681575f5ffd5b5061034f61069036600461331e565b610efb565b3480156106a0575f5ffd5b5061034f6106af366004612fc5565b6110c3565b3480156106bf575f5ffd5b506102b6611106565b3480156106d3575f5ffd5b506102b66106e2366004613384565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610743575f5ffd5b506102b67faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b348015610776575f5ffd5b5061034f6107853660046133ac565b611175565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f403f67f10000000000000000000000000000000000000000000000000000000014806107df57506107df82611225565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461081b90613432565b80601f016020809104026020016040519081016040528092919081815260200182805461084790613432565b80156108925780601f1061086957610100808354040283529160200191610892565b820191905f5260205f20905b81548152906001019060200180831161087557829003601f168201915b505050505091505090565b5f336108aa81858561122f565b5060019392505050565b5f6108c084848461123c565b90505b9392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546109038161125f565b61090d838361126c565b50505050565b5f61091c61138a565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610970576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61097a8282611393565b505050565b61098833610e1f565b6109c5576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6109cf828261146f565b5050565b6109db6114c9565b6109e4826115cf565b6109cf82826115d9565b5f6109f7611712565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b606061091c7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611781565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f610aa061178d565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610acc5750825b90505f8267ffffffffffffffff166001148015610ae85750303b155b905081158015610af6575080155b15610b2d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610b8e5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b8551602014610bc9576040517f75d507f600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f86806020019051810190610bde9190613483565b9050805f03610c3b576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f636170000000000000000000000000000000000000000000000000000000000060048201526024016109bc565b87811015610c7f576040517f59c338ce00000000000000000000000000000000000000000000000000000000815260048101829052602481018990526044016109bc565b610c88816117b5565b610c968d8d8d8d8d8d61183c565b508315610cf85784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b610d11823383611876565b6109cf828261195d565b5f6107df826119b7565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610d6357506001810154155b610dc9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016109bc565b610dd16119df565b610dd9611a30565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6107df7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611a59565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161081b90613432565b5f6107df7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611a87565b5f61091c7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611a92565b5f6108c38383611a9b565b83421115610f38576040517f62791302000000000000000000000000000000000000000000000000000000008152600481018590526024016109bc565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610faf8c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61101682611aa8565b90505f61102582878787611aef565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110ac576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b1660248201526044016109bc565b6110b78a8a8a61122f565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546110fc8161125f565b61090d8383611393565b5f5f6111307f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f61115b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b905081811161116a575f61116e565b8181035b9250505090565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61119f8161125f565b612e7f826111af57611b1b6111b3565b611b6e5b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b8581101561121c57611214828888848181106111f6576111f661349a565b905060200201602081019061120b919061316e565b8563ffffffff16565b6001016111d8565b50505050505050565b5f6107df82611c30565b61097a8383836001611cc6565b5f33611249858285611876565b611254858585611e30565b506001949350505050565b6112698133611ed9565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611381575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561131d3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506107df565b5f9150506107df565b5f61091c611f7f565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611381575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506107df565b73ffffffffffffffffffffffffffffffffffffffff82166114be576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b6109cf5f8383611ff2565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061159657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661157d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156115cd576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f6109cf8161125f565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561165e575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261165b91810190613483565b60015b6116ac576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016109bc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611708576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b61097a83836120b2565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146115cd576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f6108c383612114565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006107df565b6117bd61216d565b805f03611818576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f636170000000000000000000000000000000000000000000000000000000000060048201526024016109bc565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b61184461216d565b61184e86866121ab565b61185b8685858585612314565b6118658484612408565b61186e8461241a565b505050505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561090d578181101561194f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016109bc565b61090d84848484035f611cc6565b73ffffffffffffffffffffffffffffffffffffffff82166119ac576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b6109cf825f83611ff2565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610a6b565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161081b90613432565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061080a565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156108c3565b5f6108c38383612461565b5f6107df825490565b5f336108aa818585611e30565b5f6107df611ab461138a565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611aff88888888612487565b925092509250611b0f828261257a565b50909695505050505050565b611b25828261267d565b156109cf5760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611bdd576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f72676500000000000000000000000000000000000000000000000000000060048201526024016109bc565b611be7828261269e565b156109cf5760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107df57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146107df565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516611d36576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8416611d85576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115611e29578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051611e2091815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8316611e7f576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8216611ece576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b61097a838383611ff2565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109cf576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016109bc565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611fa96126bf565b611fb161273a565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611ffd83838361278f565b73ffffffffffffffffffffffffffffffffffffffff831661097a575f6120417f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f61206c7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081811115611e29576040517f168a7f6e00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016109bc565b6120bb8261279a565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561210c5761097a8282612868565b6109cf6128e7565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561216157602002820191905f5260205f20905b81548152602001906001019080831161214d575b50505050509050919050565b61217561291f565b6115cd576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6121b361216d565b6121bb61293d565b6121c361293d565b73ffffffffffffffffffffffffffffffffffffffff8216612232576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024016109bc565b61223c5f8361126c565b506122677faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361126c565b5073ffffffffffffffffffffffffffffffffffffffff8116158015906122b957508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156122ea576122e87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8261126c565b505b6109cf7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f612945565b61231c61216d565b83515f03612378576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d650000000000000000000000000000000000000000000000000000000060048201526024016109bc565b82515f036123d4576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c000000000000000000000000000000000000000000000000000060048201526024016109bc565b817f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055805f14611e2957611e29858261146f565b61241061216d565b6109cf82826129e6565b61242261216d565b611269816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250612a49565b5f825f0182815481106124765761247661349a565b905f5260205f200154905092915050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156124c057505f91506003905082612570565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612511573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661256757505f925060019150829050612570565b92505f91508190505b9450945094915050565b5f82600381111561258d5761258d6134c7565b03612596575050565b60018260038111156125aa576125aa6134c7565b036125e1576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156125f5576125f56134c7565b0361262f576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b6003826003811115612643576126436134c7565b036109cf576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b5f6108c38373ffffffffffffffffffffffffffffffffffffffff8416612abb565b5f6108c38373ffffffffffffffffffffffffffffffffffffffff8416612b95565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816126ea6119df565b80519091501561270257805160209091012092915050565b81548015612711579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612765611a30565b80519091501561277d57805160209091012092915050565b60018201548015612711579392505050565b61097a838383612be1565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612802576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109bc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161289191906134f4565b5f60405180830381855af49150503d805f81146128c9576040519150601f19603f3d011682016040523d82523d5f602084013e6128ce565b606091505b50915091506128de858383612dae565b95945050505050565b34156115cd576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61292861178d565b5468010000000000000000900460ff16919050565b6115cd61216d565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f61299e845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b6129ee61216d565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612a3a848261354e565b506004810161090d838261354e565b612a5161216d565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102612a9d848261354e565b5060038101612aac838261354e565b505f8082556001909101555050565b5f8181526001830160205260408120548015611381575f612add600183613692565b85549091505f90612af090600190613692565b9050808214612b4f575f865f018281548110612b0e57612b0e61349a565b905f5260205f200154905080875f018481548110612b2e57612b2e61349a565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612b6057612b606136a5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506107df565b5f818152600183016020526040812054612bda57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556107df565b505f6107df565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416612c3b5781816002015f828254612c3091906136d2565b90915550612ceb9050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020829052604090205482811015612cc0576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101829052604481018490526064016109bc565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316612d16576002810180548390039055612d41565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612da091815260200190565b60405180910390a350505050565b606082612dc357612dbe82612e3d565b6108c3565b8151158015612de7575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612e36576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016109bc565b50806108c3565b805115612e4d5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115cd6136e5565b5f60208284031215612e97575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146108c3575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6108c36020830184612ec6565b803573ffffffffffffffffffffffffffffffffffffffff81168114612f47575f5ffd5b919050565b5f5f60408385031215612f5d575f5ffd5b612f6683612f24565b946020939093013593505050565b5f5f5f60608486031215612f86575f5ffd5b612f8f84612f24565b9250612f9d60208501612f24565b929592945050506040919091013590565b5f60208284031215612fbe575f5ffd5b5035919050565b5f5f60408385031215612fd6575f5ffd5b82359150612fe660208401612f24565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f83011261302b575f5ffd5b8135602083015f5f67ffffffffffffffff84111561304b5761304b612fef565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff8211171561309857613098612fef565b6040528381529050808284018710156130af575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156130dc575f5ffd5b6130e583612f24565b9150602083013567ffffffffffffffff811115613100575f5ffd5b61310c8582860161301c565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561316357835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161312f565b509095945050505050565b5f6020828403121561317e575f5ffd5b6108c382612f24565b803560ff81168114612f47575f5ffd5b5f5f5f5f5f5f5f60e0888a0312156131ad575f5ffd5b6131b688612f24565b96506131c460208901612f24565b9550604088013567ffffffffffffffff8111156131df575f5ffd5b6131eb8a828b0161301c565b955050606088013567ffffffffffffffff811115613207575f5ffd5b6132138a828b0161301c565b94505061322260808901613187565b925060a0880135915060c088013567ffffffffffffffff811115613244575f5ffd5b6132508a828b0161301c565b91505092959891949750929550565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61329960e0830189612ec6565b82810360408401526132ab8189612ec6565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561330d5783518352602093840193909201916001016132ef565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215613334575f5ffd5b61333d88612f24565b965061334b60208901612f24565b9550604088013594506060880135935061336760808901613187565b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215613395575f5ffd5b61339e83612f24565b9150612fe660208401612f24565b5f5f5f604084860312156133be575f5ffd5b833567ffffffffffffffff8111156133d4575f5ffd5b8401601f810186136133e4575f5ffd5b803567ffffffffffffffff8111156133fa575f5ffd5b8660208260051b840101111561340e575f5ffd5b6020918201945092508401358015158114613427575f5ffd5b809150509250925092565b600181811c9082168061344657607f821691505b60208210810361347d577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f60208284031215613493575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f82111561097a57805f5260205f20601f840160051c8101602085101561352f5750805b601f840160051c820191505b81811015611e29575f815560010161353b565b815167ffffffffffffffff81111561356857613568612fef565b61357c816135768454613432565b8461350a565b6020601f8211600181146135cd575f83156135975750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611e29565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561361a57878501518255602094850194600190920191016135fa565b508482101561365657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156107df576107df613665565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b808201808211156107df576107df613665565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220283ea168b47eb848e1e7a5deab0add0fe57c1b0cff886b896877ae1114f882c264736f6c634300081c0033",
}

// ERC20Capped is an auto generated Go binding around an Ethereum contract.
type ERC20Capped struct {
	abi abi.ABI
}

// NewERC20Capped creates a new instance of ERC20Capped.
func NewERC20Capped() *ERC20Capped {
	parsed, err := ERC20CappedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Capped{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Capped) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20Capped.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20Capped *ERC20Capped) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Capped.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20Capped *ERC20Capped) PackMANAGERROLE() []byte {
	enc, err := eRC20Capped.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20Capped *ERC20Capped) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20Capped.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20Capped *ERC20Capped) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20Capped.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (eRC20Capped *ERC20Capped) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("allowance", data)
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
func (eRC20Capped *ERC20Capped) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Capped *ERC20Capped) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("balanceOf", data)
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
func (eRC20Capped *ERC20Capped) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20Capped *ERC20Capped) PackCap() []byte {
	enc, err := eRC20Capped.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("cap", data)
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
func (eRC20Capped *ERC20Capped) PackDecimals() []byte {
	enc, err := eRC20Capped.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Capped *ERC20Capped) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Capped.abi.Unpack("decimals", data)
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
func (eRC20Capped *ERC20Capped) PackEip712Domain() []byte {
	enc, err := eRC20Capped.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Capped *ERC20Capped) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Capped.abi.Unpack("eip712Domain", data)
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
func (eRC20Capped *ERC20Capped) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Capped *ERC20Capped) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("forgeByIndex", data)
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
func (eRC20Capped *ERC20Capped) PackForgeCount() []byte {
	enc, err := eRC20Capped.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("forgeCount", data)
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
func (eRC20Capped *ERC20Capped) PackForges() []byte {
	enc, err := eRC20Capped.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Capped *ERC20Capped) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("forges", data)
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
func (eRC20Capped *ERC20Capped) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("getRoleAdmin", data)
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
func (eRC20Capped *ERC20Capped) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Capped *ERC20Capped) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("hasRole", data)
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
func (eRC20Capped *ERC20Capped) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, data []byte) []byte {
	enc, err := eRC20Capped.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Capped *ERC20Capped) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20Capped *ERC20Capped) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Capped *ERC20Capped) PackName() []byte {
	enc, err := eRC20Capped.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Capped *ERC20Capped) UnpackName(data []byte) (string, error) {
	out, err := eRC20Capped.abi.Unpack("name", data)
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
func (eRC20Capped *ERC20Capped) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("nonces", data)
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
func (eRC20Capped *ERC20Capped) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Capped *ERC20Capped) PackProxiableUUID() []byte {
	enc, err := eRC20Capped.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("proxiableUUID", data)
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
func (eRC20Capped *ERC20Capped) PackRemainingSupply() []byte {
	enc, err := eRC20Capped.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("remainingSupply", data)
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
func (eRC20Capped *ERC20Capped) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20Capped *ERC20Capped) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20Capped *ERC20Capped) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Capped.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Capped *ERC20Capped) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("supportsInterface", data)
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
func (eRC20Capped *ERC20Capped) PackSymbol() []byte {
	enc, err := eRC20Capped.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Capped *ERC20Capped) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Capped.abi.Unpack("symbol", data)
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
func (eRC20Capped *ERC20Capped) PackTotalSupply() []byte {
	enc, err := eRC20Capped.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("totalSupply", data)
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
func (eRC20Capped *ERC20Capped) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("transfer", data)
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
func (eRC20Capped *ERC20Capped) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("transferFrom", data)
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
func (eRC20Capped *ERC20Capped) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20Capped.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20CappedApproval represents a Approval event raised by the ERC20Capped contract.
type ERC20CappedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20CappedApproval) ContractEventName() string {
	return ERC20CappedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Capped *ERC20Capped) UnpackApprovalEvent(log *types.Log) (*ERC20CappedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedApproval)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Capped contract.
type ERC20CappedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedEIP712DomainChanged) ContractEventName() string {
	return ERC20CappedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Capped *ERC20Capped) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20CappedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedForgeAdded represents a ForgeAdded event raised by the ERC20Capped contract.
type ERC20CappedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20CappedForgeAdded) ContractEventName() string {
	return ERC20CappedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Capped *ERC20Capped) UnpackForgeAddedEvent(log *types.Log) (*ERC20CappedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedForgeRemoved represents a ForgeRemoved event raised by the ERC20Capped contract.
type ERC20CappedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20CappedForgeRemoved) ContractEventName() string {
	return ERC20CappedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Capped *ERC20Capped) UnpackForgeRemovedEvent(log *types.Log) (*ERC20CappedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedInitialized represents a Initialized event raised by the ERC20Capped contract.
type ERC20CappedInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20CappedInitialized) ContractEventName() string {
	return ERC20CappedInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20Capped *ERC20Capped) UnpackInitializedEvent(log *types.Log) (*ERC20CappedInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedInitialized)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Capped contract.
type ERC20CappedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleAdminChanged) ContractEventName() string {
	return ERC20CappedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20Capped *ERC20Capped) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20CappedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleGranted represents a RoleGranted event raised by the ERC20Capped contract.
type ERC20CappedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleGranted) ContractEventName() string {
	return ERC20CappedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Capped *ERC20Capped) UnpackRoleGrantedEvent(log *types.Log) (*ERC20CappedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleRevoked represents a RoleRevoked event raised by the ERC20Capped contract.
type ERC20CappedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleRevoked) ContractEventName() string {
	return ERC20CappedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Capped *ERC20Capped) UnpackRoleRevokedEvent(log *types.Log) (*ERC20CappedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedTransfer represents a Transfer event raised by the ERC20Capped contract.
type ERC20CappedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20CappedTransfer) ContractEventName() string {
	return ERC20CappedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Capped *ERC20Capped) UnpackTransferEvent(log *types.Log) (*ERC20CappedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedUpgraded represents a Upgraded event raised by the ERC20Capped contract.
type ERC20CappedUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20CappedUpgraded) ContractEventName() string {
	return ERC20CappedUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20Capped *ERC20Capped) UnpackUpgradedEvent(log *types.Log) (*ERC20CappedUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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
func (eRC20Capped *ERC20Capped) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20CappedCapTooLow"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20CappedCapTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20CappedInvalidCapData"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20CappedInvalidCapDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20CappedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20Capped contract.
type ERC20CappedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20CappedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20Capped *ERC20Capped) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20CappedAccessControlBadConfirmation, error) {
	out := new(ERC20CappedAccessControlBadConfirmation)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20Capped contract.
type ERC20CappedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20CappedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20Capped *ERC20Capped) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20CappedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20CappedAccessControlUnauthorizedAccount)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20Capped contract.
type ERC20CappedAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20CappedAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20Capped *ERC20Capped) UnpackAddressEmptyCodeError(raw []byte) (*ERC20CappedAddressEmptyCode, error) {
	out := new(ERC20CappedAddressEmptyCode)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20CappedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20CappedECDSAInvalidSignature, error) {
	out := new(ERC20CappedECDSAInvalidSignature)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20CappedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20CappedECDSAInvalidSignatureLength, error) {
	out := new(ERC20CappedECDSAInvalidSignatureLength)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20CappedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20CappedECDSAInvalidSignatureS, error) {
	out := new(ERC20CappedECDSAInvalidSignatureS)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20Capped contract.
type ERC20CappedERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20CappedERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20Capped *ERC20Capped) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20CappedERC1967InvalidImplementation, error) {
	out := new(ERC20CappedERC1967InvalidImplementation)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20Capped contract.
type ERC20CappedERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20CappedERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20Capped *ERC20Capped) UnpackERC1967NonPayableError(raw []byte) (*ERC20CappedERC1967NonPayable, error) {
	out := new(ERC20CappedERC1967NonPayable)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20Capped contract.
type ERC20CappedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20CappedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20Capped *ERC20Capped) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20CappedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20CappedERC20CapableERC20ExceededCap)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20CappedCapTooLow represents a ERC20Capped__CapTooLow error raised by the ERC20Capped contract.
type ERC20CappedERC20CappedCapTooLow struct {
	Cap           *big.Int
	InitialSupply *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capped__CapTooLow(uint256 cap, uint256 initialSupply)
func ERC20CappedERC20CappedCapTooLowErrorID() common.Hash {
	return common.HexToHash("0x59c338ce8a81b32c41082da628338f2729ddbb619bec9e2a204954ae57d12112")
}

// UnpackERC20CappedCapTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capped__CapTooLow(uint256 cap, uint256 initialSupply)
func (eRC20Capped *ERC20Capped) UnpackERC20CappedCapTooLowError(raw []byte) (*ERC20CappedERC20CappedCapTooLow, error) {
	out := new(ERC20CappedERC20CappedCapTooLow)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20CappedCapTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20CappedInvalidCapData represents a ERC20Capped__InvalidCapData error raised by the ERC20Capped contract.
type ERC20CappedERC20CappedInvalidCapData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capped__InvalidCapData()
func ERC20CappedERC20CappedInvalidCapDataErrorID() common.Hash {
	return common.HexToHash("0x75d507f6fcfe3c719de318414fef27edb869ff383efd4affd57403c4f9f3f359")
}

// UnpackERC20CappedInvalidCapDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capped__InvalidCapData()
func (eRC20Capped *ERC20Capped) UnpackERC20CappedInvalidCapDataError(raw []byte) (*ERC20CappedERC20CappedInvalidCapData, error) {
	out := new(ERC20CappedERC20CappedInvalidCapData)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20CappedInvalidCapData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Capped contract.
type ERC20CappedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20CappedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Capped *ERC20Capped) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20CappedERC20InsufficientAllowance, error) {
	out := new(ERC20CappedERC20InsufficientAllowance)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Capped contract.
type ERC20CappedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20CappedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Capped *ERC20Capped) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20CappedERC20InsufficientBalance, error) {
	out := new(ERC20CappedERC20InsufficientBalance)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20CappedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidApproverError(raw []byte) (*ERC20CappedERC20InvalidApprover, error) {
	out := new(ERC20CappedERC20InvalidApprover)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20CappedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20CappedERC20InvalidReceiver, error) {
	out := new(ERC20CappedERC20InvalidReceiver)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20CappedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidSenderError(raw []byte) (*ERC20CappedERC20InvalidSender, error) {
	out := new(ERC20CappedERC20InvalidSender)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20CappedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20CappedERC20InvalidSpender, error) {
	out := new(ERC20CappedERC20InvalidSpender)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Capped contract.
type ERC20CappedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20CappedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Capped *ERC20Capped) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20CappedERC2612ExpiredSignature, error) {
	out := new(ERC20CappedERC2612ExpiredSignature)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Capped contract.
type ERC20CappedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20CappedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Capped *ERC20Capped) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20CappedERC2612InvalidSigner, error) {
	out := new(ERC20CappedERC2612InvalidSigner)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedFailedCall represents a FailedCall error raised by the ERC20Capped contract.
type ERC20CappedFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20CappedFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20Capped *ERC20Capped) UnpackFailedCallError(raw []byte) (*ERC20CappedFailedCall, error) {
	out := new(ERC20CappedFailedCall)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Capped contract.
type ERC20CappedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20CappedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Capped *ERC20Capped) UnpackInvalidAccountNonceError(raw []byte) (*ERC20CappedInvalidAccountNonce, error) {
	out := new(ERC20CappedInvalidAccountNonce)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedInvalidInitialization represents a InvalidInitialization error raised by the ERC20Capped contract.
type ERC20CappedInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20CappedInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20Capped *ERC20Capped) UnpackInvalidInitializationError(raw []byte) (*ERC20CappedInvalidInitialization, error) {
	out := new(ERC20CappedInvalidInitialization)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedNotInitializing represents a NotInitializing error raised by the ERC20Capped contract.
type ERC20CappedNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20CappedNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20Capped *ERC20Capped) UnpackNotInitializingError(raw []byte) (*ERC20CappedNotInitializing, error) {
	out := new(ERC20CappedNotInitializing)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Capped contract.
type ERC20CappedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20CappedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Capped *ERC20Capped) UnpackTokenBaseNullInputError(raw []byte) (*ERC20CappedTokenBaseNullInput, error) {
	out := new(ERC20CappedTokenBaseNullInput)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Capped contract.
type ERC20CappedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20CappedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Capped *ERC20Capped) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20CappedTokenBaseOnlyForge, error) {
	out := new(ERC20CappedTokenBaseOnlyForge)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20Capped contract.
type ERC20CappedUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20CappedUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20Capped *ERC20Capped) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20CappedUUPSUnauthorizedCallContext, error) {
	out := new(ERC20CappedUUPSUnauthorizedCallContext)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20Capped contract.
type ERC20CappedUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20CappedUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20Capped *ERC20Capped) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20CappedUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20CappedUUPSUnsupportedProxiableUUID)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
