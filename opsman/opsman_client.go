package opsman

import (
	"bytes"
	"encoding/json"
	"fmt"

	http "github.com/pivotalservices/datadog-dashboard-gen/http"
)

type Client struct {
	opsmanIP       string
	opsmanUsername string
	opsmanPassword string
}

func New(opsmanIP, opsmanUsername, opsmanPassword string) *Client {
	return &Client{
		opsmanIP:       opsmanIP,
		opsmanUsername: opsmanUsername,
		opsmanPassword: opsmanPassword,
	}
}

func (c *Client) GetAPIVersion() (string, error) {
	url := "https://" + c.opsmanIP + "/api/api_version"
	resp, err := http.SendRequest("GET", url, c.opsmanUsername, c.opsmanPassword, "")
	if err != nil {
		return "", err
	}
	res := bytes.NewBufferString(resp)
	decoder := json.NewDecoder(res)
	var ver Version
	err = decoder.Decode(&ver)
	if err != nil {
		return "", err
	}

	return ver.Version, nil
}

// GetCFDeployment returns the Elastic-Runtime deployment created by your Ops Manager
func (c *Client) GetCFDeployment(installation *InstallationSettings, products []Products) (*Deployment, error) {
	cfRelease := getProductGUID(products, "cf")
	if cfRelease == "" {
		return nil, fmt.Errorf("cf release not found")
	}

	return NewDeployment(installation, cfRelease), nil
}

// GetInstallationSettings retrieves installation settings for cf deployment
func (c *Client) GetInstallationSettings() (*InstallationSettings, error) {
	url := "https://" + c.opsmanIP + "/api/installation_settings/"
	resp, err := http.SendRequest("GET", url, c.opsmanUsername, c.opsmanPassword, "")
	res := bytes.NewBufferString(resp)
	decoder := json.NewDecoder(res)
	var installation *InstallationSettings
	err = decoder.Decode(&installation)

	return installation, err
}

// GetProducts returns all the products in an OpsMan installation
func (c *Client) GetProducts() ([]Products, error) {
	url := "https://" + c.opsmanIP + "/api/installation_settings/products"
	resp, err := http.SendRequest("GET", url, c.opsmanUsername, c.opsmanPassword, "")
	if err != nil {
		return nil, err
	}
	res := bytes.NewBufferString(resp)
	decoder := json.NewDecoder(res)
	var products []Products
	err = decoder.Decode(&products)

	return products, err
}

// gets the product GUID for a given product type
func getProductGUID(products []Products, productType string) string {
	for prod := range products {
		if products[prod].Type == productType {
			return products[prod].GUID
		}
	}
	return ""
}
