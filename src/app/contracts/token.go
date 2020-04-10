// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BuyerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BuyerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"LogicAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"LogicRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SellerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SellerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBuyer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addLogic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addSeller\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountBuyers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLastSubsidy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBuyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLogic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSeller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceBuyer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceSeller\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"setLastSubsidy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405260006006553480156200001657600080fd5b5060405162001afb38038062001afb833981810160405260608110156200003c57600080fd5b81019080805160405193929190846401000000008211156200005d57600080fd5b9083019060208201858111156200007357600080fd5b82516401000000008111828201881017156200008e57600080fd5b82525081516020918201929091019080838360005b83811015620000bd578181015183820152602001620000a3565b50505050905090810190601f168015620000eb5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010f57600080fd5b9083019060208201858111156200012557600080fd5b82516401000000008111828201881017156200014057600080fd5b82525081516020918201929091019080838360005b838110156200016f57818101518382015260200162000155565b50505050905090810190601f1680156200019d5780820380516001836020036101000a031916815260200191505b506040526020015191508390508282620001d9620001c36001600160e01b036200024b16565b60036200025060201b6200136e1790919060201c565b620001ff620001f06001600160e01b036200024b16565b6001600160e01b03620002dd16565b82516200021490600990602086019062000398565b5081516200022a90600a90602085019062000398565b50600b805460ff191660ff92909216919091179055506200043a9350505050565b335b90565b6200026582826001600160e01b036200032f16565b15620002b8576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b620002f88160046200025060201b6200136e1790919060201c565b6040516001600160a01b038216907fce0ca4df7a7ec71282174625246800d852b6926b9760ef984c9fa6611e71724a90600090a250565b60006001600160a01b038216620003785760405162461bcd60e51b815260040180806020018281038252602281526020018062001ad96022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003db57805160ff19168380011785556200040b565b828001600101855582156200040b579182015b828111156200040b578251825591602001919060010190620003ee565b50620004199291506200041d565b5090565b6200024d91905b8082111562000419576000815560010162000424565b61168f806200044a6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806379cc6790116100f9578063bbf8980311610097578063dd62ed3e11610071578063dd62ed3e146104db578063ee95feaf14610509578063f3ae24151461052f578063f8b91abe14610555576101a9565b8063bbf8980314610487578063c204a0cf146104ad578063d93fabfa146104b5576101a9565b806395d89b41116100d357806395d89b411461044b5780639abd2d0614610453578063a457c2d71461022b578063a9059cbb1461045b576101a9565b806379cc6790146103cd57806379ef704e146103f95780637b881f811461041f576101a9565b80632cf56891116101665780633950935111610140578063395093511461022b57806340c10f191461037357806364c78f8c1461039f57806370a08231146103a7576101a9565b80632cf56891146103075780632d06177a1461032d578063313ce56714610355576101a9565b806306fdde03146101ae578063095ea7b31461022b57806318160ddd1461026b57806323b872dd14610285578063257f0556146102bb5780632a55feec146102e1575b600080fd5b6101b661055d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f05781810151838201526020016101d8565b50505050905090810190601f16801561021d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102576004803603604081101561024157600080fd5b506001600160a01b0381351690602001356105f3565b604080519115158252519081900360200190f35b6102736105fb565b60408051918252519081900360200190f35b6102576004803603606081101561029b57600080fd5b506001600160a01b03813581169160208101359091169060400135610601565b610273600480360360208110156102d157600080fd5b50356001600160a01b031661068e565b610257600480360360208110156102f757600080fd5b50356001600160a01b03166106f7565b6102576004803603602081101561031d57600080fd5b50356001600160a01b0316610710565b6103536004803603602081101561034357600080fd5b50356001600160a01b0316610723565b005b61035d61077a565b6040805160ff9092168252519081900360200190f35b6102576004803603604081101561038957600080fd5b506001600160a01b038135169060200135610783565b6103536107de565b610273600480360360208110156103bd57600080fd5b50356001600160a01b03166107f0565b610353600480360360408110156103e357600080fd5b506001600160a01b03813516906020013561080b565b6103536004803603602081101561040f57600080fd5b50356001600160a01b031661085f565b6102576004803603604081101561043557600080fd5b506001600160a01b0381351690602001356108be565b6101b6610923565b610273610984565b6102576004803603604081101561047157600080fd5b506001600160a01b03813516906020013561098a565b6103536004803603602081101561049d57600080fd5b50356001600160a01b0316610a68565b610353610ab7565b610353600480360360208110156104cb57600080fd5b50356001600160a01b0316610ac7565b610273600480360360408110156104f157600080fd5b506001600160a01b0381358116916020013516610b16565b6102576004803603602081101561051f57600080fd5b50356001600160a01b0316610b41565b6102576004803603602081101561054557600080fd5b50356001600160a01b0316610b54565b610353610b67565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105e95780601f106105be576101008083540402835291602001916105e9565b820191906000526020600020905b8154815290600101906020018083116105cc57829003601f168201915b5050505050905090565b600092915050565b60025490565b600061060e848484610b77565b6106848461061a610cd3565b61067f856040518060600160405280602881526020016115a7602891396001600160a01b038a16600090815260016020526040812090610658610cd3565b6001600160a01b03168152602081019190915260400160002054919063ffffffff610cd716565b610d6e565b5060019392505050565b60006106a061069b610cd3565b610710565b6106db5760405162461bcd60e51b815260040180806020018281038252602e815260200180611500602e913960400191505060405180910390fd5b506001600160a01b031660009081526007602052604090205490565b600061070a60058363ffffffff610e5a16565b92915050565b600061070a60048363ffffffff610e5a16565b61073361072e610cd3565b610b54565b61076e5760405162461bcd60e51b81526004018080602001828103825260328152602001806115756032913960400191505060405180910390fd5b61077781610ec1565b50565b600b5460ff1690565b600061079061069b610cd3565b6107cb5760405162461bcd60e51b815260040180806020018281038252602e815260200180611500602e913960400191505060405180910390fd5b6107d58383610f09565b50600192915050565b6107ee6107e9610cd3565b610ff9565b565b6001600160a01b031660009081526020819052604090205490565b61081661069b610cd3565b6108515760405162461bcd60e51b815260040180806020018281038252602e815260200180611500602e913960400191505060405180910390fd5b61085b8282611041565b5050565b61086a61069b610cd3565b6108a55760405162461bcd60e51b815260040180806020018281038252602e815260200180611500602e913960400191505060405180910390fd5b6108ae8161113d565b6107776108b9610cd3565b611185565b60006108cb61069b610cd3565b6109065760405162461bcd60e51b815260040180806020018281038252602e815260200180611500602e913960400191505060405180910390fd5b6001600160a01b0390921660009081526007602052604090205590565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105e95780601f106105be576101008083540402835291602001916105e9565b60065490565b6000610995336106f7565b806109a457506109a433610710565b6109f5576040805162461bcd60e51b815260206004820152601760248201527f4e6f7420616c6c6f77656420746f207472616e73666572000000000000000000604482015290519081900360640190fd5b6109fe83610b41565b80610a0d5750610a0d33610710565b610a57576040805162461bcd60e51b81526020600482015260166024820152754e6f7420616c6c6f77656420746f207265636569766560501b604482015290519081900360640190fd5b610a6183836111cd565b9392505050565b610a7361072e610cd3565b610aae5760405162461bcd60e51b81526004018080602001828103825260328152602001806115756032913960400191505060405180910390fd5b610777816111e1565b6107ee610ac2610cd3565b611232565b610ad261072e610cd3565b610b0d5760405162461bcd60e51b81526004018080602001828103825260328152602001806115756032913960400191505060405180910390fd5b61077781611284565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600061070a60088363ffffffff610e5a16565b600061070a60038363ffffffff610e5a16565b6107ee610b72610cd3565b6112cc565b6001600160a01b038316610bbc5760405162461bcd60e51b81526004018080602001828103825260258152602001806116126025913960400191505060405180910390fd5b6001600160a01b038216610c015760405162461bcd60e51b81526004018080602001828103825260238152602001806114996023913960400191505060405180910390fd5b610c448160405180606001604052806026815260200161152e602691396001600160a01b038616600090815260208190526040902054919063ffffffff610cd716565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610c79908263ffffffff61131416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b3390565b60008184841115610d665760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d2b578181015183820152602001610d13565b50505050905090810190601f168015610d585780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b038316610db35760405162461bcd60e51b81526004018080602001828103825260248152602001806116376024913960400191505060405180910390fd5b6001600160a01b038216610df85760405162461bcd60e51b81526004018080602001828103825260228152602001806114de6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b60006001600160a01b038216610ea15760405162461bcd60e51b81526004018080602001828103825260228152602001806115cf6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610ed260038263ffffffff61136e16565b6040516001600160a01b038216907f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a90600090a250565b6001600160a01b038216610f64576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610f77908263ffffffff61131416565b6002556001600160a01b038216600090815260208190526040902054610fa3908263ffffffff61131416565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b61100a60088263ffffffff6113ef16565b6040516001600160a01b038216907fa6f3b3030df964d974376ba4db21cbe49207bd79998f0d619c1c3a422259707f90600090a250565b6001600160a01b0382166110865760405162461bcd60e51b81526004018080602001828103825260218152602001806115f16021913960400191505060405180910390fd5b6110c9816040518060600160405280602281526020016114bc602291396001600160a01b038516600090815260208190526040902054919063ffffffff610cd716565b6001600160a01b0383166000908152602081905260409020556002546110f5908263ffffffff61145616565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b61114e60048263ffffffff61136e16565b6040516001600160a01b038216907fce0ca4df7a7ec71282174625246800d852b6926b9760ef984c9fa6611e71724a90600090a250565b61119660048263ffffffff6113ef16565b6040516001600160a01b038216907f9f60fd9ac36634c8a086fb62b8d5f9b088ed08a56805bb49b467012264d1444c90600090a250565b60006107d56111da610cd3565b8484610b77565b6111f260058263ffffffff61136e16565b6006805460010190556040516001600160a01b038216907f6599d9dd7f97ccdc3d99ab507544f9f6ebfca20c63a7e3c018da49eeaefb164a90600090a250565b61124360058263ffffffff6113ef16565b600680546000190190556040516001600160a01b038216907f83fd4cbc7bc934bdfbae994ac6e72fcc0e50abd6995c1949520e761ba650d58290600090a250565b61129560088263ffffffff61136e16565b6040516001600160a01b038216907f3d78526f2568d430fd8e20719601f7bc3f979936f8880236cc284e8127251b4290600090a250565b6112dd60038263ffffffff6113ef16565b6040516001600160a01b038216907fef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd3190600090a250565b600082820183811015610a61576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6113788282610e5a565b156113ca576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6113f98282610e5a565b6114345760405162461bcd60e51b81526004018080602001828103825260218152602001806115546021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b6000610a6183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610cd756fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f20616464726573734c6f676963526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204c6f67696320726f6c6545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654d616e61676572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d616e6167657220726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a72315820f1b8b83f35289fa5bfd90c8d92d8e6fa238d034c53b8f80cfb935ea02dd1255864736f6c634300050f0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Store *StoreCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Store *StoreSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Store *StoreCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// AmountBuyers is a free data retrieval call binding the contract method 0x9abd2d06.
//
// Solidity: function amountBuyers() constant returns(uint256)
func (_Store *StoreCaller) AmountBuyers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "amountBuyers")
	return *ret0, err
}

