package opsman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	http "github.com/pivotalservices/datadog-dashboard-gen/http"
)

// GetCFDeployment returns the Elastic-Runtime deployment created by your Ops Manager
func GetCFDeployment(user string, passwd string, ipaddr string) (*Deployment, error) {
	//get OpsMan Version
	url := "https://" + ipaddr + "/api/api_version"
	response, err := http.SendRequest("GET", url, user, passwd, "")
	if err != nil {
		return nil, err
	}
	res := bytes.NewBufferString(response)
	decoder := json.NewDecoder(res)
	var opsmanAPI Version
	err = decoder.Decode(&opsmanAPI)
	if err != nil {
		return nil, err
	}

	if opsmanAPI.Version != "2.0" {
		log.Fatal("This version of Ops Man API version=" + opsmanAPI.Version + " is Not Supported")
	}

	// Get Ops Man CF Deployment GUID
	url = "https://" + ipaddr + "/api/installation_settings/products"
	response, err = http.SendRequest("GET", url, user, passwd, "")
	res = bytes.NewBufferString(response)
	decoder = json.NewDecoder(res)
	var products []Products
	err = decoder.Decode(&products)
	if err != nil {
		return nil, err
	}
	var cfRelease string
	for prod := range products {
		if products[prod].Type == "cf" {
			cfRelease = products[prod].GUID
		}
	}

	if cfRelease == "" {
		return nil, fmt.Errorf("cf release not found")
	}

	// Get installation settings for cf deployment
	url = "https://" + ipaddr + "/api/installation_settings/"
	response, err = http.SendRequest("GET", url, user, passwd, "")
	res = bytes.NewBufferString(response)
	decoder = json.NewDecoder(res)
	var installation *InstallationSettings
	err = decoder.Decode(&installation)
	if err != nil {
		return nil, err
	}

	//
	// Get Ops Manager Installation Metadata
	//
	var uaaJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "uaa") {
		//fmt.Println("UAA Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			uaaJobParts = append(uaaJobParts, p.InstallationName)
		}
	}
	var routerJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "router") {
		//fmt.Println("Router Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			routerJobParts = append(routerJobParts, p.InstallationName)
		}
	}
	var ccJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "cloud_controller") {
		//fmt.Println("CloudController Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			ccJobParts = append(ccJobParts, p.InstallationName)
		}
	}
	var diegoBrainParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_brain") {
		//fmt.Println("Diego Brain Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			diegoBrainParts = append(diegoBrainParts, p.InstallationName)
		}
	}
	var diegoCellParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_cell") {
		//fmt.Println("Diego Cell Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			diegoCellParts = append(diegoCellParts, p.InstallationName)
		}
	}
	var diegoDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_database") {
		//fmt.Println("Diego Database Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			diegoDatabaseParts = append(diegoDatabaseParts, p.InstallationName)
		}
	}
	var uaaDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "uaadb") {
		//fmt.Println("UAA Database Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			uaaDatabaseParts = append(uaaDatabaseParts, p.InstallationName)
		}
	}
	var ccJobDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "ccdb") {
		//fmt.Println("CC Database Partition Name:", p.InstallationName)
		if p.InstanceCount > 0 {
			ccJobDatabaseParts = append(ccJobDatabaseParts, p.InstallationName)
		}
	}

	deployment := &Deployment{
		Release:                     cfRelease,
		UaaJobs:                     uaaJobParts,
		RouterJobs:                  routerJobParts,
		CloudControllerJobs:         ccJobParts,
		CloudControllerDatabaseJobs: ccJobDatabaseParts,
		DiegoBrainJobs:              diegoBrainParts,
		DiegoCellJobs:               diegoCellParts,
		DiegoDatabaseJobs:           diegoDatabaseParts,
		UaaDatabaseJobs:             uaaDatabaseParts,
	}

	return deployment, nil
}

func getPartitions(installation *InstallationSettings, productName, jobInstallationName string) []Partition {
	for _, product := range installation.Products {
		if product.Name == productName {
			for _, job := range product.Jobs {
				if job.InstallationName == jobInstallationName {
					return job.Partition
				}
			}
		}
	}
	return nil
}
