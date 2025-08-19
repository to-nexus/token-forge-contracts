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

// TokenFactoryMetaData contains all meta data concerning the TokenFactory contract.
var TokenFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"logic\",\"type\":\"address\"}],\"name\":\"deployERC1155\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"logic\",\"type\":\"address\"}],\"name\":\"deployERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"logic\",\"type\":\"address\"}],\"name\":\"deployERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"name\":\"getPresetLogics\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"erc20Impls\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"erc721Impls\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"erc1155Impls\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"logicAddress\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setPresetLogics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"logicAddress\",\"type\":\"address\"}],\"name\":\"PresetLogicRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"logicAddress\",\"type\":\"address\"}],\"name\":\"PresetLogicSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"logicAddress\",\"type\":\"address\"}],\"name\":\"TokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenFactory__DeployFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumITokenFactory.TokenType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenFactory__InvalidLogic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenFactory__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "TokenFactory",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051612d156100395f395f81816113ba015281816113e301526116010152612d155ff3fe608060405260043610610109575f3560e01c80634f1ef286116100a1578063a217fddf11610071578063d547741f11610057578063d547741f146103a8578063ec87621c146103c7578063f3c8bccc146103fa575f5ffd5b8063a217fddf14610340578063ad3cb1cc14610353575f5ffd5b80634f1ef2861461028a57806352d1902d1461029d57806391d14854146102b15780639c952a5014610321575f5ffd5b80632f26b1de116100dc5780632f26b1de1461020c5780632f2ff15d1461022d57806336568abe1461024c5780634b6116991461026b575f5ffd5b806301ffc9a71461010d57806319b3c5761461014157806323b4a56914610185578063248a9ca3146101b1575b5f5ffd5b348015610118575f5ffd5b5061012c610127366004611f21565b610419565b60405190151581526020015b60405180910390f35b34801561014c575f5ffd5b5061016061015b366004612064565b6104b1565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610138565b348015610190575f5ffd5b506101a461019f366004612107565b6106dd565b6040516101389190612120565b3480156101bc575f5ffd5b506101fe6101cb366004612178565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610138565b348015610217575f5ffd5b5061022b6102263660046121d7565b610759565b005b348015610238575f5ffd5b5061022b610247366004612286565b6108de565b348015610257575f5ffd5b5061022b610266366004612286565b610927565b348015610276575f5ffd5b506101606102853660046122b0565b610985565b61022b610298366004612399565b610bae565b3480156102a8575f5ffd5b506101fe610bcd565b3480156102bc575f5ffd5b5061012c6102cb366004612286565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561032c575f5ffd5b5061022b61033b3660046123f1565b610bfb565b34801561034b575f5ffd5b506101fe5f81565b34801561035e575f5ffd5b5061039b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610138919061249f565b3480156103b3575f5ffd5b5061022b6103c2366004612286565b610d49565b3480156103d2575f5ffd5b506101fe7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b348015610405575f5ffd5b506101606104143660046124c1565b610d8c565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104ab57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6104dc81610f98565b7fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f006105277fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f0485610fa5565b61056b576002846040517f18b71fe50000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b60405180910390fd5b83888888886040516024016105839493929190612635565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9712d48000000000000000000000000000000000000000000000000000000001790525161060390611f14565b61060e92919061269c565b604051809103905ff080158015610627573d5f5f3e3d5ffd5b50925073ffffffffffffffffffffffffffffffffffffffff831661067c576002846040517fddefc3810000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b60026040805173ffffffffffffffffffffffffffffffffffffffff868116825287811660208301528b16917f2f41f0ca9296aeb83faf1cc044bf64ab3dd85e0cff69ff07fc6045095cf6e749910160405180910390a3505095945050505050565b60607fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f005f836002811115610713576107136125aa565b036107285761072181610fd3565b9392505050565b600183600281111561073c5761073c6125aa565b0361074d5761072181600201610fd3565b61072181600401610fd3565b5f610762610fdf565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561078e5750825b90505f8267ffffffffffffffff1660011480156107aa5750303b155b9050811580156107b8575080155b156107ef576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156108505784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610858611007565b610860611007565b61086f8c8c8c8c8c8c8c611011565b83156108d05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461091781610f98565b61092183836111a8565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610976576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61098082826112c6565b505050565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c6109b081610f98565b7fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f006109fb7fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f0285610fa5565b610a36576001846040517f18b71fe50000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b838a8a8a8a8a8a604051602401610a52969594939291906126d2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe346306b0000000000000000000000000000000000000000000000000000000017905251610ad290611f14565b610add92919061269c565b604051809103905ff080158015610af6573d5f5f3e3d5ffd5b50925073ffffffffffffffffffffffffffffffffffffffff8316610b4b576001846040517fddefc3810000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b60016040805173ffffffffffffffffffffffffffffffffffffffff868116825287811660208301528d16917f2f41f0ca9296aeb83faf1cc044bf64ab3dd85e0cff69ff07fc6045095cf6e749910160405180910390a35050979650505050505050565b610bb66113a2565b610bbf826114a6565b610bc982826114b0565b5050565b5f610bd66115e9565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610c2581610f98565b5f8080876002811115610c3a57610c3a6125aa565b03610c8957507fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f0090507f89145aae00000000000000000000000000000000000000000000000000000000610d32565b6001876002811115610c9d57610c9d6125aa565b03610cec57507fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f0290507fda4b58e200000000000000000000000000000000000000000000000000000000610d32565b507fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f0490507f37d9f405000000000000000000000000000000000000000000000000000000005b610d40878284898989611658565b50505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d8281610f98565b61092183836112c6565b5f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610db781610f98565b7fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f00610de28185610fa5565b610e1c575f846040517f18b71fe50000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b838c8c8c8c8c8c8c8c604051602401610e3c989796959493929190612763565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fb01932270000000000000000000000000000000000000000000000000000000017905251610ebc90611f14565b610ec792919061269c565b604051809103905ff080158015610ee0573d5f5f3e3d5ffd5b50925073ffffffffffffffffffffffffffffffffffffffff8316610f34575f846040517fddefc3810000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b5f6040805173ffffffffffffffffffffffffffffffffffffffff868116825287811660208301528f16917f2f41f0ca9296aeb83faf1cc044bf64ab3dd85e0cff69ff07fc6045095cf6e749910160405180910390a350509998505050505050505050565b610fa281336118f8565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610721565b60605f6107218361199e565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006104ab565b61100f6119f7565b565b6110196119f7565b73ffffffffffffffffffffffffffffffffffffffff8716611088576040517f50bbb23f0000000000000000000000000000000000000000000000000000000081527f6f776e65720000000000000000000000000000000000000000000000000000006004820152602401610562565b6110b27faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c5f611a35565b6110dc7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c886111a8565b506110e75f886111a8565b507fd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f006111385f7f89145aae00000000000000000000000000000000000000000000000000000000838a8a6001611658565b61116b60017fda4b58e2000000000000000000000000000000000000000000000000000000008360020188886001611658565b61119e60027f37d9f405000000000000000000000000000000000000000000000000000000008360040186866001611658565b5050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166112bd575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556112593390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506104ab565b5f9150506104ab565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156112bd575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506104ab565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061146f57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166114567f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561100f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610bc981610f98565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611535575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526115329181019061280f565b60015b611583576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610562565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146115df576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610562565b6109808383611ad6565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461100f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81811561185a575f5b81811015611854575f85858381811061167c5761167c612826565b90506020020160208101906116919190612853565b905073ffffffffffffffffffffffffffffffffffffffff8116611702576040517f50bbb23f0000000000000000000000000000000000000000000000000000000081527f696d706c730000000000000000000000000000000000000000000000000000006004820152602401610562565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008916600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa15801561178c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117b0919061286c565b6117ea5788816040517f18b71fe50000000000000000000000000000000000000000000000000000000081526004016105629291906125d7565b6117f48782611b38565b1561184b578073ffffffffffffffffffffffffffffffffffffffff16896002811115611822576118226125aa565b6040517f6253639bce3b622551aefbac97837a6f9ed2fff95c665232b1f5f85572ded776905f90a35b50600101611661565b50610d40565b5f5b8181101561119e575f85858381811061187757611877612826565b905060200201602081019061188c9190612853565b90506118988782611b59565b156118ef578073ffffffffffffffffffffffffffffffffffffffff168960028111156118c6576118c66125aa565b6040517fe7a7b3980d246dc33cf9402c0e30a3219269d7ebdebda0457fc520409054921e905f90a35b5060010161185c565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610bc9576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610562565b6060815f018054806020026020016040519081016040528092919081815260200182805480156119eb57602002820191905f5260205f20905b8154815260200190600101908083116119d7575b50505050509050919050565b6119ff611b7a565b61100f576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268005f611a8e845f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b5f85815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b611adf82611b98565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611b30576109808282611c66565b610bc9611ce5565b5f6107218373ffffffffffffffffffffffffffffffffffffffff8416611d1d565b5f6107218373ffffffffffffffffffffffffffffffffffffffff8416611d69565b5f611b83610fdf565b5468010000000000000000900460ff16919050565b8073ffffffffffffffffffffffffffffffffffffffff163b5f03611c00576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610562565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611c8f9190612887565b5f60405180830381855af49150503d805f8114611cc7576040519150601f19603f3d011682016040523d82523d5f602084013e611ccc565b606091505b5091509150611cdc858383611e43565b95945050505050565b341561100f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600183016020526040812054611d6257508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556104ab565b505f6104ab565b5f81815260018301602052604081205480156112bd575f611d8b60018361289d565b85549091505f90611d9e9060019061289d565b9050808214611dfd575f865f018281548110611dbc57611dbc612826565b905f5260205f200154905080875f018481548110611ddc57611ddc612826565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611e0e57611e0e6128d5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506104ab565b606082611e5857611e5382611ed2565b610721565b8151158015611e7c575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611ecb576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610562565b5092915050565b805115611ee25780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dd8061290383390190565b5f60208284031215611f31575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610721575f5ffd5b803573ffffffffffffffffffffffffffffffffffffffff81168114611f83575f5ffd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611fc4575f5ffd5b8135602083015f5f67ffffffffffffffff841115611fe457611fe4611f88565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff8211171561203157612031611f88565b604052838152905080828401871015612048575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f5f5f5f60a08688031215612078575f5ffd5b61208186611f60565b945061208f60208701611f60565b9350604086013567ffffffffffffffff8111156120aa575f5ffd5b6120b688828901611fb5565b935050606086013567ffffffffffffffff8111156120d2575f5ffd5b6120de88828901611fb5565b9250506120ed60808701611f60565b90509295509295909350565b803560038110611f83575f5ffd5b5f60208284031215612117575f5ffd5b610721826120f9565b602080825282518282018190525f918401906040840190835b8181101561216d57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612139565b509095945050505050565b5f60208284031215612188575f5ffd5b5035919050565b5f5f83601f84011261219f575f5ffd5b50813567ffffffffffffffff8111156121b6575f5ffd5b6020830191508360208260051b85010111156121d0575f5ffd5b9250929050565b5f5f5f5f5f5f5f6080888a0312156121ed575f5ffd5b6121f688611f60565b9650602088013567ffffffffffffffff811115612211575f5ffd5b61221d8a828b0161218f565b909750955050604088013567ffffffffffffffff81111561223c575f5ffd5b6122488a828b0161218f565b909550935050606088013567ffffffffffffffff811115612267575f5ffd5b6122738a828b0161218f565b989b979a50959850939692959293505050565b5f5f60408385031215612297575f5ffd5b823591506122a760208401611f60565b90509250929050565b5f5f5f5f5f5f5f60e0888a0312156122c6575f5ffd5b6122cf88611f60565b96506122dd60208901611f60565b9550604088013567ffffffffffffffff8111156122f8575f5ffd5b6123048a828b01611fb5565b955050606088013567ffffffffffffffff811115612320575f5ffd5b61232c8a828b01611fb5565b945050608088013567ffffffffffffffff811115612348575f5ffd5b6123548a828b01611fb5565b93505060a088013567ffffffffffffffff811115612370575f5ffd5b61237c8a828b01611fb5565b92505061238b60c08901611f60565b905092959891949750929550565b5f5f604083850312156123aa575f5ffd5b6123b383611f60565b9150602083013567ffffffffffffffff8111156123ce575f5ffd5b6123da85828601611fb5565b9150509250929050565b8015158114610fa2575f5ffd5b5f5f5f5f60608587031215612404575f5ffd5b61240d856120f9565b9350602085013567ffffffffffffffff811115612428575f5ffd5b6124348782880161218f565b9094509250506040850135612448816123e4565b939692955090935050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6107216020830184612453565b803560ff81168114611f83575f5ffd5b5f5f5f5f5f5f5f5f5f6101208a8c0312156124da575f5ffd5b6124e38a611f60565b98506124f160208b01611f60565b975060408a013567ffffffffffffffff81111561250c575f5ffd5b6125188c828d01611fb5565b97505060608a013567ffffffffffffffff811115612534575f5ffd5b6125408c828d01611fb5565b96505061254f60808b016124b1565b945060a08a0135935061256460c08b01611f60565b925060e08a013567ffffffffffffffff81111561257f575f5ffd5b61258b8c828d01611fb5565b92505061259b6101008b01611f60565b90509295985092959850929598565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6040810160038410612610577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b92815273ffffffffffffffffffffffffffffffffffffffff9190911660209091015290565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201525f61267f6080830185612453565b82810360608401526126918185612453565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6126ca6040830184612453565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015260c060408201525f61271c60c0830187612453565b828103606084015261272e8187612453565b905082810360808401526127428186612453565b905082810360a08401526127568185612453565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8916815273ffffffffffffffffffffffffffffffffffffffff8816602082015261010060408201525f6127af610100830189612453565b82810360608401526127c18189612453565b905060ff871660808401528560a084015273ffffffffffffffffffffffffffffffffffffffff851660c084015282810360e08401526128008185612453565b9b9a5050505050505050505050565b5f6020828403121561281f575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215612863575f5ffd5b61072182611f60565b5f6020828403121561287c575f5ffd5b8151610721816123e4565b5f82518060208501845e5f920191825250919050565b818103818111156104ab577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfe60806040526040516103dd3803806103dd8339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea264697066735822122097b96e571989a97f4a6a6c0ed50363e8d85d9a1f35d668f039becfc1a6d19a1064736f6c634300081c0033a26469706673582212202ce48a35bc09cc69811555be9bec6e6b7f44a95b5b02dcd4aedc3ac2107f3f3b64736f6c634300081c0033",
}

