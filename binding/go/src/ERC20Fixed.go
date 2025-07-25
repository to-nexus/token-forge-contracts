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

// ERC20FixedMetaData contains all meta data concerning the ERC20Fixed contract.
var ERC20FixedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Fixed__BurningNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Fixed__MintingNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20Fixed",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051612bda6100395f395f818161111701528181611140015261135c0152612bda5ff3fe6080604052600436106101af575f3560e01c806379cc6790116100e7578063a40283d511610087578063d505accf11610062578063d505accf1461052e578063dd62ed3e1461054d578063f2fde38b146105bd578063fe2df3e8146105dc575f5ffd5b8063a40283d5146104b3578063a9059cbb146104c7578063ad3cb1cc146104e6575f5ffd5b806387f45353116100c257806387f45353146103f75780638da5cb5b1461041657806395d89b41146104805780639ca92df914610494575f5ffd5b806379cc6790146103925780637ecebe00146103b157806384b0196e146103d0575f5ffd5b806340c10f19116101525780635473d6f81161012d5780635473d6f81461031f5780635c4e62c41461033e57806370a082311461035f578063715018a61461037e575f5ffd5b806340c10f19146102d75780634f1ef286146102f857806352d1902d1461030b575f5ffd5b806318160ddd1161018d57806318160ddd1461022757806323b872dd14610264578063313ce567146102835780633644e515146102c3575f5ffd5b806301ffc9a7146101b357806306fdde03146101e7578063095ea7b314610208575b5f5ffd5b3480156101be575f5ffd5b506101d26101cd3660046124a6565b6105fb565b60405190151581526020015b60405180910390f35b3480156101f2575f5ffd5b506101fb610656565b6040516101de9190612531565b348015610213575f5ffd5b506101d261022236600461256b565b61070e565b348015610232575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016101de565b34801561026f575f5ffd5b506101d261027e366004612593565b610725565b34801561028e575f5ffd5b507f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf005460405160ff90911681526020016101de565b3480156102ce575f5ffd5b5061025661074a565b3480156102e2575f5ffd5b506102f66102f136600461256b565b610758565b005b6102f66103063660046126a9565b61078a565b348015610316575f5ffd5b506102566107a9565b34801561032a575f5ffd5b506102f6610339366004612704565b6107d7565b348015610349575f5ffd5b506103526109a9565b6040516101de91906127bb565b34801561036a575f5ffd5b50610256610379366004612813565b6109d4565b348015610389575f5ffd5b506102f6610a24565b34801561039d575f5ffd5b506102f66103ac36600461256b565b610a37565b3480156103bc575f5ffd5b506102566103cb366004612813565b610a69565b3480156103db575f5ffd5b506103e4610a73565b6040516101de979695949392919061282c565b348015610402575f5ffd5b506101d2610411366004612813565b610b6d565b348015610421575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101de565b34801561048b575f5ffd5b506101fb610b98565b34801561049f575f5ffd5b5061045b6104ae3660046128eb565b610be9565b3480156104be575f5ffd5b50610256610c14565b3480156104d2575f5ffd5b506101d26104e136600461256b565b610c3e565b3480156104f1575f5ffd5b506101fb6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610539575f5ffd5b506102f6610548366004612902565b610c4b565b348015610558575f5ffd5b50610256610567366004612968565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b3480156105c8575f5ffd5b506102f66105d7366004612813565b610e13565b3480156105e7575f5ffd5b506102f66105f6366004612999565b610e76565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f6d7ebe71000000000000000000000000000000000000000000000000000000001480610650575061065082610f03565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461068c90612a1f565b80601f01602080910402602001604051908101604052809291908181526020018280546106b890612a1f565b80156107035780601f106106da57610100808354040283529160200191610703565b820191905f5260205f20905b8154815290600101906020018083116106e657829003601f168201915b505050505091505090565b5f3361071b818585610f4e565b5060019392505050565b5f33610732858285610f60565b61073d85858561104d565b60019150505b9392505050565b5f6107536110f6565b905090565b6040517f21b8298300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107926110ff565b61079b82611203565b6107a5828261120b565b5050565b5f6107b2611344565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6107e06113b3565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561080c5750825b90505f8267ffffffffffffffff1660011480156108285750303b155b905081158015610836575080155b1561086d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108ce5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b865f0361092e576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f696e697469616c537570706c790000000000000000000000000000000000000060048201526024015b60405180910390fd5b61093b8b8b8b8b8b6113db565b831561099c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60606107537f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400611400565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b610a2c61140c565b610a355f61149a565b565b6040517fd342df9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6106508261152f565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610ab157506001810154155b610b17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610925565b610b1f611557565b610b276115a8565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6106507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400836115d1565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161068c90612a1f565b5f6106507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb2747400836115ff565b5f6107537f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb274740061160a565b5f3361071b81858561104d565b83421115610c88576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610925565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610cff8c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610d6682611613565b90505f610d758287878761165a565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610dfc576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610925565b610e078a8a8a610f4e565b50505050505050505050565b610e1b61140c565b73ffffffffffffffffffffffffffffffffffffffff8116610e6a576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b610e738161149a565b50565b610e7e61140c565b61249e81610e8e57611686610e92565b6116d95b90507f226d89218e2374ee473292146d1b80fabdb6747a8e13760116d34b5eb27474005f5b84811015610efb57610ef382878784818110610ed557610ed5612a70565b9050602002016020810190610eea9190612813565b8563ffffffff16565b600101610eb7565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610650565b610f5b838383600161179b565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156110475781811015611039576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610925565b61104784848484035f61179b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661109c576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b73ffffffffffffffffffffffffffffffffffffffff82166110eb576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b610f5b838383611904565b5f610753611ad1565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806111cc57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166111b37f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610a35576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e7361140c565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611290575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261128d91810190612a9d565b60015b6112de576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610925565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461133a576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610925565b610f5b8383611b44565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a35576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610650565b6113e3611ba6565b6113ec85611be4565b6113f98585858585611c05565b5050505050565b60605f61074383611cf9565b3361144b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614610a35576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610925565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006109f8565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161068c90612a1f565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061067b565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610743565b5f6107438383611d52565b5f610650825490565b5f61065061161f6110f6565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f61166a88888888611d78565b92509250925061167a8282611e6b565b50909695505050505050565b6116908282611f6e565b156107a55760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611748576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610925565b6117528282611f8f565b156107a55760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff851661180b576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b73ffffffffffffffffffffffffffffffffffffffff841661185a576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b73ffffffffffffffffffffffffffffffffffffffff8086165f908152600183016020908152604080832093881683529290522083905581156113f9578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516118f591815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661195e5781816002015f8282546119539190612ae1565b90915550611a0e9050565b73ffffffffffffffffffffffffffffffffffffffff84165f90815260208290526040902054828110156119e3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610925565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316611a39576002810180548390039055611a64565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611ac391815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611afb611fb0565b611b0361202b565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611b4d82612080565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611b9e57610f5b828261214e565b6107a56121cd565b611bae612205565b610a35576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bec611ba6565b611bf581612223565b611bfd612234565b610e73612234565b611c0d611ba6565b83515f03611c69576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610925565b82515f03611cc5576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610925565b817f7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf0055805f146113f9576113f9858261223c565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611d4657602002820191905f5260205f20905b815481526020019060010190808311611d32575b50505050509050919050565b5f825f018281548110611d6757611d67612a70565b905f5260205f200154905092915050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611db157505f91506003905082611e61565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611e02573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611e5857505f925060019150829050611e61565b92505f91508190505b9450945094915050565b5f826003811115611e7e57611e7e612af4565b03611e87575050565b6001826003811115611e9b57611e9b612af4565b03611ed2576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115611ee657611ee6612af4565b03611f20576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610925565b6003826003811115611f3457611f34612af4565b036107a5576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610925565b5f6107438373ffffffffffffffffffffffffffffffffffffffff8416612296565b5f6107438373ffffffffffffffffffffffffffffffffffffffff8416612379565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10081611fdb611557565b805190915015611ff357805160209091012092915050565b81548015612002579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816120566115a8565b80519091501561206e57805160209091012092915050565b60018201548015612002579392505050565b8073ffffffffffffffffffffffffffffffffffffffff163b5f036120e8576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610925565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516121779190612b21565b5f60405180830381855af49150503d805f81146121af576040519150601f19603f3d011682016040523d82523d5f602084013e6121b4565b606091505b50915091506121c48583836123c5565b95945050505050565b3415610a35576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61220e6113b3565b5468010000000000000000900460ff16919050565b61222b611ba6565b610e7381612454565b610a35611ba6565b73ffffffffffffffffffffffffffffffffffffffff821661228b576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610925565b6107a55f8383611904565b5f8181526001830160205260408120548015612370575f6122b8600183612b37565b85549091505f906122cb90600190612b37565b905080821461232a575f865f0182815481106122e9576122e9612a70565b905f5260205f200154905080875f01848154811061230957612309612a70565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061233b5761233b612b4a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610650565b5f915050610650565b5f8181526001830160205260408120546123be57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610650565b505f610650565b6060826123da576123d58261245c565b610743565b81511580156123fe575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561244d576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610925565b5080610743565b610e1b611ba6565b80511561246c5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a35612b77565b5f602082840312156124b6575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610743575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61074360208301846124e5565b803573ffffffffffffffffffffffffffffffffffffffff81168114612566575f5ffd5b919050565b5f5f6040838503121561257c575f5ffd5b61258583612543565b946020939093013593505050565b5f5f5f606084860312156125a5575f5ffd5b6125ae84612543565b92506125bc60208501612543565b929592945050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612609575f5ffd5b8135602083015f5f67ffffffffffffffff841115612629576126296125cd565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715612676576126766125cd565b60405283815290508082840187101561268d575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156126ba575f5ffd5b6126c383612543565b9150602083013567ffffffffffffffff8111156126de575f5ffd5b6126ea858286016125fa565b9150509250929050565b803560ff81168114612566575f5ffd5b5f5f5f5f5f5f60c08789031215612719575f5ffd5b61272287612543565b9550602087013567ffffffffffffffff81111561273d575f5ffd5b61274989828a016125fa565b955050604087013567ffffffffffffffff811115612765575f5ffd5b61277189828a016125fa565b945050612780606088016126f4565b92506080870135915060a087013567ffffffffffffffff8111156127a2575f5ffd5b6127ae89828a016125fa565b9150509295509295509295565b602080825282518282018190525f918401906040840190835b8181101561280857835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016127d4565b509095945050505050565b5f60208284031215612823575f5ffd5b61074382612543565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61286660e08301896124e5565b828103604084015261287881896124e5565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156128da5783518352602093840193909201916001016128bc565b50909b9a5050505050505050505050565b5f602082840312156128fb575f5ffd5b5035919050565b5f5f5f5f5f5f5f60e0888a031215612918575f5ffd5b61292188612543565b965061292f60208901612543565b9550604088013594506060880135935061294b608089016126f4565b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215612979575f5ffd5b61298283612543565b915061299060208401612543565b90509250929050565b5f5f5f604084860312156129ab575f5ffd5b833567ffffffffffffffff8111156129c1575f5ffd5b8401601f810186136129d1575f5ffd5b803567ffffffffffffffff8111156129e7575f5ffd5b8660208260051b84010111156129fb575f5ffd5b6020918201945092508401358015158114612a14575f5ffd5b809150509250925092565b600181811c90821680612a3357607f821691505b602082108103612a6a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215612aad575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561065057610650612ab4565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b8181038181111561065057610650612ab4565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220f5c9a7694b24da216e785cf33fef4a93f81cc2f4ef16d7211918f4cdcd42704a64736f6c634300081c0033",
}

