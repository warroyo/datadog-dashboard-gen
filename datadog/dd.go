package datadog

import "github.com/pivotalservices/datadog-dashboard-gen/util"

func CreateStoplightDashboard(apikey string, appkey string, dddash string) string {
	url := "https://app.datadoghq.com/api/v1/screen?api_key=" + apikey + "&application_key=" + appkey
	response := util.SendRequest("POST", url, "", "", dddash)
	return response
}
