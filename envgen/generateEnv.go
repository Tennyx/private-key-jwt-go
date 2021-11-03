package envgen

import (
	"github.com/joho/godotenv"
	"fmt"
)

func LoadEnv(){
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GenerateEnv(privateKey string, publicKey string, modulus string, apiKey string, clientId string, scopes string, oktaUrl string) {
	vars := "PRIVATE_KEY=" + privateKey +
			"\nPUBLIC_KEY=" + publicKey +
			"\nMODULUS=" + modulus +
			"\nAPI_KEY=" + apiKey +
			"\nSCOPES=" + scopes +
			"\nCLIENT_ID=" + clientId +
			"\nOKTA_URL=" + oktaUrl

	env, err := godotenv.Unmarshal(vars)

	if err != nil {
		fmt.Println(err)
	}

	godotenv.Write(env, "./.env")
}