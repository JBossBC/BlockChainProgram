package main

import (
	"encoding/json"
	"github.com/metarticle/heroes-service/chaincode/vender/chaincode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)



func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func Test(t *testing.T) {
	// SimpleChaincode为链码逻辑中实现的实际struct
	cc := new(chaincode.SimpleChaincode)
	// 获取MockStub对象， 传入名称和链码实体
	stub := shim.NewMockStub("SimpleChaincode", cc)

	// 初始化链码
	checkInit(t,stub,[][]byte{[]byte("init")})

	article := chaincode.Article{Key: "1",Name: "a",DataFinger: "1",Author: "a",Submitter: "b",CompletionTime: "1",ReadPrice: "0.1",DownloadPrice: "0.5",CheckStatus: "1",Status: "1"}
	marshal,_ := json.Marshal(article)
	fmt.Println("--------------添加文章---------------")
	checkInvoke(t, stub, [][]byte{[]byte("createArticle"), marshal})
	
	article = Article{Key: "1",Name: "a",DataFinger: "1",Author: "a",Submitter: "a",CompletionTime: "1",ReadPrice: "0.1",DownloadPrice: "0.5",CheckStatus: "1",Status: "1"}
	marshal,_ = json.Marshal(article)
	fmt.Println("--------------更新文章1---------------")
	fmt.Println(marshal)
	checkInvoke(t, stub, [][]byte{[]byte("updateArticle"), marshal})
	checkInvoke(t, stub, [][]byte{[]byte("queryArticle"), []byte("1")})

	article = Article{Key: "1",Name: "",DataFinger: "",Author: "",Submitter: "",CompletionTime: "2",ReadPrice: "",DownloadPrice: "",CheckStatus: "",Status: ""}
	marshal,_ = json.Marshal(article)
	fmt.Println("--------------更新文章2---------------")
	fmt.Println(marshal)
	checkInvoke(t, stub, [][]byte{[]byte("updateArticle"), marshal})
	checkInvoke(t, stub, [][]byte{[]byte("queryArticle"), []byte("1")})
}
