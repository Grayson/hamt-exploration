package hamt

type Trie[TValue comparable] struct {
	root nodeType
}

func NewTrie[TValue comparable]() Trie[TValue] {
	return Trie[TValue]{}
}

func (t *Trie[TValue]) Insert(key uint8, value TValue) Trie[TValue] {
	return Trie[TValue]{
		valueNode[uint8, TValue]{key, value},
	}
}

func (t *Trie[TValue]) Retrieve(key uint8) *TValue {
	switch n := t.root.(type) {
	case trampolineNode:
		panic("Unimplemented")
	case valueNode[uint8, TValue]:
		return &n.value
	default:
		panic("Unknown node type!")
	}
}
