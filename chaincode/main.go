package main

import (
	"github.com/metarticle/heroes-service/chaincode/vender/chaincode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting HelloBsiService: %s", err)
	}
}
