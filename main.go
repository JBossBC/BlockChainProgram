package main

import (
	"fmt"
	"github.com/metarticle/heroes-service/blockchain"
	//"github.com/metarticle/heroes-service/web"
	//"github.com/metarticle/heroes-service/web/controllers"
	//"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.chainhero.io",

		// Channel parameters
		ChannelID:     "chainhero",
		//ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",
		ChannelConfig: "/home/sissice/go/src/github.com/metarticle/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		//ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodeGoPath: "/home/sissice/go/",
		ChaincodePath:   "github.com/metarticle/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
	
	// Query the chaincode
	response, err := fSetup.CreateArticle()
	if err != nil {
		fmt.Printf("Unable to create article on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the create article: %s\n", response)
	}
	
	res, err := fSetup.QueryArticle()
	if err != nil {
		fmt.Printf("Unable to query article on the chaincode: %v\n", err)
	} else {
		fmt.Println("Response from the query article: ", res)
	}

	// Launch the web application listening
	/*
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
	*/
}
