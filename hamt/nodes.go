package hamt

const ARRAY_NODE_SIZE = 8

type nodeType[TValue any] interface {
	insert(Key, TValue) nodeType[TValue]
	retrieve(Key) (TValue, bool)
}

/* Value */

type valueNode[TValue any] struct {
	key      Hashable
	value    TValue
	hasValue bool
}

func (n valueNode[TValue]) insert(key Key, value TValue) nodeType[TValue] {
	if n.hasValue {
		return n.promote(key, value, n.key, n.value)
	}

	return valueNode[TValue]{
		key.item,
		value,
		true,
	}
}

func (n valueNode[TValue]) promote(key Key, value TValue, origKey Hashable, origValue TValue) nodeType[TValue] {
	if key.canMoveToNextLevel() {
		key.moveToNextLevel()
		updatedKey := NewKey(origKey)
		updatedKey.depth = key.depth
		a := arrayNode[TValue]{}.insert(updatedKey, origValue)
		return a.insert(key, value)
	}

	// TODO: Save space!

	// TODO: Collisions
	panic("Need to handle collisions at edges of available space")
}

func (n valueNode[TValue]) retrieve(key Key) (TValue, bool) {
	return n.value, true
}

/* Terminal/Collision */

type terminalNode[TValue any] struct {
	keys   []Hashable
	values []TValue
}

func (t terminalNode[TValue]) insert(key Key, value TValue) nodeType[TValue] {
	keys := append(t.keys, key.item)
	values := append(t.values, value)
	return terminalNode[TValue]{
		keys,
		values,
	}
}

func (t terminalNode[TValue]) retrieve(key Key) (out TValue, ok bool) {
	for idx, k := range t.keys {
		if k.Equals(key.item) {
			return t.values[idx], true
		}
	}

	return
}

/* Array */

type arrayNode[TValue any] struct {
	children [ARRAY_NODE_SIZE]nodeType[TValue]
}

func (n arrayNode[TValue]) insert(key Key, value TValue) nodeType[TValue] {
	index := key.currentMaskedValue() % ARRAY_NODE_SIZE
	children := n.children

	if children[index] == nil {
		children[index] = valueNode[TValue]{}
	}

	children[index] = children[index].insert(key, value)
	return arrayNode[TValue]{
		children: children,
	}
}

func (n arrayNode[TValue]) retrieve(key Key) (out TValue, ok bool) {
	index := key.currentMaskedValue() % ARRAY_NODE_SIZE
	child := n.children[index]

	switch n := child.(type) {
	case arrayNode[TValue]:
		key.moveToNextLevel()
		return n.retrieve(key)
	case valueNode[TValue]:
		return n.retrieve(key)
	default:
		return
	}
}
