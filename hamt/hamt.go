package hamt

type Trie[TValue any] struct {
	root nodeType[TValue]
}

func NewTrie[TValue comparable]() Trie[TValue] {
	return Trie[TValue]{
		arrayNode[TValue]{},
	}
}

func (t *Trie[TValue]) Insert(key Hashable, value TValue) Trie[TValue] {
	root := t.root.insert(NewKey(key), value)
	return Trie[TValue]{root}
}

func (t *Trie[TValue]) Retrieve(key Hashable) (TValue, bool) {
	return t.root.retrieve(NewKey(key))
}
