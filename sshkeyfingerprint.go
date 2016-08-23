package main

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	sshkey "github.com/yosida95/golang-sshkey"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	pub, err := sshkey.UnmarshalPublicKey(string(b))
	if err != nil {
		panic(err)
	}

	fingerprint, err := sshkey.Fingerprint(pub, crypto.SHA256)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d SHA256:%s %s (%s)\n", pub.Length(), base64.RawStdEncoding.EncodeToString(fingerprint), pub.Comment(), keyType(pub.Type()))
}

func keyType(keyType sshkey.Type) string {
	switch keyType {
	case sshkey.KEY_DSA:
		return "DSA"
	case sshkey.KEY_RSA:
		return "RSA"
	case sshkey.KEY_ECDSA:
		return "ECDSA"
	default:
		return "UNKNOWN"
	}
}
