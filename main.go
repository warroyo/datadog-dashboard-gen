package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/pivotalservices/datadog-dashboard-gen/datadog"
	"github.com/pivotalservices/datadog-dashboard-gen/opsman"
)

func main() {
	// Declare Flags
	opsmanUser := flag.String("opsman_user", "admin", "Ops Manager User")
	opsmanPassword := flag.String("opsman_password", "password", "Ops Manager Password")
	opsmanIP := flag.String("opsman_ip", "192.168.200.10", "Ops Manager IP")
	ddAPIKey := flag.String("ddapikey", "12345-your-api-key-6789", "Datadog API Key")
	ddAppKey := flag.String("ddappkey", "12345-your-app-key-6789", "Datadog Application Key")

	flag.Parse()

	opsmanClient := opsman.New(*opsmanIP, *opsmanUser, *opsmanPassword)

	// Check we are using a supported Ops Man
	err := opsman.ValidateAPIVersion(opsmanClient)
	if err != nil {
		log.Fatal(err)
	}

	// Get installation settings from Ops Man foundation
	installation, err := opsmanClient.GetInstallationSettings()
	if err != nil {
		log.Fatal(err)
	}

	products, err := opsmanClient.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	deployment, err := opsmanClient.GetCFDeployment(installation, products)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = datadog.StopLightsTemplate(&buf, deployment)
	if err != nil {
		log.Fatal(err)
	}
	metadata := buf.String()

	if _, err := datadog.CreateStoplightDashboard(*ddAPIKey, *ddAppKey, metadata); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your PCF Stoplights Datadog dashboard has been published ... Go Fetch @ https://app.datadoghq.com/dash/list :)")
}
