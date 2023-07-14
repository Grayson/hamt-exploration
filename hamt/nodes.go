package hamt

type nodeType interface {
	canInsert() bool
}

type trampolineNode struct {
}

func (trampolineNode) canInsert() { panic("Unimplemented") }

type valueNode[TKey comparable, TValue comparable] struct {
	key   TKey
	value TValue
}

func (valueNode[TKey, TValue]) canInsert() bool {
	return false
}
