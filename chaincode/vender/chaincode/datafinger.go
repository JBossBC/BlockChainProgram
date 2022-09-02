package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//新增文章数据指纹
func AddDataFinger(stub shim.ChaincodeStubInterface, dataFinger string, key string) pb.Response {
	fmt.Println("新增文章数据指纹", dataFinger, key)
	res := QueryDataFinger(stub, dataFinger)
	if res.Status == shim.OK {
		return shim.Error(fmt.Sprintf("数据指纹已存在:%s",dataFinger))
	}
	newDataFinger := DataFinger{FileDataFinger: dataFinger, Key: key}
	jsonData, err := json.Marshal(newDataFinger)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-序列化json数据失败出错: %s", jsonData, err))

	}
	err = stub.PutState(dataFinger, jsonData)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s数据指纹上链失败:%s",err,dataFinger))
	}
	return shim.Success(nil)
}

//查询文章数据指纹
func QueryDataFinger(stub shim.ChaincodeStubInterface, dataFinger string) pb.Response {
	fmt.Println("查询文章数据指纹", dataFinger)
	res, err := stub.GetState(dataFinger)
	if err != nil || res == nil{
		return shim.Error(fmt.Sprintf("没有获取到对应的数据指纹:%s",err))
	}
	fmt.Println("查询出的数据指纹对象为",res)
	return shim.Success(nil)
}
