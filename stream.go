package merkle

import (
	"sync"

	flattree "github.com/kiambogo/flat-tree"
)

type stream struct {
	Hasher         // hashing implementation to use when building the merkle tree
	roots  *[]Node // the current set of root nodes in the tree
	nodes  *[]Node // the current set of all nodes in the tree
	blocks int     // size of the tree (might be ok to just len(nodes))
	wg     *sync.Mutex
}

func NewStream(hasher Hasher, roots *[]Node, nodes *[]Node) *stream {
	if roots == nil {
		roots = new([]Node)
	}
	if nodes == nil {
		nodes = new([]Node)
	}
	return &stream{
		Hasher: hasher,
		roots:  roots,
		nodes:  nodes,
		blocks: 0,
		wg:     &sync.Mutex{},
	}
}

func (s stream) Roots() *[]Node {
	return s.roots
}

func (s stream) Nodes() *[]Node {
	return s.nodes
}

func (s *stream) Append(data []byte) {
	s.wg.Lock()
	defer s.wg.Unlock()

	index := uint64(s.blocks * 2)
	leaf := unhashedNode{
		index:  index,
		parent: flattree.Parent(index),
		data:   data,
		kind:   leaf,
	}
	*s.roots = append(*s.roots, leaf)

	s.blocks++

	for len(*s.roots) > 1 {
		numRoots := len(*s.roots)
		left := (*s.roots)[numRoots-2]
		right := (*s.roots)[numRoots-1]

		if left.Parent() != right.Parent() {
			break
		}

		parentHash := s.HashParent(left, right)
		newParent := hashedNode{}

	}

	// hash new new node
	// leaf.hash = s.HashLeaf(leaf)

	// append hashed node to roots, and nodes
	// *s.nodes = append(*s.nodes, leaf)

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
