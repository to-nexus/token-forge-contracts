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

// ERC20MintingFeeMetaData contains all meta data concerning the ERC20MintingFee contract.
var ERC20MintingFeeMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"feeBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllForgeFeeBPS\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"removeForgeFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBPS\",\"type\":\"uint256\"}],\"name\":\"setFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeBPS\",\"type\":\"uint256\"}],\"name\":\"setForgeFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBPS\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBPS\",\"type\":\"uint256\"}],\"name\":\"FeeBPSUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBPS\",\"type\":\"uint256\"}],\"name\":\"ForgeFeeBPSRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBPS\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBPS\",\"type\":\"uint256\"}],\"name\":\"ForgeFeeBPSUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"ERC20Fee__InvalidFeeBPS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20MintingFee",
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516137cb6100395f395f8181611917015281816119400152611b2601526137cb5ff3fe608060405260043610610229575f3560e01c806379cc679011610131578063a40283d5116100ac578063d505accf1161007c578063e74b981b11610062578063e74b981b146106ed578063f2fde38b1461070c578063fe2df3e81461072b575f5ffd5b8063d505accf1461065e578063dd62ed3e1461067d575f5ffd5b8063a40283d5146105c4578063a9059cbb146105d8578063ad3cb1cc146105f7578063cfa8f58d1461063f575f5ffd5b806387f45353116101015780638e824877116100e75780638e8248771461057257806395d89b41146105915780639ca92df9146105a5575f5ffd5b806387f453531461050a5780638da5cb5b14610529575f5ffd5b806379cc67901461048657806379e116f2146104a55780637ecebe00146104c457806384b0196e146104e3575f5ffd5b806346904840116101c15780635c4e62c41161019157806360df3dd71161017757806360df3dd71461043157806370a0823114610453578063715018a614610472575f5ffd5b80635c4e62c4146103f1578063604b6a9c14610412575f5ffd5b806346904840146103725780634f1ef286146103ab57806352d1902d146103be5780635473d6f8146103d2575f5ffd5b806323b872dd116101fc57806323b872dd146102de578063313ce567146102fd5780633644e5151461033d57806340c10f1914610351575f5ffd5b806301ffc9a71461022d57806306fdde0314610261578063095ea7b31461028257806318160ddd146102a1575b5f5ffd5b348015610238575f5ffd5b5061024c610247366004612fd3565b61074a565b60405190151581526020015b60405180910390f35b34801561026c575f5ffd5b506102756107a5565b604051610258919061305e565b34801561028d575f5ffd5b5061024c61029c366004613091565b61085d565b3480156102ac575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b604051908152602001610258565b3480156102e9575f5ffd5b5061024c6102f83660046130bb565b610874565b348015610308575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff9091168152602001610258565b348015610348575f5ffd5b506102d0610899565b34801561035c575f5ffd5b5061037061036b366004613091565b6108a7565b005b34801561037d575f5ffd5b50610386610969565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610258565b6103706103b93660046131d5565b6109a8565b3480156103c9575f5ffd5b506102d06109c7565b3480156103dd575f5ffd5b506103706103ec366004613237565b6109f5565b3480156103fc575f5ffd5b50610405610b8d565b6040516102589190613340565b34801561041d575f5ffd5b5061037061042c366004613352565b610bb8565b34801561043c575f5ffd5b50610445610cd3565b604051610258929190613399565b34801561045e575f5ffd5b506102d061046d3660046133bd565b610e03565b34801561047d575f5ffd5b50610370610e53565b348015610491575f5ffd5b506103706104a0366004613091565b610e66565b3480156104b0575f5ffd5b506103706104bf366004613091565b610e7b565b3480156104cf575f5ffd5b506102d06104de3660046133bd565b610fd7565b3480156104ee575f5ffd5b506104f7610fe1565b60405161025897969594939291906133d8565b348015610515575f5ffd5b5061024c6105243660046133bd565b6110db565b348015610534575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff16610386565b34801561057d575f5ffd5b506102d061058c3660046133bd565b611106565b34801561059c575f5ffd5b50610275611195565b3480156105b0575f5ffd5b506103866105bf366004613352565b6111e6565b3480156105cf575f5ffd5b506102d0611211565b3480156105e3575f5ffd5b5061024c6105f2366004613091565b61123b565b348015610602575f5ffd5b506102756040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561064a575f5ffd5b506103706106593660046133bd565b611248565b348015610669575f5ffd5b5061037061067836600461346e565b6112e9565b348015610688575f5ffd5b506102d06106973660046134d8565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b3480156106f8575f5ffd5b506103706107073660046133bd565b6114b1565b348015610717575f5ffd5b506103706107263660046133bd565b6115d3565b348015610736575f5ffd5b5061037061074536600461350f565b611636565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f6d7ebe7100000000000000000000000000000000000000000000000000000000148061079f575061079f826116c3565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b90508060030180546107db90613595565b80601f016020809104026020016040519081016040528092919081815260200182805461080790613595565b80156108525780601f1061082957610100808354040283529160200191610852565b820191905f5260205f20905b81548152906001019060200180831161083557829003601f168201915b505050505091505090565b5f3361086a81858561170e565b5060019392505050565b5f3361088185828561171b565b61088c858585611802565b60019150505b9392505050565b5f6108a26118ab565b905090565b5f6108b133611106565b9050805f036108c9576108c483836118b4565b505050565b5f6108d2610969565b90505f6127106108e28486613613565b6108ec919061362a565b90508015610958576108fe8185613662565b935061090a82826118b4565b60405181815273ffffffffffffffffffffffffffffffffffffffff83169033907ff228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f469060200160405180910390a35b61096285856118b4565b5050505050565b5f7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d005b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6109b06118ff565b6109b9826119cd565b6109c382826119d5565b5050565b5f6109d0611b0e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6109fe611b7d565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610a2a5750825b90505f8267ffffffffffffffff166001148015610a465750303b155b905081158015610a54575080155b15610a8b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610aec5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610af98b8b8b8b8b611ba5565b5f5f87806020019051810190610b0f9190613675565b91509150610b1d8282611bc3565b50508315610b805784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60606108a27f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611d48565b610bc0611d54565b612710811115610c04576040517fca24260b000000000000000000000000000000000000000000000000000000008152600481018290526024015b60405180910390fd5b5f7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d00805460408051740100000000000000000000000000000000000000009092046bffffffffffffffffffffffff168252602082018590529192507fb04438781e840362c78ce8cc4b97de7afc9aaccde53ad4c035c61d532474c660910160405180910390a180546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6060807fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d015f610d0182611de2565b90508067ffffffffffffffff811115610d1c57610d1c6130f9565b604051908082528060200260200182016040528015610d45578160200160208202803683370190505b5093508067ffffffffffffffff811115610d6157610d616130f9565b604051908082528060200260200182016040528015610d8a578160200160208202803683370190505b5092505f5b81811015610dfc57610da18382611dec565b868381518110610db357610db36136a1565b60200260200101868481518110610dcc57610dcc6136a1565b602090810291909101019190915273ffffffffffffffffffffffffffffffffffffffff9091169052600101610d8f565b5050509091565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b610e5b611d54565b610e645f611e09565b565b610e7182338361171b565b6109c38282611e9e565b610e83611d54565b73ffffffffffffffffffffffffffffffffffffffff8216610ef2576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610bfb565b612710811115610f31576040517fca24260b00000000000000000000000000000000000000000000000000000000815260048101829052602401610bfb565b7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d015f80610f5e8386611ef8565b915091505f82610f6e575f610f70565b815b9050610f7d848787611f1c565b50604080518281526020810187905273ffffffffffffffffffffffffffffffffffffffff8816917fffb3fdadd9eff1590a51efedd12e5b2dec1aafbda900b1871e3a20556968ebab910160405180910390a2505050505050565b5f61079f82611f46565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100805490915015801561101f57506001810154155b611085576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610bfb565b61108d611f6e565b611095611fbf565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f61079f7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083611fe8565b5f7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d0081806111547fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d0186611ef8565b915091508161118a5782547401000000000000000000000000000000000000000090046bffffffffffffffffffffffff1661118c565b805b95945050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00916107db90613595565b5f61079f7f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740083612016565b5f6108a27f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400612021565b5f3361086a818585611802565b611250611d54565b7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d015f8061127d8385611ef8565b9150915081156112e357611291838561202a565b508373ffffffffffffffffffffffffffffffffffffffff167f6d83a066e09d89d3e22a6df1b68cd2109f54546a11242e3f943ae9dde323b528826040516112da91815260200190565b60405180910390a25b50505050565b83421115611326576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610bfb565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861139d8c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6114048261204b565b90505f61141382878787612092565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461149a576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610bfb565b6114a58a8a8a61170e565b50505050505050505050565b6114b9611d54565b73ffffffffffffffffffffffffffffffffffffffff8116611528576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666565526563697069656e7400000000000000000000000000000000000000006004820152602401610bfb565b7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d00805460405173ffffffffffffffffffffffffffffffffffffffff8481169216907faaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3905f90a380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6115db611d54565b73ffffffffffffffffffffffffffffffffffffffff811661162a576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b61163381611e09565b50565b61163e611d54565b612fcb8161164e576120be611652565b6121115b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b848110156116bb576116b382878784818110611695576116956136a1565b90506020020160208101906116aa91906133bd565b8563ffffffff16565b600101611677565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461079f565b6108c483838360016121d3565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156112e357818110156117f4576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610bfb565b6112e384848484035f6121d3565b73ffffffffffffffffffffffffffffffffffffffff8316611851576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b73ffffffffffffffffffffffffffffffffffffffff82166118a0576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b6108c483838361233c565b5f6108a2612509565b6118bd336110db565b6118f5576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610bfb565b6109c3828261257c565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061199657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661197d6125d6565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610e64576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611633611d54565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611a5a575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611a57918101906136ce565b60015b611aa8576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610bfb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611b04576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610bfb565b6108c483836125fd565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610e64576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061079f565b611bad61265f565b611bb68561269d565b61096285858585856126be565b611bcb61265f565b73ffffffffffffffffffffffffffffffffffffffff8216611c3a576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666565526563697069656e7400000000000000000000000000000000000000006004820152602401610bfb565b612710811115611c79576040517fca24260b00000000000000000000000000000000000000000000000000000000815260048101829052602401610bfb565b5f7fd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d0073ffffffffffffffffffffffffffffffffffffffff8416740100000000000000000000000000000000000000006bffffffffffffffffffffffff85160281178255604051919250905f907faaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3908290a3604080515f8152602081018490527fb04438781e840362c78ce8cc4b97de7afc9aaccde53ad4c035c61d532474c660910160405180910390a1505050565b60605f610892836127b2565b33611d937f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610e64576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610bfb565b5f61079f8261280b565b5f808080611dfa8686612815565b909450925050505b9250929050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b73ffffffffffffffffffffffffffffffffffffffff8216611eed576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b6109c3825f8361233c565b5f808080611dfa8673ffffffffffffffffffffffffffffffffffffffff871661283e565b5f611f3e8473ffffffffffffffffffffffffffffffffffffffff851684612876565b949350505050565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00610e27565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100916107db90613595565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1006107ca565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610892565b5f6108928383612892565b5f61079f825490565b5f6108928373ffffffffffffffffffffffffffffffffffffffff84166128b8565b5f61079f6120576118ab565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f6120a2888888886128d4565b9250925092506120b282826129c7565b50909695505050505050565b6120c88282612aca565b156109c35760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116612180576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610bfb565b61218a8282612aeb565b156109c35760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516612243576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b73ffffffffffffffffffffffffffffffffffffffff8416612292576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260018301602090815260408083209388168352929052208390558115610962578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161232d91815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff84166123965781816002015f82825461238b91906136e5565b909155506124469050565b73ffffffffffffffffffffffffffffffffffffffff84165f908152602082905260409020548281101561241b576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610bfb565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff831661247157600281018054839003905561249c565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516124fb91815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612533612b0c565b61253b612b87565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b73ffffffffffffffffffffffffffffffffffffffff82166125cb576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610bfb565b6109c35f838361233c565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61098c565b61260682612bdc565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612657576108c48282612caa565b6109c3612d20565b612667612d58565b610e64576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6126a561265f565b6126ae81612d76565b6126b6612d87565b611633612d87565b6126c661265f565b83515f03612722576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610bfb565b82515f0361277e576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610bfb565b817f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055805f1461096257610962858261257c565b6060815f018054806020026020016040519081016040528092919081815260200182805480156127ff57602002820191905f5260205f20905b8154815260200190600101908083116127eb575b50505050509050919050565b5f61079f82612021565b5f80806128228585612016565b5f81815260029690960160205260409095205494959350505050565b5f81815260028301602052604081205481908061286b5761285f8585612d8f565b92505f9150611e029050565b600192509050611e02565b5f8281526002840160205260408120829055611f3e8484612d9a565b5f825f0182815481106128a7576128a76136a1565b905f5260205f200154905092915050565b5f81815260028301602052604081208190556108928383612da5565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561290d57505f915060039050826129bd565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561295e573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166129b457505f9250600191508290506129bd565b92505f91508190505b9450945094915050565b5f8260038111156129da576129da6136f8565b036129e3575050565b60018260038111156129f7576129f76136f8565b03612a2e576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115612a4257612a426136f8565b03612a7c576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610bfb565b6003826003811115612a9057612a906136f8565b036109c3576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610bfb565b5f6108928373ffffffffffffffffffffffffffffffffffffffff8416612dac565b5f6108928373ffffffffffffffffffffffffffffffffffffffff8416612e8f565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612b37611f6e565b805190915015612b4f57805160209091012092915050565b81548015612b5e579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081612bb2611fbf565b805190915015612bca57805160209091012092915050565b60018201548015612b5e579392505050565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03612c44576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610bfb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051612cd39190613725565b5f60405180830381855af49150503d805f8114612d0b576040519150601f19603f3d011682016040523d82523d5f602084013e612d10565b606091505b509150915061118c858383612edb565b3415610e64576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f612d61611b7d565b5468010000000000000000900460ff16919050565b612d7e61265f565b61163381612f6a565b610e6461265f565b5f6108928383612f72565b5f6108928383612e8f565b5f61089283835b5f8181526001830160205260408120548015612e86575f612dce600183613662565b85549091505f90612de190600190613662565b9050808214612e40575f865f018281548110612dff57612dff6136a1565b905f5260205f200154905080875f018481548110612e1f57612e1f6136a1565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612e5157612e5161373b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061079f565b5f91505061079f565b5f818152600183016020526040812054612ed457508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561079f565b505f61079f565b606082612ef057612eeb82612f89565b610892565b8151158015612f14575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612f63576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610bfb565b5080610892565b6115db61265f565b5f8181526001830160205260408120541515610892565b805115612f995780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e64613768565b5f60208284031215612fe3575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610892575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6108926020830184613012565b73ffffffffffffffffffffffffffffffffffffffff81168114611633575f5ffd5b5f5f604083850312156130a2575f5ffd5b82356130ad81613070565b946020939093013593505050565b5f5f5f606084860312156130cd575f5ffd5b83356130d881613070565b925060208401356130e881613070565b929592945050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112613135575f5ffd5b8135602083015f5f67ffffffffffffffff841115613155576131556130f9565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156131a2576131a26130f9565b6040528381529050808284018710156131b9575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156131e6575f5ffd5b82356131f181613070565b9150602083013567ffffffffffffffff81111561320c575f5ffd5b61321885828601613126565b9150509250929050565b803560ff81168114613232575f5ffd5b919050565b5f5f5f5f5f5f60c0878903121561324c575f5ffd5b863561325781613070565b9550602087013567ffffffffffffffff811115613272575f5ffd5b61327e89828a01613126565b955050604087013567ffffffffffffffff81111561329a575f5ffd5b6132a689828a01613126565b9450506132b560608801613222565b92506080870135915060a087013567ffffffffffffffff8111156132d7575f5ffd5b6132e389828a01613126565b9150509295509295509295565b5f8151808452602084019350602083015f5b8281101561333657815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613302565b5093949350505050565b602081525f61089260208301846132f0565b5f60208284031215613362575f5ffd5b5035919050565b5f8151808452602084019350602083015f5b8281101561333657815186526020958601959091019060010161337b565b604081525f6133ab60408301856132f0565b828103602084015261118c8185613369565b5f602082840312156133cd575f5ffd5b813561089281613070565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61341260e0830189613012565b82810360408401526134248189613012565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c08401526134608185613369565b9a9950505050505050505050565b5f5f5f5f5f5f5f60e0888a031215613484575f5ffd5b873561348f81613070565b9650602088013561349f81613070565b955060408801359450606088013593506134bb60808901613222565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156134e9575f5ffd5b82356134f481613070565b9150602083013561350481613070565b809150509250929050565b5f5f5f60408486031215613521575f5ffd5b833567ffffffffffffffff811115613537575f5ffd5b8401601f81018613613547575f5ffd5b803567ffffffffffffffff81111561355d575f5ffd5b8660208260051b8401011115613571575f5ffd5b602091820194509250840135801515811461358a575f5ffd5b809150509250925092565b600181811c908216806135a957607f821691505b6020821081036135e0577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808202811582820484141761079f5761079f6135e6565b5f8261365d577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8181038181111561079f5761079f6135e6565b5f5f60408385031215613686575f5ffd5b825161369181613070565b6020939093015192949293505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082840312156136de575f5ffd5b5051919050565b8082018082111561079f5761079f6135e6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212203ce51058723c32be9498007a76d6669c9b8e519a506bba78a09aa99500d2503a64736f6c634300081c0033",
}

