package merkle

type nodeKind int

const (
	leaf nodeKind = iota
	parent
)

type Node interface {
	Index() uint64
	Parent() uint64
	Data() []byte
	Kind() nodeKind
}

type unhashedNode struct {
	index  uint64
	parent uint64
	data   []byte
	kind   nodeKind
}

func (un unhashedNode) Index() uint64 {
	return un.index
}
func (un unhashedNode) Parent() uint64 {
	return un.parent
}
func (un unhashedNode) Data() []byte {
	return un.data
}
func (un unhashedNode) Kind() nodeKind {
	return un.kind
}

type hashedNode struct {
	unhashedNode
	hash []byte
}

func (hn hashedNode) Index() uint64 {
	return hn.index
}
func (hn hashedNode) Parent() uint64 {
	return hn.parent
}
func (hn hashedNode) Data() []byte {
	return hn.data
}
func (hn hashedNode) Kind() nodeKind {
	return hn.kind
}
