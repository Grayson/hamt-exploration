package hamt

type nodeType interface {
	hamtNodeType()
}

type trampolineNode struct {
}

func (trampolineNode) hamtNodeType() {}

type valueNode[TKey comparable, TValue comparable] struct {
	key   TKey
	value TValue
}

func (valueNode[TKey, TValue]) hamtNodeType() {}
