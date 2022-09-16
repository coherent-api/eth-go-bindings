// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc721ABI is the input ABI used to generate the binding from.
const Erc721ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Erc721Bin is the compiled bytecode used for deploying new contracts.
var Erc721Bin = "60806040523480156200001157600080fd5b50604051620024fb380380620024fb83398181016040528101906200003791906200019f565b81600090805190602001906200004f92919062000071565b5080600190805190602001906200006892919062000071565b505050620003a8565b8280546200007f90620002b9565b90600052602060002090601f016020900481019282620000a35760008555620000ef565b82601f10620000be57805160ff1916838001178555620000ef565b82800160010185558215620000ef579182015b82811115620000ee578251825591602001919060010190620000d1565b5b509050620000fe919062000102565b5090565b5b808211156200011d57600081600090555060010162000103565b5090565b60006200013862000132846200024d565b62000224565b90508281526020810184848401111562000157576200015662000388565b5b6200016484828562000283565b509392505050565b600082601f83011262000184576200018362000383565b5b81516200019684826020860162000121565b91505092915050565b60008060408385031215620001b957620001b862000392565b5b600083015167ffffffffffffffff811115620001da57620001d96200038d565b5b620001e8858286016200016c565b925050602083015167ffffffffffffffff8111156200020c576200020b6200038d565b5b6200021a858286016200016c565b9150509250929050565b60006200023062000243565b90506200023e8282620002ef565b919050565b6000604051905090565b600067ffffffffffffffff8211156200026b576200026a62000354565b5b620002768262000397565b9050602081019050919050565b60005b83811015620002a357808201518184015260208101905062000286565b83811115620002b3576000848401525b50505050565b60006002820490506001821680620002d257607f821691505b60208210811415620002e957620002e862000325565b5b50919050565b620002fa8262000397565b810181811067ffffffffffffffff821117156200031c576200031b62000354565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b61214380620003b86000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb46514610224578063b88d4fde14610240578063c87b56dd1461025c578063e985e9c51461028c576100cf565b80636352211e146101a657806370a08231146101d657806395d89b4114610206576100cf565b806301ffc9a7146100d457806306fdde0314610104578063081812fc14610122578063095ea7b31461015257806323b872dd1461016e57806342842e0e1461018a575b600080fd5b6100ee60048036038101906100e9919061173d565b6102bc565b6040516100fb9190611a5a565b60405180910390f35b61010c61039e565b6040516101199190611a75565b60405180910390f35b61013c60048036038101906101379190611797565b610430565b60405161014991906119f3565b60405180910390f35b61016c600480360381019061016791906116fd565b610476565b005b610188600480360381019061018391906115e7565b61058e565b005b6101a4600480360381019061019f91906115e7565b6105ee565b005b6101c060048036038101906101bb9190611797565b61060e565b6040516101cd91906119f3565b60405180910390f35b6101f060048036038101906101eb919061157a565b610695565b6040516101fd9190611bb7565b60405180910390f35b61020e61074d565b60405161021b9190611a75565b60405180910390f35b61023e600480360381019061023991906116bd565b6107df565b005b61025a6004803603810190610255919061163a565b6107f5565b005b61027660048036038101906102719190611797565b610857565b6040516102839190611a75565b60405180910390f35b6102a660048036038101906102a191906115a7565b6108bf565b6040516102b39190611a5a565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610397575061039682610953565b5b9050919050565b6060600080546103ad90611d21565b80601f01602080910402602001604051908101604052809291908181526020018280546103d990611d21565b80156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b5050505050905090565b600061043b826109bd565b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104818261060e565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156104f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e990611b77565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610511610a08565b73ffffffffffffffffffffffffffffffffffffffff161480610540575061053f8161053a610a08565b6108bf565b5b61057f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057690611b97565b60405180910390fd5b6105898383610a10565b505050565b61059f610599610a08565b82610ac9565b6105de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d590611a97565b60405180910390fd5b6105e9838383610b5e565b505050565b610609838383604051806020016040528060008152506107f5565b505050565b60008061061a83610e54565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561068c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068390611b57565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610706576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106fd90611b37565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606001805461075c90611d21565b80601f016020809104026020016040519081016040528092919081815260200182805461078890611d21565b80156107d55780601f106107aa576101008083540402835291602001916107d5565b820191906000526020600020905b8154815290600101906020018083116107b857829003601f168201915b5050505050905090565b6107f16107ea610a08565b8383610e91565b5050565b610806610800610a08565b83610ac9565b610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083c90611a97565b60405180910390fd5b61085184848484610ffe565b50505050565b6060610862826109bd565b600061086c61105a565b9050600081511161088c57604051806020016040528060008152506108b7565b8061089684611071565b6040516020016108a79291906119cf565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6109c681611149565b610a05576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fc90611b57565b60405180910390fd5b50565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610a838361060e565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610ad58361060e565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610b175750610b1681856108bf565b5b80610b5557508373ffffffffffffffffffffffffffffffffffffffff16610b3d84610430565b73ffffffffffffffffffffffffffffffffffffffff16145b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610b7e8261060e565b73ffffffffffffffffffffffffffffffffffffffff1614610bd4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bcb90611ad7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3b90611af7565b60405180910390fd5b610c4f83838361118a565b8273ffffffffffffffffffffffffffffffffffffffff16610c6f8261060e565b73ffffffffffffffffffffffffffffffffffffffff1614610cc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbc90611ad7565b60405180910390fd5b6004600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4610e4f83838361118f565b505050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610f00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef790611b17565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610ff19190611a5a565b60405180910390a3505050565b611009848484610b5e565b61101584848484611194565b611054576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104b90611ab7565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b6060600060016110808461132b565b01905060008167ffffffffffffffff81111561109f5761109e611de2565b5b6040519080825280601f01601f1916602001820160405280156110d15781602001600182028036833780820191505090505b509050600082602001820190505b60011561113e578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a858161112857611127611d84565b5b04945060008514156111395761113e565b6110df565b819350505050919050565b60008073ffffffffffffffffffffffffffffffffffffffff1661116b83610e54565b73ffffffffffffffffffffffffffffffffffffffff1614159050919050565b505050565b505050565b60006111b58473ffffffffffffffffffffffffffffffffffffffff1661147e565b1561131e578373ffffffffffffffffffffffffffffffffffffffff1663150b7a026111de610a08565b8786866040518563ffffffff1660e01b81526004016112009493929190611a0e565b602060405180830381600087803b15801561121a57600080fd5b505af192505050801561124b57506040513d601f19601f82011682018060405250810190611248919061176a565b60015b6112ce573d806000811461127b576040519150601f19603f3d011682016040523d82523d6000602084013e611280565b606091505b506000815114156112c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112bd90611ab7565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050611323565b600190505b949350505050565b600080600090507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611389577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161137f5761137e611d84565b5b0492506040810190505b6d04ee2d6d415b85acef810000000083106113c6576d04ee2d6d415b85acef810000000083816113bc576113bb611d84565b5b0492506020810190505b662386f26fc1000083106113f557662386f26fc1000083816113eb576113ea611d84565b5b0492506010810190505b6305f5e100831061141e576305f5e100838161141457611413611d84565b5b0492506008810190505b612710831061144357612710838161143957611438611d84565b5b0492506004810190505b60648310611466576064838161145c5761145b611d84565b5b0492506002810190505b600a8310611475576001810190505b80915050919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60006114b46114af84611bf7565b611bd2565b9050828152602081018484840111156114d0576114cf611e16565b5b6114db848285611cdf565b509392505050565b6000813590506114f2816120b1565b92915050565b600081359050611507816120c8565b92915050565b60008135905061151c816120df565b92915050565b600081519050611531816120df565b92915050565b600082601f83011261154c5761154b611e11565b5b813561155c8482602086016114a1565b91505092915050565b600081359050611574816120f6565b92915050565b6000602082840312156115905761158f611e20565b5b600061159e848285016114e3565b91505092915050565b600080604083850312156115be576115bd611e20565b5b60006115cc858286016114e3565b92505060206115dd858286016114e3565b9150509250929050565b600080600060608486031215611600576115ff611e20565b5b600061160e868287016114e3565b935050602061161f868287016114e3565b925050604061163086828701611565565b9150509250925092565b6000806000806080858703121561165457611653611e20565b5b6000611662878288016114e3565b9450506020611673878288016114e3565b935050604061168487828801611565565b925050606085013567ffffffffffffffff8111156116a5576116a4611e1b565b5b6116b187828801611537565b91505092959194509250565b600080604083850312156116d4576116d3611e20565b5b60006116e2858286016114e3565b92505060206116f3858286016114f8565b9150509250929050565b6000806040838503121561171457611713611e20565b5b6000611722858286016114e3565b925050602061173385828601611565565b9150509250929050565b60006020828403121561175357611752611e20565b5b60006117618482850161150d565b91505092915050565b6000602082840312156117805761177f611e20565b5b600061178e84828501611522565b91505092915050565b6000602082840312156117ad576117ac611e20565b5b60006117bb84828501611565565b91505092915050565b6117cd81611c6b565b82525050565b6117dc81611c7d565b82525050565b60006117ed82611c28565b6117f78185611c3e565b9350611807818560208601611cee565b61181081611e25565b840191505092915050565b600061182682611c33565b6118308185611c4f565b9350611840818560208601611cee565b61184981611e25565b840191505092915050565b600061185f82611c33565b6118698185611c60565b9350611879818560208601611cee565b80840191505092915050565b6000611892602d83611c4f565b915061189d82611e36565b604082019050919050565b60006118b5603283611c4f565b91506118c082611e85565b604082019050919050565b60006118d8602583611c4f565b91506118e382611ed4565b604082019050919050565b60006118fb602483611c4f565b915061190682611f23565b604082019050919050565b600061191e601983611c4f565b915061192982611f72565b602082019050919050565b6000611941602983611c4f565b915061194c82611f9b565b604082019050919050565b6000611964601883611c4f565b915061196f82611fea565b602082019050919050565b6000611987602183611c4f565b915061199282612013565b604082019050919050565b60006119aa603d83611c4f565b91506119b582612062565b604082019050919050565b6119c981611cd5565b82525050565b60006119db8285611854565b91506119e78284611854565b91508190509392505050565b6000602082019050611a0860008301846117c4565b92915050565b6000608082019050611a2360008301876117c4565b611a3060208301866117c4565b611a3d60408301856119c0565b8181036060830152611a4f81846117e2565b905095945050505050565b6000602082019050611a6f60008301846117d3565b92915050565b60006020820190508181036000830152611a8f818461181b565b905092915050565b60006020820190508181036000830152611ab081611885565b9050919050565b60006020820190508181036000830152611ad0816118a8565b9050919050565b60006020820190508181036000830152611af0816118cb565b9050919050565b60006020820190508181036000830152611b10816118ee565b9050919050565b60006020820190508181036000830152611b3081611911565b9050919050565b60006020820190508181036000830152611b5081611934565b9050919050565b60006020820190508181036000830152611b7081611957565b9050919050565b60006020820190508181036000830152611b908161197a565b9050919050565b60006020820190508181036000830152611bb08161199d565b9050919050565b6000602082019050611bcc60008301846119c0565b92915050565b6000611bdc611bed565b9050611be88282611d53565b919050565b6000604051905090565b600067ffffffffffffffff821115611c1257611c11611de2565b5b611c1b82611e25565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c7682611cb5565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611d0c578082015181840152602081019050611cf1565b83811115611d1b576000848401525b50505050565b60006002820490506001821680611d3957607f821691505b60208210811415611d4d57611d4c611db3565b5b50919050565b611d5c82611e25565b810181811067ffffffffffffffff82111715611d7b57611d7a611de2565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560008201527f72206f7220617070726f76656400000000000000000000000000000000000000602082015250565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b7f4552433732313a207472616e736665722066726f6d20696e636f72726563742060008201527f6f776e6572000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b7f4552433732313a2061646472657373207a65726f206973206e6f74206120766160008201527f6c6964206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20696e76616c696420746f6b656e2049440000000000000000600082015250565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60008201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000602082015250565b6120ba81611c6b565b81146120c557600080fd5b50565b6120d181611c7d565b81146120dc57600080fd5b50565b6120e881611c89565b81146120f357600080fd5b50565b6120ff81611cd5565b811461210a57600080fd5b5056fea264697066735822122005bad367ceb9dabc164da3b24ec50c7956a04486f3ea0d13c5a4f71ae8f0702864736f6c63430008070033"