// ERC20Fixed is an auto generated Go binding around an Ethereum contract.
type ERC20Fixed struct {
	abi abi.ABI
}

// NewERC20Fixed creates a new instance of ERC20Fixed.
func NewERC20Fixed() *ERC20Fixed {
	parsed, err := ERC20FixedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Fixed{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Fixed) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Fixed.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20Fixed *ERC20Fixed) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20Fixed.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20Fixed *ERC20Fixed) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20Fixed.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (eRC20Fixed *ERC20Fixed) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("allowance", data)
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
func (eRC20Fixed *ERC20Fixed) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("approve", data)
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
func (eRC20Fixed *ERC20Fixed) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address , uint256 ) pure returns()
func (eRC20Fixed *ERC20Fixed) PackBurnFrom(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("burnFrom", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Fixed *ERC20Fixed) PackDecimals() []byte {
	enc, err := eRC20Fixed.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Fixed *ERC20Fixed) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Fixed.abi.Unpack("decimals", data)
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
func (eRC20Fixed *ERC20Fixed) PackEip712Domain() []byte {
	enc, err := eRC20Fixed.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Fixed *ERC20Fixed) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Fixed.abi.Unpack("eip712Domain", data)
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
func (eRC20Fixed *ERC20Fixed) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("forgeByIndex", data)
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
func (eRC20Fixed *ERC20Fixed) PackForgeCount() []byte {
	enc, err := eRC20Fixed.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("forgeCount", data)
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
func (eRC20Fixed *ERC20Fixed) PackForges() []byte {
	enc, err := eRC20Fixed.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Fixed *ERC20Fixed) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5473d6f8.
//
// Solidity: function initialize(address _owner, string _name, string _symbol, uint8 _decimals, uint256 _initialSupply, bytes ) returns()
func (eRC20Fixed *ERC20Fixed) PackInitialize(owner common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, arg5 []byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("initialize", owner, name, symbol, decimals, initialSupply, arg5)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Fixed *ERC20Fixed) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) pure returns()
func (eRC20Fixed *ERC20Fixed) PackMint(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("mint", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Fixed *ERC20Fixed) PackName() []byte {
	enc, err := eRC20Fixed.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Fixed *ERC20Fixed) UnpackName(data []byte) (string, error) {
	out, err := eRC20Fixed.abi.Unpack("name", data)
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
func (eRC20Fixed *ERC20Fixed) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("nonces", data)
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
func (eRC20Fixed *ERC20Fixed) PackOwner() []byte {
	enc, err := eRC20Fixed.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("owner", data)
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
func (eRC20Fixed *ERC20Fixed) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackProxiableUUID() []byte {
	enc, err := eRC20Fixed.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (eRC20Fixed *ERC20Fixed) PackRenounceOwnership() []byte {
	enc, err := eRC20Fixed.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20Fixed *ERC20Fixed) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Fixed.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Fixed *ERC20Fixed) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("supportsInterface", data)
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
func (eRC20Fixed *ERC20Fixed) PackSymbol() []byte {
	enc, err := eRC20Fixed.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Fixed *ERC20Fixed) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Fixed.abi.Unpack("symbol", data)
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
func (eRC20Fixed *ERC20Fixed) PackTotalSupply() []byte {
	enc, err := eRC20Fixed.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("totalSupply", data)
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
func (eRC20Fixed *ERC20Fixed) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("transfer", data)
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
func (eRC20Fixed *ERC20Fixed) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("transferFrom", data)
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
func (eRC20Fixed *ERC20Fixed) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (eRC20Fixed *ERC20Fixed) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20FixedApproval represents a Approval event raised by the ERC20Fixed contract.
type ERC20FixedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20FixedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20FixedApproval) ContractEventName() string {
	return ERC20FixedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Fixed *ERC20Fixed) UnpackApprovalEvent(log *types.Log) (*ERC20FixedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedApproval)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Fixed contract.
type ERC20FixedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20FixedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20FixedEIP712DomainChanged) ContractEventName() string {
	return ERC20FixedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Fixed *ERC20Fixed) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20FixedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedForgeAdded represents a ForgeAdded event raised by the ERC20Fixed contract.
type ERC20FixedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20FixedForgeAdded) ContractEventName() string {
	return ERC20FixedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Fixed *ERC20Fixed) UnpackForgeAddedEvent(log *types.Log) (*ERC20FixedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedForgeRemoved represents a ForgeRemoved event raised by the ERC20Fixed contract.
type ERC20FixedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20FixedForgeRemoved) ContractEventName() string {
	return ERC20FixedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Fixed *ERC20Fixed) UnpackForgeRemovedEvent(log *types.Log) (*ERC20FixedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedInitialized represents a Initialized event raised by the ERC20Fixed contract.
type ERC20FixedInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20FixedInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20FixedInitialized) ContractEventName() string {
	return ERC20FixedInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20Fixed *ERC20Fixed) UnpackInitializedEvent(log *types.Log) (*ERC20FixedInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedInitialized)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Fixed contract.
type ERC20FixedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ERC20FixedOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ERC20FixedOwnershipTransferred) ContractEventName() string {
	return ERC20FixedOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eRC20Fixed *ERC20Fixed) UnpackOwnershipTransferredEvent(log *types.Log) (*ERC20FixedOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedTransfer represents a Transfer event raised by the ERC20Fixed contract.
type ERC20FixedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20FixedTransfer) ContractEventName() string {
	return ERC20FixedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Fixed *ERC20Fixed) UnpackTransferEvent(log *types.Log) (*ERC20FixedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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

// ERC20FixedUpgraded represents a Upgraded event raised by the ERC20Fixed contract.
type ERC20FixedUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20FixedUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20FixedUpgraded) ContractEventName() string {
	return ERC20FixedUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20Fixed *ERC20Fixed) UnpackUpgradedEvent(log *types.Log) (*ERC20FixedUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
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
func (eRC20Fixed *ERC20Fixed) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20FixedBurningNotAllowed"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20FixedBurningNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20FixedMintingNotAllowed"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20FixedMintingNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20FixedAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20Fixed contract.
type ERC20FixedAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20FixedAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20Fixed *ERC20Fixed) UnpackAddressEmptyCodeError(raw []byte) (*ERC20FixedAddressEmptyCode, error) {
	out := new(ERC20FixedAddressEmptyCode)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20FixedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20FixedECDSAInvalidSignature, error) {
	out := new(ERC20FixedECDSAInvalidSignature)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20FixedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20FixedECDSAInvalidSignatureLength, error) {
	out := new(ERC20FixedECDSAInvalidSignatureLength)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20FixedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20FixedECDSAInvalidSignatureS, error) {
	out := new(ERC20FixedECDSAInvalidSignatureS)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20Fixed contract.
type ERC20FixedERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20FixedERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20Fixed *ERC20Fixed) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20FixedERC1967InvalidImplementation, error) {
	out := new(ERC20FixedERC1967InvalidImplementation)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20Fixed contract.
type ERC20FixedERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20FixedERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20Fixed *ERC20Fixed) UnpackERC1967NonPayableError(raw []byte) (*ERC20FixedERC1967NonPayable, error) {
	out := new(ERC20FixedERC1967NonPayable)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20FixedBurningNotAllowed represents a ERC20Fixed__BurningNotAllowed error raised by the ERC20Fixed contract.
type ERC20FixedERC20FixedBurningNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Fixed__BurningNotAllowed()
func ERC20FixedERC20FixedBurningNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xd342df9019e4f1569194e8182eed24aea86c9425b082b66fd407767649e53ed5")
}