// TokenFactory is an auto generated Go binding around an Ethereum contract.
type TokenFactory struct {
	abi abi.ABI
}

// NewTokenFactory creates a new instance of TokenFactory.
func NewTokenFactory() *TokenFactory {
	parsed, err := TokenFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TokenFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TokenFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (tokenFactory *TokenFactory) PackDEFAULTADMINROLE() []byte {
	enc, err := tokenFactory.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (tokenFactory *TokenFactory) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := tokenFactory.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (tokenFactory *TokenFactory) PackMANAGERROLE() []byte {
	enc, err := tokenFactory.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (tokenFactory *TokenFactory) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := tokenFactory.abi.Unpack("MANAGER_ROLE", data)
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
func (tokenFactory *TokenFactory) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := tokenFactory.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (tokenFactory *TokenFactory) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := tokenFactory.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackDeployERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19b3c576.
//
// Solidity: function deployERC1155(address owner, address manager, string uri, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) PackDeployERC1155(owner common.Address, manager common.Address, uri string, data []byte, logic common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("deployERC1155", owner, manager, uri, data, logic)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployERC1155 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19b3c576.
//
// Solidity: function deployERC1155(address owner, address manager, string uri, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) UnpackDeployERC1155(data []byte) (common.Address, error) {
	out, err := tokenFactory.abi.Unpack("deployERC1155", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3c8bccc.
//
// Solidity: function deployERC20(address owner, address manager, string name, string symbol, uint8 decimals, uint256 initialSupply, address initialRecipient, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) PackDeployERC20(owner common.Address, manager common.Address, name string, symbol string, decimals uint8, initialSupply *big.Int, initialRecipient common.Address, data []byte, logic common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("deployERC20", owner, manager, name, symbol, decimals, initialSupply, initialRecipient, data, logic)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployERC20 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf3c8bccc.
//
// Solidity: function deployERC20(address owner, address manager, string name, string symbol, uint8 decimals, uint256 initialSupply, address initialRecipient, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) UnpackDeployERC20(data []byte) (common.Address, error) {
	out, err := tokenFactory.abi.Unpack("deployERC20", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4b611699.
//
// Solidity: function deployERC721(address owner, address manager, string name, string symbol, string baseTokenURI, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) PackDeployERC721(owner common.Address, manager common.Address, name string, symbol string, baseTokenURI string, data []byte, logic common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("deployERC721", owner, manager, name, symbol, baseTokenURI, data, logic)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployERC721 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4b611699.
//
// Solidity: function deployERC721(address owner, address manager, string name, string symbol, string baseTokenURI, bytes data, address logic) returns(address token)
func (tokenFactory *TokenFactory) UnpackDeployERC721(data []byte) (common.Address, error) {
	out, err := tokenFactory.abi.Unpack("deployERC721", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetPresetLogics is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b4a569.
//
// Solidity: function getPresetLogics(uint8 tokenType) view returns(address[])
func (tokenFactory *TokenFactory) PackGetPresetLogics(tokenType uint8) []byte {
	enc, err := tokenFactory.abi.Pack("getPresetLogics", tokenType)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPresetLogics is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b4a569.
//
// Solidity: function getPresetLogics(uint8 tokenType) view returns(address[])
func (tokenFactory *TokenFactory) UnpackGetPresetLogics(data []byte) ([]common.Address, error) {
	out, err := tokenFactory.abi.Unpack("getPresetLogics", data)
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
func (tokenFactory *TokenFactory) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := tokenFactory.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (tokenFactory *TokenFactory) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := tokenFactory.abi.Unpack("getRoleAdmin", data)
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
func (tokenFactory *TokenFactory) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (tokenFactory *TokenFactory) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (tokenFactory *TokenFactory) UnpackHasRole(data []byte) (bool, error) {
	out, err := tokenFactory.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f26b1de.
//
// Solidity: function initialize(address owner, address[] erc20Impls, address[] erc721Impls, address[] erc1155Impls) returns()
func (tokenFactory *TokenFactory) PackInitialize(owner common.Address, erc20Impls []common.Address, erc721Impls []common.Address, erc1155Impls []common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("initialize", owner, erc20Impls, erc721Impls, erc1155Impls)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenFactory *TokenFactory) PackProxiableUUID() []byte {
	enc, err := tokenFactory.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (tokenFactory *TokenFactory) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := tokenFactory.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (tokenFactory *TokenFactory) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (tokenFactory *TokenFactory) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := tokenFactory.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPresetLogics is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c952a50.
//
// Solidity: function setPresetLogics(uint8 tokenType, address[] logicAddress, bool add) returns()
func (tokenFactory *TokenFactory) PackSetPresetLogics(tokenType uint8, logicAddress []common.Address, add bool) []byte {
	enc, err := tokenFactory.abi.Pack("setPresetLogics", tokenType, logicAddress, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenFactory *TokenFactory) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := tokenFactory.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (tokenFactory *TokenFactory) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := tokenFactory.abi.Unpack("supportsInterface", data)
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
func (tokenFactory *TokenFactory) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := tokenFactory.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TokenFactoryInitialized represents a Initialized event raised by the TokenFactory contract.
type TokenFactoryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenFactoryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TokenFactoryInitialized) ContractEventName() string {
	return TokenFactoryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (tokenFactory *TokenFactory) UnpackInitializedEvent(log *types.Log) (*TokenFactoryInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryInitialized)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryPresetLogicRemoved represents a PresetLogicRemoved event raised by the TokenFactory contract.
type TokenFactoryPresetLogicRemoved struct {
	TokenType    uint8
	LogicAddress common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const TokenFactoryPresetLogicRemovedEventName = "PresetLogicRemoved"

// ContractEventName returns the user-defined event name.
func (TokenFactoryPresetLogicRemoved) ContractEventName() string {
	return TokenFactoryPresetLogicRemovedEventName
}

// UnpackPresetLogicRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PresetLogicRemoved(uint8 indexed tokenType, address indexed logicAddress)
func (tokenFactory *TokenFactory) UnpackPresetLogicRemovedEvent(log *types.Log) (*TokenFactoryPresetLogicRemoved, error) {
	event := "PresetLogicRemoved"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryPresetLogicRemoved)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryPresetLogicSet represents a PresetLogicSet event raised by the TokenFactory contract.
type TokenFactoryPresetLogicSet struct {
	TokenType    uint8
	LogicAddress common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const TokenFactoryPresetLogicSetEventName = "PresetLogicSet"

// ContractEventName returns the user-defined event name.
func (TokenFactoryPresetLogicSet) ContractEventName() string {
	return TokenFactoryPresetLogicSetEventName
}

// UnpackPresetLogicSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PresetLogicSet(uint8 indexed tokenType, address indexed logicAddress)
func (tokenFactory *TokenFactory) UnpackPresetLogicSetEvent(log *types.Log) (*TokenFactoryPresetLogicSet, error) {
	event := "PresetLogicSet"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryPresetLogicSet)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the TokenFactory contract.
type TokenFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const TokenFactoryRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (TokenFactoryRoleAdminChanged) ContractEventName() string {
	return TokenFactoryRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (tokenFactory *TokenFactory) UnpackRoleAdminChangedEvent(log *types.Log) (*TokenFactoryRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryRoleGranted represents a RoleGranted event raised by the TokenFactory contract.
type TokenFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenFactoryRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (TokenFactoryRoleGranted) ContractEventName() string {
	return TokenFactoryRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (tokenFactory *TokenFactory) UnpackRoleGrantedEvent(log *types.Log) (*TokenFactoryRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryRoleGranted)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryRoleRevoked represents a RoleRevoked event raised by the TokenFactory contract.
type TokenFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenFactoryRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (TokenFactoryRoleRevoked) ContractEventName() string {
	return TokenFactoryRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (tokenFactory *TokenFactory) UnpackRoleRevokedEvent(log *types.Log) (*TokenFactoryRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryRoleRevoked)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryTokenDeployed represents a TokenDeployed event raised by the TokenFactory contract.
type TokenFactoryTokenDeployed struct {
	Owner        common.Address
	TokenType    uint8
	TokenAddress common.Address
	LogicAddress common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const TokenFactoryTokenDeployedEventName = "TokenDeployed"

// ContractEventName returns the user-defined event name.
func (TokenFactoryTokenDeployed) ContractEventName() string {
	return TokenFactoryTokenDeployedEventName
}

// UnpackTokenDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenDeployed(address indexed owner, uint8 indexed tokenType, address tokenAddress, address logicAddress)
func (tokenFactory *TokenFactory) UnpackTokenDeployedEvent(log *types.Log) (*TokenFactoryTokenDeployed, error) {
	event := "TokenDeployed"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryTokenDeployed)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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

// TokenFactoryUpgraded represents a Upgraded event raised by the TokenFactory contract.
type TokenFactoryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TokenFactoryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TokenFactoryUpgraded) ContractEventName() string {
	return TokenFactoryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (tokenFactory *TokenFactory) UnpackUpgradedEvent(log *types.Log) (*TokenFactoryUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != tokenFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenFactoryUpgraded)
	if len(log.Data) > 0 {
		if err := tokenFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenFactory.abi.Events[event].Inputs {
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
func (tokenFactory *TokenFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["TokenFactoryDeployFailed"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackTokenFactoryDeployFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["TokenFactoryInvalidLogic"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackTokenFactoryInvalidLogicError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["TokenFactoryZeroAddress"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackTokenFactoryZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], tokenFactory.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return tokenFactory.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TokenFactoryAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the TokenFactory contract.
type TokenFactoryAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func TokenFactoryAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (tokenFactory *TokenFactory) UnpackAccessControlBadConfirmationError(raw []byte) (*TokenFactoryAccessControlBadConfirmation, error) {
	out := new(TokenFactoryAccessControlBadConfirmation)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the TokenFactory contract.
type TokenFactoryAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func TokenFactoryAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (tokenFactory *TokenFactory) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*TokenFactoryAccessControlUnauthorizedAccount, error) {
	out := new(TokenFactoryAccessControlUnauthorizedAccount)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryAddressEmptyCode represents a AddressEmptyCode error raised by the TokenFactory contract.
type TokenFactoryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TokenFactoryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (tokenFactory *TokenFactory) UnpackAddressEmptyCodeError(raw []byte) (*TokenFactoryAddressEmptyCode, error) {
	out := new(TokenFactoryAddressEmptyCode)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TokenFactory contract.
type TokenFactoryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TokenFactoryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (tokenFactory *TokenFactory) UnpackERC1967InvalidImplementationError(raw []byte) (*TokenFactoryERC1967InvalidImplementation, error) {
	out := new(TokenFactoryERC1967InvalidImplementation)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryERC1967NonPayable represents a ERC1967NonPayable error raised by the TokenFactory contract.
type TokenFactoryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TokenFactoryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (tokenFactory *TokenFactory) UnpackERC1967NonPayableError(raw []byte) (*TokenFactoryERC1967NonPayable, error) {
	out := new(TokenFactoryERC1967NonPayable)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryFailedCall represents a FailedCall error raised by the TokenFactory contract.
type TokenFactoryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TokenFactoryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (tokenFactory *TokenFactory) UnpackFailedCallError(raw []byte) (*TokenFactoryFailedCall, error) {
	out := new(TokenFactoryFailedCall)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryInvalidInitialization represents a InvalidInitialization error raised by the TokenFactory contract.
type TokenFactoryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TokenFactoryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (tokenFactory *TokenFactory) UnpackInvalidInitializationError(raw []byte) (*TokenFactoryInvalidInitialization, error) {
	out := new(TokenFactoryInvalidInitialization)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryNotInitializing represents a NotInitializing error raised by the TokenFactory contract.
type TokenFactoryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TokenFactoryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (tokenFactory *TokenFactory) UnpackNotInitializingError(raw []byte) (*TokenFactoryNotInitializing, error) {
	out := new(TokenFactoryNotInitializing)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryTokenFactoryDeployFailed represents a TokenFactory__DeployFailed error raised by the TokenFactory contract.
type TokenFactoryTokenFactoryDeployFailed struct {
	Arg0 uint8
	Arg1 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenFactory__DeployFailed(uint8 arg0, address arg1)
func TokenFactoryTokenFactoryDeployFailedErrorID() common.Hash {
	return common.HexToHash("0xddefc38196a03f87021c761b50d14b1107b57389499550de04e97cf32cd0fdfd")
}

// UnpackTokenFactoryDeployFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenFactory__DeployFailed(uint8 arg0, address arg1)
func (tokenFactory *TokenFactory) UnpackTokenFactoryDeployFailedError(raw []byte) (*TokenFactoryTokenFactoryDeployFailed, error) {
	out := new(TokenFactoryTokenFactoryDeployFailed)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "TokenFactoryDeployFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryTokenFactoryInvalidLogic represents a TokenFactory__InvalidLogic error raised by the TokenFactory contract.
type TokenFactoryTokenFactoryInvalidLogic struct {
	Arg0 uint8
	Arg1 common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenFactory__InvalidLogic(uint8 arg0, address arg1)
func TokenFactoryTokenFactoryInvalidLogicErrorID() common.Hash {
	return common.HexToHash("0x18b71fe536a3586382120d2a7d20c6a0b22c27a408aae1b56e0e5c307336b046")
}

// UnpackTokenFactoryInvalidLogicError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenFactory__InvalidLogic(uint8 arg0, address arg1)
func (tokenFactory *TokenFactory) UnpackTokenFactoryInvalidLogicError(raw []byte) (*TokenFactoryTokenFactoryInvalidLogic, error) {
	out := new(TokenFactoryTokenFactoryInvalidLogic)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "TokenFactoryInvalidLogic", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryTokenFactoryZeroAddress represents a TokenFactory__ZeroAddress error raised by the TokenFactory contract.
type TokenFactoryTokenFactoryZeroAddress struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenFactory__ZeroAddress(bytes32 field)
func TokenFactoryTokenFactoryZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x50bbb23fc3c9caa34b10360e46f232db41fff60895c36a03bfaa0b875c39300a")
}

// UnpackTokenFactoryZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenFactory__ZeroAddress(bytes32 field)
func (tokenFactory *TokenFactory) UnpackTokenFactoryZeroAddressError(raw []byte) (*TokenFactoryTokenFactoryZeroAddress, error) {
	out := new(TokenFactoryTokenFactoryZeroAddress)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "TokenFactoryZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TokenFactory contract.
type TokenFactoryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TokenFactoryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (tokenFactory *TokenFactory) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TokenFactoryUUPSUnauthorizedCallContext, error) {
	out := new(TokenFactoryUUPSUnauthorizedCallContext)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TokenFactoryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TokenFactory contract.
type TokenFactoryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TokenFactoryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (tokenFactory *TokenFactory) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TokenFactoryUUPSUnsupportedProxiableUUID, error) {
	out := new(TokenFactoryUUPSUnsupportedProxiableUUID)
	if err := tokenFactory.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}