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

// ERC20MintableMetaData contains all meta data concerning the ERC20Mintable contract.
var ERC20MintableMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_forges\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "ERC20Mintable",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051612bd56100395f395f81816111120152818161113b01526113570152612bd55ff3fe6080604052600436106101af575f3560e01c806379cc6790116100e7578063a40283d511610087578063d505accf11610062578063d505accf1461052e578063dd62ed3e1461054d578063f2fde38b146105bd578063fe2df3e8146105dc575f5ffd5b8063a40283d5146104b3578063a9059cbb146104c7578063ad3cb1cc146104e6575f5ffd5b806387f45353116100c257806387f45353146103f75780638da5cb5b1461041657806395d89b41146104805780639ca92df914610494575f5ffd5b806379cc6790146103925780637ecebe00146103b157806384b0196e146103d0575f5ffd5b806340c10f19116101525780635473d6f81161012d5780635473d6f81461031f5780635c4e62c41461033e57806370a082311461035f578063715018a61461037e575f5ffd5b806340c10f19146102d75780634f1ef286146102f857806352d1902d1461030b575f5ffd5b806318160ddd1161018d57806318160ddd1461022757806323b872dd14610264578063313ce567146102835780633644e515146102c3575f5ffd5b806301ffc9a7146101b357806306fdde03146101e7578063095ea7b314610208575b5f5ffd5b3480156101be575f5ffd5b506101d26101cd3660046124a1565b6105fb565b60405190151581526020015b60405180910390f35b3480156101f2575f5ffd5b506101fb610656565b6040516101de919061252c565b348015610213575f5ffd5b506101d2610222366004612566565b61070e565b348015610232575f5ffd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016101de565b34801561026f575f5ffd5b506101d261027e36600461258e565b610725565b34801561028e575f5ffd5b507f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c8005460405160ff90911681526020016101de565b3480156102ce575f5ffd5b5061025661074a565b3480156102e2575f5ffd5b506102f66102f1366004612566565b610758565b005b6102f66103063660046126a4565b6107ac565b348015610316575f5ffd5b506102566107c7565b34801561032a575f5ffd5b506102f66103393660046126ff565b6107f5565b348015610349575f5ffd5b50610352610967565b6040516101de91906127b6565b34801561036a575f5ffd5b5061025661037936600461280e565b610992565b348015610389575f5ffd5b506102f66109e2565b34801561039d575f5ffd5b506102f66103ac366004612566565b6109f5565b3480156103bc575f5ffd5b506102566103cb36600461280e565b610a0a565b3480156103db575f5ffd5b506103e4610a14565b6040516101de9796959493929190612827565b348015610402575f5ffd5b506101d261041136600461280e565b610b0e565b348015610421575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101de565b34801561048b575f5ffd5b506101fb610b39565b34801561049f575f5ffd5b5061045b6104ae3660046128e6565b610b8a565b3480156104be575f5ffd5b50610256610bb5565b3480156104d2575f5ffd5b506101d26104e1366004612566565b610bdf565b3480156104f1575f5ffd5b506101fb6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610539575f5ffd5b506102f66105483660046128fd565b610bec565b348015610558575f5ffd5b50610256610567366004612963565b73ffffffffffffffffffffffffffffffffffffffff9182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b3480156105c8575f5ffd5b506102f66105d736600461280e565b610db4565b3480156105e7575f5ffd5b506102f66105f6366004612994565b610e17565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f6d7ebe71000000000000000000000000000000000000000000000000000000001480610650575061065082610ea4565b92915050565b60605f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b905080600301805461068c90612a1a565b80601f01602080910402602001604051908101604052809291908181526020018280546106b890612a1a565b80156107035780601f106106da57610100808354040283529160200191610703565b820191905f5260205f20905b8154815290600101906020018083116106e657829003601f168201915b505050505091505090565b5f3361071b818585610eef565b5060019392505050565b5f33610732858285610f01565b61073d858585610fee565b60019150505b9392505050565b5f610753611097565b905090565b61076133610b0e565b61079e576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6107a882826110a0565b5050565b6107b46110fa565b6107bd826111fe565b6107a88282611206565b5f6107d061133f565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6107fe6113ae565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561082a5750825b90505f8267ffffffffffffffff1660011480156108465750303b155b905081158015610854575080155b1561088b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108ec5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6108f98b8b8b8b8b6113d6565b831561095a5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60606107537f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c8006113fb565b5f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace005b73ffffffffffffffffffffffffffffffffffffffff9093165f9081526020939093525050604090205490565b6109ea611407565b6109f35f611495565b565b610a00823383610f01565b6107a8828261152a565b5f61065082611584565b5f60608082808083817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008054909150158015610a5257506001810154155b610ab8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152606401610795565b610ac06115ac565b610ac86115fd565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009c939b5091995046985030975095509350915050565b5f6106507f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c80083611626565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace009161068c90612a1a565b5f6106507f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c80083611654565b5f6107537f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c80061165f565b5f3361071b818585610fee565b83421115610c29576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610795565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610ca08c73ffffffffffffffffffffffffffffffffffffffff165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610d0782611668565b90505f610d16828787876116af565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d9d576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610795565b610da88a8a8a610eef565b50505050505050505050565b610dbc611407565b73ffffffffffffffffffffffffffffffffffffffff8116610e0b576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b610e1481611495565b50565b610e1f611407565b61249981610e2f576116db610e33565b61172e5b90507f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c8005f5b84811015610e9c57610e9482878784818110610e7657610e76612a6b565b9050602002016020810190610e8b919061280e565b8563ffffffff16565b600101610e58565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610650565b610efc83838360016117f0565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610fe85781811015610fda576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610795565b610fe884848484035f6117f0565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661103d576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b73ffffffffffffffffffffffffffffffffffffffff821661108c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b610efc838383611959565b5f610753611b26565b73ffffffffffffffffffffffffffffffffffffffff82166110ef576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b6107a85f8383611959565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806111c757507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166111ae7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b156109f3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e14611407565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561128b575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261128891810190612a98565b60015b6112d9576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610795565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611335576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610795565b610efc8383611b99565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146109f3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610650565b6113de611bfb565b6113e785611c39565b6113f48585858585611c5a565b5050505050565b60605f61074383611d4e565b336114467f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16146109f3576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610795565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b73ffffffffffffffffffffffffffffffffffffffff8216611579576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b6107a8825f83611959565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006109b6565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060917fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1009161068c90612a1a565b60605f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10061067b565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610743565b5f6107438383611da7565b5f610650825490565b5f610650611674611097565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f6116bf88888888611dcd565b9250925092506116cf8282611ec0565b50909695505050505050565b6116e58282611fc3565b156107a85760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff811661179d576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610795565b6117a78282611fe4565b156107a85760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516611860576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b73ffffffffffffffffffffffffffffffffffffffff84166118af576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610795565b73ffffffffffffffffffffffffffffffffffffffff8086165f908152600183016020908152604080832093881683529290522083905581156113f4578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161194a91815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff84166119b35781816002015f8282546119a89190612adc565b90915550611a639050565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020829052604090205482811015611a38576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610795565b73ffffffffffffffffffffffffffffffffffffffff85165f9081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316611a8e576002810180548390039055611ab9565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611b1891815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611b50612005565b611b58612080565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611ba2826120d5565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611bf357610efc82826121a3565b6107a8612222565b611c0361225a565b6109f3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c41611bfb565b611c4a81612278565b611c52612289565b610e14612289565b611c62611bfb565b83515f03611cbe576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e616d65000000000000000000000000000000000000000000000000000000006004820152602401610795565b82515f03611d1a576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f73796d626f6c00000000000000000000000000000000000000000000000000006004820152602401610795565b817f5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c80055805f146113f4576113f485826110a0565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611d9b57602002820191905f5260205f20905b815481526020019060010190808311611d87575b50505050509050919050565b5f825f018281548110611dbc57611dbc612a6b565b905f5260205f200154905092915050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611e0657505f91506003905082611eb6565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611e57573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611ead57505f925060019150829050611eb6565b92505f91508190505b9450945094915050565b5f826003811115611ed357611ed3612aef565b03611edc575050565b6001826003811115611ef057611ef0612aef565b03611f27576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115611f3b57611f3b612aef565b03611f75576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610795565b6003826003811115611f8957611f89612aef565b036107a8576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610795565b5f6107438373ffffffffffffffffffffffffffffffffffffffff8416612291565b5f6107438373ffffffffffffffffffffffffffffffffffffffff8416612374565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816120306115ac565b80519091501561204857805160209091012092915050565b81548015612057579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100816120ab6115fd565b8051909150156120c357805160209091012092915050565b60018201548015612057579392505050565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361213d576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610795565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516121cc9190612b1c565b5f60405180830381855af49150503d805f8114612204576040519150601f19603f3d011682016040523d82523d5f602084013e612209565b606091505b50915091506122198583836123c0565b95945050505050565b34156109f3576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6122636113ae565b5468010000000000000000900460ff16919050565b612280611bfb565b610e148161244f565b6109f3611bfb565b5f818152600183016020526040812054801561236b575f6122b3600183612b32565b85549091505f906122c690600190612b32565b9050808214612325575f865f0182815481106122e4576122e4612a6b565b905f5260205f200154905080875f01848154811061230457612304612a6b565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061233657612336612b45565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610650565b5f915050610650565b5f8181526001830160205260408120546123b957508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610650565b505f610650565b6060826123d5576123d082612457565b610743565b81511580156123f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612448576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610795565b5080610743565b610dbc611bfb565b8051156124675780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f3612b72565b5f602082840312156124b1575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610743575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61074360208301846124e0565b803573ffffffffffffffffffffffffffffffffffffffff81168114612561575f5ffd5b919050565b5f5f60408385031215612577575f5ffd5b6125808361253e565b946020939093013593505050565b5f5f5f606084860312156125a0575f5ffd5b6125a98461253e565b92506125b76020850161253e565b929592945050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612604575f5ffd5b8135602083015f5f67ffffffffffffffff841115612624576126246125c8565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715612671576126716125c8565b604052838152905080828401871015612688575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f604083850312156126b5575f5ffd5b6126be8361253e565b9150602083013567ffffffffffffffff8111156126d9575f5ffd5b6126e5858286016125f5565b9150509250929050565b803560ff81168114612561575f5ffd5b5f5f5f5f5f5f60c08789031215612714575f5ffd5b61271d8761253e565b9550602087013567ffffffffffffffff811115612738575f5ffd5b61274489828a016125f5565b955050604087013567ffffffffffffffff811115612760575f5ffd5b61276c89828a016125f5565b94505061277b606088016126ef565b92506080870135915060a087013567ffffffffffffffff81111561279d575f5ffd5b6127a989828a016125f5565b9150509295509295509295565b602080825282518282018190525f918401906040840190835b8181101561280357835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016127cf565b509095945050505050565b5f6020828403121561281e575f5ffd5b6107438261253e565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f61286160e08301896124e0565b828103604084015261287381896124e0565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156128d55783518352602093840193909201916001016128b7565b50909b9a5050505050505050505050565b5f602082840312156128f6575f5ffd5b5035919050565b5f5f5f5f5f5f5f60e0888a031215612913575f5ffd5b61291c8861253e565b965061292a6020890161253e565b95506040880135945060608801359350612946608089016126ef565b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215612974575f5ffd5b61297d8361253e565b915061298b6020840161253e565b90509250929050565b5f5f5f604084860312156129a6575f5ffd5b833567ffffffffffffffff8111156129bc575f5ffd5b8401601f810186136129cc575f5ffd5b803567ffffffffffffffff8111156129e2575f5ffd5b8660208260051b84010111156129f6575f5ffd5b6020918201945092508401358015158114612a0f575f5ffd5b809150509250925092565b600181811c90821680612a2e57607f821691505b602082108103612a65577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215612aa8575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561065057610650612aaf565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f82518060208501845e5f920191825250919050565b8181038181111561065057610650612aaf565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212202c5e0e04ade9310812a43082128a990de24ae4001125e5ea3743baa491a4c22f64736f6c634300081c0033",
}