// UnpackERC20FixedBurningNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Fixed__BurningNotAllowed()
func (eRC20Fixed *ERC20Fixed) UnpackERC20FixedBurningNotAllowedError(raw []byte) (*ERC20FixedERC20FixedBurningNotAllowed, error) {
	out := new(ERC20FixedERC20FixedBurningNotAllowed)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20FixedBurningNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20FixedMintingNotAllowed represents a ERC20Fixed__MintingNotAllowed error raised by the ERC20Fixed contract.
type ERC20FixedERC20FixedMintingNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Fixed__MintingNotAllowed()
func ERC20FixedERC20FixedMintingNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x21b82983de03f0653d31e4c2f78db5241001ff57ffdbc14aa9d7a41d9c814b1e")
}

// UnpackERC20FixedMintingNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Fixed__MintingNotAllowed()
func (eRC20Fixed *ERC20Fixed) UnpackERC20FixedMintingNotAllowedError(raw []byte) (*ERC20FixedERC20FixedMintingNotAllowed, error) {
	out := new(ERC20FixedERC20FixedMintingNotAllowed)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20FixedMintingNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Fixed contract.
type ERC20FixedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20FixedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20FixedERC20InsufficientAllowance, error) {
	out := new(ERC20FixedERC20InsufficientAllowance)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Fixed contract.
type ERC20FixedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20FixedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20FixedERC20InsufficientBalance, error) {
	out := new(ERC20FixedERC20InsufficientBalance)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20FixedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidApproverError(raw []byte) (*ERC20FixedERC20InvalidApprover, error) {
	out := new(ERC20FixedERC20InvalidApprover)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20FixedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20FixedERC20InvalidReceiver, error) {
	out := new(ERC20FixedERC20InvalidReceiver)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20FixedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidSenderError(raw []byte) (*ERC20FixedERC20InvalidSender, error) {
	out := new(ERC20FixedERC20InvalidSender)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20FixedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20FixedERC20InvalidSpender, error) {
	out := new(ERC20FixedERC20InvalidSpender)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Fixed contract.
type ERC20FixedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20FixedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Fixed *ERC20Fixed) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20FixedERC2612ExpiredSignature, error) {
	out := new(ERC20FixedERC2612ExpiredSignature)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Fixed contract.
type ERC20FixedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20FixedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Fixed *ERC20Fixed) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20FixedERC2612InvalidSigner, error) {
	out := new(ERC20FixedERC2612InvalidSigner)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedFailedCall represents a FailedCall error raised by the ERC20Fixed contract.
type ERC20FixedFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20FixedFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20Fixed *ERC20Fixed) UnpackFailedCallError(raw []byte) (*ERC20FixedFailedCall, error) {
	out := new(ERC20FixedFailedCall)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Fixed contract.
type ERC20FixedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20FixedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Fixed *ERC20Fixed) UnpackInvalidAccountNonceError(raw []byte) (*ERC20FixedInvalidAccountNonce, error) {
	out := new(ERC20FixedInvalidAccountNonce)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedInvalidInitialization represents a InvalidInitialization error raised by the ERC20Fixed contract.
type ERC20FixedInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20FixedInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20Fixed *ERC20Fixed) UnpackInvalidInitializationError(raw []byte) (*ERC20FixedInvalidInitialization, error) {
	out := new(ERC20FixedInvalidInitialization)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedNotInitializing represents a NotInitializing error raised by the ERC20Fixed contract.
type ERC20FixedNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20FixedNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20Fixed *ERC20Fixed) UnpackNotInitializingError(raw []byte) (*ERC20FixedNotInitializing, error) {
	out := new(ERC20FixedNotInitializing)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ERC20Fixed contract.
type ERC20FixedOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ERC20FixedOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eRC20Fixed *ERC20Fixed) UnpackOwnableInvalidOwnerError(raw []byte) (*ERC20FixedOwnableInvalidOwner, error) {
	out := new(ERC20FixedOwnableInvalidOwner)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ERC20Fixed contract.
type ERC20FixedOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ERC20FixedOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eRC20Fixed *ERC20Fixed) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ERC20FixedOwnableUnauthorizedAccount, error) {
	out := new(ERC20FixedOwnableUnauthorizedAccount)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Fixed contract.
type ERC20FixedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20FixedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Fixed *ERC20Fixed) UnpackTokenBaseNullInputError(raw []byte) (*ERC20FixedTokenBaseNullInput, error) {
	out := new(ERC20FixedTokenBaseNullInput)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Fixed contract.
type ERC20FixedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20FixedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Fixed *ERC20Fixed) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20FixedTokenBaseOnlyForge, error) {
	out := new(ERC20FixedTokenBaseOnlyForge)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20Fixed contract.
type ERC20FixedUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20FixedUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20Fixed *ERC20Fixed) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20FixedUUPSUnauthorizedCallContext, error) {
	out := new(ERC20FixedUUPSUnauthorizedCallContext)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20Fixed contract.
type ERC20FixedUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20FixedUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20Fixed *ERC20Fixed) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20FixedUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20FixedUUPSUnsupportedProxiableUUID)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
