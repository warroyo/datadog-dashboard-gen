package opsman_test

import (
	"fmt"
	"io/ioutil"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDatadogDashboardGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DatadogDashboardGen Suite")
}

func fixture(name string) string {
	filePath := path.Join("../fixtures", name)
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("could not read fixture: %s", name))
	}

	return string(contents)
}
