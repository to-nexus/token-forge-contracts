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

// ForgeV3MetaData contains all meta data concerning the ForgeV3 contract.
var ForgeV3MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"burnERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"mintERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC1155Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"}],\"name\":\"transferFromERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseForge__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBPS\",\"type\":\"uint256\"}],\"name\":\"BaseForge__InvalidFeeData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidPermitSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__InvalidValidatorSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseForge__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721MintForge__InvalidRecipientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ForgeV3",
	Bin: "0x6080806040523460aa575f516020612f145f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b604051612e6590816100af8239f35b6001600160401b0319166001600160401b039081175f516020612f145f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a714611bc157508063098dad4914611962578063150b7a02146118d5578063199a40451461172b5780632ee543ed1461159057806331e2f9b4146113765780633644e515146113355780633a5381b5146112c457806342ce35a11461105b57806353307fb014610e1c5780637b07e73914610bde5780637ecebe0014610b5c57806384b0196e146109dc578063aab8034814610781578063b172aa00146104bc578063bc197c81146103ed578063d81de88f146101755763f23a6e61146100e4575f80fd5b346101725760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725761011b611c7d565b50610124611ca0565b5060843567ffffffffffffffff811161017057610145903690600401611e0a565b5060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b505b80fd5b50346101725780806102a46101bf61032f61029e6103036102996102c361019b3661213a565b9e869f97859f979893879f829e96979c9495828a858e6101ba846122f8565b612721565b73ffffffffffffffffffffffffffffffffffffffff806102258473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9b81604051956020870197507f0b4de94c63368b482149a27594e4b35fc83e159d5641e4972f601e8a3dd4c4f98852166040860152169c8d60608501528960808501521660a08301528660c08301528960e0830152610100820152610100815261029161012082611cc3565b51902061232d565b61236e565b8a61261b565b50926102b288308884612c4f565b8389816103dc575b5050508461240e565b95604051958694602086019094939273ffffffffffffffffffffffffffffffffffffffff9060609382608085019816845260208401521660408201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282611cc3565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937f3a15d07000000000000000000000000000000000000000000000000000000000855260048501612455565b03925af180156103cd576103bc5750f35b816103c691611cc3565b6101725780f35b6040513d84823e3d90fd5b8480fd5b6103e592612695565b5f83896102ba565b50346101725760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257610425611c7d565b5061042e611ca0565b5060443567ffffffffffffffff81116101705761044f903690600401611d31565b5060643567ffffffffffffffff811161017057610470903690600401611d31565b5060843567ffffffffffffffff811161017057610491903690600401611e0a565b5060206040517fbc197c81000000000000000000000000000000000000000000000000000000008152f35b50346101725780610299816105c26105bc6104d636611faa565b859a849a88846104ea899d99989c966122f8565b73ffffffffffffffffffffffffffffffffffffffff806105508473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9b81604051956020870197507f29fb0f21ef9e1026ffc6321253b615d00e4bede79417f58d0fbebf2eb55eeef38852166040860152169c8d60608501528960808501521660a08301528660c08301528960e0830152610100820152610100815261029161012082611cc3565b8761261b565b8491943b1561077d576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810191909152888160448183865af190811561077257899161075d575b5050836106c3575b50610646926102c3610303928461240e565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937fe98a5784000000000000000000000000000000000000000000000000000000008552856004860161248d565b803b15610759576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87166004820152602481018590529088908290604490829084905af190811561074e578891610735575b50610634565b8161073f91611cc3565b61074a57865f61072f565b8680fd5b6040513d8a823e3d90fd5b8780fd5b8161076791611cc3565b61075957875f61062c565b6040513d8b823e3d90fd5b8880fd5b5034610172578061088d61079436612062565b9685976107a58495939697946122f8565b6108876107f88973ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9661087f8c8b875160208901209973ffffffffffffffffffffffffffffffffffffffff80604051936020850195507f9533ee965a40f21d53123db5a4eb531bff3f3c0aa35e2839edd63086691566318652169a8b6040850152169a8b60608401528d608084015260a08301528b60c083015260e082015260e0815261029161010082611cc3565b923691611dd4565b9061236e565b823b15610759576108e9928892836040518096819582947fb88d4fde00000000000000000000000000000000000000000000000000000000845230600485015260248401528a6044840152608060648401526084830190612214565b03925af19081156109d15786916109bc575b505061090a61093e918461240e565b6040805173ffffffffffffffffffffffffffffffffffffffff90951660208601528401929092529091908160608101610303565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937f0b7e40c600000000000000000000000000000000000000000000000000000000855260016004860161248d565b816109c691611cc3565b6103d857845f6108fb565b6040513d88823e3d90fd5b503461017257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610172577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580610b33575b15610ad557610aa490610ad1610a48612809565b91610a5161295e565b610ab260405191610a63602084611cc3565b8383525f3681376040519687967f0f00000000000000000000000000000000000000000000000000000000000000885260e0602089015260e0880190612214565b908682036040880152612214565b9146606086015230608086015260a085015283820360c0850152612257565b0390f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415610a34565b50346101725760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257604060209173ffffffffffffffffffffffffffffffffffffffff610bae611c7d565b1681527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0083522054604051908152f35b5034610172578080610c19610ceb6102996105bc610bfb3661213a565b9b8783869d899f828a9f9b9c9d98999e9697858e6101ba8e936122f8565b73ffffffffffffffffffffffffffffffffffffffff80610c7f8473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9b81604051956020870197507fc742fe2727dd2ec0d2f49da6c2ddc688c79d1fb71500b54045cb8bac9b33a5608852166040860152169c8d60608501528960808501521660a08301528660c08301528960e0830152610100820152610100815261029161012082611cc3565b8491943b1561077d576040517f79cc679000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810191909152888160448183865af1908115610772578991610e03575b5050926102c361030392610d759583898782610df1575b9250505061240e565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937f3ea113d000000000000000000000000000000000000000000000000000000000855260048501612455565b610dfa93612c4f565b5f838987610d6c565b81610e1091969396611cc3565b6107595792875f610d55565b503461017257610f1f81610e2f36611f17565b918498829593610e43889a969998946122f8565b61029988610e978173ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9885516020870120906040519173ffffffffffffffffffffffffffffffffffffffff8060208501957f1cf386885696073114b875bcc0ca5841b01f81d9cf119e96624c8ddac5b158e2875216998a604086015216998a606085015260808401528d60a084015260c08301528a60e0830152610100820152610100815261029161012082611cc3565b823b1561105757610f829289858094604051968795869485937ff242432a000000000000000000000000000000000000000000000000000000008552306004860152602485015260448401528c606484015260a0608484015260a4830190612214565b03925af180156103cd57611036575b50611033611002868661102e8a610fe689610303610faf8b8361240e565b956040519485936020850160409194939273ffffffffffffffffffffffffffffffffffffffff606083019616825260208201520152565b6040519485918860208401526040808401526060830190612214565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284611cc3565b6124fe565b80f35b81611045919695949396611cc3565b61105357909192855f610f91565b8580fd5b8380fd5b5034610172578061106b36611e28565b9294809694979261107b836122f8565b61118e6110ce8773ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9561029986898d6040516110ea8161030360208201809561228a565b519020978c6040516111048161030360208201809561228a565b519020875160208901209073ffffffffffffffffffffffffffffffffffffffff80604051956020870197507f1891b74327c5901e0986c8ac5f98aff01d90354ca2a7d5fe7ef6fcff03bd51798852169a8b6040870152169a8b6060860152608085015260a084015260c08301528a60e0830152610100820152610100815261029161012082611cc3565b823b15611057578692898580946112486111e895611218604051998a98899788967f1f7fdffa0000000000000000000000000000000000000000000000000000000088526004880152608060248801526084870190612257565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016044870152612257565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152612214565b03925af180156103cd576112a7575b5061103361100286866112a28a611285896103036112758b8361240e565b95604051948593602085016122b7565b604051948591600160208401526040808401526060830190612214565b61259b565b816112b6919695949396611cc3565b61105357909192855f611257565b503461017257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257602073ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca015416604051908152f35b503461017257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257602061136e612bee565b604051908152f35b5034610172578061087f6108878261147a6114c3602061139536612062565b919892959a90859c859b6113ab8c9897986122f8565b8d6113fc8173ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9c8a518c8c012073ffffffffffffffffffffffffffffffffffffffff80604051948f860196507f3c30a679068e4bb67e250491cb883b798e16e7f5f3c412fe63f2729171591e8d8752169a8b6040860152169a8b6060850152608084015260a08301528d60c083015260e082015260e0815261029161010082611cc3565b896040518096819582947f94d008ef00000000000000000000000000000000000000000000000000000000845260048401528a6024840152606060448401526064830190612214565b03925af180156109d15761155e575b5061090a6114e0918461240e565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937fe98a578400000000000000000000000000000000000000000000000000000000855260016004860161248d565b6020813d602011611588575b8161157760209383611cc3565b81010312611053575061090a6114d2565b3d915061156a565b50346101725780610299816116a46116ae61169e6103036102b26102c36115b636611faa565b859e849e88879f979698999d946115cc816122f8565b73ffffffffffffffffffffffffffffffffffffffff806116328473ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9c81604051956020870197507f9dd4d5c883252e1fb1700d3da5d6b337b54fcdb459e7780e8ce54cf87a8211bf8852166040860152169a8b60608501528960808501521660a08301528660c08301528a60e0830152610100820152610100815261029161012082611cc3565b8b61261b565b8783969296612695565b73ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b156103d8576103ab93858094604051968795869485937f0b7e40c6000000000000000000000000000000000000000000000000000000008552856004860161248d565b50346101725761182e8161173e36611f17565b918498829593611752889a969998946122f8565b610299886117a68173ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9885516020870120906040519173ffffffffffffffffffffffffffffffffffffffff8060208501957f3583c0cae54afbc9a9fa379eb68ad426af9ae25cd06eea4ae16a8e22db5a23d6875216998a604086015216998a606085015260808401528d60a084015260c08301528a60e0830152610100820152610100815261029161012082611cc3565b823b156110575761188b9289858094604051968795869485937f731133e9000000000000000000000000000000000000000000000000000000008552600485015260248401528c6044840152608060648401526084830190612214565b03925af180156103cd576118b8575b5061103361100286866112a28a610fe689610303610faf8b8361240e565b816118c7919695949396611cc3565b61105357909192855f61189a565b50346101725760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725761190d611c7d565b50611916611ca0565b5060643567ffffffffffffffff811161017057611937903690600401611e0a565b5060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b5034611bbd5761197136611e28565b919480969395611980836122f8565b611a916119d38673ffffffffffffffffffffffffffffffffffffffff165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260405f2080549060018201905590565b9461029987896040516119ee8161030360208201809561228a565b519020968b604051611a088161030360208201809561228a565b51902090865160208801206040519273ffffffffffffffffffffffffffffffffffffffff8060208601967fe5e0659cf0947bacce7219f90e80df40880da61440bf0c5adc17f9f0b9aeb8d78852169a8b6040870152169a8b6060860152608085015260a084015260c08301528960e0830152610100820152610100815261029161012082611cc3565b823b15611bbd578692865f8094611b51611af195611b21604051998a98899788967f2eb2c2d6000000000000000000000000000000000000000000000000000000008852306004890152602488015260a0604488015260a4870190612257565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016064870152612257565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016084850152612214565b03925af18015611bb257611b7e575b5061103394926112858361030361127561102e95611002999761240e565b61102e919650926112856110029593611b9b5f6110339997611cc3565b6103036112755f9a95505050939550509294611b60565b6040513d5f823e3d90fd5b5f80fd5b34611bbd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611bbd57600435907fffffffff000000000000000000000000000000000000000000000000000000008216809203611bbd57817f4e2312e00000000000000000000000000000000000000000000000000000000060209314908115611c53575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611c4c565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611bbd57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203611bbd57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611d0457604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b9080601f83011215611bbd5781359167ffffffffffffffff8311611d04578260051b906020820193611d666040519586611cc3565b8452602080850192820101928311611bbd57602001905b828210611d8a5750505090565b8135815260209182019101611d7d565b67ffffffffffffffff8111611d0457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192611de082611d9a565b91611dee6040519384611cc3565b829481845281830111611bbd578281602093845f960137010152565b9080601f83011215611bbd57816020611e2593359101611dd4565b90565b60e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611bbd5760043573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160243573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160443567ffffffffffffffff8111611bbd5781611eb391600401611d31565b9160643567ffffffffffffffff8111611bbd5782611ed391600401611d31565b9160843567ffffffffffffffff8111611bbd5781611ef391600401611e0a565b9160a4359160c4359067ffffffffffffffff8211611bbd57611e2591600401611e0a565b60e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611bbd5760043573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160243573ffffffffffffffffffffffffffffffffffffffff81168103611bbd5791604435916064359160843567ffffffffffffffff8111611bbd5781611ef391600401611e0a565b60e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611bbd5760043573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160243573ffffffffffffffffffffffffffffffffffffffff81168103611bbd57916044359160643573ffffffffffffffffffffffffffffffffffffffff81168103611bbd57916084359160a4359160c4359067ffffffffffffffff8211611bbd57611e2591600401611e0a565b60c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611bbd5760043573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160243573ffffffffffffffffffffffffffffffffffffffff81168103611bbd57916044359160643567ffffffffffffffff8111611bbd57826120f191600401611e0a565b916084359160a43567ffffffffffffffff8111611bbd5782602382011215611bbd5780600401359267ffffffffffffffff8411611bbd5760248483010111611bbd576024019190565b906101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112611bbd5760043573ffffffffffffffffffffffffffffffffffffffff81168103611bbd579160243573ffffffffffffffffffffffffffffffffffffffff81168103611bbd57916044359160643573ffffffffffffffffffffffffffffffffffffffff81168103611bbd57916084359160a4359160c43567ffffffffffffffff8111611bbd57826121f491600401611e0a565b9160e4359067ffffffffffffffff8211611bbd57611e2591600401611e0a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b90602080835192838152019201905f5b8181106122745750505090565b8251845260209384019390920191600101612267565b80516020909101905f5b8181106122a15750505090565b8251845260209384019390920191600101612294565b916122ea9073ffffffffffffffffffffffffffffffffffffffff611e259593168452606060208501526060840190612257565b916040818403910152612257565b8042116123025750565b7fa564eef9000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b604290612338612bee565b90604051917f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201522090565b9073ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca0154169182156123e6576123b792612a48565b156123be57565b7fc0299d90000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f95d18bbf000000000000000000000000000000000000000000000000000000005f5260045ffd5b60408051306020820190815273ffffffffffffffffffffffffffffffffffffffff9093169181019190915260608101929092529061244f8160808101610303565b51902090565b73ffffffffffffffffffffffffffffffffffffffff611e2594936080935f845260208401521660408201528160608201520190612214565b9093929360038110156124d157611e259460809373ffffffffffffffffffffffffffffffffffffffff92845260208401521660408201528160608201520190612214565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b909173ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15611bbd5761257e935f8094604051968795869485937f0b7e40c600000000000000000000000000000000000000000000000000000000855260026004860161248d565b03925af18015611bb25761258f5750565b5f61259991611cc3565b565b909173ffffffffffffffffffffffffffffffffffffffff7f2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca005416803b15611bbd5761257e935f8094604051968795869485937fe98a578400000000000000000000000000000000000000000000000000000000855260026004860161248d565b929192811561268e5773ffffffffffffffffffffffffffffffffffffffff1680158015612683575b61265557506127109083020480920390565b7fc028e60c000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b506127108211612643565b50505f9190565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff909216602483015260448201929092526125999161271c82606481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283611cc3565b612b67565b9193929173ffffffffffffffffffffffffffffffffffffffff169081156123e65760418451036127e157602084015192606060408601519501515f1a92803b15611bbd575f9573ffffffffffffffffffffffffffffffffffffffff9560e49588946040519a8b998a987fd505accf000000000000000000000000000000000000000000000000000000008a5216600489015230602489015260448801526064870152608486015260a485015260c48401525af18015611bb25761258f5750565b7fd114b229000000000000000000000000000000000000000000000000000000005f5260045ffd5b604051905f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102548060011c9160018216918215612954575b6020841083146129275783865285929081156128ea575060011461286c575b61259992500383611cc3565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106128ce57505090602061259992820101612860565b60209193508060019154838589010152019101909184926128b6565b602092506125999491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101612860565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692612841565b604051905f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103548060011c9160018216918215612a3e575b6020841083146129275783865285929081156128ea57506001146129c05761259992500383611cc3565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b818310612a2257505090602061259992820101612860565b6020919350806001915483858901015201910190918492612a0a565b92607f1692612996565b9190823b612a905790612a5a91612cb7565b5060048110156124d157159182612a7057505090565b73ffffffffffffffffffffffffffffffffffffffff919250811691161490565b915f92610303612adf859460405192839160208301957f1626ba7e0000000000000000000000000000000000000000000000000000000087526024840152604060448401526064830190612214565b51915afa3d15612b60573d612af381611d9a565b90612b016040519283611cc3565b81523d5f602083013e5b81612b52575b81612b1a575090565b9050602081805181010312611bbd57602001517f1626ba7e000000000000000000000000000000000000000000000000000000001490565b905060208151101590612b11565b6060612b0b565b905f602091828151910182855af115611bb2575f513d612be5575073ffffffffffffffffffffffffffffffffffffffff81163b155b612ba35750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b60011415612b9c565b612bf6612cf1565b612bfe612d5b565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261244f60c082611cc3565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff928316602482015292909116604483015260648201929092526125999161271c82608481016126f0565b8151919060418303612ce757612ce09250602082015190606060408401519301515f1a90612da0565b9192909190565b50505f9160029190565b612cf9612809565b8051908115612d09576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100548015612d365790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b612d6361295e565b8051908115612d73576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015612d365790565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411612e24579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15611bb2575f5173ffffffffffffffffffffffffffffffffffffffff811615612e1a57905f905f90565b505f906001905f90565b5050505f916003919056fea2646970667358221220bed26f3c1ec7a2c85110a550e10b441a96586718b15fb37c4d1764056a95a7e964736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ForgeV3 is an auto generated Go binding around an Ethereum contract.
