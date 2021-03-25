package merkle

import "crypto"

type NodeHasher interface {
	Node() Node
	HashLeaf(data []byte) []byte
	HashParent(left, right []byte) []byte
}

type BLAKE2b512 struct{}

func (b2b BLAKE2b512) Node() Node {
	return &DefaultNode{}
}
func (b2b BLAKE2b512) HashLeaf(data []byte) []byte {
	return crypto.BLAKE2b_512.New().Sum(data)
}
func (b2b BLAKE2b512) HashParent(left, right []byte) []byte {
	return crypto.BLAKE2b_512.New().Sum(append(left, right...))
}
