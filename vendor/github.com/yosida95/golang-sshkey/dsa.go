package sshkey

import (
	"bytes"
	"crypto"
	"crypto/dsa"
	"math/big"
)

type dsaPublicKey struct {
	pub *dsa.PublicKey
	basePublicKey
}

func (r *dsaPublicKey) Length() int {
	return r.pub.P.BitLen()
}

func (r *dsaPublicKey) Public() crypto.PublicKey {
	return r.pub
}

func marshalDSAPublicKey(k PublicKey) (prefix string, content []byte, err error) {
	key, ok := k.Public().(*dsa.PublicKey)
	if !ok {
		err = ErrUnsupportedKey
		return
	}

	prefix = "ssh-dss"

	buf := bytes.NewBuffer(nil)
	buf.Write(encodeByteSlice([]byte(prefix)))
	buf.Write(encodeByteSlice([]byte{0}, key.Parameters.P.Bytes()))
	buf.Write(encodeByteSlice([]byte{0}, key.Parameters.Q.Bytes()))
	buf.Write(encodeByteSlice(key.Parameters.G.Bytes()))
	buf.Write(encodeByteSlice(key.Y.Bytes()))
	content = buf.Bytes()

	return
}

func unmarshalDSAPublicKey(c []byte, comment string) (*dsaPublicKey, error) {
	var p, q, g, y []byte

	alg, c := decodeByteSlice(c)
	if alg == nil || string(alg) != "ssh-dss" {
		return nil, ErrMalformedKey
	}

	p, c = decodeByteSlice(c)
	if p == nil {
		return nil, ErrMalformedKey
	}

	q, c = decodeByteSlice(c)
	if q == nil {
		return nil, ErrMalformedKey
	}

	g, c = decodeByteSlice(c)
	if g == nil {
		return nil, ErrMalformedKey
	}

	y, c = decodeByteSlice(c)
	if y == nil {
		return nil, ErrMalformedKey
	}

	key := &dsaPublicKey{
		pub: &dsa.PublicKey{
			Parameters: dsa.Parameters{
				P: new(big.Int).SetBytes(p),
				Q: new(big.Int).SetBytes(q),
				G: new(big.Int).SetBytes(g),
			},
			Y: new(big.Int).SetBytes(y),
		},
		basePublicKey: basePublicKey{
			keyType: KEY_DSA,
			comment: comment,
		},
	}
	return key, nil
}
