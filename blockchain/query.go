package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"encoding/json"
)

type Article struct {
	Key string //文章存证号
	Name string //文章名称
	DataFinger string //文章数据指纹
	Author string //作者
	Submitter string //提交者
	CompletionTime string //文章完成时间
	ReadPrice string //阅读价格
	DownloadPrice string //下载价格
	CheckStatus string //审核状态
	Status string //链上文章有效状态
}

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) CreateArticle() (string, error) {

	article := Article{Key: "1",Name: "a",DataFinger: "1",Author: "a",Submitter: "b",CompletionTime: "1",ReadPrice: "0.1",DownloadPrice: "0.5",CheckStatus: "1",Status: "1"}
	marshal,_ := json.Marshal(article)

	// Prepare arguments
	var args []string
	args = append(args, "createArticle")
	//args = append(args, marshal)

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "createArticle", Args: [][]byte{marshal}})
	if err != nil {
		return "", fmt.Errorf("failed to create: %v", err)
	}
	fmt.Printf("Invoke chaincode createArticle response:\n"+
		"id: %v\nvalidate: %v\nchaincode status: %v\n\n",
		response.TransactionID,
		response.TxValidationCode,
		response.ChaincodeStatus)

	return string(response.TransactionID), nil
}

func (setup *FabricSetup) QueryArticle() ([]byte, error) {


	// Prepare arguments
	
	var args []string
	args = append(args, "queryArticle")
	//args = append(args, marshal)

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte("1")}})
	
	/*
	req := channel.Request{
		ChaincodeID: setup.ChainCodeID,
		Fcn:         "queryArticle",
		Args:        packArgs([]string{"1"}),
	}

	// send request and handle response
	reqPeers := channel.WithTargetEndpoints("peer0.org1.hf.chainhero.io")
	response, err := setup.client.Query(req, reqPeers)
	*/
	if err != nil {
		return []byte{}, fmt.Errorf("failed to query: %v", err)
	}
	return response.Payload, nil
}

func packArgs(paras []string) [][]byte {
	var args [][]byte
	for _, k := range paras {
		args = append(args, []byte(k))
	}
	return args
}