// DeployErc721 deploys a new Ethereum contract, binding an instance of Erc721 to it.
func DeployErc721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Erc721, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Erc721Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc721{Erc721Caller: Erc721Caller{contract: contract}, Erc721Transactor: Erc721Transactor{contract: contract}, Erc721Filterer: Erc721Filterer{contract: contract}}, nil
}

// Erc721 is an auto generated Go binding around an Ethereum contract.
type Erc721 struct {
	Erc721Caller     // Read-only binding to the contract
	Erc721Transactor // Write-only binding to the contract
	Erc721Filterer   // Log filterer for contract events
}

// Erc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721Session struct {
	Contract     *Erc721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721CallerSession struct {
	Contract *Erc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721TransactorSession struct {
	Contract     *Erc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721Raw struct {
	Contract *Erc721 // Generic contract binding to access the raw methods on
}

// Erc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721CallerRaw struct {
	Contract *Erc721Caller // Generic read-only contract binding to access the raw methods on
}

// Erc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721TransactorRaw struct {
	Contract *Erc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721 creates a new instance of Erc721, bound to a specific deployed contract.
func NewErc721(address common.Address, backend bind.ContractBackend) (*Erc721, error) {
	contract, err := bindErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721{Erc721Caller: Erc721Caller{contract: contract}, Erc721Transactor: Erc721Transactor{contract: contract}, Erc721Filterer: Erc721Filterer{contract: contract}}, nil
}

// NewErc721Caller creates a new read-only instance of Erc721, bound to a specific deployed contract.
func NewErc721Caller(address common.Address, caller bind.ContractCaller) (*Erc721Caller, error) {
	contract, err := bindErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721Caller{contract: contract}, nil
}

// NewErc721Transactor creates a new write-only instance of Erc721, bound to a specific deployed contract.
func NewErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc721Transactor, error) {
	contract, err := bindErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721Transactor{contract: contract}, nil
}

