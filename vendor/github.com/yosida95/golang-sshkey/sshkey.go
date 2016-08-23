package sshkey

import (
	"crypto"
)

type Type int

const (
	KEY_RSA Type = iota
	KEY_DSA
	KEY_ECDSA
)

type PublicKey interface {
	Type() Type
	Public() crypto.PublicKey
	Length() int
	Comment() string
}

type basePublicKey struct {
	keyType Type
	comment string
}

func (k *basePublicKey) Type() Type {
	return k.keyType
}

func (k *basePublicKey) Comment() string {
	return k.comment
}
