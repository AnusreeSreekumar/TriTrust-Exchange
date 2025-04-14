package contracts

import (
	"encoding/json"
	"fmt"
	"time"

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

	InsuranceEligible bool    `json:"insuranceEligible"`
	PolicyNumber      string  `json:"policyNumber"`
	CoverageAmount    float64 `json:"coverageAmount"`
	InsuranceOrg      string  `json:"insuranceOrg"`

	LastUpdated string `json:"lastUpdated"`
}

func (tc *TrustChaincode) AccountExists(ctx contractapi.TransactionContextInterface, accountID string) (bool, error) {
	accountData, err := ctx.GetStub().GetState(accountID)
	if err != nil {
		return false, err
	}
	return accountData != nil, nil
}

func (tc *TrustChaincode) CreateAccount(ctx contractapi.TransactionContextInterface,
	accountID string, customerName string, age int, address string, balance float64, accountType string, createdByBank string, insuranceEligible bool) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "bank-org-com" {

		exists, err := tc.AccountExists(ctx, accountID)
		if err != nil {
			return "", fmt.Errorf("failed to read from world state: %v", err)
		} else if exists {
			return fmt.Sprintf("account %s already exists", accountID), nil
		}

		account := BankAccount{
			AccountID:         accountID,
			CustomerName:      customerName,
			Age:               age,
			Address:           address,
			Balance:           balance,
			AccountType:       accountType,
			CreatedByBank:     createdByBank,
			InsuranceEligible: false,
			LastUpdated:       time.Now().Format(time.RFC3339),
		}

		bytes, _ := json.Marshal(account)

		err = ctx.GetStub().PutState(accountID, bytes)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("successfully added account %v", accountID), nil
		}
	} else {
		return "", fmt.Errorf("only banks can create accounts")
	}
}

func (tc *TrustChaincode) ReadAccount(ctx contractapi.TransactionContextInterface, accountID string) (*BankAccount, error) {

	bytes, err := ctx.GetStub().GetState(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to read account %s: %v", accountID, err)
	} else if bytes == nil {
		return nil, fmt.Errorf("account %s does not exist", accountID)
	}
	var account BankAccount
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %v", err)
	}
	return &account, nil

}

func (tc *TrustChaincode) UpdateAccount(ctx contractapi.TransactionContextInterface,
	accountId string, address string, balance string, insuranceEligible bool, policyNumber string, coverageAmount float64, insuranceOrg string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err	
	}
	if clientOrgId == "insurance-org-com" {	

		account, err := tc.ReadAccount(ctx, accountId)
		if err != nil {		
			return "", fmt.Errorf("failed to read account %s: %v", accountId, err)
		}
		account.InsuranceEligible = true
		account.PolicyNumber = policyNumber
		account.CoverageAmount = coverageAmount
		account.InsuranceOrg = insuranceOrg
		account.LastUpdated = time.Now().Format(time.RFC3339)

		accountJSON, err := json.Marshal(account)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated account: %v", err)
		}

		err = ctx.GetStub().PutState(accountId, accountJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update account in ledger: %v", err)
		}

		return fmt.Sprintf("insurance details updated for account %s", accountId), nil
	} else {
		return "", fmt.Errorf("only insurance organizations can update accounts")
	}
}
