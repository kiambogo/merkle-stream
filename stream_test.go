package merkle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHasher struct{}

func (th testHasher) Leaf() []byte {
	return []byte("leaf-hash")
}

func (th testHasher) Parent() []byte {
	return []byte("parent-hash")
}

func Test_NewStream(t *testing.T) {
	stream := NewStream(testHasher{})
	assert.Equal(t, stream.Roots(), []uint64(nil))

	stream = NewStream(testHasher{}, 1)
	assert.Equal(t, stream.Roots(), []uint64{1})

	stream = NewStream(testHasher{}, 1, 2, 3, 4, 5)
	assert.Equal(t, stream.Roots(), []uint64{1, 2, 3, 4, 5})
}
