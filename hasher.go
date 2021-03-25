package merkle

import "crypto"

type NodeHasher interface {
	Node() Node
	HashLeaf(node PartialNode) []byte
	HashParent(left, right Node) []byte
}

type BLAKE2b512 struct{}

func (b2b BLAKE2b512) Node() Node {
	return &DefaultNode{}
}
func (b2b BLAKE2b512) HashLeaf(node PartialNode) []byte {
	return crypto.BLAKE2b_512.New().Sum(node.data)
}
func (b2b BLAKE2b512) HashParent(left, right Node) []byte {
	return crypto.BLAKE2b_512.New().Sum(append(left.Hash(), right.Hash()...))
}
