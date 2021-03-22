package merkle

type Hasher interface {
	Leaf() []byte
	Parent() []byte
}

type stream struct {
	roots []uint64
}

func NewStream(hasher Hasher, roots ...uint64) *stream {
	return &stream{
		roots: roots,
	}
}

func (s stream) Roots() []uint64 {
	return s.roots
}
