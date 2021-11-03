package createjwt

import (
	"github.com/Tennyx/private-key-jwt-go/envgen"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
	"fmt"
	b64 "encoding/base64"
	"strings"
	"net/http"
	"io/ioutil"
)

func init() {
	envgen.LoadEnv()
}

func CreateJwt(url string, clientId string) string {
	b64PrivateKey := os.Getenv("PRIVATE_KEY")
	privateKey, _ := b64.StdEncoding.DecodeString(b64PrivateKey)
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		fmt.Println(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": clientId,
		"sub": clientId,
		"aud": url + "/oauth2/v1/token",
		"exp": time.Now().Unix() + 3600,
	})

	tokenString, err := token.SignedString(pkey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func GetAccessToken(){
	oktaUrl := os.Getenv("OKTA_URL")
	scopes := os.Getenv("SCOPES")
	clientId := os.Getenv("CLIENT_ID")

	privateKeyJwt := CreateJwt(oktaUrl, clientId)

	tokenUrl := oktaUrl + "/oauth2/v1/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&scope="+ scopes +"&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer&client_assertion=" + privateKeyJwt)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, tokenUrl, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}	
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n///////////////// Access Token: \n \n", string(body))
}