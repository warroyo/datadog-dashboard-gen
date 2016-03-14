package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SendRequest sends http requests
func SendRequest(method string, url string, user string, passwd string, data string) (string, error) {
	//Ignore Self Signed SSL
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//make Request Object
	req, err := http.NewRequest(method, url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}

	//Set Auth
	if user != "" && passwd != "" {
		req.SetBasicAuth(user, passwd)
	}

	//If POST set header
	if method == "POST" {
		req.Header.Add("Content-type", "application/json")
	}

	//Make Client http Request
	client := http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	//If POST verify Dashboard was published
	if method == "POST" && res.Status != "200 OK" {
		return "", fmt.Errorf("got " + res.Status + " when sending dashboard to datadog; expecting 200")
	}

	return string(body), nil
}
