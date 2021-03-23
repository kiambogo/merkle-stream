package merkle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHasher struct{}

func (th testHasher) HashLeaf(node Node) []byte {
	return []byte("leaf-hash")
}

func (th testHasher) HashParent(left, right Node) []byte {
	return []byte("parent-hash")
}

func Test_NewStream(t *testing.T) {
	stream := NewStream(testHasher{}, nil, nil)
	assert.Equal(t, stream.Roots(), new([]uint64))

	stream = NewStream(testHasher{}, &[]uint64{1}, nil)
	assert.Equal(t, stream.Roots(), &[]uint64{1})

	stream = NewStream(testHasher{}, &[]uint64{1, 2, 3, 4, 5}, nil)
	assert.Equal(t, stream.Roots(), &[]uint64{1, 2, 3, 4, 5})
}
