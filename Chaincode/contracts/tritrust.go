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
    AccountID    string `json:"accountID"`
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

    LoanEligible   bool    `json:"loanEligible"`
    LoanAmount     float64 `json:"loanAmount"`
    InterestRate   float64 `json:"interestRate"`
    RepaymentStatus string `json:"repaymentStatus"`

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

    if clientOrgId == "BankOrgMSP" {

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
            LoanEligible:      false,
            LastUpdated:       time.Now().Format(time.RFC3339),
        }

        bytes, _ := json.Marshal(account)

        err = ctx.GetStub().PutState(accountID, bytes)
        if err != nil {
            return "", err
        } else {
            return fmt.Sprintf("successfully added account %v", accountID), nil
        }
    }
    return fmt.Sprintf("Account creation attempted by %s: %v", clientOrgId, accountID), nil
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

func (tc *TrustChaincode) UpdateInsuranceDetails(ctx contractapi.TransactionContextInterface,
    accountID string, policyNumber string, coverageAmount float64, insuranceOrg string) (string, error) {

    clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return "", err
    }
    if clientOrgId == "InsuranceOrgMSP" {

        account, err := tc.ReadAccount(ctx, accountID)
        if err != nil {
            return "", fmt.Errorf("failed to read account %s: %v", accountID, err)
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

        err = ctx.GetStub().PutState(accountID, accountJSON)
        if err != nil {
            return "", fmt.Errorf("failed to update account in ledger: %v", err)
        }

        return fmt.Sprintf("insurance details updated for account %s", accountID), nil
    }
    return fmt.Sprintf("Insurance update attempted by %s: %v", clientOrgId, accountID), nil
}

func (tc *TrustChaincode) ApproveLoan(ctx contractapi.TransactionContextInterface,
    accountID string, loanAmount float64, interestRate float64) (string, error) {

    clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return "", err
    }
    if clientOrgId == "LoanProviderOrgMSP" {

        account, err := tc.ReadAccount(ctx, accountID)
        if err != nil {
            return "", fmt.Errorf("failed to read account %s: %v", accountID, err)
        }
        if account.Balance < loanAmount/2 {
            return "", fmt.Errorf("account %s is not eligible for the requested loan amount", accountID)
        }

        account.LoanEligible = true
        account.LoanAmount = loanAmount
        account.InterestRate = interestRate
        account.RepaymentStatus = "Pending"
        account.LastUpdated = time.Now().Format(time.RFC3339)

        accountJSON, err := json.Marshal(account)
        if err != nil {
            return "", fmt.Errorf("failed to marshal updated account: %v", err)
        }

        err = ctx.GetStub().PutState(accountID, accountJSON)
        if err != nil {
            return "", fmt.Errorf("failed to update account in ledger: %v", err)
        }

        return fmt.Sprintf("loan approved for account %s", accountID), nil
    }
    return fmt.Sprintf("Loan approval attempted by %s: %v", clientOrgId, accountID), nil
}

func (tc *TrustChaincode) UpdateLoanDetails(ctx contractapi.TransactionContextInterface,
    accountID string, repaymentStatus string) (string, error) {

    clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return "", err
    }
    if clientOrgId == "LoanProviderOrgMSP" {

        account, err := tc.ReadAccount(ctx, accountID)
        if err != nil {
            return "", fmt.Errorf("failed to read account %s: %v", accountID, err)
        }
        account.RepaymentStatus = repaymentStatus
        account.LastUpdated = time.Now().Format(time.RFC3339)

        accountJSON, err := json.Marshal(account)
        if err != nil {
            return "", fmt.Errorf("failed to marshal updated account: %v", err)
        }

        err = ctx.GetStub().PutState(accountID, accountJSON)
        if err != nil {
            return "", fmt.Errorf("failed to update account in ledger: %v", err)
        }

        return fmt.Sprintf("loan details updated for account %s", accountID), nil
    }
    return fmt.Sprintf("Loan update attempted by %s: %v", clientOrgId, accountID), nil
}