package main

import (
	"log"
	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
)

func main() {

	client := cloudcraft.NewClient("", "https://api.cloudcraft.co", 3)
	cfurl := client.GetBaseURL()

	//Create a CloudCraft Azure integration
	cloudcraftAzureIntegrationname  := "AzureIntegration"
	applicationId 					:= ""
	directoryId 					:= ""
	subscriptionId 					:= ""
	clientSecret 					:= ""

	newAccInfo := cloudcraft.AccountIntegrationAzure{
		Name: &cloudcraftAzureIntegrationname, 
		ApplicationId: &applicationId,
		DirectoryId: &directoryId,
		SubscriptionId: &subscriptionId,
		ClientSecret: &clientSecret,
	}

	log.Printf("CloudCraft Base URL %s\n", cfurl)
	errNewAcc := client.AccountIntegrationAzureCreate(&newAccInfo)
	{
		if errNewAcc != nil {
			log.Fatal(errNewAcc)
		}
	}

	log.Printf("Integration ID %s", *newAccInfo.ID)
	log.Printf("Integration Name%s", *newAccInfo.Name)

	//Get a CloudCraft Azure integration
	Info, err := client.AccountIntegrationAzure(*newAccInfo.ID)
	{
		if err != nil {
			log.Fatal(err)
		}
	}
	ID := *Info.ID
	Name := *Info.Name
	log.Printf("ID - %s, Name - %s", ID, Name)

	//Get all CloudCraft Azure integrations
	accounts, err := client.AccountIntegrationsAllAzure()
	{
		if err != nil {
			log.Fatal(err)
		}
	}
	//log.Printf("accounts%v - length%d", accounts, len(accounts))

	for i := range accounts {
		if *accounts[i].ID == *newAccInfo.ID {
			ID := *accounts[i].ID
			Name := *accounts[i].Name
			log.Printf("ID - %s, Name - %s", ID, Name)
		}
	}

	//Update a CloudCraft Azure integration
	updID := *newAccInfo.ID
	updname := "AzureIntegrationUpdate"

	updAccInfo := cloudcraft.AccountIntegrationAzure{
		ID: &updID, 
		Name: &updname, 
		ApplicationId: &applicationId,
		DirectoryId: &directoryId,
		SubscriptionId: &subscriptionId,
		ClientSecret: &clientSecret,
	}
	errUpdAcc := client.AccountIntegrationAzureUpdate(&updAccInfo)
	{
		if errUpdAcc != nil {
			log.Fatal(errUpdAcc)
		}
	}

	log.Printf("Integration ID %s", *updAccInfo.ID)
	log.Printf("Integration Name %s", *updAccInfo.Name)

	//Delete a CloudCraft Azure integration
	errDelAcc := client.AccountIntegrationAzureDelete(&updAccInfo)
	{
		if errDelAcc != nil {
			log.Fatal(errDelAcc)
		}
	}
}
