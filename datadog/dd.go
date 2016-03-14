package datadog

import http "github.com/pivotalservices/datadog-dashboard-gen/http"

// CreateStoplightDashboard sends dashboard to Datadog
func CreateStoplightDashboard(apikey string, appkey string, dddash string) (string, error) {
	url := "https://app.datadoghq.com/api/v1/screen?api_key=" + apikey + "&application_key=" + appkey
	return http.SendRequest("POST", url, "", "", dddash)
}