// NewErc721Filterer creates a new log filterer instance of Erc721, bound to a specific deployed contract.
func NewErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc721Filterer, error) {
	contract, err := bindErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721Filterer{contract: contract}, nil
}

// bindErc721 binds a generic wrapper to an already deployed contract.
func bindErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721 *Erc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721.Contract.Erc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721 *Erc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.Contract.Erc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721 *Erc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721.Contract.Erc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721 *Erc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721 *Erc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721 *Erc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721.Contract.BalanceOf(&_Erc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721.Contract.BalanceOf(&_Erc721.CallOpts, owner)
}

//TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
//Solidity: function totalSupply() constant returns(uint256)
func (_Erc721 *Erc721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "totalSupply")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}

//TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
//Solidity: function totalSupply() constant returns(uint256)
func (_Erc721 *Erc721Session) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	return _Erc721.Contract.TotalSupply(&_Erc721.CallOpts)
}

//TotalSupply is a free data retrieval call binding the Contract method 0x18160ddd.
//
//Solidity: function totalSupply() constant returns(uint256)
func (_Erc721 *Erc721CallerSession) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	return _Erc721.Contract.TotalSupply(&_Erc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.GetApproved(&_Erc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.GetApproved(&_Erc721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721.Contract.IsApprovedForAll(&_Erc721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721.Contract.IsApprovedForAll(&_Erc721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721Session) Name() (string, error) {
	return _Erc721.Contract.Name(&_Erc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721CallerSession) Name() (string, error) {
	return _Erc721.Contract.Name(&_Erc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.OwnerOf(&_Erc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.OwnerOf(&_Erc721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721.Contract.SupportsInterface(&_Erc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721.Contract.SupportsInterface(&_Erc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721Session) Symbol() (string, error) {
	return _Erc721.Contract.Symbol(&_Erc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721CallerSession) Symbol() (string, error) {
	return _Erc721.Contract.Symbol(&_Erc721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721.Contract.TokenURI(&_Erc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721.Contract.TokenURI(&_Erc721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.Approve(&_Erc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.Approve(&_Erc721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom0(&_Erc721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom0(&_Erc721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetApprovalForAll(&_Erc721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetApprovalForAll(&_Erc721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.TransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.TransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// Erc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721 contract.
type Erc721ApprovalIterator struct {
	Event *Erc721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721Approval represents a Approval event raised by the Erc721 contract.
type Erc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ApprovalIterator{contract: _Erc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721Approval)
				if err := _Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) ParseApproval(log types.Log) (*Erc721Approval, error) {
	event := new(Erc721Approval)
	if err := _Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc721 contract.
type Erc721ApprovalForAllIterator struct {
	Event *Erc721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ApprovalForAll represents a ApprovalForAll event raised by the Erc721 contract.
type Erc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ApprovalForAllIterator{contract: _Erc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ApprovalForAll)
				if err := _Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) ParseApprovalForAll(log types.Log) (*Erc721ApprovalForAll, error) {
	event := new(Erc721ApprovalForAll)
	if err := _Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721 contract.
type Erc721TransferIterator struct {
	Event *Erc721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721Transfer represents a Transfer event raised by the Erc721 contract.
type Erc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721TransferIterator{contract: _Erc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721Transfer)
				if err := _Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) ParseTransfer(log types.Log) (*Erc721Transfer, error) {
	event := new(Erc721Transfer)
	if err := _Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
