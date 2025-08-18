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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacities\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriod\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodBlocks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartBlocks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"updateMintLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"oldLimits\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"ERC20MultiMintLimited__CapTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20MultiMintLimited__InvalidInitialData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodsMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20MultiMintLimited",
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516145106100395f395f8181611ca801528181611cd10152611ef101526145105ff3fe608060405260043610610291575f3560e01c806370a0823111610165578063a40283d5116100c6578063d547741f1161007c578063dd62ed3e11610062578063dd62ed3e146107c6578063ec87621c14610836578063fe2df3e814610869575f5ffd5b8063d547741f14610793578063da0239a6146107b2575f5ffd5b8063ad3cb1cc116100ac578063ad3cb1cc14610718578063c3ac4c2d14610760578063d505accf14610774575f5ffd5b8063a40283d5146106e5578063a9059cbb146106f9575f5ffd5b806387f453531161011b57806395d89b411161010157806395d89b411461067a5780639ca92df91461068e578063a217fddf146106d2575f5ffd5b806387f45353146105eb57806391d148541461060a575f5ffd5b806379cc67901161014b57806379cc6790146105865780637ecebe00146105a557806384b0196e146105c4575f5ffd5b806370a082311461054857806379320f7814610567575f5ffd5b8063313ce5671161020f57806340c10f19116101c557806352d1902d116101ab57806352d1902d146104ff5780635c4e62c4146105135780636148fc1514610534575f5ffd5b806340c10f19146104cd5780634f1ef286146104ec575f5ffd5b80633644e515116101f55780633644e5151461047b57806336568abe1461048f57806337e653ea146104ae575f5ffd5b8063313ce56714610408578063355274ea14610448575f5ffd5b806323b872dd116102645780632a71222c1161024a5780632a71222c146103b25780632d9ed80d146103d35780632f2ff15d146103e7575f5ffd5b806323b872dd14610346578063248a9ca314610365575f5ffd5b806301ffc9a71461029557806306fdde03146102c9578063095ea7b3146102ea57806318160ddd14610309575b5f5ffd5b3480156102a0575f5ffd5b506102b46102af366004613a05565b610888565b60405190151581526020015b60405180910390f35b3480156102d4575f5ffd5b506102dd6108e3565b6040516102c09190613a90565b3480156102f5575f5ffd5b506102b4610304366004613aca565b61099b565b348015610314575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016102c0565b348015610351575f5ffd5b506102b4610360366004613af2565b6109b2565b348015610370575f5ffd5b5061033861037f366004613b2c565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156103bd575f5ffd5b506103c66109c8565b6040516102c09190613b7d565b3480156103de575f5ffd5b506103c6610a40565b3480156103f2575f5ffd5b50610406610401366004613b8f565b610ab6565b005b348015610413575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff90911681526020016102c0565b348015610453575f5ffd5b507f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30054610338565b348015610486575f5ffd5b50610338610aff565b34801561049a575f5ffd5b506104066104a9366004613b8f565b610b0d565b3480156104b9575f5ffd5b506104066104c8366004613c01565b610b6b565b3480156104d8575f5ffd5b506104066104e7366004613aca565b610d50565b6104066104fa366004613d4c565b610d5e565b34801561050a575f5ffd5b50610338610d79565b34801561051e575f5ffd5b50610527610da7565b6040516102c09190613d97565b34801561053f575f5ffd5b506103c6610dd2565b348015610553575f5ffd5b50610338610562366004613de4565b610f19565b348015610572575f5ffd5b50610406610581366004613e0d565b610f69565b348015610591575f5ffd5b506104066105a0366004613aca565b61118f565b3480156105b0575f5ffd5b506103386105bf366004613de4565b6111a4565b3480156105cf575f5ffd5b506105d86111ae565b6040516102c09796959493929190613ed5565b3480156105f6575f5ffd5b506102b4610605366004613de4565b6112a8565b348015610615575f5ffd5b506102b4610624366004613b8f565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b348015610685575f5ffd5b506102dd6112d3565b348015610699575f5ffd5b506106ad6106a8366004613b2c565b611324565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c0565b3480156106dd575f5ffd5b506103385f81565b3480156106f0575f5ffd5b5061033861134f565b348015610704575f5ffd5b506102b4610713366004613aca565b611379565b348015610723575f5ffd5b506102dd6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561076b575f5ffd5b506103c6611384565b34801561077f575f5ffd5b5061040661078e366004613f6b565b611529565b34801561079e575f5ffd5b506104066107ad366004613b8f565b6116f1565b3480156107bd575f5ffd5b50610338611734565b3480156107d1575f5ffd5b506103386107e0366004613fd1565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610841575f5ffd5b506103387faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b348015610874575f5ffd5b50610406610883366004613ff9565b6117a3565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f403f67f10000000000000000000000000000000000000000000000000000000014806108dd57506108dd8261184a565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461091990614050565b80601f016020809104026020016040519081016040528092919081815260200182805461094590614050565b80156109905780601f1061096757610100808354040283529160200191610990565b820191905f5260205f20905b81548152906001019060200180831161097357829003601f168201915b505050505091505090565b5f336109a8818585611854565b5060019392505050565b5f6109be848484611861565b90505b9392505050565b60607f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00600201805480602002602001604051908101604052809291908181526020018280548015610a3657602002820191905f5260205f20905b815481526020019060010190808311610a22575b5050505050905090565b60607f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00600101805480602002602001604051908101604052809291908181526020018280548015610a3657602002820191905f5260205f2090815481526020019060010190808311610a22575050505050905090565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610aef81611884565b610af98383611891565b50505050565b5f610b086119af565b905090565b73ffffffffffffffffffffffffffffffffffffffff81163314610b5c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b6682826119b8565b505050565b5f610b7581611884565b815f819003610bd7576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d697473000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d0080548214610c32576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805b83811015610cfa575f878783818110610c5057610c506140a1565b905060200201359050805f03610cb4576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69747300000000000000000000000000000000000000000000006004820152602401610bce565b828111610cf0576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101839052602401610bce565b9150600101610c35565b507f89be4a9eb7586334d719b1c87a0fd1e8306235335bf257d957c53aa7b22fcb74826001018787604051610d31939291906140ce565b60405180910390a1610d47600183018787613967565b50505050505050565b610d5a8282611a94565b5050565b610d66611c90565b610d6f82611d96565b610d5a8282611da0565b5f610d82611ed9565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6060610b087f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611f48565b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d028054604080516020808402820181019092528281526060937f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00935f93830182828015610e5c57602002820191905f5260205f20905b815481526020019060010190808311610e48575b505050505090505f815190505f8167ffffffffffffffff811115610e8257610e82613c40565b604051908082528060200260200182016040528015610eab578160200160208202803683370190505b509050435f5b83811015610f0e575f858281518110610ecc57610ecc6140a1565b60200260200101519050808381610ee557610ee561415d565b068303848381518110610efa57610efa6140a1565b602090810291909101015250600101610eb1565b509095945050505050565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b5f610f72611f54565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610f9e5750825b90505f8267ffffffffffffffff166001148015610fba5750303b155b905081158015610fc8575080155b15610fff576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156110605784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b606086511161109b576040517f94c8879f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f888060200190518101906110b2919061420a565b925092509250898310156110fc576040517fce739feb00000000000000000000000000000000000000000000000000000000815260048101849052602481018b9052604401610bce565b6111068282611f7c565b61110f836121ed565b61111d8f8f8f8f8f8f612274565b50505083156111815784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b61119a8233836122ae565b610d5a8282612395565b5f6108dd826123ef565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10080549091501580156111ec57506001810154155b611252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610bce565b61125a612417565b611262612468565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6108dd7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083612491565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161091990614050565b5f6108dd7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400836124bf565b5f610b087f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474006124ca565b5f6109c183836124d3565b60607f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d005f6113b0610dd2565b90505f826003018054806020026020016040519081016040528092919081815260200182805480156113ff57602002820191905f5260205f20905b8154815260200190600101908083116113eb575b505086549394505f92508391505067ffffffffffffffff81111561142557611425613c40565b60405190808252806020026020018201604052801561144e578160200160208202803683370190505b5090505f5b8281101561151f575f84828151811061146e5761146e6140a1565b6020026020010151905085828151811061148a5761148a6140a1565b602002602001015181036114d9578660040182815481106114ad576114ad6140a1565b905f5260205f2001548383815181106114c8576114c86140a1565b602002602001018181525050611516565b8660010182815481106114ee576114ee6140a1565b905f5260205f200154838381518110611509576115096140a1565b6020026020010181815250505b50600101611453565b5095945050505050565b83421115611566576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610bce565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886115dd8c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f611644826124e0565b90505f61165382878787612527565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146116da576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610bce565b6116e58a8a8a611854565b50505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461172a81611884565b610af983836119b8565b5f5f61175e7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90505f6117897f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b9050818111611798575f61179c565b8181035b9250505090565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6117cd81611884565b6139b0826117dd576125536117e1565b6125a65b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b85811015610d475761184282888884818110611824576118246140a1565b90506020020160208101906118399190613de4565b8563ffffffff16565b600101611806565b5f6108dd82612668565b610b6683838360016126fe565b5f3361186e8582856122ae565b611879858585612868565b506001949350505050565b61188e8133612911565b50565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166119a6575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556119423390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506108dd565b5f9150506108dd565b5f610b086129b7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156119a6575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506108dd565b7f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d007f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d037f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d045f611b00610dd2565b84549091505f5b81811015611c85575f5f5f868481548110611b2457611b246140a1565b905f5260205f200154888581548110611b3f57611b3f6140a1565b905f5260205f200154878681518110611b5a57611b5a6140a1565b6020026020010151925092509250808214611bee575f896001018581548110611b8557611b856140a1565b905f5260205f200154905081898681548110611ba357611ba36140a1565b905f5260205f200181905550809350817ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d182604051611be491815260200190565b60405180910390a2505b89831015611c5a57886002018481548110611c0b57611c0b6140a1565b5f918252602090912001546040517faace2cbe0000000000000000000000000000000000000000000000000000000081526004810191909152602481018b905260448101849052606401610bce565b898303878581548110611c6f57611c6f6140a1565b5f91825260209091200155505050600101611b07565b50610d478787612a2a565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611d5d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611d447f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611d94576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610d5a81611884565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611e25575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611e2291810190614278565b60015b611e73576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610bce565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611ecf576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610bce565b610b668383612a75565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611d94576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f6109c183612ad7565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006108dd565b611f84612b30565b8051801580611f94575082518114155b15611fcb576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80805b838110156120cc575f5f878381518110611feb57611feb6140a1565b6020026020010151878481518110612005576120056140a1565b602002602001015191509150815f148061201d575080155b15612076576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6c696d697473206f7220706572696f64730000000000000000000000000000006004820152602401610bce565b84821115806120855750838111155b156120bf576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101849052602401610bce565b9093509150600101611fcf565b507f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d008381558451612123907f70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d019060208801906139b8565b50855161213990600283019060208901906139b8565b508367ffffffffffffffff81111561215357612153613c40565b60405190808252806020026020018201604052801561217c578160200160208202803683370190505b5080516121939160038401916020909101906139b8565b508367ffffffffffffffff8111156121ad576121ad613c40565b6040519080825280602002602001820160405280156121d6578160200160208202803683370190505b508051610d479160048401916020909101906139b8565b6121f5612b30565b805f03612250576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f63617000000000000000000000000000000000000000000000000000000000006004820152602401610bce565b7f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b30055565b61227c612b30565b6122868686612b6e565b6122938685858585612cd7565b61229d8484612dcb565b6122a684612ddd565b505050505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610af95781811015612387576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610bce565b610af984848484035f6126fe565b73ffffffffffffffffffffffffffffffffffffffff82166123e4576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b610d5a825f83612e24565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610f3d565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161091990614050565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100610908565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156109c1565b5f6109c18383612e2f565b5f6108dd825490565b5f336109a8818585612868565b5f6108dd6124ec6119af565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f61253788888888612e55565b9250925092506125478282612f48565b50909695505050505050565b61255d828261304b565b15610d5a5760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116612615576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610bce565b61261f828261306c565b15610d5a5760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108dd57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146108dd565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff851661276e576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b73ffffffffffffffffffffffffffffffffffffffff84166127bd576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115612861578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161285891815260200190565b60405180910390a35b5050505050565b73ffffffffffffffffffffffffffffffffffffffff83166128b7576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b73ffffffffffffffffffffffffffffffffffffffff8216612906576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b610b66838383612e24565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610d5a576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610bce565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6129e161308d565b6129e9613108565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b612a33336112a8565b612a6b576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610bce565b610d5a828261315d565b612a7e826131b7565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612acf57610b668282613285565b610d5a613304565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612b2457602002820191905f5260205f20905b815481526020019060010190808311612b10575b50505050509050919050565b612b3861333c565b611d94576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b76612b30565b612b7e61335a565b612b8661335a565b73ffffffffffffffffffffffffffffffffffffffff8216612bf5576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610bce565b612bff5f83611891565b50612c2a7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c83611891565b5073ffffffffffffffffffffffffffffffffffffffff811615801590612c7c57508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b15612cad57612cab7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c82611891565b505b610d5a7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f613362565b612cdf612b30565b83515f03612d3b576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610bce565b82515f03612d97576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610bce565b817f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055805f1461286157612861858261315d565b612dd3612b30565b610d5a8282613403565b612de5612b30565b61188e816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250613466565b610b668383836134d8565b5f825f018281548110612e4457612e446140a1565b905f5260205f200154905092915050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115612e8e57505f91506003905082612f3e565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612edf573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612f3557505f925060019150829050612f3e565b92505f91508190505b9450945094915050565b5f826003811115612f5b57612f5b61428f565b03612f64575050565b6001826003811115612f7857612f7861428f565b03612faf576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115612fc357612fc361428f565b03612ffd576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610bce565b60038260038111156130115761301161428f565b03610d5a576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610bce565b5f6109c18373ffffffffffffffffffffffffffffffffffffffff8416613598565b5f6109c18373ffffffffffffffffffffffffffffffffffffffff8416613672565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816130b8612417565b8051909150156130d057805160209091012092915050565b815480156130df579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081613133612468565b80519091501561314b57805160209091012092915050565b600182015480156130df579392505050565b73ffffffffffffffffffffffffffffffffffffffff82166131ac576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bce565b610d5a5f8383612e24565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361321f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610bce565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516132ae91906142bc565b5f60405180830381855af49150503d805f81146132e6576040519150601f19603f3d011682016040523d82523d5f602084013e6132eb565b606091505b50915091506132fb8583836136be565b95945050505050565b3415611d94576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613345611f54565b5468010000000000000000900460ff16919050565b611d94612b30565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f6133bb845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b61340b612b30565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036134578482614316565b5060048101610af98382614316565b61346e612b30565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026134ba8482614316565b50600381016134c98382614316565b505f8082556001909101555050565b6134e383838361374d565b73ffffffffffffffffffffffffffffffffffffffff8316610b66575f6135277f8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b3005490565b90505f6135527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081811115612861576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610bce565b5f81815260018301602052604081205480156119a6575f6135ba60018361445a565b85549091505f906135cd9060019061445a565b905080821461362c575f865f0182815481106135eb576135eb6140a1565b905f5260205f200154905080875f01848154811061360b5761360b6140a1565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061363d5761363d61446d565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506108dd565b5f8181526001830160205260408120546136b757508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556108dd565b505f6108dd565b6060826136d3576136ce82613758565b6109c1565b81511580156136f7575073ffffffffffffffffffffffffffffffffffffffff84163b155b15613746576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610bce565b50806109c1565b610b6683838361379a565b8051156137685780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff84166137f45781816002015f8282546137e9919061449a565b909155506138a49050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020829052604090205482811015613879576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610bce565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff83166138cf5760028101805483900390556138fa565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161395991815260200190565b60405180910390a350505050565b828054828255905f5260205f209081019282156139a0579160200282015b828111156139a0578235825591602001919060010190613985565b506139ac9291506139f1565b5090565b611d946144ad565b828054828255905f5260205f209081019282156139a0579160200282015b828111156139a05782518255916020019190600101906139d6565b5b808211156139ac575f81556001016139f2565b5f60208284031215613a15575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109c1575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109c16020830184613a44565b803573ffffffffffffffffffffffffffffffffffffffff81168114613ac5575f5ffd5b919050565b5f5f60408385031215613adb575f5ffd5b613ae483613aa2565b946020939093013593505050565b5f5f5f60608486031215613b04575f5ffd5b613b0d84613aa2565b9250613b1b60208501613aa2565b929592945050506040919091013590565b5f60208284031215613b3c575f5ffd5b5035919050565b5f8151808452602084019350602083015f5b82811015613b73578151865260209586019590910190600101613b55565b5093949350505050565b602081525f6109c16020830184613b43565b5f5f60408385031215613ba0575f5ffd5b82359150613bb060208401613aa2565b90509250929050565b5f5f83601f840112613bc9575f5ffd5b50813567ffffffffffffffff811115613be0575f5ffd5b6020830191508360208260051b8501011115613bfa575f5ffd5b9250929050565b5f5f60208385031215613c12575f5ffd5b823567ffffffffffffffff811115613c28575f5ffd5b613c3485828601613bb9565b90969095509350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613cb457613cb4613c40565b604052919050565b5f82601f830112613ccb575f5ffd5b8135602083015f5f67ffffffffffffffff841115613ceb57613ceb613c40565b50601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001613d1e81613c6d565b915050828152858383011115613d32575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f60408385031215613d5d575f5ffd5b613d6683613aa2565b9150602083013567ffffffffffffffff811115613d81575f5ffd5b613d8d85828601613cbc565b9150509250929050565b602080825282518282018190525f918401906040840190835b81811015610f0e57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613db0565b5f60208284031215613df4575f5ffd5b6109c182613aa2565b803560ff81168114613ac5575f5ffd5b5f5f5f5f5f5f5f60e0888a031215613e23575f5ffd5b613e2c88613aa2565b9650613e3a60208901613aa2565b9550604088013567ffffffffffffffff811115613e55575f5ffd5b613e618a828b01613cbc565b955050606088013567ffffffffffffffff811115613e7d575f5ffd5b613e898a828b01613cbc565b945050613e9860808901613dfd565b925060a0880135915060c088013567ffffffffffffffff811115613eba575f5ffd5b613ec68a828b01613cbc565b91505092959891949750929550565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f613f0f60e0830189613a44565b8281036040840152613f218189613a44565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c0840152613f5d8185613b43565b9a9950505050505050505050565b5f5f5f5f5f5f5f60e0888a031215613f81575f5ffd5b613f8a88613aa2565b9650613f9860208901613aa2565b95506040880135945060608801359350613fb460808901613dfd565b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215613fe2575f5ffd5b613feb83613aa2565b9150613bb060208401613aa2565b5f5f5f6040848603121561400b575f5ffd5b833567ffffffffffffffff811115614021575f5ffd5b61402d86828701613bb9565b90945092505060208401358015158114614045575f5ffd5b809150509250925092565b600181811c9082168061406457607f821691505b60208210810361409b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b604080825284549082018190525f8581526020812090916060840190835b8181101561410a5783548352600193840193602090930192016140ec565b505083810360208501528481527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff851115614143575f5ffd5b8460051b9150818660208301370160200195945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82601f830112614199575f5ffd5b815167ffffffffffffffff8111156141b3576141b3613c40565b8060051b6141c360208201613c6d565b918252602081850181019290810190868411156141de575f5ffd5b6020860192505b838310156142005782518252602092830192909101906141e5565b9695505050505050565b5f5f5f6060848603121561421c575f5ffd5b8351602085015190935067ffffffffffffffff81111561423a575f5ffd5b6142468682870161418a565b925050604084015167ffffffffffffffff811115614262575f5ffd5b61426e8682870161418a565b9150509250925092565b5f60208284031215614288575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b601f821115610b6657805f5260205f20601f840160051c810160208510156142f75750805b601f840160051c820191505b81811015612861575f8155600101614303565b815167ffffffffffffffff81111561433057614330613c40565b6143448161433e8454614050565b846142d2565b6020601f821160018114614395575f831561435f5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612861565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156143e257878501518255602094850194600190920191016143c2565b508482101561441e57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156108dd576108dd61442d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b808201808211156108dd576108dd61442d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea264697066735822122093a16ac4b72468e78773c76791197b2d3474ebe020af67b26e2de85ce176b7c364736f6c634300081c0033",
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
// the contract method with ID 0x79320f78.
//
// Solidity: function initialize(address _owner, address _manager, string _name, string _symbol, uint8 _decimals, uint256 _initialSupply, bytes _data) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackInitialize(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, data []byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("initialize", owner, manager, name, symbol, decimals, initialSupply, data)
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

// PackMaxMintPerPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMaxMintPerPeriod() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("maxMintPerPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMaxMintPerPeriod(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("maxMintPerPeriod", data)
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

// PackPeriodBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a71222c.
//
// Solidity: function periodBlocks() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodBlocks() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a71222c.
//
// Solidity: function periodBlocks() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodBlocks(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodBlocks", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackPeriodStartBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6148fc15.
//
// Solidity: function periodStartBlocks() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodStartBlocks() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodStartBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6148fc15.
//
// Solidity: function periodStartBlocks() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodStartBlocks(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodStartBlocks", data)
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
	PeriodStartBlock  *big.Int
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
// Solidity: event PeriodStarted(uint256 indexed periodStartBlock, uint256 availableCapacity)
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
