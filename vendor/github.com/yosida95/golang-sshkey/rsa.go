package sshkey

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"encoding/binary"
	"math/big"
)

type rsaPublicKey struct {
	pub *rsa.PublicKey
	basePublicKey
}

func (r *rsaPublicKey) Length() int {
	return r.pub.N.BitLen()
}

func (r *rsaPublicKey) Public() crypto.PublicKey {
	return r.pub
}

func marshalRSAPublicKey(k PublicKey) (prefix string, content []byte, err error) {
	key, ok := k.Public().(*rsa.PublicKey)
	if !ok {
		err = ErrUnsupportedKey
		return
	}

	prefix = "ssh-rsa"

	buf := bytes.NewBuffer(nil)
	buf.Write(encodeByteSlice([]byte(prefix)))

	e := make([]byte, 4)
	binary.BigEndian.PutUint32(e, uint32(key.E))
	buf.Write(encodeByteSlice(bytes.TrimLeft(e, "\x00")))

	buf.Write(encodeByteSlice([]byte{0}, key.N.Bytes()))
	content = buf.Bytes()

	return
}

func unmarshalRSAPublicKey(c []byte, comment string) (*rsaPublicKey, error) {
	var alg, exp, mod []byte

	alg, c = decodeByteSlice(c)
	if alg == nil || string(alg) != "ssh-rsa" {
		return nil, ErrMalformedKey
	}

	exp, c = decodeByteSlice(c)
	if exp == nil {
		return nil, ErrMalformedKey
	}
	if len(exp) < 4 {
		newExp := make([]byte, 4)
		copy(newExp[4-len(exp):4], exp)
		exp = newExp
	}

	mod, _ = decodeByteSlice(c)
	if mod == nil {
		return nil, ErrMalformedKey
	}

	key := &rsaPublicKey{
		pub: &rsa.PublicKey{
			E: int(binary.BigEndian.Uint32(exp)),
			N: new(big.Int).SetBytes(mod),
		},
		basePublicKey: basePublicKey{
			keyType: KEY_RSA,
			comment: comment,
		},
	}
	return key, nil
}