type ForgeV3 struct {
	abi abi.ABI
}

// NewForgeV3 creates a new instance of ForgeV3.
func NewForgeV3() *ForgeV3 {
	parsed, err := ForgeV3MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ForgeV3{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ForgeV3) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV3 *ForgeV3) PackDOMAINSEPARATOR() []byte {
	enc, err := forgeV3.abi.Pack("DOMAIN_SEPARATOR")
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
func (forgeV3 *ForgeV3) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return forgeV3.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (forgeV3 *ForgeV3) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := forgeV3.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackBurnERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b07e739.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig, bytes permitSig) returns()
func (forgeV3 *ForgeV3) PackBurnERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("burnERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b07e739.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig, bytes permitSig) returns()
func (forgeV3 *ForgeV3) TryPackBurnERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte, permitSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("burnERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig)
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV3 *ForgeV3) PackEip712Domain() []byte {
	enc, err := forgeV3.abi.Pack("eip712Domain")
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
func (forgeV3 *ForgeV3) TryPackEip712Domain() ([]byte, error) {
	return forgeV3.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (forgeV3 *ForgeV3) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := forgeV3.abi.Unpack("eip712Domain", data)
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
// the contract method with ID 0x199a4045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("mintERC1155", recipient, token, tokenID, amount, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x199a4045.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackMintERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("mintERC1155", recipient, token, tokenID, amount, data, deadline, validatorSig)
}

// PackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42ce35a1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42ce35a1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackMintERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("mintERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, validatorSig)
}

// PackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb172aa00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackMintERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("mintERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb172aa00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackMintERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("mintERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, validatorSig)
}

// PackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31e2f9b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("mintERC721", recipient, token, tokenID, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMintERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31e2f9b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mintERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackMintERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("mintERC721", recipient, token, tokenID, data, deadline, validatorSig)
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV3 *ForgeV3) PackNonces(owner common.Address) []byte {
	enc, err := forgeV3.abi.Pack("nonces", owner)
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
func (forgeV3 *ForgeV3) TryPackNonces(owner common.Address) ([]byte, error) {
	return forgeV3.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (forgeV3 *ForgeV3) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := forgeV3.abi.Unpack("nonces", data)
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
func (forgeV3 *ForgeV3) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := forgeV3.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
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
func (forgeV3 *ForgeV3) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return forgeV3.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (forgeV3 *ForgeV3) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := forgeV3.abi.Unpack("onERC1155BatchReceived", data)
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
func (forgeV3 *ForgeV3) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := forgeV3.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
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
func (forgeV3 *ForgeV3) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return forgeV3.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (forgeV3 *ForgeV3) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := forgeV3.abi.Unpack("onERC1155Received", data)
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
func (forgeV3 *ForgeV3) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := forgeV3.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
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
func (forgeV3 *ForgeV3) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return forgeV3.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (forgeV3 *ForgeV3) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := forgeV3.abi.Unpack("onERC721Received", data)
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
func (forgeV3 *ForgeV3) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := forgeV3.abi.Pack("supportsInterface", interfaceId)
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
func (forgeV3 *ForgeV3) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return forgeV3.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (forgeV3 *ForgeV3) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := forgeV3.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x53307fb0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("transferERC1155", recipient, token, tokenID, amount, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x53307fb0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC1155(address recipient, address token, uint256 tokenID, uint256 amount, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackTransferERC1155(recipient common.Address, token common.Address, tokenID *big.Int, amount *big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("transferERC1155", recipient, token, tokenID, amount, data, deadline, validatorSig)
}

// PackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x098dad49.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC1155Batch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x098dad49.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC1155Batch(address recipient, address token, uint256[] tokenIDs, uint256[] amounts, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackTransferERC1155Batch(recipient common.Address, token common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("transferERC1155Batch", recipient, token, tokenIDs, amounts, data, deadline, validatorSig)
}

// PackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ee543ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackTransferERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("transferERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ee543ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC20(address recipient, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackTransferERC20(recipient common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("transferERC20", recipient, token, amount, feeRecipient, feeBPS, deadline, validatorSig)
}

// PackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaab80348.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) PackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, validatorSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("transferERC721", recipient, token, tokenID, data, deadline, validatorSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaab80348.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC721(address recipient, address token, uint256 tokenID, bytes data, uint256 deadline, bytes validatorSig) returns()
func (forgeV3 *ForgeV3) TryPackTransferERC721(recipient common.Address, token common.Address, tokenID *big.Int, data []byte, deadline *big.Int, validatorSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("transferERC721", recipient, token, tokenID, data, deadline, validatorSig)
}

// PackTransferFromERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd81de88f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFromERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig, bytes permitSig) returns()
func (forgeV3 *ForgeV3) PackTransferFromERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte, permitSig []byte) []byte {
	enc, err := forgeV3.abi.Pack("transferFromERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFromERC20Permit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd81de88f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFromERC20Permit(address from, address token, uint256 amount, address feeRecipient, uint256 feeBPS, uint256 deadline, bytes validatorSig, bytes permitSig) returns()
func (forgeV3 *ForgeV3) TryPackTransferFromERC20Permit(from common.Address, token common.Address, amount *big.Int, feeRecipient common.Address, feeBPS *big.Int, deadline *big.Int, validatorSig []byte, permitSig []byte) ([]byte, error) {
	return forgeV3.abi.Pack("transferFromERC20Permit", from, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig)
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validator() view returns(address)
func (forgeV3 *ForgeV3) PackValidator() []byte {
	enc, err := forgeV3.abi.Pack("validator")
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
func (forgeV3 *ForgeV3) TryPackValidator() ([]byte, error) {
	return forgeV3.abi.Pack("validator")
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (forgeV3 *ForgeV3) UnpackValidator(data []byte) (common.Address, error) {
	out, err := forgeV3.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// ForgeV3EIP712DomainChanged represents a EIP712DomainChanged event raised by the ForgeV3 contract.
type ForgeV3EIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ForgeV3EIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ForgeV3EIP712DomainChanged) ContractEventName() string {
	return ForgeV3EIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (forgeV3 *ForgeV3) UnpackEIP712DomainChangedEvent(log *types.Log) (*ForgeV3EIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV3.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV3EIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := forgeV3.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV3.abi.Events[event].Inputs {
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

// ForgeV3Initialized represents a Initialized event raised by the ForgeV3 contract.
type ForgeV3Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ForgeV3InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ForgeV3Initialized) ContractEventName() string {
	return ForgeV3InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (forgeV3 *ForgeV3) UnpackInitializedEvent(log *types.Log) (*ForgeV3Initialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV3.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV3Initialized)
	if len(log.Data) > 0 {
		if err := forgeV3.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV3.abi.Events[event].Inputs {
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

// ForgeV3ValidatorUpdated represents a ValidatorUpdated event raised by the ForgeV3 contract.
type ForgeV3ValidatorUpdated struct {
	Validator common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const ForgeV3ValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (ForgeV3ValidatorUpdated) ContractEventName() string {
	return ForgeV3ValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed validator)
func (forgeV3 *ForgeV3) UnpackValidatorUpdatedEvent(log *types.Log) (*ForgeV3ValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != forgeV3.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ForgeV3ValidatorUpdated)
	if len(log.Data) > 0 {
		if err := forgeV3.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range forgeV3.abi.Events[event].Inputs {
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
func (forgeV3 *ForgeV3) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["BaseForgeExpiredSignature"].ID.Bytes()[:4]) {
		return forgeV3.UnpackBaseForgeExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["BaseForgeInvalidFeeData"].ID.Bytes()[:4]) {
		return forgeV3.UnpackBaseForgeInvalidFeeDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["BaseForgeInvalidPermitSignatureLength"].ID.Bytes()[:4]) {
		return forgeV3.UnpackBaseForgeInvalidPermitSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["BaseForgeInvalidValidatorSignature"].ID.Bytes()[:4]) {
		return forgeV3.UnpackBaseForgeInvalidValidatorSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["BaseForgeZeroAddress"].ID.Bytes()[:4]) {
		return forgeV3.UnpackBaseForgeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["ERC721MintForgeInvalidRecipientSignature"].ID.Bytes()[:4]) {
		return forgeV3.UnpackERC721MintForgeInvalidRecipientSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return forgeV3.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return forgeV3.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return forgeV3.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], forgeV3.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return forgeV3.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ForgeV3BaseForgeExpiredSignature represents a BaseForge__ExpiredSignature error raised by the ForgeV3 contract.
type ForgeV3BaseForgeExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func ForgeV3BaseForgeExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0xa564eef9562665c75321203e373a16e1260bbcf280dfa3881cb6983d5ab0497b")
}

// UnpackBaseForgeExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ExpiredSignature(uint256 deadline)
func (forgeV3 *ForgeV3) UnpackBaseForgeExpiredSignatureError(raw []byte) (*ForgeV3BaseForgeExpiredSignature, error) {
	out := new(ForgeV3BaseForgeExpiredSignature)
	if err := forgeV3.abi.UnpackIntoInterface(out, "BaseForgeExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3BaseForgeInvalidFeeData represents a BaseForge__InvalidFeeData error raised by the ForgeV3 contract.
type ForgeV3BaseForgeInvalidFeeData struct {
	FeeRecipient common.Address
	FeeBPS       *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func ForgeV3BaseForgeInvalidFeeDataErrorID() common.Hash {
	return common.HexToHash("0xc028e60c3834e172f648e56b5588dbdbb2caa156ad1931ff196c489c710a7363")
}

// UnpackBaseForgeInvalidFeeDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS)
func (forgeV3 *ForgeV3) UnpackBaseForgeInvalidFeeDataError(raw []byte) (*ForgeV3BaseForgeInvalidFeeData, error) {
	out := new(ForgeV3BaseForgeInvalidFeeData)
	if err := forgeV3.abi.UnpackIntoInterface(out, "BaseForgeInvalidFeeData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3BaseForgeInvalidPermitSignatureLength represents a BaseForge__InvalidPermitSignatureLength error raised by the ForgeV3 contract.
type ForgeV3BaseForgeInvalidPermitSignatureLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func ForgeV3BaseForgeInvalidPermitSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xd114b2291dcfec2182b0d6c24f580b96bcae4149c53fa4278ac88775ec7bb032")
}

// UnpackBaseForgeInvalidPermitSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidPermitSignatureLength()
func (forgeV3 *ForgeV3) UnpackBaseForgeInvalidPermitSignatureLengthError(raw []byte) (*ForgeV3BaseForgeInvalidPermitSignatureLength, error) {
	out := new(ForgeV3BaseForgeInvalidPermitSignatureLength)
	if err := forgeV3.abi.UnpackIntoInterface(out, "BaseForgeInvalidPermitSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3BaseForgeInvalidValidatorSignature represents a BaseForge__InvalidValidatorSignature error raised by the ForgeV3 contract.
type ForgeV3BaseForgeInvalidValidatorSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func ForgeV3BaseForgeInvalidValidatorSignatureErrorID() common.Hash {
	return common.HexToHash("0xc0299d906fef6fff02cc5dc92034a345928d480697ea46384b87962e5a4e2b11")
}

// UnpackBaseForgeInvalidValidatorSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__InvalidValidatorSignature()
func (forgeV3 *ForgeV3) UnpackBaseForgeInvalidValidatorSignatureError(raw []byte) (*ForgeV3BaseForgeInvalidValidatorSignature, error) {
	out := new(ForgeV3BaseForgeInvalidValidatorSignature)
	if err := forgeV3.abi.UnpackIntoInterface(out, "BaseForgeInvalidValidatorSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3BaseForgeZeroAddress represents a BaseForge__ZeroAddress error raised by the ForgeV3 contract.
type ForgeV3BaseForgeZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BaseForge__ZeroAddress()
func ForgeV3BaseForgeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x95d18bbf8a479b49818b711f9f314df2a718572816d9e33be371ebfb827d0d55")
}

// UnpackBaseForgeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BaseForge__ZeroAddress()
func (forgeV3 *ForgeV3) UnpackBaseForgeZeroAddressError(raw []byte) (*ForgeV3BaseForgeZeroAddress, error) {
	out := new(ForgeV3BaseForgeZeroAddress)
	if err := forgeV3.abi.UnpackIntoInterface(out, "BaseForgeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3ERC721MintForgeInvalidRecipientSignature represents a ERC721MintForge__InvalidRecipientSignature error raised by the ForgeV3 contract.
type ForgeV3ERC721MintForgeInvalidRecipientSignature struct {
	Recipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721MintForge__InvalidRecipientSignature(address recipient)
func ForgeV3ERC721MintForgeInvalidRecipientSignatureErrorID() common.Hash {
	return common.HexToHash("0xcd3457babac82634ab3742a6056df59537d2eba6a801e097bd1a5c85f54409e6")
}

// UnpackERC721MintForgeInvalidRecipientSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721MintForge__InvalidRecipientSignature(address recipient)
func (forgeV3 *ForgeV3) UnpackERC721MintForgeInvalidRecipientSignatureError(raw []byte) (*ForgeV3ERC721MintForgeInvalidRecipientSignature, error) {
	out := new(ForgeV3ERC721MintForgeInvalidRecipientSignature)
	if err := forgeV3.abi.UnpackIntoInterface(out, "ERC721MintForgeInvalidRecipientSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3InvalidAccountNonce represents a InvalidAccountNonce error raised by the ForgeV3 contract.
type ForgeV3InvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ForgeV3InvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (forgeV3 *ForgeV3) UnpackInvalidAccountNonceError(raw []byte) (*ForgeV3InvalidAccountNonce, error) {
	out := new(ForgeV3InvalidAccountNonce)
	if err := forgeV3.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3InvalidInitialization represents a InvalidInitialization error raised by the ForgeV3 contract.
type ForgeV3InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ForgeV3InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (forgeV3 *ForgeV3) UnpackInvalidInitializationError(raw []byte) (*ForgeV3InvalidInitialization, error) {
	out := new(ForgeV3InvalidInitialization)
	if err := forgeV3.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3NotInitializing represents a NotInitializing error raised by the ForgeV3 contract.
type ForgeV3NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ForgeV3NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (forgeV3 *ForgeV3) UnpackNotInitializingError(raw []byte) (*ForgeV3NotInitializing, error) {
	out := new(ForgeV3NotInitializing)
	if err := forgeV3.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ForgeV3SafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ForgeV3 contract.
type ForgeV3SafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ForgeV3SafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (forgeV3 *ForgeV3) UnpackSafeERC20FailedOperationError(raw []byte) (*ForgeV3SafeERC20FailedOperation, error) {
	out := new(ForgeV3SafeERC20FailedOperation)
	if err := forgeV3.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
