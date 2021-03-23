package merkle

type Hasher interface {
	HashLeaf(a Node) []byte
	HashParent(a, b Node) []byte
}
