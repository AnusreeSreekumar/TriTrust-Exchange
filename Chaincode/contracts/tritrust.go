package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TrustChaincode struct {
	contractapi.Contract
}

type BankAccount struct {
	AccountID    string `json:"accountId"`
	CustomerName string `json:"customerName"`
	Age          int    `json:"age"`
	Address      string `json:"address"`

	Balance       float64 `json:"balance"`
	AccountType   string  `json:"accountType"`
	CreatedByBank string  `json:"createdByBank"`

	CreditScore  int     `json:"creditScore"`
	LoanEligible bool    `json:"loanEligible"`
	LoanAmount   float64 `json:"loanAmount"`
	CreditAgency string  `json:"creditAgency"`

	InsuranceEligible bool    `json:"insuranceEligible"`
	PolicyNumber      string  `json:"policyNumber"`
	CoverageAmount    float64 `json:"coverageAmount"`
	InsuranceOrg      string  `json:"insuranceOrg"`

	LastUpdated string `json:"lastUpdated"`
}
