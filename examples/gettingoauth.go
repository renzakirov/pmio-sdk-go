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
