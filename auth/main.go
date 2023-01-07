package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	keys()
}

func keys() {
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
}

/*
func generatePrivateKey() {
	// Generate a new RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// Encode the private key as a PEM block
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write the PEM block to a file
	if err := ioutil.WriteFile("private.pem", pem.EncodeToMemory(privateKeyPEM), 0644); err != nil {
		log.Fatal(err)
	}

}

func generatePublicKey() {
	// Load the RSA private key from a file
	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	// Extract the public key from the private key
	publicKey := rsaPrivateKey.PublicKey

	// Encode the public key as a PEM block
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&publicKey),
	}

	// Write the PEM block to a file
	if err := ioutil.WriteFile("public.pem", pem.EncodeToMemory(publicKeyPEM), 0644); err != nil {
		fmt.Println("3")
		log.Fatal(err)
	}

}
*/
