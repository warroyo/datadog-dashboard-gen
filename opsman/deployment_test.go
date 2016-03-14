package opsman_test

import (
	"bytes"
	"encoding/json"

	. "github.com/pivotalservices/datadog-dashboard-gen/opsman"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deployment", func() {
	var settingsJSON string
	var deployment *Deployment
	var is *InstallationSettings

	BeforeEach(func() {
		settingsJSON = fixture("installation_settings.json")
	})

	JustBeforeEach(func() {
		is = NewInstallationSettingsJSON(settingsJSON)
		deployment = NewDeployment(is, "cf-6455120728b109a1086c")
	})

	Describe("calling NewDeployment", func() {
		Context("when returns succesfully", func() {

			It("should populate the fields correctly", func() {
				Expect(deployment.Release).To(Equal("cf-6455120728b109a1086c"))
				Expect(deployment.UaaJobs).To(HaveLen(1))
			})

			// It("should not error", func() {
			// 	Expect(err).NotTo(HaveOccurred())
			// })
		})
	})
})

func NewInstallationSettingsJSON(metadata string) *InstallationSettings {
	resp := bytes.NewBufferString(metadata)
	decoder := json.NewDecoder(resp)
	var is *InstallationSettings
	decoder.Decode(&is)
	return is
}