// ERC20Mintable is an auto generated Go binding around an Ethereum contract.
type ERC20Mintable struct {
	abi abi.ABI
}

// NewERC20Mintable creates a new instance of ERC20Mintable.
func NewERC20Mintable() *ERC20Mintable {
	parsed, err := ERC20MintableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Mintable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Mintable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Mintable.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20Mintable *ERC20Mintable) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := eRC20Mintable.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (eRC20Mintable *ERC20Mintable) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := eRC20Mintable.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (eRC20Mintable *ERC20Mintable) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("allowance", data)
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
func (eRC20Mintable *ERC20Mintable) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("approve", data)
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
func (eRC20Mintable *ERC20Mintable) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("balanceOf", data)
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
func (eRC20Mintable *ERC20Mintable) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Mintable *ERC20Mintable) PackDecimals() []byte {
	enc, err := eRC20Mintable.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Mintable *ERC20Mintable) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Mintable.abi.Unpack("decimals", data)
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
func (eRC20Mintable *ERC20Mintable) PackEip712Domain() []byte {
	enc, err := eRC20Mintable.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Mintable *ERC20Mintable) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Mintable.abi.Unpack("eip712Domain", data)
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
func (eRC20Mintable *ERC20Mintable) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("forgeByIndex", data)
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
func (eRC20Mintable *ERC20Mintable) PackForgeCount() []byte {
	enc, err := eRC20Mintable.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("forgeCount", data)
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
func (eRC20Mintable *ERC20Mintable) PackForges() []byte {
	enc, err := eRC20Mintable.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Mintable *ERC20Mintable) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("forges", data)
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
func (eRC20Mintable *ERC20Mintable) PackInitialize(owner common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, arg5 []byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("initialize", owner, name, symbol, decimals, initialSupply, arg5)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Mintable *ERC20Mintable) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("isForge", data)
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
func (eRC20Mintable *ERC20Mintable) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Mintable *ERC20Mintable) PackName() []byte {
	enc, err := eRC20Mintable.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Mintable *ERC20Mintable) UnpackName(data []byte) (string, error) {
	out, err := eRC20Mintable.abi.Unpack("name", data)
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
func (eRC20Mintable *ERC20Mintable) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("nonces", data)
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
func (eRC20Mintable *ERC20Mintable) PackOwner() []byte {
	enc, err := eRC20Mintable.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("owner", data)
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
func (eRC20Mintable *ERC20Mintable) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) PackProxiableUUID() []byte {
	enc, err := eRC20Mintable.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("proxiableUUID", data)
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
func (eRC20Mintable *ERC20Mintable) PackRenounceOwnership() []byte {
	enc, err := eRC20Mintable.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] _forges, bool add) returns()
func (eRC20Mintable *ERC20Mintable) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Mintable.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Mintable *ERC20Mintable) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("supportsInterface", data)
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
func (eRC20Mintable *ERC20Mintable) PackSymbol() []byte {
	enc, err := eRC20Mintable.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Mintable *ERC20Mintable) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Mintable.abi.Unpack("symbol", data)
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
func (eRC20Mintable *ERC20Mintable) PackTotalSupply() []byte {
	enc, err := eRC20Mintable.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("totalSupply", data)
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
func (eRC20Mintable *ERC20Mintable) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("transfer", data)
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
func (eRC20Mintable *ERC20Mintable) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("transferFrom", data)
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
func (eRC20Mintable *ERC20Mintable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (eRC20Mintable *ERC20Mintable) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20MintableApproval represents a Approval event raised by the ERC20Mintable contract.
type ERC20MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintableApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MintableApproval) ContractEventName() string {
	return ERC20MintableApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Mintable *ERC20Mintable) UnpackApprovalEvent(log *types.Log) (*ERC20MintableApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableApproval)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Mintable contract.
type ERC20MintableEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintableEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintableEIP712DomainChanged) ContractEventName() string {
	return ERC20MintableEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Mintable *ERC20Mintable) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MintableEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableForgeAdded represents a ForgeAdded event raised by the ERC20Mintable contract.
