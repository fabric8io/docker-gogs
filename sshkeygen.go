package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	priv, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Get der format. priv_der []byte
	priv_der := x509.MarshalPKCS1PrivateKey(priv)

	// pem.Block
	// blk pem.Block
	priv_blk := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priv_der,
	}

	// Resultant private key in PEM format.
	// priv_pem string
	privBytes := pem.EncodeToMemory(&priv_blk)

	// Public Key generation
	sshPublicKey, err := ssh.NewPublicKey(&priv.PublicKey)
	pubBytes := ssh.MarshalAuthorizedKey(sshPublicKey)

	if err = writeFile("gogs.rsa", privBytes); err != nil && !os.IsExist(err) {
		panic(err)
	}
	if err = writeFile("gogs.rsa.pub", pubBytes); err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func writeFile(path string, contents []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(contents); err != nil {
		return err
	}

	return nil
}
