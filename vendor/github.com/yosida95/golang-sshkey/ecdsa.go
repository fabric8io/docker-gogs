package sshkey

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
)

type ecdsaPublicKey struct {
	pub *ecdsa.PublicKey
	basePublicKey
}

func (k *ecdsaPublicKey) Length() int {
	return k.pub.Curve.Params().BitSize
}

func (k *ecdsaPublicKey) Public() crypto.PublicKey {
	return k.pub
}

func marshalECDSAPublicKey(k PublicKey) (prefix string, content []byte, err error) {
	key, ok := k.Public().(*ecdsa.PublicKey)
	if !ok {
		err = ErrUnsupportedKey
		return
	}

	var cName string
	switch key.Curve.Params().BitSize {
	case 256:
		cName = "nistp256"
	case 384:
		cName = "nistp384"
	case 521:
		cName = "nistp521"
	default:
		err = ErrUnsupportedKey
		return
	}

	prefix = "ecdsa-sha2-" + cName

	buf := bytes.NewBuffer(nil)
	buf.Write(encodeByteSlice([]byte(prefix)))
	buf.Write(encodeByteSlice([]byte(cName)))
	buf.Write(encodeByteSlice(elliptic.Marshal(key.Curve, key.X, key.Y)))
	content = buf.Bytes()

	return
}

func unmarshalECDSAPublicKey(c []byte, comment string) (*ecdsaPublicKey, error) {
	var alg, cName, data []byte

	alg, c = decodeByteSlice(c)
	if alg == nil || !bytes.HasPrefix(alg, []byte("ecdsa-sha2-")) {
		return nil, ErrMalformedKey
	}

	cName, c = decodeByteSlice(c)
	if cName == nil {
		return nil, ErrMalformedKey
	}

	data, c = decodeByteSlice(c)
	if data == nil {
		return nil, ErrMalformedKey
	}

	var curve elliptic.Curve
	switch string(cName) {
	case "nistp256":
		curve = elliptic.P256()
	case "nistp384":
		curve = elliptic.P384()
	case "nistp521":
		curve = elliptic.P521()
	default:
		return nil, ErrUnsupportedKey
	}

	pub := &ecdsa.PublicKey{
		Curve: curve,
	}
	pub.X, pub.Y = elliptic.Unmarshal(curve, data)
	if pub.X == nil {
		return nil, ErrUnsupportedKey
	}

	key := &ecdsaPublicKey{
		pub: pub,
		basePublicKey: basePublicKey{
			keyType: KEY_ECDSA,
			comment: comment,
		},
	}
	return key, nil
}
