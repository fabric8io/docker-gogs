package sshkey

import (
	"crypto"
	"io/ioutil"
	"testing"
)

type prettyFingerPrintCase struct {
	path string
	fp   string
}

var (
	prettyFingerPrintCases = []prettyFingerPrintCase{
		{
			path: "testdata/rsa.pub",
			fp:   "6c:e4:66:c2:56:10:72:a4:fc:da:5c:0d:4b:30:a6:55",
		},
		{
			path: "testdata/dsa.pub",
			fp:   "04:65:87:9f:ae:6d:d6:56:6d:b3:a9:e6:4f:d7:e9:25",
		},
	}
)

func TestPrettyFingerprint(t *testing.T) {
	for _, c := range prettyFingerPrintCases {
		key, err := ioutil.ReadFile(c.path)
		if err != nil {
			panic(err)
		}
		pub, err := UnmarshalPublicKey(string(key))
		if err != nil {
			t.Error(err)
			continue
		}

		ret, err := PrettyFingerprint(pub, crypto.MD5)
		if err != nil {
			t.Error(err)
			continue
		}

		if ret != c.fp {
			t.Fail()
		}
	}
}
