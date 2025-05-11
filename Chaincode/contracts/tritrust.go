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

	InsuranceEligible bool   `json:"insuranceEligible"`
	LoanEligible      bool   `json:"loanEligible"`
	LastUpdated       string `json:"lastUpdated"`
}

type LoanAccount struct {
	AccountID        string  `json:"accountID"`
	RequestedAmount  float64 `json:"requestedAmount"`
	EmploymentStatus string  `json:"employmentStatus"`
	PayrollDocument  string  `json:"payrollDocument"`
	CIBILScore       int     `json:"cibilScore"`
	RequestStatus    string  `json:"requestStatus"`
	RequestedAt      string  `json:"requestedAt"`
}

type InsuranceApplication struct {
	ApplicationID  string  `json:"applicationID"`
	AccountID      string  `json:"accountID"`
	PolicyType     string  `json:"policyType"`
	CoverageAmount float64 `json:"coverageAmount"`
	PremiumAmount  float64 `json:"premiumAmount"`
	PaymentMethod  string  `json:"paymentMethod"`
	InsuranceOrg   string  `json:"insuranceOrg"`
	Status         string  `json:"status"`
}

type InsuranceAccount struct {
	PolicyNumber   string  `json:"policyNumber"`
	AccountID      string  `json:"accountID"`
	InsuranceOrg   string  `json:"insuranceOrg"`
	PolicyType     string  `json:"policyType"`
	CoverageAmount float64 `json:"coverageAmount"`
	PremiumAmount  float64 `json:"premiumAmount"`
	PaymentMethod  string  `json:"paymentMethod"`
	IssuedAt       string  `json:"issuedAt"`
	ValidTill      string  `json:"validTill"`
}

// BANK ORG:
// **********

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

