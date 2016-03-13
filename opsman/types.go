package opsman

// Ops Manager installation json types
type (
	// Version information
	Version struct {
		Version string `json:"version"`
	}

	// Deployment information
	Deployment struct {
		Release                     string
		UaaJobs                     []string
		RouterJobs                  []string
		CloudControllerJobs         []string
		CloudControllerDatabaseJobs []string
		DiegoBrainJobs              []string
		DiegoCellJobs               []string
		DiegoDatabaseJobs           []string
		UaaDatabaseJobs             []string
	}

	// Partition information
	Partition struct {
		JobReference              string `json:"job_reference"`
		InstallationName          string `json:"installation_name"`
		InstanceCount             int    `json:"instance_count"`
		AvailabilityZoneReference string `json:"availability_zone_reference"`
	}

	// Products contains all the installed products in an installation
	Products struct {
		Name string              `json:"installation_name"`
		GUID string              `json:"guid"`
		Type string              `json:"type"`
		IPS  map[string][]string `json:"ips"`
		Jobs []Jobs              `json:"jobs"`
	}

	// InstallationSettings contains the installationsettings elements from the json
	InstallationSettings struct {
		Infrastructure Infrastructure `json:"infrastructure"`
		Products       []Products     `json:"products"`
	}

	// Infrastructure contains Infrastructure block elements from the json
	Infrastructure struct {
		Type       string            `json:"type"`
		IaaSConfig IaaSConfiguration `json:"iaas_configuration"`
	}

	// IaaSConfiguration contains the IaaSConfiguration block elements from the json
	IaaSConfiguration struct {
		SSHPrivateKey string `json:"ssh_private_key"`
	}

	// Product contains installation settings for a product
	Product struct {
		Identifer      string              `json:"identifier"`
		IPS            map[string][]string `json:"ips"`
		Jobs           []Jobs              `json:"jobs"`
		ProductVersion string              `json:"product_version"`
	}

	// Jobs contains job settings for a product
	Jobs struct {
		Identifier       string       `json:"identifier"`
		InstallationName string       `json:"installation_name"`
		Properties       []Properties `json:"properties"`
		Instances        []Instances  `json:"instances"`
		Type             string       `json:"type"`
		GUID             string       `json:"guid"`
		Partition        []Partition  `json:"partitions"`
	}

	// Properties contains property settings for a job
	Properties struct {
		Definition string      `json:"definition"`
		Value      interface{} `json:"value"`
	}

	// Instances contains instances for a job
	Instances struct {
		Identifier string `json:"identifier"`
		Value      int    `json:"value"`
	}
)
