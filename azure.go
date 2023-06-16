package cloudcraft

import (
	"fmt"
)

// AccountIntegrationAzure struct represents the data of an Account.
type AccountIntegrationAzure struct {
    ID              *string `json:"id,omitempty"`
    Name            *string `json:"name,omitempty"`
    ApplicationId   *string `json:"applicationId,omitempty"`
    DirectoryId     *string `json:"directoryId,omitempty"`
    SubscriptionId  *string `json:"subscriptionId,omitempty"`
    ClientSecret    *string `json:"clientSecret,omitempty"`
    ExternalID      *string `json:"externalId,omitempty"`
    CreatedAt       *string `json:"createdAt,omitempty"`
    UpdatedAt       *string `json:"updatedAt,omitempty"`
    CreatorID       *string `json:"CreatorId,omitempty"`
}


// AccountsAzureInfoResp contains data of all Account.
type AccountsAzureInfoResp struct {
	AccountsAzureInfo []AccountIntegrationAzure `json:"accounts,omitempty"`
}

// AccountIntegrationsAllAzure get and return the Azure Accounts.
func (client *Client) AccountIntegrationsAllAzure() ([]AccountIntegrationAzure, error) {
	var out AccountsAzureInfoResp
	err := client.RequestResponse("GET", "/azure/account", nil, &out)
	{
		if err != nil {
			return nil, err
		}
	}
	if len(out.AccountsAzureInfo) == 0 {
		return nil, fmt.Errorf("Cloudcraft Azure Account Integrations not found or there are none")
	}
	return out.AccountsAzureInfo, nil
}

// AccountIntegrationAzure get and return the Azure Account Information.
func (client *Client) AccountIntegrationAzure(AzureAccountId string) (AccountIntegrationAzure, error) {
	accounts, err := client.AccountIntegrationsAllAzure()
	{
		if err != nil {
			return AccountIntegrationAzure{}, err
		}
	}
	//log.Printf("accounts%v - length%d", accounts, len(accounts))
	AccountIntegrationAzureItem := AccountIntegrationAzure{}
	for i := range accounts {

		if *accounts[i].ID == AzureAccountId {
			AccountIntegrationAzureItem = accounts[i]
			//ID := *accounts[i].ID
			//Name := *accounts[i].Name
			//log.Printf("ID %s, Name %s", ID, Name)
			return AccountIntegrationAzureItem, nil
		}
	}
	if (AccountIntegrationAzure{}) == AccountIntegrationAzureItem {
		return AccountIntegrationAzureItem, fmt.Errorf("Cloudcraft Azure Account Integration not found")
	}
	return AccountIntegrationAzureItem, nil
}

// AccountIntegrationAzureCreate updates Azure Account Integration.
func (client *Client) AccountIntegrationAzureCreate(acc *AccountIntegrationAzure) error {

	return client.RequestResponse("POST", "/azure/account", acc, &acc)
}

// AccountIntegrationAzureUpdate updates Azure Account Integrations.
func (client *Client) AccountIntegrationAzureUpdate(acc *AccountIntegrationAzure) error {
	return client.RequestResponse("PUT", fmt.Sprintf("/azure/account/%v", *acc.ID), acc, &acc)
}

// AccountAzureDelAccountIntegrationAzureDeleteete updates Azure Account Integrations.
func (client *Client) AccountIntegrationAzureDelete(acc *AccountIntegrationAzure) error {
	return client.RequestResponse("DELETE", fmt.Sprintf("/azure/account/%v", *acc.ID), nil, nil)
}
