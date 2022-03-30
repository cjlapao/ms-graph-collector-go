# Microsoft Graph Collector POC

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) ![Build](https://github.com/cjlapao/ms-graph-collector-go/workflows/Build/badge.svg) ![Release](https://github.com/cjlapao/ms-graph-collector-go/workflows/Release/badge.svg) ![Security](https://github.com/cjlapao/ms-graph-collector-go/workflows/CodeQL/badge.svg)  

This tool allows you to collect data from Microsoft Graph API.

## How to Run the API

- Open the docker-compose.yaml file
- Change the following configuration

```yaml
 version: "3.7"

 services:
   app:
     image: cjlapao/ms-graph-collector
     ports:
       - 10000:10000
     environment:
      MONGODB_CONNECTION_STRING: "" # Your mongodb connection string, this is required
      MONGODB_DATABASENAME: "saas_demo" # the api database, we will store the credentials here
      HTTP_PORT: 10000 # API port, if this is changed then you will need to change the ports in the port section
      API_PREFIX: "/api" # API prefix
      DUMP_TO_DATABASE: false # set this to true for saving all of the graph api query results into your database
      DUMP_TO_FILE: false # set this to true to save to a file the result of the collector
      DUMP_TO_FILE_PATH: ".\\" #if the DUMP_TO_FILE is true set the folder path for the file dump, we create a file per user
      PROCESS_SINGLE_USAGE: true # set this to false to collect all of the usage in the user, slower
      POST_TO_NEURONS: true # set this to false if you do not want the result to be posted to neurons
```

- Run the docker compose command

```bash
docker-compose up --force-recreate
```

This will download the latest image and start a containerized api on port 10000

## How To Use

### Create a tenant credential document

We need some credentials to be able to access the **Microsoft Graph API**, this will be a service principal created in the azure active directory and we will need the following details.

```yaml
TenantId
ClientId # Service Principal ID
ClientSecret # Service Principal Password
```

If you will need to post to Neurons we will need the following details

```yaml
LoginAppClientSecret
LoginAppClientId
NeuronsTenantId
NeuronsTenantUrl
UnoBaseUrl # normally the SFC url
UnoLoginUrl # normally the SFC url
```

With all this details you can use the endpoint ```http://localhost:10000/api/credentials``` and send a **POST** request with the following body

```json
{
    "name": "{{Tenant Name}}",
    "tenantId": "{{ServicePrincipalTenantId}}",
    "clientId": "{{ServicePrincipalClientId}}",
    "clientSecret": "{{ServicePrincipalClientSecret}}",
    "loginAppClientSecret": "{{LoginAppClientSecret}}",
    "loginAppClientId": "{{LoginAppClientId}}",
    "neuronsTenantId": "{{NeuronsTenantId}}",
    "neuronsTenantUrl": "{{NeuronsTenantUrl}}",
    "unoBaseUrl": "{{UnoBaseUrl}}",
    "unoLoginUrl": "{{UnoLoginUrl}"
}
```

### Start the collection

Using the API endpoint **GET** ```http://localhost:10000/api/{{ServicePrincipalTenantId}}/start``` this will start a collection for that specific tenant, this process can take a long time depending on the size of the tenant