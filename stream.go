package merkle

import (
	"github.com/kiambogo/flat-tree"
)

type stream struct {
	Hasher
	roots  *[]uint64
	nodes  *[]Node
	blocks int
}

func NewStream(hasher Hasher, roots *[]uint64, nodes *[]Node) *stream {
	if roots == nil {
		roots = new([]uint64)
	}
	if nodes == nil {
		nodes = new([]Node)
	}
	return &stream{
		Hasher: hasher,
		roots:  roots,
		nodes:  nodes,
		blocks: 0,
	}
}

func (s stream) Roots() *[]uint64 {
	return s.roots
}

func (s stream) Nodes() *[]Node {
	return s.nodes
}

func (s *stream) Append(data []byte) {
	// build new node, with index, parent, and data
	index := uint64(s.blocks) * 2
	leaf := Node{
		index:  index,
		parent: flattree.Parent(index),
		data:   data,
		kind:   leaf,
	}
	*s.nodes = append(*s.nodes, leaf)

	s.blocks++
	// hash new new node
	// append hashed node to roots, and nodes
	//
	// rehash stuff:
	// take last and second last roots
	// if parents of those roots is NOT the same, done (why not just check offset?)
	// else, recalc has of the parent
	// pop the 2 elems off roots
	// push the new leaf (which is a parent..?) onto roots, and nodes
	//
	//
	// end result:
	// new node added to nodes, and recalc'd parent added to nodes
	// recalc'd parent added to roots
}
