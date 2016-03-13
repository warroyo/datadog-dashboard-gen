package opsman

import (
	"bytes"
	"encoding/json"

	"github.com/pivotalservices/datadog-dashboard-gen/util"

	"log"
	"strings"
)

// GetFoundationMetadata returns the datadog dashboard configuration
func GetFoundationMetadata(user string, passwd string, ipaddr string) *Deployment {
	//get OpsMan Version
	url := "https://" + ipaddr + "/api/api_version"
	response := util.SendRequest("GET", url, user, passwd, "")
	res := bytes.NewBufferString(response)
	decoder := json.NewDecoder(res)
	var opsmanver Version
	err := decoder.Decode(&opsmanver)
	if err != nil {
		log.Fatal(err)
	}

	if opsmanver.Version != "2.0" {
		log.Fatal("This version of Ops Man API version=" + opsmanver.Version + " is Not Supported")
	}

	// Get Ops Man CF Deployment GUID
	url = "https://" + ipaddr + "/api/installation_settings/products"
	response = util.SendRequest("GET", url, user, passwd, "")
	res = bytes.NewBufferString(response)
	decoder = json.NewDecoder(res)
	var products []Products
	err = decoder.Decode(&products)
	if err != nil {
		log.Fatal(err)
	}
	var cfRelease string
	for x := range products {
		if products[x].Type == "cf" {
			cfRelease = products[x].GUID
		}
	}
	//Test if cf Release Found
	if cfRelease == "" {
		log.Fatal("CF Release Not Found")
	}

	// Get installation settings for cf deployment
	url = "https://" + ipaddr + "/api/installation_settings/"
	response = util.SendRequest("GET", url, user, passwd, "")
	res = bytes.NewBufferString(response)
	decoder = json.NewDecoder(res)
	var installation *InstallationSettings
	err = decoder.Decode(&installation)
	if err != nil {
		log.Fatal(err)
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

	return deployment
}

func getIPAddress(installation *InstallationSettings, productName string, jobName string) string {
	for _, product := range installation.Products {
		// fmt.Println("product.Name:", product.Name)
		// fmt.Println("productName:", productName)
		if product.Name == productName {
			// fmt.Println("product.Type:", productType)
			for k, vals := range product.IPS {
				// fmt.Println("k:", k)
				if strings.Contains(k, jobName) {
					return vals[0]
				}
			}
		}
	}
	return ""
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

func getPassword(installation *InstallationSettings, productName, jobGUID, identity string) string {
	for _, product := range installation.Products {
		if product.Name == productName {
			// fmt.Println("productName:", productName)
			for _, job := range product.Jobs {
				if job.GUID == jobGUID {
					// fmt.Println("jobGUID:", jobGUID)
					for _, property := range job.Properties {
						switch property.Value.(type) {
						case map[string]interface{}:
							propertyValue := property.Value.(map[string]interface{})
							// fmt.Println("propertyValue:", propertyValue)
							field := propertyValue["identity"]
							value := propertyValue["password"]
							if field == identity {
								return value.(string)
							}
						}
					}
				}
			}
		}
	}
	return ""
}
