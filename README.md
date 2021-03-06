# PMIO

PMIO - Go client for PMIO
This ProcessMaker I/O API provides access to a BPMN 2.0 compliant workflow engine API that is designed to be used as a microservice to support enterprise cloud applications. The current Alpha 1.0 version supports most of the descriptive classes of the BPMN 2.0 specification.
This SDK is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 1.0.0
- Package version: 1.0.0
- Build date: 2017-09-26T12:29:56.762+03:00

For more information, please visit [https://www.processmaker.io/](https://www.processmaker.io/)

## Installation

```shell
go get github.com/ProcessMaker/pmio-sdk-go
```
### For glide
```
glide get github.com/ProcessMaker/pmio-sdk-go
```
### For dep
```
dep ensure -add github.com/ProcessMaker/pmio-sdk-go
```

## Install the required libraries:
### [Resty](https://github.com/go-resty/resty)

```
go get -u gopkg.in/go-resty/resty.v0
```

### [Testify](https://github.com/stretchr/testify)
```
go get github.com/stretchr/testify
```

## Getting Started

You can run this code in examples directory

```
package main

import (
	"os"
    "log"
    "github.com/ProcessMaker/pmio-sdk-go/pmio"
)


func main() {
    //You should define the server URL (PMIO_API_HOST) and authorization API key (PMIO_ACCESS_TOKEN)
    //as environment variables or write its inplace
    /// you can find this variables in your Processmaker.io account
    var apiHost string = os.Getenv("PMIO_API_HOST")
    var accessToken string = os.Getenv("PMIO_ACCESS_TOKEN")

    if apiHost == "" || accessToken == "" {
        log.Fatal("You should define the server URL (PMIO_API_HOST) and authorization API key (PMIO_ACCESS_TOKEN)")
    }

    //make API client and configurate them
    api := pmio.NewClient()
    api.Configuration.BasePath = "https://" + apiHost + "/api/v1"
    api.Configuration.Host = apiHost
    api.Configuration.APIKey["Authorization"] = accessToken
    api.Configuration.APIKeyPrefix["Authorization"] = "Bearer"
    api.Configuration.AccessToken = accessToken

    //call API method. This method returns user information using a token
    userItem, _, err := api.MyselfUser(1,100)

    if err != nil {
        log.Println(err)
    } else {
        log.Println("userItem = ", userItem)
    }
}
```

## Getting Oauth authorization key example
Having both `client_id` and `client_secret` you may retrieve `access_token` using password grant. Additionally `username` and `password` are required to perform the operation.

```
// +build ignore

package main

import (
	"log"
	"encoding/json"
	"os"
	"gopkg.in/go-resty/resty.v0"
)


type AuthConfig struct {
	GrantType string `json:"grant_type"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
}

