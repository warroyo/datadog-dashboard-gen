## Datadog Dashboard Generator

This is a command line utility that can be used to deploy a MVP "StopLights" Dashboard to a given Users Datadog Subscription from a template.

## Logical Flow

1. Tool queries PCF Ops Manager for the following:

        cf-release string (e.g. cf-76553423523ab)
        job indexes/partitions from CF manifest

2. Reads template of PCF Stoplights Dashboard

3. Uploads a generated Datadog dashboard combining Ops Man vars & template

## Example usage
```bash
    datadog-dashboard-gen -opsman_user=admin -opsman_passwd=blah -opsmanip=192.168.100.10 -ddapikey=243235r23435435345 -ddappkey=564758643636
```

## Build & Run

1. Clone repo

        git clone https://github.com/pivotalservices/datadog-dashboard-gen.git

1. Build binary

        cd datadog-dashboard-gen
        go install

1. Run program to upload the Stoplights dashboard

        $GOPATH/bin/datadog-dashboard-gen -opsman_user=<OPSMAN_USER> -opsman_password=<OPSMAN_PASSWORD> -opsman_ip=<OPSMAN_IP> \
        -ddapikey=<DATADOG_API_KEY> -ddappkey=<DATADOG_APP_KEY>

## Generate code from template

1. Install [ego](https://github.com/benbjohnson/ego)

1. Run `ego` (template is located under `templates/screen`)
```bash
ego -package datadog -o datadog/stoplights.go
```