// AmountBuyers is a free data retrieval call binding the contract method 0x9abd2d06.
//
// Solidity: function amountBuyers() constant returns(uint256)
func (_Store *StoreSession) AmountBuyers() (*big.Int, error) {
	return _Store.Contract.AmountBuyers(&_Store.CallOpts)
}

// AmountBuyers is a free data retrieval call binding the contract method 0x9abd2d06.
//
// Solidity: function amountBuyers() constant returns(uint256)
func (_Store *StoreCallerSession) AmountBuyers() (*big.Int, error) {
	return _Store.Contract.AmountBuyers(&_Store.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Store *StoreSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Store *StoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Store *StoreSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Store *StoreCallerSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// IsBuyer is a free data retrieval call binding the contract method 0x2a55feec.
//
// Solidity: function isBuyer(address account) constant returns(bool)
func (_Store *StoreCaller) IsBuyer(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isBuyer", account)
	return *ret0, err
}

// IsBuyer is a free data retrieval call binding the contract method 0x2a55feec.
//
// Solidity: function isBuyer(address account) constant returns(bool)
func (_Store *StoreSession) IsBuyer(account common.Address) (bool, error) {
	return _Store.Contract.IsBuyer(&_Store.CallOpts, account)
}

// IsBuyer is a free data retrieval call binding the contract method 0x2a55feec.
//
// Solidity: function isBuyer(address account) constant returns(bool)
func (_Store *StoreCallerSession) IsBuyer(account common.Address) (bool, error) {
	return _Store.Contract.IsBuyer(&_Store.CallOpts, account)
}

// IsLogic is a free data retrieval call binding the contract method 0x2cf56891.
//
// Solidity: function isLogic(address account) constant returns(bool)
func (_Store *StoreCaller) IsLogic(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isLogic", account)
	return *ret0, err
}

// IsLogic is a free data retrieval call binding the contract method 0x2cf56891.
//
// Solidity: function isLogic(address account) constant returns(bool)
func (_Store *StoreSession) IsLogic(account common.Address) (bool, error) {
	return _Store.Contract.IsLogic(&_Store.CallOpts, account)
}

// IsLogic is a free data retrieval call binding the contract method 0x2cf56891.
//
// Solidity: function isLogic(address account) constant returns(bool)
func (_Store *StoreCallerSession) IsLogic(account common.Address) (bool, error) {
	return _Store.Contract.IsLogic(&_Store.CallOpts, account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) constant returns(bool)
func (_Store *StoreCaller) IsManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isManager", account)
	return *ret0, err
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) constant returns(bool)
func (_Store *StoreSession) IsManager(account common.Address) (bool, error) {
	return _Store.Contract.IsManager(&_Store.CallOpts, account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) constant returns(bool)
func (_Store *StoreCallerSession) IsManager(account common.Address) (bool, error) {
	return _Store.Contract.IsManager(&_Store.CallOpts, account)
}

// IsSeller is a free data retrieval call binding the contract method 0xee95feaf.
//
// Solidity: function isSeller(address account) constant returns(bool)
func (_Store *StoreCaller) IsSeller(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isSeller", account)
	return *ret0, err
}

// IsSeller is a free data retrieval call binding the contract method 0xee95feaf.
//
// Solidity: function isSeller(address account) constant returns(bool)
func (_Store *StoreSession) IsSeller(account common.Address) (bool, error) {
	return _Store.Contract.IsSeller(&_Store.CallOpts, account)
}

// IsSeller is a free data retrieval call binding the contract method 0xee95feaf.
//
// Solidity: function isSeller(address account) constant returns(bool)
func (_Store *StoreCallerSession) IsSeller(account common.Address) (bool, error) {
	return _Store.Contract.IsSeller(&_Store.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// AddBuyer is a paid mutator transaction binding the contract method 0xbbf89803.
//
// Solidity: function addBuyer(address account) returns()
func (_Store *StoreTransactor) AddBuyer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addBuyer", account)
}

// AddBuyer is a paid mutator transaction binding the contract method 0xbbf89803.
//
// Solidity: function addBuyer(address account) returns()
func (_Store *StoreSession) AddBuyer(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBuyer(&_Store.TransactOpts, account)
}

// AddBuyer is a paid mutator transaction binding the contract method 0xbbf89803.
//
// Solidity: function addBuyer(address account) returns()
func (_Store *StoreTransactorSession) AddBuyer(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBuyer(&_Store.TransactOpts, account)
}

// AddLogic is a paid mutator transaction binding the contract method 0x79ef704e.
//
// Solidity: function addLogic(address account) returns()
func (_Store *StoreTransactor) AddLogic(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addLogic", account)
}

// AddLogic is a paid mutator transaction binding the contract method 0x79ef704e.
//
// Solidity: function addLogic(address account) returns()
func (_Store *StoreSession) AddLogic(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddLogic(&_Store.TransactOpts, account)
}

// AddLogic is a paid mutator transaction binding the contract method 0x79ef704e.
//
// Solidity: function addLogic(address account) returns()
func (_Store *StoreTransactorSession) AddLogic(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddLogic(&_Store.TransactOpts, account)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_Store *StoreTransactor) AddManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addManager", account)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_Store *StoreSession) AddManager(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddManager(&_Store.TransactOpts, account)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_Store *StoreTransactorSession) AddManager(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddManager(&_Store.TransactOpts, account)
}

// AddSeller is a paid mutator transaction binding the contract method 0xd93fabfa.
//
// Solidity: function addSeller(address account) returns()
func (_Store *StoreTransactor) AddSeller(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addSeller", account)
}

// AddSeller is a paid mutator transaction binding the contract method 0xd93fabfa.
//
// Solidity: function addSeller(address account) returns()
func (_Store *StoreSession) AddSeller(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddSeller(&_Store.TransactOpts, account)
}

// AddSeller is a paid mutator transaction binding the contract method 0xd93fabfa.
//
// Solidity: function addSeller(address account) returns()
func (_Store *StoreTransactorSession) AddSeller(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddSeller(&_Store.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Store *StoreTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Store *StoreSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnFrom(&_Store.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Store *StoreTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnFrom(&_Store.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// GetLastSubsidy is a paid mutator transaction binding the contract method 0x257f0556.
//
// Solidity: function getLastSubsidy(address account) returns(uint256)
func (_Store *StoreTransactor) GetLastSubsidy(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getLastSubsidy", account)
}

// GetLastSubsidy is a paid mutator transaction binding the contract method 0x257f0556.
//
// Solidity: function getLastSubsidy(address account) returns(uint256)
func (_Store *StoreSession) GetLastSubsidy(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GetLastSubsidy(&_Store.TransactOpts, account)
}

// GetLastSubsidy is a paid mutator transaction binding the contract method 0x257f0556.
//
// Solidity: function getLastSubsidy(address account) returns(uint256)
func (_Store *StoreTransactorSession) GetLastSubsidy(account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GetLastSubsidy(&_Store.TransactOpts, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Store *StoreSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, account, amount)
}

// RenounceBuyer is a paid mutator transaction binding the contract method 0xc204a0cf.
//
// Solidity: function renounceBuyer() returns()
func (_Store *StoreTransactor) RenounceBuyer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceBuyer")
}

// RenounceBuyer is a paid mutator transaction binding the contract method 0xc204a0cf.
//
// Solidity: function renounceBuyer() returns()
func (_Store *StoreSession) RenounceBuyer() (*types.Transaction, error) {
	return _Store.Contract.RenounceBuyer(&_Store.TransactOpts)
}

// RenounceBuyer is a paid mutator transaction binding the contract method 0xc204a0cf.
//
// Solidity: function renounceBuyer() returns()
func (_Store *StoreTransactorSession) RenounceBuyer() (*types.Transaction, error) {
	return _Store.Contract.RenounceBuyer(&_Store.TransactOpts)
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_Store *StoreTransactor) RenounceManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceManager")
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_Store *StoreSession) RenounceManager() (*types.Transaction, error) {
	return _Store.Contract.RenounceManager(&_Store.TransactOpts)
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_Store *StoreTransactorSession) RenounceManager() (*types.Transaction, error) {
	return _Store.Contract.RenounceManager(&_Store.TransactOpts)
}

// RenounceSeller is a paid mutator transaction binding the contract method 0x64c78f8c.
//
// Solidity: function renounceSeller() returns()
func (_Store *StoreTransactor) RenounceSeller(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceSeller")
}

// RenounceSeller is a paid mutator transaction binding the contract method 0x64c78f8c.
//
// Solidity: function renounceSeller() returns()
func (_Store *StoreSession) RenounceSeller() (*types.Transaction, error) {
	return _Store.Contract.RenounceSeller(&_Store.TransactOpts)
}

// RenounceSeller is a paid mutator transaction binding the contract method 0x64c78f8c.
//
// Solidity: function renounceSeller() returns()
func (_Store *StoreTransactorSession) RenounceSeller() (*types.Transaction, error) {
	return _Store.Contract.RenounceSeller(&_Store.TransactOpts)
}

// SetLastSubsidy is a paid mutator transaction binding the contract method 0x7b881f81.
//
// Solidity: function setLastSubsidy(address account, uint256 day) returns(bool)
func (_Store *StoreTransactor) SetLastSubsidy(opts *bind.TransactOpts, account common.Address, day *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setLastSubsidy", account, day)
}

// SetLastSubsidy is a paid mutator transaction binding the contract method 0x7b881f81.
//
// Solidity: function setLastSubsidy(address account, uint256 day) returns(bool)
func (_Store *StoreSession) SetLastSubsidy(account common.Address, day *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastSubsidy(&_Store.TransactOpts, account, day)
}

// SetLastSubsidy is a paid mutator transaction binding the contract method 0x7b881f81.
//
// Solidity: function setLastSubsidy(address account, uint256 day) returns(bool)
func (_Store *StoreTransactorSession) SetLastSubsidy(account common.Address, day *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastSubsidy(&_Store.TransactOpts, account, day)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, sender, recipient, amount)
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

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
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
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
		it.Event = new(StoreApproval)
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
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreBuyerAddedIterator is returned from FilterBuyerAdded and is used to iterate over the raw logs and unpacked data for BuyerAdded events raised by the Store contract.
type StoreBuyerAddedIterator struct {
	Event *StoreBuyerAdded // Event containing the contract specifics and raw log

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
func (it *StoreBuyerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBuyerAdded)
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
		it.Event = new(StoreBuyerAdded)
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
func (it *StoreBuyerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBuyerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBuyerAdded represents a BuyerAdded event raised by the Store contract.
type StoreBuyerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyerAdded is a free log retrieval operation binding the contract event 0x6599d9dd7f97ccdc3d99ab507544f9f6ebfca20c63a7e3c018da49eeaefb164a.
//
// Solidity: event BuyerAdded(address indexed account)
func (_Store *StoreFilterer) FilterBuyerAdded(opts *bind.FilterOpts, account []common.Address) (*StoreBuyerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "BuyerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreBuyerAddedIterator{contract: _Store.contract, event: "BuyerAdded", logs: logs, sub: sub}, nil
}

// WatchBuyerAdded is a free log subscription operation binding the contract event 0x6599d9dd7f97ccdc3d99ab507544f9f6ebfca20c63a7e3c018da49eeaefb164a.
//
// Solidity: event BuyerAdded(address indexed account)
func (_Store *StoreFilterer) WatchBuyerAdded(opts *bind.WatchOpts, sink chan<- *StoreBuyerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "BuyerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBuyerAdded)
				if err := _Store.contract.UnpackLog(event, "BuyerAdded", log); err != nil {
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

// ParseBuyerAdded is a log parse operation binding the contract event 0x6599d9dd7f97ccdc3d99ab507544f9f6ebfca20c63a7e3c018da49eeaefb164a.
//
// Solidity: event BuyerAdded(address indexed account)
func (_Store *StoreFilterer) ParseBuyerAdded(log types.Log) (*StoreBuyerAdded, error) {
	event := new(StoreBuyerAdded)
	if err := _Store.contract.UnpackLog(event, "BuyerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreBuyerRemovedIterator is returned from FilterBuyerRemoved and is used to iterate over the raw logs and unpacked data for BuyerRemoved events raised by the Store contract.
type StoreBuyerRemovedIterator struct {
	Event *StoreBuyerRemoved // Event containing the contract specifics and raw log

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
func (it *StoreBuyerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBuyerRemoved)
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
		it.Event = new(StoreBuyerRemoved)
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
func (it *StoreBuyerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBuyerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBuyerRemoved represents a BuyerRemoved event raised by the Store contract.
type StoreBuyerRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyerRemoved is a free log retrieval operation binding the contract event 0x83fd4cbc7bc934bdfbae994ac6e72fcc0e50abd6995c1949520e761ba650d582.
//
// Solidity: event BuyerRemoved(address indexed account)
func (_Store *StoreFilterer) FilterBuyerRemoved(opts *bind.FilterOpts, account []common.Address) (*StoreBuyerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "BuyerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreBuyerRemovedIterator{contract: _Store.contract, event: "BuyerRemoved", logs: logs, sub: sub}, nil
}

// WatchBuyerRemoved is a free log subscription operation binding the contract event 0x83fd4cbc7bc934bdfbae994ac6e72fcc0e50abd6995c1949520e761ba650d582.
//
// Solidity: event BuyerRemoved(address indexed account)
func (_Store *StoreFilterer) WatchBuyerRemoved(opts *bind.WatchOpts, sink chan<- *StoreBuyerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "BuyerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBuyerRemoved)
				if err := _Store.contract.UnpackLog(event, "BuyerRemoved", log); err != nil {
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

// ParseBuyerRemoved is a log parse operation binding the contract event 0x83fd4cbc7bc934bdfbae994ac6e72fcc0e50abd6995c1949520e761ba650d582.
//
// Solidity: event BuyerRemoved(address indexed account)
func (_Store *StoreFilterer) ParseBuyerRemoved(log types.Log) (*StoreBuyerRemoved, error) {
	event := new(StoreBuyerRemoved)
	if err := _Store.contract.UnpackLog(event, "BuyerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreLogicAddedIterator is returned from FilterLogicAdded and is used to iterate over the raw logs and unpacked data for LogicAdded events raised by the Store contract.
type StoreLogicAddedIterator struct {
	Event *StoreLogicAdded // Event containing the contract specifics and raw log

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
func (it *StoreLogicAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreLogicAdded)
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
		it.Event = new(StoreLogicAdded)
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
func (it *StoreLogicAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreLogicAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreLogicAdded represents a LogicAdded event raised by the Store contract.
type StoreLogicAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogicAdded is a free log retrieval operation binding the contract event 0xce0ca4df7a7ec71282174625246800d852b6926b9760ef984c9fa6611e71724a.
//
// Solidity: event LogicAdded(address indexed account)
func (_Store *StoreFilterer) FilterLogicAdded(opts *bind.FilterOpts, account []common.Address) (*StoreLogicAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "LogicAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreLogicAddedIterator{contract: _Store.contract, event: "LogicAdded", logs: logs, sub: sub}, nil
}

// WatchLogicAdded is a free log subscription operation binding the contract event 0xce0ca4df7a7ec71282174625246800d852b6926b9760ef984c9fa6611e71724a.
//
// Solidity: event LogicAdded(address indexed account)
func (_Store *StoreFilterer) WatchLogicAdded(opts *bind.WatchOpts, sink chan<- *StoreLogicAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "LogicAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreLogicAdded)
				if err := _Store.contract.UnpackLog(event, "LogicAdded", log); err != nil {
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

// ParseLogicAdded is a log parse operation binding the contract event 0xce0ca4df7a7ec71282174625246800d852b6926b9760ef984c9fa6611e71724a.
//
// Solidity: event LogicAdded(address indexed account)
func (_Store *StoreFilterer) ParseLogicAdded(log types.Log) (*StoreLogicAdded, error) {
	event := new(StoreLogicAdded)
	if err := _Store.contract.UnpackLog(event, "LogicAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreLogicRemovedIterator is returned from FilterLogicRemoved and is used to iterate over the raw logs and unpacked data for LogicRemoved events raised by the Store contract.
type StoreLogicRemovedIterator struct {
	Event *StoreLogicRemoved // Event containing the contract specifics and raw log

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
func (it *StoreLogicRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreLogicRemoved)
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
		it.Event = new(StoreLogicRemoved)
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
func (it *StoreLogicRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreLogicRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreLogicRemoved represents a LogicRemoved event raised by the Store contract.
type StoreLogicRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogicRemoved is a free log retrieval operation binding the contract event 0x9f60fd9ac36634c8a086fb62b8d5f9b088ed08a56805bb49b467012264d1444c.
//
// Solidity: event LogicRemoved(address indexed account)
func (_Store *StoreFilterer) FilterLogicRemoved(opts *bind.FilterOpts, account []common.Address) (*StoreLogicRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "LogicRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreLogicRemovedIterator{contract: _Store.contract, event: "LogicRemoved", logs: logs, sub: sub}, nil
}

// WatchLogicRemoved is a free log subscription operation binding the contract event 0x9f60fd9ac36634c8a086fb62b8d5f9b088ed08a56805bb49b467012264d1444c.
//
// Solidity: event LogicRemoved(address indexed account)
func (_Store *StoreFilterer) WatchLogicRemoved(opts *bind.WatchOpts, sink chan<- *StoreLogicRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "LogicRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreLogicRemoved)
				if err := _Store.contract.UnpackLog(event, "LogicRemoved", log); err != nil {
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

// ParseLogicRemoved is a log parse operation binding the contract event 0x9f60fd9ac36634c8a086fb62b8d5f9b088ed08a56805bb49b467012264d1444c.
//
// Solidity: event LogicRemoved(address indexed account)
func (_Store *StoreFilterer) ParseLogicRemoved(log types.Log) (*StoreLogicRemoved, error) {
	event := new(StoreLogicRemoved)
	if err := _Store.contract.UnpackLog(event, "LogicRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the Store contract.
type StoreManagerAddedIterator struct {
	Event *StoreManagerAdded // Event containing the contract specifics and raw log

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
func (it *StoreManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreManagerAdded)
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
		it.Event = new(StoreManagerAdded)
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
func (it *StoreManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreManagerAdded represents a ManagerAdded event raised by the Store contract.
type StoreManagerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_Store *StoreFilterer) FilterManagerAdded(opts *bind.FilterOpts, account []common.Address) (*StoreManagerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreManagerAddedIterator{contract: _Store.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_Store *StoreFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *StoreManagerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreManagerAdded)
				if err := _Store.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_Store *StoreFilterer) ParseManagerAdded(log types.Log) (*StoreManagerAdded, error) {
	event := new(StoreManagerAdded)
	if err := _Store.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the Store contract.
type StoreManagerRemovedIterator struct {
	Event *StoreManagerRemoved // Event containing the contract specifics and raw log

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
func (it *StoreManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreManagerRemoved)
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
		it.Event = new(StoreManagerRemoved)
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
func (it *StoreManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreManagerRemoved represents a ManagerRemoved event raised by the Store contract.
type StoreManagerRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_Store *StoreFilterer) FilterManagerRemoved(opts *bind.FilterOpts, account []common.Address) (*StoreManagerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreManagerRemovedIterator{contract: _Store.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_Store *StoreFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *StoreManagerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreManagerRemoved)
				if err := _Store.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_Store *StoreFilterer) ParseManagerRemoved(log types.Log) (*StoreManagerRemoved, error) {
	event := new(StoreManagerRemoved)
	if err := _Store.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreSellerAddedIterator is returned from FilterSellerAdded and is used to iterate over the raw logs and unpacked data for SellerAdded events raised by the Store contract.
type StoreSellerAddedIterator struct {
	Event *StoreSellerAdded // Event containing the contract specifics and raw log

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
func (it *StoreSellerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSellerAdded)
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
		it.Event = new(StoreSellerAdded)
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
func (it *StoreSellerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSellerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSellerAdded represents a SellerAdded event raised by the Store contract.
type StoreSellerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSellerAdded is a free log retrieval operation binding the contract event 0x3d78526f2568d430fd8e20719601f7bc3f979936f8880236cc284e8127251b42.
//
// Solidity: event SellerAdded(address indexed account)
func (_Store *StoreFilterer) FilterSellerAdded(opts *bind.FilterOpts, account []common.Address) (*StoreSellerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "SellerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreSellerAddedIterator{contract: _Store.contract, event: "SellerAdded", logs: logs, sub: sub}, nil
}

// WatchSellerAdded is a free log subscription operation binding the contract event 0x3d78526f2568d430fd8e20719601f7bc3f979936f8880236cc284e8127251b42.
//
// Solidity: event SellerAdded(address indexed account)
func (_Store *StoreFilterer) WatchSellerAdded(opts *bind.WatchOpts, sink chan<- *StoreSellerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "SellerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSellerAdded)
				if err := _Store.contract.UnpackLog(event, "SellerAdded", log); err != nil {
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

// ParseSellerAdded is a log parse operation binding the contract event 0x3d78526f2568d430fd8e20719601f7bc3f979936f8880236cc284e8127251b42.
//
// Solidity: event SellerAdded(address indexed account)
func (_Store *StoreFilterer) ParseSellerAdded(log types.Log) (*StoreSellerAdded, error) {
	event := new(StoreSellerAdded)
	if err := _Store.contract.UnpackLog(event, "SellerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreSellerRemovedIterator is returned from FilterSellerRemoved and is used to iterate over the raw logs and unpacked data for SellerRemoved events raised by the Store contract.
type StoreSellerRemovedIterator struct {
	Event *StoreSellerRemoved // Event containing the contract specifics and raw log

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
func (it *StoreSellerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSellerRemoved)
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
		it.Event = new(StoreSellerRemoved)
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
func (it *StoreSellerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSellerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSellerRemoved represents a SellerRemoved event raised by the Store contract.
type StoreSellerRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSellerRemoved is a free log retrieval operation binding the contract event 0xa6f3b3030df964d974376ba4db21cbe49207bd79998f0d619c1c3a422259707f.
//
// Solidity: event SellerRemoved(address indexed account)
func (_Store *StoreFilterer) FilterSellerRemoved(opts *bind.FilterOpts, account []common.Address) (*StoreSellerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "SellerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoreSellerRemovedIterator{contract: _Store.contract, event: "SellerRemoved", logs: logs, sub: sub}, nil
}

// WatchSellerRemoved is a free log subscription operation binding the contract event 0xa6f3b3030df964d974376ba4db21cbe49207bd79998f0d619c1c3a422259707f.
//
// Solidity: event SellerRemoved(address indexed account)
func (_Store *StoreFilterer) WatchSellerRemoved(opts *bind.WatchOpts, sink chan<- *StoreSellerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "SellerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSellerRemoved)
				if err := _Store.contract.UnpackLog(event, "SellerRemoved", log); err != nil {
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

// ParseSellerRemoved is a log parse operation binding the contract event 0xa6f3b3030df964d974376ba4db21cbe49207bd79998f0d619c1c3a422259707f.
//
// Solidity: event SellerRemoved(address indexed account)
func (_Store *StoreFilterer) ParseSellerRemoved(log types.Log) (*StoreSellerRemoved, error) {
	event := new(StoreSellerRemoved)
	if err := _Store.contract.UnpackLog(event, "SellerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
