package hamt

type Trie[TValue any] struct {
	root nodeType[TValue]
}

func NewTrie[TValue comparable]() Trie[TValue] {
	return Trie[TValue]{
		arrayNode[TValue]{},
	}
}

func (t *Trie[TValue]) Insert(key uint8, value TValue) Trie[TValue] {
	root := t.root.insert(key, value)
	return Trie[TValue]{root}
}

func (t *Trie[TValue]) Retrieve(key uint8) TValue {
	return t.root.retrieve(key)
}