<<<<<<< HEAD
	if clientOrgId == "bank-fin-com" {
=======
	if clientOrgId == "Org1MSP" {
>>>>>>> 53f2682f8d391422205eedd979c70cd074bbb714

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
<<<<<<< HEAD
=======
	// else {
	// 	return "", fmt.Errorf("only banks can create accounts")
	// }
>>>>>>> 53f2682f8d391422205eedd979c70cd074bbb714
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

<<<<<<< HEAD
// LOAN ORG:
//***********

func (tc *TrustChaincode) RequestLoan(ctx contractapi.TransactionContextInterface, applicationId string, accountID string, requestedAmount float64, employmentStatus string, payrollDocument string, cibilScore int) (string, error) {
=======
func (tc *TrustChaincode) UpdateAccount(ctx contractapi.TransactionContextInterface,
	accountID string, address string, balance string, insuranceEligible bool, policyNumber string, coverageAmount float64, insuranceOrg string) (string, error) {
>>>>>>> 53f2682f8d391422205eedd979c70cd074bbb714

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}
<<<<<<< HEAD

	if clientOrgId == "loanprovider-fin-com" {

		accountJSON, err := ctx.GetStub().GetState(accountID)
		if err != nil {
			return "", fmt.Errorf("failed to read bank account: %v", err)
		}
		if accountJSON == nil {
			return "", fmt.Errorf("bank account with ID %s does not exist", accountID)
=======
	if clientOrgId == "Org2MSP" {

		account, err := tc.ReadAccount(ctx, accountID)
		if err != nil {
			return "", fmt.Errorf("failed to read account %s: %v", accountID, err)
>>>>>>> 53f2682f8d391422205eedd979c70cd074bbb714
		}

		loan := LoanAccount{
			AccountID:        accountID,
			RequestedAmount:  requestedAmount,
			EmploymentStatus: employmentStatus,
			PayrollDocument:  payrollDocument,
			CIBILScore:       cibilScore,
			RequestStatus:    "PENDING",
			RequestedAt:      time.Now().Format(time.RFC3339),
		}

		loanKey := "LOAN_" + applicationId
		loanJSON, err := json.Marshal(loan)
		if err != nil {
			return "", fmt.Errorf("failed to marshal loan request: %v", err)
		}

		err = ctx.GetStub().PutState(loanKey, loanJSON)
		if err != nil {
			return "", fmt.Errorf("failed to put loan request in world state: %v", err)
		}

		return fmt.Sprintf("Loan request %s saved successfully for account %s", applicationId, accountID), nil
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func IsLoanEligible(age int, balance float64, employmentStatus string, cibilScore int, requestedAmount float64) (bool, string) {
	if age < 21 || age > 60 {
		return false, "age must be between 21 and 60"
	}
	if balance < 1000 {
		return false, "account balance must be at least 1000"
	}
	if employmentStatus != "employed" {
		return false, "employment status must be 'employed'"
	}
	if cibilScore < 700 {
		return false, "CIBIL score must be at least 700"
	}
	if requestedAmount > 10*balance {
		return false, "requested amount exceeds allowed limit (10x balance)"
	}
	return true, ""
}

func (tc *TrustChaincode) ProcessLoanApplication(ctx contractapi.TransactionContextInterface,
	applicationId string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "loanprovider-fin-com" {

		loanKey := "LOAN_" + applicationId

		loanJSON, err := ctx.GetStub().GetState(loanKey)
		if err != nil {
			return "", fmt.Errorf("failed to read loan request: %v", err)
		}
		if loanJSON == nil {
			return "", fmt.Errorf("loan request with application ID %s does not exist", applicationId)
		}

		var loan LoanAccount
		if err := json.Unmarshal(loanJSON, &loan); err != nil {
			return "", fmt.Errorf("failed to unmarshal loan request: %v", err)
		}

		accountJSON, err := ctx.GetStub().GetState(loan.AccountID)
		if err != nil {
			return "", fmt.Errorf("failed to read bank account: %v", err)
		}
		if accountJSON == nil {
			return "", fmt.Errorf("bank account with ID %s does not exist", loan.AccountID)
		}

		var account BankAccount
		if err := json.Unmarshal(accountJSON, &account); err != nil {
			return "", fmt.Errorf("failed to unmarshal bank account: %v", err)
		}

		eligible, reason := IsLoanEligible(account.Age, account.Balance, loan.EmploymentStatus, loan.CIBILScore, loan.RequestedAmount)

		loan.RequestedAt = time.Now().Format(time.RFC3339)

		if eligible {
			loan.RequestStatus = "APPROVED"
			account.LoanEligible = true
		} else {
			loan.RequestStatus = "REJECTED"
			account.LoanEligible = false
		}

		updatedLoanJSON, err := json.Marshal(loan)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated loan request: %v", err)
		}
		if err := ctx.GetStub().PutState(loanKey, updatedLoanJSON); err != nil {
			return "", fmt.Errorf("failed to update loan request: %v", err)
		}

		account.LastUpdated = time.Now().Format(time.RFC3339)
		updatedAccountJSON, err := json.Marshal(account)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated account: %v", err)
		}
<<<<<<< HEAD
		if err := ctx.GetStub().PutState(account.AccountID, updatedAccountJSON); err != nil {
			return "", fmt.Errorf("failed to update account: %v", err)
		}

		if eligible {
			return fmt.Sprintf("Loan request %s has been approved.", applicationId), nil
		} else {
			return "", fmt.Errorf("loan request %s has been rejected: %s", applicationId, reason)
		}
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func (tc *TrustChaincode) DisburseLoan(ctx contractapi.TransactionContextInterface, applicationId string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "loanprovider-fin-com" {
		loanKey := "LOAN_" + applicationId
		loanJSON, err := ctx.GetStub().GetState(loanKey)
		if err != nil {
			return "", fmt.Errorf("failed to read loan request: %v", err)
		}
		if loanJSON == nil {
			return "", fmt.Errorf("loan request with application ID %s does not exist", applicationId)
		}

		var loan LoanAccount
		err = json.Unmarshal(loanJSON, &loan)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal loan request: %v", err)
		}

		if loan.RequestStatus != "APPROVED" {
			return "", fmt.Errorf("loan request %s is not approved yet, cannot disburse", applicationId)
		}

		accountJSON, err := ctx.GetStub().GetState(loan.AccountID)
		if err != nil {
			return "", fmt.Errorf("failed to read bank account: %v", err)
		}
		if accountJSON == nil {
			return "", fmt.Errorf("bank account with ID %s does not exist", loan.AccountID)
		}

		var account BankAccount
		err = json.Unmarshal(accountJSON, &account)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal bank account: %v", err)
		}

		account.Balance += loan.RequestedAmount
		account.LastUpdated = time.Now().Format(time.RFC3339)

		loan.RequestStatus = "DISBURSED"
		loan.RequestedAt = time.Now().Format(time.RFC3339)

		updatedLoanJSON, _ := json.Marshal(loan)
		err = ctx.GetStub().PutState(loanKey, updatedLoanJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update loan status to DISBURSED: %v", err)
		}

		updatedAccountJSON, _ := json.Marshal(account)
		err = ctx.GetStub().PutState(loan.AccountID, updatedAccountJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update bank account balance: %v", err)
		}

		return fmt.Sprintf("Loan %s has been successfully disbursed to account %s", applicationId, loan.AccountID), nil
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func (tc *TrustChaincode) ReadLoanAccount(ctx contractapi.TransactionContextInterface, applicationId string) (*LoanAccount, error) {

	bytes, err := ctx.GetStub().GetState(applicationId)
	if err != nil {
		return nil, fmt.Errorf("failed to read account %s: %v", applicationId, err)
	} else if bytes == nil {
		return nil, fmt.Errorf("account %s does not exist", applicationId)
	}
	var loan LoanAccount
	err = json.Unmarshal(bytes, &loan)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %v", err)
	}
	return &loan, nil
}

// INSURANCE ORG:
//****************

func (tc *TrustChaincode) ApplyForInsurancePolicy(ctx contractapi.TransactionContextInterface, accountID string, policyType string, coverageAmount float64, premiumAmount float64, paymentMethod string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "insurance-fin-com" {

		accountJSON, err := ctx.GetStub().GetState(accountID)
		if err != nil {
			return "", fmt.Errorf("failed to read bank account: %v", err)
		}
		if accountJSON == nil {
			return "", fmt.Errorf("account with ID %s does not exist", accountID)
		}

		applicationKey := "INSURANCE_" + accountID + "_" + policyType
		applicationJSON, err := ctx.GetStub().GetState(applicationKey)
		if err == nil && applicationJSON != nil {
			return "", fmt.Errorf("an application for this policy type already exists for account %s", accountID)
		}

		insuranceApplication := InsuranceApplication{
			AccountID:      accountID,
			PolicyType:     policyType,
			CoverageAmount: coverageAmount,
			PremiumAmount:  premiumAmount,
			PaymentMethod:  paymentMethod,
			Status:         "Pending",
		}

		applicationJSON, err = json.Marshal(insuranceApplication)
		if err != nil {
			return "", fmt.Errorf("failed to marshal insurance application: %v", err)
		}

		err = ctx.GetStub().PutState(applicationKey, applicationJSON)
		if err != nil {
			return "", fmt.Errorf("failed to put insurance application in world state: %v", err)
		}

		return "Insurance application submitted successfully. Awaiting verification.", nil
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func CheckInsuranceEligibility(account BankAccount, policyType string) bool {

	if policyType == "Pension" {
		if account.Age < 18 || account.Age > 60 {
			return false
		}
		if account.Balance < 10000 {
			return false
		}
	} else if policyType == "TermLife" {

		if account.Age < 18 || account.Age > 70 {
			return false
		}
	} else if policyType == "Health" {

		if account.Balance < 5000 {
			return false
		}
	}
	return true
}

func (tc *TrustChaincode) VerifyInsuranceApplication(ctx contractapi.TransactionContextInterface, applicationID string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "insurance-fin-com" {

		applicationJSON, err := ctx.GetStub().GetState(applicationID)
		if err != nil {
			return "", fmt.Errorf("failed to read insurance application: %v", err)
		}
		if applicationJSON == nil {
			return "", fmt.Errorf("insurance application with ID %s does not exist", applicationID)
		}

		var application InsuranceApplication
		err = json.Unmarshal(applicationJSON, &application)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal insurance application: %v", err)
		}

		accountJSON, err := ctx.GetStub().GetState(application.AccountID)
		if err != nil {
			return "", fmt.Errorf("failed to read bank account: %v", err)
		}
		if accountJSON == nil {
			return "", fmt.Errorf("account with ID %s does not exist", application.AccountID)
		}

		var account BankAccount
		err = json.Unmarshal(accountJSON, &account)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal bank account: %v", err)
		}

		isEligible := CheckInsuranceEligibility(account, application.PolicyType)
		if !isEligible {
			application.Status = "Rejected"
		} else {
			application.Status = "Approved"
			account.InsuranceEligible = true
		}

		updatedApplicationJSON, err := json.Marshal(application)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated insurance application: %v", err)
		}

		err = ctx.GetStub().PutState(applicationID, updatedApplicationJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update insurance application in world state: %v", err)
		}

		updateAccount, err := json.Marshal(account)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated bank account: %v", err)
		}

		err = ctx.GetStub().PutState(account.AccountID, updateAccount)
		if err != nil {
			return "", fmt.Errorf("failed to update bank account in world state: %v", err)
		}

		return fmt.Sprintf("Insurance application %s has been %s.", applicationID, application.Status), nil
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func (tc *TrustChaincode) IssueInsurancePolicy(ctx contractapi.TransactionContextInterface, applicationID string) (string, error) {

	clientOrgId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgId == "insurance-fin-com" {

		applicationJSON, err := ctx.GetStub().GetState(applicationID)
		if err != nil {
			return "", fmt.Errorf("failed to read insurance application: %v", err)
		}
		if applicationJSON == nil {
			return "", fmt.Errorf("insurance application with ID %s does not exist", applicationID)
		}

		var application InsuranceApplication
		err = json.Unmarshal(applicationJSON, &application)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal insurance application: %v", err)
		}

		if application.Status != "Approved" {
			return "", fmt.Errorf("cannot issue policy as the application is not approved")
		}

		policyNumber := generatePolicyNumber(application.AccountID, application.PolicyType)

		validTill := time.Now().AddDate(1, 0, 0).Format("2006-01-02")

		insuranceAccount := InsuranceAccount{
			PolicyNumber:   policyNumber,
			AccountID:      application.AccountID,
			InsuranceOrg:   "InsuranceOrg",
			PolicyType:     application.PolicyType,
			CoverageAmount: application.CoverageAmount,
			PremiumAmount:  application.PremiumAmount,
			PaymentMethod:  application.PaymentMethod,
			IssuedAt:       time.Now().Format("2006-01-02"),
			ValidTill:      validTill,
		}

		insuranceAccountJSON, err := json.Marshal(insuranceAccount)
		if err != nil {
			return "", fmt.Errorf("failed to marshal insurance account: %v", err)
		}

		err = ctx.GetStub().PutState(policyNumber, insuranceAccountJSON)
		if err != nil {
			return "", fmt.Errorf("failed to put insurance account in world state: %v", err)
		}

		application.Status = "Issued"
		applicationJSON, err = json.Marshal(application)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated insurance application: %v", err)
		}

		err = ctx.GetStub().PutState(applicationID, applicationJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update insurance application in world state: %v", err)
		}

		return fmt.Sprintf("Policy issued successfully. Policy Number: %s", policyNumber), nil
	}
	return "", fmt.Errorf("operation not permitted for organization: %s", clientOrgId)
}

func generatePolicyNumber(accountID string, policyType string) string {

	return fmt.Sprintf("%s-%s-%s", accountID, policyType, time.Now().Format("20060102-150405"))
}

func (tc *TrustChaincode) ReadInsuranceAccount(ctx contractapi.TransactionContextInterface, policyNumber string) (*InsuranceAccount, error) {

	accountJSON, err := ctx.GetStub().GetState(policyNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to read insurance account from world state: %v", err)
	}
	if accountJSON == nil {
		return nil, fmt.Errorf("insurance account with policy number %s does not exist", policyNumber)
	}

	var insuranceAccount InsuranceAccount
	err = json.Unmarshal(accountJSON, &insuranceAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal insurance account JSON: %v", err)
	}

	return &insuranceAccount, nil
}

func (tc *TrustChaincode) ReadInsuranceApplication(ctx contractapi.TransactionContextInterface, applicationID string) (*InsuranceApplication, error) {
	applicationJSON, err := ctx.GetStub().GetState(applicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to read insurance application from world state: %v", err)
	}
	if applicationJSON == nil {
		return nil, fmt.Errorf("insurance application with ID %s does not exist", applicationID)
	}

	var application InsuranceApplication
	err = json.Unmarshal(applicationJSON, &application)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal insurance application JSON: %v", err)
	}

	return &application, nil
=======

		err = ctx.GetStub().PutState(accountID, accountJSON)
		if err != nil {
			return "", fmt.Errorf("failed to update account in ledger: %v", err)
		}

		return fmt.Sprintf("insurance details updated for account %s", accountID), nil
	}
	// else {
	// 	return "", fmt.Errorf("only insurance organizations can update accounts")
	// }
	return fmt.Sprintf("Account updation attempted by %s: %v", clientOrgId, accountID), nil
>>>>>>> 53f2682f8d391422205eedd979c70cd074bbb714
}
