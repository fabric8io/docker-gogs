package sshkey

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func Fingerprint(k PublicKey, alg crypto.Hash) ([]byte, error) {
	var (
		c   []byte
		err error
	)
	switch k.Type() {
	case KEY_RSA:
		_, c, err = marshalRSAPublicKey(k)
	case KEY_DSA:
		_, c, err = marshalDSAPublicKey(k)
	case KEY_ECDSA:
		_, c, err = marshalECDSAPublicKey(k)
	default:
		return nil, ErrUnsupportedKey
	}
	if err != nil {
		return nil, err
	}

	var h hash.Hash
	switch alg {
	case crypto.MD5:
		h = md5.New()
	case crypto.SHA1:
		h = sha1.New()
	case crypto.SHA256:
		h = sha256.New()
	case crypto.SHA512:
		h = sha512.New()
	default:
		return nil, ErrUnsupportedKey
	}

	h.Write(c)
	return h.Sum(nil), nil
}

func PrettyFingerprint(k PublicKey, alg crypto.Hash) (string, error) {
	const (
		hextable = "0123456789abcdef"
	)

	fp, err := Fingerprint(k, alg)
	if err != nil {
		return "", err
	}

	dest := make([]byte, len(fp)*3)
	for i, v := range fp {
		dest[i*3] = hextable[v>>4]
		dest[i*3+1] = hextable[v&0xf]
		dest[i*3+2] = ':'
	}

	return string(dest[0 : len(fp)*3-1]), nil
}