// ERC20MintingFee is an auto generated Go binding around an Ethereum contract.
type ERC20MintingFee struct {
	abi abi.ABI
}

// NewERC20MintingFee creates a new instance of ERC20MintingFee.
func NewERC20MintingFee() *ERC20MintingFee {
	parsed, err := ERC20MintingFeeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20MintingFee{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20MintingFee) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20MintingFee *ERC20MintingFee) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20MintingFee.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20MintingFee *ERC20MintingFee) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20MintingFee.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20MintingFee.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20MintingFee *ERC20MintingFee) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20MintingFee.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("allowance", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20MintingFee *ERC20MintingFee) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20MintingFee.abi.Unpack("approve", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("balanceOf", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MintingFee *ERC20MintingFee) PackDecimals() []byte {
	enc, err := eRC20MintingFee.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MintingFee *ERC20MintingFee) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20MintingFee.abi.Unpack("decimals", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackEip712Domain() []byte {
	enc, err := eRC20MintingFee.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20MintingFee *ERC20MintingFee) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20MintingFee.abi.Unpack("eip712Domain", data)
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

// PackFeeBPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e824877.
//
// Solidity: function feeBPS(address forge) view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) PackFeeBPS(forge common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("feeBPS", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFeeBPS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e824877.
//
// Solidity: function feeBPS(address forge) view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackFeeBPS(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("feeBPS", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (eRC20MintingFee *ERC20MintingFee) PackFeeRecipient() []byte {
	enc, err := eRC20MintingFee.abi.Pack("feeRecipient")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFeeRecipient is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (eRC20MintingFee *ERC20MintingFee) UnpackFeeRecipient(data []byte) (common.Address, error) {
	out, err := eRC20MintingFee.abi.Unpack("feeRecipient", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20MintingFee *ERC20MintingFee) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20MintingFee.abi.Unpack("forgeByIndex", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackForgeCount() []byte {
	enc, err := eRC20MintingFee.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("forgeCount", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackForges() []byte {
	enc, err := eRC20MintingFee.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20MintingFee *ERC20MintingFee) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20MintingFee.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetAllForgeFeeBPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60df3dd7.
//
// Solidity: function getAllForgeFeeBPS() view returns(address[] forges, uint256[] fees)
func (eRC20MintingFee *ERC20MintingFee) PackGetAllForgeFeeBPS() []byte {
	enc, err := eRC20MintingFee.abi.Pack("getAllForgeFeeBPS")
	if err != nil {
		panic(err)
	}
	return enc
}

// GetAllForgeFeeBPSOutput serves as a container for the return parameters of contract
// method GetAllForgeFeeBPS.
type GetAllForgeFeeBPSOutput struct {
	Forges []common.Address
	Fees   []*big.Int
}

// UnpackGetAllForgeFeeBPS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x60df3dd7.
//
// Solidity: function getAllForgeFeeBPS() view returns(address[] forges, uint256[] fees)
func (eRC20MintingFee *ERC20MintingFee) UnpackGetAllForgeFeeBPS(data []byte) (GetAllForgeFeeBPSOutput, error) {
	out, err := eRC20MintingFee.abi.Unpack("getAllForgeFeeBPS", data)
	outstruct := new(GetAllForgeFeeBPSOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Forges = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Fees = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5473d6f8.
// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5473d6f8.
//
// Solidity: function initialize(address _owner, string _name, string _symbol, uint8 _decimals, uint256 _initialSupply, bytes _data) returns()
func (eRC20MintingFee *ERC20MintingFee) PackInitialize(owner common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, data []byte) []byte {
	enc, err := eRC20MintingFee.abi.Pack("initialize", owner, name, symbol, decimals, initialSupply, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MintingFee *ERC20MintingFee) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MintingFee *ERC20MintingFee) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20MintingFee.abi.Unpack("isForge", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MintingFee *ERC20MintingFee) PackName() []byte {
	enc, err := eRC20MintingFee.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MintingFee *ERC20MintingFee) UnpackName(data []byte) (string, error) {
	out, err := eRC20MintingFee.abi.Unpack("name", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackNonces(owner common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20MintingFee *ERC20MintingFee) PackOwner() []byte {
	enc, err := eRC20MintingFee.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20MintingFee *ERC20MintingFee) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20MintingFee.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20MintingFee *ERC20MintingFee) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20MintingFee.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20MintingFee *ERC20MintingFee) PackProxiableUUID() []byte {
	enc, err := eRC20MintingFee.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20MintingFee *ERC20MintingFee) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20MintingFee.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
cfa8f58d.
//
// Solidity: function removeForgeFeeBPS(address forge) returns()
func (eRC20MintingFee *ERC20MintingFee) PackRemoveForgeFeeBPS(forge common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("removeForgeFeeBPS", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for 
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (eRC20MintingFee *ERC20MintingFee) PackRenounceOwnership() []byte {
	enc, err := eRC20MintingFee.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeBPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x604b6a9c.
//
// Solidity: function setFeeBPS(uint256 _feeBPS) returns()
func (eRC20MintingFee *ERC20MintingFee) PackSetFeeBPS(feeBPS *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("setFeeBPS", feeBPS)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (eRC20MintingFee *ERC20MintingFee) PackSetFeeRecipient(feeRecipient common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("setFeeRecipient", feeRecipient)
	if err != nil {
		panic(err)
	}
red for calling
// the contract method with ID 0x79e116f2.
//
// Solidity: function setForgeFeeBPS(address forge, uint256 _feeBPS) returns()
func (eRC20MintingFee *ERC20MintingFee) PackSetForgeFeeBPS(forge common.Address, feeBPS *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("setForgeFeeBPS", forge, feeBPS)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to 
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20MintingFee *ERC20MintingFee) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20MintingFee.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MintingFee *ERC20MintingFee) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20MintingFee.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MintingFee *ERC20MintingFee) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20MintingFee.abi.Unpack("supportsInterface", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackSymbol() []byte {
	enc, err := eRC20MintingFee.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20MintingFee *ERC20MintingFee) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20MintingFee.abi.Unpack("symbol", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackTotalSupply() []byte {
	enc, err := eRC20MintingFee.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20MintingFee *ERC20MintingFee) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20MintingFee.abi.Unpack("totalSupply", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20MintingFee *ERC20MintingFee) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20MintingFee.abi.Unpack("transfer", data)
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
func (eRC20MintingFee *ERC20MintingFee) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20MintingFee.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20MintingFee *ERC20MintingFee) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20MintingFee.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (eRC20MintingFee *ERC20MintingFee) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eRC20MintingFee.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (eRC20MintingFee *ERC20MintingFee) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20MintingFee.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20MintingFeeApproval represents a Approval event raised by the ERC20MintingFee contract.
type ERC20MintingFeeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeApproval) ContractEventName() string {
	return ERC20MintingFeeApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20MintingFee *ERC20MintingFee) UnpackApprovalEvent(log *types.Log) (*ERC20MintingFeeApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeApproval)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20MintingFee contract.
type ERC20MintingFeeEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeEIP712DomainChanged) ContractEventName() string {
	return ERC20MintingFeeEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20MintingFee *ERC20MintingFee) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MintingFeeEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeFeeBPSUpdated represents a FeeBPSUpdated event raised by the ERC20MintingFee contract.
type ERC20MintingFeeFeeBPSUpdated struct {
	OldBPS *big.Int
	NewBPS *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeFeeBPSUpdatedEventName = "FeeBPSUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeFeeBPSUpdated) ContractEventName() string {
	return ERC20MintingFeeFeeBPSUpdatedEventName
}

// UnpackFeeBPSUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeBPSUpdated(uint256 oldBPS, uint256 newBPS)
func (eRC20MintingFee *ERC20MintingFee) UnpackFeeBPSUpdatedEvent(log *types.Log) (*ERC20MintingFeeFeeBPSUpdated, error) {
	event := "FeeBPSUpdated"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeFeeBPSUpdated)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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


// ERC20MintingFeeFeeCollected represents a FeeCollected event raised by the ERC20MintingFee contract.
type ERC20MintingFeeFeeCollected struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeFeeCollectedEventName = "FeeCollected"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeFeeCollected) ContractEventName() string {
	return ERC20MintingFeeFeeCollectedEventName
}

// Solidity: event FeeCollected(address indexed caller, address indexed recipient, uint256 amount)
// by contract.
//
// Solidity: event FeeCollected(address indexed recipient, uint256 amount)
func (eRC20MintingFee *ERC20MintingFee) UnpackFeeCollectedEvent(log *types.Log) (*ERC20MintingFeeFeeCollected, error) {
	event := "FeeCollected"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeFeeCollected)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the ERC20MintingFee contract.
type ERC20MintingFeeFeeRecipientUpdated struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeFeeRecipientUpdatedEventName = "FeeRecipientUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeFeeRecipientUpdated) ContractEventName() string {
	return ERC20MintingFeeFeeRecipientUpdatedEventName
}

// UnpackFeeRecipientUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (eRC20MintingFee *ERC20MintingFee) UnpackFeeRecipientUpdatedEvent(log *types.Log) (*ERC20MintingFeeFeeRecipientUpdated, error) {
	event := "FeeRecipientUpdated"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeFeeRecipientUpdated)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeForgeAdded represents a ForgeAdded event raised by the ERC20MintingFee contract.
type ERC20MintingFeeForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeForgeAdded) ContractEventName() string {
	return ERC20MintingFeeForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeAddedEvent(log *types.Log) (*ERC20MintingFeeForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
ar indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeForgeFeeBPSRemoved represents a ForgeFeeBPSRemoved event raised by the ERC20MintingFee contract.
type ERC20MintingFeeForgeFeeBPSRemoved struct {
	Forge  common.Address
	OldBPS *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeForgeFeeBPSRemovedEventName = "ForgeFeeBPSRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeForgeFeeBPSRemoved) ContractEventName() string {
	return ERC20MintingFeeForgeFeeBPSRemovedEventName
}

// UnpackForgeFeeBPSRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeFeeBPSRemoved(address indexed forge, uint256 oldBPS)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeFeeBPSRemovedEvent(log *types.Log) (*ERC20MintingFeeForgeFeeBPSRemoved, error) {
	event := "ForgeFeeBPSRemoved"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeForgeFeeBPSRemoved)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeForgeFeeBPSUpdated represents a ForgeFeeBPSUpdated event raised by the ERC20MintingFee contract.
type ERC20MintingFeeForgeFeeBPSUpdated struct {
	Forge  common.Address
	OldBPS *big.Int
	NewBPS *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeForgeFeeBPSUpdatedEventName = "ForgeFeeBPSUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeForgeFeeBPSUpdated) ContractEventName() string {
	return ERC20MintingFeeForgeFeeBPSUpdatedEventName
}

// UnpackForgeFeeBPSUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeFeeBPSUpdated(address indexed forge, uint256 oldBPS, uint256 newBPS)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeFeeBPSUpdatedEvent(log *types.Log) (*ERC20MintingFeeForgeFeeBPSUpdated, error) {
	event := "ForgeFeeBPSUpdated"
	if log.Topics[0] != eRC20MintingFee.abi.Even
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeForgeRemoved represents a ForgeRemoved event raised by the ERC20MintingFee contract.
type ERC20MintingFeeForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeForgeRemoved) ContractEventName() string {
	return ERC20MintingFeeForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20MintingFee *ERC20MintingFee) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MintingFeeForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeInitialized represents a Initialized event raised by the ERC20MintingFee contract.
type ERC20MintingFeeInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeInitialized) ContractEventName() string {
	return ERC20MintingFeeInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20MintingFee *ERC20MintingFee) UnpackInitializedEvent(log *types.Log) (*ERC20MintingFeeInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeInitialized)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20MintingFee contract.
type ERC20MintingFeeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeOwnershipTransferred) ContractEventName() string {
	return ERC20MintingFeeOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eRC20MintingFee *ERC20MintingFee) UnpackOwnershipTransferredEvent(log *types.Log) (*ERC20MintingFeeOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeTransfer represents a Transfer event raised by the ERC20MintingFee contract.
type ERC20MintingFeeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeTransfer) ContractEventName() string {
	return ERC20MintingFeeTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20MintingFee *ERC20MintingFee) UnpackTransferEvent(log *types.Log) (*ERC20MintingFeeTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeTransfer)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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

// ERC20MintingFeeUpgraded represents a Upgraded event raised by the ERC20MintingFee contract.
type ERC20MintingFeeUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintingFeeUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20MintingFeeUpgraded) ContractEventName() string {
	return ERC20MintingFeeUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20MintingFee *ERC20MintingFee) UnpackUpgradedEvent(log *types.Log) (*ERC20MintingFeeUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20MintingFee.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintingFeeUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20MintingFee.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintingFee.abi.Events[event].Inputs {
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
func (eRC20MintingFee *ERC20MintingFee) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20FeeInvalidFeeBPS"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20FeeInvalidFeeBPSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintingFee.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20MintingFee.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MintingFeeAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20MintingFee contract.
type ERC20MintingFeeAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20MintingFeeAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20MintingFee *ERC20MintingFee) UnpackAddressEmptyCodeError(raw []byte) (*ERC20MintingFeeAddressEmptyCode, error) {
	out := new(ERC20MintingFeeAddressEmptyCode)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20MintingFee contract.
type ERC20MintingFeeECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MintingFeeECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20MintingFee *ERC20MintingFee) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MintingFeeECDSAInvalidSignature, error) {
	out := new(ERC20MintingFeeECDSAInvalidSignature)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20MintingFee contract.
type ERC20MintingFeeECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MintingFeeECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20MintingFee *ERC20MintingFee) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MintingFeeECDSAInvalidSignatureLength, error) {
	out := new(ERC20MintingFeeECDSAInvalidSignatureLength)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20MintingFee contract.
type ERC20MintingFeeECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MintingFeeECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20MintingFee *ERC20MintingFee) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MintingFeeECDSAInvalidSignatureS, error) {
	out := new(ERC20MintingFeeECDSAInvalidSignatureS)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20MintingFeeERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20MintingFeeERC1967InvalidImplementation, error) {
	out := new(ERC20MintingFeeERC1967InvalidImplementation)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20MintingFeeERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20MintingFee *ERC20MintingFee) UnpackERC1967NonPayableError(raw []byte) (*ERC20MintingFeeERC1967NonPayable, error) {
	out := new(ERC20MintingFeeERC1967NonPayable)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20FeeInvalidFeeBPS represents a ERC20Fee__InvalidFeeBPS error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20FeeInvalidFeeBPS struct {
	FeeBPS *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Fee__InvalidFeeBPS(uint256 feeBPS)
func ERC20MintingFeeERC20FeeInvalidFeeBPSErrorID() common.Hash {
	return common.HexToHash("0xca24260b1eef995eb83a05adfda6c8ba89d4596eb4455231b6381101c6b92e4d")
}

// UnpackERC20FeeInvalidFeeBPSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Fee__InvalidFeeBPS(uint256 feeBPS)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20FeeInvalidFeeBPSError(raw []byte) (*ERC20MintingFeeERC20FeeInvalidFeeBPS, error) {
	out := new(ERC20MintingFeeERC20FeeInvalidFeeBPS)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20FeeInvalidFeeBPS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MintingFeeERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MintingFeeERC20InsufficientAllowance, error) {
	out := new(ERC20MintingFeeERC20InsufficientAllowance)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MintingFeeERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MintingFeeERC20InsufficientBalance, error) {
	out := new(ERC20MintingFeeERC20InsufficientBalance)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MintingFeeERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MintingFeeERC20InvalidApprover, error) {
	out := new(ERC20MintingFeeERC20InvalidApprover)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MintingFeeERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MintingFeeERC20InvalidReceiver, error) {
	out := new(ERC20MintingFeeERC20InvalidReceiver)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MintingFeeERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MintingFeeERC20InvalidSender, error) {
	out := new(ERC20MintingFeeERC20InvalidSender)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MintingFeeERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MintingFeeERC20InvalidSpender, error) {
	out := new(ERC20MintingFeeERC20InvalidSpender)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MintingFeeERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MintingFeeERC2612ExpiredSignature, error) {
	out := new(ERC20MintingFeeERC2612ExpiredSignature)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20MintingFee contract.
type ERC20MintingFeeERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MintingFeeERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20MintingFee *ERC20MintingFee) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MintingFeeERC2612InvalidSigner, error) {
	out := new(ERC20MintingFeeERC2612InvalidSigner)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeFailedCall represents a FailedCall error raised by the ERC20MintingFee contract.
type ERC20MintingFeeFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20MintingFeeFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20MintingFee *ERC20MintingFee) UnpackFailedCallError(raw []byte) (*ERC20MintingFeeFailedCall, error) {
	out := new(ERC20MintingFeeFailedCall)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20MintingFee contract.
type ERC20MintingFeeInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MintingFeeInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20MintingFee *ERC20MintingFee) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MintingFeeInvalidAccountNonce, error) {
	out := new(ERC20MintingFeeInvalidAccountNonce)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeInvalidInitialization represents a InvalidInitialization error raised by the ERC20MintingFee contract.
type ERC20MintingFeeInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20MintingFeeInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20MintingFee *ERC20MintingFee) UnpackInvalidInitializationError(raw []byte) (*ERC20MintingFeeInvalidInitialization, error) {
	out := new(ERC20MintingFeeInvalidInitialization)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeNotInitializing represents a NotInitializing error raised by the ERC20MintingFee contract.
type ERC20MintingFeeNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20MintingFeeNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20MintingFee *ERC20MintingFee) UnpackNotInitializingError(raw []byte) (*ERC20MintingFeeNotInitializing, error) {
	out := new(ERC20MintingFeeNotInitializing)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ERC20MintingFee contract.
type ERC20MintingFeeOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ERC20MintingFeeOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eRC20MintingFee *ERC20MintingFee) UnpackOwnableInvalidOwnerError(raw []byte) (*ERC20MintingFeeOwnableInvalidOwner, error) {
	out := new(ERC20MintingFeeOwnableInvalidOwner)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ERC20MintingFee contract.
type ERC20MintingFeeOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ERC20MintingFeeOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eRC20MintingFee *ERC20MintingFee) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ERC20MintingFeeOwnableUnauthorizedAccount, error) {
	out := new(ERC20MintingFeeOwnableUnauthorizedAccount)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20MintingFee contract.
type ERC20MintingFeeTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MintingFeeTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20MintingFee *ERC20MintingFee) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MintingFeeTokenBaseNullInput, error) {
	out := new(ERC20MintingFeeTokenBaseNullInput)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20MintingFee contract.
type ERC20MintingFeeTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MintingFeeTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20MintingFee *ERC20MintingFee) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MintingFeeTokenBaseOnlyForge, error) {
	out := new(ERC20MintingFeeTokenBaseOnlyForge)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20MintingFee contract.
type ERC20MintingFeeUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20MintingFeeUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20MintingFee *ERC20MintingFee) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20MintingFeeUUPSUnauthorizedCallContext, error) {
	out := new(ERC20MintingFeeUUPSUnauthorizedCallContext)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintingFeeUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20MintingFee contract.
type ERC20MintingFeeUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20MintingFeeUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20MintingFee *ERC20MintingFee) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20MintingFeeUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20MintingFeeUUPSUnsupportedProxiableUUID)
	if err := eRC20MintingFee.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
