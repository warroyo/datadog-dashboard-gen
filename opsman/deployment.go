package opsman

import "fmt"

// NewDeployment creates a new installation deployment
func NewDeployment(installation *InstallationSettings, cfRelease string) *Deployment {

	var uaaJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "uaa") {
		if p.InstanceCount > 0 {
			uaaJobParts = append(uaaJobParts, p.InstallationName)
		}
	}
	var routerJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "router") {
		if p.InstanceCount > 0 {
			routerJobParts = append(routerJobParts, p.InstallationName)
		}
	}
	var ccJobParts []string
	for _, p := range getPartitions(installation, cfRelease, "cloud_controller") {
		if p.InstanceCount > 0 {
			ccJobParts = append(ccJobParts, p.InstallationName)
		}
	}
	var diegoBrainParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_brain") {
		if p.InstanceCount > 0 {
			diegoBrainParts = append(diegoBrainParts, p.InstallationName)
		}
	}
	var diegoCellParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_cell") {
		if p.InstanceCount > 0 {
			diegoCellParts = append(diegoCellParts, p.InstallationName)
		}
	}
	var diegoDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "diego_database") {
		if p.InstanceCount > 0 {
			diegoDatabaseParts = append(diegoDatabaseParts, p.InstallationName)
		}
	}
	var uaaDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "uaadb") {
		if p.InstanceCount > 0 {
			uaaDatabaseParts = append(uaaDatabaseParts, p.InstallationName)
		}
	}
	var ccJobDatabaseParts []string
	for _, p := range getPartitions(installation, cfRelease, "ccdb") {
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

// ValidateAPIVersion checks for a supported API version
func ValidateAPIVersion(client *Client) error {
	version, err := client.GetAPIVersion()
	if err != nil {
		return err
	}

	if version != "2.0" {
		return fmt.Errorf("This version of Ops Manager (using api version ''" + version + "') is not supported")
	}

	return nil
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