type ERC20MintableForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MintableForgeAdded) ContractEventName() string {
	return ERC20MintableForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Mintable *ERC20Mintable) UnpackForgeAddedEvent(log *types.Log) (*ERC20MintableForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableForgeRemoved represents a ForgeRemoved event raised by the ERC20Mintable contract.
type ERC20MintableForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MintableForgeRemoved) ContractEventName() string {
	return ERC20MintableForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Mintable *ERC20Mintable) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MintableForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableInitialized represents a Initialized event raised by the ERC20Mintable contract.
type ERC20MintableInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintableInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ERC20MintableInitialized) ContractEventName() string {
	return ERC20MintableInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (eRC20Mintable *ERC20Mintable) UnpackInitializedEvent(log *types.Log) (*ERC20MintableInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableInitialized)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Mintable contract.
type ERC20MintableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ERC20MintableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ERC20MintableOwnershipTransferred) ContractEventName() string {
	return ERC20MintableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eRC20Mintable *ERC20Mintable) UnpackOwnershipTransferredEvent(log *types.Log) (*ERC20MintableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableTransfer represents a Transfer event raised by the ERC20Mintable contract.
type ERC20MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MintableTransfer) ContractEventName() string {
	return ERC20MintableTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Mintable *ERC20Mintable) UnpackTransferEvent(log *types.Log) (*ERC20MintableTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableUpgraded represents a Upgraded event raised by the ERC20Mintable contract.
type ERC20MintableUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintableUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (ERC20MintableUpgraded) ContractEventName() string {
	return ERC20MintableUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (eRC20Mintable *ERC20Mintable) UnpackUpgradedEvent(log *types.Log) (*ERC20MintableUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableUpgraded)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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
func (eRC20Mintable *ERC20Mintable) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MintableAddressEmptyCode represents a AddressEmptyCode error raised by the ERC20Mintable contract.
type ERC20MintableAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func ERC20MintableAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (eRC20Mintable *ERC20Mintable) UnpackAddressEmptyCodeError(raw []byte) (*ERC20MintableAddressEmptyCode, error) {
	out := new(ERC20MintableAddressEmptyCode)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MintableECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MintableECDSAInvalidSignature, error) {
	out := new(ERC20MintableECDSAInvalidSignature)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MintableECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MintableECDSAInvalidSignatureLength, error) {
	out := new(ERC20MintableECDSAInvalidSignatureLength)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MintableECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MintableECDSAInvalidSignatureS, error) {
	out := new(ERC20MintableECDSAInvalidSignatureS)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the ERC20Mintable contract.
type ERC20MintableERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func ERC20MintableERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (eRC20Mintable *ERC20Mintable) UnpackERC1967InvalidImplementationError(raw []byte) (*ERC20MintableERC1967InvalidImplementation, error) {
	out := new(ERC20MintableERC1967InvalidImplementation)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC1967NonPayable represents a ERC1967NonPayable error raised by the ERC20Mintable contract.
type ERC20MintableERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func ERC20MintableERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (eRC20Mintable *ERC20Mintable) UnpackERC1967NonPayableError(raw []byte) (*ERC20MintableERC1967NonPayable, error) {
	out := new(ERC20MintableERC1967NonPayable)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Mintable contract.
type ERC20MintableERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MintableERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MintableERC20InsufficientAllowance, error) {
	out := new(ERC20MintableERC20InsufficientAllowance)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Mintable contract.
type ERC20MintableERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MintableERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MintableERC20InsufficientBalance, error) {
	out := new(ERC20MintableERC20InsufficientBalance)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MintableERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MintableERC20InvalidApprover, error) {
	out := new(ERC20MintableERC20InvalidApprover)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MintableERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MintableERC20InvalidReceiver, error) {
	out := new(ERC20MintableERC20InvalidReceiver)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MintableERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MintableERC20InvalidSender, error) {
	out := new(ERC20MintableERC20InvalidSender)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MintableERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MintableERC20InvalidSpender, error) {
	out := new(ERC20MintableERC20InvalidSpender)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Mintable contract.
type ERC20MintableERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MintableERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Mintable *ERC20Mintable) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MintableERC2612ExpiredSignature, error) {
	out := new(ERC20MintableERC2612ExpiredSignature)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Mintable contract.
type ERC20MintableERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MintableERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Mintable *ERC20Mintable) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MintableERC2612InvalidSigner, error) {
	out := new(ERC20MintableERC2612InvalidSigner)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableFailedCall represents a FailedCall error raised by the ERC20Mintable contract.
type ERC20MintableFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func ERC20MintableFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (eRC20Mintable *ERC20Mintable) UnpackFailedCallError(raw []byte) (*ERC20MintableFailedCall, error) {
	out := new(ERC20MintableFailedCall)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Mintable contract.
type ERC20MintableInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MintableInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Mintable *ERC20Mintable) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MintableInvalidAccountNonce, error) {
	out := new(ERC20MintableInvalidAccountNonce)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableInvalidInitialization represents a InvalidInitialization error raised by the ERC20Mintable contract.
type ERC20MintableInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ERC20MintableInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (eRC20Mintable *ERC20Mintable) UnpackInvalidInitializationError(raw []byte) (*ERC20MintableInvalidInitialization, error) {
	out := new(ERC20MintableInvalidInitialization)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableNotInitializing represents a NotInitializing error raised by the ERC20Mintable contract.
type ERC20MintableNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ERC20MintableNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (eRC20Mintable *ERC20Mintable) UnpackNotInitializingError(raw []byte) (*ERC20MintableNotInitializing, error) {
	out := new(ERC20MintableNotInitializing)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ERC20Mintable contract.
type ERC20MintableOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ERC20MintableOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eRC20Mintable *ERC20Mintable) UnpackOwnableInvalidOwnerError(raw []byte) (*ERC20MintableOwnableInvalidOwner, error) {
	out := new(ERC20MintableOwnableInvalidOwner)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ERC20Mintable contract.
type ERC20MintableOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ERC20MintableOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eRC20Mintable *ERC20Mintable) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ERC20MintableOwnableUnauthorizedAccount, error) {
	out := new(ERC20MintableOwnableUnauthorizedAccount)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Mintable contract.
type ERC20MintableTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MintableTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Mintable *ERC20Mintable) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MintableTokenBaseNullInput, error) {
	out := new(ERC20MintableTokenBaseNullInput)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Mintable contract.
type ERC20MintableTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MintableTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Mintable *ERC20Mintable) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MintableTokenBaseOnlyForge, error) {
	out := new(ERC20MintableTokenBaseOnlyForge)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the ERC20Mintable contract.
type ERC20MintableUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func ERC20MintableUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (eRC20Mintable *ERC20Mintable) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*ERC20MintableUUPSUnauthorizedCallContext, error) {
	out := new(ERC20MintableUUPSUnauthorizedCallContext)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the ERC20Mintable contract.
type ERC20MintableUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func ERC20MintableUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (eRC20Mintable *ERC20Mintable) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*ERC20MintableUUPSUnsupportedProxiableUUID, error) {
	out := new(ERC20MintableUUPSUnsupportedProxiableUUID)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