type AuthSuccess struct {
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	var cfg AuthConfig
	configFile, err := os.Open("authconfig.json")
	defer configFile.Close()
	if err != nil {
		log.Println("Please rename and set fields of authconfig.example.json. More info in readme...")
		log.Fatal(err)
	}
	
	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	args := make(map[string]string)
	args["grant_type"] = cfg.GrantType
	args["client_id"] = cfg.ClientID
	args["client_secret"] = cfg.ClientSecret
	args["username"] = cfg.UserName
	args["password"] = cfg.Password
	
	auth := &AuthSuccess{}

	_, err = resty.R().
	SetFormData(args).SetResult(auth).Post("https://" + cfg.Host + "/oauth/access_token")

	if err != nil {
		log.Println("err = ", err)
	} else {
		log.Println("TokenType = ", auth.TokenType)
		log.Println("ExpiresIn = ", auth.ExpiresIn)
		log.Println("AccessToken = ", auth.AccessToken)
		log.Println("RefreshToken = ", auth.RefreshToken)
	}
}
```
This code is located in the folder `examples`

To run this code you need:

1. Rename file `authconfig.example.json` to `authconfig.json`

2. Write down your values, which you can find in your Processmaker.io account (`client_id`, `client_secret`, `username` and `password`)

3. Run:  
```go run examples/gettingoauth.go ```


## How run Go-SDK Tests

The test files are in the directory `tests`. 
To run tests code you need:

1. Rename file `env.example.json` to `env.json`

2. Write down your values domain name and authorization key

3. Run:
```go test github.com/ProcessMaker/pmio-sdk-go/tests``` or ```go test -v github.com/ProcessMaker/pmio-sdk-go/tests```


## Documentation for API Endpoints

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Client* | [**AddEvent**](pmio/docs/Client.md#addevent) | **Post** /processes/{process_id}/events | 
*Client* | [**AddEventConnector**](pmio/docs/Client.md#addeventconnector) | **Post** /processes/{process_id}/events/{event_id}/connectors | 
*Client* | [**AddFlow**](pmio/docs/Client.md#addflow) | **Post** /processes/{process_id}/flows | 
*Client* | [**AddGateway**](pmio/docs/Client.md#addgateway) | **Post** /processes/{process_id}/gateways | 
*Client* | [**AddGroup**](pmio/docs/Client.md#addgroup) | **Post** /groups | 
*Client* | [**AddGroupsToTask**](pmio/docs/Client.md#addgroupstotask) | **Put** /processes/{process_id}/tasks/{task_id}/groups | 
*Client* | [**AddInputOutput**](pmio/docs/Client.md#addinputoutput) | **Post** /processes/{process_id}/tasks/{task_id}/inputoutput | 
*Client* | [**AddInstance**](pmio/docs/Client.md#addinstance) | **Post** /processes/{process_id}/instances | 
*Client* | [**AddOauthClient**](pmio/docs/Client.md#addoauthclient) | **Post** /users/{user_id}/clients | 
*Client* | [**AddProcess**](pmio/docs/Client.md#addprocess) | **Post** /processes | 
*Client* | [**AddTask**](pmio/docs/Client.md#addtask) | **Post** /processes/{process_id}/tasks | 
*Client* | [**AddTaskConnector**](pmio/docs/Client.md#addtaskconnector) | **Post** /processes/{process_id}/tasks/{task_id}/connectors | 
*Client* | [**AddUser**](pmio/docs/Client.md#adduser) | **Post** /users | 
*Client* | [**AddUsersToGroup**](pmio/docs/Client.md#adduserstogroup) | **Put** /groups/{id}/users | 
*Client* | [**DeleteEvent**](pmio/docs/Client.md#deleteevent) | **Delete** /processes/{process_id}/events/{event_id} | 
*Client* | [**DeleteEventConnector**](pmio/docs/Client.md#deleteeventconnector) | **Delete** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*Client* | [**DeleteFlow**](pmio/docs/Client.md#deleteflow) | **Delete** /processes/{process_id}/flows/{flow_id} | 
*Client* | [**DeleteGateway**](pmio/docs/Client.md#deletegateway) | **Delete** /processes/{process_id}/gateways/{gateway_id} | 
*Client* | [**DeleteGroup**](pmio/docs/Client.md#deletegroup) | **Delete** /groups/{id} | 
*Client* | [**DeleteInputOutput**](pmio/docs/Client.md#deleteinputoutput) | **Delete** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*Client* | [**DeleteInstance**](pmio/docs/Client.md#deleteinstance) | **Delete** /processes/{process_id}/instances/{instance_id} | 
*Client* | [**DeleteOauthClient**](pmio/docs/Client.md#deleteoauthclient) | **Delete** /users/{user_id}/clients/{client_id} | 
*Client* | [**DeleteProcess**](pmio/docs/Client.md#deleteprocess) | **Delete** /processes/{id} | 
*Client* | [**DeleteTask**](pmio/docs/Client.md#deletetask) | **Delete** /processes/{process_id}/tasks/{task_id} | 
*Client* | [**DeleteTaskConnector**](pmio/docs/Client.md#deletetaskconnector) | **Delete** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*Client* | [**DeleteUser**](pmio/docs/Client.md#deleteuser) | **Delete** /users/{id} | 
*Client* | [**EventTrigger**](pmio/docs/Client.md#eventtrigger) | **Post** /processes/{process_id}/events/{event_id}/trigger | 
*Client* | [**EventWebhook**](pmio/docs/Client.md#eventwebhook) | **Post** /processes/{process_id}/events/{event_id}/webhook | 
*Client* | [**FindByFieldInsideDataModel**](pmio/docs/Client.md#findbyfieldinsidedatamodel) | **Get** /processes/{process_id}/datamodels/search/{search_param} | 
*Client* | [**FindDataModel**](pmio/docs/Client.md#finddatamodel) | **Get** /processes/{process_id}/instances/{instance_id}/datamodel | 
*Client* | [**FindEventById**](pmio/docs/Client.md#findeventbyid) | **Get** /processes/{process_id}/events/{event_id} | 
*Client* | [**FindEventConnectorById**](pmio/docs/Client.md#findeventconnectorbyid) | **Get** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*Client* | [**FindEventConnectors**](pmio/docs/Client.md#findeventconnectors) | **Get** /processes/{process_id}/events/{event_id}/connectors | 
*Client* | [**FindEvents**](pmio/docs/Client.md#findevents) | **Get** /processes/{process_id}/events | 
*Client* | [**FindFlowById**](pmio/docs/Client.md#findflowbyid) | **Get** /processes/{process_id}/flows/{flow_id} | 
*Client* | [**FindFlows**](pmio/docs/Client.md#findflows) | **Get** /processes/{process_id}/flows | 
*Client* | [**FindGatewayById**](pmio/docs/Client.md#findgatewaybyid) | **Get** /processes/{process_id}/gateways/{gateway_id} | 
*Client* | [**FindGateways**](pmio/docs/Client.md#findgateways) | **Get** /processes/{process_id}/gateways | 
*Client* | [**FindGroupById**](pmio/docs/Client.md#findgroupbyid) | **Get** /groups/{id} | 
*Client* | [**FindGroups**](pmio/docs/Client.md#findgroups) | **Get** /groups | 
*Client* | [**FindInputOutputById**](pmio/docs/Client.md#findinputoutputbyid) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*Client* | [**FindInputOutputs**](pmio/docs/Client.md#findinputoutputs) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput | 
*Client* | [**FindInstanceById**](pmio/docs/Client.md#findinstancebyid) | **Get** /processes/{process_id}/instances/{instance_id} | 
*Client* | [**FindInstances**](pmio/docs/Client.md#findinstances) | **Get** /processes/{process_id}/instances | 
*Client* | [**FindOauthClientById**](pmio/docs/Client.md#findoauthclientbyid) | **Get** /users/{user_id}/clients/{client_id} | 
*Client* | [**FindOauthClients**](pmio/docs/Client.md#findoauthclients) | **Get** /users/{user_id}/clients | 
*Client* | [**FindProcessById**](pmio/docs/Client.md#findprocessbyid) | **Get** /processes/{id} | 
*Client* | [**FindProcesses**](pmio/docs/Client.md#findprocesses) | **Get** /processes | 
*Client* | [**FindTaskById**](pmio/docs/Client.md#findtaskbyid) | **Get** /processes/{process_id}/tasks/{task_id} | 
*Client* | [**FindTaskConnectorById**](pmio/docs/Client.md#findtaskconnectorbyid) | **Get** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*Client* | [**FindTaskConnectors**](pmio/docs/Client.md#findtaskconnectors) | **Get** /processes/{process_id}/tasks/{task_id}/connectors | 
*Client* | [**FindTaskInstanceById**](pmio/docs/Client.md#findtaskinstancebyid) | **Get** /task_instances/{task_instance_id} | 
*Client* | [**FindTaskInstances**](pmio/docs/Client.md#findtaskinstances) | **Get** /task_instances | 
*Client* | [**FindTaskInstancesByInstanceAndTaskId**](pmio/docs/Client.md#findtaskinstancesbyinstanceandtaskid) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances | 
*Client* | [**FindTaskInstancesByInstanceAndTaskIdDelegated**](pmio/docs/Client.md#findtaskinstancesbyinstanceandtaskiddelegated) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/delegated | 
*Client* | [**FindTaskInstancesByInstanceAndTaskIdStarted**](pmio/docs/Client.md#findtaskinstancesbyinstanceandtaskidstarted) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/started | 
*Client* | [**FindTasks**](pmio/docs/Client.md#findtasks) | **Get** /processes/{process_id}/tasks | 
*Client* | [**FindTokens**](pmio/docs/Client.md#findtokens) | **Get** /processes/{process_id}/instances/{instance_id}/tokens | 
*Client* | [**FindUserById**](pmio/docs/Client.md#finduserbyid) | **Get** /users/{id} | 
*Client* | [**FindUsers**](pmio/docs/Client.md#findusers) | **Get** /users | 
*Client* | [**ImportBpmnFile**](pmio/docs/Client.md#importbpmnfile) | **Post** /processes/import | 
*Client* | [**MyselfUser**](pmio/docs/Client.md#myselfuser) | **Get** /users/myself | 
*Client* | [**RemoveGroupsFromTask**](pmio/docs/Client.md#removegroupsfromtask) | **Delete** /processes/{process_id}/tasks/{task_id}/groups | 
*Client* | [**RemoveUsersFromGroup**](pmio/docs/Client.md#removeusersfromgroup) | **Delete** /groups/{id}/users | 
*Client* | [**SyncGroupsToTask**](pmio/docs/Client.md#syncgroupstotask) | **Post** /processes/{process_id}/tasks/{task_id}/groups | 
*Client* | [**SyncUsersToGroup**](pmio/docs/Client.md#syncuserstogroup) | **Post** /groups/{id}/users | 
*Client* | [**UpdateEvent**](pmio/docs/Client.md#updateevent) | **Put** /processes/{process_id}/events/{event_id} | 
*Client* | [**UpdateEventConnector**](pmio/docs/Client.md#updateeventconnector) | **Put** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*Client* | [**UpdateFlow**](pmio/docs/Client.md#updateflow) | **Put** /processes/{process_id}/flows/{flow_id} | 
*Client* | [**UpdateGateway**](pmio/docs/Client.md#updategateway) | **Put** /processes/{process_id}/gateways/{gateway_id} | 
*Client* | [**UpdateGroup**](pmio/docs/Client.md#updategroup) | **Put** /groups/{id} | 
*Client* | [**UpdateInputOutput**](pmio/docs/Client.md#updateinputoutput) | **Put** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*Client* | [**UpdateInstance**](pmio/docs/Client.md#updateinstance) | **Put** /processes/{process_id}/instances/{instance_id} | 
*Client* | [**UpdateOauthClient**](pmio/docs/Client.md#updateoauthclient) | **Put** /users/{user_id}/clients/{client_id} | 
*Client* | [**UpdateProcess**](pmio/docs/Client.md#updateprocess) | **Put** /processes/{id} | 
*Client* | [**UpdateTask**](pmio/docs/Client.md#updatetask) | **Put** /processes/{process_id}/tasks/{task_id} | 
*Client* | [**UpdateTaskConnector**](pmio/docs/Client.md#updatetaskconnector) | **Put** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*Client* | [**UpdateTaskInstance**](pmio/docs/Client.md#updatetaskinstance) | **Patch** /task_instances/{task_instance_id} | 
*Client* | [**UpdateUser**](pmio/docs/Client.md#updateuser) | **Put** /users/{id} | 
*EventApi* | [**AddEvent**](pmio/docs/EventApi.md#addevent) | **Post** /processes/{process_id}/events | 
*EventApi* | [**AddEventConnector**](pmio/docs/EventApi.md#addeventconnector) | **Post** /processes/{process_id}/events/{event_id}/connectors | 
*EventApi* | [**DeleteEvent**](pmio/docs/EventApi.md#deleteevent) | **Delete** /processes/{process_id}/events/{event_id} | 
*EventApi* | [**DeleteEventConnector**](pmio/docs/EventApi.md#deleteeventconnector) | **Delete** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*EventApi* | [**EventTrigger**](pmio/docs/EventApi.md#eventtrigger) | **Post** /processes/{process_id}/events/{event_id}/trigger | 
*EventApi* | [**EventWebhook**](pmio/docs/EventApi.md#eventwebhook) | **Post** /processes/{process_id}/events/{event_id}/webhook | 
*EventApi* | [**FindEventById**](pmio/docs/EventApi.md#findeventbyid) | **Get** /processes/{process_id}/events/{event_id} | 
*EventApi* | [**FindEventConnectorById**](pmio/docs/EventApi.md#findeventconnectorbyid) | **Get** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*EventApi* | [**FindEventConnectors**](pmio/docs/EventApi.md#findeventconnectors) | **Get** /processes/{process_id}/events/{event_id}/connectors | 
*EventApi* | [**FindEvents**](pmio/docs/EventApi.md#findevents) | **Get** /processes/{process_id}/events | 
*EventApi* | [**UpdateEvent**](pmio/docs/EventApi.md#updateevent) | **Put** /processes/{process_id}/events/{event_id} | 
*EventApi* | [**UpdateEventConnector**](pmio/docs/EventApi.md#updateeventconnector) | **Put** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
*FlowApi* | [**AddFlow**](pmio/docs/FlowApi.md#addflow) | **Post** /processes/{process_id}/flows | 
*FlowApi* | [**DeleteFlow**](pmio/docs/FlowApi.md#deleteflow) | **Delete** /processes/{process_id}/flows/{flow_id} | 
*FlowApi* | [**FindFlowById**](pmio/docs/FlowApi.md#findflowbyid) | **Get** /processes/{process_id}/flows/{flow_id} | 
*FlowApi* | [**FindFlows**](pmio/docs/FlowApi.md#findflows) | **Get** /processes/{process_id}/flows | 
*FlowApi* | [**UpdateFlow**](pmio/docs/FlowApi.md#updateflow) | **Put** /processes/{process_id}/flows/{flow_id} | 
*GatewayApi* | [**AddGateway**](pmio/docs/GatewayApi.md#addgateway) | **Post** /processes/{process_id}/gateways | 
*GatewayApi* | [**DeleteGateway**](pmio/docs/GatewayApi.md#deletegateway) | **Delete** /processes/{process_id}/gateways/{gateway_id} | 
*GatewayApi* | [**FindGatewayById**](pmio/docs/GatewayApi.md#findgatewaybyid) | **Get** /processes/{process_id}/gateways/{gateway_id} | 
*GatewayApi* | [**FindGateways**](pmio/docs/GatewayApi.md#findgateways) | **Get** /processes/{process_id}/gateways | 
*GatewayApi* | [**UpdateGateway**](pmio/docs/GatewayApi.md#updategateway) | **Put** /processes/{process_id}/gateways/{gateway_id} | 
*GroupsApi* | [**AddGroup**](pmio/docs/GroupsApi.md#addgroup) | **Post** /groups | 
*GroupsApi* | [**AddUsersToGroup**](pmio/docs/GroupsApi.md#adduserstogroup) | **Put** /groups/{id}/users | 
*GroupsApi* | [**DeleteGroup**](pmio/docs/GroupsApi.md#deletegroup) | **Delete** /groups/{id} | 
*GroupsApi* | [**FindGroupById**](pmio/docs/GroupsApi.md#findgroupbyid) | **Get** /groups/{id} | 
*GroupsApi* | [**FindGroups**](pmio/docs/GroupsApi.md#findgroups) | **Get** /groups | 
*GroupsApi* | [**RemoveUsersFromGroup**](pmio/docs/GroupsApi.md#removeusersfromgroup) | **Delete** /groups/{id}/users | 
*GroupsApi* | [**SyncUsersToGroup**](pmio/docs/GroupsApi.md#syncuserstogroup) | **Post** /groups/{id}/users | 
*GroupsApi* | [**UpdateGroup**](pmio/docs/GroupsApi.md#updategroup) | **Put** /groups/{id} | 
*InputoutputApi* | [**AddInputOutput**](pmio/docs/InputoutputApi.md#addinputoutput) | **Post** /processes/{process_id}/tasks/{task_id}/inputoutput | 
*InputoutputApi* | [**DeleteInputOutput**](pmio/docs/InputoutputApi.md#deleteinputoutput) | **Delete** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*InputoutputApi* | [**FindInputOutputById**](pmio/docs/InputoutputApi.md#findinputoutputbyid) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*InputoutputApi* | [**FindInputOutputs**](pmio/docs/InputoutputApi.md#findinputoutputs) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput | 
*InputoutputApi* | [**UpdateInputOutput**](pmio/docs/InputoutputApi.md#updateinputoutput) | **Put** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
*OauthApi* | [**AddOauthClient**](pmio/docs/OauthApi.md#addoauthclient) | **Post** /users/{user_id}/clients | 
*OauthApi* | [**FindOauthClientById**](pmio/docs/OauthApi.md#findoauthclientbyid) | **Get** /users/{user_id}/clients/{client_id} | 
*OauthApi* | [**FindOauthClients**](pmio/docs/OauthApi.md#findoauthclients) | **Get** /users/{user_id}/clients | 
*ProcessApi* | [**AddProcess**](pmio/docs/ProcessApi.md#addprocess) | **Post** /processes | 
*ProcessApi* | [**DeleteProcess**](pmio/docs/ProcessApi.md#deleteprocess) | **Delete** /processes/{id} | 
*ProcessApi* | [**FindProcessById**](pmio/docs/ProcessApi.md#findprocessbyid) | **Get** /processes/{id} | 
*ProcessApi* | [**FindProcesses**](pmio/docs/ProcessApi.md#findprocesses) | **Get** /processes | 
*ProcessApi* | [**ImportBpmnFile**](pmio/docs/ProcessApi.md#importbpmnfile) | **Post** /processes/import | 
*ProcessApi* | [**UpdateProcess**](pmio/docs/ProcessApi.md#updateprocess) | **Put** /processes/{id} | 
*ProcessInstanceApi* | [**AddInstance**](pmio/docs/ProcessInstanceApi.md#addinstance) | **Post** /processes/{process_id}/instances | 
*ProcessInstanceApi* | [**DeleteInstance**](pmio/docs/ProcessInstanceApi.md#deleteinstance) | **Delete** /processes/{process_id}/instances/{instance_id} | 
*ProcessInstanceApi* | [**FindByFieldInsideDataModel**](pmio/docs/ProcessInstanceApi.md#findbyfieldinsidedatamodel) | **Get** /processes/{process_id}/datamodels/search/{search_param} | 
*ProcessInstanceApi* | [**FindDataModel**](pmio/docs/ProcessInstanceApi.md#finddatamodel) | **Get** /processes/{process_id}/instances/{instance_id}/datamodel | 
*ProcessInstanceApi* | [**FindInstanceById**](pmio/docs/ProcessInstanceApi.md#findinstancebyid) | **Get** /processes/{process_id}/instances/{instance_id} | 
*ProcessInstanceApi* | [**FindInstances**](pmio/docs/ProcessInstanceApi.md#findinstances) | **Get** /processes/{process_id}/instances | 
*ProcessInstanceApi* | [**FindTaskInstancesByInstanceAndTaskId**](pmio/docs/ProcessInstanceApi.md#findtaskinstancesbyinstanceandtaskid) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances | 
*ProcessInstanceApi* | [**FindTaskInstancesByInstanceAndTaskIdDelegated**](pmio/docs/ProcessInstanceApi.md#findtaskinstancesbyinstanceandtaskiddelegated) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/delegated | 
*ProcessInstanceApi* | [**FindTaskInstancesByInstanceAndTaskIdStarted**](pmio/docs/ProcessInstanceApi.md#findtaskinstancesbyinstanceandtaskidstarted) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/started | 
*ProcessInstanceApi* | [**FindTokens**](pmio/docs/ProcessInstanceApi.md#findtokens) | **Get** /processes/{process_id}/instances/{instance_id}/tokens | 
*ProcessInstanceApi* | [**UpdateInstance**](pmio/docs/ProcessInstanceApi.md#updateinstance) | **Put** /processes/{process_id}/instances/{instance_id} | 
*TaskApi* | [**AddGroupsToTask**](pmio/docs/TaskApi.md#addgroupstotask) | **Put** /processes/{process_id}/tasks/{task_id}/groups | 
*TaskApi* | [**AddTask**](pmio/docs/TaskApi.md#addtask) | **Post** /processes/{process_id}/tasks | 
*TaskApi* | [**AddTaskConnector**](pmio/docs/TaskApi.md#addtaskconnector) | **Post** /processes/{process_id}/tasks/{task_id}/connectors | 
*TaskApi* | [**DeleteTask**](pmio/docs/TaskApi.md#deletetask) | **Delete** /processes/{process_id}/tasks/{task_id} | 
*TaskApi* | [**DeleteTaskConnector**](pmio/docs/TaskApi.md#deletetaskconnector) | **Delete** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*TaskApi* | [**FindTaskById**](pmio/docs/TaskApi.md#findtaskbyid) | **Get** /processes/{process_id}/tasks/{task_id} | 
*TaskApi* | [**FindTaskConnectorById**](pmio/docs/TaskApi.md#findtaskconnectorbyid) | **Get** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*TaskApi* | [**FindTaskConnectors**](pmio/docs/TaskApi.md#findtaskconnectors) | **Get** /processes/{process_id}/tasks/{task_id}/connectors | 
*TaskApi* | [**FindTaskInstanceById**](pmio/docs/TaskApi.md#findtaskinstancebyid) | **Get** /task_instances/{task_instance_id} | 
*TaskApi* | [**FindTaskInstances**](pmio/docs/TaskApi.md#findtaskinstances) | **Get** /task_instances | 
*TaskApi* | [**FindTasks**](pmio/docs/TaskApi.md#findtasks) | **Get** /processes/{process_id}/tasks | 
*TaskApi* | [**RemoveGroupsFromTask**](pmio/docs/TaskApi.md#removegroupsfromtask) | **Delete** /processes/{process_id}/tasks/{task_id}/groups | 
*TaskApi* | [**SyncGroupsToTask**](pmio/docs/TaskApi.md#syncgroupstotask) | **Post** /processes/{process_id}/tasks/{task_id}/groups | 
*TaskApi* | [**UpdateTask**](pmio/docs/TaskApi.md#updatetask) | **Put** /processes/{process_id}/tasks/{task_id} | 
*TaskApi* | [**UpdateTaskConnector**](pmio/docs/TaskApi.md#updatetaskconnector) | **Put** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
*TaskApi* | [**UpdateTaskInstance**](pmio/docs/TaskApi.md#updatetaskinstance) | **Patch** /task_instances/{task_instance_id} | 
*UsersApi* | [**AddUser**](pmio/docs/UsersApi.md#adduser) | **Post** /users | 
*UsersApi* | [**DeleteUser**](pmio/docs/UsersApi.md#deleteuser) | **Delete** /users/{id} | 
*UsersApi* | [**FindUserById**](pmio/docs/UsersApi.md#finduserbyid) | **Get** /users/{id} | 
*UsersApi* | [**FindUsers**](pmio/docs/UsersApi.md#findusers) | **Get** /users | 
*UsersApi* | [**MyselfUser**](pmio/docs/UsersApi.md#myselfuser) | **Get** /users/myself | 
*UsersApi* | [**UpdateUser**](pmio/docs/UsersApi.md#updateuser) | **Put** /users/{id} | 


## Documentation For Models

 - [BpmnFile](pmio/docs/BpmnFile.md)
 - [BpmnFileAttributes](pmio/docs/BpmnFileAttributes.md)
 - [BpmnImportItem](pmio/docs/BpmnImportItem.md)
 - [DataModel](pmio/docs/DataModel.md)
 - [DataModelAttributes](pmio/docs/DataModelAttributes.md)
 - [DataModelCollection](pmio/docs/DataModelCollection.md)
 - [DataModelItem](pmio/docs/DataModelItem.md)
 - [DataModelItem1](pmio/docs/DataModelItem1.md)
 - [DataModelItemAttributes](pmio/docs/DataModelItemAttributes.md)
 - [ErrorArray](pmio/docs/ErrorArray.md)
 - [Event](pmio/docs/Event.md)
 - [EventAttributes](pmio/docs/EventAttributes.md)
 - [EventCollection](pmio/docs/EventCollection.md)
 - [EventConnector](pmio/docs/EventConnector.md)
 - [EventConnector1](pmio/docs/EventConnector1.md)
 - [EventConnectorAttributes](pmio/docs/EventConnectorAttributes.md)
 - [EventConnectorCreateItem](pmio/docs/EventConnectorCreateItem.md)
 - [EventConnectorUpdateItem](pmio/docs/EventConnectorUpdateItem.md)
 - [EventConnectorsCollection](pmio/docs/EventConnectorsCollection.md)
 - [EventCreateItem](pmio/docs/EventCreateItem.md)
 - [EventItem](pmio/docs/EventItem.md)
 - [EventUpdateItem](pmio/docs/EventUpdateItem.md)
 - [Flow](pmio/docs/Flow.md)
 - [FlowAttributes](pmio/docs/FlowAttributes.md)
 - [FlowCollection](pmio/docs/FlowCollection.md)
 - [FlowCreateItem](pmio/docs/FlowCreateItem.md)
 - [FlowItem](pmio/docs/FlowItem.md)
 - [FlowUpdateItem](pmio/docs/FlowUpdateItem.md)
 - [Gateway](pmio/docs/Gateway.md)
 - [GatewayAttributes](pmio/docs/GatewayAttributes.md)
 - [GatewayCollection](pmio/docs/GatewayCollection.md)
 - [GatewayCreateItem](pmio/docs/GatewayCreateItem.md)
 - [GatewayItem](pmio/docs/GatewayItem.md)
 - [GatewayUpdateItem](pmio/docs/GatewayUpdateItem.md)
 - [Group](pmio/docs/Group.md)
 - [GroupAddUsersItem](pmio/docs/GroupAddUsersItem.md)
 - [GroupAttributes](pmio/docs/GroupAttributes.md)
 - [GroupCollection](pmio/docs/GroupCollection.md)
 - [GroupCreateItem](pmio/docs/GroupCreateItem.md)
 - [GroupIds](pmio/docs/GroupIds.md)
 - [GroupItem](pmio/docs/GroupItem.md)
 - [GroupRemoveUsersItem](pmio/docs/GroupRemoveUsersItem.md)
 - [GroupSyncUsersItem](pmio/docs/GroupSyncUsersItem.md)
 - [GroupUpdateItem](pmio/docs/GroupUpdateItem.md)
 - [InlineResponse200](pmio/docs/InlineResponse200.md)
 - [InputOutput](pmio/docs/InputOutput.md)
 - [InputOutputAttributes](pmio/docs/InputOutputAttributes.md)
 - [InputOutputCollection](pmio/docs/InputOutputCollection.md)
 - [InputOutputCreateItem](pmio/docs/InputOutputCreateItem.md)
 - [InputOutputItem](pmio/docs/InputOutputItem.md)
 - [InputOutputUpdateItem](pmio/docs/InputOutputUpdateItem.md)
 - [Instance](pmio/docs/Instance.md)
 - [InstanceAttributes](pmio/docs/InstanceAttributes.md)
 - [InstanceCollection](pmio/docs/InstanceCollection.md)
 - [InstanceCreateItem](pmio/docs/InstanceCreateItem.md)
 - [InstanceItem](pmio/docs/InstanceItem.md)
 - [InstanceUpdateItem](pmio/docs/InstanceUpdateItem.md)
 - [Meta](pmio/docs/Meta.md)
 - [MetaLog](pmio/docs/MetaLog.md)
 - [ModelError](pmio/docs/ModelError.md)
 - [OauthClient](pmio/docs/OauthClient.md)
 - [OauthClientAttributes](pmio/docs/OauthClientAttributes.md)
 - [OauthClientCollection](pmio/docs/OauthClientCollection.md)
 - [OauthClientCreateItem](pmio/docs/OauthClientCreateItem.md)
 - [OauthClientItem](pmio/docs/OauthClientItem.md)
 - [OauthClientUpdateItem](pmio/docs/OauthClientUpdateItem.md)
 - [Pagination](pmio/docs/Pagination.md)
 - [PaginationLinks](pmio/docs/PaginationLinks.md)
 - [Process](pmio/docs/Process.md)
 - [ProcessAttributes](pmio/docs/ProcessAttributes.md)
 - [ProcessCollection](pmio/docs/ProcessCollection.md)
 - [ProcessCollection1](pmio/docs/ProcessCollection1.md)
 - [ProcessCreateItem](pmio/docs/ProcessCreateItem.md)
 - [ProcessItem](pmio/docs/ProcessItem.md)
 - [ProcessUpdateItem](pmio/docs/ProcessUpdateItem.md)
 - [ResultSuccess](pmio/docs/ResultSuccess.md)
 - [ResultSuccessMeta](pmio/docs/ResultSuccessMeta.md)
 - [Task](pmio/docs/Task.md)
 - [TaskAddGroupsItem](pmio/docs/TaskAddGroupsItem.md)
 - [TaskAttributes](pmio/docs/TaskAttributes.md)
 - [TaskCollection](pmio/docs/TaskCollection.md)
 - [TaskConnector](pmio/docs/TaskConnector.md)
 - [TaskConnector1](pmio/docs/TaskConnector1.md)
 - [TaskConnectorAttributes](pmio/docs/TaskConnectorAttributes.md)
 - [TaskConnectorCreateItem](pmio/docs/TaskConnectorCreateItem.md)
 - [TaskConnectorUpdateItem](pmio/docs/TaskConnectorUpdateItem.md)
 - [TaskConnectorsCollection](pmio/docs/TaskConnectorsCollection.md)
 - [TaskCreateItem](pmio/docs/TaskCreateItem.md)
 - [TaskInstance](pmio/docs/TaskInstance.md)
 - [TaskInstanceAttributes](pmio/docs/TaskInstanceAttributes.md)
 - [TaskInstanceCollection](pmio/docs/TaskInstanceCollection.md)
 - [TaskInstanceUpdateItem](pmio/docs/TaskInstanceUpdateItem.md)
 - [TaskItem](pmio/docs/TaskItem.md)
 - [TaskRemoveGroupsItem](pmio/docs/TaskRemoveGroupsItem.md)
 - [TaskSyncGroupsItem](pmio/docs/TaskSyncGroupsItem.md)
 - [TaskUpdateItem](pmio/docs/TaskUpdateItem.md)
 - [Token](pmio/docs/Token.md)
 - [TokenAttributes](pmio/docs/TokenAttributes.md)
 - [TokenCollection](pmio/docs/TokenCollection.md)
 - [TriggerEventCreateItem](pmio/docs/TriggerEventCreateItem.md)
 - [User](pmio/docs/User.md)
 - [UserAttributes](pmio/docs/UserAttributes.md)
 - [UserCollection](pmio/docs/UserCollection.md)
 - [UserCreateItem](pmio/docs/UserCreateItem.md)
 - [UserIds](pmio/docs/UserIds.md)
 - [UserItem](pmio/docs/UserItem.md)
 - [UserUpdateItem](pmio/docs/UserUpdateItem.md)


## Documentation For Authorization


## PasswordGrant

- **Type**: API key 
- **API key parameter name**: Authorization
- **Location**: HTTP header
