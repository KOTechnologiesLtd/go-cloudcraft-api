package main

import (
	"log"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
)

func main() {

	client := cloudcraft.NewClient("APIKEY", "https://api.cloudcraft.co", 3)
	cfurl := client.GetBaseURL()

	//Create a CloudCraft AWS integration
	cloudcraftAwsIntegrationname := "AWSIntegration"
	roleArn := "YOURAWSROLEARN"
	newAccInfo := cloudcraft.AccountIntegrationAws{Name: &cloudcraftAwsIntegrationname, RoleArn: &roleArn}

	log.Printf("CloudCraft Base URL %s\n", cfurl)
	errNewAcc := client.AccountIntegrationAwsCreate(&newAccInfo)
	{
		if errNewAcc != nil {
			log.Fatal(errNewAcc)
		}
	}

	log.Printf("Integration ID %s", *newAccInfo.ID)
	log.Printf("Integration Name%s", *newAccInfo.Name)

	//Get a CloudCraft AWS integration
	Info, err := client.AccountIntegrationAws(*newAccInfo.ID)
	{
		if err != nil {
			log.Fatal(err)
		}
	}
	ID := *Info.ID
	Name := *Info.Name
	log.Printf("ID - %s, Name - %s", ID, Name)

	//Get all CloudCraft AWS integrations
	accounts, err := client.AccountIntegrationsAllAws()
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

	//Update a CloudCraft AWS integration
	updID := *newAccInfo.ID
	updname := "AWSIntegrationUpdate"

	updAccInfo := cloudcraft.AccountIntegrationAws{ID: &updID, Name: &updname, RoleArn: &roleArn}
	errUpdAcc := client.AccountIntegrationAwsUpdate(&updAccInfo)
	{
		if errUpdAcc != nil {
			log.Fatal(errUpdAcc)
		}
	}

	log.Printf("Integration ID %s", *updAccInfo.ID)
	log.Printf("Integration Name %s", *updAccInfo.Name)

	//Delete a CloudCraft AWS integration
	errDelAcc := client.AccountIntegrationAwsDelete(&updAccInfo)
	{
		if errDelAcc != nil {
			log.Fatal(errDelAcc)
		}
	}

	//Get CloudCraft AWS integration IAM Params
	iamParams, err := client.AccountIntegrationAwsIamParams()
	{
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("IAM Param AccountID %s", *iamParams.AccountID)
	log.Printf("IAM Param ExternalID %s", *iamParams.ExternalID)
	log.Printf("IAM Param AwsConsoleUrl %s", *iamParams.AwsConsoleUrl)
}
