package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type AccountService struct {
	sling *sling.Sling
}

func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling,
	}
}

type Accounts struct {
	Items []Account `json:"Items"`
	PagedResults
}

type Account struct {
	// Common Account fields

	ID                              string                 `json:"Id"`
	AccountType                     AccountType            `json:"AccountType"`
	Description                     string                 `json:"Description,omitempty"`
	EnvironmentIDs                  []string               `json:"EnvironmentIds,omitempty"`
	Name                            string                 `json:"Name" validate:"required"`
	TenantedDeploymentParticipation TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantTags                      []string               `json:"TenantTags,omitempty"`
	Token                           SensitiveValue         `json:"Token,omitempty"`

	// Azure Service Principal fields

	AzureEnvironment                  string         `json:"AzureEnvironment,omitempty"`
	ActiveDirectoryEndpointBaseURI    string         `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ClientID                          string         `json:"ClientId,omitempty"`
	Password                          SensitiveValue `json:"Password,omitempty"`
	ResourceManagementEndpointBaseURI string         `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SubscriptionNumber                string         `json:"SubscriptionNumber,omitempty"`
	TenantID                          string         `json:"TenantId,omitempty"`
}

func (t *Account) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf(`%s failed validation. Validation type: %s Field type: %s`, err.Namespace(), err.Tag(), err.Type())
			fmt.Println()
		}
		return err
	}

	switch t.AccountType {
	case AzureServicePrincipal:
		return validateAzureServicePrincipalAccount(t)
	default:
		return nil
	}
}

func validateAzureServicePrincipalAccount(acc *Account) error {
	validations := []error{
		ValidateRequiredPropertyValue("ClientID", acc.ClientID),
		ValidateRequiredPropertyValue("TenantID", acc.TenantID),
		ValidateRequiredPropertyValue("SubscriptionNumber", acc.SubscriptionNumber),
	}

	if acc.Password.HasValue {
		validations = append(validations, ValidateRequiredPropertyValue("Password", acc.Password.NewValue))
	}

	return ValidateMultipleProperties(validations)
}

func NewAccount(name string, accountType AccountType) *Account {
	return &Account{
		Name:        name,
		AccountType: accountType,
	}
}

func (s *AccountService) Get(accountId string) (*Account, error) {
	path := fmt.Sprintf("accounts/%s", accountId)
	resp, err := apiGet(s.sling, new(Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}

func (s *AccountService) GetAll() (*[]Account, error) {
	var p []Account

	path := "accounts"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Accounts), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Accounts)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *AccountService) GetByName(accountName string) (*Account, error) {
	var foundAccount Account
	accounts, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, account := range *accounts {
		if account.Name == accountName {
			return &account, nil
		}
	}

	return &foundAccount, fmt.Errorf("no account found with account name %s", accountName)
}

func (s *AccountService) Add(account *Account) (*Account, error) {
	resp, err := apiAdd(s.sling, account, new(Account), "accounts")

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}

func (s *AccountService) Delete(accountId string) error {
	path := fmt.Sprintf("accounts/%s", accountId)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Update(account *Account) (*Account, error) {
	path := fmt.Sprintf("accounts/%s", account.ID)
	resp, err := apiUpdate(s.sling, account, new(Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}
