package merkle

type nodeKind int

const (
	leaf nodeKind = iota
	parent
)

type Node struct {
	index  uint64
	parent uint64
	data   []byte
	hash   []byte
	kind   nodeKind
}
