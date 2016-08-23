package sshkey

import (
	"crypto/rsa"
	"fmt"
)

const (
	marshaledPub = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDMjH3YZMNFG8cnl98t6w6Ca152cnTsWyrZ56WYSYNkEax1grChZB3P4NcxmtqFxrN2wMXuATiqp62cNkj8wAQUIwRgUnqKkkaQTDyLEDVaTZ75RsZIE4vM/YJ5AzmbCIHK8u6YvfM8fIlv4PKzbMHIIcZvuG9ZYQ+ZEKmSIVxIKZNVfUYyoRK6RFPEMjZPGGoOFRBo8sifsJDLDIBLWOgR4Nf2rWuV+ZuySXX9wjsv42iIdp9RVJcjQXHmi7AKVifKfFJwM+6aPiQcAaWnINzvUnqQK5yrWEp5tVH49bFL92UNriT+LTozloILCj5SdqXQ+JbKp/6EobY96bWhkwyZ yosida95@yosida95"
)

func Example() {
	pubkey, err := UnmarshalPublicKey(marshaledPub)
	if err != nil {
		panic(err)
	}
	nativePub := pubkey.Public().(*rsa.PublicKey)

	fmt.Println(pubkey.Type() == KEY_RSA)
	fmt.Println(nativePub.E)
	fmt.Println(pubkey.Length())
	fmt.Println(pubkey.Comment())

	// Output:
	// true
	// 65537
	// 2048
	// yosida95@yosida95
}
