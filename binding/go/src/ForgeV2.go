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

// ForgeV2MetaData contains all meta data concerning the ForgeV2 contract.
var ForgeV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"burnERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"BaseForge__InvalidFeeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidPermitSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC1155ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC721ForgeV2__InvalidAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ForgeV2",
	Bin: "0x6080806040523460aa575f5160206142395f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b60405161418a90816100af8239f35b6001600160401b0319166001600160401b039081175f5160206142395f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a71461291957508063150b7a021461288d5780632af4c9d0146128685780632f2ddc1f1461258a5780633081f746146122555780633644e5151461221457806338033a8b14611edf5780633a5381b514611e6e57806343dc7f7814611bce57806360cfd5011461192857806372033c8c146116665780637ecebe00146115e457806382a9bb38146115c057806384b0196e146112f85780638abb7071146110db578063939cefc1146110c1578063a84adb3214610d7f578063b55ba9c414610d48578063ba32f859146109d8578063bc197c8114610909578063bf6bdd9714610623578063db686a78146101ac5763f23a6e611461011b575f80fd5b346101a95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576101526129d5565b5061015b6129f8565b5060843567ffffffffffffffff81116101a75761017c903690600401612af9565b5060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b505b80fd5b50346101a95760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576101e46129d5565b906101ed6129f8565b6044359260643567ffffffffffffffff811161061f57610211903690600401612bf1565b60849591953560a43567ffffffffffffffff811161061b57610237903690600401612bf1565b60c43567ffffffffffffffff811161061757899493926102c39261035d61034b8c8c6102cc61026b8e983690600401612bf1565b9990986102778c612fec565b73ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9b8b3691612ac3565b602081519101209873ffffffffffffffffffffffffffffffffffffffff604051926020840194507f8c9fe637b6999b95af5c3d94f9198143c67fc87f05c92635e2623e9d030902e4855216998a60408401528d606084015260808301528b60a083015260c082015260c0815261034360e082612a1b565b519020613368565b610356368486612ac3565b908b6133a9565b156105d5576084896103f48d9e966103ee8f95976103e66103848f9960209d9b3691612ac3565b8c81519101209560405173ffffffffffffffffffffffffffffffffffffffff8e8201927f2dd43c6051e2af91d874ee86be1d3d1dabd6562aa4f73852cebc8779fcbfc64984521697886040830152606082015260608152610343608082612a1b565b923691612ac3565b906134f5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f896040519a8b998a9889967f94d008ef00000000000000000000000000000000000000000000000000000000885260048801526024870152606060448701528160648701528686013785858286010152011681010301925af180156105ca57610594575b506104896104e59184613595565b6040805173ffffffffffffffffffffffffffffffffffffffff909516602086015284019290925290919081606081015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612a1b565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156105905761056393858094604051968795869485937fe98a57840000000000000000000000000000000000000000000000000000000085526001600486016135dc565b03925af18015610585576105745750f35b8161057e91612a1b565b6101a95780f35b6040513d84823e3d90fd5b8480fd5b6020813d6020116105c2575b816105ad60209383612a1b565b810103126105be575061048961047b565b8580fd5b3d91506105a0565b6040513d88823e3d90fd5b60248b73ffffffffffffffffffffffffffffffffffffffff8b7ffb1b8f0f00000000000000000000000000000000000000000000000000000000835216600452fd5b8880fd5b8680fd5b8380fd5b50346101a95761063236612def565b909282936106438799969794612fec565b61071f816107198c8b61069c8b73ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b998751602089012073ffffffffffffffffffffffffffffffffffffffff604051936020850195507f71afff252ff99c81e2e331f598fa5b6726b8fede3187b1e76ccddde23814194a8652169a8b604085015260608401528d608084015260a08301528a60c083015260e082015260e0815261034361010082612a1b565b886133a9565b156108c757906107998260208c969594519101209261079460405160208101907ffb69794a96dc00d8da302bbdbf1b9af3b713bfeb066c705c1f55ba40dd7eabed825273ffffffffffffffffffffffffffffffffffffffff8b1696876040830152606082015260608152610343608082612a1b565b6134f5565b823b1561061f576107f69289858094604051968795869485937f731133e9000000000000000000000000000000000000000000000000000000008552600485015260248401528c6044840152608060648401526084830190612d79565b03925af18015610585576108aa575b506108a761087686866108a28a61085a896104b96108238b83613595565b956040519485936020850160409194939273ffffffffffffffffffffffffffffffffffffffff606083019616825260208201520152565b6040519485918860208401526040808401526060830190612d79565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284612a1b565b613737565b80f35b816108b9919695949396612a1b565b6105be57909192855f610805565b60248a73ffffffffffffffffffffffffffffffffffffffff887fa424a94100000000000000000000000000000000000000000000000000000000835216600452fd5b50346101a95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576109416129d5565b5061094a6129f8565b5060443567ffffffffffffffff81116101a75761096b903690600401612c1f565b5060643567ffffffffffffffff81116101a75761098c903690600401612c1f565b5060843567ffffffffffffffff81116101a7576109ad903690600401612af9565b5060206040517fbc197c81000000000000000000000000000000000000000000000000000000008152f35b50346101a95760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957610a106129d5565b90610a196129f8565b6044359260643567ffffffffffffffff811161061f57610a3d903690600401612af9565b9360843560a43567ffffffffffffffff81116105be57610a61903690600401612bf1565b919060c43567ffffffffffffffff8111610d0257610a83903690600401612bf1565b989091610a8f84612fec565b610adf8773ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b94610b71610b5f8b8b865160208801209873ffffffffffffffffffffffffffffffffffffffff604051926020840194507f3babb3dee300cb5c1296b0da361f7f00e8b4572bd64f6e0ecd782c87c264f830855216998a60408401528c606084015260808301528a60a083015260c082015260c0815261034360e082612a1b565b610b6a368486612ac3565b908a6133a9565b15610d0657996103ee610b8a610bf2938c9d3691612ac3565b60208151910120946103e660405160208101907f7a638779cbc1a1dd3218e65032adb415412015122b2788070739a2ace3509a31825273ffffffffffffffffffffffffffffffffffffffff8d1698896040830152606082015260608152610343608082612a1b565b823b15610d0257610c4e928892836040518096819582947fb88d4fde00000000000000000000000000000000000000000000000000000000845230600485015260248401528a6044840152608060648401526084830190612d79565b03925af19081156105ca578691610ced575b5050610489610c6f9184613595565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156105905761056393858094604051968795869485937f0b7e40c60000000000000000000000000000000000000000000000000000000085526001600486016135dc565b81610cf791612a1b565b61059057845f610c60565b8780fd5b60248a73ffffffffffffffffffffffffffffffffffffffff8a7ffb1b8f0f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346101a9576108a7610d7a610d5d36612e84565b610d718498949a929a999399979597612fec565b87858486613c87565b6138c3565b50346101a95760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957610db76129d5565b610dbf6129f8565b60443567ffffffffffffffff811161061f57610ddf903690600401612c1f565b9160643567ffffffffffffffff811161059057610e00903690600401612c1f565b9060843560a43567ffffffffffffffff811161061b57610e24903690600401612af9565b9060c43567ffffffffffffffff8111610d0257610e45903690600401612af9565b610e4e82612fec565b610e9e8473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b92610f4e81610f48898b604051610ebd816104b9602082018095612f7e565b519020968a604051610ed7816104b9602082018095612f7e565b519020906040519173ffffffffffffffffffffffffffffffffffffffff60208401947fda6dbf9c08df397da92be8fda42638fad009adf2a495b08f3526089ad02e7dd1865216998a6040850152606084015260808301528960a083015260c082015260c0815261034360e082612a1b565b876133a9565b1561107f57889291610794826020610fbf945191012060405160208101917fa221f8b9ccdf4675afa9893d3a99032eb17f5b0c291026fd55d97bb9993fa4e3835273ffffffffffffffffffffffffffffffffffffffff8a166040830152606082015260608152610343608082612a1b565b803b156101a7576040517fe08ba4bb00000000000000000000000000000000000000000000000000000000815290829082908183816110038b8e8c60048501612fab565b03925af1801561058557611062575b506108a7610876868661105d8a611040896104b96110308b83613595565b9560405194859360208501612fab565b604051948591600160208401526040808401526060830190612d79565b613620565b81611071919695949396612a1b565b6105be57909192855f611012565b60248973ffffffffffffffffffffffffffffffffffffffff877fa424a94100000000000000000000000000000000000000000000000000000000835216600452fd5b50346101a9576108a76110d6610d5d36612e84565b613021565b50346101a9576110ea36612def565b909282936110fb8799969794612fec565b6111d1816107198c8b6111548b73ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b998751602089012073ffffffffffffffffffffffffffffffffffffffff604051936020850195507f3991a11da0d98ebde4668c7c2716088d93b9497b88d2b4cda98f19cfc97acd528652169a8b604085015260608401528d608084015260a08301528a60c083015260e082015260e0815261034361010082612a1b565b156108c757906112468260208c969594519101209261079460405160208101907fadcb1a40c0ace98e86ce9038ac7bb2e74cec43bdaa9e5727e2ea08c07a1f36dd825273ffffffffffffffffffffffffffffffffffffffff8b1696876040830152606082015260608152610343608082612a1b565b823b1561061f576112a99289858094604051968795869485937ff242432a000000000000000000000000000000000000000000000000000000008552306004860152602485015260448401528c606484015260a0608484015260a4830190612d79565b03925af18015610585576112db575b506108a761087686866112d68a61085a896104b96108238b83613595565b6137b7565b816112ea919695949396612a1b565b6105be57909192855f6112b8565b50346101a957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580611597575b1561153957604051817fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102548060011c9060018116801561152f575b602083108114611502578285529081156114c05750600114611448575b50611417926113bc83611444930384612a1b565b6113c4613b32565b611425604051916113d6602084612a1b565b8383525f3681376040519687967f0f00000000000000000000000000000000000000000000000000000000000000885260e0602089015260e0880190612d79565b908682036040880152612d79565b9146606086015230608086015260a085015283820360c0850152612dbc565b0390f35b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10284528391507f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106114a657505081016020016114176113a8565b600181602092949394548385880101520191019190611490565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915061141790506113a8565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f169161138b565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415611350565b50346101a9576108a76115d236612b17565b96610d7a839693979297959495612fec565b50346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957604060209173ffffffffffffffffffffffffffffffffffffffff6116366129d5565b1681527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0083522054604051908152f35b50346101a95761167536612b17565b61168783999899969596949294612fec565b6116d78973ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b936117718161176b60405160208101907fab25e47ab9e13a956bdee490006a87ac9cbe78ae734a45c431fe3938cbd08581825273ffffffffffffffffffffffffffffffffffffffff8d16988960408301528b606083015273ffffffffffffffffffffffffffffffffffffffff8d1660808301528860a08301528a60c083015260e082015260e0815261034361010082612a1b565b8c6133a9565b156118e657611858948a946117fd886104b9966117f78e9f97610794611807998960206118189b519101206040519073ffffffffffffffffffffffffffffffffffffffff60208301937f7ae019cd8a06bf0d82d3d0063c026c8cfe203e2686c8de51ae89d5645e022d618552166040830152606082015260608152610343608082612a1b565b8b6136bd565b8783969296613837565b8389816118d5575b50505084613595565b95604051958694602086019094939273ffffffffffffffffffffffffffffffffffffffff9060609382608085019816845260208401521660408201520152565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156105905761056393858094604051968795869485937f0b7e40c600000000000000000000000000000000000000000000000000000000855285600486016135dc565b6118de92613837565b5f838961180f565b60248973ffffffffffffffffffffffffffffffffffffffff8c7f9c0f933100000000000000000000000000000000000000000000000000000000835216600452fd5b50346101a95761193736612c88565b90926119468398959698612fec565b6119968573ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b93611a4f816107198c8b8d6040516119b6816104b9602082018095612f7e565b519020988c6040516119d0816104b9602082018095612f7e565b519020885160208a01209073ffffffffffffffffffffffffffffffffffffffff604051946020860196507f53b86f7b5453d9de3449a5a75167b5d95aaf0968acc8570e67d9cb44e4c0a09d8752169b8c60408601526060850152608084015260a08301528a60c083015260e082015260e0815261034361010082612a1b565b156108c75790611ac48260208c969594519101209261079460405160208101907fbe51d3aafe3852e9cbe2ce5b4d184540bcb68868fc9307bf8b45d1f170d6739f825273ffffffffffffffffffffffffffffffffffffffff8b1696876040830152606082015260608152610343608082612a1b565b823b1561061f57869289858094611b84611b2495611b54604051998a98899788967f2eb2c2d6000000000000000000000000000000000000000000000000000000008852306004890152602488015260a0604488015260a4870190612dbc565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016064870152612dbc565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016084850152612d79565b03925af1801561058557611bb1575b506108a761087686866112d68a611040896104b96110308b83613595565b81611bc0919695949396612a1b565b6105be57909192855f611b93565b50346101a957611bdd36612c88565b9092611bec8398959698612fec565b611c3c8573ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b93611cf5816107198c8b8d604051611c5c816104b9602082018095612f7e565b519020988c604051611c76816104b9602082018095612f7e565b519020885160208a01209073ffffffffffffffffffffffffffffffffffffffff604051946020860196507fae67447a4abe000721c47402f895d79b71458082b87e1a69d5837b6241fed4858752169b8c60408601526060850152608084015260a08301528a60c083015260e082015260e0815261034361010082612a1b565b156108c75790611d6a8260208c969594519101209261079460405160208101907f811c2af50c1ad6d46b121ba5a43024be13e428134866be726da388c2357856a5825273ffffffffffffffffffffffffffffffffffffffff8b1696876040830152606082015260608152610343608082612a1b565b823b1561061f57869289858094611e24611dc495611df4604051998a98899788967f1f7fdffa0000000000000000000000000000000000000000000000000000000088526004880152608060248801526084870190612dbc565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016044870152612dbc565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152612d79565b03925af1801561058557611e51575b506108a761087686866108a28a611040896104b96110308b83613595565b81611e60919695949396612a1b565b6105be57909192855f611e33565b50346101a957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602073ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015416604051908152f35b50346101a95760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957611f176129d5565b90611f206129f8565b6064359260443560843567ffffffffffffffff811161059057611f47903690600401612bf1565b909560a43567ffffffffffffffff811161061b57611f69903690600401612bf1565b97611f7383612fec565b611fc38673ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9361204561203360405160208101907f4e6479369d4c42395619be8224f096421cd6800a840e701655040f5a933c33cd825273ffffffffffffffffffffffffffffffffffffffff8c16978860408301528a606083015289608083015260a082015260a0815261034360c082612a1b565b61203e368486612ac3565b90896133a9565b156121d2576120c4926103e66120628b9c946103ee943691612ac3565b6020815191012060405160208101917f42dec6f4478d32028294c542bb3ee6cb16fb5523fd8f13922dca479fe7647513835273ffffffffffffffffffffffffffffffffffffffff8c166040830152606082015260608152610343608082612a1b565b803b156105be576040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018490529086908290604490829084905af19081156105ca5786916121bd575b505061048961213f9184613595565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156105905761056393858094604051968795869485937f3ea113d00000000000000000000000000000000000000000000000000000000085526001600486016135dc565b816121c791612a1b565b61059057845f612130565b60248973ffffffffffffffffffffffffffffffffffffffff897ffb1b8f0f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346101a957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602061224d613e11565b604051908152f35b50346101a95761226436612b17565b61227683949293999899969596612fec565b6122c68973ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9261235a8161176b60405160208101907f9eff2bc106f16f11dab69de8c466346a0a369dd1d3a0878683947093830994fa825273ffffffffffffffffffffffffffffffffffffffff8d16998a60408301528b606083015273ffffffffffffffffffffffffffffffffffffffff8d1660808301528860a08301528960c083015260e082015260e0815261034361010082612a1b565b156118e657889993926123d16123d7936107948460208b965191012060405160208101917f340f6ac2b427cb47a7172eaa76879e09534311918b4cff11daea1028ca325aa8835273ffffffffffffffffffffffffffffffffffffffff8b166040830152606082015260608152610343608082612a1b565b876136bd565b8491943b15610617576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810191909152888160448183865af190811561257f57899161256a575b5050836124d8575b5061245b926118186104b99284613595565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156105905761056393858094604051968795869485937fe98a578400000000000000000000000000000000000000000000000000000000855285600486016135dc565b803b15610d02576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87166004820152602481018590529088908290604490829084905af190811561255f57889161254a575b50612449565b8161255491612a1b565b61061b57865f612544565b6040513d8a823e3d90fd5b8161257491612a1b565b610d0257875f612441565b6040513d8b823e3d90fd5b50346128225760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112612822576125c26129d5565b6125ca6129f8565b906044356064359160843560a43567ffffffffffffffff8111612822576125f5903690600401612af9565b9060c43567ffffffffffffffff811161282257612616903690600401612af9565b61261f82612fec565b61266f8473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b926126e681610f4860405160208101907f4c82fe77677b753d6e8e9d6fd479564b08b135757219f2044e54db1882de3aa6825273ffffffffffffffffffffffffffffffffffffffff8d16978860408301528b60608301528c60808301528960a083015260c082015260c0815261034360e082612a1b565b156128265790610794826020612755945191012060405160208101917fe9f21759dec582500f25edee01bab124984bac6771cf3f3dcb23232955508347835273ffffffffffffffffffffffffffffffffffffffff89166040830152606082015260608152610343608082612a1b565b803b15612822576040517f124d91e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810185905260448101869052905f908290606490829084905af18015612817576127e3575b506108a7949261085a836104b961082361105d956108769997613595565b61105d9196509261085a61087695936128005f6108a79997612a1b565b6104b96108235f9a955050509395505092946127c5565b6040513d5f823e3d90fd5b5f80fd5b73ffffffffffffffffffffffffffffffffffffffff857fa424a941000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b346128225761288b61287936612b17565b966110d6839693979297959495612fec565b005b346128225760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112612822576128c46129d5565b506128cd6129f8565b5060643567ffffffffffffffff8111612822576128ee903690600401612af9565b5060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346128225760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261282257600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361282257817f4e2312e000000000000000000000000000000000000000000000000000000000602093149081156129ab575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836129a4565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361282257565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361282257565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612a5c57604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b67ffffffffffffffff8111612a5c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612acf82612a89565b91612add6040519384612a1b565b829481845281830111612822578281602093845f960137010152565b9080601f8301121561282257816020612b1493359101612ac3565b90565b906101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126128225760043573ffffffffffffffffffffffffffffffffffffffff81168103612822579160243573ffffffffffffffffffffffffffffffffffffffff8116810361282257916044359160643573ffffffffffffffffffffffffffffffffffffffff8116810361282257916084359160a4359160c43567ffffffffffffffff81116128225782612bd191600401612af9565b9160e4359067ffffffffffffffff821161282257612b1491600401612af9565b9181601f840112156128225782359167ffffffffffffffff8311612822576020838186019501011161282257565b9080601f830112156128225781359167ffffffffffffffff8311612a5c578260051b906020820193612c546040519586612a1b565b845260208085019282010192831161282257602001905b828210612c785750505090565b8135815260209182019101612c6b565b906101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126128225760043573ffffffffffffffffffffffffffffffffffffffff81168103612822579160243573ffffffffffffffffffffffffffffffffffffffff81168103612822579160443567ffffffffffffffff81116128225782612d1591600401612c1f565b9160643567ffffffffffffffff81116128225781612d3591600401612c1f565b9160843567ffffffffffffffff81116128225782612d5591600401612af9565b9160a4359160c43567ffffffffffffffff81116128225782612bd191600401612af9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b90602080835192838152019201905f5b818110612dd95750505090565b8251845260209384019390920191600101612dcc565b906101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126128225760043573ffffffffffffffffffffffffffffffffffffffff81168103612822579160243573ffffffffffffffffffffffffffffffffffffffff811681036128225791604435916064359160843567ffffffffffffffff81116128225782612d5591600401612af9565b6101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126128225760043573ffffffffffffffffffffffffffffffffffffffff81168103612822579160243573ffffffffffffffffffffffffffffffffffffffff8116810361282257916044359160643573ffffffffffffffffffffffffffffffffffffffff8116810361282257916084359160a4359160c43567ffffffffffffffff81116128225781612f3d91600401612af9565b9160e43567ffffffffffffffff81116128225782612f5d91600401612af9565b91610104359067ffffffffffffffff821161282257612b1491600401612af9565b80516020909101905f5b818110612f955750505090565b8251845260209384019390920191600101612f88565b91612fde9073ffffffffffffffffffffffffffffffffffffffff612b149593168452606060208501526060840190612dbc565b916040818403910152612dbc565b804211612ff65750565b7fa564eef9000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b929491959690936130788473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9161310c89610f4860405160208101907f50c95a183fa9071740cc40b876855b5397200fe0dec9b35ebbeabeb42e6fe5ff825273ffffffffffffffffffffffffffffffffffffffff8b16988960408301528d606083015273ffffffffffffffffffffffffffffffffffffffff8d1660808301528760a08301528860c083015260e082015260e0815261034361010082612a1b565b15985f9961332657916123d18992610794856020613183975191012060405160208101917f92b6ebbf8a8d1a4a46407e547752332e48ebdb6dbadef77e897e681f10a9c845835273ffffffffffffffffffffffffffffffffffffffff8c166040830152606082015260608152610343608082612a1b565b8391933b15612822576040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260248101919091525f8160448183865af1801561281757613300575b509161320a610876928594838961324a98826132ee575b92505050613595565b95604051978894602086019094939273ffffffffffffffffffffffffffffffffffffffff9060609382608085019816845260208401521660408201520152565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00541691823b1561059057916132ca9391858094604051968795869485937f3ea113d000000000000000000000000000000000000000000000000000000000855285600486016135dc565b03925af18015610585576132dc575050565b6132e7828092612a1b565b6101a95750565b6132f793613d6f565b5f838987613201565b61324a9493919850610876926133195f61320a93612a1b565b5f999294955092506131ea565b73ffffffffffffffffffffffffffffffffffffffff867f9c0f9331000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b604290613373613e11565b90604051917f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201522090565b9190823b61341e57906133bb91613dd7565b5060048110156133f1571591826133d157505090565b73ffffffffffffffffffffffffffffffffffffffff919250811691161490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b915f926104b961346d859460405192839160208301957f1626ba7e0000000000000000000000000000000000000000000000000000000087526024840152604060448401526064830190612d79565b51915afa3d156134ee573d61348181612a89565b9061348f6040519283612a1b565b81523d5f602083013e5b816134e0575b816134a8575090565b905060208180518101031261282257602001517f1626ba7e000000000000000000000000000000000000000000000000000000001490565b90506020815110159061349f565b6060613499565b9073ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca01541691821561356d5761353e926133a9565b1561354557565b7fc0299d90000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f95d18bbf000000000000000000000000000000000000000000000000000000005f5260045ffd5b60408051306020820190815273ffffffffffffffffffffffffffffffffffffffff909316918101919091526060810192909252906135d681608081016104b9565b51902090565b9093929360038110156133f157612b149460809373ffffffffffffffffffffffffffffffffffffffff92845260208401521660408201528160608201520190612d79565b909173ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15612822576136a0935f8094604051968795869485937f3ea113d00000000000000000000000000000000000000000000000000000000085526002600486016135dc565b03925af18015612817576136b15750565b5f6136bb91612a1b565b565b92919281156137305773ffffffffffffffffffffffffffffffffffffffff1680158015613725575b6136f757506127109083020480920390565b7fc028e60c000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b5061271082116136e5565b50505f9190565b909173ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15612822576136a0935f8094604051968795869485937fe98a57840000000000000000000000000000000000000000000000000000000085526002600486016135dc565b909173ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15612822576136a0935f8094604051968795869485937f0b7e40c60000000000000000000000000000000000000000000000000000000085526002600486016135dc565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff909216602483015260448201929092526136bb916138be82606481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283612a1b565b613e72565b939194969092956139ad81610f486139218873ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9660405173ffffffffffffffffffffffffffffffffffffffff60208201927fe02fdfc6897b4a06000941baacdeb844de72778c2eef3e70571d3e5920463ec18452169c8d60408301528b606083015273ffffffffffffffffffffffffffffffffffffffff8d1660808301528860a08301528960c083015260e082015260e0815261034361010082612a1b565b15613af057613a32866104b994613a2c613a4c989795610794866020611818985191012060405160208101917fcadc7cd2eeaec3e47314c811398d3aa8a9ecdf0c276677031bb365dc962e99e7835273ffffffffffffffffffffffffffffffffffffffff8d166040830152606082015260608152610343608082612a1b565b896136bd565b5091613a408730878c613d6f565b82613ae0575b84613595565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15612822576136a0935f8094604051968795869485937f3a15d07000000000000000000000000000000000000000000000000000000000855285600486015260248501526044840152608060648401526084830190612d79565b613aeb83898b613837565b613a46565b73ffffffffffffffffffffffffffffffffffffffff857f9c0f9331000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b604051905f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103548060011c9160018216918215613c7d575b602084108314613c50578386528592908115613c135750600114613b95575b6136bb92500383612a1b565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b818310613bf75750509060206136bb92820101613b89565b6020919350806001915483858901015201910190918492613bdf565b602092506136bb9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101613b89565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692613b6a565b9193929173ffffffffffffffffffffffffffffffffffffffff1690811561356d576041845103613d4757602084015192606060408601519501515f1a92803b15612822575f9573ffffffffffffffffffffffffffffffffffffffff9560e49588946040519a8b998a987fd505accf000000000000000000000000000000000000000000000000000000008a5216600489015230602489015260448801526064870152608486015260a485015260c48401525af18015612817576136b15750565b7fd114b229000000000000000000000000000000000000000000000000000000005f5260045ffd5b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff928316602482015292909116604483015260648201929092526136bb916138be8260848101613892565b8151919060418303613e0757613e009250602082015190606060408401519301515f1a90613ef9565b9192909190565b50505f9160029190565b613e19613f88565b613e2161410f565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a081526135d660c082612a1b565b905f602091828151910182855af115612817575f513d613ef0575073ffffffffffffffffffffffffffffffffffffffff81163b155b613eae5750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b60011415613ea7565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411613f7d579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612817575f5173ffffffffffffffffffffffffffffffffffffffff811615613f7357905f905f90565b505f906001905f90565b5050505f9160039190565b6040515f907fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102548060011c9060018116938415614105575b602083108514613c50578284526020840194849281156140cb575060011461404d575b613fef92500382612a1b565b51908115613ffb572090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005480156140285790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106140af575050906020613fef92820101613fe3565b6020919350806001915483858801015201910190918392614097565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016865250613fef92151560051b82016020019050613fe3565b91607f1691613fc0565b614117613b32565b8051908115614127576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015614028579056fea26469706673582212208d575c499d18c1d1389833302386b040fb13be68502c699d6c27d240c09d3bec64736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ForgeV2 is an auto generated Go binding around an Ethereum contract.
