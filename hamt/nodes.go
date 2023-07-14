package hamt

const ARRAY_NODE_SIZE = 8

type nodeType[TValue any] interface {
	insert(Key, TValue) nodeType[TValue]
	retrieve(Key) TValue
}

/* Value */

type valueNode[TValue any] struct {
	key   uint8
	value TValue
}

func (valueNode[TValue]) insert(key Key, value TValue) nodeType[TValue] {
	// TODO: Collisions
	// TODO: Expansions
	return valueNode[TValue]{
		key.hash,
		value,
	}
}

func (n valueNode[TValue]) retrieve(key Key) TValue {
	return n.value
}

/* Array */

type arrayNode[TValue any] struct {
	children [ARRAY_NODE_SIZE]nodeType[TValue]
}

func (n arrayNode[TValue]) insert(key Key, value TValue) nodeType[TValue] {
	index := key.hash % ARRAY_NODE_SIZE
	children := n.children
	children[index] = valueNode[TValue]{key.hash, value}
	return arrayNode[TValue]{
		children: children,
	}
}

func (n arrayNode[TValue]) retrieve(key Key) TValue {
	index := key.hash % ARRAY_NODE_SIZE
	return n.children[index].retrieve(key)
}
