package hamt

type Hashable interface {
	Hash() uint8
	Equals(Hashable) bool
}