type ForgeV2 struct {
	abi abi.ABI
}

// NewForgeV2 creates a new instance of ForgeV2.
func NewForgeV2() *ForgeV2 {
	parsed, err := ForgeV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ForgeV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ForgeV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV2 *ForgeV2) PackDOMAINSEPARATOR() []byte {
	enc, err := forgeV2.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV2 *ForgeV2) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return forgeV2.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV2 *ForgeV2) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := forgeV2.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackBurnERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ddc1f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC1155(address from, address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC1155(from common.Address, token common.Address, tokenID *big.Int, amount *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC1155", from, token, tokenID, amount, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ddc1f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC1155(address from, address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackBurnERC1155(from common.Address, token common.Address, tokenID *big.Int, amount *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("burnERC1155", from, token, tokenID, amount, deadline, fromSig, validatorSig)
}

// PackBurnERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa84adb32.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC1155Batch(address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC1155Batch(from common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC1155Batch", from, token, tokenIDs, amounts, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa84adb32.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC1155Batch(address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackBurnERC1155Batch(from common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("burnERC1155Batch", from, token, tokenIDs, amounts, deadline, fromSig, validatorSig)
}

// PackBurnERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2af4c9d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2af4c9d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackBurnERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("burnERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
}

// PackBurnERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x939cefc1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x939cefc1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) TryPackBurnERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("burnERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
}

// PackBurnERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38033a8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC721(address from, address token, uint256 tokenID, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackBurnERC721(from common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("burnERC721", from, token, tokenID, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38033a8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC721(address from, address token, uint256 tokenID, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackBurnERC721(from common.Address, token common.Address, tokenID *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("burnERC721", from, token, tokenID, deadline, fromSig, validatorSig)
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV2 *ForgeV2) PackEip712Domain() []byte {
	enc, err := forgeV2.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV2 *ForgeV2) TryPackEip712Domain() ([]byte, error) {
	return forgeV2.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV2 *ForgeV2) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := forgeV2.abi.Unpack("eip712Domain", data)
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
	return *outstruct, nil
}

// PackMintERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf6bdd97.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf6bdd97.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("mintERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
}

// PackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43dc7f78.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43dc7f78.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
}

// PackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3081f746.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3081f746.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackMintERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("mintERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
}

// PackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb686a78.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("mintERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb686a78.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("mintERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV2 *ForgeV2) PackNonces(owner common.Address) []byte {
	enc, err := forgeV2.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV2 *ForgeV2) TryPackNonces(owner common.Address) ([]byte, error) {
	return forgeV2.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV2 *ForgeV2) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := forgeV2.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return forgeV2.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return forgeV2.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := forgeV2.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return forgeV2.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV2 *ForgeV2) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := forgeV2.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV2 *ForgeV2) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := forgeV2.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV2 *ForgeV2) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return forgeV2.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV2 *ForgeV2) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := forgeV2.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8abb7071.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8abb7071.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferERC1155", recipient, token, tokenID, amount, data, deadline, recipientSig, validatorSig)
}

// PackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60cfd501.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60cfd501.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig)
}

// PackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72033c8c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72033c8c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig)
}

// PackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba32f859.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba32f859.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes recipientSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, recipientSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferERC721", recipient, token, tokenID, data, deadline, recipientSig, validatorSig)
}

// PackTransferFromERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a9bb38.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFromERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) PackTransferFromERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferFromERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFromERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a9bb38.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFromERC20(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferFromERC20(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferFromERC20", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig)
}

// PackTransferFromERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb55ba9c4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFromERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) PackTransferFromERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV2.abi.Pack("transferFromERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFromERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb55ba9c4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFromERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes fromSig, bytes validatorSig, bytes permitSig) returns()
func (forgeV2 *ForgeV2) TryPackTransferFromERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, fromSig []byte, validatorSig []byte, permitSig []byte) ([]byte, error) {
	return forgeV2.abi.Pack("transferFromERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig)
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validator() view returns(address)
func (forgeV2 *ForgeV2) PackValidator() []byte {
	enc, err := forgeV2.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validator() view returns(address)
func (forgeV2 *ForgeV2) TryPackValidator() ([]byte, error) {
	return forgeV2.abi.Pack("validator")
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (forgeV2 *ForgeV2) UnpackValidator(data []byte) (common.Address, error) {
	out, err := forgeV2.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// ForgeV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the ForgeV2 contract.
type ForgeV2EIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ForgeV2EIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ForgeV2EIP712DomainChanged) ContractEventName() string {
	return ForgeV2EIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (forgeV2 *ForgeV2) UnpackEIP712DomainChangedEvent(log *types.Log) (*ForgeV2EIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2EIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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

// ForgeV2Initialized represents a Initialized event raised by the ForgeV2 contract.
type ForgeV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ForgeV2Initialized) ContractEventName() string {
	return ForgeV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (forgeV2 *ForgeV2) UnpackInitializedEvent(log *types.Log) (*ForgeV2Initialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2Initialized)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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

// ForgeV2ValidatorUpdated represents a ValidatorUpdated event raised by the ForgeV2 contract.
type ForgeV2ValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const ForgeV2ValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (ForgeV2ValidatorUpdated) ContractEventName() string {
	return ForgeV2ValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (forgeV2 *ForgeV2) UnpackValidatorUpdatedEvent(log *types.Log) (*ForgeV2ValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV2ValidatorUpdated)
	if len(log.Data) > 0 {
		if err := forgeV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV2.abi.Events[event].Inputs {
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
func (forgeV2 *ForgeV2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeInvalidFeeData"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeInvalidFeeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeInvalidPermitSignatureLength"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeInvalidPermitSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return forgeV2.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC1155ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC1155ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC20ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC20ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["ERC721ForgeV2InvalidAccountSignature"].ID.Bytes()[:4]) {
		return forgeV2.UnpackERC721ForgeV2InvalidAccountSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return forgeV2.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return forgeV2.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return forgeV2.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV2.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return forgeV2.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ForgeV2BaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the ForgeV2 contract.
type ForgeV2BaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func ForgeV2BaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (forgeV2 *ForgeV2) UnpackBaseForgeExpiredSignatureError(raw []byte) (*ForgeV2BaseForgeExpiredSignature, error) {
	out := new(ForgeV2BaseForgeExpiredSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeInvalidFeeData represents a BaseForge__InvalidFeeData error raised by the ForgeV2 contract.
type ForgeV2BaseForgeInvalidFeeData struct {
	FeeRecipient common.Address
	FeeBPS       *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func ForgeV2BaseForgeInvalidFeeDataErrorID() common.Hash {
	return common.HexToHash("0xc028e60c3834e172f648e56b5588dbdbb2caa156ad1931ff196c489c710a7363")
}

// UnpackBaseForgeInvalidFeeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func (forgeV2 *ForgeV2) UnpackBaseForgeInvalidFeeDataError(raw []byte) (*ForgeV2BaseForgeInvalidFeeData, error) {
	out := new(ForgeV2BaseForgeInvalidFeeData)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeInvalidFeeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeInvalidPermitSignatureLength represents a BaseForge__InvalidPermitSignatureLength error raised by the ForgeV2 contract.
type ForgeV2BaseForgeInvalidPermitSignatureLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func ForgeV2BaseForgeInvalidPermitSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xd114b2291dcfec2182b0d6c24f580b96bcae4149c53fa4278ac88775ec7bb032")
}

// UnpackBaseForgeInvalidPermitSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func (forgeV2 *ForgeV2) UnpackBaseForgeInvalidPermitSignatureLengthError(raw []byte) (*ForgeV2BaseForgeInvalidPermitSignatureLength, error) {
	out := new(ForgeV2BaseForgeInvalidPermitSignatureLength)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeInvalidPermitSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeInvalidValidatorSignature represents a BaseForge__InvalidValidatorSignature error raised by the ForgeV2 contract.
type ForgeV2BaseForgeInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func ForgeV2BaseForgeInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0xc0299d906fef6fff02cc5dc92034a345928d480697ea46384b87962e5a4e2b11")
}

// UnpackBaseForgeInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func (forgeV2 *ForgeV2) UnpackBaseForgeInvalidValidatorSignatureError(raw []byte) (*ForgeV2BaseForgeInvalidValidatorSignature, error) {
	out := new(ForgeV2BaseForgeInvalidValidatorSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2BaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the ForgeV2 contract.
type ForgeV2BaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func ForgeV2BaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (forgeV2 *ForgeV2) UnpackBaseForgeZeroAddressError(raw []byte) (*ForgeV2BaseForgeZeroAddress, error) {
	out := new(ForgeV2BaseForgeZeroAddress)
	if err := forgeV2.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC1155ForgeV2InvalidAccountSignature represents a ERC1155ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC1155ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC1155ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0xa424a9419d285008b7909b6400faa2d8734c2da12fd92cf72aa0ba5a815c12e6")
}

// UnpackERC1155ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC1155ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC1155ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC1155ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC1155ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC20ForgeV2InvalidAccountSignature represents a ERC20ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC20ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC20ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0x9c0f9331d31a14bc40f7b3648f1dd167e28f4c2ce6f3452830ca126ff2afbd53")
}

// UnpackERC20ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC20ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC20ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC20ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC20ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2ERC721ForgeV2InvalidAccountSignature represents a ERC721ForgeV2__InvalidAccountSignature error raised by the ForgeV2 contract.
type ForgeV2ERC721ForgeV2InvalidAccountSignature struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721ForgeV2__InvalidAccountSignature(address account)
func ForgeV2ERC721ForgeV2InvalidAccountSignatureErrorID() common.Hash {
	return common.HexToHash("0xfb1b8f0f5fbb9ddf1a7ea566d6daa430cb3b1d4f882ecfb03a31df2c5b0c1008")
}

// UnpackERC721ForgeV2InvalidAccountSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721ForgeV2__InvalidAccountSignature(address account)
func (forgeV2 *ForgeV2) UnpackERC721ForgeV2InvalidAccountSignatureError(raw []byte) (*ForgeV2ERC721ForgeV2InvalidAccountSignature, error) {
	out := new(ForgeV2ERC721ForgeV2InvalidAccountSignature)
	if err := forgeV2.abi.UnpackIntoInterface(out, "ERC721ForgeV2InvalidAccountSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2InvalidAccountNonce represents a InvalidAccountNonce error raised by the ForgeV2 contract.
type ForgeV2InvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ForgeV2InvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (forgeV2 *ForgeV2) UnpackInvalidAccountNonceError(raw []byte) (*ForgeV2InvalidAccountNonce, error) {
	out := new(ForgeV2InvalidAccountNonce)
	if err := forgeV2.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2InvalidInitialization represents a InvalidInitialization error raised by the ForgeV2 contract.
type ForgeV2InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ForgeV2InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (forgeV2 *ForgeV2) UnpackInvalidInitializationError(raw []byte) (*ForgeV2InvalidInitialization, error) {
	out := new(ForgeV2InvalidInitialization)
	if err := forgeV2.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2NotInitializing represents a NotInitializing error raised by the ForgeV2 contract.
type ForgeV2NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ForgeV2NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (forgeV2 *ForgeV2) UnpackNotInitializingError(raw []byte) (*ForgeV2NotInitializing, error) {
	out := new(ForgeV2NotInitializing)
	if err := forgeV2.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV2SafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ForgeV2 contract.
type ForgeV2SafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ForgeV2SafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (forgeV2 *ForgeV2) UnpackSafeERC20FailedOperationError(raw []byte) (*ForgeV2SafeERC20FailedOperation, error) {
	out := new(ForgeV2SafeERC20FailedOperation)
	if err := forgeV2.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
