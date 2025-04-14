package main

import (
	"log"
	"tritrustmod/contracts"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	TrustChaincode := new(contracts.TrustChaincode)
	chaincode, err := contractapi.NewChaincode(TrustChaincode)
	if err != nil {	
		log.Panicf("Error creating TrustChaincode chaincode: %v", err)
	}
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting TrustChaincode chaincode: %v", err)
	}
}