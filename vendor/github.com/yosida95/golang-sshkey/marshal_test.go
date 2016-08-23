package sshkey

import (
	"io/ioutil"
	"strings"
	"testing"
)

var (
	marshalPublicKeyCases = []string{
		"testdata/rsa.pub",
		"testdata/dsa.pub",
		"testdata/ecdsa.pub",
	}
)

func TestMarshalPublicKey(t *testing.T) {
	for _, c := range marshalPublicKeyCases {
		keyB, err := ioutil.ReadFile(c)
		if err != nil {
			panic(err)
		}
		key := strings.TrimSpace(string(keyB))

		pub, err := UnmarshalPublicKey(key)
		if err != nil {
			panic(err)
		}

		ret, err := MarshalPublicKey(pub)
		if err != nil {
			t.Error(err)
			continue
		}

		if ret != key {
			t.Fail()
		}
	}
}
