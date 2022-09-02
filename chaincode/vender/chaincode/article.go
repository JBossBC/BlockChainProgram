package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *SimpleChaincode) CreateArticle(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	//获取java sdk请求体
	fmt.Println("文章上链入参：{}", args[0])
	var jsonString = args[0]
	var article Article
	var err error
	//将请求体用机构对象进行接收
	err = json.Unmarshal([]byte(jsonString), &article)
	//如果无法转化，则抛出异常
	if err != nil {
		return shim.Error(fmt.Sprintf("执行失败,合约内部错误.参数验证失败:%s",err))
	}
	//验证文章是否存在
	res := QueryArticleInfo(stub, article.Key)
	if res.Status == shim.OK {
		return shim.Error(fmt.Sprintf("执行失败,key已存在:%s",err))
	}

	//验证数据指纹是否存在
	if article.DataFinger == "" {
		return shim.Error(fmt.Sprintf("数据指纹为空"))
	} else {
		res = AddDataFinger(stub, article.DataFinger, article.Key)
		if res.Status != shim.OK {
			return res
		}
	}

	//上链
	jsonData, err := json.Marshal(article)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-序列化json数据失败出错: %s", jsonData, err))

	}
	err = stub.PutState(article.Key, jsonData)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s文章上链失败:%s",err,article))
	}
	fmt.Println("添加文章成功,key=", article.Key)
	return shim.Success(nil)

}

func QueryArticleInfo(stub shim.ChaincodeStubInterface, key string) pb.Response {
	res, err := stub.GetState(key)
	if err != nil || res == nil {
		return shim.Error(fmt.Sprintf("没有获取到对应的文章信息:%s", err))
	}
	fmt.Println("查询出的文章信息为",res)
	return shim.Success(nil)
}


func (t *SimpleChaincode) QueryArticle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取java sdk请求体
	fmt.Println("文章查询入参：{}", args[0])
	var key = args[0]
	res, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("没有获取到对应的文章信息:%s",err))
	}
	/*
	var article Article
	err = json.Unmarshal([]byte(res), &article)
	//如果无法转化，则抛出异常
	if err != nil {
		return shim.Error(fmt.Sprintf("执行失败,合约内部错误.参数验证失败:%s",err))
	}
	*/
	fmt.Println("查询出的文章信息为",res)
	return shim.Success(res)
}


