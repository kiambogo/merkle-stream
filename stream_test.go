package merkle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var blake2bHasher = BLAKE2b512{}

func Test_NewStream_Empty(t *testing.T) {
	t.Parallel()

	stream := NewStream(blake2bHasher, nil, nil)

	assert.Equal(t, stream.Roots(), new([]Node))
}

func Test_NewStream_SetRootsAndNodes(t *testing.T) {
	t.Parallel()

	roots := &[]Node{
		DefaultNode{
			index:  0,
			parent: 1,
			kind:   leaf,
			data:   []byte("hello, i'm node a"),
			hash:   []byte{},
		},
		DefaultNode{
			index:  2,
			parent: 1,
			kind:   leaf,
			data:   []byte("hello, i'm node b"),
			hash:   []byte{},
		},
	}

	stream := NewStream(blake2bHasher, roots, roots)
	assert.Equal(t, stream.Roots(), roots)
	assert.Equal(t, stream.Nodes(), roots)
}
