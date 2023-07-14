package hamt

type Hashable interface {
	Hash() uint8
}

type HashableInt int

func (i HashableInt) Hash() uint8 {
	return uint8(i)
}
