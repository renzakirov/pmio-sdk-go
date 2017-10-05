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
