package client

import (
	"github.com/Tennyx/private-key-jwt-go/envgen"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

type OktaApp struct {
	ClientID string `json:"client_id"`
}

func init() {
	envgen.LoadEnv()
}

func grantScopesInServiceApp(okta_url string, client_id string, scope string, apiKey string){
	url := okta_url + "/api/v1/apps/" + client_id + "/grants"
	method := "POST"

	payload := strings.NewReader(`{
		"issuer": "` + okta_url +`",
		"scopeId": "`+ scope + `"
	}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "SSWS " + apiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n///////////////// Scope granted: \n \n", string(body))
}

func RegisterOktaClient(){
	modulus := os.Getenv("MODULUS")
	apiKey := os.Getenv("API_KEY")
	scopes := os.Getenv("SCOPES")
	oktaUrl := os.Getenv("OKTA_URL")
	clientUrl := oktaUrl + "/oauth2/v1/clients"

	method := "POST"

	payload := strings.NewReader(`{
		"client_name": "Private Key JWT Service App (Go)",
		"response_types": [
			"token"
		],
		"grant_types": [
			"client_credentials"
		],
		"token_endpoint_auth_method": "private_key_jwt",
		"application_type": "service",
		"jwks": {
			"keys": [
				{
					"kty": "RSA",
					"e": "AQAB",
					"use": "sig",
					"kid": "O4O",
					"alg": "RS256",
					"n": "`+ modulus + `"
				}
			]
		}
	}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, clientUrl, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "SSWS " + os.Getenv("API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var result OktaApp

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	envgen.GenerateEnv(
					os.Getenv("PRIVATE_KEY"),
					os.Getenv("PUBLIC_KEY"),
					os.Getenv("MODULUS"),
					os.Getenv("API_KEY"),
					result.ClientID,
					scopes,
					oktaUrl,
				)

	scopesArray := strings.Fields(scopes)
	
	fmt.Println("\n///////////////// Okta Service app created: \n \n", string(body))

	for i := 0; i < len(scopesArray); i++ {
		grantScopesInServiceApp(oktaUrl, result.ClientID, scopesArray[i], apiKey)
	}
}