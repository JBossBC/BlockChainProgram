package chaincode

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {

}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("SimpleChaincode Init")
	//args := stub.GetStringArgs()

	args := stub.GetStringArgs()
	if len(args) != 1 && args[0] == "" {
		return shim.Error("Incorrect arguments")
	}

	// Initialize the HelloBsiService
	operator := args[0]
	fmt.Printf("operator = %v\n", operator)
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("SimpleChaincode Invoke")
	//提取调用链码交易中的参数，其中第一个作为被调用的函数名称，剩下的参数作为函数的执行参数
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "createArticle":
		return t.CreateArticle(stub, args)
	case "queryArticle":
		return t.QueryArticle(stub, args)
	}

	return shim.Error("请输入正确的函数名称")
}
