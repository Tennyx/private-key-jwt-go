package keygen

import (
	"github.com/Tennyx/private-key-jwt-go/envgen"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	b64 "encoding/base64"
	"fmt"
)

func init() {
	envgen.LoadEnv()
}

func GenerateKeys() {
	apiKey := os.Getenv("API_KEY")
	clientId := os.Getenv("CLIENT_ID")
	scopes := os.Getenv("SCOPES")
	oktaUrl := os.Getenv("OKTA_URL")
	bitSize := 2048

	key, err := rsa.GenerateKey(rand.Reader, bitSize)

	if err != nil {
		panic(err)
	}

	pub := key.Public()
	modulus := key.PublicKey.N

	privateKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	publicKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	// keys need to be encoded/decoded to/from base64 due to newline limitations in godotenv
	envgen.GenerateEnv(
					b64.StdEncoding.EncodeToString(privateKey),
					b64.StdEncoding.EncodeToString(publicKey),
					b64.RawURLEncoding.EncodeToString(modulus.Bytes()),
					apiKey,
					clientId,
					scopes,
					oktaUrl,
				)

	fmt.Println("\n///////////////// Keys generated and added to .env file.\n")
}