package tests

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

func (suite *ClientTestSuite) SetupTest() {
	cfg := OauthConfig{}
	configFile, err := os.Open("../env.json")
	defer configFile.Close()
	suite.Require().NoError(err)

	err = json.NewDecoder(configFile).Decode(&cfg)
	suite.Require().NoError(err)

	api := pmio.NewClient()
	api.Configuration.BasePath = "https://" + cfg.Host + "/api/v1"
	api.Configuration.Host = cfg.Host
	api.Configuration.APIKey["Authorization"] = cfg.Key
	api.Configuration.APIKeyPrefix["Authorization"] = "Bearer"
	api.Configuration.AccessToken = cfg.Key
	suite.client = api
	rand.Seed(time.Now().UnixNano())
}
