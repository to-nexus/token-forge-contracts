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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20Capped__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Capped__InvalidCapData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20Capped",
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516137c16100395f395f81816114e30152818161150c015261172c01526137c15ff3fe60806040526004361061020f575f3560e01c80637ecebe0011610117578063a9059cbb116100ac578063d547741f1161007c578063dd62ed3e11610062578063dd62ed3e146106c8578063ec87621c14610738578063fe2df3e81461076b575f5ffd5b8063d547741f14610695578063da0239a6146106b4575f5ffd5b8063a9059cbb146105f0578063ad3cb1cc1461060f578063b019322714610657578063d505accf14610676575f5ffd5b806395d89b41116100e757806395d89b41146105715780639ca92df914610585578063a217fddf146105c9578063a40283d5146105dc575f5ffd5b80637ecebe001461049c57806384b0196e146104bb57806387f45353146104e257806391d1485414610501575f5ffd5b8063355274ea116101a75780634f1ef286116101775780635c4e62c41161015d5780635c4e62c41461043d57806370a082311461045e57806379cc67901461047d575f5ffd5b80634f1ef2861461041657806352d1902d14610429575f5ffd5b8063355274ea146103915780633644e515146103c457806336568abe146103d857806340c10f19146103f7575f5ffd5b806323b872dd116101e257806323b872dd146102c4578063248a9ca3146102e35780632f2ff15d14610330578063313ce56714610351575f5ffd5b806301ffc9a71461021357806306fdde0314610247578063095ea7b31461026857806318160ddd14610287575b5f5ffd5b34801561021e575f5ffd5b5061023261022d366004612eef565b61078a565b60405190151581526020015b60405180910390f35b348015610252575f5ffd5b5061025b6107e5565b60405161023e9190612f7a565b348015610273575f5ffd5b50610232610282366004612fb4565b61089d565b348015610292575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161023e565b3480156102cf575f5ffd5b506102326102de366004612fdc565b6108b4565b3480156102ee575f5ffd5b506102b66102fd366004613016565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561033b575f5ffd5b5061034f61034a36600461302d565b6108ca565b005b34801561035c575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff909116815260200161023e565b34801561039c575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b300546102b6565b3480156103cf575f5ffd5b506102b6610913565b3480156103e3575f5ffd5b5061034f6103f236600461302d565b610921565b348015610402575f5ffd5b5061034f610411366004612fb4565b61097f565b61034f610424366004613133565b6109d3565b348015610434575f5ffd5b506102b66109ee565b348015610448575f5ffd5b50610451610a1c565b60405161023e919061317e565b348015610469575f5ffd5b506102b66104783660046131d6565b610a47565b348015610488575f5ffd5b5061034f610497366004612fb4565b610a97565b3480156104a7575f5ffd5b506102b66104b63660046131d6565b610aac565b3480156104c6575f5ffd5b506104cf610ab6565b60405161023e97969594939291906131ef565b3480156104ed575f5ffd5b506102326104fc3660046131d6565b610bb0565b34801561050c575f5ffd5b5061023261051b36600461302d565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561057c575f5ffd5b5061025b610bdb565b348015610590575f5ffd5b506105a461059f366004613016565b610c2c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161023e565b3480156105d4575f5ffd5b506102b65f81565b3480156105e7575f5ffd5b506102b6610c57565b3480156105fb575f5ffd5b5061023261060a366004612fb4565b610c81565b34801561061a575f5ffd5b5061025b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610662575f5ffd5b5061034f6106713660046132be565b610c8c565b348015610681575f5ffd5b5061034f610690366004613397565b610efd565b3480156106a0575f5ffd5b5061034f6106af36600461302d565b6110c5565b3480156106bf575f5ffd5b506102b6611108565b3480156106d3575f5ffd5b506102b66106e23660046133fd565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610743575f5ffd5b506102b67faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b348015610776575f5ffd5b5061034f610785366004613425565b611177565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f89145aae0000000000000000000000000000000000000000000000000000000014806107df57506107df82611227565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461081b906134ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610847906134ab565b80156108925780601f1061086957610100808354040283529160200191610892565b820191905f5260205f20905b81548152906001019060200180831161087557829003601f168201915b505050505091505090565b5f336108aa818585611231565b5060019392505050565b5f6108c084848461123e565b90505b9392505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461090381611261565b61090d838361126e565b50505050565b5f61091c61138c565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610970576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61097a8282611395565b505050565b61098833610bb0565b6109c5576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6109cf8282611471565b5050565b6109db6114cb565b6109e4826115d1565b6109cf82826115db565b5f6109f7611714565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b606061091c7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611783565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b610aa282338361178f565b6109cf8282611876565b5f6107df826118d0565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610af457506001810154155b610b5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016109bc565b610b626118f8565b610b6a611949565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6107df7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611972565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161081b906134ab565b5f6107df7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400836119a0565b5f61091c7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474006119ab565b5f6108c383836119b4565b5f610c956119c1565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610cc15750825b90505f8267ffffffffffffffff166001148015610cdd5750303b155b905081158015610ceb575080155b15610d22576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610d835784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610d928d8d8d8d8d8d8d6119e9565b8551602014610dcd576040517f75d507f600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f86806020019051810190610de291906134fc565b9050805f03610e3f576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f636170000000000000000000000000000000000000000000000000000000000060048201526024016109bc565b88811015610e83576040517f59c338ce00000000000000000000000000000000000000000000000000000000815260048101829052602481018a90526044016109bc565b610e8c81611ad1565b508315610eee5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b83421115610f3a576040517f62791302000000000000000000000000000000000000000000000000000000008152600481018590526024016109bc565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610fb18c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61101882611b58565b90505f61102782878787611b9f565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110ae576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b1660248201526044016109bc565b6110b98a8a8a611231565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546110fe81611261565b61090d8383611395565b5f5f6111327f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f61115d7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b905081811161116c575f611170565b8181035b9250505090565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6111a181611261565b612ee7826111b157611bcb6111b5565b611c1e5b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b8581101561121e57611216828888848181106111f8576111f8613513565b905060200201602081019061120d91906131d6565b8563ffffffff16565b6001016111da565b50505050505050565b5f6107df82611ce0565b61097a8383836001611d76565b5f3361124b85828561178f565b611256858585611ee0565b506001949350505050565b61126b8133611f89565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611383575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561131f3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506107df565b5f9150506107df565b5f61091c61202f565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611383575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506107df565b73ffffffffffffffffffffffffffffffffffffffff82166114c0576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b6109cf5f83836120a2565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061159857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661157f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156115cf576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f6109cf81611261565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611660575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261165d918101906134fc565b60015b6116ae576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016109bc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461170a576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b61097a8383612162565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146115cf576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f6108c3836121c4565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561090d5781811015611868576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016109bc565b61090d84848484035f611d76565b73ffffffffffffffffffffffffffffffffffffffff82166118c5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b6109cf825f836120a2565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610a6b565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161081b906134ab565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061080a565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156108c3565b5f6108c3838361221d565b5f6107df825490565b5f336108aa818585611ee0565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006107df565b6119f1612243565b84515f03611a4d576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d650000000000000000000000000000000000000000000000000000000060048201526024016109bc565b83515f03611aa9576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c000000000000000000000000000000000000000000000000000060048201526024016109bc565b611ab38787612281565b611abd85856123ea565b611ac6856123fc565b61121e838383612443565b611ad9612243565b805f03611b34576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f636170000000000000000000000000000000000000000000000000000000000060048201526024016109bc565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b5f6107df611b6461138c565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611baf888888886124ef565b925092509250611bbf82826125e2565b50909695505050505050565b611bd582826126e5565b156109cf5760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611c8d576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f72676500000000000000000000000000000000000000000000000000000060048201526024016109bc565b611c978282612706565b156109cf5760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107df57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146107df565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516611de6576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8416611e35576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115611ed9578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051611ed091815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8316611f2f576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b73ffffffffffffffffffffffffffffffffffffffff8216611f7e576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016109bc565b61097a8383836120a2565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109cf576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016109bc565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612059612727565b6120616127a2565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6120ad8383836127f7565b73ffffffffffffffffffffffffffffffffffffffff831661097a575f6120f17f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f61211c7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081811115611ed9576040517f168a7f6e00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016109bc565b61216b82612802565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156121bc5761097a82826128d0565b6109cf61294f565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561221157602002820191905f5260205f20905b8154815260200190600101908083116121fd575b50505050509050919050565b5f825f01828154811061223257612232613513565b905f5260205f200154905092915050565b61224b612987565b6115cf576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612289612243565b6122916129a5565b6122996129a5565b73ffffffffffffffffffffffffffffffffffffffff8216612308576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e657200000000000000000000000000000000000000000000000000000060048201526024016109bc565b6123125f8361126e565b5061233d7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361126e565b5073ffffffffffffffffffffffffffffffffffffffff81161580159061238f57508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156123c0576123be7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8261126e565b505b6109cf7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f6129ad565b6123f2612243565b6109cf8282612a4e565b612404612243565b61126b816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250612ab1565b61244b612243565b827f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055815f1461097a5773ffffffffffffffffffffffffffffffffffffffff81166124e4576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f696e697469616c526563697069656e740000000000000000000000000000000060048201526024016109bc565b61097a5f8284612b23565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561252857505f915060039050826125d8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612579573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166125cf57505f9250600191508290506125d8565b92505f91508190505b9450945094915050565b5f8260038111156125f5576125f5613540565b036125fe575050565b600182600381111561261257612612613540565b03612649576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561265d5761265d613540565b03612697576040517ffce698f7000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b60038260038111156126ab576126ab613540565b036109cf576040517fd78bce0c000000000000000000000000000000000000000000000000000000008152600481018290526024016109bc565b5f6108c38373ffffffffffffffffffffffffffffffffffffffff8416612cf0565b5f6108c38373ffffffffffffffffffffffffffffffffffffffff8416612dca565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816127526118f8565b80519091501561276a57805160209091012092915050565b81548015612779579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816127cd611949565b8051909150156127e557805160209091012092915050565b60018201548015612779579392505050565b61097a838383612b23565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361286a576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109bc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516128f9919061356d565b5f60405180830381855af49150503d805f8114612931576040519150601f19603f3d011682016040523d82523d5f602084013e612936565b606091505b5091509150612946858383612e16565b95945050505050565b34156115cf576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6129906119c1565b5468010000000000000000900460ff16919050565b6115cf612243565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f612a06845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b612a56612243565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612aa284826135c7565b506004810161090d83826135c7565b612ab9612243565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102612b0584826135c7565b5060038101612b1483826135c7565b505f8082556001909101555050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416612b7d5781816002015f828254612b72919061370b565b90915550612c2d9050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020829052604090205482811015612c02576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101829052604481018490526064016109bc565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316612c58576002810180548390039055612c83565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612ce291815260200190565b60405180910390a350505050565b5f8181526001830160205260408120548015611383575f612d1260018361371e565b85549091505f90612d259060019061371e565b9050808214612d84575f865f018281548110612d4357612d43613513565b905f5260205f200154905080875f018481548110612d6357612d63613513565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612d9557612d95613731565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506107df565b5f818152600183016020526040812054612e0f57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556107df565b505f6107df565b606082612e2b57612e2682612ea5565b6108c3565b8151158015612e4f575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612e9e576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016109bc565b50806108c3565b805115612eb55780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115cf61375e565b5f60208284031215612eff575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146108c3575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6108c36020830184612f2e565b803573ffffffffffffffffffffffffffffffffffffffff81168114612faf575f5ffd5b919050565b5f5f60408385031215612fc5575f5ffd5b612fce83612f8c565b946020939093013593505050565b5f5f5f60608486031215612fee575f5ffd5b612ff784612f8c565b925061300560208501612f8c565b929592945050506040919091013590565b5f60208284031215613026575f5ffd5b5035919050565b5f5f6040838503121561303e575f5ffd5b8235915061304e60208401612f8c565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112613093575f5ffd5b8135602083015f5f67ffffffffffffffff8411156130b3576130b3613057565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff8211171561310057613100613057565b604052838152905080828401871015613117575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f60408385031215613144575f5ffd5b61314d83612f8c565b9150602083013567ffffffffffffffff811115613168575f5ffd5b61317485828601613084565b9150509250929050565b602080825282518282018190525f918401906040840190835b818110156131cb57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613197565b509095945050505050565b5f602082840312156131e6575f5ffd5b6108c382612f8c565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61322960e0830189612f2e565b828103604084015261323b8189612f2e565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561329d57835183526020938401939092019160010161327f565b50909b9a5050505050505050505050565b803560ff81168114612faf575f5ffd5b5f5f5f5f5f5f5f5f610100898b0312156132d6575f5ffd5b6132df89612f8c565b97506132ed60208a01612f8c565b9650604089013567ffffffffffffffff811115613308575f5ffd5b6133148b828c01613084565b965050606089013567ffffffffffffffff811115613330575f5ffd5b61333c8b828c01613084565b95505061334b60808a016132ae565b935060a0890135925061336060c08a01612f8c565b915060e089013567ffffffffffffffff81111561337b575f5ffd5b6133878b828c01613084565b9150509295985092959890939650565b5f5f5f5f5f5f5f60e0888a0312156133ad575f5ffd5b6133b688612f8c565b96506133c460208901612f8c565b955060408801359450606088013593506133e0608089016132ae565b9699959850939692959460a0840135945060c09093013592915050565b5f5f6040838503121561340e575f5ffd5b61341783612f8c565b915061304e60208401612f8c565b5f5f5f60408486031215613437575f5ffd5b833567ffffffffffffffff81111561344d575f5ffd5b8401601f8101861361345d575f5ffd5b803567ffffffffffffffff811115613473575f5ffd5b8660208260051b8401011115613487575f5ffd5b60209182019450925084013580151581146134a0575f5ffd5b809150509250925092565b600181811c908216806134bf57607f821691505b6020821081036134f6577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f6020828403121561350c575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f82111561097a57805f5260205f20601f840160051c810160208510156135a85750805b601f840160051c820191505b81811015611ed9575f81556001016135b4565b815167ffffffffffffffff8111156135e1576135e1613057565b6135f5816135ef84546134ab565b84613583565b6020601f821160018114613646575f83156136105750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611ed9565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156136935787850151825560209485019460019092019101613673565b50848210156136cf57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156107df576107df6136de565b818103818111156107df576107df6136de565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220a35b0c426beadc0765b73cba96fc6277e2ee1a4d00413bcaef4d7121352c6d2a64736f6c634300081c0033",
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
// the contract method with ID 0xb0193227.
//
// Solidity: function initialize(address _owner, address manager, string name, string symbol, uint8 decimals, uint256 initialSupply, address initialRecipient, bytes data) returns()
func (eRC20Capped *ERC20Capped) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, initialRecipient common.Address, data []byte) []byte {
	enc, err := eRC20Capped.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, initialRecipient, data)
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
